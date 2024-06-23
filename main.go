package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"

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

		sampleData, err := mariaDB.ListSampleData(db)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(sampleData)

		for _, row := range sampleData {
			fmt.Println(row)
		}

	}
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	if len(sqsEvent.Records) == 0 {
		return errors.New("no SQS message passed to function")
	}

	for _, msg := range sqsEvent.Records {
		fmt.Printf("Got SQS message %q with body %q\n", msg.MessageId, msg.Body)
		// TODO: Add application logic here
	}

	return nil
}
