// Solution 03

package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// Average case: when the tree is balanced
// O(n) time | O(n) space - where n is the number of nodes in the Binary Tree
func AllKindsOfNodeDepths(root *BinaryTree) int {
	nodeCounts, nodeDepths := map[*BinaryTree]int{}, map[*BinaryTree]int{}
	addNodeCounts(root, nodeCounts)
	addNodeDepths(root, nodeDepths, nodeCounts)
	return sumAllNodeDepths(root, nodeDepths)
}

func sumAllNodeDepths(node *BinaryTree, nodeDepths map[*BinaryTree]int) int {
	if node == nil {
		return 0
	}
	return sumAllNodeDepths(node.Left, nodeDepths) + sumAllNodeDepths(node.Right, nodeDepths) + nodeDepths[node]
}

func addNodeDepths(node *BinaryTree, nodeDepths, nodeCounts map[*BinaryTree]int) {
	nodeDepths[node] = 0
	if node.Left != nil {
		addNodeDepths(node.Left, nodeDepths, nodeCounts)
		nodeDepths[node] = nodeDepths[node] + nodeDepths[node.Left] + nodeCounts[node.Left]
	}
	if node.Right != nil {
		addNodeDepths(node.Right, nodeDepths, nodeCounts)
		nodeDepths[node] = nodeDepths[node] + nodeDepths[node.Right] + nodeCounts[node.Right]
	}
}

func addNodeCounts(node *BinaryTree, nodeCounts map[*BinaryTree]int) {
	nodeCounts[node] = 1
	if node.Left != nil {
		addNodeCounts(node.Left, nodeCounts)
		nodeCounts[node] = nodeCounts[node] + nodeCounts[node.Left]
	}
	if node.Right != nil {
		addNodeCounts(node.Right, nodeCounts)
		nodeCounts[node] = nodeCounts[node] + nodeCounts[node.Right]
	}
}
