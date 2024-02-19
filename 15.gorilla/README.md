# Gorilla (WebSocket)
This project shows how to use Gorilla to build a send-receive model with WebSocket connection.

# Steps to use
1. Start the server
```go
go run main.go
```
2. Use a WebSocket client([Postman](https://www.postman.com/), [Insomnia](https://insomnia.rest/), CURL, and so on) connect to `localhost:888`
3. Send a message from the client, and you'll got a receive message like:

    Send: `123`
    
    Response: `We've got your message: 123`
