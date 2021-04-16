<?php

$arr = [1, 2, 3, 5, 11, 15, 17, 19, 30, 40, 50];

function binSearch($arr, $v)
{
    $low = 0;
    $high = count($arr) - 1;

    while ($low <= $high) {
       $mid = floor(($low + $high) / 2);
       if ($v < $arr[$mid]) {
           $high = $mid - 1;
       }
       if ($v > $arr[$mid]) {
           $low = $mid + 1;
       }
       if ($v == $arr[$mid]) {
           return $mid;
       }
    }

    return -1;
}
$a = binSearch($arr, 30);
print_r($a);

