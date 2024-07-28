# ms-metawebhooks

This repository contains a microservice for handling meta webhooks.

### [🎥 Confira a Playlist do Projeto no YouTube! 🚀](https://www.youtube.com/watch?v=XbZ6neR4oms&list=PLOE-AgqKOXHKJxjDoywG5Y4xjGidITAyp)

### References

[Go (Programming Language)](https://go.dev/)
Go is an open-source programming language developed by Google. It is designed for simplicity, efficiency, and scalability, making it ideal for building reliable and efficient software. Go features strong typing, garbage collection, concurrency support through goroutines, and a rich standard library, making it well-suited for building web services, system tools, and large-scale applications.

[ngrok (Tunneling Software)](https://ngrok.com/docs/what-is-ngrok/)
ngrok is a tool that creates secure tunnels to localhost, allowing you to expose local servers behind NATs and firewalls to the public internet over secure tunnels. This enables you to easily share and test webhooks, APIs, and other HTTP/S services during development without deploying them to a public server.

[RabbitMQ (Message Broker)](https://www.rabbitmq.com/tutorials/tutorial-one-go)
RabbitMQ is a powerful open-source message broker software that implements the Advanced Message Queuing Protocol (AMQP). It acts as a mediator between applications, allowing them to communicate by sending and receiving messages in a decoupled manner. RabbitMQ is highly scalable, supports multiple messaging protocols, and is widely used for building distributed systems and implementing reliable message queuing.

[Facebook Graph API Webhooks](https://developers.facebook.com/docs/graph-api/webhooks/getting-started/)
Facebook Graph API Webhooks allow you to receive real-time updates from Facebook when specific events occur. By setting up webhooks, you can subscribe to changes in user data, page updates, or other events on Facebook, and Facebook will send HTTP POST requests to your server whenever these events occur. This enables developers to build applications that react to changes on Facebook in real time, such as updating content or sending notifications.

### Clone the Repository

```bash
git clone https://github.com/joaogcamilojr/ms-metawebhooks.git
cd ms-metawebhooks
cp .env.example .env
```

### Running Rabbitmmq with Docker

```nash
docker run -d -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password --name rabbitmq-1 rabbitmq:management
```

Feel free to customize this README with additional details. This README will help users understand how to use and set up your microservice and also how to embed YouTube widgets into their projects.
