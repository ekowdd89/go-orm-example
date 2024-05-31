package main

import (
	cmd "github.com/ekowdd89/go-orm-example/cmd"
)

func main() {
	svr, _ := cmd.New()
	svr.Run()
}
