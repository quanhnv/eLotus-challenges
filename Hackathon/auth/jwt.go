package jwtHelper

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateAndSetCookie(w http.ResponseWriter, userName string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	expiredTime, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_TIME"))

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = userName
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expiredTime)).Unix()
	fmt.Println("Expired = ", time.Now().Add(time.Hour*time.Duration(expiredTime)*time.Second))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	finalToken, err := token.SignedString([]byte(secretKey))

	setCookie(w, finalToken)

	return finalToken, err
}

func Verify(r *http.Request) error {
	secretKey := os.Getenv("SECRET_KEY")

	tokenString, err := getTokenFromCookie(r)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}

	//Check token expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return errors.New("Token expired")
		}
	}

	return nil
}

func ExtractUsernameFromToken(r *http.Request) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	var username string
	tokenString, err := getTokenFromCookie(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return username, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username = fmt.Sprintf("%v", claims["username"])
	}

	return username, nil
}

func getTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func setCookie(w http.ResponseWriter, tokenString string) {
	expiredTime, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_TIME"))
	// Thiết lập cookie với JWT token
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * time.Duration(expiredTime)),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}
