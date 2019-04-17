# Golang CRUD REST API App with JWT Token authentication

### About Application
- App is based on **Go 1.12**
- DB in develop branch is **MySQL**
- In order to test application with **PostgreSQL** db please check **MNT-1_REST_API_on_PostgreSQL_db** branch
- Main application functionality is available in **Docker**
- For development purposes **[REST API Tutorial](https://github.com/gin-gonic/gin)** HTTP Golang framework is available

# [REST API Tutorial](https://github.com/Maksim1990/Golang_REST_API__JWT_App/blob/develop/API_GUIDE.md)

### How To Run

1) Build and start Docker containers for start application

```
docker-compose build && docker-compose up -d
```

2) Enter in running application container

```
docker exec -it golang_app bash 
```

3) Navigate to app folder and run app with GIN framework command
```
cd src/github.com/goRESTapi && gin -i -all rin main.go
```

4) Now your application is running on **9090** port [http://localhost:9090/](http://localhost:9090/)

5) You can use [Postman](https://www.getpostman.com/) or any other tool for testing REST API app
