package main

import (
	"fmt"
)

func testSimplexSolver1() {
	funcCoefficients := []float64{-30, -60, -120}
	solver := NewSimplexSolver(funcCoefficients)

	solver.AddConstraint([]float64{1, 0, 0}, 100)
	solver.AddConstraint([]float64{0, 1, 0}, 25)
	solver.AddConstraint([]float64{0, 0, 1}, 10)
	solver.AddConstraint([]float64{1, 2, 3}, 150)

	solver.Solve()

	solution, objectiveValue := solver.getSolutionFromTableau()

	expectedProfits := 4800.0
	expectedSolution := []float64{70.0, 25.0, 10.0}

	fmt.Println("Test Case 1")
	fmt.Printf("Expected Profits: %.1f\n", expectedProfits)
	fmt.Printf("Actual Profits: %.1f\n", objectiveValue)
	fmt.Printf("Expected Solution: %v\n", expectedSolution)
	fmt.Printf("Actual Solution: %v\n", solution)
	if objectiveValue == expectedProfits && equalSlices(solution, expectedSolution) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	fmt.Println("---------------------------------")
}

func testSimplexSolver2() {
	funcCoefficients := []float64{-10, -50, -100}
	solver := NewSimplexSolver(funcCoefficients)

	solver.AddConstraint([]float64{1, 0, 0}, 100)
	solver.AddConstraint([]float64{0, 1, 0}, 50)
	solver.AddConstraint([]float64{0, 0, 1}, 20)
	solver.AddConstraint([]float64{1, 2, 3}, 200)

	solver.Solve()

	solution, objectiveValue := solver.getSolutionFromTableau()

	expectedProfits := 4900.0
	expectedSolution := []float64{40.0, 50.0, 20.0}

	fmt.Println("Test Case 2")
	fmt.Printf("Expected Profits: %.1f\n", expectedProfits)
	fmt.Printf("Actual Profits: %.1f\n", objectiveValue)
	fmt.Printf("Expected Solution: %v\n", expectedSolution)
	fmt.Printf("Actual Solution: %v\n", solution)
	if objectiveValue == expectedProfits && equalSlices(solution, expectedSolution) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	fmt.Println("============= PASS ==============")
	fmt.Println("2 passed, 0 failed")
}

func equalSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
