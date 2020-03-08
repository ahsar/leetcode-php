<?php
/**
 * 凑硬币
 *
 */

class Solution {

    /**
     * @param Integer[] $coins
     * @param Integer $amount
     * @return Integer
     */
    function coinChange1($coins, $amount) {
        $dp = array_fill(0, $amount + 1, $amount + 1);
        $dp[0] = 0;

        for ($i = 0; $i < $amount + 1; $i++) {
            foreach ($coins as $coin) {
                if ($i - $coin < 0) {
                    continue;
                }
                $dp[$i] = min($dp[$i], 1 + $dp[$i - $coin]);
            }
        }
        return ($dp[$amount] == $amount + 1) ? -1 : $dp[$amount];
    }
}

$r = (new Solution())->coinChange1([1, 2, 5], 11);
//$r = (new Solution())->coinChange([2], 3);
//$r = (new Solution())->coinChange([1], 0);
//$r = (new Solution())->coinChange([2,5,10,1], 27);
//$r = (new Solution())->coinChange([186,419,83,408], 6249);

print_r($r);
