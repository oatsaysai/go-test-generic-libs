package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/oatsaysai/go-test-generic-libs/mariaDB"
	"github.com/spf13/viper"
)

var db *sql.DB

func init() {
	viper.AutomaticEnv()
}

func main() {

	// Read flag
	runLambda := flag.Bool("lambda", true, "Run as lambda")
	flag.Parse()

	// Init instance
	var err error
	db, err = mariaDB.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if *runLambda {
		lambda.Start(handler)
	} else {
		genReport001()
	}
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	if len(sqsEvent.Records) == 0 {
		return errors.New("no SQS message passed to function")
	}

	for _, msg := range sqsEvent.Records {
		fmt.Printf("Got SQS message %q with body %q\n", msg.MessageId, msg.Body)
		// TODO: Add application logic here
		genReport001()
	}

	return nil
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
