FROM golang:latest

ADD . /go/src/github.com/goRESTapi


RUN apt-get update
RUN apt-get install vim -y
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/joho/godotenv"
RUN go get "github.com/dgrijalva/jwt-go"
RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/codegangsta/gin"
RUN go get "github.com/gorilla/context"
RUN go get "golang.org/x/crypto/bcrypt"
RUN go get "github.com/lib/pq"

# Install api binary globally within container
#RUN go install github.com/goRESTapi

# Set entrypoint
RUN cd /go/src/github.com/goRESTapi

# Expose default port (9090)
EXPOSE 9090

