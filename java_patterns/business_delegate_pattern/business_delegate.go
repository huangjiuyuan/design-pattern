package delegate

type BusinessDelegate struct {
	LookUpService   *BusinessLookUp
	BusinessService BusinessService
	ServiceType     string
}

func NewBusinessDelegate() *BusinessDelegate {
	return &BusinessDelegate{
		LookUpService: &BusinessLookUp{},
	}
}

func (bd *BusinessDelegate) SetServiceType(service string) {
	bd.ServiceType = service
}

func (bd *BusinessDelegate) DoTask() {
	bd.BusinessService = bd.LookUpService.GetBusinessService(bd.ServiceType)
	bd.BusinessService.DoProcessing()
}
