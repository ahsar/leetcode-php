<?php
require_once('../structure/linklist.php');

$link = new ListNode(1);
insert($link, 2);
insert($link, 3);
insert($link, 4);
insert($link, 5);
insert($link, 6);
reverse($link);

function reverse($link)
{
    $next = $prev = null;
    while ($link) {
        $next = $link->next;
        $link->next = $prev;
        $prev = $link;
        $link = $next;
    }
    return $prev;
}

