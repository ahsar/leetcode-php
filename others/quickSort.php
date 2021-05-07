<?php

function quickSort(&$arr)
{
    quickSortC($arr, 0, count($arr) - 1);
}

function quickSortC(&$arr, $l, $r)
{
    if ($l >= $r) {
        return;
    }

    $mid = partion($arr, $l, $r);
    quickSortC($arr, $l, $mid - 1);
    quickSortC($arr, $mid + 1, $r);
}

function partion(&$arr, $l, $r)
{
    $j = $i = $l;
    $pivot = $arr[$r];

    for (; $j < $r; $j++) {
       if ($arr[$j] < $pivot) {
           [$arr[$i], $arr[$j]] = [$arr[$j], $arr[$i]];
           $i++;
       }
    }

    [$arr[$i], $arr[$r]] = [$arr[$r], $arr[$i]];
    return $i;
}

$arr = [6,1,4,5,9,7,10,8];
//$arr = [3,2,3,1,2,4,5,5,6];
//$arr = [0,0,0, 3, 0, 6, 5];
quickSort($arr);
print_r($arr);
