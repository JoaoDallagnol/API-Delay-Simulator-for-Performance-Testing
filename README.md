# API Delay Simulator for Performance Testing

## Overview

This project provides a simple API simulator that introduces configurable delays and random failures to simulate real-world conditions for testing purposes. It allows developers to test how their systems handle delayed responses and failures, which is crucial for performance optimization and resilience testing.

The project has three main endpoints:
- `/delaye`: Simulates a fixed delay of 5 seconds before returning a response.
- `/unstable`: Simulates an unstable endpoint where there's a 30% chance of failure (HTTP 500).
- `/custom-delay`: Accepts a query parameter to simulate a custom delay in seconds.
- `/timeout`: Accepts a query parameter to simulate a timeout error in seconds.

The API is built using Go's standard `net/http` library to keep it lightweight and straightforward, making it a perfect starting point for performance tests.

## Project Structure

- `main.go`: The entry point that starts the HTTP server.
- `internal/handler`: Contains the logic for handling incoming requests.
- `internal/service`: Contains the business logic for introducing delays and failures.
- `internal/model`: Defines data structures used in responses.

## Requirements

- [Go 1.20+](https://golang.org/dl/)

## How to Run

1. Clone this repository:
   ```bash
   git clone https://github.com/JoaoDallagnol/API-Delay-Simulator-for-Performance-Testing.git

2. Run the Go application:
   ```bash
   go run main.go

3. The server wil start on port 8080:
   ```bash
   Server is running on port 8080

## Endpoints

### 1. Fixed Delay (5 seconds)

**Endpoint**: `/delaye`

**Description**: Introduces a fixed delay of 5 seconds before returning the response.

#### Example:
```bash
curl -X GET http://localhost:8080/delaye
```

#### Response:

```json
{
  "message": "This response was delayed by 5 seconds"
}
```

### 2. Unstable Endpoint (30% failure rate)
**Endpoint** `/unstable`

**Description**: Simutes an unstable endpoint where there's a 30% chance of failure (return HTTP 500)

#### Example:
```bash
curl -X GET http://localhost:8080/unstable
```

#### Possible Responses:

If failure occurs (30% chance):
```json
{
  "error": "Internal Server Error"
}
```

If successful (70% chance):
```json
{
  "message": "Success"
}
```

### 3. Customer Delay
**Endpoint** `/customer-delay?dealy={seconds}`

**Description**: Simulates a custom delay based on the value of the delay query parameter. If the delay parameter is missing, the default delay is 2 seconds. If the delay is invalid or negative, the response will return 400 Bad Request.

#### Example:
```bash
curl -X GET "http://localhost:8080/custom-delay?delay=10"
```

#### Possible Responses:

```json
{
  "message": "Response delayed by 10 seconds"
}
```

```json
{
  "error": "invalid delay parameter"
}
```

### 4. Timeout
**Endpoint** `/timeout?timeout={seconds}`

**Description**: Simulates a timeout after the specified duration (in seconds). The timeout query parameter defines how long the request should take before returning a 504 Gateway Timeout. If no parameter is passed, the default timeout is 5 seconds.

#### Example:
```bash
curl -X GET "http://localhost:8080/timeout?timeout=5"
```

#### Possible Response:

```json
{
  "message": "Request timed out after 5 seconds"
}
```

```json
{
  "error": "Invalid timeout parameter"
}
```
