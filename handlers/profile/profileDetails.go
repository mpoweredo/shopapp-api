package profile

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"shop-app-API/database"
	"shop-app-API/models"
	"shop-app-API/resources"
	"shop-app-API/utils"
)

type UpdateProfileDetailsRequest struct {
	Description string  `json:"description" validate:"required,min=1,max=500"`
	Image       *[]byte `json:"image"`
}

func UpdateProfileDetails(c *fiber.Ctx) error {
	var ctx = context.Background()

	var user models.User

	u := utils.GetUser(c)

	b := UpdateProfileDetailsRequest{}

	if err := c.BodyParser(&b); err != nil {
		return utils.ReturnFiberError(c, err.Error())
	}

	errors := utils.ValidateStruct(b)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	database.DB.Model(&user).Where("id = ?", u.Id).Update("description", b.Description)

	file, _ := c.FormFile("image")

	if file != nil {

		buffer, _ := file.Open()
		defer buffer.Close()

		cld, err := utils.GetCld()

		if err != nil {
			return utils.ReturnFiberError(c, err.Error())
		}

		r, err := cld.Upload.Upload(ctx, buffer, uploader.UploadParams{Folder: "profile_images"})

		if err != nil {
			return utils.ReturnFiberError(c, "Couldn't save image")
		}

		database.DB.Model(&user).Where("id = ?", u.Id).Update("photo", r.URL)
	}

	return c.Status(fiber.StatusOK).JSON(resources.SnackbarResponse{
		Message:     "Updated",
		Type:        resources.SUCCESS,
		Description: "Successfully updated profile details",
	})
}
