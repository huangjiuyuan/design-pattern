package person

import (
	"fmt"
	"testing"
)

func TestMeetCriteria(t *testing.T) {
	people := make([]Person, 0)
	people = append(people, Person{"Robert", "Male", "Single"})
	people = append(people, Person{"John", "Male", "Married"})
	people = append(people, Person{"Laura", "Female", "Married"})
	people = append(people, Person{"Diana", "Female", "Single"})
	people = append(people, Person{"Mike", "Male", "Single"})
	people = append(people, Person{"Bobby", "Male", "Single"})

	male := &CriteriaMale{}
	female := &CriteriaFemale{}
	single := &CriteriaSingle{}
	married := &CriteriaMarried{}
	singleMale := &AndCriteria{single, male}
	marriedOrFemale := &OrCriteria{married, female}

	fmt.Printf("Male:\n %v\n", male.MeetCriteria(people))
	fmt.Printf("Female:\n %v\n", female.MeetCriteria(people))
	fmt.Printf("Single:\n %v\n", single.MeetCriteria(people))
	fmt.Printf("Married:\n %v\n", married.MeetCriteria(people))
	fmt.Printf("Single Male:\n %v\n", singleMale.MeetCriteria(people))
	fmt.Printf("Married or Female:\n %v\n", marriedOrFemale.MeetCriteria(people))
}
