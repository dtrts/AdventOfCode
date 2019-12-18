<?php
include '../php_util/opcode_computer.php';

$input = rtrim(fgets(fopen("input.txt", "r")), "\n");
$originalInstructions = array_map('intval', preg_split("/,/", $input));
// var_dump($originalInstructions);


$comp = new opcodeComputer($originalInstructions);
$comp->setNounAndVerb(12, 2);
$comp->runInstructions();
echo "Part 1: " . $comp->instructions[0] . PHP_EOL;
// var_dump($comp->instructions[0]);

for ($noun = 0; $noun <= 99; $noun++) {
  for ($verb = 0; $verb <= 99; $verb++) {
    $comp->resetInstructions();
    $comp->setNounAndVerb($noun, $verb);
    $comp->runInstructions();
    if ($comp->instructions[0] == 19690720) {
      echo "Part 2: " . (100 * $noun + $verb) . PHP_EOL;
    }
  }
}
