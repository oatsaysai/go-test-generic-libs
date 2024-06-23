package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/minio/minio-go/v7"
	"github.com/oatsaysai/go-test-generic-libs/mariaDB"
	s3Client "github.com/oatsaysai/go-test-generic-libs/s3"
	"github.com/spf13/viper"
)

var db *sql.DB
var s3 *minio.Client

func init() {
	viper.AutomaticEnv()
}

func main() {
	// Init instances
	var err error
	db, err = mariaDB.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s3, err = s3Client.NewS3Client()
	if err != nil {
		log.Fatal(err)
	}

	if runtime_api, _ := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); runtime_api != "" {
		lambda.Start(handler)
	} else {
		createReport001()
	}
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	if len(sqsEvent.Records) == 0 {
		return errors.New("no SQS message passed to function")
	}

	for _, msg := range sqsEvent.Records {
		fmt.Printf("Got SQS message %q with body %q\n", msg.MessageId, msg.Body)
		// TODO: Add application logic here
		createReport001()
	}

	return nil
}

func createReport001() {
	// Get data from DB
	// Gen to CSV
	genReport001()

	// Upload to S3
	err := s3Client.UploadFile(
		s3,
		"report_001.csv",
		"report_001.csv",
	)
	if err != nil {
		log.Fatal(err)
	}
}

func genReport001() {
	sampleData, err := mariaDB.ListSampleData(db)
	if err != nil {
		log.Fatal(err)
	}

	// Gen CSV
	file, err := os.Create("report_001.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	// Write CSV header
	w.Write([]string{
		"name",
		"data_001",
		"data_002",
		"created_time",
		"updated_time",
	})

	for _, data := range sampleData {
		row := []string{
			data.Name,
			data.Data_001,
			data.Data_002,
			data.CreatedTime.Local().String(),
			data.UpdatedTime.Local().String(),
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
