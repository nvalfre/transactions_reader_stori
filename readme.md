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
   - email_service/config.go
    - SMTP Configuration:
        - `DefaultSMTPHost`: Hostname of the SMTP server
        - `DefaultSMTPPort`: Port number of the SMTP server
        - `DefaultSMTPUsername`: Username for authenticating with the SMTP server
        - `DefaultSMTPPassword`: Password for authenticating with the SMTP server

    - Sender Details:
        - `DefaultSenderName`: Name of the email sender
        - `DefaultSenderEmail`: Email address of the email sender

4. Build the microservice:

   ```bash
   go build
   ```

## Usage

1. Start the microservice:

   ```bash
   ./transactions-summary
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
   docker build -t <image> .
   ie: 
   docker build -t transactions_summary .
   ```

2. Run the Docker container:
   ```bash
   build:
    - docker build -t <image>:<version>
    - docker tag <image>:latest <user>/<image>:latest
    - docker push <user>/<image>:<version>
   
   on aws ec2 instance
   install docker:
    - sudo yum install docker
    - sudo usermod -a -G docker ec2-user
    - id ec2-user
    - # Reload a Linux user's group assignments to docker w/o logout
    - newgrp docker
    - sudo systemctl enable docker.service
    - sudo systemctl start docker.service
    - verify
      - sudo systemctl status docker.service
      - docker <version>
   docker run -e GIN_MODE=release -p 8080:8080 --name <image> <user>/<image>
    - logs: docker logs <image>
   ```
3. Examples:
   ```
   curl --location --request POST 'http://localhost:8080/file/process/transactions?account_id=1&name=accname&email=testmailnv23@gmail.com' \
   --header 'Content-Type: application/json' \
   --form 'file=@"/C:/Users/KTUFi5-Desk/Documents/Nico/dev/transactions_reader_stori/balance.csv"'
   ```
   
## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
