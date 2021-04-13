<?php

class Solution
{
    public function getLeastNumbers($arr, $k)
    {
        $this->pation($arr, 0, count($arr) - 1, $k);
        return array_slice($arr, 0, $k);
    }

    private function pation(&$arr, $lo, $hi, $k)
    {
        if ($hi < $lo) {
            return;
        }

        $i = $lo;
        $j = $hi;
        $pvot = $arr[$i];

        while ($i < $j) {
            while ($i < $j && $arr[$j] >= $pvot) {
                $j--;
            }
            if ($i < $j) {
                $arr[$i++] = $arr[$j];
            }
            while ($i < $j && $arr[$i] <= $pvot) {
                $j++;
            }
        }

        $arr[$i] = $pvot;

        if ($i < $k -1) {
            $this->pation($arr, $i + 1, $hi, $k); /* 右区间递归调用 */ 
        } else {
            
            $this->pation($arr, $lo, $hi-1, $k); /* 右区间递归调用 */ 
        }
    }
}
