FROM golang:1.19-alpine
#Define Work Dir for all image files to be built
WORKDIR /app
#Install Necessary Modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy Source Files to work dir
COPY *.go ./
#
RUN go build -o /ms-template
EXPOSE 8080
CMD [ "/ms-template" ]

