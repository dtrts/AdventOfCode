<?php
class opcodeComputer
{
  public $originalInstructions;
  public $instructions;

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

  function runInstructions()
  {
    // echo "Operations Started" . PHP_EOL;
    $opLoc = 0;
    while ($opLoc < count($this->instructions)) {

      switch ($this->instructions[$opLoc]) {
        case 99:
          // echo "Operations Complete" . PHP_EOL;
          return;

        case 1:
          $this->instructions[$this->instructions[$opLoc + 3]] = $this->instructions[$this->instructions[$opLoc + 1]] + $this->instructions[$this->instructions[$opLoc + 2]];
          $opLoc += 4;
          break;

        case 2:
          $this->instructions[$this->instructions[$opLoc + 3]] = $this->instructions[$this->instructions[$opLoc + 1]] * $this->instructions[$this->instructions[$opLoc + 2]];
          $opLoc += 4;
          break;

        default:
          echo "Oh fex" . PHP_EOL;
          return;
      }
    }
  }
}
