<?php
/**
 * 将数组分成和相等的三个部分
 */

class Solution {

    /**
     * @param Integer[] $A
     * @return Boolean
     */
    public function canThreePartsEqualSum($A)
    {
        $sum = array_sum($A);
        if (0 != ($sum % 3)) {
            return false;
        }
        $part = $sum / 3;

        $total = 0;
        
        $count = 0;
        foreach ($A as $v) {
            $total += $v;
            if ($total == $part) {
                $count++;
                $total = 0;
            }
        }

        return $count == 3 || ($count > 3 && $sum == 0);
    }
}

//$r = (new Solution())->canThreePartsEqualSum([0,2,1,-6,6,-7,9,1,2,0,1]);
//$r = (new Solution())->canThreePartsEqualSum([0,2,1,-6,6,7,9,-1,2,0,1]);
//$r = (new Solution())->canThreePartsEqualSum([3,3,6,5,-2,2,5,1,-9,4]);
//$r = (new Solution())->canThreePartsEqualSum([12,-4,16,-5,9,-3,3,8,0]);
$r = (new Solution())->canThreePartsEqualSum([1,-1,1,-1]);
//$r = (new Solution())->canThreePartsEqualSum([10,-10,10,-10,10,-10,10,-10]);


var_dump($r);
