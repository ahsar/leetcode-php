<?php
/**
 * 圆圈中最后剩下的数字
 */

class Solution {

    /**
     * @param Integer $n
     * @param Integer $m
     * @return Integer
     */
    public function lastRemaining($n, $m) {
        $p = 0;

        for ($i = 2; $i < $n; $i++) {
            $p = ($p + $m) % $i;
        }

        return $p;
    }
}

$n = [0,1,2,3,4];
$m = 3;

$n = [5];
$m = 3;

$n = [5];
$m = 3;

$r = (new Solution())->lastRemaining($n, $m);
print_r($r);
