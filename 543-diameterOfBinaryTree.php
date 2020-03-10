<?php
require_once __DIR__ . '/structure/bintree.php';
/**
 * 二叉树的直径
 */

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {
    private $max = 0;

    /**
     * @param TreeNode $root
     * @return Integer
     */
    public function diameterOfBinaryTree($root)
    {
        $this->maxPath($root);
        return $this->max;
    }

    public function maxPath($root)
    {
        if (!$root->left && !$root->right) {
            return 0;
        }

        $left = $root->left ? $this->maxPath($root->left) + 1 : 0;
        $right = $root->right ? $this->maxPath($root->right) + 1 : 0;
        $this->max = max($this->max, $left + $right);
        return max($left, $right);
    }
}

$root = (new BinTree(7));
$root->insert(8);
$root->insert(5);
$root->insert(6);
$root->insert(3);

$r = (new Solution())->maxPath($root->root);
print_r($r);
die;
