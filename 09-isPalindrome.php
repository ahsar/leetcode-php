<?php
/**
 * 回文数
 */

function isPalindrome($x)
{
    if ($x < 0) {
        return false;
    }

    $y = $x;
    $n = 0;
    while ($x) {
        $n = $n * 10 + $x % 10;
        $x = intval($x / 10);
    }

    return $y == $n;
}

$r = isPalindrome(-121);
var_dump($r);
