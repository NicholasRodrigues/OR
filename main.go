package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

type SimplexSolver struct {
	objective   []float64
	rows        [][]float64
	constraints []float64
	mu          sync.Mutex
}

func NewSimplexSolver(funcCoefficients []float64) *SimplexSolver {
	return &SimplexSolver{
		objective:   funcCoefficients,
		rows:        [][]float64{},
		constraints: []float64{},
	}
}

func (s *SimplexSolver) AddConstraint(coefficients []float64, value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	row := append([]float64{}, coefficients...)
	s.rows = append(s.rows, row)
	s.constraints = append(s.constraints, value)
}

func (s *SimplexSolver) getPivotCol() int {
	low := 0.0
	pivotIdx := 0
	for i := 0; i < len(s.objective)-1; i++ {
		if s.objective[i] < low {
			low = s.objective[i]
			pivotIdx = i
		}
	}
	return pivotIdx
}

func (s *SimplexSolver) getPivotRow(colIdx int) (int, error) {
	lastCol := make([]float64, len(s.rows))
	pivotCol := make([]float64, len(s.rows))

	var wg sync.WaitGroup
	wg.Add(len(s.rows))

	for i := 0; i < len(s.rows); i++ {
		go func(i int) {
			defer wg.Done()
			lastCol[i] = s.rows[i][len(s.rows[i])-1]
			pivotCol[i] = s.rows[i][colIdx]
		}(i)
	}

	wg.Wait()

	minRatio := math.Inf(1)
	minRatioIdx := -1

	for i := range lastCol {
		ratio := math.Inf(1)
		if pivotCol[i] != 0 {
			ratio = lastCol[i] / pivotCol[i]
		}
		if ratio >= 0 && ratio < minRatio {
			minRatio = ratio
			minRatioIdx = i
		}
	}

	if minRatioIdx == -1 {
		return -1, errors.New("no non-negative ratios, problem doesn't have a solution")
	}
	return minRatioIdx, nil
}

func (s *SimplexSolver) pivot(pivotRowIdx, pivotColIdx int) {
	pivotVal := s.rows[pivotRowIdx][pivotColIdx]
	for i := range s.rows[pivotRowIdx] {
		s.rows[pivotRowIdx][i] /= pivotVal
	}

	var wg sync.WaitGroup
	wg.Add(len(s.rows) - 1)

	for i := range s.rows {
		if i == pivotRowIdx {
			continue
		}
		go func(i int) {
			defer wg.Done()
			mul := s.rows[i][pivotColIdx]
			for j := range s.rows[i] {
				s.rows[i][j] -= mul * s.rows[pivotRowIdx][j]
			}
		}(i)
	}

	wg.Wait()

	mul := s.objective[pivotColIdx]
	for i := range s.objective {
		s.objective[i] -= mul * s.rows[pivotRowIdx][i]
	}
}

func (s *SimplexSolver) shouldPivot() bool {
	minVal := s.objective[0]
	for _, val := range s.objective[:len(s.objective)-1] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal < 0
}

func (s *SimplexSolver) addSlackVariables() {
	var wg sync.WaitGroup
	for i := range s.rows {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.mu.Lock()
			defer s.mu.Unlock()
			s.objective = append(s.objective, 0)
			basicCols := make([]float64, len(s.rows))
			basicCols[i] = 1
			basicCols = append(basicCols, s.constraints[i])
			s.rows[i] = append(s.rows[i], basicCols...)
		}(i)
	}
	wg.Wait()
	s.objective = append(s.objective, 0)
}

func (s *SimplexSolver) getSolutionFromTableau() ([]float64, float64) {
	cols := make([][]float64, len(s.rows[0]))

	var wg sync.WaitGroup
	for colI := range cols {
		wg.Add(1)
		go func(colI int) {
			defer wg.Done()
			col := make([]float64, len(s.rows))
			for rowI := range s.rows {
				col[rowI] = s.rows[rowI][colI]
			}
			s.mu.Lock()
			defer s.mu.Unlock()
			cols[colI] = col
		}(colI)
	}
	wg.Wait()

	results := make([]float64, len(cols)-1)
	for i := 0; i < len(cols)-1; i++ {
		if countZeros(cols[i]) == len(cols[i])-1 && containsOne(cols[i]) {
			results[i] = cols[len(cols)-1][indexOfOne(cols[i])]
		} else {
			results[i] = 0
		}
	}
	return results, s.objective[len(s.objective)-1]
}

func (s *SimplexSolver) Solve() {
	s.addSlackVariables()

	for s.shouldPivot() {
		pivotCol := s.getPivotCol()
		pivotRow, err := s.getPivotRow(pivotCol)
		if err != nil {
			fmt.Println(err)
			return
		}
		s.pivot(pivotRow, pivotCol)
	}
}

func countZeros(col []float64) int {
	count := 0
	for _, val := range col {
		if val == 0 {
			count++
		}
	}
	return count
}

func containsOne(col []float64) bool {
	for _, val := range col {
		if val == 1 {
			return true
		}
	}
	return false
}

func indexOfOne(col []float64) int {
	for i, val := range col {
		if val == 1 {
			return i
		}
	}
	return -1
}

func main() {
	testSimplexSolver1()
	testSimplexSolver2()
}
