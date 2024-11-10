package test

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
	"time"
)

func RunTestOne[T any](testCase GenericTestCaseBoolOutput[T], testFunc func(T) bool) bool {
	result := testFunc(testCase.Input)
	if result == testCase.Expected {
		fmt.Println("Test Passed!")
		return true
	} else {
		fmt.Printf("Test Failed! Expected: %v, Got: %v\n", testCase.Expected, result)
		return false
	}
}

// Define a function that takes a list of GenericTestCaseBoolOutput as a parameter
func RunTestCases[T any](testCases []GenericTestCaseBoolOutput[T], testFunc func(T) bool) {
	// Start time tracking
	startTime := time.Now()

	// Track memory before tests
	var memStart runtime.MemStats
	runtime.ReadMemStats(&memStart)

	successCount := 0
	failureCount := 0

	// Create color formats for messages
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()
	reset := color.New(color.Reset).SprintFunc()

	// Track failed test cases to display at the end
	var failedTestCases []string

	// Iterate through the test cases and run each one
	for i, testCase := range testCases {
		// Start time for each test case
		caseStartTime := time.Now()

		// Run the test
		result := testFunc(testCase.Input)

		// Track test case result and success/failure count
		if result == testCase.Expected {
			successCount++
			fmt.Printf("%sTest Case %d Passed!%s\n", green(""), i+1, reset())
		} else {
			failureCount++
			failedTestCases = append(failedTestCases, fmt.Sprintf("Test Case %d: Expected %v, Got %v, Input: %v", i+1, testCase.Expected, result, testCase.Input))
			fmt.Printf("%sTest Case %d Failed: Expected %v, Got %v%s\n", red(""), i+1, testCase.Expected, result, reset())
			fmt.Printf("Input: %v\n", testCase.Input)
		}

		// Track time for this test case
		caseDuration := time.Since(caseStartTime)
		fmt.Printf("%sTest Case %d took %v%s\n", cyan(""), i+1, caseDuration, reset())
	}

	// End time tracking
	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)

	// Track memory after tests
	var memEnd runtime.MemStats
	runtime.ReadMemStats(&memEnd)
	memUsed := memEnd.Alloc - memStart.Alloc

	// Display metrics
	totalTests := len(testCases)
	fmt.Printf("\n%s%sSummary:%s %s%d/%d tests passed%s (%.2f%% success rate)\n", bold(""), green(""), reset(), bold(""), successCount, totalTests, reset(), float64(successCount)/float64(totalTests)*100)
	if failureCount > 0 {
		fmt.Printf("%s%d tests failed%s\n", red(""), failureCount, reset())
	} else {
		fmt.Printf("%sAll tests passed! Great job!%s\n", green(""), reset())
	}

	// Display time and memory metrics
	fmt.Printf("\n%sTotal Time Taken: %v%s\n", cyan(""), totalDuration, reset())
	fmt.Printf("%sTotal Memory Used: %v bytes%s\n", cyan(""), memUsed, reset())

	// If there are failed test cases, list them at the end
	if len(failedTestCases) > 0 {
		fmt.Printf("\n%sFailed Test Cases:%s\n", red(""), reset())
		for _, failedCase := range failedTestCases {
			fmt.Println(failedCase)
		}
	}
}
