<?php
/**
 * ip 地址无效化
 */
function defangIPaddr($address) {
    return str_replace('.', '[.]', $address);
}
