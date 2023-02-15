package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Emplyee struct {
	gorm.Model
	Name      string
	Email     string
	EmplyeeID string
}

func TestDataValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check Data validate", func(t *testing.T) {

		u := Emplyee{
			Name:      "kkk",
			Email:     "kkk@gmail.com",
			EmplyeeID: "12345678",
		}

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

}
