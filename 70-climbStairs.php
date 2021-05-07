<?php

/**
 * Solution 
 * 
 * @package 爬楼梯
 * @version $id$
 * @author letwhip<letwhip@gmail.com>
 */
class Solution
{

    /**
     * @param Integer $n
     * @return Integer
     */
    public function climbStairs($n)
    {
        $r = [
            0 => 1,
            1 => 1,
        ];
        for ($i = 2; $i <= $n; $i++) {
            $r[$i] = $r[$i - 1] + $r[$i - 2];
        }

        return $r[$n];
    }
}

$n = 0;
$r = (new Solution())->climbStairs($n);
print_r($r);
