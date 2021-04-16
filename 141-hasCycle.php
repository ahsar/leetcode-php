<?php
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
    function detectCycle($head)
    {
       if (!$head || !$head->next) {
           return null; 
       } 

       $slow = $fast = $head;
       $flag = 0;
       while ($fast->next->next) {
           $slow = $slow->next;
           $fast = $fast->next->next;

           if ($fast == $slow) {
               $flag = 1;
               break; 
           }
       }
       if (!$flag) {
          return null; 
       }

       for ($slow = $head; $fast != $slow; $slow=$slow->next,$fast =$fast->next)
           return $fast;

       return null;
    }
}

//$head = new ListNode(1);
//insert($head, 2);
//insert($head, 0);
//insert($head, -4);
//insert($head, -3);
//insert($head, -5);
//insert($head, -6);
//$circle = $head->next->next->next;
//$head->next->next->next->next->next->next = $circle;
//lPrint($head);
//die;
//
$head = new ListNode(3);
insert($head, 2);
insert($head, 0);
insert($head, -4);

$r = (new Solution())->detectCycle($head);
print_r($r);
