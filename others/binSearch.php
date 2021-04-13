<?php

$a = [1, 3, 5, 7, 9];
function binSearch($arr, $find)
{
    $left = 0;
    $right = count($arr) - 1;
    while ($left <= $right) {
        $mid = intval(($left + $right) / 2);
        if ($find < $arr[$mid]) {
            $right = $mid - 1;
        }
        if ($find > $arr[$mid]) {
            $left = $mid + 1;
        }
        if ($find == $arr[$mid]) {
            return $mid;
        }
    }

    return -1;
}

$c = binSearch($a, 3);
print_r($c);
