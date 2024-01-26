package services

import (
	"bytes"
	"context"
	"encoding/base64"

	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/fabiolaguna/twitter-go/dao"
	"github.com/fabiolaguna/twitter-go/models"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.Response {
	var response models.Response
	response.Status = 500
	userId := claim.Id.Hex()
	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))

	var fileName string
	var user models.User

	switch uploadType {
	case "avatar":
		fileName = "avatars/" + userId + ".jpg"
		user.Avatar = fileName
	case "banner":
		fileName = "banners/" + userId + ".jpg"
		user.Banner = fileName
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		response.Message = "[image service][method:UploadImage] Error has occurred parsing header: " + err.Error()
		fmt.Println(response.Message)
		return response
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			response.Message = "[image service][method:UploadImage] Error decoding string: " + err.Error()
			fmt.Println(response.Message)
			return response
		}

		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		p, err := mr.NextPart()
		if err != nil && err != io.EOF {
			response.Message = "[image service][method:UploadImage] Error trying to access next part: " + err.Error()
			fmt.Println(response.Message)
			return response
		}

		if err != io.EOF {
			if p.FileName() != "" {
				buffer := bytes.NewBuffer(nil)
				if _, err := io.Copy(buffer, p); err != nil {
					response.Message = "[image service][method:UploadImage] Error copying to buffer: " + err.Error()
					fmt.Println(response.Message)
					return response
				}

				sess, err := session.NewSession(&aws.Config{
					Region: aws.String("us-east-1")})
				if err != nil {
					response.Message = "[image service][method:UploadImage] Error creating new aws session: " + err.Error()
					fmt.Println(response.Message)
					return response
				}

				uploader := s3manager.NewUploader(sess)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(fileName),
					Body:   &readSeeker{buffer},
				})
				if err != nil {
					response.Message = "[image service][method:UploadImage] Error uploading image to S3: " + err.Error()
					fmt.Println(response.Message)
					return response
				}
			}
		}

		status, err := dao.UpdateProfile(user, userId)
		if err != nil || !status {
			response.Message = "[image service][method:UploadImage] Error updating profile: " + err.Error()
			fmt.Println(response.Message)
			return response
		}
	} else {
		response.Status = 400
		response.Message = "[image service][method:UploadImage] Incorrect Content-Type header, must be 'multipart/'"
		fmt.Println(response.Message)
		return response
	}

	response.Status = 200
	response.Message = "Image uploaded"
	return response
}
