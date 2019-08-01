package main


/*
二叉树是每个结点最多有两个子树的树结构。通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）。二叉树常被用于实现二叉查找树和二叉堆。
一棵深度为k，且有2^k-1个节点的二叉树，称为满二叉树。这种树的特点是每一层上的节点数都是最大节点数。而在一棵二叉树中，除最后一层外，若其余层都是满的，并且最后一层或者是满的，或者是在右边缺少连续若干节点，则此二叉树为完全二叉树。具有n个节点的完全二叉树的深度为floor(log2n)+1。深度为k的完全二叉树，至少有2k-1个叶子节点，至多有2k-1个节点。
*/
func main() {

	
}


type Node struct {
	data int
	parent *Node
	left *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

// 向二叉排序树节点添加一个节点
func (tree *BinaryTree)InsertItem(i int){
	if tree.root == nil {
		tree.root = &Node{data:i}
		return
	}

	currentNode := tree.root

	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data:i,parent:currentNode}
				return
			}
			currentNode = currentNode.right
		}else{
			if currentNode.left == nil {
				currentNode.left = &Node{data:i,parent:currentNode}
				return
			}
			currentNode = currentNode.left
		}
	}
}


// 二叉树中某个位置的节点
func (tree *BinaryTree) SearchItem(i int)(*Node , bool){
	if tree.root == nil {
		return nil,false
	}

	currentNode := tree.root

	for currentNode != nil {
		if i == currentNode.data {
			return  currentNode,true
		}else if i > currentNode.data {
			currentNode = currentNode.right
		}else if i < currentNode.data{
			currentNode = currentNode.left
		}
	}
	return nil,false
}

// 先序遍历 先访问根，再遍历左子树，再遍历右子树

func (tree *BinaryTree) PreorderTraversal(subtree *Node,callback func(int)){
	callback(subtree.data)
	if subtree.left != nil {
		tree.PreorderTraversal(subtree.left,callback)
	}
	if subtree.right != nil {
		tree.PreorderTraversal(subtree.right,callback)
	}
}

//  中序遍历 先遍历左子树，再访问根，再遍历右子树
func (tree *BinaryTree) InorderTraversal(subtree *Node, callback func(int)) {
	if subtree.left != nil {
		tree.InorderTraversal(subtree.left, callback)
	}
	callback(subtree.data)
	if subtree.right != nil {
		tree.InorderTraversal(subtree.right, callback)
	}
}

// 后序遍历 先遍历左子树，再遍历右子树，再访问根

func (tree *BinaryTree) PostorderTraversal(subtree *Node, callback func(int)) {
	if subtree.left != nil {
		tree.PostorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PostorderTraversal(subtree.right, callback)
	}
	callback(subtree.data)
}




