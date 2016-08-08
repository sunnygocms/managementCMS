package util

import (
	"errors"

	"github.com/sunnygocms/managementCMS/models"
)

//节点结构
type Node struct {
	Data models.SunnyNavigation
	Next *Node
}

var Chain Node

func ChainClear() {
	var mynode Node
	mynode.Data = models.SunnyNavigation{}
	mynode.Next = nil
	Chain = mynode
}

/*
* 返回第一个节点
* h 头结点
 */
func GetFirst() *Node {
	h := &Chain
	if h.Next == nil {
		return nil
	}
	return h.Next
}

/*
* 返回最后一个节点
* h 头结点
 */
func GetLast() *Node {
	h := &Chain
	if h.Next == nil {
		return nil
	}
	i := h
	for i.Next != nil {
		i = i.Next
		if i.Next == nil {
			return i
		}
	}
	return nil
}

//取长度
func GetChainLength() int {
	var i int = 0
	n := &Chain
	for n.Next != nil {
		//		if n.Data.Name != "" {
		//			fmt.Println("^^^^^^^^", n.Data.Id, n.Data.Level, n.Data.ParentId, n.Data.Name)
		//		}

		i++
		n = n.Next
	}
	//	if n.Data.Name != "" {
	//		fmt.Println("^^^^^^^^", n.Data.Id, n.Data.Level, n.Data.ParentId, n.Data.Name)
	//	}
	return i
}

//把链表换成数组
func ChainToSunnyNavigation() (ml []models.SunnyNavigation, err error) {
	if GetChainLength() == 0 {
		err = errors.New("length is nil.")
	} else {
		n := &Chain
		for n.Next != nil {
			if n.Data.Name != "" {
				ml = append(ml, n.Data)
				//				fmt.Println("^^^^^^^^", n.Data.Id, n.Data.Level, n.Data.ParentId, n.Data.Name)
			}
			n = n.Next
		}
		if n.Data.Name != "" {
			ml = append(ml, n.Data)
			//			fmt.Println("^^^^^^^^", n.Data.Id, n.Data.Level, n.Data.ParentId, n.Data.Name)
		}
		err = errors.New("")
	}
	return

}

//插入一个节点
//h: 头结点
//d:要插入的节点
//p:要插入的位置
func Insert(p models.SunnyNavigation) (result bool) {
	node := &Chain
	if node.Next == nil && GetChainLength() == 0 {
		var mynode Node
		mynode.Data = p
		mynode.Next = nil
		node.Next = &mynode
		result = true
	} else {
		result = false
		for node.Next != nil {
			if node.Data.Id == p.ParentId {
				//				fmt.Println("---------", p, node.Data.ParentId, node.Data.Name)
				var mynode Node
				mynode.Data = p
				mynode.Next = node.Next
				node.Next = &mynode
				result = true
				break
			}
			node = node.Next
			if node.Next == nil {
				var mynode Node
				mynode.Data = p
				mynode.Next = nil
				node.Next = &mynode
				result = true
				break
			}
		}
	}

	return
}
