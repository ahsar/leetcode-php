<?php

/*
 * Definition for a binary tree node.
 */
class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($val = 0, $left = null, $right = null) {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
    }
}

class Solution {

    /**
     * @param Integer[] $preorder
     * @param Integer[] $inorder
     * @return TreeNode
     */
    public function buildTree($preorder, $inorder)
    {
        if (!$preorder) {
            return null;
        }

        $root = new TreeNode($preorder[0]);
        foreach ($inorder as $i => $v) {
            if ($v == $preorder[0])
                break;
        }
        print_r($i);
        die;

        $root->left = $this->buildTree(array_slice($preorder, 1, $i), array_slice($inorder, 0, $i));
        $root->right = $this->buildTree(array_slice($preorder, $i+1), array_slice($inorder, $i+1));
        return $root;
    }
}

$pre = [3,9,20,15,7]; // 根->左->右
$in = [9,3,15,20,7];  // 左->根->右

$r = (new Solution())->buildTree($pre, $in);
