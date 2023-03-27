package controller

import (
	"encoding/json"
	"fmt"
	"go-postgres-crud/libs"
	"go-postgres-crud/models" //models package dimana Buku didefinisikan
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	_ "github.com/lib/pq" // postgres golang driver
)

type responseUpload struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type responsesImage struct {
	Data []models.Image `json:"data"`
}

type ResponseUpload struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Book `json:"data"`
}

func UploadAsset(w http.ResponseWriter, r *http.Request) {
	file, _, _ := r.FormFile("file")

	if file == nil {

	}
	files := r.MultipartForm.File["file"]
	// create a channel to receive responses from goroutines
	resChan := make(chan responseUpload, len(files))
	sizeAllFile := 0 // initialize a counter variable
	for _, fileHeader := range files {
		go func(fileHeader *multipart.FileHeader) {
			filee, err := fileHeader.Open()
			if err != nil {
				log.Println(err)
				resChan <- responseUpload{Message: "Failed to read file"}
				return
			}
			defer filee.Close()

			fileBytes, err := ioutil.ReadAll(filee)
			if err != nil {
				log.Println(err)
				resChan <- responseUpload{Message: "Failed to read file"}
				return
			}
			sizeAllFile += int(fileHeader.Size)
			// Upload the file to S3
			var randomStr = libs.RandStringBytes(12) + fileHeader.Filename
			err = libs.UploadFile("lars-storage", randomStr, fileBytes)
			if err != nil {
				log.Println(err)
				resChan <- responseUpload{Message: "Failed to upload file to S3"}
				return
			}
			var filepath = "https://" + "lars-storage" + "." + "is3.cloudhost.id/" + randomStr
			fmt.Println("File successfully uploaded to S3! :", randomStr)

			insertID := models.AddAssets(randomStr, filepath)
			res := responseUpload{
				ID:      insertID,
				Message: "Success Upload File: " + filepath,
			}
			resChan <- res
		}(fileHeader)
	}

	// create a slice to hold the responses
	var resList []responseUpload

	// wait for all the goroutines to complete and receive their responses
	for i := 0; i < len(files); i++ {
		res := <-resChan
		resList = append(resList, res)
	}

	// close the response channel
	close(resChan)
	fmt.Println("Final value of counter:", sizeAllFile)
	// return the list of responses
	json.NewEncoder(w).Encode(resList)
}
