# OR
---
# GOLANG RULES!

# Simplex Solver - README

## Overview

This project is a simplified implementation of the Simplex algorithm in Go. It solves linear programming problems where the objective is to maximize or minimize a linear function subject to a set of linear inequalities (constraints). This implementation is specifically tailored to handle problems that are already in canonical form with all constraints as "≤" and non-negative variables.
`Plus, golang is the best ever known language!`

## What This Solver Can Do

- **Maximize/Minimize Linear Objectives**: The solver can find the optimal solution for linear programming problems with linear objective functions.
- **Handle Basic Constraints**: It supports linear inequalities of the form "≤" and assumes non-negative variables.
- **Add Slack Variables**: Slack variables are added automatically to convert inequalities into equations.
- **Run Concurrently**: The implementation utilizes Go’s concurrency features, allowing some parts of the algorithm to run in parallel.

## What This Solver Cannot Do (Yet)

- **Infeasibility/Unboundedness Detection**: The solver does not detect when a problem is infeasible (no solution exists) or unbounded (the solution can grow indefinitely).
- **Handling of Non-Standard Forms**: It does not support problems with "≥" or "=" constraints or those that require artificial variables.
- **Degeneracy and Cycling**: There is no anti-cycling mechanism to handle degeneracy, which can cause the algorithm to loop infinitely without finding a solution.
- **Post-Optimality Analysis**: No sensitivity analysis or dual pricing is implemented.

## Constraints

- **Canonical Form**: All constraints must be "≤", and all variables must be non-negative.
- **Input Requirements**: Coefficients must be provided as slices of floats, and constraints must be specified as linear inequalities with a corresponding value.

## Next Steps

To make the solver fully functional for more general linear programming problems, the following features will be added:

1. **Infeasibility/Unboundedness Detection**: Implement checks to identify and handle infeasible or unbounded problems.
2. **Two-Phase Simplex Method**: Add support for problems with "≥" or "=" constraints using a two-phase method.
3. **Anti-Cycling Mechanism**: Implement Bland's rule or another method to prevent cycling.
4. **Floating Point Precision Handling**: Improve handling of floating-point arithmetic to reduce precision errors.
5. **Post-Optimality Analysis**: Add support for sensitivity analysis and dual variables.
6. **Better Error Handling**: Enhance the robustness of the solver by adding comprehensive error handling and edge case management.
