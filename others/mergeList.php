<?php

include_once __DIR__ . '../structure/linklist.php';

function mergeList($a, $b) {
    if (!$a) {
        return $b;
    }
    if (!$b) {
        return $b;
    }

    $newLink = $cur = new ListNode(0);

    while ($a && $b) {
        if ($a->val <= $b->val) {
            $cur->next = $a;
            $cur = $cur->next; // 向后搬移指针
            $a = $a->next;
        } else {
            $cur->next = $b;
            $cur = $cur->next;
            $b = $b->next;
        }
    }

    if ($a) {
        $cur->next = $a;
    }
    if ($b) {
        $cur->next = $b;
    }
    return $newLink->next;
}
