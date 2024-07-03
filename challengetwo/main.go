package main

import (
	"fmt"
	"math"
	"sort"
)

func minimizeMaxRisk(dataCenters []int, fragments int) int {
	// Sort data centers by their base risk factor in ascending order
	sort.Ints(dataCenters)

	// Binary search for the minimized maximum risk
	low, high := 1, int(math.Pow(10, 18)) // Adjust upper limit for practical binary search
	var result int

	for low <= high {
		mid := (low + high) / 2
		if checkValidDistribution(dataCenters, fragments, mid) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}

func checkValidDistribution(dataCenters []int, fragments, maxRiskThreshold int) bool {
	totalRisk := 0
	for _, baseRiskFactor := range dataCenters {
		risk := calculateTotalRisk(baseRiskFactor, fragments)
		totalRisk += risk
		if totalRisk > maxRiskThreshold {
			return false
		}
	}
	return true
}

func calculateTotalRisk(baseRiskFactor, fragments int) int {
	return int(math.Pow(float64(baseRiskFactor), float64(fragments)))
}

// Secure Data Fragment Allocation Scenario
func main() {
	dataCenters := []int{10, 20, 30}
	fragments := 5

	// Find the minimized maximum risk achievable
	minimizedMaxRisk := minimizeMaxRisk(dataCenters, fragments)

	fmt.Printf("Minimized maximum risk: %d\n", minimizedMaxRisk)
}
