### Notes


### Dev
Set environment variables in docker-compose.yaml:
 * LOG=DEBUG
 * MONGO=mongodb://127.0.0.1:27017

### Mongo
docker container run -d --name my-mongo -p 27017:27017 -v data:/data/db mongo