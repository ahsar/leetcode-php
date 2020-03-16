<?php
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
require_once __DIR__ . '/structure/linklist.php';

class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $n1 = $l1;
        $n2 = $l2;
        while ($n1 && $n2) {
            $n1->val += $n2->val;
            if ($n1->val >= 10) {
                if (!$n1->next) {
                    $n1->next = new ListNode(0);
                }
                $n1->next->val++;
                $n1->val -= 10;
            }

            $prev = $n1;
            $n1 = $n1->next;
            $n2 = $n2->next;
        }

        while ($n1) {
            if ($n1->val >= 10) {
                if (!$n1->next) {
                    $n1->next = new ListNode(0);
                }
                $n1->next->val++;
                $n1->val -= 10;
            }
            $n1 = $n1->next;
        }

        if ($n2) {
            $prev->next = $n2;
        }
        return $l1;
    }
}

$l1 = (new ListNode(2));
insert($l1, 4);
insert($l1, 3);
//insert($l1, 3);
//insert($l1, 9);

echo "\n";

$l2 = (new ListNode(5));
insert($l2, 6);
insert($l2, 4);
//insert($l2, 1);
//insert($l2, 3);

$r = (new Solution())->addTwoNumbers($l1, $l2);
lPrint($r);
