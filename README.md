# Go Echo Boilerplate

Go Boilerplate with Echo Framework and GORM

## Installation

```
git clone https://github.com/farhanmn/go-echo-boilerplate
cd go-echo-boilerplate
## rename .env.sample to .env
## adjust the value inside as you needed
go get install
go run cmd/main.go
```

## Folder Structure

```
.
├── /cmd                    # main file for app running
│ └── main.go               # app entrypoint
├── /config                 # app configuration
│ ├── config.go             # config DB
│ └── migration.go          # migration config
├── /internal               # main module
│ ├── controller            # controller
│ ├── model                 # DB model/ structure
│ ├── repository            # DB layer
│ ├── routes                # HTTP handler
│ ├── service               # business logic
├── /pkg                    # additional packages
│ ├── middleware            # custom middleware
│ ├── response              # API response std
│ ├── utils                 # Utility/ helper func
├── /test                   # unit test
├── .env.sample             # environment variable
├── .gitignore              # file ignore
├── go.mod                  # go module file
├── go.sum                  # dependency list
└── README.md
```

## Environment Variables

`DB_HOST`
`DB_USER`
`DB_PASS`
`DB_NAME`
`DB_PORT`

## Tech

- [Echo](https://echo.labstack.com/) - High performance, extensible, minimalist Go web framework
- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- [Logrus](https://github.com/sirupsen/logrus) - Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger
- [Color](https://pkg.go.dev/github.com/fatih/color) - Color lets you use colorized outputs in terms of ANSI Escape Codes in Go (Golang). It has support for Windows too! The API can be used in several ways, pick one that suits you
- [JWT](https://jwt.io/) - a compact URL-safe means of representing claims to be transferred between two parties

## Script

- `go run cmd/main.go --migrate` - Run migrations
- `go run cmd/main.go` - Run the application

## Authors

- [@farhamn](https://github.com/farhanmn)
