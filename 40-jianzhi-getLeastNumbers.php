<?php

/**
 * Solution 
 * 
 * 剑指 Offer 40. 最小的k个数
 * 使用大顶堆实现
 *
 * @package 
 * @author letwhip<letwhip@gmail.com> 
 */
class Solution
{

    /**
     * @param Integer[] $arr
     * @param Integer $k
     * @return Integer[]
     */
    public function getLeastNumbers($arr, $k)
    {
        $l = count($arr);
        if (!$l || !$k || $k > $l) {
            return [];
        }

        $heap = new \SplMaxHeap();
        for ($i = 0; $i < $k; $i++) {
            $heap->insert($arr[$i]);
        }

        for ($i = $k; $i < $l; $i++) {
            if ($heap->top() > $arr[$i]) {
                $heap->extract();
                $heap->insert($arr[$i]);
            }
        }

        $r = [];
        for ($i = 0; $i < $k; $i++) {
            $r[] = $heap->extract();
        }
        return $r;
    }
}

$arr = [3,2,1];
$k = 2;

$arr = [0,1,2,1];
$k = 1;

$arr = [0,0,1,3,4,5,0,7,6,7];
$k = 9;
$r = (new Solution())->getLeastNumbers($arr, $k);
print_r($r);
