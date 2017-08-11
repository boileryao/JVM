package rtdz

type Stack struct {
	maxSize uint
	size    uint
	top     *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.maxSize {
		panic("StackOverflowError, stack of runtime data zone is already full")
	}

	frame.lower = stack.top
	stack.top = frame
	stack.size++
}

func (stack *Stack) pop() *Frame {
	if stack.top == nil {
		panic("EmptyStackPopError, can't pop an empty stack")
	}

	top := stack.top

	stack.top = top.lower
	top.lower = nil

	stack.size--
	return stack.top
}

func (stack *Stack) peek() *Frame {
	return stack.top
}
