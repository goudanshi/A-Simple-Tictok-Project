package util

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"time"
)

var client *minio.Client

func InitOSS() error {
	useSSL := false
	minioClient, err := minio.New(MINIO_END_POINT, &minio.Options{
		Creds:  credentials.NewStaticV4(MINIO_ACCESS_KEY, MINIO_SECRET_ACCESS_KEY, ""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}
	client = minioClient

	exist, err := client.BucketExists(context.Background(), MINIO_BUCKET)
	if err != nil {
		return err
	}
	if !exist {
		err = client.MakeBucket(ctx, MINIO_BUCKET, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func PutVideo(name string, reader io.Reader, size int64) error {
	return PutObject(name, reader, size, "video/mp4")
}

func PutJpg(name string, reader io.Reader, size int64) error {
	return PutObject(name, reader, size, "image/jpeg")
}

func PutObject(name string, reader io.Reader, size int64, contentType string) error {
	_, err := client.PutObject(ctx, MINIO_BUCKET, name, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func GetObject(name string) (io.Reader, error) {
	return client.GetObject(ctx, MINIO_BUCKET, name, minio.GetObjectOptions{})
}

func StatObject(name string) (minio.ObjectInfo, error) {
	return client.StatObject(ctx, MINIO_BUCKET, name, minio.StatObjectOptions{})
}

func GetUrl(name string) string {
	presignedURL, err := client.PresignedGetObject(context.Background(), MINIO_BUCKET, name, time.Second*24*60*60, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return "http://" + presignedURL.Host + presignedURL.RequestURI()
}

func GetObjectWithSize(name string) (io.Reader, int64, time.Time, error) {
	info, err := StatObject(name)
	if err != nil {
		return nil, -1, time.Now(), err
	}
	object, err := GetObject(name)
	if err != nil {
		return nil, -1, time.Now(), err
	}
	return object, info.Size, info.LastModified, nil
}
