package esepunittests

import (
	"fmt"
	"testing"
)

func TestPassFailEvaluation(t *testing.T) {
	for _, tc := range []struct {
		score int
		want  string
	}{
		{0, "Fail"}, {59, "Fail"}, {60, "Fail"}, {69, "Fail"},
		{70, "Pass"}, {80, "Pass"}, {90, "Pass"}, {100, "Pass"},
	} {
		t.Run(fmt.Sprint(tc.score), func(t *testing.T) {
			gc := NewGradeCalculator(PassFail)
			gc.AddGrade("assignment", tc.score, Assignment)
			gc.AddGrade("exam", tc.score, Exam)
			gc.AddGrade("essay", tc.score, Essay)
			if got := gc.GetFinalGrade(); got != tc.want {
				t.Errorf("score %d: got %q, want %q", tc.score, got, tc.want)
			}
		})
	}
}

func TestPassFailWeightedBoundary(t *testing.T) {
	for _, tc := range []struct {
		assignment, exam, essay int
		want                    string
	}{
		{70, 70, 69, "Fail"},  // 69.85: no rounding up to Pass
		{60, 80, 80, "Pass"},  // weighted score exactly 70
		{100, 0, 100, "Fail"}, // weighted score 65
	} {
		gc := NewGradeCalculator(PassFail)
		gc.AddGrade("assignment", tc.assignment, Assignment)
		gc.AddGrade("exam", tc.exam, Exam)
		gc.AddGrade("essay", tc.essay, Essay)
		if got := gc.GetFinalGrade(); got != tc.want {
			t.Errorf("weighted case %+v: got %q", tc, got)
		}
	}
}

func TestPassFailEmptyCalculator(t *testing.T) {
	if got := NewGradeCalculator(PassFail).GetFinalGrade(); got != "Fail" {
		t.Errorf("empty calculator: got %q, want Fail", got)
	}
}

func TestExplicitLetterMode(t *testing.T) {
	gc := NewGradeCalculator(LetterGrades)
	gc.AddGrade("assignment", 70, Assignment)
	gc.AddGrade("exam", 70, Exam)
	gc.AddGrade("essay", 70, Essay)
	if got := gc.GetFinalGrade(); got != "C" {
		t.Errorf("explicit letter mode: got %q, want C", got)
	}
}
