package school

import "sort"

// Grade represents a school grade and students who are a part of the grade.
type Grade struct {
	number   int
	students []string
}

// School represents a school which can have grades, each of which can have students.
type School struct {
	grades map[int]*Grade
}

// New creates a new School.
func New() *School {
	return &School{grades: map[int]*Grade{}}
}

// Add adds a new student to a specified grade in School s. If the grade
// does not yet exist in s, then a new Grade is created and the student
// is added to that grade.
func (s *School) Add(studentName string, gradeNumber int) {
	if g, ok := s.grades[gradeNumber]; ok {
		g.students = append(g.students, studentName)
	} else {
		s.grades[gradeNumber] = &Grade{number: gradeNumber, students: []string{studentName}}
	}
}

// Grade returns a list of students in a given grade at School s.
func (s *School) Grade(gradeNumber int) []string {
	if g, ok := s.grades[gradeNumber]; ok {
		return g.students
	}
	return []string{}
}

// Enrollment returns a sorted list of all students in all grades.
func (s *School) Enrollment() []Grade {
	var sortedGrades []Grade

	sortedGradeNumbers := make([]int, 0, len(s.grades))
	for k := range s.grades {
		sortedGradeNumbers = append(sortedGradeNumbers, k)
	}
	sort.Ints(sortedGradeNumbers)

	for _, gradeNumber := range sortedGradeNumbers {
		sort.Strings(s.grades[gradeNumber].students)
		sortedGrades = append(sortedGrades, *(s.grades[gradeNumber]))
	}
	return sortedGrades
}
