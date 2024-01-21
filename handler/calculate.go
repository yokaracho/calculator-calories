package handler

import (
	"blocker/loggers"
	hellopage "blocker/templates/hello"
	"fmt"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"strconv"
)

func render(c echo.Context, component templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return component.Render(c.Request().Context(), c.Response())
}

func CalculateHandler(c echo.Context) error {
	logger := loggers.GetLogger()
	heights := c.FormValue("height")
	ages := c.FormValue("age")
	weights := c.FormValue("weight")
	males := c.FormValue("gender")
	actives := c.FormValue("activity")

	age, err := strconv.Atoi(ages)
	if err != nil {
		fmt.Println("Ошибка при преобразовании возраста в int:", err)
		logger.Info("Ошибка при преобразовании возраста в int:")
		return err
	}

	height, err := strconv.Atoi(heights)
	if err != nil {
		fmt.Println("Ошибка при преобразовании роста в float64:", err)
		logger.Info("Ошибка при преобразовании возраста в float64:")

		return err
	}

	weight, err := strconv.Atoi(weights)
	if err != nil {
		fmt.Println("Ошибка при преобразовании веса в int:", err)
		logger.Info("Ошибка при преобразовании возраста в int:")

		return err
	}
	var bmrMen, bmrWomen, bmrRes int

	if males == "male" {
		bmrMen = int(10*float64(weight) + 6.25*float64(height) - 5*float64(age) + 5)
		bmrRes = bmrMen
	} else {
		bmrWomen = int(10*float64(weight) + 6.25*float64(height) - 5*float64(age) - 161)
		bmrRes = bmrWomen
	}
	var result float64
	switch actives {
	case "sedentary":
		result = float64(bmrRes) * 1.2

	case "light":
		result = float64(bmrRes) * 1.375

	case "moderate":
		result = float64(bmrRes) * 1.55

	case "active":
		result = float64(bmrRes) * 1.725

	case "hard":
		result = float64(bmrRes) * 1.9
	}

	resultStr := strconv.Itoa(int(result))
	return render(c, hellopage.HelloName(resultStr))
}
