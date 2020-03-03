<?php
/**
 * 移除元素 
 */

function removeElement(&$nums, $val) {
    $len = count($nums);
    $y = 0;

    for ($i = 0; $i < $len; $i++) {
        if ($nums[$i] != $val) {
            $nums[$y++] = $nums[$i];
        }
    }
    return $y;
}

$arr = [0,1,2,2,3,0,4,2];
$arr = [3,2,2,3];
$r = removeElement($arr, 3);
print_r($arr);
print_r($r);
