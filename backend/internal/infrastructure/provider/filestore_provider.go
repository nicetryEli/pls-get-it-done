package provider_impl

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
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

func (provider *FilestoreProviderImpl) UploadFile(ctx context.Context, bucketName string, fileName string, file multipart.File, fileSize int64, options *minio.PutObjectOptions) (string, error) {
	uploadInfo, err := provider.minioClient.PutObject(ctx, bucketName, fileName, file, fileSize, *options)
	if err != nil {
		return "", err
	}
	return uploadInfo.Key, nil
}

func (provider *FilestoreProviderImpl) GetPresignedUrl(ctx context.Context, bucketName string, fileName string, expiraton time.Duration, downloadable bool, downloadableFilename string) (string, error) {
	reqParams := make(url.Values)
	if downloadable {
		reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", downloadableFilename))
	}
	presignUrl, err := provider.minioClient.PresignedGetObject(ctx, bucketName, fileName, expiraton, reqParams)
	if err != nil {
		return "", err
	}
	return presignUrl.String(), nil
}

func (provider *FilestoreProviderImpl) DeleteFile(ctx context.Context, bucketName string, fileName string, options *minio.RemoveObjectOptions) error {
	err := provider.minioClient.RemoveObject(ctx, bucketName, fileName, *options)
	if err != nil {
		return err
	}
	return nil
}

func (provider *FilestoreProviderImpl) GetBucketNames(ctx context.Context) ([]string, error) {
	buckets, err := provider.minioClient.ListBuckets(ctx)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(buckets))
	for _, bucket := range buckets {
		names = append(names, bucket.Name)
	}
	return names, nil
}
