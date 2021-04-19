<?php

/**
 * Solutio
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

    /**
     * getLeastNumbers2
     *
     * 基于二分法最小k个数
     * @param array $arr
     * @param int $k
     * @access public
     * @return array
     */
    public function getLeastNumbers2($arr, $k)
    {
        $this->quickSortc($arr, $k, 0, count($arr) - 1);
        return array_slice($arr, 0, $k);
    }

    protected function quickSortc(&$arr, $k, $l, $r)
    {
        if ($l >= $r) {
            return;
        }

        $mid = $this->partion($arr, $l, $r);
        if ($k == $mid) {
            return;
        }
        if ($mid > $k) {
            $this->quickSortc($arr, $k, $l, $mid - 1);
        } else {
            $this->quickSortc($arr, $k, $mid + 1, $r);
        }
    }

    protected function partion(&$arr, $l, $r)
    {
        $i = $l;
        $pivot = $arr[$r];

        for ($j = $l; $j < $r; $j++) {
            if ($arr[$j] < $pivot) {
                [$arr[$i], $arr[$j]] = [$arr[$j], $arr[$i]];
                $i++;
            }
        }

        [$arr[$r], $arr[$i]] = [$arr[$i], $arr[$r]];
        return $i;
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
