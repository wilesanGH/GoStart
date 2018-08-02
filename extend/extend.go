package extend

import "GoStart/tree"

type MyTreeNode struct {
	Node *tree.TreeNode
}

func (myNode *MyTreeNode) PostOrder(){
	if myNode ==nil ||myNode.Node == nil{
		return
	}
	left:=MyTreeNode{myNode.Node.Left}
	left.PostOrder()
	right:=MyTreeNode{myNode.Node.Rigth}
	right.PostOrder()
	myNode.Node.Print()
}
