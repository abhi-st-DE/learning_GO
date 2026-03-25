module github.com/abhi-st-DE/learning_GO // This is the module path for the Go project. It indicates that the code in this project belongs to the "learning_GO" module, which is hosted on GitHub under the user "abhi-st-DE".

go 1.25.4

// Package that allows us to grab environment variables from a .env file.
require github.com/joho/godotenv v1.5.1

require (
	github.com/go-chi/chi v1.5.5
	github.com/go-chi/cors v1.2.2
	github.com/google/uuid v1.6.0
)

require github.com/lib/pq v1.12.0
