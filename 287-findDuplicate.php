<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    public function findDuplicate($nums)
    {
        if (!$nums) {
            return 0; 
        }

        $fast = $slow = 0;

        do {
            $slow = $nums[$slow];
            $fast = $nums[$nums[$fast]];
        } while ($fast != $slow);

        $res = 0;
        while ($slow != $res) {
           $slow = $nums[$slow];
           $res = $nums[$res];
        }

        return $res;
    }
}

$arr = [1,3,4,2,2];
//$arr = [3,1,3,4,2];
//$arr = [1,1];
//$arr = [1,1, 2];
$r = (new Solution())->findDuplicate($arr);
print_r($r);
