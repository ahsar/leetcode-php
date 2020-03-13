<?php

class Solution {

    /**
     * hash table
     *
     * @param Integer[] $nums
     * @return Integer
     */
    public function majorityElement($nums) {
        if (!$nums) {
            return 0;
        }

        $most = round(count($nums) / 2);
        $map = [];

        foreach ($nums as $v) {
            isset($map[$v]) || $map[$v] = 0;
            $map[$v]++;
            if ($map[$v] >= $most) {
                return $v;
            }
        }
        return $nums[0];
    }

    /**
     * majorityElement1
     *
     * 排序法
     * @param array $nums
     * @access public
     * @return int
     */
    public function majorityElement1(array $nums)
    {
        if (!$nums) {
            return 0;
        }

        $len = count($nums);
        if ($len == 1) {
            return $nums[0];
        }

        sort($nums);
        return $nums[floor($len / 2)];
    }
}

//$a = [3,2,3];
$a = [1, 2];
//$a = [1];
$a = [2,2,1,1,1,2,2];
//$a = [1,2,3];
$a = [6,5,5];
$r = (new Solution())->majorityElement1($a);
print_r($r);
