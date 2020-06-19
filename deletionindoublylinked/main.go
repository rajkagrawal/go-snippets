package deletionindoublylinked


type Node struct {
	data  int
	Left *Node
	Right *Node
}

type DoublyLinked struct {
	Head *Node
	Tail *Node
}
func NewDoublyLinked() *DoublyLinked{
	return &DoublyLinked{}
}

//This just containse the logic
//TODO : Need to create linked list first.
func main() {
	delete(8, NewDoublyLinked())
}

func delete(i int,dou *DoublyLinked){
		head := dou.Head
		temp := head
		for temp != nil {
			if temp.data == i {
				if temp == dou.Head {
					dou.Head = temp.Right
					dou.Head.Left = nil
				}else if temp == dou.Tail{
					dou.Tail = dou.Tail.Left
					dou.Tail.Left = nil
				}else{
					temp.Right.Left = temp.Left
					temp.Left.Right = temp.Right
				}
			}
		}
}