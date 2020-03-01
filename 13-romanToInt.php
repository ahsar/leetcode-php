<?php
/**
 * 罗马数字转整数
 */

function romanToInt($s)
{
    $map = [
        'I' => 1,
        'V' => 5,
        'X' => 10,
        'L' => 50,
        'C' => 100,
        'D' => 500,
        'M' => 1000,
    ];

    for ($i = strlen($s) - 1, $n = $map[$s[$i]]; $i > 0; $i--) {
        $cur = $map[$s[$i]];
        $next = $map[$s[$i-1]];
        if ($next < $cur) {
            $n -= $next;
        } else {
            $n += $next;
        }
    }

    return $n;
}

//$r = romanToInt('III');
//$r = romanToInt('IV');
//$r = romanToInt('LVIII');
$r = romanToInt('MCMXCIV');

print_r($r);
