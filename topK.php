<?php

class Solution
{
    public function getLeastNumbers($arr, $k)
    {
        $this->quickSortC($arr, $k, 0, count($arr) - 1);
        return array_slice($arr, 0, $k);
    }

    private function quickSortC(&$arr, $k , $l, $r)
    {
        if ($l >= $r) {
            return [];
        }

        $index = $this->pation($arr, $l, $r);

        if ($index == $k -1) {
            return $arr;
        }
        if ($index > $k - 1) {
            $this->quickSortC($arr, $k, $l, $index -1);
        } else {
            $this->quickSortC($arr, $k, $index + 1, $r);
        }
    }

    private function pation(&$arr, $l, $r)
    {
        $i = $l;
        $j = $r - 1;

        $pivot = $arr[$r];
        while (1) {
            // 左侧节点小则停止
            while ($i < $r && $arr[$i] > $pivot) {
                $i++;
            }
            // 右侧节点大则停止
            while ($j > 0 && $arr[$j] < $pivot) {
                $j--;
            }

            if ($i >= $j) {
                break;
            }
            [$arr[$i], $arr[$j]] = [$arr[$j], $arr[$i]];
        }

        [$arr[$i], $arr[$r]] = [$arr[$r], $arr[$i]];
        return $i;
    }
}

$arr = [3,2,1,5,6,4];
$k = 2;

$arr = [3,2,3,1,2,4,5,5,6];
$k = 4;

$arr = [0,0,0,2,0,5];
$k = 0;
$r = (new Solution())->getLeastNumbers($arr, $k);
print_r($r);
