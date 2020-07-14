package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	// add()
	low()
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	fmt.Println("low is start")
	EvChan := hook.Start()
	defer hook.End()
	// 可以监听所有的事件
	for ev := range EvChan {
		fmt.Println("hook: ", ev)
	}
}
