package config

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	MinioClient     *minio.Client
	minioClientOnce sync.Once
)

func init() {
	minioClientOnce.Do(func() {
		minioClient, err := minio.New(
			fmt.Sprintf("%s:%s", Env.MINIO_HOST, Env.MINIO_API_PORT_NUMBER),
			&minio.Options{
				Creds:  credentials.NewStaticV4(Env.MINIO_ROOT_USER, Env.MINIO_ROOT_PASSWORD, ""),
				Secure: Env.MINIO_USE_SSL,
			},
		)
		if err != nil {
			log.Fatalln(err)
			return
		}
		for _, bucket := range []string{"user-avatar"} {
			found, err := minioClient.BucketExists(context.Background(), bucket)
			if err != nil {
				log.Fatalln(err)
				return
			}
			if found {
				continue
			}
			err = minioClient.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{ObjectLocking: true, Region: ""})
			if err != nil {
				log.Fatalln(err)
				return
			}
		}
		MinioClient = minioClient
	})
}
