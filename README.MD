rodar localstack -> docker run --rm -it -d -p 4566:4566 localstack/localstack start

criar fila -> aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name catalog-emit