package main

import (
	"fmt"
	"time"

	"github.com/ananrafs/goerranger/goerranger"
	"github.com/ananrafs/goerranger/zord/wopool"
)

func main() {
	goer, dispose := goerranger.Init(wopool.New, goerranger.Options{Count: 2})
	defer dispose()
	for i := 0; i < 10; i++ {
		goer.Hit(func() {
			fmt.Println("hello guys")
			time.Sleep(2 * time.Second)
		})
	}

}
