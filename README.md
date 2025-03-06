# Backend Challenge Solutions

This repository contains solutions to three programming challenges implemented in Go.

## Challenges

### 1. Maximum Sum Path

Finds the path with the maximum sum in a triangle of numbers.

### 2. Left-Right-Equal Sequence

Solves the problem of finding the minimum digit sequence 

### 3. Pie Fire Dire API

A REST API service that handles beef-related data.
- Built using Gin framework
- Runs on port 3002

## Running the Application

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/main.go
```

## API Endpoints

- GET `/` - Returns beef summary data

## Project Structure
```
.
├── cmd
│   └── main.go
└── pkg
    ├── leftrightequal
    ├── maxsumpath
    └── piefiredire
```

## Dependencies

- gin-gonic/gin - Web framework
- Standard Go libraries
