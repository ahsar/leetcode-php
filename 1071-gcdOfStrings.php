<?php

class Solution {

    /**
     * @param String $str1
     * @param String $str2
     * @return String
     */
    public function gcdOfStrings($str1, $str2) {
        if ($str1 . $str2 != $str2 . $str1) {
            return '';
        }

        $l1 = strlen($str1);
        $l2 = strlen($str2);
        return substr($str1, 0, $this->gcd($l1, $l2));
    }

    public function gcd(int $a, int $b)
    {
        while ($a != $b) {
            if ($a > $b) {
                $a = $a - $b;
            } else {
                $b = $b - $a;
            }
        }

        return $a;
    }
}

//$a = 'ABCABC';
//$b = 'ABC';
$a = 'LEET';
$b = 'CODE';

$r = (new Solution())->gcdOfStrings($a, $b);
print_r($r);
