<?php

class Solution {

    /**
     * @param Integer[][] $matrix
     * @param Integer $target
     * @return Boolean
     */
    public function findNumberIn2DArray($matrix, $target) 
    {
        if (!$matrix) 
            return false;

        $col = count($matrix) - 1; //行
        $l_len = count($matrix[0]) - 1;
        $row = 0; // 列

        // 当列索引为0, 行索引大于行结束搜索
        while ($col >= 0 && $row <= $l_len) {
            printf("col is %d, row is %d\n", $col, $row);
            if ($target == $matrix[$col][$row]) 
                return true;

            if ($target > $matrix[$col][$row])
                $row++;

            if ($target < $matrix[$col][$row])
                $col--;
        }

        return false;
    }
}

$arr = [
    [1, 4, 7, 11, 15],
    [2, 5, 8, 12, 19],
    [3, 6, 9, 16, 22],
    [10, 13, 14, 17, 24],
    [18, 21, 23, 26, 30]
];
$arr = [
    [1]
];
$arr = [
    []
];

//$r = (new Solution())->findNumberIn2DArray($arr, 5);
$r = (new Solution())->findNumberIn2DArray($arr, 1);
var_dump($r);
