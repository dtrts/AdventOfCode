<?php
include '../php_util/opcode_computer.php';

$input = rtrim(fgets(fopen("input.txt", "r")), "\n");
$originalInstructions = array_map('intval', preg_split("/,/", $input));

$comp = new opcodeComputer($originalInstructions);
$comp->setInputs(array(1));
$comp->runInstructions();
