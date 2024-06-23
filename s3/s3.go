package s3

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func NewS3Client() (*minio.Client, error) {

	ctx := context.Background()
	endpoint := viper.GetString("S3_ENDPOINT")
	accessKeyID := viper.GetString("S3_ACCESS_KEY")
	secretAccessKey := viper.GetString("S3_SECRET_KEY")
	useSSL := viper.GetBool("S3_USE_SSL")
	bucketName := viper.GetString("S3_BUCKET")
	location := viper.GetString("S3_LOCATION")

	// Initialize minio client object.
	minioClient, err := minio.New(
		endpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				accessKeyID,
				secretAccessKey,
				"",
			),
			Secure: useSSL,
		},
	)
	if err != nil {
		return nil, nil
	}

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	return minioClient, nil
}

func UploadFile(client *minio.Client, fileName, filePath string) error {

	ctx := context.Background()
	bucketName := viper.GetString("S3_BUCKET")

	info, err := client.FPutObject(
		ctx,
		bucketName,
		fileName,
		filePath,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", fileName, info.Size)

	return nil
}
