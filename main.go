package main

import (
	"bytes"
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"os"
)

func main() {
	setLogLevel(LoglevelInfo)

	bucket := flag.String("bucket", "", "S3 bucket to which the files will be sent")
	region := flag.String("region", "us-east-1", "S3 bucket region")
	file := flag.String("file", "", "Path of the file")

	flag.Parse()

	if *bucket == "" || *file == "" {
		logError("All flags must be present. Exiting.")
		os.Exit(-1)
	}

	session, err := session.NewSession(&aws.Config{Region: aws.String(*region)})
	if err != nil {
		log.Fatal(err)
	}

	err = uploadFile(session, *file, *bucket)
	if err != nil {
		log.Fatal(err)
	}
}

func uploadFile(session *session.Session, uploadFileDir string, bucket string) error {
	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	fileSize := upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)

	_, err = upFile.Read(fileBuffer)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(uploadFileDir),
		ACL:                  aws.String("public"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}