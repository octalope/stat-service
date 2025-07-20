# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.24-alpine AS build

RUN adduser -u 1001 -D -H dockeruser

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY ./modules ./modules

# compile application
RUN go build -o /stat-service ./modules

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

COPY --from=build /stat-service /stat-service

COPY --from=build /etc/passwd /etc/passwd

USER 1001

EXPOSE 8080

ENTRYPOINT ["/stat-service"]
