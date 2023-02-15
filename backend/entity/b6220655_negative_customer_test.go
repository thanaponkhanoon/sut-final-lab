package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerCorrect(t *testing.T) {
	g := NewGomegaWithT(t)
	cus := Customer{
		Name:       "thanapon",
		CustomerID: "L1234567",
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(cus)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ err ได้
	g.Expect(ok).To(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).To((BeNil()))

}
