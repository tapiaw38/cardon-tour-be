package bisiness_image

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
	"github.com/tapiaw38/cardon-tour-be/internal/platform/appcontext"
	"sync"
)

type (
	CreateUsecase interface {
		Execute(context.Context, string, []domain.BusinessImageFile) (*CreateOutput, error)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateOutput struct {
		Data []BusinessImageOutputData `json:"data"`
	}
)

func NewCreateUseCase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{
		contextFactory: contextFactory,
	}
}

func (u *createUsecase) Execute(ctx context.Context, businessID string, files []domain.BusinessImageFile) (*CreateOutput, error) {
	app := u.contextFactory()

	businessImagesChannel := make(chan domain.BusinessImage, len(files))
	errChannel := make(chan error, len(files))

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file domain.BusinessImageFile) {
			defer wg.Done()

			id, err := uuid.NewUUID()
			if err != nil {
				errChannel <- err
				return
			}

			fileName, err := app.StoreService.PutObject(file.File, file.FileHeader, id.String())
			if err != nil {
				errChannel <- err
				return
			}

			fileURL := app.StoreService.GenerateUrl(fileName)

			businessImage := domain.BusinessImage{
				ID:         id.String(),
				URL:        fileURL,
				BusinessID: businessID,
			}

			businessImagesChannel <- businessImage
		}(file)
	}

	go func() {
		wg.Wait()
		close(businessImagesChannel)
		close(errChannel)
	}()

	var businessImagesMemory []domain.BusinessImage
	for businessImage := range businessImagesChannel {
		businessImagesMemory = append(businessImagesMemory, businessImage)
	}

	select {
	case err := <-errChannel:
		if err != nil {
			return nil, err
		}
	default:
		fmt.Printf("the routine for creating business images is completed\n")
	}

	var businessImagesOutput []BusinessImageOutputData
	for _, images := range businessImagesMemory {
		idCreated, err := app.Repositories.BusinessImage.Create(ctx, images)
		if err != nil {
			return nil, err
		}

		businessImageCreated, err := app.Repositories.BusinessImage.Get(ctx, idCreated)
		if err != nil {
			return nil, err
		}

		businessImagesOutput = append(
			businessImagesOutput,
			toBusinessImageOutputData(businessImageCreated),
		)
	}

	return &CreateOutput{
		Data: businessImagesOutput,
	}, nil
}
