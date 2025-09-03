package handlers

import (
	model "github.com/ChanatpakornS/sa-REST-example/internal/models"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetConcertsAll(ctx fiber.Ctx, db *gorm.DB) error {
	var concerts []model.Concert

	result := db.Find(&concerts)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Concert not found",
		})
	}

	ctx.JSON(concerts)

	return ctx.Status(fiber.StatusOK).JSON(concerts)
}
func GetConcert(ctx fiber.Ctx, db *gorm.DB) error {
	var (
		concert model.Concert
		param   = struct {
			ID uint `uri:"id"`
		}{}
	)

	if err := ctx.Bind().URI(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID parameter",
		})
	}
	id := param.ID

	result := db.First(&concert, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Concert not found",
		})
	}

	ctx.JSON(concert)

	return ctx.Status(fiber.StatusOK).JSON(concert)
}
func CreateConcert(ctx fiber.Ctx, db *gorm.DB) error {
	var concert model.Concert
	if err := ctx.Bind().Body(&concert); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := db.Create(&concert).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create concert",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(concert)
}
func UpdateConcert(ctx fiber.Ctx, db *gorm.DB) error {
	var (
		concert model.Concert
		param   = struct {
			ID uint `uri:"id"`
		}{}
	)
	if err := ctx.Bind().URI(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID parameter",
		})
	}
	id := param.ID

	if err := ctx.Bind().Body(&concert); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := db.Where("id = ?", id).Updates(&concert).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update concert",
		})
	}

	result := db.First(&concert, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Concert not found",
		})
	}

	ctx.JSON(concert)

	return ctx.Status(fiber.StatusOK).JSON(concert)
}
func DeleteConcert(ctx fiber.Ctx, db *gorm.DB) error {
	var (
		concert model.Concert
		param   = struct {
			ID uint `uri:"id"`
		}{}
	)

	if err := ctx.Bind().URI(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID parameter",
		})
	}
	id := param.ID

	if err := db.Delete(&concert, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete concert",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
