package main

import "fmt"

func main() {
	mem := NewMemory(10)
	mem.Write(5, 2)
	mem.Write(6, 3)
	mem.Write(7, 1)
	mem.Write(8, 0)

	mem.Write(0, OpAd2Ins(LOAD, 5))
	mem.Write(1, OpAd2Ins(MUL, 6))
	mem.Write(2, OpAd2Ins(ADD, 7))
	mem.Write(3, OpAd2Ins(STORE, 8))
	mem.Write(4, OpAd2Ins(STOP, 0))

	cpu := NewCPU(mem)
	cpu.Run()
	fmt.Println(mem.data)
}
