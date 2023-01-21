package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tyange/pian-fiber/database"
	"github.com/tyange/pian-fiber/models"
)

func GetAllBurger(c *fiber.Ctx) error {
	var burgers []models.Burger

	database.DBConn.Find(&burgers)

	return c.Status(200).JSON(burgers)
}

func GetBurger(c *fiber.Ctx) error {
	burger := models.Burger{}
	id := c.Params("id")

	if database.DBConn.First(&burger, id) == nil {
		return c.Status(400).JSON(fiber.Map{"error": "버거를 찾지 못했습니다."})
	}

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
