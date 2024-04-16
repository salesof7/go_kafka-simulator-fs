[![Banner](/placeholder.jpg)](https://github.com/xavimondev/supaplay)
<p align="center">
  <img src="https://img.shields.io/github/languages/top/salesof7/go_kafka-simulator-fs "Language"" alt=" Language" />
  <img src="https://img.shields.io/github/stars/salesof7/go_kafka-simulator-fs "Stars"" alt=" Stars" />
  <img src="https://img.shields.io/github/issues-pr/salesof7/go_kafka-simulator-fs "Pull Requests"" alt=" Pull Requests" />
  <img src="https://img.shields.io/github/issues/salesof7/go_kafka-simulator-fs "Issues"" alt=" Issues" />
  <img src="https://img.shields.io/github/contributors/salesof7/go_kafka-simulator-fs "Contributors"" alt=" Contributors" />
</p>

## Table of Contents

- [Stack](#stack)
- [Project Summary](#project-summary)
- [Setting Up](#setting-up)
- [Run Locally](#run-locally)
- [FAQ](#faq)

## Stack

- [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go): Library for Kafka integration, essential for data fetching and server-client communication.
- [godotenv](https://github.com/joho/godotenv): Package for loading environment variables, crucial for configuration and authentication.

## Project Summary

- [app](app): Main application code and logic implementation.
- [app/routes](app/routes): Contains routing configurations for the application.
- [infra](infra): Infrastructure-related code and configurations.
- [infra/kafka](infra/kafka): Infrastructure configurations specific to Kafka.
- [destinations](destinations): Code related to defining destination endpoints.
- [.docker](.docker): Contains Docker configurations for the project.
- [.docker/kafka](.docker/kafka): Docker configurations specific to Kafka setup.
- [app/kafka](app/kafka): Code related to Kafka integration in the application.

## Setting Up

Insert your environment variables.

## Run Locally

1. Clone go_kafka-simulator-fs repository:  
```bash  
git clone https://github.com/salesof7/go_kafka-simulator-fs  
```
2. Install the dependencies with one of the package managers listed below:  
```bash  
go build -o myapp  
```
3. Start the development mode:  
```bash  
go run main.go  
```

## Contributors

[![Contributors](https://contrib.rocks/image?repo=salesof7/go_kafka-simulator-fs)](https://github.com/salesof7/go%5Fkafka-simulator-fs/graphs/contributors)

## FAQ

#### 1.What is this project about?

This project aims to **briefly describe your project's purpose and goals.**

#### 2.How can I contribute to this project?

Yes, we welcome contributions! Please refer to our [Contribution Guidelines](CONTRIBUTING.md) for more information on how to contribute.

#### 3.What is this project about?

Your answer.
