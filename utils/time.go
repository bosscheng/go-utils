package utils

import "time"

func NowWithMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
