<?php
/**
 * 整数反转 
 */

function reverse($x) {
    $n = 0;
    $max = 0x7fffffff;

    while ($x) {
        $n = $n * 10 + $x % 10;
        $x = intval($x / 10);
    }
    return (($n > $max - 1) || $n < -$max) ? 0 : $n;
}

//$r = reverse(123);
//$r = reverse( -123);
$r = reverse(1534236469);

print_r($r);
