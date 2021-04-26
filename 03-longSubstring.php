<?php

class Solution
{
    public function lengthOfLongestSubstring2($s)
    {
        if (!$s) {
            return 0;
        }

        $l = strlen($s);
        $maxL = 0;
        $win = [];

        // 使用hash
        for ($i = 0; $i < $l; $i++) {
            // 非重复, 加入数组
            if (!in_array($s[$i], $win)) {
                //if (!isset($win[$s[$i]])) {
                //$win[$s[$i]] = $s[$i];
                $win[] = $s[$i];
                $maxL = max($maxL, count($win));
                $win[] = $s[$i];
            } else {
                // 重复缩小数组
                if (count($win) > 1)
                    array_shift($win);
            }
            print_r($win);
        }

        return $maxL;
    }

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
