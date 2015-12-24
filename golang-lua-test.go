package main

import "github.com/yuin/gopher-lua"

func Double(L *lua.LState) int {
	lv := L.ToInt(1)            /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1                    /* number of results */
}

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

	// --- define a custom go function and call it in lua
	L.SetGlobal("double", L.NewFunction(Double))
	L.DoFile("calldouble.lua")

}
