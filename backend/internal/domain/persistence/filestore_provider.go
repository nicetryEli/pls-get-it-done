package persistence

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
)

type FilestoreProvider interface {
	UploadFile(ctx context.Context, bucketName string, fileName string, file multipart.File, fileSize int64, options *minio.PutObjectOptions) (string, error)
	GetPresignedUrl(ctx context.Context, bucketName string, fileName string, expiraton time.Duration, downloadable bool, downloadableFilename string) (string, error)
	DeleteFile(ctx context.Context, bucketName string, fileName string, versionId string) error
	GetBucketNames(ctx context.Context) ([]string, error)
}
