package main

type MinMaxStack struct {
	sta []int
	max []int
	min []int
}

func NewMinMaxStack() *MinMaxStack {
	// Write your code here.
	return &MinMaxStack{}
}
func (stack *MinMaxStack) IsEmpty() bool {
	if len(stack.sta) == 0 {
		return true
	}
	return false
}
func (stack *MinMaxStack) Peek() int {
	if !stack.IsEmpty(){
		return stack.sta[len(stack.sta)-1]
	}
	return -1
}

func (stack *MinMaxStack) Pop() int {
	if stack.IsEmpty() {
		return  -1
	}
	poppedValue := stack.sta[len(stack.sta)-1]
	stack.sta = stack.sta[:len(stack.sta)-1]
	if len(stack.max)>0 &&  stack.max[len(stack.max)-1] == poppedValue {
		stack.max = stack.max[:len(stack.max)-1]
	}
	if len(stack.min)>0 && stack.min[len(stack.min)-1] == poppedValue {
		stack.min = stack.min[:len(stack.min)-1]
	}


	return poppedValue

}

func (stack *MinMaxStack) Push(number int) {
	stack.sta = append(stack.sta, number)
	if len(stack.max)==0 {
		stack.max = append(stack.max, number)
	} else if stack.max[len(stack.max)-1] <= number {
		stack.max = append(stack.max, number)
	}

	if len(stack.min)==0 {
		stack.min = append(stack.min, number)
	} else if  stack.min[len(stack.min)-1] >= number {
		stack.min = append(stack.min, number)
	}

}

func (stack *MinMaxStack) GetMin() int {
	if len(stack.min) == 0 {
		return -1
	}
	return stack.min[len(stack.min)-1]

}

func (stack *MinMaxStack) GetMax() int {
	if len(stack.max) == 0 {
		return -1
	}
	return stack.max[len(stack.max)-1]
}

//This is just a dummy main , test cases cover the required functionalities
func main()  {
	stack := NewMinMaxStack()
	stack.Push(5)
}