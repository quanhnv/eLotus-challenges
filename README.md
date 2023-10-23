# eLotus-challenges
- DSA tasks
- Hackathon - Simple upload file app

## Simple upload file application by Golang

## Features

- Simple JWT authentication (HS256)
    - User register by Username & Password
    - User login
    - Revoke token by time
- API handler a form upload file
    - Validate request & file data
    - Save file to folder
    - Store file metadata to database (SQLite)

## Tech

- Vanilla Golang
- Postman

## Installation

Requires [Go](https://go.dev/dl/) v1.21.3 to run.

```sh
cd hackathon
go mod tidy
go run main.go
```

## Test

Any api client tool like [Postman](https://www.postman.com/) to run.

```sh
Import file https://github.com/quanhnv/eLotus-challenges/blob/main/Hackathon/eLotus-hackathon-golang-upload-app.postman_collection.json to Postman
Request to test
```
