docker run --name mongo -p 27017:27017 -d mongo:latest
docker run --name genghis --link mongo:db -p 5678:5678 -d dockervan/genghis
