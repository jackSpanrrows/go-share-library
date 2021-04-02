package common

import "time"

// 获取格式化时间
func GetTimeFormat(stamp uint32, fmt string) (string, error) {
	var utc time.Time
	if stamp == 0 {
		utc = time.Now().UTC()
	} else {
		utc = time.Unix(int64(stamp), 0).UTC()
	}

	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return "", err
	}

	ntime := utc.In(loc)
	return ntime.Format(fmt), nil
}
