package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{grades: make([]Grade, 0)}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{Name: name, Grade: grade, Type: gradeType})
	}
}

// Only one list is stored; filtered slices are temporary calculation inputs.
func (gc *GradeCalculator) gradesOfType(gradeType GradeType) []Grade {
	var selected []Grade
	for _, grade := range gc.grades {
		if grade.Type == gradeType {
			selected = append(selected, grade)
		}
	}
	return selected
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	assignment_average := computeAverage(gc.gradesOfType(Assignment))
	exam_average := computeAverage(gc.gradesOfType(Exam))
	essay_average := computeAverage(gc.gradesOfType(Essay))

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return weighted_grade
}

// Empty categories contribute zero without redistributing their fixed weight.
func computeAverage(grades []Grade) float64 {
	if len(grades) == 0 {
		return 0
	}
	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return float64(sum) / float64(len(grades))
}
