# Transaction Summary Microservice

The Transaction Summary Microservice is a Go-based application that processes a file containing debit and credit transactions on an account and sends a summary of the transactions via email. It provides an API endpoint to upload the transaction file and generate the summary.

## Features

- Upload a file containing debit and credit transactions
- Process the file and save transaction and account information to a database
- Generate a summary of the transactions including total balance, number of transactions grouped by month, and average credit and debit amounts grouped by month
- Send the summary information to a user via email

## Prerequisites

- Go (version 1.20 or higher)

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   ```

2. Install the required dependencies:

   ```bash
   go mod download
   ```

3. Set up the necessary environment variables:

    - SMTP Configuration:
        - `SMTP_HOST`: Hostname of the SMTP server
        - `SMTP_PORT`: Port number of the SMTP server
        - `SMTP_USERNAME`: Username for authenticating with the SMTP server
        - `SMTP_PASSWORD`: Password for authenticating with the SMTP server

    - Sender Details:
        - `SENDER_NAME`: Name of the email sender
        - `SENDER_EMAIL`: Email address of the email sender

4. Build the microservice:

   ```bash
   go build
   ```

## Usage

1. Start the microservice:

   ```bash
   ./transaction-summary-microservice
   ```

2. Upload a transaction file using the API endpoint:

   ```
   POST /process-file
   ```

    - Payload: `file` (multipart/form-data)

3. The microservice will process the file, save the transaction and account information to the database, generate a summary, and send an email to the specified recipient.

## Deployment

The microservice can be deployed using various methods such as containerization with Docker, running it as a standalone binary, or using a container orchestration platform like Kubernetes or Amazon ECS. Choose the deployment method that best suits your infrastructure and requirements.

### Docker

1. Build the Docker image:

   ```bash
   docker build -t transaction-summary-microservice .
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 -e SMTP_HOST=<host> -e SMTP_PORT=<port> -e SMTP_USERNAME=<username> -e SMTP_PASSWORD=<password> -e SENDER_NAME=<sender-name> -e SENDER_EMAIL=<sender-email> transaction-summary-microservice
   ```

   Replace `<host>`, `<port>`, `<username>`, `<password>`, `<sender-name>`, and `<sender-email>` with the actual values.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
