package handler

import (
	"github.com/altuntasfatih/car-service-backend/internal/repair/service"
	"github.com/altuntasfatih/car-service-backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// CreateRepair godoc
// @Summary Create Repair
// @Description CreateRepair
// @ID RepairCreate
// @Tags Repairs
// @Accept json
// @Produce json
// @Param request body models.CreateRepairRequest true "Request"
// @Success 200 {object} models.GetRepairResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/repairs [post]
func CreateRepair(service service.RepairService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request models.CreateRepairRequest
		if err := ctx.BodyParser(&request); err != nil {
			return err
		}

		repair, err := service.CreateRepair(&request)
		if err != nil {
			return err
		}
		return ctx.JSON(&models.GetRepairResponse{Repair: repair})
	}
}

// GetRepair godoc
// @Summary Get Repair by repairId
// @Description GetRepair
// @ID GetRepair
// @Tags Repairs
// @Accept json
// @Produce json
// @Param repairId path string true "repairId"
// @Success 200 {object} models.GetRepairResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/repairs/{repairId} [get]
func GetRepair(service service.RepairService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Params("repairId", "")
		repair, err := service.GetRepair(userId)
		if err != nil {
			return err
		}

		return ctx.JSON(&models.GetRepairResponse{Repair: repair})
	}
}

// Getrepairs godoc
// @Summary Get All Repairs
// @Description GetRepairs
// @ID GetRepairs
// @Tags Repairs
// @Accept json
// @Produce json
// @Success 200
// @Success 200 {object} models.GetRepairsResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/repairs [get]
func GetRepairs(service service.RepairService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repairs, err := service.GetRepairs()
		if err != nil {
			return err
		}
		return ctx.JSON(&models.GetRepairsResponse{
			Repairs: repairs,
		})
	}
}

// DeleteRepair godoc
// @Summary Delete repair by repairId
// @Description DeleteRepair
// @ID DeleteRepair
// @Tags Repairs
// @Accept json
// @Produce json
// @Param repairId path string true "repairId"
// @Success 200
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/repairs/{repairId} [delete]
func DeleteRepair(service service.RepairService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		repairId := ctx.Params("repairId", "")
		if err := service.DeleteRepair(repairId); err != nil {
			return err
		}
		return ctx.Send(nil)
	}
}
