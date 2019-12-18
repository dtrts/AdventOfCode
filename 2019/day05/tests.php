<?php
include '../php_util/opcode_computer.php';

$comp = new opcodeComputer(array(3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8));
for ($i = 0; $i < 10; $i++) {
  echo "Test " . $i . PHP_EOL;
  $comp->resetInstructions();
  $comp->setInputs(array($i));
  $comp->runInstructions();
  if ($i == 8) {
    if ($comp->outputs[0] != 1) {
      var_dump($comp->outputs);
      echo "ERROR 8 :" . PHP_EOL;
    } else {
      echo "PASS" . PHP_EOL;
    }
  } else {
    if ($comp->outputs[0] != 0) {
      var_dump($comp->outputs);
      echo "ERROR " . $i . PHP_EOL;
    } else {
      echo "PASS" . PHP_EOL;
    }
  }
}
$comp->setInputs(array(2));
$comp->runInstructions();
