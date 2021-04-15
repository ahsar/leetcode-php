<?php
/**
 * 两数之和
 *
 * 思路: hash
 */

function twoSum($nums, $target) {
    // 构造查找数组
    $hash = [];
    foreach ($nums as $i => $v) {
        $hash[$v] = $i;
    }

    foreach ($nums as $k => $v) {
        $dist = $target - $v;
        if (!isset($hash[$dist]) || $k == $hash[$dist]) {
            continue;
        }
        return [$k, $hash[$dist]];
    }

    return [];
}

$r = twoSum([-3,4,3,90], 0);
print_r($r);
