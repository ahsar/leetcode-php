<?php

class Solution {

    /**
     * @param Integer[][] $matrix
     * @return Integer[]
     */
    public function spiralOrder($matrix)
    {
        if (!$matrix)
            return [];

        //$r = [];
        $row = $col = 0;
        $row_right = count($matrix[0]);
        // 定义边界, 边界到了退出矩阵遍历
        while (true) {
            print_r($matrix[$i][$j]);
            if ($row == $row_right) // 行坐标抵达右边界,纵坐标增加
                $col++;

        }
        //return $r;
    }
}
/**
 *  [1,2,3]
 *  [4,5,6]
 *  [7,8,9]
 */

$r = (new Solution())->spiralOrder([[1,2,3],[4,5,6],[7,8,9]]);
print_r($r);
