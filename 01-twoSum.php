<?php

/**
 * Solution
 *
 * 两数之和
 *
 * 一次遍历 hash
 * @package
 * @version $id$
 * @author letwhip<letwhip@gmail.com>
 */
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    public function twoSum($nums, $target)
    {
        $res = $hash = [];
        $l = count($nums);

        for ($i = 0; $i < $l; $i++) {
            $dist = $target - $nums[$i];

            if (isset($hash[$nums[$i]]) ) {
                // 当前项索引和值索引
                $res = [$i, $hash[$nums[$i]]];
                return $res;
            }

            $hash[$dist] = $i;
        }

        return $res;
    }
}

$nums = [2,7,11,15];
$target = 9;

//$nums = [3,2,4];
//$target = 6;

$nums = [3,3];
$target = 6;

$r = (new Solution())->twoSum($nums, $target);
print_r($r);
