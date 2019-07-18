# Golang CRUD REST API App with JWT Token authentication 

### About Application
- App is based on **Go 1.12**
- DB in develop branch is **MySQL**
- In order to test application with **PostgreSQL** db please check **MNT-1_REST_API_on_PostgreSQL_db** branch
- Main application functionality is available in **Docker**
- For development purposes **[REST API Tutorial](https://github.com/gin-gonic/gin)** HTTP Golang framework is available 

# CHECK ALSO USEFUL [REST API Tutorial](https://github.com/Maksim1990/Golang_REST_API__JWT_App/blob/develop/API_GUIDE.md)

### How To Run

1) Build and start Docker containers for start application

```
docker-compose build && docker-compose up -d
```

2) Enter in running application container

```
docker exec -it golang_app bash 
```

3) Navigate to app folder
```
cd src/github.com/goRESTapi
```
4) Set required credentials in **.env** file:
```
cp .env.dist .env
```
5) Run app with GIN framework command
```
gin -i -all run main.go
```

6) Now your application is running on **9090** port [http://localhost:9090/](http://localhost:9090/)

7) You can use [Postman](https://www.getpostman.com/) or any other tool for testing REST API app
