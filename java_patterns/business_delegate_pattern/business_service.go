package delegate

import (
	"fmt"
)

type BusinessService interface {
	DoProcessing()
}

type EJBService struct{}

func (ejb *EJBService) DoProcessing() {
	fmt.Println("Processing task by invoking EJB Service")
}

type JMSService struct{}

func (jms *JMSService) DoProcessing() {
	fmt.Println("Processing task by invoking JMS Service")
}
