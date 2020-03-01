<?php
/**
 * 最长前缀 
 */

function longestCommonPrefix($strs) {
    if (empty($strs)) {
        return '';
    }

    $r = '';
    $len = strlen($strs[0]);
    for ($i = 0; $i < $len ; $i++) {
        $str1 = $strs[0][$i];
        foreach ($strs as $v) {
            if ($v[$i] != $str1) {
                return $r; 
            }
        }
        $r .= $v[$i];
    }
    return $r;
}
$r = longestCommonPrefix(['dog','racecar','car']);

var_dump($r);
