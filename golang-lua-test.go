package main

import "github.com/yuin/gopher-lua"

func main() {
	L := lua.NewState()
	defer L.Close()

	// execute a line of lua code in go
	if err := L.DoString(`print("Hello World from go!")`); err != nil {
		panic(err)
	}

	// execute a lua file
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
}
