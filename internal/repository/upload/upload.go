package upload

import (
	"context"
	"fmt"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/ecintiawan/loan-service/pkg/file"
)

type (
	// repoImpl implements Upload interface
	repoImpl struct {
		config      *config.Config
		fileManager file.File
	}
)

// New creates a new instance of repoImpl
func New(
	config *config.Config,
	fileManager file.File,
) repository.Upload {
	return &repoImpl{
		config:      config,
		fileManager: fileManager,
	}
}

// Upload will upload files based on model, to simulate uploading files from client
func (r *repoImpl) Upload(
	ctx context.Context,
	model *entity.File,
) (string, error) {
	var (
		err error
	)

	model.FilePath = r.config.Vendor.Upload.Path
	err = r.fileManager.Write(model.File, model.FilePath, model.FileName)
	if err != nil {
		return "", errorwrapper.E(err, errorwrapper.CodeInternal)
	}

	return fmt.Sprintf(r.config.Vendor.Upload.URL, model.FileName), nil
}
