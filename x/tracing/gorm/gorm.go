package gorm

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	ParentSpanGormKey = "opentracingParentSpan"
	SpanGormKey       = "opentracingSpan"
	SpanGormTracer    = "opentracingTracer"
)

func Regist(db *gorm.DB) {
	callbacks := NewCallbacks(db)
	callbacks.RegisterCallbacks("create")
	callbacks.RegisterCallbacks("update")
	callbacks.RegisterCallbacks("delete")
	callbacks.RegisterCallbacks("query")
	callbacks.RegisterCallbacks("row_query")
}

func SetContext(db *gorm.DB, ctx context.Context) *gorm.DB {
	if ctx == nil {
		return db
	} else {
		parentSpan := opentracing.SpanFromContext(ctx)
		return db.Set(ParentSpanGormKey, parentSpan)
	}
}

type callbacks struct {
	db *gorm.DB
}

func NewCallbacks(db *gorm.DB) *callbacks {
	return &callbacks{
		db: db,
	}
}

func (c *callbacks) beforeCreate(scope *gorm.Scope)   { c.before(scope) }
func (c *callbacks) afterCreate(scope *gorm.Scope)    { c.after(scope, "INSERT") }
func (c *callbacks) beforeQuery(scope *gorm.Scope)    { c.before(scope) }
func (c *callbacks) afterQuery(scope *gorm.Scope)     { c.after(scope, "SELECT") }
func (c *callbacks) beforeUpdate(scope *gorm.Scope)   { c.before(scope) }
func (c *callbacks) afterUpdate(scope *gorm.Scope)    { c.after(scope, "UPDATE") }
func (c *callbacks) beforeDelete(scope *gorm.Scope)   { c.before(scope) }
func (c *callbacks) afterDelete(scope *gorm.Scope)    { c.after(scope, "DELETE") }
func (c *callbacks) beforeRowQuery(scope *gorm.Scope) { c.before(scope) }
func (c *callbacks) afterRowQuery(scope *gorm.Scope)  { c.after(scope, "") }

func (c *callbacks) before(scope *gorm.Scope) {
	parentSpanValue, ok := scope.Get(ParentSpanGormKey)
	if !ok {
		return
	}
	if parentSpanValue == nil {
		tr := opentracing.GlobalTracer()
		span := tr.StartSpan("sql")
		ext.DBType.Set(span, "sql")
		scope.Set(SpanGormKey, span)
	} else {
		parentSpan := parentSpanValue.(opentracing.Span)
		tr := opentracing.GlobalTracer()
		span := tr.StartSpan("sql", opentracing.ChildOf(parentSpan.Context()))
		ext.DBType.Set(span, "sql")
		scope.Set(SpanGormKey, span)
	}
}

func (c *callbacks) after(scope *gorm.Scope, operation string) {
	val, ok := scope.Get(SpanGormKey)
	if !ok || val == nil {
		return
	}
	span := val.(opentracing.Span)
	if operation == "" {
		operation = strings.ToUpper(strings.Split(scope.SQL, " ")[0])
	}
	span.SetTag("db.table", scope.TableName())
	span.SetTag("db.method", operation)
	span.LogFields(log.String("db.sql", formatSQL(scope)))
	span.LogFields(log.Object("db.pool.stats", c.db.DB().Stats()))
	if scope.HasError() {
		ext.Error.Set(span, true)
		span.LogFields(log.String("db.sql.error", scope.DB().Error.Error()))
	}
	span.Finish()
}

func (c *callbacks) RegisterCallbacks(name string) {
	beforeName := fmt.Sprintf("tracing:%v_before", name)
	afterName := fmt.Sprintf("tracing:%v_after", name)
	gormCallbackName := fmt.Sprintf("gorm:%v", name)
	// gorm does some magic, if you pass CallbackProcessor here - nothing works
	switch name {
	case "create":
		c.db.Callback().Create().Before(gormCallbackName).Register(beforeName, c.beforeCreate)
		c.db.Callback().Create().After(gormCallbackName).Register(afterName, c.afterCreate)
	case "update":
		c.db.Callback().Update().Before(gormCallbackName).Register(beforeName, c.beforeUpdate)
		c.db.Callback().Update().After(gormCallbackName).Register(afterName, c.afterUpdate)
	case "delete":
		c.db.Callback().Delete().Before(gormCallbackName).Register(beforeName, c.beforeDelete)
		c.db.Callback().Delete().After(gormCallbackName).Register(afterName, c.afterDelete)
	case "query":
		c.db.Callback().Query().Before(gormCallbackName).Register(beforeName, c.beforeQuery)
		c.db.Callback().Query().After(gormCallbackName).Register(afterName, c.afterQuery)
	case "row_query":
		c.db.Callback().RowQuery().Before(gormCallbackName).Register(beforeName, c.beforeRowQuery)
		c.db.Callback().RowQuery().After(gormCallbackName).Register(afterName, c.afterRowQuery)
	}
	//c.db.Callback().Register("gorm:commit_or_rollback_transaction", commitOrRollbackTransactionCallback)
	//c.db.Callback().Create().Register("gorm:begin_transaction", beginTransactionCallback)
}

func formatSQL(scope *gorm.Scope) string {
	var (
		sql             string
		formattedValues []string
	)
	var (
		sqlRegexp                = regexp.MustCompile(`\?`)
		numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
	)
	for _, value := range scope.SQLVars {
		indirectValue := reflect.Indirect(reflect.ValueOf(value))
		if indirectValue.IsValid() {
			value = indirectValue.Interface()
			if t, ok := value.(time.Time); ok {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
			} else if b, ok := value.([]byte); ok {
				if str := string(b); isPrintable(str) {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
				} else {
					formattedValues = append(formattedValues, "'<binary>'")
				}
			} else if r, ok := value.(driver.Valuer); ok {
				if value, err := r.Value(); err == nil && value != nil {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			} else {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
			}
		} else {
			formattedValues = append(formattedValues, "NULL")
		}
	}

	// differentiate between $n placeholders or else treat like ?
	if numericPlaceHolderRegexp.MatchString(scope.SQL) {
		sql = scope.SQL
		for index, value := range formattedValues {
			placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
			sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
		}
	} else {
		formattedValuesLength := len(formattedValues)
		for index, value := range sqlRegexp.Split(scope.SQL, -1) {
			sql += value
			if index < formattedValuesLength {
				sql += formattedValues[index]
			}
		}
	}
	return sql
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
