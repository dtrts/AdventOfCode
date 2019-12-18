<?php
class opcodeComputer
{
  public $originalInstructions;
  public $instructions;
  public $inputs;
  public $outputs;


  function __construct($instructions)
  {
    $this->originalInstructions = $instructions;
    $this->instructions = $instructions;
  }
  function resetInstructions()
  {
    $this->instructions = $this->originalInstructions;
  }

  function setNounAndVerb($noun, $verb)
  {
    $this->instructions[1] = $noun;
    $this->instructions[2] = $verb;
  }
  function setInputs($inputs)
  {
    $this->inputs = $inputs;
  }
  function set($loc, $value)
  {
    $this->instructions[$this->instructions[$loc]] = $value;
  }
  function get($loc, $mode)
  {
    if ($mode == 0) {
      return $this->instructions[$this->instructions[$loc]];
    } else if ($mode == 1) {
      return $this->instructions[$loc];
    }
  }

  function runInstructions()
  {
    // echo "Operations Started" . PHP_EOL;
    $opLoc = 0;

    while ($opLoc < count($this->instructions)) {

      $operation = $this->instructions[$opLoc];
      $mode1 = $operation / 100 % 10;
      $mode2 = $operation / 1000 % 10;
      $mode3 = $operation / 10000 % 10;
      $operation = $operation % 100;


      switch ($operation) {
        case 99:
          // echo "Operations Complete" . PHP_EOL;
          return;

        case 1:
          $this->set(
            $opLoc + 3,
            $this->get($opLoc + 1, $mode1) + $this->get($opLoc + 2, $mode2)
          );
          $opLoc += 4;
          break;

        case 2:
          $this->set(
            $opLoc + 3,
            $this->get($opLoc + 1, $mode1) * $this->get($opLoc + 2, $mode2)
          );
          $opLoc += 4;
          break;

        case 3:
          $input = array_shift($this->inputs);
          echo "Reading Input: " . $input . PHP_EOL;
          $this->set($opLoc + 1, $input);
          $opLoc += 2;
          break;
        case 4:
          $output = $this->get($opLoc + 1, $mode1);
          echo "Generating output: " . $output . PHP_EOL;
          $this->outputs[] = $output;
          $opLoc += 2;
          break;
        default:
          echo "Oh fex" . PHP_EOL;
          return;
      }
    }
  }
}
