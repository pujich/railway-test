package controller

import (
	"echo/config"
	"echo/model"
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hellooooo")
}

func JsonMap(c echo.Context) error {
	data := model.M{
		"message":    "Hello",
		"counter":    2,
		"statusCode": http.StatusOK,
	}
	return c.JSON(http.StatusOK, data)
}

func Param(c echo.Context) error {
	name := c.QueryParam("name")
	data := "Hello " + name
	// result := fmt.Sprintf("%s", data)
	return c.JSON(http.StatusOK, data)
}

func User(c echo.Context) error {
	user := model.Item{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create Employee
// @Description Create Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param model.Employee body model.Employee true "Employee"
// @Success 200 {object} string "ok"
// @Router /create [post]
func CreateUser(c echo.Context) error {
	db := config.GetDB()
	employee := model.Employee{}

	if err := c.Bind(&employee); err != nil { //ngebind object echo context yg masuk ke objek employee
		return err
	}

	db.Debug().Create(&employee)

	return c.JSON(http.StatusOK, employee)
	// db, err := config.Connect()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// employee := model.Employee{}
	// if err := c.Bind(&employee); err != nil {
	// 	return err
	// }

	// sqlStatement := `INSERT INTO employees (name, email) VALUES ($1,$2)`

	// _, err = db.Exec(sqlStatement, employee.Name, employee.Email)
	// if err != nil {
	// 	panic(err)
	// }
}

func CreateItem(c echo.Context) error {
	db := config.GetDB()
	item := model.Item{}

	userData, ok := c.Get("userData").(jwt.MapClaims)

	if !ok {
		//in case userData is not jwt.MapClass
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     userData,
			"message": "failed to get user data",
		})
	}

	userID := uint(userData["id"].(float64))

	if err := c.Bind(&item); err != nil {
		return err
	}

	item.EmployeeId = int(userID)
	db.Debug().Create(&item)

	return c.JSON(http.StatusOK, item)

}

func UpdateEmployee(c echo.Context) error {
	db := config.GetDB()
	employee := model.Employee{}
	if err := c.Bind(&employee); err != nil {
		return err
	}
	// db.Debug().Save(&employee)

	db.Debug().Model(&employee).Where("id=?", employee.ID).Updates(model.Employee{
		Name:     employee.Name,
		Email:    employee.Email,
		Division: employee.Division,
	})

	return c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c echo.Context) error {
	db := config.GetDB()
	employee := model.Employee{}

	delResp := model.DeleteResponse{
		Status:  http.StatusOK,
		Message: "Delete Success!",
	}

	paramId := c.Param("id")

	if err := c.Bind(&employee); err != nil {
		return err
	}

	db.Debug().Model(&employee).Where("id=?", paramId).Delete(&employee)

	return c.JSON(http.StatusOK, delResp)
}

func Index(c echo.Context) error {
	tmpl :=
		template.Must(template.ParseGlob("template/*.html"))

	type M map[string]interface{}
	data := make(M)

	data[config.CSRFKey] = c.Get(config.CSRFKey)
	return tmpl.Execute(c.Response(), data)
}

func SayHello(c echo.Context) error {
	type M map[string]interface{}
	data := make(M)

	if err := c.Bind(&data); err != nil {
		return err
	}

	message := fmt.Sprintf("Hello %s , My Gender %s", data["name"], data["gender"])

	return c.JSON(http.StatusOK, message)

}
