Реализация gRPC обёртки над сайтом https://www.rusprofile.ru/

## API

Сервис реализовывает один метод, принимающий на вход ИНН компании, ищущий компанию на rusprofile, и возвращающий её ИНН, КПП, название, ФИО руководителя

## Технологии

* сервис написан на Go с использованием Go Modules
* предоставляет API через [gRPC](https://grpc.io/docs/languages/go/quickstart/)
* предоставляет API через HTTP с помощью [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
* предоставляет Swagger UI с документацией, сгенерированной из .proto файла с помощью protoc-gen-swagger
* упакован в docker контейнер