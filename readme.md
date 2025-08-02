## Docker
```
docker pull khemmathat141/simple-api
```

## Endpoints

| Endpoint                                    | Description                |
|---------------------------------------------|----------------------------|
| `GET /hello`                                | Get hello response         |
| `POST /send_request`                        | Send a request             |

// Example POST /send_request body
```
{
    "method": "GET",
    "url": "http://localhost:8080/hello",
    "body": "",
    "headers": {
        "X-Test": "hello"
        ....
    }
}
```
