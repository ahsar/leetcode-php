<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    public function findRepeatNumber($nums)
    {
        $len = count($nums);
        for ($i = 0; $i < $len; $i++) {
            // 跳过
            if ($i == $nums[$i]) {
                continue;
            } else {
                // 交换且交换过程判断对应位置占用
                if ($nums[$nums[$i]] == $nums[$i]) {
                    return $nums[$i];
                }

                $tmp = $nums[$nums[$i]];
                $nums[$nums[$i]] = $nums[$i];
                $nums[$i] = $tmp;
            }
        }
    }
}

$arr = [1, 3, 2, 0, 1, 5];
$r = (new Solution())->findRepeatNumber($arr);
print_r($r);
