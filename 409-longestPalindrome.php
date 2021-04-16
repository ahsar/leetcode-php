<?php
/**
 *  最长回文字符串
 */

class Solution {
    /**
     * @param String $s
     * @return Integer
     */
    public function longestPalindrome($s) {
        $hash = [];
        $len = strlen($s);

        for ($i = 0; $i < $len; $i++) {
            if (!isset($hash[$s[$i]])) {
                $hash[$s[$i]] = 1;
            } else {
                ++$hash[$s[$i]];
            }
        }

        $sum = 0;
        $flag = 0;
        foreach($hash as $v){
            // 奇数
            if($v % 2){
                $flag = 1;
                $sum += $v - 1;
            } else {
                $sum += $v;
            }
        }

        return $flag == 1 ? $sum + 1 : $sum;
    }
}

$s = 'bananas';
$s = 'abccccdd';
$r = (new Solution())->longestPalindrome($s);
print_r($r);
