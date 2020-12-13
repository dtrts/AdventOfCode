package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input = inputTest
	inputSplit := strings.Split(input, "\n")
	earliestDepartTimestamp, _ := strconv.Atoi(inputSplit[0])
	serviceIds := strings.Split(inputSplit[1], ",")

	serviceIdsInt := []int{}
	for _, serviceId := range serviceIds {
		serviceIdInt := 0
		if serviceId == "x" {
			serviceId = "0"
		}
		serviceIdInt, _ = strconv.Atoi(serviceId)
		serviceIdsInt = append(serviceIdsInt, serviceIdInt)
	}

	fmt.Println(earliestDepartTimestamp, serviceIds, serviceIdsInt)

	departTimestamp := 0
	departServiceId := 0
	for departTimestampTmp := earliestDepartTimestamp; departTimestampTmp < earliestDepartTimestamp+1000; departTimestampTmp++ {
		for _, serviceIdInt := range serviceIdsInt {
			// fmt.Println(earliestDepartTimestamp, departTimestamp, serviceIdInt)
			if serviceIdInt > 0 && departTimestampTmp%serviceIdInt == 0 {
				fmt.Println(earliestDepartTimestamp, departTimestampTmp, serviceIdInt) //939 944 59
				departTimestamp = departTimestampTmp
				departServiceId = serviceIdInt
				break
			}
		}

		if departTimestamp > 0 {
			break
		}

	}

	fmt.Println("Part1:", (departTimestamp-earliestDepartTimestamp)*departServiceId)

	// 7,13,x,x,59,x,31,19 , 1068781
	// The earliest timestamp that matches the list 17,x,13,19 is 3417.
	// 67,7,59,61 first occurs at timestamp 754018.
	// 67,x,7,59,61 first occurs at timestamp 779210.
	// 67,7,x,59,61 first occurs at timestamp 1261476.
	// 1789,37,47,1889 first occurs at timestamp 1202161486.

	bases := []int{}
	testServiceIdsInt := []int{17, 0, 13, 19}
	for i, serviceId := range testServiceIdsInt {
		// if i == 0 {
		// 	continue
		// }
		if serviceId > 0 {
			bases = append(bases, combi(testServiceIdsInt[0], serviceId, i))
		}

	}

	fmt.Println(bases, LCM(bases[0], bases[1], bases[2:]...))

	minValue := 0
	runningProduct := 1
	for index, ServiceId := range serviceIdsInt {
		if ServiceId > 0 {

			for (minValue+index)%ServiceId != 0 {
				minValue += runningProduct
			}
			runningProduct *= ServiceId
		}
	}

	fmt.Println(serviceIdsInt, minValue, runningProduct)

}

func combi(baseService int, nextService int, diff int) int {
	fmt.Println("Start combi", baseService, nextService, diff)
	for times := 1; times <= nextService; times++ {
		// fmt.Println("loop", times, baseService*times, baseService*times+diff, ((baseService*times)+diff)%nextService)
		if ((baseService*times)+diff)%nextService == 0 {
			fmt.Println("break", times)
			return baseService * times
		}
	}

	return 0

}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := (a * b) / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
