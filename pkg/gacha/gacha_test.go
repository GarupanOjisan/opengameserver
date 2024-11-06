package gacha

import (
	"testing"
)

func TestWeightedRandomDistribution(t *testing.T) {
	items := []*Item{
		{ID: "1", Name: "Item1", Weight: 50},
		{ID: "2", Name: "Item2", Weight: 30},
		{ID: "3", Name: "Item3", Weight: 20},
	}

	// Number of trials
	trials := 1000000

	// Map to hold the count of each item's occurrences
	counts := make(map[string]int)

	var totalWeight int64
	for _, item := range items {
		totalWeight += item.Weight
	}

	// Execute weightedRandom function multiple times and aggregate occurrence counts
	for i := 0; i < trials; i++ {
		item := weightedRandom(items)
		if item != nil {
			counts[item.ID]++
		}
	}

	// Variables needed for Chi-squared statistic calculation
	var chiSquare float64

	// Compare expected and observed values for each item
	for _, item := range items {
		expected := float64(trials) * float64(item.Weight) / float64(totalWeight)
		observed := float64(counts[item.ID])
		diff := observed - expected
		chiSquare += diff * diff / expected

		t.Logf("Item ID: %s, Expected: %.2f, Observed: %.2f, Occurrence Rate: %.4f%%", item.ID, expected, observed, (observed/float64(trials))*100)
	}

	// Degrees of freedom is the number of item types minus one
	degreesOfFreedom := len(items) - 1

	// Set significance level (alpha) for Chi-squared test
	alpha := 0.01

	// Calculate critical value from Chi-squared distribution (using a statistics library)
	criticalValue := getChiSquareCriticalValue(degreesOfFreedom, alpha)

	t.Logf("Chi-squared Statistic: %.4f, Critical Value: %.4f, Degrees of Freedom: %d", chiSquare, criticalValue, degreesOfFreedom)

	// If Chi-squared statistic exceeds critical value, the distribution differs from expected
	if chiSquare > criticalValue {
		t.Errorf("Chi-squared test failed. There is a statistically significant difference between the observed and expected distributions.")
	} else {
		t.Logf("Chi-squared test passed. The distribution matches the expected values.")
	}
}

// Function to get critical value from Chi-squared distribution
func getChiSquareCriticalValue(degreesOfFreedom int, alpha float64) float64 {
	// Simplified: hard-code critical values for common degrees of freedom and significance levels
	criticalValues := map[int]map[float64]float64{
		1: {
			0.10: 2.7055,
			0.05: 3.8415,
			0.01: 6.6349,
		},
		2: {
			0.10: 4.6052,
			0.05: 5.9915,
			0.01: 9.2103,
		},
		3: {
			0.10: 6.2514,
			0.05: 7.8147,
			0.01: 11.3449,
		},
		// Add more as needed
	}

	if values, ok := criticalValues[degreesOfFreedom]; ok {
		if cv, ok := values[alpha]; ok {
			return cv
		}
	}

	// Return 0 if undefined (you may want to add error handling)
	return 0.0
}

func TestService_Execute(t *testing.T) {
	// Register a gacha
	items := []*Item{
		{ID: "1", Name: "Item1", Weight: 50},
		{ID: "2", Name: "Item2", Weight: 30},
		{ID: "3", Name: "Item3", Weight: 20},
	}

	// Execute the gacha
	n := 10
	result, err := Execute(items, n)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Output the result
	for i, item := range result {
		t.Logf("Result %d: %+v", i+1, item)
	}
}
