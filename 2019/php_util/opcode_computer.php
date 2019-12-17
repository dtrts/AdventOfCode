<?php
class opcodeComputer
{
  public $originalInstructions;
  public $instructions;

  function runInstructions()
  {
    $opLoc = 0;
    while ($opLoc < count($this->instructions)) {


      switch ($this->instructions[$opLoc]) {
        case 99:
          echo "Operations Complete" . PHP_EOL;
          return;


        default:
          echo "Oh fex" . PHP_EOL;
          return;
      }
    }
  }
}
