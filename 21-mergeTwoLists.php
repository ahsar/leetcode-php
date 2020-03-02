<?php
/**
 * 合并两个有序链表
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 * 注意指针的运用
 */

include_once __DIR__ . '/structure/linklist.php';

function mergeTwoLists($l1, $l2) {
    $new_link = $cur = new ListNode(0);
    while ($l1 && $l2) {
        if ($l1->val <= $l2->val) {
            $cur->next = $l1;
            $cur = $cur->next;
            $l1 = $l1->next;
        } else {
            $cur->next = $l2;
            $cur = $cur->next;
            $l2 = $l2->next;
        }
    }

    if ($l1) {
        $cur->next = $l1;
    } else {
        $cur->next = $l2;
    }
    return $new_link->next;
}

$link1 = $link2 = null;
insert($link1, 1);
insert($link1, 3);
insert($link1, 5);
insert($link1, 7);
insert($link1, 9);

insert($link2, 2);
insert($link2, 4);
insert($link2, 6);
insert($link2, 8);
insert($link2, 10);

$r = mergeTwoLists($link2, $link1);
lPrint($r);
