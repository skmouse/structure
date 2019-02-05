/**
* 动态链表
*/
package main

import "fmt"

type List struct {
	data int
	next *List
}
func main() {
	var head List
	head.data = 1
	var nodeArray [] int
	for i:=0; i<10; i++  {
		nodeArray = append(nodeArray, int(i*100))
		add(&head, nodeArray[i])
	}

	get(&head)
	search(&head,900)
}
/**
* 插入节点
 */
func add(head *List, data int)  {
	point :=head
	for point.next != nil {
		point = point.next
	}
	var node List
	point.next = &node
	point.data = data
}
func GetLength (head *List) int {
	point := head
	var length int
	for point.next != nil {
		length ++
		point = point.next
	}
	return length
}
func search(head *List, data int)  {
	point :=head
	index :=0
	for head.next !=nil {
		if point.data ==data {
			fmt.Println(data,"exist", index,"th")
			break
		} else {
			index ++
			point = point.next
			if index > GetLength(head)-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
}
/**
* 获取所有节点
*/
func get(head *List)  {
	point := head
	for point.next !=nil {
		println(point.data)
		point = point.next
	}
}
