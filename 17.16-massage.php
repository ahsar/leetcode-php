<?php
/**
 * 按摩师
 */

class Solution {

    /**
     * 动态规划
     *
     * @param Integer[] $nums
     * @return Integer
     */
    public function massage($nums)
    {
        if (!$nums) {
            return 0;
        }

        $len = count($nums);
        if ($len <= 2) {
            return max($nums);
        }

        $dp[0] = $nums[0];
        $dp[1] = max($nums[1], $nums[0]);
        for ($i = 2; $i < $len; $i++) {
            // 前一项选中, 当前项没有
            // 前2项选中, 当前项选中
            $dp[$i] = max($dp[$i-1], $dp[$i-2] + $nums[$i]);
        }
        return end($dp);
    }
}

//$order = [1, 2, 3, 1];
$order = [2, 7, 9, 3, 1];
//$order = [2, 1, 4, 5, 3, 1, 1, 3];
//$order = [2,1, 1, 2];

$r = (new Solution())->massage($order);
print_r($r);
