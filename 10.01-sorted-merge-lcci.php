<?php
/**
 * 合并排序的数组
 *
 * 给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
 * 初始化 A 和 B 的元素数量分别为 m 和 n。
 */

/**
 * @param Integer[] $A
 * @param Integer $m
 * @param Integer[] $B
 * @param Integer $n
 * @return NULL
 */
function merge(&$A, $m, $B, $n) {
    $i = $m - 1;
    $y = $n - 1;
    $k = count($A);

    // 从后向前遍历, 避免从前向后需要搬移数据
    while ($i >= 0 && $y >= 0) {
        if ($B[$y] > $A[$i]) {
            $A[--$k] = $B[$y--];
        } else {
            $A[--$k] = $A[$i--];
        }
    }

    while ($y >= 0) {
        $A[--$k] = $B[$y--];
    }
}

$a = [2,0];
$m = 1;
$b = [1];
$n = 1;

//$a = [0];
//$m = 0;
//$b = [1];
//$n = 1;

//$a = [1,2,3,0, 0, 0];
//$m = 3;
//$b = [2, 5,6];
//$n = 3;

//$a = [1,2,3,0];
//$m = 3;
//$b = [2];
//$n = 1;

//$a = [1,0];
//$m = 1;
//$b = [2];
//$n = 1;

merge($a, $m, $b, $n);
print_r($a);
