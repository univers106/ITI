FROM alpine:3.23.3

RUN ["apk", "update"]
RUN ["apk", "upgrade"]
RUN ["apk", "add", "--no-cache", "bash"]

RUN ["apk", "add", "go=1.25.7-r0"]
RUN ["apk", "add", "npm"]

RUN ["mkdir", "/src_backend/"]
RUN ["mkdir", "/src_frontend/"]
RUN ["mkdir", "/app/"]

WORKDIR /src_frontend
COPY frontend .
RUN ["npm", "install"]
RUN ["npm", "build"]

WORKDIR /src_backend

COPY backend .

RUN ["go", "build",  "-o", "/app/main", "main.go"]

EXPOSE 8080

WORKDIR /app

CMD ["./main"]
