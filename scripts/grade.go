package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type TestResult struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"`
	Output  string    `json:"Output"`
}

type GradingResult struct {
	TestName string `json:"test_name"`
	Points   int    `json:"points"`
	MaxPoints int   `json:"max_points"`
	Status   string `json:"status"`
	Output   string `json:"output"`
}

func main() {
	// Define test suites and their point values
	testSuites := map[string]int{
		"TestFIFOCache": 10,
		"TestLRUCache":  10,
		"TestLFUCache":  10,
		"TestTTLCache":  10,
	}

	var results []TestResult
	var gradingResults []GradingResult

	// Run each test suite
	for testName, maxPoints := range testSuites {
		fmt.Printf("Running %s...\n", testName)
		
		// Run the specific test
		cmd := exec.Command("go", "test", "./tests", "-run", testName, "-v", "-json")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("  Warning: Error running %s: %v\n", testName, err)
		}
		
		// Parse test results
		lines := strings.Split(string(output), "\n")
		var testResults []TestResult
		var passed bool
		var testOutput strings.Builder
		
		for _, line := range lines {
			if line == "" {
				continue
			}
			
			var result TestResult
			if err := json.Unmarshal([]byte(line), &result); err == nil {
				testResults = append(testResults, result)
				results = append(results, result)
				
				if result.Action == "pass" && strings.Contains(result.Test, testName) {
					passed = true
				}
				
				if result.Output != "" {
					testOutput.WriteString(result.Output)
				}
			}
		}
		
		// Determine points based on test results
		points := 0
		status := "FAIL"
		if passed {
			points = maxPoints
			status = "PASS"
		}
		
		gradingResults = append(gradingResults, GradingResult{
			TestName:  testName,
			Points:    points,
			MaxPoints: maxPoints,
			Status:    status,
			Output:    testOutput.String(),
		})
		
		fmt.Printf("  %s: %d/%d points\n", testName, points, maxPoints)
	}
	
	// Calculate total score
	totalPoints := 0
	totalMaxPoints := 0
	for _, result := range gradingResults {
		totalPoints += result.Points
		totalMaxPoints += result.MaxPoints
	}
	
	// Write results to files
	writeTestResults(results)
	writeGradingSummary(gradingResults, totalPoints, totalMaxPoints)
	
	fmt.Printf("\n=== FINAL SCORE ===\n")
	fmt.Printf("Total: %d/%d points (%.1f%%)\n", totalPoints, totalMaxPoints, float64(totalPoints)/float64(totalMaxPoints)*100)
}

func writeTestResults(results []TestResult) {
	file, err := os.Create("test-results.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	for _, result := range results {
		encoder.Encode(result)
	}
}

func writeGradingSummary(results []GradingResult, totalPoints, totalMaxPoints int) {
	file, err := os.Create("grading-summary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	fmt.Fprintf(file, "=== CACHE STRATEGY GRADING SUMMARY ===\n\n")
	
	for _, result := range results {
		fmt.Fprintf(file, "%s: %s (%d/%d points)\n", 
			result.TestName, result.Status, result.Points, result.MaxPoints)
		if result.Output != "" {
			fmt.Fprintf(file, "  Output: %s\n", strings.TrimSpace(result.Output))
		}
		fmt.Fprintf(file, "\n")
	}
	
	fmt.Fprintf(file, "=== FINAL SCORE ===\n")
	fmt.Fprintf(file, "Total: %d/%d points (%.1f%%)\n", 
		totalPoints, totalMaxPoints, float64(totalPoints)/float64(totalMaxPoints)*100)
} 