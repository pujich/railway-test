package model

import (
	"echo/helpers"
	"fmt"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type M map[string]interface{}

type Item struct {
	Name       string `json:"name" form:"name"`
	EmployeeId int
	Employee   *Employee
}

type Employee struct {
	ID       int    `json:"id" form:"id" swagger:"description(ID)"`
	Name     string `json:"name" form:"name" swagger:"description(Name)" valid:"required"`
	Email    string `json:"email" form:"email" swagger:"description(Email)" valid:"required"`
	Password string `json:"password" form:"password" swagger:"description(Password)" valid:"required"`
	Division string `json:"division" form:"division" swagger:"description(Division)" valid:"required"`
	Item     []Item
}

type DeleteResponse struct {
	Status  int    `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) { //kebaca duluan, gaperlu dipanggil. yg njalanin gorm
	fmt.Println("masuk ke before create")
	_, errCreate := govalidator.ValidateStruct(e)

	if errCreate != nil {
		err = errCreate
		return
	}
	e.Password = helpers.HassPass(e.Password)
	err = nil
	return
}
