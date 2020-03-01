<?php
/**
 * 链表
 */

 class ListNode {
     public $val = 0;
     public $next = null;
     function __construct($val) { $this->val = $val; }
 }

function insert(&$link, int $v) {
    $node = new ListNode($v);
    if ($link->val == 0) {
        $link = $node;
        return;
    }

    $cur = $link;
    while($cur->next != null) {
        if ($cur->next == null) {
            break;
        }
        $cur = $cur->next;
    }

    $cur->next = $node;
}

function lPrint(&$link) {
    $cur = $link;
    while ($cur) {
        print_r($cur->val);
        echo ' -> ';
        $cur = $cur->next;
    }
}

