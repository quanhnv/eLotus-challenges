package sqliteHelper

import (
	"database/sql"
	"errors"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
}

type Image struct {
	Id           int
	Name         string
	FilePath     string
	ImageType    string
	Size         int
	UploadedUser string
	UploadedDate time.Time
}

func InsertUser(userName string, passWord string) error {
	dbFilePath := os.Getenv("SQLITE_DB_FILE_PATH")
	// Connect to sqlite
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return err
	}
	defer db.Close() //Alway close connection after finish func

	hashPw, err := hash(passWord)
	if err != nil {
		return err
	}

	//insert to db
	_, err = db.Exec("INSERT INTO User (Username, Password) VALUES (?, ?)", userName, hashPw)
	if err != nil {
		return err
	}
	return nil
}

func Login(userName string, passWord string) error {
	dbFilePath := os.Getenv("SQLITE_DB_FILE_PATH")
	// Connect to sqlite
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return err
	}
	defer db.Close() //Alway close connection after finish func

	statement := "SELECT * FROM User WHERE Username = ? LIMIT 1"
	var user User
	err = db.QueryRow(statement, userName).Scan(&user.Username, &user.Password)
	if err != nil {
		return err
	}

	checkPass := checkPasswordHash(user.Password, passWord)
	if checkPass != nil {
		return errors.New("Password wrong")
	}

	return nil
}

func InsertImage(image Image) error {
	dbFilePath := os.Getenv("SQLITE_DB_FILE_PATH")
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO Image (Name, FilePath, ImageType, Size, UploadedUser, UploadedDate) VALUES (?, ?, ?, ?, ?, ?)", image.Name, image.FilePath, image.ImageType, image.Size, image.UploadedUser, image.UploadedDate)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
