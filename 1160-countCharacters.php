<?php
/**
 * 拼写单词
 */

class Solution {

    /**
     * @param String[] $words
     * @param String $chars
     * @return Integer
     */
    public function countCharacters($words, $chars) {
        $l1 = strlen($chars);
        $hash = [];
        for ($i = 0; $i < $l1; $i++) {
            if (!isset($hash[$chars[$i]])) {
                $hash[$chars[$i]] = 1;
            } else {
                ++$hash[$chars[$i]];
            }
        }

        $total = 0;
        foreach ($words as $strings) {
            $tmp = $hash;
            $len = strlen($strings);
            $flag = false;
            $sum = 0;
            for ($i = 0; $i < $len; $i++) {
                // 从字母表未找到
                if (!isset($tmp[$strings[$i]])) {
                    $flag = false;
                    break 1;
                } else if ($tmp[$strings[$i]] == 0) {
                    $flag = false;
                    // 超出最大使用次数
                    break 1;
                }
                ++$sum;
                --$tmp[$strings[$i]];
                $flag = true;
            }
            $flag && $total += $sum;
        }

        return $total;
    }
}


$words = ['hello','world','leetcode'];
$chars = 'welldonehoneyr';

$words = ['cat','bt','hat','tree'];
$chars = 'atach';
//$words = ['caat'];
//$chars = 'atch';

$r = (new Solution())->countCharacters($words, $chars);
print_r($r);
