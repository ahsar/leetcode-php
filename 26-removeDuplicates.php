<?php
/**
 * 删除排序数组中的重复项
 */
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
function removeDuplicates1(&$nums) {
    $l = count($nums);
    $j = 0;

    for ($i = 1; $i < $l; $i++) {
        if ($nums[$j] != $nums[$i]) {
            $nums[++$j] = $nums[$i];
        }
    }
    return ++$j;
}

    /**
     * @param Integer[] $nums
     * @return Integer
     */
function removeDuplicates(&$nums) {
    $l = count($nums);
    $i = 1;
    $j = 0;

    while ($i < $l) {
        if ($nums[$i] != $nums[$j]) {
            $nums[++$j] = $nums[$i++];
        } else {
            $i++;
        }
    }
    return ++$j;
}
}

$arr = [
    //1,2
    //1,1,2
    //1,2,2,3
    1,2,3,3,4
];
$r = (new Solution())->removeDuplicates($arr);
print_r($arr);
print_r($r);
