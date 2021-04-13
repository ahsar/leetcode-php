<?php

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    public function lengthOfLongestSubstring($s)
    {
        if (!$s) {
            return 0; 
        }

        $l = strlen($s);
        $maxL = 0;
        $win = [];

        // 使用hash
        for ($i = 0; $i < $l; $i++) {
            // 非重复, 加入数组
            if (!in_array($s[$i], $win)) {
                //if (!isset($win[$s[$i]])) {
                //$win[$s[$i]] = $s[$i];
                $win[] = $s[$i];
                $maxL = max($maxL, count($win));
                $win[] = $s[$i];
            } else {
                // 重复缩小数组
                if (count($win) > 1)
                    array_shift($win);
            }
            print_r($win);
        }

        return $maxL;
    }
}

//$r = (new Solution())->lengthOfLongestSubstring('aab');
//$r = (new Solution())->lengthOfLongestSubstring('abcabcacb');
//$r = (new Solution())->lengthOfLongestSubstring('bbbbbb');
//$r = (new Solution())->lengthOfLongestSubstring('pwwkew');
$r = (new Solution())->lengthOfLongestSubstring('dvdf');
print_r($r);
