package timer

import "time"

// GetNowTime 获取时间
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// GetCalculateTime 推算时间
func GetCalculateTime(currTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currTimer.Add(duration), nil
}
