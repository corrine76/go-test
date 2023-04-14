package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
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

type NodeConfig struct {
	ExecUserType   int            `json:"exec_user_type"`   // [动作]执行对象
	RoleSelection  []string       `json:"role_selection"`   // [动作]执行角色
	ExecUserFilter ExecUserFilter `json:"exec_user_filter"` // [动作]执行对象筛选
}

type ExecUserFilter struct {
	Logic int                    `json:"logic"` // 筛选逻辑关系
	Items []ExecUserFilterSingle `json:"items"` // 内容
}

type ExecUserFilterSingle struct {
	Type        int      `json:"type"`        // 筛选项类型
	UserFollows []string `json:"userFollows"` // 员工跟进+共享关系
	UserRoles   []string `json:"userRoles"`   // 员工角色
	UserDepts   []int64  `json:"userDepts"`   // 员工部门
}

func main() {
	str := "eyJhZGRfdGFncyI6IFtdLCAiYXBwb2ludF90aW1lIjogeyJkYXRlX3N0ciI6ICIiLCAiZGF5IjogMH0sICJjdXN0b21lcl9maWxlZCI6IG51bGwsICJkZWxheV90aW1lIjogIiIsICJkZWxheV90aW1lX3R5cGUiOiAwLCAiZXhlY191c2VyX3R5cGUiOiA1LCAiZXhlY3V0ZV9kZXB0IjogW10sICJleGVjdXRlX3VzZXIiOiBudWxsLCAiZ3JvdXBPd25lclNlbGVjdGVkIjogMCwgImluZm9ybV92YWx1ZSI6IG51bGwsICJpbmZvcm1fd2F5IjogMSwgImlzX2FsbG93X21hc3NfaGVscGVyIjogZmFsc2UsICJpc19oaWRlX25vdGljZV9ibG9jayI6IHRydWUsICJub2RlX3Rhc2tfaWQiOiAiIiwgIm92ZXJkdWVfY29udGludWUiOiAyLCAib3ZlcmR1ZV90aW1lX3NldCI6ICIwMDcwMDAwMDAiLCAib3ZlcmR1ZV93YXJuIjogIiIsICJyZW1vdmVfdGFncyI6IFtdLCAicm9sZV9zZWxlY3Rpb24iOiBbImZvbGxvd2VyIl0sICJzaGFwZSI6ICJ2dWUtc2hhcGUifQ=="
	fmt.Sprintln(str)
	// 编写base64解码
	// 解码
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(err)
	}
	nodeConfig := &NodeConfig{}
	if err := json.Unmarshal(decodeBytes, nodeConfig); err != nil {
		fmt.Println(err)
	}
	fmt.Println(nodeConfig)
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

func encrypt(str string) string {
	b := []byte(str)
	return base64.StdEncoding.EncodeToString(b)
}

func decrypt(str string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
