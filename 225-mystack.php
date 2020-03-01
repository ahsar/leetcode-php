<?php
/**
 * MyStack 
 * 
 * 双队列实现栈
 */
class MyStack {
    private $arr1 = [];
    private $arr2 = [];

    /**
     * Initialize your data structure here.
     */
    function __construct() {
    }

    /**
     * Push element x onto stack.
     * @param Integer $x
     * @return NULL
     */
    function push($x) {
        $this->arr1 = [$x];
        while ($this->arr2) {
            $tmp = array_shift($this->arr2);
            $this->arr1[] = $tmp;
        }
        $this->arr2 = $this->arr1;
    }

    /**
     * Removes the element on top of the stack and returns that element.
     * @return Integer
     */
    function pop() {
        return array_shift($this->arr1);
    }

    /**
     * Get the top element.
     * @return Integer
     */
    function top() {
        return $this->arr1[0];
    }

    /**
     * Returns whether the stack is empty.
     * @return Boolean
     */
    function empty() {
        return empty($this->arr1);
    }
}

 //Your MyStack object will be instantiated and called as such:
$obj = new MyStack();
$obj->push(1);
$obj->push(2);
$obj->push(3);
$obj->push(4);
//$obj->push(5);
$ret_2 = $obj->top();
print_r($ret_2);
echo "\n";
$ret_2 = $obj->pop();
print_r($ret_2);
//$ret_2 = $obj->pop();
//$ret_2 = $obj->pop();
//$ret_2 = $obj->pop();
//$ret_2 = $obj->pop();
//$ret_4 = $obj->empty();
