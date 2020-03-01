<?php
/**
 * 有效的括号
 */

function isValid($s)
{
    if ($s == '') {
        return true;
    }
    $s = str_split($s);

    $arr = [];
    foreach ($s as $v) {
        if (!in_array($v, [')', '}', ']'])) {
            $arr[] = $v;
            continue;
        }

        $tmp = array_pop($arr);
        switch ($v) {
        case ')':
            if ($tmp != '(')
                return false;
            break;
        case ']':
            if ($tmp != '[')
                return false;
            break;
        case '}':
            if ($tmp != '{')
                return false;
            break;
        }
    }
    return $arr == [];
}

//$r = isValid('()[]{}');
//$r = isValid('(]');
//$r = isValid('([)]');
//$r = isValid('{[]}');
//$r = isValid('((');
$r = isValid('(');

var_dump($r);
