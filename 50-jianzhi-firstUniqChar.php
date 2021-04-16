<?php

class Solution {

    /**
     * @param String $s
     * @return String
     */
    public function firstUniqChar2($s)
    {
        if (!$s) 
            return ' ';

        $l = strlen($s);
        $queue = $hash = [];
        for ($i = 0; $i < $l; $i++) {
            // 存在重复的
            if (isset($hash[$s[$i]]) && $queue[0] == $s[$i]) {
                array_shift($queue);
            } else {
                $hash[$s[$i]] = $i;
                $queue[] = $s[$i];
            }
        }

        return $queue ? $queue[0] : '';
    }

    public function firstUniqChar($s)
    {
        if (!$s) {
            return ' ';
        }

        $hash = [];
        $l = strlen($s);
        for ($i = 0; $i < $l; $i++) {
            if (isset($hash[$s[$i]])) {
                $hash[$s[$i]]++;
            } else {
                $hash[$s[$i]] = 0;
            }
        }

        for ($i = 0; $i < $l; $i++) {
            if ($hash[$s[$i]] == 0) {
                return $s[$i];
            }
        }
        return ' ';
    }
}

//$s = 'abaccdeff';
$s = 'abbca';
$r = (new Solution())->firstUniqChar($s);
print_r($r);
