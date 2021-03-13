package link

// 1.链表
import (
	"fmt"
	"os"
	"testing"
)

// Element 定义元素类型
type Element int64

// LinkNode 定义链表结构
type LinkNode struct {
	Data Element   // 数据域
	Next *LinkNode // 指针域
}

// Append 链表追加元素
func Append(head *LinkNode, data Element) {
	point := head
	for point.Next != nil {
		point = point.Next
	}
	var node LinkNode  // 创建一个新节点
	point.Next = &node // 创建节点指针复制
	node.Data = data   // 赋值给新节点
}

// Delete 链表删除元素
func Delete(head *LinkNode, index int) bool {
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return false
	}
	point := head
	for i := 0; i < index-1; i++ {
		point = point.Next
	}
	point.Next = point.Next.Next
	return true
}

// Insert 插入位置
func Insert(head *LinkNode, index int, data Element) bool {
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return false
	}
	// 查找到插入位置
	point := head
	for i := 0; i < index-1; i++ {
		point = point.Next
	}
	var node LinkNode
	node.Data = data
	node.Next = point.Next
	point.Next = &node

	return true
}

// GetLength 获取链表长度
func GetLength(head *LinkNode) int {
	point := head.Next
	var length int
	for point.Next != nil {
		length++
		point = point.Next
	}
	return length
}

func SearchNode(head *LinkNode, data Element) {
	point := head
	for point.Next != nil {
		if point.Data == data {
			break
		}
	}
}

// Traverse 遍历链表
func Traverse(head *LinkNode) {
	point := head.Next
	for point.Next != nil {
		fmt.Println(point.Data)
		point = point.Next
	}
	fmt.Println("Traverse OK!")
}

func TestLink(t *testing.T) {
	t.Log("start run test")
	// 头节点
	var head LinkNode = LinkNode{Data: 0, Next: nil}
	head.Data = 0
	// 批量赋值
	var nodeArray []Element
	for i := 0; i < 5; i++ {
		nodeArray = append(nodeArray, Element(i+1+i*100))
		Append(&head, nodeArray[i])
	}
	Traverse(&head)

	// 删除一个元素
	Delete(&head, 3)
	Traverse(&head)
	// 插入一个元素
	Insert(&head, 6, 10010)
	Insert(&head, 6, 10011)
	Traverse(&head)

	fmt.Println("length:", GetLength(&head))
	os.Exit(0)
}
