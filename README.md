## demo-app
The app is built as a project to learn new microservice intercommunication techniques and new storage solutions.

Features:
* Event Sourcing and Command Query Responsibility Segregation (CQRS) using Apache Kafka
* Elasticsearch / LogStash / Kibana logging solution (ELK Stack)


### Event Sourcing and CQRS
Event Sourcing means keeping your data as a series of immutable events in an event log such as Apache Kafka.
CQRS just means that we separate the Read and Write part of the system. In this demo app the Events microservice
doesn't own the data; it's just a materialized view of the events (as in party, concert etc.) in the event log
(the other type of event).

Advantages:
* Maximum flexibility for analysis of business data and error scenarios since all data is kept
* Data integration is easy
* Loose coupling between services
* You can optimize Read and Write/Update/Delete separately.

Links:
* [Stream processing, Event sourcing, Reactive, CEP… and making sense of it al](https://www.confluent.io/blog/making-sense-of-stream-processing/)


### Apache Kafka
Apache Kafka is a distributed, highly available, fault-tolerant event log. Setting this up is out of scope of this simple demo.
To test it, set up a Kafka instance running on localhost:9092. I used [wurstmeister/kafka-docker](https://github.com/wurstmeister/kafka-docker) for development.

For the project to work in production, remember to set the retention time for your topics to a reasonable time. For data like users you might want to disable this entirely, i.e. keep the user data forever.

### ELK Stack
Todo


### Development
Run once to test:
`$ make testd`

To run tests multiple times without building and removing containers, run:
```
$ make up
$ make test
...
$ make down
```


### Security
I used the following command to generate a self-signed certificate for development:
`go run /usr/local/go/src/crypto/tls/generate_cert.go --host=localhost`

In production you need to replace the files `cert.pem` and `key.pem` with certificates issued from a CA (Certificate Authority).


### Todo
*