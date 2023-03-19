package profile

import (
	"github.com/gofiber/fiber/v2"
	"shop-app-API/database"
	"shop-app-API/models"
	"shop-app-API/resources"
	"shop-app-API/utils"
)

type AddUpdateDeliveryAddressRequest struct {
	Firstname     string `json:"firstname" validate:"required,min=2,max=35"`
	Lastname      string `json:"lastname" validate:"required,min=2,max=35"`
	StreetAddress string `json:"streetAddress" validate:"required,min=2,max=50"`
	Building      string `json:"building" validate:"required,min=1,max=30"`
	City          string `json:"city" validate:"required,min=2,max=60"`
	PostCode      string `json:"postCode" validate:"required,len=6"`
	Country       string `json:"country" validate:"required,min=2,max=60"`
	Province      string `json:"province" validate:"required,min=2,max=60"`
	Phone         string `json:"phone" validate:"required,len=9"`
}

func AddDeliveryAddress(c *fiber.Ctx) error {
	user := utils.GetUser(c)

	b := AddUpdateDeliveryAddressRequest{}

	if err := c.BodyParser(&b); err != nil {
		return utils.ReturnFiberError(c, err.Error())
	}

	errors := utils.ValidateStruct(b)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	deliveryAddress := models.DeliveryAddresses{
		Firstname:     b.Firstname,
		Lastname:      b.Lastname,
		StreetAddress: b.StreetAddress,
		Building:      b.Building,
		City:          b.City,
		PostCode:      b.PostCode,
		Country:       b.Country,
		Province:      b.Province,
		Phone:         b.Phone,
		UserId:        user.Id,
	}

	result := database.DB.Create(&deliveryAddress)

	if result.Error != nil {
		return utils.ReturnFiberError(c, "Something went wrong while adding new delivery address")
	}

	return c.Status(200).JSON(fiber.Map{
		"snackbar": resources.SnackbarResponse{Message: "Success!", Type: resources.SUCCESS, Description: "Successfully added delivery address"},
	})
}

func GetDeliveryAddresses(c *fiber.Ctx) error {
	user := utils.GetUser(c)

	var deliveryAddresses []models.DeliveryAddresses

	database.DB.Find(&deliveryAddresses, "user_id = ?", user.Id)

	return c.Status(200).JSON(fiber.Map{
		"data": deliveryAddresses,
	})
}

func DeleteDeliveryAddress(c *fiber.Ctx) error {
	id := c.Params("id")

	var deliveryAddresses models.DeliveryAddresses

	database.DB.Delete(deliveryAddresses, "id = ?", id)

	return c.Status(200).JSON(fiber.Map{
		"snackbar": resources.SnackbarResponse{
			Message:     "Success",
			Type:        resources.SUCCESS,
			Description: "Successfully deleted delivery address"},
	})
}

func UpdateDeliveryAddress(c *fiber.Ctx) error {
	id := c.Params("id")

	user := utils.GetUser(c)

	b := AddUpdateDeliveryAddressRequest{}

	if err := c.BodyParser(&b); err != nil {
		return utils.ReturnFiberError(c, err.Error())
	}

	errors := utils.ValidateStruct(b)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	updatedDeliveryAddress := models.DeliveryAddresses{
		Firstname:     b.Firstname,
		Lastname:      b.Lastname,
		StreetAddress: b.StreetAddress,
		Building:      b.Building,
		City:          b.City,
		PostCode:      b.PostCode,
		Country:       b.Country,
		Province:      b.Province,
		Phone:         b.Phone,
	}

	var deliveryAddress models.DeliveryAddresses

	database.DB.Model(&deliveryAddress).
		Where("id = ?", id).
		Where("user_id = ?", user.Id).
		Updates(updatedDeliveryAddress)

	return nil
}
