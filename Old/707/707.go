package main


type MyLinkedList struct {
	head *Node
}
type Node struct {
	Val int
	Pre *Node
	Next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	curr:=this.head
	for i := 0; i < index&&curr!=nil; i++ {
		curr=curr.Next
	}
	if curr!=nil{
		return curr.Val
	}else {
		return -1
	}
}


func (this *MyLinkedList) AddAtHead(val int)  {
	newNode:=&Node{Val: val}
	newNode.Next=this.head
	if this.head!=nil{
		this.head.Pre=newNode
	}
	this.head=newNode
}


func (this *MyLinkedList) AddAtTail(val int)  {
	newNode:=&Node{Val: val,}
	if this.head==nil{
		this.AddAtHead(val)
		return
	}
	curr:=this.head
	for curr!=nil&&curr.Next!=nil{
		curr=curr.Next
	}
	newNode.Pre=curr
	curr.Next=newNode
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index==0{
		this.AddAtHead(val)
	}else{
		newNode:=&Node{Val: val,}
		curr:=this.head
		for i := 0; i < index-1&&curr!=nil; i++ {
			curr=curr.Next
		}
		if curr!=nil{
			newNode.Pre=curr
			newNode.Next=curr.Next
			curr.Next=newNode
			if newNode.Next!=nil{
				newNode.Next.Pre=newNode
			}
		}
	}
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index==0{
		this.head=this.head.Next
		if this.head != nil {
			this.head.Pre = nil
		}
	}else {
		curr:=this.head
		for i := 0; i < index-1&&curr!=nil; i++ {
			curr=curr.Next
		}
		if curr!=nil&&curr.Next!=nil{
			curr.Next=curr.Next.Next
			if curr.Next!=nil{
				curr.Next.Pre=curr
			}
		}
	}

}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {

}
