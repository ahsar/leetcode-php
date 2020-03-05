<?php
/**
 * 分糖果2
 */

class Solution {

    /**
     * @param Integer $candies
     * @param Integer $num_people
     * @return Integer[]
     */
    function distributeCandies($candies, $num_people) {
        if ($candies <= 0 || $num_people <= 0) {
            return [];
        }

        $distribute = array_fill(0, $num_people, 0);
        $y = 0;
        while ($candies > 0) {
            for ($i = 0; $i < $num_people; $i++) {

                $should = $i+1 + $y * $num_people;
                $fact = ($candies > $should) ? $should : $candies;
                $distribute[$i] += $fact;
                $candies -= $fact;
                if ($candies <= 0) {
                    break 2; 
                }
            }
            $y++;
        }
        return $distribute;
    }
}

$r = (new Solution())->distributeCandies(7, 4);
print_r($r);
