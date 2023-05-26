package main

import "github.com/labstack/echo/v4"

type People struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type BaseResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"name"`
}

func main(){ 
	e := echo.New()
	e.GET("/peoples", GetPeopleController)
	e.GET("/peoples/:id", GetDetailPeopleController)
	e.POST("/login", LoginController)

	e.Start(":8000")
}

func LoginController(c echo.Context) error {
	var loginRequest LoginRequest
	c.Bind(&loginRequest)
	return c.JSON(200, loginRequest)
}

func GetDetailPeopleController(c echo.Context) error {
	var id = c.Param("id")
	var people = People{Name: id, Age: 1}
	return c.JSON(200, people)
}

func GetPeopleController(c echo.Context) error {
	var location = c.QueryParam("location")


	var peopleData []People

	var people People
	people.Name = location
	people.Age = 20
	peopleData = append(peopleData, people)

	people.Name = location
	people.Age = 30
	peopleData = append(peopleData, people)

	var response BaseResponse
	response.Data = peopleData
	response.Message = "Berhasil"
	return c.JSON(200, response)
}
