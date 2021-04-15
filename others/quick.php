<?php
/**
 * Class quickSort
 * @author yourname
 */
class QuickSort
{
    public function quick_sort(&$arr)
    {
        $this->quickSortC($arr, 0, count($arr) - 1);
    }

    private function quickSortC(&$arr, $l, $r)
    {
        if ($l >= $r) {
            return;
        }

        $q = $this->pation($arr, $l, $r);
        $this->quickSortC($arr, $l, $q - 1);
        $this->quickSortC($arr, $q + 1, $r);
    }

    /**
     * pation
     *
     * 分区函数
     * @param array $arr
     * @param int $l
     * @param int $r
     * @access private
     * @return int
     */
    private function pation(&$arr, $l, $r)
    {
        $i = $l;
        $j = $r - 1;

        // 标的元素
        $pivot = $arr[$r];
        while (1) {
            // 如果右侧元素小于标的, 则停止
            while ($j > $l && $arr[$j] > $pivot) {
                $j--;
            }

            // 如果右侧元素小于标的, 则停止
            while ($i < $r - 1 && $arr[$i] < $pivot) {
                $i++;
            };

            if ($i >= $j)
                break;

            // 交换左右元素
            [$arr[$i], $arr[$j]] = [$arr[$j], $arr[$i]];
        }

        // 将标记元素放入中间位置
        [$arr[$r], $arr[$i]] = [$arr[$i], $arr[$r]];
        return $i;
    }
}


$arr = [8, 10, 9, 11, 12, 17, 5];
//$arr = [8, 10, 6, 3, 2, 1, 5];
(new QuickSort())->quick_sort($arr);
print_r($arr);
