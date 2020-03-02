<?php
/**
 * 反转链表
 */

class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function reverseList($head) {
        $prev = $next = null;

        while ($head) {
            $next = $head->next;
            $head->next = $prev;
            $prev = $head;
            $head = $next;
        }
        return $prev;
    }
}
