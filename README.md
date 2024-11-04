## Go Worker Pool App
This is a Go application using a pool of workers that performs tasks in parallel using a common buffered channel and displays the results of execution.

## Getting Started
Here's how to get the project up and running on your local machine for development and testing.

### Prerequisites
- Docker
- Docker Compose

### Installation

1. Clone the repository:
    ```bash
    git https://github.com/forsitet/worker-pool.git
    cd worker-pool
    ```
2. Build and run with Docker Compose:
    ```bash
    docker-compose up --build
    ```
    This spins up the following services:
    - `app`: The main marketplace service application.

### Usage

The application will be available at `http://localhost:8080` after startup.
