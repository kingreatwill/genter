package strings

import (
	"strings"
	"unicode"
)

func Count(s, substr string) int {
	return strings.Count(s, substr)
}

func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

func ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

func IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}

func IndexRune(s string, r rune) int {
	return strings.IndexRune(s, r)
}

func IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}
func LastIndexAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}
func LastIndexByte(s string, c byte) int {
	return strings.LastIndexByte(s, c)
}
func SplitN(s, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}
func SplitAfterN(s, sep string, n int) []string {
	return strings.SplitAfterN(s, sep, n)
}
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}
func SplitAfter(s, sep string) []string {
	return strings.SplitAfter(s, sep)
}

func Fields(s string) []string {
	return strings.Fields(s)
}
func FieldsFunc(s string, f func(rune) bool) []string {
	return strings.FieldsFunc(s, f)
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func Map(mapping func(rune) rune, s string) string {
	return strings.Map(mapping, s)
}

func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToTitle(s string) string {
	return strings.ToTitle(s)
}

func ToUpperSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToUpperSpecial(c, s)
}

func ToLowerSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToLowerSpecial(c, s)
}

func ToTitleSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToTitleSpecial(c, s)
}

func ToValidUTF8(s, replacement string) string {
	return strings.ToValidUTF8(s, replacement)
}

func Title(s string) string {
	return strings.Title(s)
}

func TrimLeftFunc(s string, f func(rune) bool) string {
	return strings.TrimLeftFunc(s, f)
}

func TrimRightFunc(s string, f func(rune) bool) string {
	return strings.TrimRightFunc(s, f)
}

func TrimFunc(s string, f func(rune) bool) string {
	return strings.TrimFunc(s, f)
}

func IndexFunc(s string, f func(rune) bool) int {
	return strings.IndexFunc(s, f)
}

func LastIndexFunc(s string, f func(rune) bool) int {
	return strings.LastIndexFunc(s, f)
}

func Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

func TrimLeft(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

func TrimRight(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

func Index(s, substr string) int {
	return strings.Index(s, substr)
}
