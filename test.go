package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	var head Node
	var nodeArray [] int
	for i:=0; i<10;i++ {
		nodeArray = append(nodeArray, i*10)
		adder(&head, nodeArray[i])
	}

	gets(&head)
	length(&head)
	searcher(&head, 80)
	insert(&head,3,10086)
	gets(&head)
	delete(&head, 3)
	gets(&head)
}
/**
 * 插入节点
 */
func adder(head *Node, data int) {
	point :=head
	for point.next!=nil {
		point = point.next
	}
	var node Node
	point.data = data
	point.next = &node
}

/**
  获取所有的结点
 */
func gets(head *Node)  {
	point :=head
	for point.next !=nil {
		println(point.data)
		point = point.next
	}
}

/**
 * 获取结点长度
 */
func length(head *Node)  {
	point :=head
	i:=0
	for point.next!=nil {
		point = point.next
		i++
	}
	println(i)
}

/**
 * 查找结点
 */
func searcher(head *Node, data int)  {
	point:= head
	index:=0
	for point.next!=nil {
		if point.data == data {
			fmt.Println(data,"exist", index,"th")
			break
		}else {
			point = point.next
			index ++
		}
	}

}

/**
 * 在某个结点中插入节点
 */
func insert(head *Node, index int, data int) {
	point :=head
	for i:=0; i<index-1 ;i++  {
		point = point.next
	}
	var node Node
	node.data = data
	node.next = point.next
	point.next = &node
}

/**
* 删除指定index位置的元素
 */
func delete(head *Node, index int)   {
	point :=head
	for i:=0; i<index-1 ;i++  {
		point = point.next
	}
	point.next = point.next.next
}



