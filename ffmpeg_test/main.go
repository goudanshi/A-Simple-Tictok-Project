package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// ffmpeg URL-IO模式demo
var (
	client *minio.Client
)

// MinIO 服务的配置参数
const (
	endpoint        string = "127.0.0.1:9000/"
	accessKeyID     string = "minioadmin"
	secretAccessKey string = "minioadmin"
	useSSL          bool   = false
	bucketName      string = "video" //桶名
	picType         string = "image/jpeg"
)

func initMinio() {
	//创建 MinIO 客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("创建 MinIO 客户端失败", err)
		return
	}
	client = minioClient
	log.Printf("创建 MinIO 客户端成功")

	exist, err := client.BucketExists(context.Background(), bucketName)

	if err != nil {
		log.Fatalln("查找桶失败", err)
	}

	if !exist {
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln("创建桶失败", err)
		}
	}
	log.Printf("桶已创建")
}

func putPicture(buf *bytes.Buffer, saveName string) {
	//上传图片
	_, err := client.PutObject(context.Background(),
		bucketName,
		saveName,
		buf,
		int64(buf.Len()),
		minio.PutObjectOptions{
			ContentType: picType,
		})

	if err != nil {
		log.Fatalln("图片上传失败", err)
		return
	}
	fmt.Println("图片上传成功")
	//获取对象预签名url
	URL, err := client.PresignedGetObject(context.Background(), bucketName, saveName, time.Second*24*60*60, nil)
	if err != nil {
		log.Fatalln("获取url失败", err)
		return
	}
	fmt.Println("URL获取成功,URL为：", URL)
}

func main() {
	initMinio()

	// 设置视频源文件路径
	//inputFile := "./video/立冬.mp4"//本地视频路径
	inputFile := "https://www.w3schools.com/html/movie.mp4" //url

	// 设置 ffmpeg 命令行参数
	buf := bytes.NewBuffer(nil)
	saveName := "test_pic.jpeg"
	err := ffmpeg.Input(inputFile).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 5)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	// 运行 ffmpeg 命令
	if err != nil {
		log.Fatalln("截取图片失败", err)
		return
	}
	log.Printf("截取图片成功")

	putPicture(buf, saveName)

}
