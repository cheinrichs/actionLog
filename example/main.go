package main

import (
	"fmt"
	"os"

	actionlog "github.com/cheinrichs/actionlog"
)

func main() {

	argsWithoutProg := os.Args[1:]

	for i := 0; i < len(argsWithoutProg); i++ {
		actionlog.AddAction(argsWithoutProg[i])
	}

	fmt.Println(actionlog.GetStats())
}
