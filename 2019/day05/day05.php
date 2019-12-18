<?php
include '../php_util/opcode_computer.php';

$input = rtrim(fgets(fopen("input.txt", "r")), "\n");
$originalInstructions = array_map('intval', preg_split("/,/", $input));


echo "Part1:" . PHP_EOL;
$comp = new opcodeComputer($originalInstructions);
$comp->setInputs(array(1));
$comp->runInstructions();
var_dump($comp->outputs);
echo "Part2:" . PHP_EOL;
$comp = new opcodeComputer($originalInstructions);
$comp->setInputs(array(5));
$comp->runInstructions();
var_dump($comp->outputs);
