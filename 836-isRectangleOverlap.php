<?php

class Solution {

    /**
     * @param Integer[] $rec1
     * @param Integer[] $rec2
     * @return Boolean
     */
    public function isRectangleOverlap($rec1, $rec2) {
        // 矩形1 在矩形2 的四面之外, 则没有重叠部分
        return !(
            // 矩形 1 在 2 的左边之外, 矩形1右上角, 小于等于矩形2的左下角
            $rec1[2] <= $rec2[0]

            // 矩形 1 在 2 的下边之外, 矩形1右上角, 小于等于矩形2的左下角
            || $rec1[3] <= $rec2[1]

            // 矩形 1 在 2 的右边之外, 矩形2右上角, 小于等于矩形1的左下角
            || $rec2[2] <= $rec1[0]

            // 矩形 1 在 2 的上边之外, 矩形1左下角, 大于等于矩形2的右上角
            || $rec1[1] >= $rec2[3]
        );
    }
}

//$rec1 = [0, 0, 2, 2];
//$rec2 = [1, 1, 3, 3];

$rec1 = [0, 0, 1, 1];
$rec2 = [1, 0, 2, 1];
$r = (new Solution())->isRectangleOverlap($rec1, $rec2);
var_dump($r);
