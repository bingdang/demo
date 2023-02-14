package Students

type Students struct {
	Name  string
	Score float32
	Sex   int
	Grade int
}

func NewStudents(name string, score float32, sex, grade int) (S *Students) {
	S = &Students{
		Name:  name,
		Score: score,
		Sex:   sex,
		Grade: grade,
	}
	return
}
