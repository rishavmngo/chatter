# Chatter app
- A backend for chat app made using golang and postgres
- Decoupled databases and controller

## Directory structure
.
├── api
│   ├── server.go
│   └── user
│       ├── user.controller.go
│       └── user.routes.go
├── go.mod
├── go.sum
├── intrf
│   └── storage.go
├── jwtUtils
│   └── jwtUtils.go
├── main.go
├── README.md
├── storage
│   └── postgres.go
├── types
│   ├── Nullablestring.go
│   └── user.go
└── utils
    └── utils.go
