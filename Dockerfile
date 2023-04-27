# Use the official Golang image as the base image
FROM golang:1.17

# Set the working directory to /app
WORKDIR /app

COPY go.sum .
# Copy the current directory contents into the container at /app
COPY . .

# Install the ascp command
RUN apt-get update && apt-get install -y ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb


# Build the Go application
RUN go mod download
RUN go mod verify
RUN go build -o main .

# Set environment variables from .env file
ENV $(cat .env | xargs)



# Expose port 8080 for the Echo service to listen on
EXPOSE 8080

# Set the command to run the Go application
CMD ["./main"]

# # STAGE 1
# # use golang image as builder
# FROM golang:1.18.1 AS builder
# # FROM apline:latest 

# # golang specific variables
# ENV GO111MODULE=on \
#   CGO_ENABLED=0 \
#   GOOS=linux \
#   GOARCH=amd64

# WORKDIR /app

# COPY go.mod .
# COPY go.sum .


# RUN go mod download


# RUN apt-get update && apt-get install -y bash
# # Copy Aspera deb package to the container
# COPY ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb

# RUN apt-get update && \
#     apt-get install ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb && \
#     rm -rf ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb && \
#     apt-get clean

# # RUN apt-get update && apt-get install -y  aspera-connect

# # # Copy the temp directory and set the ownership and permissions
# # COPY temp/ /app/temp/
# # RUN chown -R root:root /app/temp/ && chmod -R 755 /app/temp/


# COPY . .


# COPY .env .


# # Build the binary
# RUN go build -o main .


# # # Stage 2: Production stage
# FROM alpine:latest


# ENV GOPATH /go
# ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
# ENV AWS_ACCESS_KEY_ID="ASIAYBJV2LLDYTERBPQJ"
# ENV AWS_SECRET_ACCESS_KEY="Dipu7In2kaXnpvep2SpOCcv8XrnyFQPtvlzdZabJ"
# ENV AWS_SESSION_TOKEN="IQoJb3JpZ2luX2VjEPH//////////wEaCXVzLWVhc3QtMSJGMEQCIESGk5pgC1Z8XVGghuGSr1kfdWR4uO2WvYUier3nUzPBAiBDRk8wUxHeKT9fL6jQ8uS0SaywYHAAOohuEs3AjtTKpiqRAwjq//////////8BEAQaDDU1MjU1MzA0NDY3OSIMpI8qj373oqk6rTizKuUCtjrkTV2CaO/JxQH/0zVoohtjL05Y6vaxguv8OouN7RsOavQerWbz73hNrE1KhWX7ePUnzLKQvXHswLD1M6QSP8XruoQiZYkSI1TGDW+9skqTjj4QWp3wh6Ev/+Db4o0nhtMfjXxSdoeVdEbHPdEwPrhtJV1xF8iZ55kJsyhLkD3IsdWHEhVF010g0MkBf0EatOZ2AjveruSqd2lu1/nyYS0N7ZP4FeSOXaXb2l0AxmM2CkPCqbdLt08cw9a++T0h3KPXYYSGs5Phx3j+EypV/s9iHfmbXvv+PmCo5Y7u0MQRa4U0xTY/6Fgk2PAeHE3sVWmLxRcZadflmjDswPpfTGcssSyjPBAlhLbs57Egi7HjEzREyT2YMJBXE1G6RiRIujmPyEUCSiEjNn7A2LeUSB8veCJ9yVh67CqyzobeFE2gUz3OB32sWgTNiNIVz2YkbA2BYhxwdwX9tPhcuE894LW6+AwgMIiHmaIGOqcBQyg7UVfiQ7wGsk80u6st0o5xo6mMh74fkgiePJWZah4PKdhfI6rbnhK6gUgnjOJT9FhIIjWsd0C0nYKwIalqxHScoZ7sO4T4zUDdb8gMGM5w3wJm/95cpk8Bh3doCjmUIgTTjt0+Gt6aEBF5fMXI8JQssUFX9zSRgC//fkkny9p9PUg9YCpHThUjyCMWdcdD6LA0GB9si5Aa5YeqhPFWdYuE1bGNR94="
# WORKDIR /app

# COPY --from=builder /app/main .


# RUN apk --no-cache add ca-certificates

# RUN mkdir -p /app/temp && \
#     chown -R nobody:nobody /app/temp && \
#     chmod -R 777 /app/temp

# USER nobody:nobody


# EXPOSE 8080

# CMD ["./main"]
# # CMD [ "go","run","main.go" ]

# # STAGE 2
# # use a small and secure base image for the final image
# # FROM alpine:3.14.2
# # WORKDIR /app

# # # Copy the binary from the builder image to the final image
# # COPY --from=builder /app/main .

# # # Run the binary on container startup
# # CMD ["./main"]

# # # STAGE 1
# # # use golang image as builder
# # FROM golang:1.18.1 AS builder

# # # golang specific variables
# # ENV GO111MODULE=on \
# #   CGO_ENABLED=0 \
# #   GOOS=linux \
# #   GOARCH=amd64

# # WORKDIR /app

# # COPY go.mod .
# # COPY go.sum .

# # RUN go mod download

# # # Copy Aspera deb package to the container
# # COPY ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb

# # RUN apt-get update && \
# #     apt-get install -y --allow-downgrades --allow-remove-essential ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb && \
# #     rm -rf ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb && \
# #     apt-get clean


# # COPY . .

# # # Build the binary
# # RUN go build -o main .


# # # STAGE 1
# # # use alpine image as builder
# # FROM golang:alpine AS builder

# # # golang specific variables
# # ENV GO111MODULE=on \
# #   CGO_ENABLED=0 \
# #   GOOS=linux \
# #   GOARCH=amd64


# # WORKDIR /app

# # COPY go.mod .
# # COPY go.sum .

# # RUN go mod download


# # COPY ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb
# # RUN dpkg -i ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb && \
# #     rm -r ./ibm-aspera-desktopclient-3.9.1.168302-linux-64.deb && apt-get purge -y wget

# # COPY . .
# # RUN go build -o main .

# # FROM scratch

# # ARG AWS_ACCESS_KEY_ID
# # ARG AWS_SECRET_ACCESS_KEY
# # ARG AWS_SESSION_KEY

# # #environment variables for the application

# # ENV AWS_ACCESS_KEY_ID={AWS_ACCESS_KEY_ID}
# # ENV AWS_SECRET_ACCESS_KEY={AWS_SECRET_ACCESS_KEY}
# # ENV AWS_SESSION_KEY={AWS_SESSION_KEY}


# # # copy from stage-1 image
# # COPY --from=builder /build/main /

# # # expose the port to run the application on
# # EXPOSE 8080

# # # Command to run
# # ENTRYPOINT ["/main"]



# # # Use an official Golang runtime as a parent image
# # FROM golang:1.17-alpine AS builder

# # # Set the working directory
# # WORKDIR /app

# # # Copy the source code into the container
# # COPY . .

# # # Download the necessary dependencies
# # RUN go mod download

# # # Build the application
# # RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# # # Use an official Alpine Linux runtime as a parent image
# # FROM alpine:latest

# # # Install the Aspera client tools
# # RUN apk add --no-cache openssh-client aspera-cli

# # # Set the working directory
# # WORKDIR /app

# # # Copy the compiled application from the builder image
# # COPY --from=builder /app/app .

# # # Expose the default port (optional)
# # EXPOSE 8080

# # # Run the application
# # CMD ["../app"]

# # # Copy the binary from the builder image to the final image
# # COPY --from=builder /app/main .

# # # Run the binary on container startup
# # CMD ["./main"]

