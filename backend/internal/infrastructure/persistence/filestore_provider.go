package persistence_impl

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
)

type FilestoreProviderImpl struct {
	minioClient *minio.Client
}

func NewFilestoreProviderImpl(minioClient *minio.Client) *FilestoreProviderImpl {
	return &FilestoreProviderImpl{
		minioClient: minioClient,
	}
}

func (persistence *FilestoreProviderImpl) UploadFile(ctx context.Context, bucketName string, fileName string, file multipart.File, fileSize int64, options *minio.PutObjectOptions) (string, error) {
	return "", nil
}

func (persistence *FilestoreProviderImpl) GetPresignedUrl(ctx context.Context, bucketName string, fileName string, expiraton time.Duration, downloadable bool, downloadableFilename string) (string, error) {
	return "", nil
}

func (persistence *FilestoreProviderImpl) DeleteFile(ctx context.Context, bucketName string, fileName string, versionId string) error {
	return nil
}

func (persistence *FilestoreProviderImpl) GetBucketNames(ctx context.Context) ([]string, error) {
	buckets, err := persistence.minioClient.ListBuckets(ctx)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(buckets))
	for _, bucket := range buckets {
		names = append(names, bucket.Name)
	}
	return names, nil
}
