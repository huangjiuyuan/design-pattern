package delegate

type BusinessLookUp struct{}

func (bl *BusinessLookUp) GetBusinessService(service string) BusinessService {
	if service == "EJB" {
		return &EJBService{}
	}
	return &JMSService{}
}
