package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

type tidbMysql struct {
	db *sql.DB
}

func main() {
	// a := substrByByte("如何做“洞察业务”的财务经营分析？", 512, "点击阅读")
	// fmt.Println(a)
	// if test, err := initDB(); err == nil {
	// 	test.Create()
	// 	test.Close()
	// }
}

func substrByByte(str string, length int, def string) string {
	bss := []byte(str)
	if l := len(bss); length > l {
		length = l
	}
	bs := bss[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return def
}

func initDB() (*tidbMysql, error) {
	driver := &tidbMysql{}
	db, err := sql.Open("mysql", "root:guo6RjqO4OBlinv2VXAt@tcp(10.100.0.12:4000)/biz_dev?charset=utf8")
	if err != nil {
		fmt.Println("database initialize error: ", err.Error())
		return nil, err
	}
	driver.db = db
	fmt.Println("init tidb success")
	return driver, nil
}

func (driver *tidbMysql) Create() {
	if driver.db == nil {
		return
	}
	fmt.Println("start to create")
	stmt, err := driver.db.Prepare("INSERT INTO biz_dev.qw_customer_follow_user (_id, cid, company_id, description, external_userid, friend_state, uid, userid) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("db prepare failed", err.Error())
		return
	}
	defer stmt.Close()
	// 循环插入数据
	insertcount := 0
	for i := 0; i < 100000; i++ {
		// 构建数据
		id := driver.randStringRunes(32) // 生成随机数
		cusid := "8fe32228-95bf-11ea-a971-0242ac110004"
		cid := "W00000000109"
		desc := "testforlock"
		extid := "wmDolHEAAAglWgLqpXaGAgJ92u0p6GvA"
		fstate := "IsFriend"
		uid := "420"
		userid := "mitchconner"
		if result, err := stmt.Exec(id, cusid, cid, desc, extid, fstate, uid, userid); err == nil {
			if _, err := result.LastInsertId(); err == nil {
				insertcount++
			}
		}
	}
	fmt.Println("insert count : ", insertcount)
}

func (driver *tidbMysql) Close() {
	if driver.db != nil {
		driver.db.Close()
	}
	fmt.Println("tidb close success")
}

func (driver *tidbMysql) randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
