# hello-backend-webapp
Simple web app

## Request

`GET /backend/`

    curl -i -H 'Accept: application/json' http://localhost:8080/backend?greeting=hello

## Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Tue, 22 Nov 2022 17:48:11 GMT
    Content-Length: 116

    {"greeting":"hello from backend","time":"2022-11-22 17:48:11.222477533 +0000 UTC m=+74.265783540","ip":"172.22.0.1"}