<?php
class MaxQueue {
    private $deque = [];
    private $enque = [];

    /**
     */
    function __construct() {
    }

    /**
     * @return Integer
     */
    function max_value() {
        if (!$this->deque) {
            return -1; 
        }
        return $this->deque[0];
    }

    /**
     * @param Integer $value
     * @return NULL
     */
    function push_back($value) {
        while ($this->deque && end($this->deque) < $value) {
            array_pop($this->deque);
        }
        $this->deque[] = $value;
        $this->enque[] = $value;
    }

    /**
     * @return Integer
     */
    function pop_front() {
        if (!$this->enque) {
            return -1; 
        }

        $res = array_shift($this->enque);
        if ($res == $this->deque[0]) {
            array_shift($this->deque);
        }
        return $res;
    }
}

//Your MaxQueue object will be instantiated and called as such:
$obj = new MaxQueue();
$obj->push_back(1);
$obj->push_back(2);
$ret_1 = $obj->max_value();
//$obj->push_back(4);
//$obj->push_back(6);
//$obj->push_back(7);
//$obj->push_back(5);
//$obj->push_back(3);
$ret_3 = $obj->pop_front();
var_dump($ret_3);
die;
