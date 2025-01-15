package test

import (
	"fmt"
	"testing"

	"github.com/Nichakorn25/testfn03.git/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestStudentID (t *testing.T){

	g := NewGomegaWithT(t)

	t.Run(`student_id is required`, func(t *testing.T){
		user := entity.User{
			StudentId: "", //false
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			LinkIn:    "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})

	t.Run(`student_id pattern is not true`, func(t *testing.T){
		user := entity.User{

			StudentID: "K6500001", //false
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			LinkIn:    "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := gavalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)",user.StudentID)))
	})

	t.Run(`student_id is valid`,func(t *testing.T){
		user := entity.User{
			StudentID: "B6525972", //true
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			LinkIn:    "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(ok).To(BeNil())
	})
}