structure of a golang web server

project/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── api/                  # API layer
│   │   ├── handlers/         # Request handlers
│   │   │   ├── auth.go
│   │   │   └── user.go
│   │   ├── middleware/       # Custom middleware
│   │   │   ├── auth.go
│   │   │   └── logging.go
│   │   └── routes/           # Route definitions
│   │       └── router.go
│   ├── config/               # Configuration
│   │   └── config.go
│   ├── models/               # Data models
│   │   └── user.go
│   ├── repository/           # Data access layer
│   │   └── user_repository.go
│   └── service/              # Business logic
│       └── user_service.go
├── pkg/                      # Shared packages
│   ├── utils/
│   │   └── response.go
│   └── validator/
│       └── validator.go
├── .env                      # Environment variables
├── go.mod                    # Go modules
└── go.sum                    # Dependencies checksum