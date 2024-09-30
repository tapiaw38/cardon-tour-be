package business_image

import (
	"github.com/gin-gonic/gin"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
	businessimage "github.com/tapiaw38/cardon-tour-be/internal/usecases/business/bisiness_image"
)

func NewCreateHandler(usecase businessimage.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		businessID := c.Param("business_id")
		if businessID == "" {
			c.JSON(400, gin.H{
				"error": "business id is required",
			})
			return
		}

		err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max
		if err != nil {
			c.JSON(400, gin.H{
				"error": "failed to parse form data",
			})
			return
		}

		formFiles := c.Request.MultipartForm.File["pictures"]
		if len(formFiles) == 0 {
			c.JSON(400, gin.H{
				"error": "no files uploaded",
			})
			return
		}

		var files []domain.BusinessImageFile

		for _, fileHeader := range formFiles {
			file, err := fileHeader.Open()
			if err != nil {
				c.JSON(400, gin.H{
					"error": "failed to open file",
				})
				return
			}
			defer file.Close()

			files = append(files, domain.BusinessImageFile{
				File:       file,
				FileHeader: fileHeader,
			})
		}

		businessImages, err := usecase.Execute(c.Request.Context(), businessID, files)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, businessImages)
	}
}
