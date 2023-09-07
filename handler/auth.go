package handler

import (
	"avenger/entity"
	"avenger/repo"
	"avenger/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DBHandler repo.DBHandler
}

func (ah AuthHandler) Register(c echo.Context) error {
	var input entity.User
	validate := validator.New()
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err.Error(),
			Details: err,
		})
	}

	err = validate.Struct(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err.Error(),
			Details: err,
		})
	}

	bytePass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
			Details: err,
		})
	}
	input.Password = string(bytePass)
	user, err := ah.DBHandler.Register(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
			Details: err,
		})
	}
	return c.JSON(http.StatusCreated, map[string]any{
		"message": "Registration successful.",
		"user":    user,
	})
}
func (ah AuthHandler) Login(c echo.Context) error { return nil }

func (ah AuthHandler) GetStores(c echo.Context) error {
	// Slice of entity.Store filled with data from DB.
	var stores []entity.Store
	result := ah.DBHandler.DB.Find(&stores)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: result.Error.Error(),
			Details: result.Error,
		})
	}
	//print
	return c.JSON(http.StatusAccepted, stores)
}
func (ah AuthHandler) GetStoreByID(c echo.Context) error {
	var store entity.Store
	id := c.Param("id")
	result := ah.DBHandler.DB.Where("id = ?", id).Find(&store)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: result.Error.Error(),
			Details: result.Error,
		})
	}

	url := fmt.Sprintf("https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?lat=%v&lon=%v", store.Latitude, store.Longitude)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("WEATHER_KEY"))
	req.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
			Details: err,
		})
	}
	defer res.Body.Close()

	var weather entity.Weather
	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
			Details: err,
		})
	}
	return c.JSON(http.StatusAccepted, map[string]any{
		"coordinate": map[string]any{
			"latitude":  store.Latitude,
			"longitude": store.Longitude,
		},
		"total_sales": store.TotalSales,
		"rating":      store.Rating,
		"weather":     weather,
	})
}
