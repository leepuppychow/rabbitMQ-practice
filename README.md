## Practice with RabbitMQ in Go 

* [Docs](https://www.rabbitmq.com/getstarted.html)

## Notes:

1. Work queues (publisher sends tasks with multiple workers/consumers waiting to get tasks popped off the queue)

  * Install RabbitMW
  * Start RabbitMQ locally: `rabbitmq-server`
  * In one shell publish a message with: `go run work-queue/publisher/send.go some message...`
  * In as many shells you want, run this to start a worker to listen to queue: `go run work-queue/consumer/worker.go`

2. Publish/Subscribe (publisher sends same message to many subscribers at the same time)

  * In one shell publish a message with: `go run pub-sub/publisher/send.go some message...`
  * In as many shells you want, run this to start a worker to listen to queue: `go run pub-sub/subscriber/subscriber.go`
  * To see current exchanges: `rabbitmqctl list_exchanges`
  * To see current bindings: `rabbitmqctl list_bindings`
