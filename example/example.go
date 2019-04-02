package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
	"time"

	_ "net/http/pprof"

	"github.com/locxiang/debugcharts"
)

func dummyCPUUsage() {
	var a uint64
	var t = time.Now()
	for {
		t = time.Now()
		a += uint64(t.Unix())
	}
}

func dummyAllocations() {
	var d []uint64

	for {
		for i := 0; i < 2*1024*1024; i++ {
			d = append(d, 42)
		}
		time.Sleep(time.Second * 10)
		fmt.Println(len(d))
		d = make([]uint64, 0)
		runtime.GC()
		time.Sleep(time.Second * 10)
	}
}

func main() {

	engine := gin.Default()
	debugcharts.Wrapper(engine)

	go dummyAllocations()
	go dummyCPUUsage()
	go func() {
		err := engine.Run(":8080")
		if err != nil {
			log.Panic(err)
		}
	}()
	log.Printf("you can now open http://localhost:8080/debug/charts/ in your browser")
	select {}
}
