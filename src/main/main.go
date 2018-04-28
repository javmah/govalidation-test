package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

// User is a Struct To Hold The User
type User struct {
	ID          int    `valid:"length(1|5)~ ID is betwinr 1 to 5 Cheracter Long"  `
	Name        string `valid:"length(3|20)~ name Should betwine 3 to 20 Charecter  , alpha~ Name should Be alphabet  "`
	Email       string `valid:"email"`
	Phrone      string `valid:"numeric~ Phone Number Should Be Number Not Anything Else ,length(9|13) "`
	Location    string `valid:"length(1|20),optional"`
	Address     string `valid:"length(10|30) , optional"`
	Designation string `valid:"in(admin|dev|user|visitor|)"`
}

func main() {
	// validnewuser := User{
	// 	ID:          1,
	// 	Name:        "Javed ",
	// 	Email:       "bisonback@hotmail.com",
	// 	Phrone:      "01720012412",
	// 	Location:    "Mirpur Dhaka ",
	// 	Address:     "240 / A , Ahmed Nagor , Paik Para Mirpur 1 , Dhaka 1216",
	// 	Designation: "dev",
	// }

	invalidUser := User{
		ID:          123456789,
		Name:        "Javed 123 ",
		Email:       "bisonback@hotmail",
		Phrone:      "AAA01720012412",
		Address:     "240 / A , Ahmed Nagor , Paik Para Mirpur 1 , Dhaka 1216",
		Designation: "Xdev",
	}

	// result, err := govalidator.ValidateStruct(validnewuser)
	result, err := govalidator.ValidateStruct(invalidUser)
	// fmt.Println(err)

	if err != nil {
		println("error: " + err.Error())
	}

	fmt.Println(result)
}
