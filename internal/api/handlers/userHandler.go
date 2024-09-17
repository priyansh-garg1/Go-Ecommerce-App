package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoute(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
	}
	handler := UserHandler{
		svc: svc,
	}

	//Public
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	//Private
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)

	app.Post("/become-seller", handler.Login)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {

	user := dto.UserSignUp{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Please provide valid input"})
	}
	token, err := h.svc.SignUp(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error While Signing"})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"message": token,
	})
}
func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Please provide valid input"})
	}
	token, err := h.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Error While Signing"})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Logedin",
		"token":   token,
	})
}
func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "verify user",
	})
}
func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Login user",
	})
}
func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Login user",
	})
}
func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Login user",
	})
}
func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Login user",
	})
}
func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Login user",
	})
}
func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Login user",
	})
}
func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Login user",
	})
}
func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Login user",
	})
}
