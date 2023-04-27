package handlers

import (
	"aspera-delivery/config"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var cfg config.Config

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read : %v", err)
	}

}

func Uploader(c echo.Context) error {

	UploadEventResult, err := returnUploadEventInfo(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, UploadEventResult)
}

func returnUploadEventInfo(c echo.Context) (UploadEventResponse, error) {

	// take input from form-data

	var credentials Credentials

	//TODO: Credentials(username,password)
	// TODO: Credentials(username,privatekey)
	credentials.Username = c.FormValue("username")

	host := c.FormValue("host")
	port := c.FormValue("port")
	authtype := c.FormValue("auth_type")
	// sourcepath := c.FormValue("source_path")
	destinationpath := c.FormValue("destination_path")

	scpath := c.FormValue("source_pathS3URI")
	sourcepath, err := downloadFromS3URI(scpath)

	if err != nil {
		fmt.Println("Failed to download the source file")
		return UploadEventResponse{}, err
	}
	defer os.RemoveAll(sourcepath)

	if authtype == "password" {
		credentials.Password = c.FormValue("password")
	} else if authtype == "privatekey" {
		credentials.KeyS3URI = c.FormValue("keyS3URI")
	}

	var command string

	if authtype == "password" {
		command = fmt.Sprint("ASPERA_SCP_PASS='", credentials.Password, "'", " ascp -m 100m -l 1000m -d ", sourcepath, " ", credentials.Username+"@"+host+":"+destinationpath)
	} else if authtype == "privatekey" {

		keyfilepath, err := downloadFromS3URI(credentials.KeyS3URI)
		if err != nil {
			log.Fatal("Failed to download key")
			return UploadEventResponse{}, err
		}

		test_path := keyfilepath
		defer os.Remove(keyfilepath)
		command = fmt.Sprint("ascp", " -m 100m -l 1000m -d -i ", test_path, " -P ", port, " ", sourcepath, " ", credentials.Username+"@"+host+":"+destinationpath)
	}

	cmd := exec.Command("sh", "-c", command)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		fmt.Println("Command output:", stdout.String())
		fmt.Println("Command error:", stderr.String())
		return UploadEventResponse{}, err
	}

	output := stdout.String()
	if output == " " {
		output = "Upload comepleted successfully"
	}
	uploadEventResponse := UploadEventResponse{
		FileName:        filepath.Base(scpath),
		SourcePath:      scpath,
		DestinationPath: destinationpath,
		OutputResult:    string(output),
	}

	return uploadEventResponse, nil

}
func downloadFromS3URI(s3URI string) (string, error) {

	// Parse the S3 URI
	uri, err := url.Parse(s3URI)
	if err != nil {
		fmt.Println("Failed to Parse s3URI")
		return "Failed to Parse s3URI", err
	}

	// Create a new AWS session with explicit credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, cfg.AWSSessionToken),
	})
	if err != nil {
		fmt.Println("Failed to create the session")
		return "Failed to create the session", err
	}

	// Set the default download directory
	downloadDir,_:=filepath.Abs("temp")
	// downloadDir := "/home/ajits/Desktop/Aspera_Delivery_Service/temp"
	// Create an S3 downloader
	downloader := s3manager.NewDownloader(sess)

	// Download the object from S3 directly to the default download directory
	tempfilePath := filepath.Join(downloadDir, filepath.Base(uri.Path))

	file, err := os.Create(tempfilePath)
	if err != nil {
		fmt.Println("Failed to create file path", "--", "Err:", err)
		return "Failed to create file path", err
	}
	defer file.Close()
	// host := uri.Host
	// path := uri.Path

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(uri.Host),
		Key:    aws.String(strings.TrimPrefix(uri.Path, "/")),
		// Bucket: aws.String(host),
		// Key:    aws.String(strings.TrimPrefix(path, "/")),

	})
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to start the session")
		return "Failed to start the session", err
	}

	return tempfilePath, err
}
