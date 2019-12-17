<?php
// echo getcwd();
$input = fopen("input.txt", "r") or die("Unable to open file");
$orbitPairs = array();
while ($line = fgets($input)) {
  $orbitPairs[] = preg_split("/\)/", rtrim($line, "\n"));
}

$orbitingObjects = array(new OrbitingObject("COM", 0, null));
$i = 0;
while ($i < count($orbitingObjects)) {

  $currentObject = $orbitingObjects[$i];
  $currentLevel = $orbitingObjects[$i]->orbitLevel + 1;
  $orbitedBy = findOrbitedBy($orbitPairs, $currentObject->objectName);
  foreach ($orbitedBy as $oby) {
    $orbitingObjects[] = new OrbitingObject($oby, $currentLevel, $currentObject);
  }
  $i++;
}

// Get sum of levels to get all orbits and sub orbits. 
$part1 = 0;
foreach ($orbitingObjects as $orb) {

  $part1 = $part1 + $orb->orbitLevel;
}


echo $part1 . PHP_EOL;


// Find object you are orbiting. 
$isItMe = pathDownTheWell($orbitingObjects, findObjByName($orbitingObjects, "YOU"));
// Find object you are orbiting. 
$santa = pathDownTheWell($orbitingObjects, findObjByName($orbitingObjects, "SAN"));

$i = 0;
$j = 0;
foreach ($isItMe as $objVisit) {
  if (array_search($objVisit, $santa)) {
    $j = array_search($objVisit, $santa);
    break;
  }
  $i++;
}

echo $i + $j - 2 . PHP_EOL;





function pathDownTheWell($orbitingObjects, $obj)
{
  if ($obj->objectName == "COM") {
    return array("COM");
  }

  return array_merge(array($obj->objectName), pathDownTheWell($orbitingObjects, $obj->orbits));
}


function findObjByName($orbitingObjects, $objName)
{
  foreach ($orbitingObjects as $orb) {
    if ($orb->objectName == $objName) {
      return $orb;
    }
  }
}

function findOrbitedBy($orbitPairs, $objectName)
{
  $orbitedBy = array();
  foreach ($orbitPairs as $orbitPair) {
    if ($orbitPair[0] === $objectName) {
      $orbitedBy[] = $orbitPair[1];
    }
  }
  return $orbitedBy;
}

class OrbitingObject
{
  public $objectName;
  public $orbitLevel;
  public $orbits;

  function __construct($objectName, $orbitLevel, $orbits)
  {
    $this->objectName = $objectName;
    $this->orbitLevel = $orbitLevel;
    $this->orbits = $orbits;
  }
}
