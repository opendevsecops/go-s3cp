package main

import (
	"bufio"
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	fromPtr := flag.String("from", "", "From source")
	toPtr := flag.String("to", "", "To destination")

	flag.Parse()

	if strings.HasPrefix(*fromPtr, "s3://") {
		u, err := url.Parse(*fromPtr)

		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(*toPtr)

		if err != nil {
			log.Fatal(err)

			return
		}

		downloader := s3manager.NewDownloader(session.New())

		_, err = downloader.Download(f, &s3.GetObjectInput{
			Bucket: aws.String(u.Host),
			Key:    aws.String(u.Path),
		})

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					log.Fatal(aerr)
				}
			} else {
				log.Fatal(err)
			}

			return
		}
	} else if strings.HasPrefix(*toPtr, "s3://") {
		u, err := url.Parse(*toPtr)

		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Open(*fromPtr)

		if err != nil {
			log.Fatal(err)

			return
		}

		uploader := s3manager.NewUploader(session.New())

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(u.Host),
			Key:    aws.String(u.Path),
			Body:   aws.ReadSeekCloser(bufio.NewReader(f)),
		})

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					log.Fatal(aerr)
				}
			} else {
				log.Fatal(err)
			}

			return
		}
	} else {
		log.Fatal("unsupported operation")
	}
}
