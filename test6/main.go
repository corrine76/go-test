package main

import (
	"context"
	"fmt"
	"git.dustess.com/mk-base/util/date"
	"sort"
	"time"
)

// intSlice ..
type intSlice []int64

func (s intSlice) Len() int           { return len(s) }
func (s intSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s intSlice) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	fmt.Println(getNearNoticeTime([]string{"16:08"}))
}

func MeasureTime(ctx context.Context, funcName string, extraData ...interface{}) func() {
	start := time.Now()
	return func() {
		dur := time.Since(start)
		fmt.Println(dur.String())
	}
}

func getNearNoticeTime(times []string) int64 {
	now := date.GetNowTimestampms()

	if len(times) == 0 {
		return 0
	}

	ti := make([]int64, 0, len(times))
	for _, t := range times {
		ti = append(ti, getTimeDay(t, 0, 0, 0))
	}
	sort.Sort(intSlice(ti))

	// 获取最后一个或者唯一一个数据与当前时间对比 如果小于当前时间，即说明今日提醒时间已过，需要设置第二天的第一个时间
	if ti[len(times)-1] < now {
		return ti[0] + 24*60*60*1000
	}

	for _, tt := range ti {
		if now > tt {
			continue
		}
		return tt
	}
	return 0
}

// getTimeDay 获取传入时间在当前天的时间戳 y 年份偏移量 m 月份偏移量 d 日期偏移量
func getTimeDay(hm string, y, m, d int) int64 {
	now := time.Now().AddDate(y, m, d)
	ns := date.GetDateFormat(now)

	ns = fmt.Sprintf("%s %s", ns, hm)
	t, _ := time.ParseInLocation("2006-01-02 15:04", ns, time.Local)
	return t.Unix() * 1e3
}
