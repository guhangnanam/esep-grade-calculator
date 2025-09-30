package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 30, Assignment)
	gradeCalculator.AddGrade("exam 1", 40, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 45, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
    expected_value := "C"

    gradeCalculator := NewGradeCalculator()
    gradeCalculator.AddGrade("assignment", 75, Assignment)
    gradeCalculator.AddGrade("exam", 70, Exam)
    gradeCalculator.AddGrade("essay", 72, Essay)

    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    }
}

func TestGetGradeD(t *testing.T) {
    expected_value := "D"

    gradeCalculator := NewGradeCalculator()
    gradeCalculator.AddGrade("assignment", 65, Assignment)
    gradeCalculator.AddGrade("exam", 60, Exam)
    gradeCalculator.AddGrade("essay", 62, Essay)

    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    }
}

func TestEmptyCategories(t *testing.T) {
    // Only add assignments, leave exams and essays empty
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 100, Assignment)

    grade := gc.GetFinalGrade()

    // Weighted: assignments = 100, exams = 0, essays = 0
    // Final = 100*0.5 + 0*0.35 + 0*0.15 = 50.0 → F
    if grade != "F" {
        t.Errorf("Expected 'F' with only assignments present; got '%s'", grade)
    }
}

func TestGradeTypeString(t *testing.T) {
    if Assignment.String() != "assignment" {
        t.Errorf("Expected 'assignment', got %s", Assignment.String())
    }
    if Exam.String() != "exam" {
        t.Errorf("Expected 'exam', got %s", Exam.String())
    }
    if Essay.String() != "essay" {
        t.Errorf("Expected 'essay', got %s", Essay.String())
    }
}
