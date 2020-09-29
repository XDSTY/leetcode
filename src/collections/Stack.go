package collections

type Stack struct {
	stack []interface{}
}

func NewStack() Stack {
	return Stack{stack: []interface{}{}}
}

func (stack *Stack) Push(val interface{}) {
	stack.stack = append(stack.stack, val)
}

func (stack *Stack) Pop() interface{} {
	node := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return node
}

func (stack *Stack) Peek() interface{} {
	return stack.stack[len(stack.stack)-1]
}

func (stack *Stack) Empty() bool {
	return len(stack.stack) == 0
}
