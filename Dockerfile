FROM golang:1.22

#set working directory
WORKDIR /Projetos/src/app

#Copy the source code
COPY . .

#Expose the port
EXPOSE 8000

#Build the Go app
RUN go build -o main cmd/main.go

#Run the executable
CMD ["./main"]
