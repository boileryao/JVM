package rtdz

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (thread *Thread) PC() int {
	return thread.pc
}

func (thread *Thread) SetPC(pc int) {
	thread.pc = pc
}

func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.top
}

func (thread *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return NewFrame(thread, maxLocals, maxStack)
}
