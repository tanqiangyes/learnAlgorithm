//顺序结构实现得线性表
package list

import "errors"

const  MAXSIZE = 20

type Node struct {
	data [MAXSIZE]int
	lens int
}

//获取线性表元素
// n Node结构
// i 获取位置
func GetNodeData(n Node,i int) (res int,err error)  {
	if n.lens == 0 || i > n.lens {
		return 0,errors.New("i error")
	}
	res = n.data[i-1]
	return res,nil
}

//插入线性表元素
// n Node结构
// i 插入位置
// d 插入数据
func InsertNode(n *Node,i int,d int) error {
	if n.lens >= MAXSIZE {//线性表已满
		return errors.New("list full")
	}
	if i > n.lens + 1 || i < 1 {//插入得位置有误
		return errors.New("i error")
	}
	if i <= n.lens {
		for k := n.lens - 1;k>= i - 1 ; k-- {
			n.data[k+1] = n.data[k]
		}
	}
	n.data[i-1] = d
	n.lens++
	return nil
}

//删除线性表元素
// n Node结构
// i 删除位置
func DeleteNode(n *Node,i int) (res int,err error) {
	if n.lens == 0 {
		return 0,errors.New("empty list")
	}
	if i > n.lens || i < 1 {
		return 0,errors.New("i error")
	}
	res = n.data[i]
	if i < n.lens {
		for k := i;k < n.lens ; k++ {
			n.data[k-1] = n.data[k]
		}
	}
	n.data[n.lens-1] = 0
	n.lens--
	return res,nil
}

type PtrNode struct {
	data int
	next *PtrNode
}
