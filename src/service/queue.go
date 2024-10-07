package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FreitasGabriel/anotai-test-consumer/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test-consumer/src/repository"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.mongodb.org/mongo-driver/mongo"
)

const QUEUE_NAME_ENV = "QUEUE_NAME"

var queue_name = ""

func initSQSSession() *session.Session {
	queue_name = os.Getenv(QUEUE_NAME_ENV)

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
	}))

	return sess

}

func ReceivedMessageFromQueue(svc *sqs.SQS) (*sqs.ReceiveMessageOutput, error) {
	receiveprams := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queue_name),
		WaitTimeSeconds:     aws.Int64(20),
		MaxNumberOfMessages: aws.Int64(1),
	}

	result, err := svc.ReceiveMessage(receiveprams)
	return result, err
}

func deleteMessageFromQueue(svc *sqs.SQS, msg *sqs.Message) error {
	deleteprams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queue_name),
		ReceiptHandle: msg.ReceiptHandle,
	}

	_, err := svc.DeleteMessage(deleteprams)
	return err
}

func InitQueue(database *mongo.Database) {
	sess := initSQSSession()
	repo := repository.NewCatalogRepository(database)
	service := NewCatalogService(repo)
	svc := sqs.New(sess)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case <-signalCh:
			fmt.Println("Saindo")
			return
		default:

			result, err := ReceivedMessageFromQueue(svc)
			if err != nil {
				logger.Error(fmt.Sprintf("error to receive message from queue %s ", queue_name), err)
				time.Sleep(1 * time.Second)
				continue
			}

			logger.Info("looking for new messages...")
			if len(result.Messages) > 0 {
				for _, msg := range result.Messages {
					logger.Info(fmt.Sprintf("Received message: %s \n", *msg.Body))
					service.PublishCatalog()
					err = deleteMessageFromQueue(svc, msg)
					if err != nil {
						logger.Error(fmt.Sprintf("error to delete message from queue %s ", queue_name), err)
						time.Sleep(1 * time.Second)
						continue
					}
				}
			}
		}
	}
}
