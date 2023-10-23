# eLotus-challenges
Exam content: https://github.com/elotusteam/challenges/blob/main/backend.md
includes:
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

## How project is structured

Structure look like

![image](https://github.com/quanhnv/eLotus-challenges/assets/51664950/835b7eaf-0c41-4b78-b9e2-ad2be8f16f47)

We have:
- Auth folder: where contain file that process Authentication of the application. Currently, we only implement simple auth by jwt.
- Middlewares: where contain middlewares that between client and server. Currently, inside that folder, we defined Check-jwt to validate Request.
- Database: where contain func as layer after daâtbase, to application work with database without caring about specific database type.
- Routes: where contain api handler.
- Tmp: file storage.
- .env: File contain app's configuration.
- eLotus-hackathon-golang-upload-app.postman_collection.json: Postman sample request.

## Installation & Run

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
