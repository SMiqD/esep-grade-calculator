# ESEP Grade Calculator

A Go library for calculating course grades from assignments, exams, and essays.
There is no executable application or external dependency.

## Requirements and setup

Install Go 1.24.4 or later: https://go.dev/doc/install

```sh
git clone https://github.com/SMiqD/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
go version
go test -v -count=1 ./...
```

Run commands from `grade-calculator/`, which contains `go.mod`. The directory
uses a hyphen, not an underscore. No dependency download step is necessary:
the module only uses the standard library.

## Files

- `README.MD`: setup, usage, testing, and grading rules.
- `grade-calculator/go.mod`: module `esep/grade-calculator`, Go version 1.24.4.
- `grade-calculator/grade_calculator.go`: grade types, calculator, category
  averages, weighting, and conversion to a final grade.
- `grade-calculator/grade_calculator_test.go`: Go unit tests using `testing`.

The package name is `esepunittests`; it differs from the module name.

## Usage

```go
import calculator "esep/grade-calculator"

gc := calculator.NewGradeCalculator()
gc.AddGrade("assignment 1", 90, calculator.Assignment)
gc.AddGrade("exam 1", 80, calculator.Exam)
gc.AddGrade("essay 1", 100, calculator.Essay)
letter := gc.GetFinalGrade() // B after Part 3 fixes (weighted score 88)
```

This import is for code that can resolve the local module, such as through a
Go workspace or a `replace` directive. It is not a published GitHub module path
suitable for `go get`.

## Grading requirements

Average scores in each category, then apply these weights:

| Category | Weight |
| --- | --- |
| Assignments | 50% |
| Exams | 35% |
| Essays | 15% |

Use A for scores >= 90, B for >= 80, C for >= 70, D for >= 60, and F otherwise.
Passing means C or better. Keep fractional averages until the final comparison;
do not round up across a letter boundary.

Empty-category behavior, invalid inputs, and configurable weights need product
clarification. The starter implementation does not safely handle empty
categories. Later assignment commits document any explicit implementation policy.

## Tests and coverage

```sh
go test -v -count=1 ./...
go test -count=1 -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
go vet ./...
```

Go discovers `Test...` functions in files ending in `_test.go`. Assertions use
comparisons and `t.Errorf` / `t.Fatalf`. Coverage measures executed statements,
not whether their results were checked correctly.

## Starting-project status

The provided tests cover expected A, B, and F outcomes. The starter code adds
slice indexes instead of scores and reads exams for the essay average. The F
test supplies scores that should produce A and can pass for the wrong reason.
Part 3 corrects the existing tests and code; Part 4 adds tests. See commit
history for the separate assignment snapshots.

## References

- https://go.dev/doc/tutorial/add-a-test
- https://go.dev/blog/cover
- https://go.dev/ref/spec#For_statements

## Completed assignment policies and Pass/Fail mode

Empty categories contribute zero and their weights are not redistributed.
Unknown grade types are ignored, preserving the starter behavior. Scores are
expected to be integers in 0..100; validation and configurable weights remain
outside this exercise's requirements.

`NewGradeCalculator()` still produces letter grades. For Pass/Fail:

```go
gc := calculator.NewGradeCalculator(calculator.PassFail)
gc.AddGrade("assignment", 70, calculator.Assignment)
gc.AddGrade("exam", 70, calculator.Exam)
gc.AddGrade("essay", 70, calculator.Essay)
result := gc.GetFinalGrade() // "Pass"
```

`NewGradeCalculator(calculator.LetterGrades)` explicitly selects letter grades.
The optional constructor mode uses the first supplied value; values other than
`PassFail` use letter evaluation. `AddGrade` and `GetFinalGrade` retain their
original signatures. Pass means a weighted score >= 70; lower scores return
"Fail". Fractional scores are not rounded up.

The calculator stores one persistent list of grades. Temporary category slices
are used when computing averages. Existing tests were unchanged for Part 6.

Assignment checkpoints:
- Part 3: corrected the three original tests/code; statement coverage 85.2%.
- Part 4: added tests; statement coverage 100.0%.
- Part 5: in a temporary copy, commented out failure assertions; coverage stayed
  100.0%. Assertions remain enabled in this repository.
- Part 6: changed internal storage only; unchanged tests and coverage 100.0%.
- Part 7.1: added new tests before implementation (intentional red snapshot).
- Part 7.2: added constructor/state; new behavioral tests still failed.
- Part 7.3: implemented Pass/Fail; all tests pass, coverage 100.0%.
