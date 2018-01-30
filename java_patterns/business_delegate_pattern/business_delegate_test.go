package delegate

import (
	"testing"
)

func TestBusinessDelegate(t *testing.T) {
	ejb := NewBusinessDelegate()
	ejb.SetServiceType("EJB")
	ejbClient := NewClient(ejb)
	ejbClient.DoTask()

	jms := NewBusinessDelegate()
	jms.SetServiceType("JMS")
	jmsClient := NewClient(jms)
	jmsClient.DoTask()
}
