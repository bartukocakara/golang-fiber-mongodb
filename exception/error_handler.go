package exception

import (
	"go-delivery-food/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(model.ApiResponse{
				Code : 400,
				Status : "BAD_REQUEST",
				Data : err.Error(),
		})
	}

	return ctx.JSON(model.ApiResponse{
				Code : 500,
				Status : "INTERNAL_SERVER_ERROR",
				Data : err.Error(),
	})
}