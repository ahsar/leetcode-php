<?php
/**
 * 链表的中间结点
 */

require_once __DIR__ . '/structure/linklist.php';

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    public function middleNode($head) {
        $slow = $cur = $head;

        $num = 0;
        while ($cur) {
            $num++;
            if (($num % 2) == 0) {
                $slow = $slow->next;
            }
            $cur = $cur->next; 
        }

        return $slow;
    }
}

$link = new ListNode(1);
insert($link, 2);
insert($link, 3);
insert($link, 4);
insert($link, 5);
insert($link, 6);

$r = (new Solution())->middleNode($link);
lPrint($link);
