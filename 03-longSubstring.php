<?php

class Solution 
{
    public function lengthOfLongestSubstring($s)
    {
        $l = strlen($s);
        if ($l < 1) {
            return 0;
        }

        $maxL = $start = 0;
        $hash = [];

        for ($i = 0; $i < $l; $i++) {
            $str = $s[$i];

            if (isset($hash[$str])) {
                $start = max($hash[$str] + 1, $start);
            }
            $maxL = max($i - $start + 1, $maxL);

            $hash[$str] = $i;
        }

        return $maxL;
    }
}

//$r = (new Solution())->lengthOfLongestSubstring('aab');
$r = (new Solution())->lengthOfLongestSubstring('bbbbbb');
//$r = (new Solution())->lengthOfLongestSubstring('pwwkew');
//$r = (new Solution())->lengthOfLongestSubstring('dvdf');
//$r = (new Solution())->lengthOfLongestSubstring(" ");
//$r = (new Solution())->lengthOfLongestSubstring("aux");
//$r = (new Solution())->lengthOfLongestSubstring("a");
echo "\n";
print_r($r);
