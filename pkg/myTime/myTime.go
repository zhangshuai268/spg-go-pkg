package myTime

import "time"

const DateTimeFormat string = "2006-01-02 15:04:05"

// MyTime 自定义时间类型
// 可直接与time.Time类型相互强制类型转换
type MyTime time.Time

func (l MyTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, DateTimeFormat)
	b = append(b, '"')
	return b, nil
}
func (l *MyTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+DateTimeFormat+`"`, string(b), time.Local)
	*l = MyTime(now)
	return err
}

// StringToMyTime 字符串转MyTime
// timeS: 日期字符串，格式: YYYY-MM-DD hh:mm:ss或YYYY-MM-DD
func StringToMyTime(timeS string) (MyTime, error) {
	location, err := time.ParseInLocation("2006-01-02 15:04:05", timeS, time.Local)
	if err != nil {
		return MyTime{}, err
	}
	return MyTime(location), nil
}
