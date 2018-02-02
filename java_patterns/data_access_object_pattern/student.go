package student

type Student struct {
	row  int
	name string
}

func NewStudent(row int, name string) *Student {
	return &Student{
		row:  row,
		name: name,
	}
}

func (s *Student) GetRow() int {
	return s.row
}

func (s *Student) SetRow(row int) {
	s.row = row
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}
