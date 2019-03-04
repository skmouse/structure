package main

import (
	"fmt"
)


type Element interface {}


var header *entry  //链表表头
var size int  //链表长度

type entry struct {
	previous *entry
	next     *entry
	element  Element
}

func newEntry(prev,next *entry,e Element) *entry {
	return  &entry{prev,next,e}
}

//初始化header  表头
func NewLinkedList() *entry {
	header = newEntry(nil,nil,nil)
	header.previous =header
	header.next = header
	return header
}

type LinkedList interface {
	add(e  Element)
	addFirst(e  Element)
	addLast(e Element)
	get(location int) Element
	getFirst() Element
	getLast() Element
	remove(location int) Element
	removeFirst()  Element
	removeLast() Element
	size() int
	removeAll() bool
	getAll() []Element
}

//将元素添加到链表的末尾
func (e *entry) add(element Element)  {
	addBefore(header,element)
}

func (e *entry)addFirst(element Element)  {
	addBefore(header.next,element)
}

func (e *entry) addLast(element Element)  {
	addBefore(header,element)
}

func (e *entry) get(location int) Element {
	entry := getEntryByIndex(location)
	if entry != nil {
		return entry.element
	}
	return nil
}

func (e *entry) getFirst() Element {
	if size == 0 {
		fmt.Println("chian is empty!")
		return nil
	}
	return header.next.element
}

func (e *entry) getLast() Element {
	if size == 0 {
		fmt.Println("chain is empty!")
		return nil
	}
	return header.previous.element
}

func (e *entry) remove(location int) Element {
	if location < 0 || location >= size {
		fmt.Println("check index error")
		return nil
	}
	rmEntry := getEntryByIndex(location)
	element := rmEntry.element
	rmEntry.previous.next = rmEntry.next
	rmEntry.next.previous = rmEntry.previous
	rmEntry = nil
	size--
	return element
}

func (e *entry) removeFirst()  Element{
	return e.remove(0)
}

func (e *entry) removeLast() Element {
	return e.remove(size-1)
}

func (e *entry) size() int {
	return size
}

func (e *entry)removeAll() bool {
	entry := header.next
	for entry != header {
		next := entry.next
		entry.next = nil
		entry.previous = nil
		entry.element = nil
		entry = next
	}
	header.next = header
	header.previous = header
	size = 0
	return true
}

func (e *entry)getAll() []Element {
	resultList := make([]Element,0,size)
	entry := header
	for i:= 0 ; i < size ;i++ {
		resultList = append(resultList,entry.next.element)
		entry = entry.next
	}
	return resultList
}

//在entry节点之前添加
func addBefore(e *entry,element Element) Element{
	newEntry := newEntry(e.previous,e,element)
	newEntry.previous.next = newEntry
	newEntry.next.previous = newEntry
	size++
	return newEntry
}

func getEntryByIndex(index int) *entry {
	if index < 0 || index >= size {
		fmt.Println("check index error")
		return nil
	}
	e := header

	if index < (size >> 1) {
		for i:=0 ; i <= index ;i++ {
			e = e.next
		}
	}else {
		for i:= size ; i > index; i-- {
			e = e.previous
		}
	}
	return e

}

func main() {
	linkedList := NewLinkedList()
	for i:=0 ;i<50;i++ {
		linkedList.add(i)
	}
	//fmt.Println(size)
	//fmt.Println(linkedList.get(40))
	fmt.Println(len(linkedList.getAll()),linkedList.getAll())
	//linkedList.addFirst(100)
	//fmt.Println(len(linkedList.getAll()),linkedList.getAll())
	//linkedList.addLast(200)
	//fmt.Println(len(linkedList.getAll()),linkedList.getAll())
	//linkedList.add(1000)
	//fmt.Println(len(linkedList.getAll()),linkedList.getAll())
	//fmt.Println(linkedList.get(50))
	//fmt.Println(linkedList.getFirst())
	//fmt.Println(linkedList.getLast())
	//
	//fmt.Println(linkedList.removeFirst())
	//fmt.Println(len(linkedList.getAll()),linkedList.getAll())
	//fmt.Println(linkedList.removeLast())
	//fmt.Println(len(linkedList.getAll()),linkedList.getAll())
	//fmt.Println(linkedList.remove(50))
	//fmt.Println(len(linkedList.getAll()),linkedList.getAll())
	//fmt.Println(linkedList.removeAll())
	//fmt.Println(len(linkedList.getAll()),linkedList.getAll())
	//
	//t := time.Now()
	//for i:=0;i<400000;i++ {
	//	linkedList.add(i)
	//}
	//fmt.Println(time.Since(t))
	//fmt.Println(linkedList.size())
	//
	//t = time.Now()
	//fmt.Println(linkedList.get(200000))
}
