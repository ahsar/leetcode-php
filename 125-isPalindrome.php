<?php
/**
 * 验证回文字符串
 * @attention 0P while作为条件
 */

class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function isPalindrome($s) {
        $s = preg_replace( '/[\W]/', '', $s);
        $s = strtolower($s);

        $len = strlen($s);
        $mid = intval($len / 2);

        $left = --$mid;
        $right = $mid + 1;
        if ($len % 2) {
            $right++;
        }

        while (isset($s[$left]) && isset($s[$right])) {
            if ($s[$left] != $s[$right]) {
                return false;
            }
            $left--;
            $right++;
        }
        return true;
    }
}
//$r = (new Solution)->isPalindrome('A man, a plan, a canal: Panama');
//$r = (new Solution)->isPalindrome('siis');
//$r = (new Solution)->isPalindrome('race a car');
$r = (new Solution)->isPalindrome('0P');
var_dump($r);
