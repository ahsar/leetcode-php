<?php
/**
 * 删除链表的倒数第N个节点
 * 快慢指针
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
     * @param Integer $n
     * @return ListNode
     */
    public function removeNthFromEnd($head, $n) {
        if (!$head || !$n) {
            return $head;
        }

        $cur = $head;
        $slow = null; // 慢指针的前项节点
        $i = 0;
        while ($cur) {
            // 获取父节点
            if ($i == $n + 1) {
                $slow = $head;
            }
            $slow && $slow = $slow->next;
            $cur = $cur->next;
            $i++;
        }

        // 处理特殊情况
        if ($n == $i -1) {
            $slow = $head;
        }
        // 超过链表长度
        if ($n >= $i) {
            $head = $head->next;
        }

        if ($slow) {
            $deletNode = $slow->next;
            $slow->next = $slow->next->next;
            unset($deletNode);
        }
        return $head;
    }
}

$link1 = new ListNode(1);
insert($link1, 9);
insert($link1, 10);
insert($link1, 11);
insert($link1, 12);
insert($link1, 15);
insert($link1, 14);
$r = (new Solution())->removeNthFromEnd($link1, 8);
lPrint($r);
