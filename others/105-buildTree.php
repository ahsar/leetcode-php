<?php
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
//class Solution {

    /**
     * @param Integer[] $preorder
     * @param Integer[] $inorder
     * @return TreeNode
     */
    //function buildTree($preorder, $inorder) {
        //if (empty($preorder)) {
            //return null;
        //}
        //$preRootValue = $preorder[0];
        //$root = new TreeNode($preRootValue);
        //$inRootIndex = array_search($preRootValue, $inorder);
        //// 用于构造左子树的中序遍历序列
        //$leftInOrder = array_slice($inorder, 0, $inRootIndex);
        //// 用于构造右子树的中序遍历序列
        //$rightInOrder = array_slice($inorder, $inRootIndex+1);
        //// 用于构造左子树的先序遍历序列
        //$leftPreOrder = array_slice($preorder, 1, $inRootIndex);
        //// 用于构造右子树的先序遍历序列
        //$rightPreOrder = array_slice($preorder, $inRootIndex+1);
        //// 构造左子树
        //$root->left = $this->buildTree($leftPreOrder, $leftInOrder);
        //// 构造右子树
        //$root->right = $this->buildTree($rightPreOrder, $rightInOrder);
        //return $root;
    //}
//}

class Solution {

    /**
     * @param Integer[] $preorder
     * @param Integer[] $inorder
     * @return TreeNode
     */
    function buildTree($preorder, $inorder) {
        $this->preorder = $preorder;
        $this->inorder = $inorder;
        $this->inmap = array_flip($inorder);
        return $this->helper(0,count($inorder)-1);
    }
    private $inmap; //反转中序数组中所有键以及它们关联的值
    private $preindex = 0; // 前序参数索引
    private $preorder;
    function helper($instart,$inend){
        if($instart > $inend) return null;
        $nodeval = $this->preorder[$this->preindex];
        $inindex = $this->inmap[$nodeval];
        $node = new TreeNode($nodeval);
        $this->preindex++;
        $node->left = $this->helper($instart,$inindex-1);
        $node->right = $this->helper($inindex+1,$inend);
        return $node;
    }
}
