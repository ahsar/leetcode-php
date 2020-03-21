<?php
/**
 * 最长上升子序列
 */

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    public function lengthOfLIS($nums) {
        if (!$nums) {
            return 0;
        }

        $len = count($nums);
        $dp = array_fill(0, $len, 1);

        $res = 0;
        for ($i = 0; $i < $len; $i++) {
            for ($j = 0; $j < $i; $j++) {
                if ($nums[$j] < $nums[$i]) {
                    $dp[$i] = max($dp[$i], $dp[$j] + 1);
                }
                $res = max($res, $dp[$i]);
            }
        }
        return $res;
    }
}

$r = (new Solution())->lengthOfLIS([10,9,2,5,3,7,101,18]);
//$r = (new Solution())->lengthOfLIS([2,3, 5,7]);
print_r($r);
