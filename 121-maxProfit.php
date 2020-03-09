<?php
class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        $maxPrice = 0;
        $minPrice = $prices[0];

        foreach ($prices as $price) {
            $maxPrice = max($maxPrice, $price - $minPrice);
            $minPrice = min($minPrice, $price);
        }
        return $maxPrice;
    }

    /**
     * maxProfit1
     *
     * 动态规划实现
     * $dp[$i] = max($dp[$i], $dp[$i-1]);
     *
     * @param array $prices
     * @access public
     * @return array
     */
    public function maxProfit1($prices)
    {
        $dp[0] = 0;
        $minPrice = $prices[0];
        /*
         * 1. 定状态
         * 2. 定义dp函数
         * 3. 明确选择(最优)
         * 4. 明确base
         */
        $len = count($prices);
        for ($i = 1; $i <= $len; $i++) {
            $minPrice = min($prices[$i], $minPrice);
            $price = $prices[$i] - $minPrice;
            $dp[$i] = max($price, $dp[$i-1]);
        }

        return end($dp);
        print_r($dp);
    }
}

$r = (new Solution())->maxProfit1([5,10,7,13]);
//$r = (new Solution())->maxProfit1([7,6,4,3,1]);
//$r = (new Solution())->maxProfit1([7,1,5,3,6,4]);

print_r($r);
