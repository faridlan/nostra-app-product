package service

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/web"
)

// Interface Of Upload S3 AWS
type UploadS3AWS interface {
	Upload(body io.Reader, driectory string) web.UploadResponse
}

// Implemantation Interfave Upload S3 AWS
type UploadS3AWSImpl struct {
}

func NewUploadS3AWS() UploadS3AWS {
	return &UploadS3AWSImpl{}
}

func (service *UploadS3AWSImpl) Upload(body io.Reader, directory string) web.UploadResponse {
	helper.GetEnvWithKey("AWS_REGION")
	helper.GetEnvWithKey("AWS_ACCESS_KEY_ID")
	helper.GetEnvWithKey("AWS_SECRET_ACCESS_KEY")

	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	s := helper.RandStringRunes(10)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(helper.GetEnvWithKey("BUCKET_NAME")),
		Key:         aws.String(directory + "/" + s + ".png"),
		Body:        body,
		ContentType: aws.String("image/png"),
		ACL:         aws.String("public-read"),
	})

	helper.PanicIfError(err)

	return web.UploadResponse{
		ImageUrl: aws.StringValue(&result.Location),
	}
}
