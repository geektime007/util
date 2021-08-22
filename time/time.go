package time

import "time"

// TimeCost 耗时统计
func TimeCost() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}
