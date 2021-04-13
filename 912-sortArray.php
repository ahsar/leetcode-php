<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function sortArray($nums) {
        sort($nums);
        return $nums; 
    }
}

$nums = [4,3,2,1];
$r = (new Solution())->sortArray($nums);
print_r($r);
