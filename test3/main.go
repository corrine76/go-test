package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"time"
)

func main() {
	// 用于监听服务退出
	done := make(chan error, 2)
	// 用于控制服务退出，传入同一个 stop，做到只要有一个服务退出了那么另外一个服务也会随之退出
	stop := make(chan string, 0)
	// for debug
	go func() {
		done <- pprof(stop)
	}()

	// 主服务
	go func() {
		done <- app(stop)
	}()

	// stoped 用于判断当前 stop 的状态
	var stoped bool
	// 这里循环读取 done 这个 channel
	// 只要有一个退出了，我们就关闭 stop channel
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			log.Printf("server exit err: %+v", err)
		}

		if !stoped {
			stoped = true
			myclose(stop)
		}
	}
}

func myclose(stop chan<- string) {
	stop <- "s"
	// close(stop)
}

func app(stop <-chan string) error {
	server("app", stop)

	time.Sleep(10 * time.Second)
	return fmt.Errorf("mock app exit")
}

func pprof(stop <-chan string) error {
	server("pprof", stop)

	time.Sleep(5 * time.Second)
	return fmt.Errorf("mock pprof exit")
}

// 启动一个服务
func server(sign string, stop <-chan string) error {
	go func() {
		<-stop
		log.Printf("server exit: %s ", sign)
	}()
	return nil
}
