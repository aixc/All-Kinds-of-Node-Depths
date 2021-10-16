// Solution 04

using namespace std;

class BinaryTree
{
public:
    int value;
    BinaryTree *left;
    BinaryTree *right;
    BinaryTree(int value)
    {
        this->value = value;
        left = NULL;
        right = NULL;
    }
};

struct TreeInfo
{
    int numNodesInTree;
    int sumOfDepths;
    int sumOfAllDepths;
};

TreeInfo getTreeInfo(BinaryTree *tree);

// Average case: when the tree is balanced
// O(n) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree
int allKindsOfNodeDepths(BinaryTree *root)
{
    return getTreeInfo(root).sumOfAllDepths;
}

TreeInfo getTreeInfo(BinaryTree *tree)
{
    if (tree == NULL)
    {
        return TreeInfo{0, 0, 0};
    }

    TreeInfo leftTreeInfo = getTreeInfo(tree->left);
    TreeInfo rightTreeInfo = getTreeInfo(tree->right);
    
    int sumOfLeftDepths = leftTreeInfo.sumOfDepths + leftTreeInfo.numNodesInTree;
    int sumOfRightDepths =
        rightTreeInfo.sumOfDepths + rightTreeInfo.numNodesInTree;
    
    int numNodesInTree =
        1 + leftTreeInfo.numNodesInTree + rightTreeInfo.numNodesInTree;
    int sumOfDepths = sumOfLeftDepths + sumOfRightDepths;
    int sumOfAllDepths =
        sumOfDepths + leftTreeInfo.sumOfAllDepths + rightTreeInfo.sumOfAllDepths;
    
    return TreeInfo{
        numNodesInTree,
        sumOfDepths,
        sumOfAllDepths,
    };
}
