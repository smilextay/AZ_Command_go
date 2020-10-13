package timer

import "time"

//GetNowTime 返回系统当前时间
func GetNowTime() time.Time {
	return time.Now()
}

//GetCalculateTime 时间计算
func GetCalculateTime(calculateTime time.Time, d string) (time.Time, error) {

	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return calculateTime.Add(duration), nil
}
