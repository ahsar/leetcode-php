<?php

/**
 * Definition for a binary tree node.
 */
class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;

    function __construct($value) {
        $this->val = $value;
    }
}

class BinTree {
    public $root;

    public function __construct($v)
    {
        $this->root = new TreeNode($v);
    }

    public function insert(int $v)
    {
        $p = $this->root;

        while ($p) {
            // 插入值大于当前值
            if ($v > $p->val) {
                if ($p->right == null) {
                    $p->right = new TreeNode($v);
                    break;
                }
                $p = $p->right;
            } else {
                if ($p->left == null) {
                    $p->left = new TreeNode($v);
                    break;
                }
                $p = $p->left;
            }
        }
    }

    public function find(int $v)
    {
        $p = $this->root;

        while ($p) {
            if ($p->val == $v) {
                return $p;
            }
            if ($v > $p->val) {
                $p = $p->right;
            }
            if ($v < $p->val) {
                $p = $p->left;
            }
        }
        return null;
    }

    public function delete(int $v)
    {
        $p = $this->root;
        $pNode = null;
        while ($p) {
            if ($p->val > $v) {
                $pNode = $p;
                $p = $p->left;
            }
            if ($p->val < $v) {
                $pNode = $p;
                $p = $p->right;
            }

            if ($p->val == $v) {
                break;
            }
        }

        // 找到需要被删除节点
        if (!$p) {
            return null;
        }

        // 被删除节点没有任何子节点
        if (!$p->left && !$p->right) {
            unset($p);
            if ($pNode->left->val == $v) {
                $pNode->left = null;
            } else {
                $pNode->right = null;
            }
            return true;
        }

        // 有两个子节点
        if ($p->left && $p->right) {
            $minP = $p->right;
            $minPP = $p;

            while ($minP->left) {
                $minPP = $minP;
                $minP = $minP->left; 
            }
            $minPP->left = null;
            $minP->left = $p->left;
            $minP->right = $p->right;
            if ($pNode->left->val == $v) {
                $pNode->left = $minP;
            } else {
                $pNode->left = $minP;
            }
            unset($p);
            print_r($this->root);
            return true;
        }

        // 有一个子节点
        if ($pNode->left->val == $v) {
            $pNode->left->val = $pNode->left ? : $p->right;
        } else {
            $pNode->right->val = $pNode->left ? : $p->right;
        }
    }

    /**
     * preOrder
     *
     * 先序遍历(根左右)
     * 递归版
     * @param TreeNode $p
     * @access public
     * @return void
     */
    public static function preOrder($p)
    {
        if (!$p) {
            return;
        }
        echo $p->val;
        echo '->';
        self::preOrder($p->left);
        self::preOrder($p->right);
        return;
    }

    /**
     * inOrder2
     *
     * 中序遍历
     * 迭代版本
     * @access public
     * @return void
     */
    public function inOrder2()
    {
        $p = $this->root;
        $stack = [];

        while ($stack || $p) {
            if ($p) {
                $stack[] = $p;
                $p = $p->left;
            } else {
                $p = array_pop($stack);
                print_r($p->val);
                echo "->";
                $p = $p->right;
            }
        }
    }

    /**
     * preOrder2
     *
     * 先序遍历非递归
     * @access public
     * @return void
     */
    public function preOrder2()
    {
        $p = $this->root;
        $stack = [];

        while ($p || $stack) {
            if ($p) {
                print_r($p->val);
                echo '->';
                $stack[] = $p;
                $p = $p->left;
            } else {
                $p = array_pop($stack);
                $p = $p->right;
            }
        }
    }

    /**
     * postOrder
     *
     * 后序遍历
     * @access public
     * @return void
     */
    public function postOrder()
    {
        $p = $this->root;
        $stack = [];
        $lastVisit = null;

        while ($p) {
            $stack[] = $p;
            $p = $p->left;
        }

        while ($stack) {
            $p = array_pop($stack);
            // 没有右边节点或者当前节点右子节点已经被遍历过
            if ($p->right == null || $lastVisit == $p->right) {
                print_r($p->val);
                echo '->';
                $lastVisit = $p;
            } else {
                $stack[] = $p;
                $p = $p->right;
                while ($p) {
                    $stack[] = $p;
                    $p = $p->left;
                }
            }
        }
    }

    /**
     * postOrder2
     *
     * 后序遍历
     * 双栈实现
     * @access public
     * @return void
     */
    public function postOrder2()
    {
        $p = $this->root;
        $s1 = $s2 = [];

        $s1[] = $p;
        while ($s1) {
            $node = array_pop($s1);
            $s2[] = $node;
            if ($node->left) {
                $s1[] = $node->left;
            }
            if ($node->right) {
                $s1[] = $node->right;
            }
        }

        while ($s2) {
            $node = array_pop($s2);
            print_r($node->val);
            echo '->';
        }
    }

    public function levelOrder()
    {
        $p = $this->root;
        $queue = [
            $p
        ];

        while ($queue) {
            $node = array_shift($queue);
            echo $node->val;
            echo '->';
            $node->left && $queue[] = $node->left;
            $node->right && $queue[] = $node->right;
        }
    }
}

$root = (new BinTree(7));
$root->insert(3);
$root->insert(5);
$root->insert(6);
$root->insert(1);
$root->insert(13);
$root->insert(17);
$root->insert(4);
$root->insert(8);
$root->insert(9);
//$root->inOrder2();
//$root->preOrder2();
//$root->postOrder();
$r = $root->delete(3);
var_dump($r);
//$root->postOrder2();
//BinTree::preOrder($root->root);
//print_r($root);
