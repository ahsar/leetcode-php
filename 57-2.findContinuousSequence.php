<?php
class Solution {

    /**
     * 暴力法
     *
     * 获取[1, (target / 2) + 1]
     * @param Integer $target
     * @return Integer[][]
     */
    function findContinuousSequence2($target) {
        if ($target <= 1) {
            return [];
        }

        $end = round($target / 2);

        $res = [];
        for ($i = 1; $i <= $end; $i++) {
            $sum = 0;
            $total = [];

            for ($y = $i; $y <= $end; $y++) {
                $sum += $y;
                if ($sum > $target) {
                    break 1;
                }
                $total[] = $y;
                if ($sum == $target) {
                    $res[] = $total;
                    break 1;
                }
            }
        }
        return $res;
    }

    /**
     * findContinuousSequence
     *
     * 和为s的连续正数序列
     * 滑动窗口法
     *
     * @param int $target
     * @access public
     * @return array
     */
    public function findContinuousSequence($target)
    {
        if ($target <= 1) {
            return [];
        }

        $end = round($target / 2);

        $res = [];
        $sum = 0;

        for ($i = 1, $j = 1; $i <= $end;) {
            if ($sum > $target) {
                $sum -= $i++;
            } else if ($sum < $target) {
                $sum += $j++;
            } else {
                $res[] = range($i, $j-1);
                $sum -= $i++;
            }
        }

        return $res;
    }
}

//$r = (new Solution())->findContinuousSequence(50252);
$r = (new Solution())->findContinuousSequence(15);
print_r($r);
