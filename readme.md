## 🚗 Car Inventory API — Go + Fiber + PostgreSQL

        A lightweight RESTful API built using Go (Fiber framework) to manage car inventory data.
        This project demonstrates real-world backend design — structured routes, PostgreSQL integration, error handling, testing, and containerization — ideal for         DevOps and Platform Engineering demonstrations.
---
## 🧩 Tech Stack
        Component	Technology
        Language	Go (v1.21+)
        Web Framework	Fiber v2
        Database	PostgreSQL
        DB Driver	database/sql
        Testing	Go testing package + stretchr/testify
        Benchmarking	Go -bench and -benchmem
        Containerization	Docker & Docker Compose
---
## ⚙️ Features

        ✅ RESTful CRUD operations for cars
        ✅ Structured Fiber handlers and routes
        ✅ PostgreSQL integration with connection pooling
        ✅ Unit and benchmark tests
        ✅ Graceful error handling and input validation
        ✅ Fully containerized with Docker & Docker Compose
---
## 🧱 API Endpoints
        Method 	   Endpoint	     Description
        POST	     /cars	        Add a new car to inventory
        GET	       /cars/:id	    Retrieve car details by ID
        PUT	       /cars/:id	    Update car details
        DELETE	   /cars/:id	    Delete a car from inventory
---
## 🧰 Local Setup (Manual Run)
### 1️⃣ Clone the repository
        git clone https://github.com/Aashritha123-lab/car_Inventory_Fiber.git
        cd car_Inventory_Fiber

### 2️⃣ Configure PostgreSQL connection

    Update the DSN in config/config.go:

    dsn := "user=postgres password=Helper@123 dbname=postgres sslmode=disable"

    Or use environment variables:

    export DB_USER=postgres
    export DB_PASS=Helper@123
    export DB_NAME=postgres
    export DB_HOST=localhost

### 3️⃣ Run the server
    go run main.go
    Server will start at 👉 http://127.0.0.1:3051

## 🧪 Testing & Benchmarking
    Run unit tests
    go test ./...

    Run performance benchmarks
    go test -benchmem . -bench .

## 🐳 Docker Setup
        Dockerfile
        # ---------- Stage 1: Build ----------
        FROM golang:1.22-alpine AS builder
        WORKDIR /app
        COPY go.mod go.sum ./
        RUN go mod download
        COPY . .
        RUN go build -o car_inventory .
        
        # ---------- Stage 2: Runtime ----------
        FROM gcr.io/distroless/base-debian12
        WORKDIR /app
        COPY --from=builder /app/car_inventory .
        EXPOSE 3051
        ENTRYPOINT ["./car_inventory"]

## 🧩 Docker Compose Setup
        docker-compose.yml
            version: "3.9"

    services:
      app:
        build: .
        container_name: car_inventory_api
        ports:
          - "3051:3051"
        environment:
          - DB_USER=postgres
          - DB_PASS=Helper@123
          - DB_NAME=car_inventory
          - DB_HOST=db
          - DB_PORT=5432
        depends_on:
          - db
    
      db:
        image: postgres:15-alpine
        container_name: car_inventory_db
        restart: always
        environment:
          POSTGRES_USER: postgres
          POSTGRES_PASSWORD: Helper@123
          POSTGRES_DB: car_inventory
        ports:
          - "5432:5432"
        volumes:
          - pgdata:/var/lib/postgresql/data
    
    volumes:
      pgdata:

## 🐳 Run the entire stack
    docker-compose up --build
  
    Once running, access the API at 👉 http://localhost:3051/cars
    
    To stop and clean up:
    
    docker-compose down

## 📊 Example Request

    POST /cars
    
    {
      "Name": "BMW X7",
      "Model": "V8",
      "Brand": "BMW",
      "Year": 2023,
      "Price": 20000000
    }
    
    
    Response
      
      {
        "message": "Car added successfully",
        "id": 13
      }

## 🧠 Learning Focus

    This project demonstrates:
    Backend microservice development in Go
    Integration of Fiber with PostgreSQL
    Writing unit tests and benchmarks
    Secure and minimal distroless Docker images
    Container orchestration with Docker Compose
    A DevOps-ready setup for CI/CD integration

### 📎 Author
    👩‍💻 Aashritha Chakka

