package controller

import (
	"echo/config"
	"echo/helpers"
	"echo/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserLogin(c echo.Context) error {
	db := config.GetDB()

	Emp := model.Employee{}
	// contentType:=helpers.GetContentType(c)

	password := ""

	if err := c.Bind(&Emp); err != nil { //ngebind object echo context yg masuk ke objek employee
		return err
	}

	password = Emp.Password

	if err := db.Debug().Where("email=?", Emp.Email).Take(&Emp).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
	}

	fmt.Println(password)
	fmt.Println(Emp.Password)
	comparePass := helpers.ComparePass([]byte(Emp.Password), []byte(password)) //argumen pertama harus yg udah dihash, argumen kedua harus yg raw

	if !comparePass {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "invalid email/password2",
		})
	}

	token := helpers.GenerateToken(uint(Emp.ID), Emp.Email)

	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"token": token,
	})
}
