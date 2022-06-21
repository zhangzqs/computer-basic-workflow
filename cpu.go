package main

type CPU struct {
	acc uint16
	pc  uint16
	ir  uint16

	stopFlag bool
	mem      *Memory
}

func NewCPU(mem *Memory) *CPU {
	return &CPU{
		mem:      mem,
		stopFlag: false,
	}
}

func (r *CPU) load(addr uint16) {
	r.acc = r.mem.Read(addr)
}

func (r *CPU) store(addr uint16) {
	r.mem.Write(addr, r.acc)
}

func (r *CPU) mul(addr uint16) {
	r.acc = r.acc * r.mem.Read(addr)
}

func (r *CPU) add(addr uint16) {
	r.acc = r.acc + r.mem.Read(addr)
}

func (r *CPU) stop(addr uint16) {
	r.stopFlag = true
}

// 取指令
func (r *CPU) fetch() (op, ad uint16) {
	r.ir = r.mem.Read(r.pc)
	return Ins2OpAd(r.ir)
}

// 指令译码并送出控制信号
func (r *CPU) decode(op uint16) func(uint16) {
	// 分析指令
	switch op {
	case LOAD:
		return r.load
	case MUL:
		return r.mul
	case ADD:
		return r.add
	case STORE:
		return r.store
	case STOP:
		return r.stop
	}
	return nil
}

// 执行一次指令
func (r *CPU) runOnce() {
	// 取指令
	op, ad := r.fetch()
	// 分析指令
	ctl := r.decode(op)
	// 执行指令
	ctl(ad)
	// 更新pc
	r.pc++
}

// Run 执行
func (r *CPU) Run() {
	for !r.stopFlag {
		r.runOnce()
	}
}
