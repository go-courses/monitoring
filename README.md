Instruction to install and run

- Устанавливаем grpc по офиц документации
- Тянем библиотеку grpc-gateway и навсяк библиотеку swagger grpc:
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
- И в нашей папке запустим:
    make all

    в результате чего будут такие действия:
        - чистка сгенерированных файлов
        - генерация proto файлов
        - генерация сервера в bin/server и консольного клиента bin/client

- Для запуска:
    в терминале с рабочей папки запустим сервер: bin/server и "Enter";
    на втором терминале если запустим bin/client получим все ответы с gRPC сервера за раз;
    а для REST сервера, с браузера нужно открыть адреса (4 шт endpoints):
        localhost:7778/disk
        localhost:7778/cpu
        localhost:7778/ram
        localhost:7778/net

        Ответы получим в виде JSON.

TODO: осталось доделать так чтоб при открытии корневого endpointa (localhost:7778/) выводились все ответы разом ...