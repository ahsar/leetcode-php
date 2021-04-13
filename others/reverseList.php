<?php

function reverseList($link)
{
    $prev = $next = null;

    while ($link) {
        $next = $link->next;
        $link->next = $prev;
        $prev = $link;
        $link = $next;
    }
    return $link;
}
