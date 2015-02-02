package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/**
 * Uploads a folder to an s3 bucket
 */

func main() {
	bucketName := requiredArg("bucket", "REQUIRED", "bucket to upload to")
	prefix := requiredArg("prefix", "REQUIRED", "prefix to upload into")
	folder := requiredArg("folder", "REQUIRED", "folder to upload")
	flag.Parse()
	err := uploadDir(bucketName(), folder(), prefix())
	if err != nil {
		log.Fatal(err)
	}
}

func requiredArg(key string, defaultValue string, desc string) func() string {
	value := flag.String(key, defaultValue, desc)
	return func() string {
		if *value == "REQUIRED" {
			log.Fatal(errors.New("value required for " + key))
		}
		return *value
	}
}

func uploadDir(bucketName string, folder string, prefix string) error {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.USEast)
	bucket := client.Bucket(bucketName)
	visit := func(filename string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			return upload(bucket, path.Join(prefix, strings.TrimPrefix(filename, folder)), filename)
		}
		return nil
	}

	return filepath.Walk(folder, visit)
}

func upload(bucket *s3.Bucket, key, filename string) error {
	contentType := mime.TypeByExtension(filepath.Ext(filename))
	fmt.Printf("upload '%s' -> '%s' (%s)\n", filename, key, contentType)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return bucket.Put(key, data, contentType, s3.BucketOwnerFull)
}
