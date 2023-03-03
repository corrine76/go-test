package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Millisecond ..
const (
	Millisecond int64 = 1
	Second            = 1000 * Millisecond
	Minute            = 60 * Second
	Hour              = 60 * Minute
	Day               = 24 * Hour
)

type node struct {
	ID       string
	ParentID string
	Path     string
	FilterID string
	Type     int8
	Desc     string
}

type treeNode struct {
	ID          string
	ParentID    string
	Path        string
	FilterID    string
	FilterMerge []string
	Type        int8
	Desc        string
	PreNode     *treeNode
	Children    []*treeNode
}

type treeNodeV2 struct {
	ID             string
	PID            string
	Path           string
	FilterID       string
	Type           int8
	Desc           string
	Children       []*treeNodeV2 // 主体层级结构
	ChildrenRecord []string      // 全部子节点集
}

type msgDataAsyncCallbackInfo struct {
	Type         string `json:"type"`         // 消息类型
	TaskDetailID int64  `json:"taskDetailID"` // 任务详情id
}

type NodeConfig struct {
	ExecUserType  int8     `json:"exec_user_type"`   // [动作]执行对象
	RoleSelection []string `json:"role_selection"`   // [动作]执行角色
	InformWay     int8     `json:"inform_way"`       // [动作]任务通知方式
	InformValue   []string `json:"inform_value"`     // [动作]自定义任务通知时间 [hh:mm, hh:mm] 循环提醒时间
	OverdueTime   string   `json:"overdue_time_set"` // [动作]逾期时间设置 9位字符串111222333 代表:111天222小时333分钟
	OverdueWarn   string   `json:"overdue_warn"`     // [动作]逾期提醒 9位字符串111222333 代表:111天222小时333分钟
	DelayTime     string   `json:"delay_time"`       // [等待]配置等待时间 9位字符串111222333 代表:111天222小时333分钟
}

// 伪代码
func main() {
	var str = "yJhZGRfdGFncyI6IFtdLCAiYXBwb2ludF90aW1lIjogeyJkYXRlX3N0ciI6ICIiLCAiZGF5IjogMH0sICJjdXN0b21lcl9maWxlZCI6IG51bGwsICJkZWxheV90aW1lIjogIiIsICJkZWxheV90aW1lX3R5cGUiOiAwLCAiZXhlY191c2VyX3R5cGUiOiA1LCAiZXhlY3V0ZV9kZXB0IjogW10sICJleGVjdXRlX3VzZXIiOiBudWxsLCAiaW5mb3JtX3ZhbHVlIjogbnVsbCwgImluZm9ybV93YXkiOiAxLCAiaXNfYWxsb3dfbWFzc19oZWxwZXIiOiBmYWxzZSwgImlzX2hpZGVfbm90aWNlX2Jsb2NrIjogdHJ1ZSwgIm5vZGVfdGFza19pZCI6ICIiLCAib3ZlcmR1ZV9jb250aW51ZSI6IDEsICJvdmVyZHVlX3RpbWVfc2V0IjogIjAwMTAwMDAwMCIsICJvdmVyZHVlX3dhcm4iOiAiIiwgInJlbW92ZV90YWdzIjogW10sICJyb2xlX3NlbGVjdGlvbiI6IFsiZm9sbG93ZXIiLCAic2hhcmVPd25lciJdLCAic2hhcGUiOiAidnVlLXNoYXBlIn0="

	config := &NodeConfig{}
	if err := json.Unmarshal([]byte(str), config); err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(config)
}

func transTimeFormat2Int(delayTime string) (int64, error) {
	repairStr := "000000000"
	strLen := 9
	// 向前补充缺失位数 "0"
	if len(delayTime) < strLen {
		delayTime = string([]byte(repairStr)[:strLen-len(delayTime)]) + delayTime
	}
	// 截取后9位
	if len(delayTime) > strLen {
		delayTime = string([]byte(repairStr)[strLen:])
	}
	d, err := strconv.ParseInt(string([]byte(delayTime)[:3]), 10, 64)
	if err != nil {
		return 0, err
	}
	h, err := strconv.ParseInt(string([]byte(delayTime)[3:6]), 10, 64)
	if err != nil {
		return 0, err
	}
	m, err := strconv.ParseInt(string([]byte(delayTime)[6:]), 10, 64)
	if err != nil {
		return 0, err
	}
	return d*Day + h*Hour + m*Minute, nil
}

func nodeRange(node *treeNode) {
	fmt.Printf("id[%s]\tpid[%s]\t类型[%d]\t条件合并%v\t\t描述[%s]\n", node.ID, node.ParentID, node.Type, node.FilterMerge, node.Desc)
	for _, v := range node.Children {
		nodeRange(v)
	}
}

func nodeRangeV2(node *treeNodeV2) {
	fmt.Printf("id=%-4s描述=%-16s子节点集合=%s\n", node.ID, node.Desc, node.ChildrenRecord)
	for _, v := range node.Children {
		nodeRangeV2(v)
	}
}

// 构造树 多个父节点
func toTree() {
	nodeList := []node{
		{ID: "1", ParentID: "", Path: "01", Type: 1, Desc: "根节点触发sop", FilterID: "f1"},
		{ID: "2", ParentID: "1", Path: "0101", Type: 2, Desc: "金额", FilterID: "f2,f3"},
		{ID: "3", ParentID: "2", Path: "010101", Type: 3, Desc: "发送优惠券3"},
		{ID: "4", ParentID: "2,7", Path: "010102", Type: 2, Desc: "性别", FilterID: "f4,f5"},
		{ID: "5", ParentID: "4", Path: "01010201", Type: 5, Desc: "延迟10分钟"},
		{ID: "6", ParentID: "4", Path: "01010202", Type: 3, Desc: "发放优惠券2"},
		{ID: "7", ParentID: "4", Path: "01010203", Type: 5, Desc: "等待10分钟 重新判断"},
		{ID: "8", ParentID: "5", Path: "0101020101", Type: 3, Desc: "发放优惠券1"},
	}

	// 来自同一个plan的node节点，按照path有序排列，构造tree结构
	var treePath2Node = make(map[string]*treeNodeV2, 0)
	var id2Node = make(map[string]*treeNodeV2, 0)
	var rootTreeNode *treeNodeV2
	for _, nodeItem := range nodeList {
		tNode := &treeNodeV2{
			ID:       nodeItem.ID,
			PID:      nodeItem.ParentID,
			Path:     nodeItem.Path,
			Type:     nodeItem.Type,
			Desc:     nodeItem.Desc,
			Children: []*treeNodeV2{},
		}
		// 记录map
		treePath2Node[tNode.Path] = tNode
		id2Node[tNode.ID] = tNode
	}

	for _, nodeItem := range nodeList {
		tNode := id2Node[nodeItem.ID]
		// 根节点
		if rootTreeNode == nil {
			rootTreeNode = tNode // 记录根节点
			continue
		}
		// 通过path拼接结构父节点
		pathLen := len(tNode.Path)
		parentPath := tNode.Path[0 : pathLen-2]
		parentNode, ok := treePath2Node[parentPath]
		if !ok {
			continue
		}
		parentNode.Children = append(parentNode.Children, tNode)

		// 通过pid更新关系父节点
		pids := strings.Split(tNode.PID, ",")
		for _, p := range pids {
			pNode, ok := id2Node[p]
			if !ok {
				continue
			}
			pNode.ChildrenRecord = append(pNode.ChildrenRecord, tNode.ID)
		}
	}
	nodeRangeV2(rootTreeNode)
}

// 构造树+条件合并 单个父节点
func filterMerge() {
	nodeList := []node{
		{ID: "1", ParentID: "", Path: "01", Type: 1, Desc: "根节点触发sop", FilterID: "f1"},
		{ID: "2", ParentID: "1", Path: "0101", Type: 2, Desc: "金额", FilterID: "f2,f3"},
		{ID: "3", ParentID: "2", Path: "010101", Type: 3, Desc: "发送优惠券3"},
		{ID: "4", ParentID: "2", Path: "010102", Type: 2, Desc: "性别", FilterID: "f4,f5"},
		{ID: "5", ParentID: "4", Path: "01010201", Type: 5, Desc: "延迟10分钟"},
		{ID: "6", ParentID: "4", Path: "01010202", Type: 3, Desc: "发放优惠券2"},
		{ID: "7", ParentID: "5", Path: "0101020101", Type: 3, Desc: "发放优惠券1"},
	}

	// 来自同一个plan的node节点，按照path有序排列，构造tree结构
	var path2Node = make(map[string]*treeNode, 0)
	var rootTreeNode *treeNode
	for _, nodeItem := range nodeList {
		tNode := &treeNode{
			ID:          nodeItem.ID,
			ParentID:    nodeItem.ParentID,
			Path:        nodeItem.Path,
			Type:        nodeItem.Type,
			Desc:        nodeItem.Desc,
			FilterID:    nodeItem.FilterID,
			FilterMerge: []string{},
			Children:    []*treeNode{},
		}
		// 记录map
		path2Node[tNode.Path] = tNode
		// 根节点
		if rootTreeNode == nil {
			tNode.FilterMerge = append(tNode.FilterMerge, tNode.FilterID)
			rootTreeNode = tNode // 记录根节点
			continue
		}
		// 通过map查找父节点
		pathLen := len(tNode.Path)
		parentPath := tNode.Path[0 : pathLen-2]
		leftPath := tNode.Path[pathLen-2:]
		index, err := strconv.Atoi(leftPath)
		if err != nil {
			fmt.Println(err, "bug!!!!")
		}
		parentNode, ok := path2Node[parentPath]
		if !ok {
			continue
		}
		// 添加到父节点children
		parentNode.Children = append(parentNode.Children, tNode)
		// 更新父节点
		tNode.PreNode = parentNode

		// 条件合并
		// 根据情况merge
		if parentNode.Type == 2 {
			filters := strings.Split(parentNode.FilterID, ",")
			tNode.FilterMerge = append(parentNode.FilterMerge, filters[index-1])
		} else {
			tNode.FilterMerge = append(parentNode.FilterMerge, tNode.FilterMerge...)
		}
		if tNode.Type > 2 {
			// todo 请求分层合并filter条件，更新node的filterID
		}
	}
	nodeRange(rootTreeNode)
}
