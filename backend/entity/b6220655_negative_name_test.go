package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerName(t *testing.T) {
	g := NewGomegaWithT(t)
	cus := Customer{
		Name:       "",
		CustomerID: "L1234567",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(cus)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("กรุณากรอกชื่อ"))

}
