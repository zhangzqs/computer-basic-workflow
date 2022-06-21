package main

// Memory 设计一个存储单元为16位的主存储器
type Memory struct {
	data []uint16
}

func NewMemory(size int) *Memory {
	return &Memory{
		data: make([]uint16, size),
	}
}

func (r *Memory) Read(addr uint16) uint16 {
	return r.data[addr]
}

func (r *Memory) Write(addr uint16, data uint16) {
	r.data[addr] = data
}
