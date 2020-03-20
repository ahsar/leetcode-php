<?php
/**
 * 最小的k个数
 */
class Solution {

    /**
     * @param Integer[] $arr
     * @param Integer $k
     * @return Integer[]
     */
    public function getLeastNumbers($arr, $k) {
        if ($k == 0) {
            return [];
        }
        sort($arr);
        return array_slice($arr, 0, $k);
    }
}

//$arr = [3,2,1];
//$k = 2;
$arr = [3,2,2,1, 8, 9,5, 6];
$k = 4;
$r = (new Solution())->getLeastNumbers($arr, $k);
print_r($r);
