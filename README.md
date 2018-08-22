## demo-app
Demo application built while reading [Cloud Native programming with Golang](https://www.safaribooksonline.com/library/view/cloud-native-programming/).

The app is built as a project to learn new microservice intercommunication techniques and new storage solutions.

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