package customer

type Customer interface {
	IsNil() bool
	GetName() string
}

type RealCustomer struct {
	Name string
}

func NewRealCustomer(name string) *RealCustomer {
	return &RealCustomer{
		Name: name,
	}
}

func (rc *RealCustomer) IsNil() bool {
	return false
}

func (rc *RealCustomer) GetName() string {
	return rc.Name
}

type NilCustomer struct {
	Name string
}

func NewNilCustomer() *NilCustomer {
	return &NilCustomer{
		Name: "",
	}
}

func (rc *NilCustomer) IsNil() bool {
	return true
}

func (rc *NilCustomer) GetName() string {
	return "Unavailable in customer database"
}

type CustomerFactory struct {
	Names []string
}

func (cf *CustomerFactory) GetCustomer(name string) Customer {
	for i := 0; i < len(cf.Names); i++ {
		if name == cf.Names[i] {
			return NewRealCustomer(name)
		}
	}
	return NewNilCustomer()
}
