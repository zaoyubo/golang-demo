package common

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type CreateStudentReq struct {
	Name        string `json:"name" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	StudentType int    `json:"student_type" validate:"required,studentType"`

	// *float64 可以区分零值和空值， 而 float64 不能
	Tuition float64  `json:"tuition" validate:"required,gte=0"`
	Loan    *float64 `json:"loan" validate:"omitempty,twoDecimal"` // omitempty：当值不为空时，才进行其他校验

	MinCredits *int `json:"minCredits"`
	MaxCredits *int `json:"max_credits" validate:"omitempty,gtefield=MinCredits"` // 若 MinCredits 和 MaxCredits 一定同时
	// 出现，则可以这么写。否则若只传MaxCredits，不会通过gtefield 校验
}

func TestValidator(t *testing.T) {
	var validate *validator.Validate = NewValidator()
	//maxCredits := 200
	req := CreateStudentReq{
		Name:        "zhangsan",
		Age:         1,
		StudentType: 1,
		Tuition:     100,
		//MaxCredits:  &maxCredits,
	}
	if err := validate.Struct(&req); err != nil {
		fmt.Printf("validate error:%v\n", err)
	}
	fmt.Println("validate success")

}
