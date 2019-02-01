package main

type ListNode struct {
	Data int
	Next *ListNode
}
func main() {

	var a ListNode
	var b ListNode
	var c ListNode

	var head *ListNode

	a.Data = 10
	b.Data = 20
	c.Data = 30
	a.Next = &b
	b.Next = &c
	c.Next = nil
	
	head = &a

	for head !=nil {
		println(head.Data)
		head = head.Next
	}

}
