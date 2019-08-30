package time

import "time"

//type Time time.Time
type Time struct {
	time.Time
}

var (
	timeJSONFormat = "2006-01-02 15:04:05"
)

func SetTimeJSONFormat(format string) {
	timeJSONFormat = format
}

func GetTimeJSONFormat() string {
	return timeJSONFormat
}

func Now() Time {
	return Time{Time: time.Now()}
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeJSONFormat+`"`, string(data), time.Local)
	t.Time = now
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeJSONFormat)+2)
	b = append(b, '"')
	b = t.Time.AppendFormat(b, timeJSONFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) String() string {
	return t.Time.Format(timeJSONFormat)
}
