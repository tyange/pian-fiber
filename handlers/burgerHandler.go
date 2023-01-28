package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tyange/pian-fiber/database"
	"github.com/tyange/pian-fiber/models"
	"math"
	"strconv"
)

type Pagination struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
}

type BurgersData struct {
	Burgers  []models.Burger
	PageData Pagination
}

func GetAllBurger(c *fiber.Ctx) error {
	var burgers []models.Burger

	perPage := 6
	page := 1
	pageStr := c.Query("page")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	var totalBurger int64
	database.DBConn.Model(&models.Burger{}).Count(&totalBurger)
	totalPages := math.Ceil(float64(totalBurger / int64(perPage)))

	offset := (page - 1) * perPage

	database.DBConn.Limit(6).Offset(offset).Find(&burgers)

	return c.Status(200).JSON(BurgersData{Burgers: burgers, PageData: Pagination{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		CurrentPage:  page,
		TotalPages:   int(totalPages),
	}})
}

func GetBurger(c *fiber.Ctx) error {
	burger := models.Burger{}
	id := c.Params("id")

	if database.DBConn.First(&burger, id) == nil {
		return c.Status(400).JSON(fiber.Map{"error": "버거를 찾지 못했습니다."})
	}

	return c.Status(200).JSON(burger)
}

func AddBurger(c *fiber.Ctx) error {
	burger := models.Burger{}

	if c.BodyParser(&burger) != nil {
		return c.Status(400).JSON(fiber.Map{"error": "입력하려는 데이터에 오류가 있습니다."})
	}

	database.DBConn.Save(&burger)

	return c.Status(200).JSON(burger)
}

func EditBurger(c *fiber.Ctx) error {
	burger := models.Burger{}
	updatedBurger := models.Burger{}
	id := c.Params("id")

	if database.DBConn.First(&burger, id) == nil {
		return c.Status(400).JSON(fiber.Map{"error": "수정하려는 버거를 찾지 못했습니다."})
	}

	if c.BodyParser(&updatedBurger) != nil {
		return c.Status(400).JSON(fiber.Map{"error": "수정하려는 데이터 입력에 오류가 있습니다."})
	}

	burger.Name = updatedBurger.Name
	burger.Brand = updatedBurger.Brand
	burger.Description = updatedBurger.Description

	database.DBConn.Save(&burger)

	return c.Status(200).JSON(burger)
}

func DeleteBurger(c *fiber.Ctx) error {
	burger := models.Burger{}
	id := c.Params("id")

	if database.DBConn.First(&burger, id) == nil {
		return c.Status(400).JSON(fiber.Map{"error": "삭제하려는 버거를 찾지 못했습니다."})
	}

	database.DBConn.Delete(&burger)

	return c.Status(200).JSON(fiber.Map{"message": "버거를 삭제했습니다."})
}
