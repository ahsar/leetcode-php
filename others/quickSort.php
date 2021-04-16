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
    $i = $l;
    $j = $r - 1;
    $pivot = $arr[$r];

    while (1) {
        // 右边界的数字大于标的,不换
        while ($j > 0 && $arr[$j] > $pivot) {
            $j--;
        }

        // 左边界数字小于边界, 不换
        while ($i < $r && $arr[$i] < $pivot)
            $i++;

        if ($i >= $j) {
            break;
        }

        [$arr[$i], $arr[$j]] = [$arr[$j], $arr[$i]];
    }

    [$arr[$i], $arr[$r]] = [$arr[$r], $arr[$i]];
    return $i;
}

//$arr = [6,1,4,5,9,7,10,8];
$arr = [3,2,3,1,2,4,5,5,6];
quickSort($arr);
print_r($arr);
