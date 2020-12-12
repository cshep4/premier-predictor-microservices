#!/bin/sh
aws sqs create-queue --queue-name=EmailQueue --profile localstack --endpoint-url=http://localhost:4576