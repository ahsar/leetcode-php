<?php

$a = [1,2,5, 7];
$b = [3,4, 10];

function merge($a, $b) {
    $i = $j = 0;
    $c = [];
    while (isset($a[$i] ) && isset($b[$j])) {
        if ($a[$i] < $b[$j]) {
            $c[] = $a[$i];
            $i++;
            continue;
        }

        if ($a[$i] >= $b[$j]) {
            $c[] = $b[$j];
            $j++;
        }
    }

    while (isset($a[$i])) {
        $c[] = $a[$i++];
    }

    while (isset($b[$j])) {
        $c[] = $b[$j++];
    }

    print_r($c);
}

merge($a, $b);
