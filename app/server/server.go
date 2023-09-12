package server

import (
	"fmt"
	"github.com/RatulRahmanRudra/goly/models"
	"github.com/RatulRahmanRudra/goly/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strconv"
)

func redirect(ctx *fiber.Ctx) error {
	golyUrl := ctx.Params("redirect")
	goly, err := models.FindByGolyUrl(golyUrl)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	goly.Clicked++
	if err := models.UpdateGoly(goly); err != nil {
		fmt.Println("error updating goly. ", err.Error())
	}
	return ctx.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllRedirects(ctx *fiber.Ctx) error {
	gollies, err := models.GetAllGollies()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error fetching redirects " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(gollies)
}

func getGoly(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing goly id " + err.Error(),
		})
	}
	goly, err := models.GetGoly(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error fetching goly " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func createGoly(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	goly := models.GolyModel{}
	if err := ctx.BodyParser(&goly); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing goly " + err.Error(),
		})
	}
	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}
	if err := models.CreateGoly(goly); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	goly := models.GolyModel{}
	err := ctx.BodyParser(&goly)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := models.UpdateGoly(goly); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func deleteGoly(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messege": err.Error(),
		})
	}

	if err := models.DeleteGoly(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goly deleted",
	})
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "origin, Content-Type, Accept",
	}))
	router.Get("/r/:redirect", redirect)
	router.Get("/goly", getAllRedirects)
	router.Get("/goly/:id", getGoly)
	router.Post("goly", createGoly)
	router.Patch("/goly", updateGoly)
	router.Delete("/goly/:id", deleteGoly)
	router.Listen(":3000")
}
