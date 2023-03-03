package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	target = `{"cid": "W00000006880", "cusId": "86391903-93f8-11ec-b6b0-2ecfd471f0ae", "state": "ECTLNCoUz4KKEnG5hbo9q5_qrmark"}
	{"cid": "W00000005178", "cusId": "02be130f-0c94-11ec-a046-8e55fcb864b9", "state": "Dq3Awm5kibSkKX7Af6UNKd_qrmark"}
	{"cid": "W00000011011", "cusId": "c220dd7f-93f8-11ec-a8ca-c6e4718230e2", "state": "sefZ76vcFUnG6pDQeav3UU_qrmark"}
	{"cid": "W00000002698", "cusId": "fbe7364d-f027-11eb-a83a-9aa7228f8d96", "state": "wb6Jvt9x7xY9wsyv8TVLUA_qrmark"}
	{"cid": "W00000001946", "cusId": "2d2fb046-93f9-11ec-8c88-be4c61d53b0d", "state": "C2ewuNguSs8PUw4ifvTfTJ_qrmark"}
	{"cid": "W00000012817", "cusId": "3542130f-93f9-11ec-b6b0-2ecfd471f0ae", "state": "ZsNwaro6qfdbjnKztgQv5B_qrmark"}`
)

func main() {
	test1()
}

func test4() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Errorf("%+v and T[%T]", err, err))
		}
	}()

	var i = 1
	var j = 0
	k := i / j
	fmt.Printf("%d / %d = %d\n", i, j, k)
}

func test3() {
	// callback
	callbackname := "./batchcallback.txt"
	callbackfile, err := os.Open(callbackname)
	if err != nil {
		panic(err)
	}
	defer callbackfile.Close()
	callbackbuf := bufio.NewReader(callbackfile)
	// callbackMap := make(map[string]string)
	for {
		data, _, err := callbackbuf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// 处理数据
		parseData := make(map[string]string)
		err = json.Unmarshal(data, &parseData)
		if err != nil {
			panic(err)
		}
		fmt.Println(parseData)
		// fmt.Println(parseData["cid"], parseData["cusId"], parseData["state"])
	}

	// // create
	// createname := "./batchcallback.txt"
	// createfile, err := os.Open(createname)
	// if err != nil {
	// 	panic(err)
	// }
	// defer callbackfile.Close()
	// createbuf := bufio.NewReader(createfile)

}

func test2() {
	list := strings.Split(target, "\n")
	for _, line := range list {
		line2 := strings.TrimSpace(line)
		parseData := make(map[string]string)
		err := json.Unmarshal([]byte(line2), &parseData)
		if err != nil {
			panic(err)
		}
		fmt.Println(parseData["cid"], parseData["cusId"], parseData["state"])
	}
}

func test1() {
	// 文本解析
	rname := "./test2.json"
	rfile, err := os.Open(rname)
	if err != nil {
		panic(err)
	}
	defer rfile.Close()
	rbuf := bufio.NewReader(rfile)
	for {
		data, _, err := rbuf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// 处理数据
		parseData := make(map[string]string)
		err = json.Unmarshal(data, &parseData)
		if err != nil {
			panic(err)
		}
		fmt.Println(parseData)
		// fmt.Println(parseData["cid"], parseData["cusId"], parseData["state"])
	}
}
