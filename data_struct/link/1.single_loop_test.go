package link

// 单向循环链表
import (
	"fmt"
	"os"
	"testing"
)

// SingleLoopLinkNode 定义链表结构
type SingleLoopLinkNode struct {
	Data int                 // 数据域
	Next *SingleLoopLinkNode // 指针域
}

// Traverse 遍历链表
func traverseLoop(head *SingleLoopLinkNode) {
	point := head
	if point.Next == head {
		fmt.Println("...")
		fmt.Println(point.Data)
	} else {
		for point.Next != head {
			point = point.Next
			fmt.Println(point.Data)
		}
	}
	fmt.Println("Traverse OK!")
}

func (l *SingleLoopLinkNode) append(data int) {
	if l.Next == nil {
		l.Data = data
		l.Next = l
	}
	temp := l
	for temp.Next != l {
		temp = temp.Next
	}
	// 创建一个新节点
	node := new(SingleLoopLinkNode)
	node.Data = data
	node.Next = l
	//  最后一个节点的值
	temp.Next = node
}

func TestLoopLink(t *testing.T) {
	loopNode := new(SingleLoopLinkNode)
	loopNode.append(1)
	loopNode.append(2)
	loopNode.append(3)
	loopNode.append(4)
	traverseLoop(loopNode)
	os.Exit(0)
}
