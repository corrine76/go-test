package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Obj struct {
	FilterID string `json:"_id"`
}

func DataParse() {
	// 1. 读取JSON文件
	data, err := ioutil.ReadFile("./funcs/data_parse.json")
	if err != nil {
		log.Fatalf("读取文件失败：%v", err)
	}

	// 2. 解析JSON对象数组
	var objs []Obj
	err = json.Unmarshal(data, &objs)
	if err != nil {
		log.Fatalf("解析JSON失败：%v", err)
	}

	// 3. 遍历对象数组，生成map
	objMap := make(map[string]string)
	for _, p := range objs {
		objMap[p.FilterID] = p.FilterID
	}

	filterIDs := make([]string, 0)
	for _, v := range objMap {
		filterIDs = append(filterIDs, v)
	}
	fmt.Println(strings.Join(filterIDs, ","))
}
