package routes

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	jwtHelper "github.com/quanhnv/eLotus-challenges/auth"
	sqliteHelper "github.com/quanhnv/eLotus-challenges/database"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	maxImageFileSize, err := strconv.Atoi(os.Getenv("MAX_IMAGE_FILE_SIZE"))
	tmpDir := os.Getenv("IMAGE_DIR")

	//Get username from token
	userName, err := jwtHelper.ExtractUsernameFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Get, read file info - Key = data
	file, handler, err := r.FormFile("data")
	if err != nil {
		message := fmt.Sprintf("Error receiving file: %s", err.Error())
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	defer file.Close()

	//Check content file type
	contentType := handler.Header.Get("Content-Type")
	if !isImage(contentType) {
		http.Error(w, "Invalid file type. Only image files are allowed.", http.StatusBadRequest)
		return
	}

	//Check file size
	if handler.Size > int64(maxImageFileSize) {
		http.Error(w, "File size exceeds the limit of 8MB.", http.StatusBadRequest)
		return
	}

	//Check directory available
	err = os.MkdirAll(tmpDir, os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("Error creating directory: %s", err.Error())
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	//create file path
	pathFile := filepath.Join(tmpDir, handler.Filename)
	_, err = os.Stat(pathFile)
	if err == nil {
		//Generate unique file name with time
		pathFile = filepath.Join(tmpDir, fmt.Sprintf("%s_%s%s", strings.TrimSuffix(handler.Filename, filepath.Ext(handler.Filename)), time.Now().Format("02_01_2006_15_04_05"), filepath.Ext(handler.Filename)))
		fmt.Println("File path", pathFile)
	}

	//Create file
	dst, err := os.Create(pathFile)
	if err != nil {
		message := fmt.Sprintf("Error creating file: %s", err.Error())
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		message := fmt.Sprintf("Error copying file data: %s", err.Error())
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	image := sqliteHelper.Image{
		Name:         filepath.Base(handler.Filename),
		FilePath:     pathFile,
		ImageType:    contentType,
		Size:         int(handler.Size),
		UploadedUser: userName,
		UploadedDate: time.Now(),
	}
	err = sqliteHelper.InsertImage(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "File uploaded successfully")
}

func isImage(contentType string) bool {
	mimeType, _, _ := mime.ParseMediaType(contentType)
	imageMimeTypes := strings.Split(os.Getenv("IMAGE_FILE_TYPE"), ",")
	for _, mtype := range imageMimeTypes {
		if mimeType == mtype {
			return true
		}
	}
	return false
}
