docker-compose build && docker-compose up -d
winpty docker exec -it golang_app bash 
cd src/github.com/goRESTapi && gin -i -all rin main.go

go run main.go


docker-compose down