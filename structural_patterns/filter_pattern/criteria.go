package person

type Criteria interface {
	MeetCriteria(people []Person) []Person
}

type CriteriaMale struct{}

func (c *CriteriaMale) MeetCriteria(people []Person) []Person {
	malePeople := make([]Person, 0)
	for _, person := range people {
		if person.GetGender() == "Male" {
			malePeople = append(malePeople, person)
		}
	}
	return malePeople
}

type CriteriaFemale struct{}

func (c *CriteriaFemale) MeetCriteria(people []Person) []Person {
	femalePeople := make([]Person, 0)
	for _, person := range people {
		if person.GetGender() == "Female" {
			femalePeople = append(femalePeople, person)
		}
	}
	return femalePeople
}

type CriteriaSingle struct{}

func (c *CriteriaSingle) MeetCriteria(people []Person) []Person {
	singlePeople := make([]Person, 0)
	for _, person := range people {
		if person.GetMaritalStatus() == "Single" {
			singlePeople = append(singlePeople, person)
		}
	}
	return singlePeople
}

type CriteriaMarried struct{}

func (c *CriteriaMarried) MeetCriteria(people []Person) []Person {
	marriedPeople := make([]Person, 0)
	for _, person := range people {
		if person.GetMaritalStatus() == "Married" {
			marriedPeople = append(marriedPeople, person)
		}
	}
	return marriedPeople
}

type AndCriteria struct {
	First  Criteria
	Second Criteria
}

func (c *AndCriteria) AndCriteria(first Criteria, second Criteria) {
	c.First = first
	c.Second = second
}

func (c *AndCriteria) MeetCriteria(people []Person) []Person {
	first := c.First.MeetCriteria(people)
	second := c.Second.MeetCriteria(first)
	return second
}

type OrCriteria struct {
	First  Criteria
	Second Criteria
}

func (c *OrCriteria) OrCriteria(first Criteria, second Criteria) {
	c.First = first
	c.Second = second
}

func (c *OrCriteria) MeetCriteria(people []Person) []Person {
	first := c.First.MeetCriteria(people)
	second := c.Second.MeetCriteria(people)
	for _, p1 := range first {
		for _, p2 := range second {
			if p1 != p2 {
				second = append(second, p1)
			}
		}
	}
	return second
}
