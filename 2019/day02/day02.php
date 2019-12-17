<?php
include '../php_util/opcode_computer.php';

$input = rtrim(fgets(fopen("input.txt", "r")), "\n");
$originalInstructions = preg_split("/,/", $input);

$comp = new opcodeComputer;
$comp->originalInstructions = $originalInstructions;
$comp->instructions = $originalInstructions;

$comp->runInstructions();
