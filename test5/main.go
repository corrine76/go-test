package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	cfredis "git.dustess.com/mk-base/redis-driver/redis"
	"github.com/go-redis/redis"
)

// if ret == 'nil' then return 1 else return 0 end
const (
	delLua = `
	local hasKey = redis.call('exists',KEYS[1])
	if	hasKey == 1
	then
		local dict = cjson.decode(ARGV[1])
		local payload = {}
		local count = 0
		for k, v in pairs(dict) do
			table.insert(payload, v)
			count = count + 1
		end
		if count == 0 
		then
			local ret = redis.call('Del', KEYS[1])
			return ret
		else
			local ret = redis.call('HDel', KEYS[1], unpack(payload))
			return ret
		end
	else
		return 0
	end
`
)

// TestA ..
type TestA struct {
	Name string
}

type TestRedis map[string]interface{}

// MarshalBinary 序列化
func (s TestRedis) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalBinary 反序列化
func (s TestRedis) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

func (t TestA) isEmpty() bool {
	return reflect.DeepEqual(t, TestA{})
}

func main() {
	_, err := cfredis.InitClient(&cfredis.Config{
		ClientName: cfredis.MKCache,
		AppName:    "test",
		Addr:       "192.168.0.128:6379",
		Password:   "dustesssetsud",
		DB:         13,
		PoolSize:   5,
	})
	if err != nil {
		panic(err)
	}

	data := TestRedis{}
	// data["a"] = "a"
	// data["b"] = "b"
	fmt.Println(data)

	client := cfredis.Client(cfredis.MKCache)
	if client == nil {
		panic("client error")
	}

	script := redis.NewScript(delLua)
	r, err := script.Run(client, []string{"ly-3"}, data).Result()
	fmt.Println(r, err)
}

func f3() {
	m := make(map[string]string)
	for _, i := range []string{"a", "b", "c"} {
		m[i] = i
	}
	fmt.Println(m)
}

func f2(s string) error {
	fmt.Println("exec f2")
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("you said: ", s)
		fmt.Println("exec f2 gofunc")
	}()
	return nil
}

func f1() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("f1 recover...")
			fmt.Println(err)
		}
	}()
	if true {
		panic("f1 panic...")
	}
	return nil
}
