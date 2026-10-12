package esepunittests

import (
	"fmt"
	"math"
	"testing"
)

// Public-behavior tests stay unchanged during the Part 6 storage refactor.
func TestLetterGradeBoundaries(t *testing.T) {
	for _, tc := range []struct {
		score int
		want  string
	}{
		{0, "F"}, {59, "F"}, {60, "D"}, {69, "D"},
		{70, "C"}, {79, "C"}, {80, "B"}, {89, "B"},
		{90, "A"}, {100, "A"},
	} {
		t.Run(fmt.Sprintf("%d_%s", tc.score, tc.want), func(t *testing.T) {
			gc := NewGradeCalculator()
			gc.AddGrade("assignment", tc.score, Assignment)
			gc.AddGrade("exam", tc.score, Exam)
			gc.AddGrade("essay", tc.score, Essay)
			if got := gc.GetFinalGrade(); got != tc.want {
				t.Errorf("score %d: got %q, want %q", tc.score, got, tc.want)
			}
		})
	}
}

func TestCategoryAveragesAndWeights(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment 1", 100, Assignment)
	gc.AddGrade("assignment 2", 61, Assignment)
	gc.AddGrade("exam", 100, Exam)
	gc.AddGrade("essay", 100, Essay)
	// 80.5*0.50 + 100*0.35 + 100*0.15 = 90.25.
	// Truncating the assignment average would incorrectly produce B.
	if got := gc.GetFinalGrade(); got != "A" {
		t.Errorf("fractional category average: got %q, want A", got)
	}

	gc = NewGradeCalculator()
	gc.AddGrade("assignment", 100, Assignment)
	gc.AddGrade("exam", 100, Exam)
	gc.AddGrade("essay", 0, Essay)
	// 100*0.50 + 100*0.35 + 0*0.15 = 85.
	if got := gc.GetFinalGrade(); got != "B" {
		t.Errorf("essay must have its own average: got %q, want B", got)
	}
}

func TestNoRoundingAcrossLetterBoundary(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 90, Assignment)
	gc.AddGrade("exam", 90, Exam)
	gc.AddGrade("essay", 89, Essay)
	// 89.85 must remain B.
	if got := gc.GetFinalGrade(); got != "B" {
		t.Errorf("got %q, want B for 89.85", got)
	}
}

func TestEmptyAndMissingCategories(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("empty calculator: got %q, want F", got)
	}
	gc.AddGrade("assignment", 100, Assignment)
	gc.AddGrade("exam", 100, Exam)
	if got := gc.GetFinalGrade(); got != "B" {
		t.Errorf("missing essay contributes zero: got %q, want B", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	for _, tc := range []struct {
		kind GradeType
		want string
	}{{Assignment, "assignment"}, {Exam, "exam"}, {Essay, "essay"}, {GradeType(99), ""}} {
		if got := tc.kind.String(); got != tc.want {
			t.Errorf("type %d: got %q, want %q", tc.kind, got, tc.want)
		}
	}
}

func TestUnknownGradeTypeIgnored(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("unknown", 100, GradeType(99))
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("unknown category: got %q, want F", got)
	}
}

func TestComputeAverage(t *testing.T) {
	for _, tc := range []struct {
		grades []Grade
		want   float64
	}{{nil, 0}, {[]Grade{{Grade: 100}, {Grade: 61}}, 80.5}} {
		if got := computeAverage(tc.grades); math.Abs(got-tc.want) > 1e-9 {
			t.Errorf("average: got %v, want %v", got, tc.want)
		}
	}
}
