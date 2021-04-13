<?php

/**
 * Definition for a binary tree node.
 */
class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

/**
 * 根据前序遍历, 中序遍历构建二叉树
 */
class Solution
{
    private $inmap; // 反转中序数组以及关联
    private $preindex = 0; // 前序参数索引
    private $preorder;

    public function buildTree($preorder, $inorder)
    {
        $this->preorder = $preorder;
        $this->inmap = array_flip($inorder);
        return $this->helper(0, count($inorder) - 1);
    }

    private function helper($instart, $inend)
    {
        if ($instart > $inend) {
            return null;
        }

        $nodeval = $this->preorder[$this->preindex];
        $inindex = $this->inmap[$nodeval];
        $node = new TreeNode($nodeval);

        $this->preindex++;
        $node->left = $this->helper($instart, $inindex - 1);
        $node->right = $this->helper($inindex+1, $inend);
        return $node;
    }
}
