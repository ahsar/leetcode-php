<?php

/**
 * 判断是否子树
 */
class Solution
{
    public function isSubStructure($A, $B)
    {
        if ($A == null || $B == null) {
            return false;
        }
        return $this->isSub($A, $B) || $this->isSubStructure($A->left, $B) || $this->isSubStructure($A->right, $B);
    }

    private function isSub($A, $B)
    {
        if ($B == null) {
            return true; 
        }
        if ($A == null || $A->val != $B->val) {
            return false;
        }

        return $this->isSub($A->left, $B->left) && $this->isSub($B->right, $B->right);
    }
}
