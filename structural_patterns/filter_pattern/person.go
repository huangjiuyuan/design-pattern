package person

type Person struct {
	Name          string
	Gender        string
	MaritalStatus string
}

func NewPerson(name string, gender string, maritalStatus string) *Person {
	return &Person{
		Name:          name,
		Gender:        gender,
		MaritalStatus: maritalStatus,
	}
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetGender() string {
	return p.Gender
}

func (p *Person) GetMaritalStatus() string {
	return p.MaritalStatus
}
