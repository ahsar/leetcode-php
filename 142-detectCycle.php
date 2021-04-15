<?php

require_once  __DIR__ . '/structure/linklist.php';

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
    public function detectCycle($head)
    {
        if (!$head || !$head->next) {
            return null;
        }

        $slow = $fast = $head;
        do {
            if ($fast == null || $fast->next == null) {
                return null; 
            }

            $slow = $slow->next;
            $fast = $fast->next->next;
        } while ($fast != $slow);

        for ($slow = $head; $fast != $slow; $slow=$slow->next,$fast =$fast->next)
            return $fast;
    }
}

$head = new ListNode(3);
insert($head, 2);
insert($head, 0);
insert($head, -4);
$circle = $head->next;
$head->next->next->next = $circle;
//lPrint($head);
//die;

$r = (new Solution())->detectCycle($head);
print_r($r);
