package todo

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func CreateApp() *fiber.App {
	app := fiber.New()

	return app
}
type HandlerService interface {
	AddNewToDo(ctx context.Context, request *AddToDoRequest) (*AddToDoResponse, error)
	GetAll() *GetAllResponse
}

type Handler struct {
	Service HandlerService

}
func NewHandler(s HandlerService) *Handler {
	handler := Handler{
		Service: s,
	}
	return &handler
}

func ( h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/health", h.Health )
	app.Post("/newtodo", h.AddNewToDo )
	app.Get("/alltodo", h.GetAll )
}

func (h *Handler) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status" : "OK",
	})
}

func (h *Handler) AddNewToDo(c *fiber.Ctx) error {
	var request AddToDoRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"Bad Request",
		})
	}
	result,err := h.Service.AddNewToDo(c.Context(), &request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":"Internal Server Error",
		})
	} 
	return c.Status(fiber.StatusCreated).JSON(result)

} 
func (h *Handler) GetAll(c *fiber.Ctx) error {
	result := h.Service.GetAll()

	return c.Status(fiber.StatusOK).JSON(result)
}