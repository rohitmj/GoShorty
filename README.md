# GoShorty - Simple URL Shortener

GoShorty is a lightweight URL shortener built in Go (Golang), designed for learning DevOps, backend, and cloud deployment best practices.

## Features

- Shorten long URLs to short, unique links
- Fast HTTP backend with Go and [gorilla/mux](https://github.com/gorilla/mux)
- Simple HTML/JS frontend for easy testing
- Dockerized for easy local development and deployment
- Ready for CI/CD and Google Cloud Run deployment

## Quick Start

1. **Clone the repository:**
    ```bash
    git clone https://github.com/rohitmj/GoShorty.git
    cd GoShorty
    ```

2. **Run locally (requires Go):**
    ```bash
    go run main.go
    # App runs at http://localhost:8080
    ```

3. **Build and run with Docker:**
    ```bash
    docker build -t goshorty .
    docker run -p 8080:8080 goshorty
    ```

4. **Access the frontend:**
    - Open your browser and go to `http://localhost:8080`
    - Paste a long URL, click “Shorten,” and get a short link

## Project Structure

  - handlers/ # Go HTTP handler functions (API)
  - static/ # Frontend (index.html, CSS, JS)
  - utils/ # Helper Go files
  - main.go # Main entrypoint
  - go.mod # Go modules/dependencies
  - Dockerfile # Container build instructions
  - README.md


## Learning Goals

- Learn Go web development basics
- Understand how REST APIs work
- Practice Dockerizing Go apps
- Prepare for CI/CD and cloud deployment (Google Cloud Run)

## Author

- Rohit Maharjan  
  [LinkedIn](https://www.linkedin.com/in/rmj619) • [GitHub](https://github.com/rohitmj)

---

> **This project was built as part of my DevOps & SRE learning journey.**



