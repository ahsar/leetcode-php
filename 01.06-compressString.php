<?php
/**
 * 压缩字符串
 */

class Solution {

    /**
     * @param String $S
     * @return String
     */
    public function compressString($S)
    {
        $len = strlen($S);
        if ($len <= 1) {
            return $S;
        }

        $new = '';
        $tmp = 1;
        for ($i = 0; $i < $len - 1; $i++) {
            if ($S[$i] == $S[$i + 1]) {
                $tmp++;
            } else {
                $tmp && $new .= $S[$i] . $tmp;
                $tmp = 1;
            }

            if ($i+1 == $len - 1) {
                $tmp && $new .= $S[$i + 1] . $tmp;
            }
        }

        return strlen($new) > $len ? $S : $new;
    }
}

//$s = 'aabcccccaaab';
$s = 'abbccd';
//$s = 'aabcccccaaaabb';
//$s = 'abcccccaaa';
$r = (new Solution())->compressString($s);
print_r($r);
