package main

import (
	"fmt"
	"testing"
)

func TestConsumes(t *testing.T) {
	cuboidTest1 := cuboid{true, -5, 5, -5, 5, -5, 5}
	cuboidTest2 := cuboid{true, -5, 5, -5, 5, -5, 5}

	result := cuboidTest1.consumes(cuboidTest2)
	if result == false {
		t.Error("Result should be false, got", result)
	}
}

func TestDeconstructBy(t *testing.T) {
	cuboidTest1 := cuboid{true, -5, 5, -5, 5, -5, 5}
	cuboidTest2 := cuboid{true, -1, 1, -1, 1, -1, 1}
	cutout, remainingCuboids := cuboidTest1.deconstructBy(cuboidTest2)
	fmt.Println(cutout, remainingCuboids)
	if cutout != cuboidTest2 {
		t.Log(cutout, cuboidTest2)
		t.Error("Cutout should be the same as the deconstructing cube")
	}
	if len(remainingCuboids) != 6 {
		t.Error("Should have 6 new cuboids.")
	}
}

func TestDeconstructBy2(t *testing.T) {
	cuboidTest1 := cuboid{true, 0, 10, 0, 10, 0, 10}
	cuboidTest2 := cuboid{true, -5, 5, -5, 5, -5, 5}
	cutout, remainingCuboids := cuboidTest1.deconstructBy(cuboidTest2)
	fmt.Println(cutout, remainingCuboids)
	desiredCutout := cuboid{true, 0, 5, 0, 5, 0, 5}
	if cutout != desiredCutout {
		t.Error("Cutout should be Just da corner")
	}
	if len(remainingCuboids) != 3 {
		t.Error("Should have 3 new cuboids.")
	}
}
