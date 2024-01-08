package routers

import (
	"github.com/gofiber/fiber/v2"
	IdentityService "service/controllers"
)

func Api(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/sign/:sender?", IdentityService.Sign)
	api.Get("/verify/:sender?", IdentityService.VerifyAttestation)
	api.Get("/claim/:sender?", IdentityService.VerifyClaim)

}
