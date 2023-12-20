# Gunakan golang:alpine3.19 sebagai base image
FROM golang:alpine3.19

# Set working directory di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum untuk mengelola dependencies
COPY go.mod go.sum ./

# Download dependencies menggunakan go mod
RUN go mod download

# Salin seluruh kode aplikasi ke dalam container
COPY . .

# Pindah ke direktori cmd/app (sesuai dengan struktur proyek)
WORKDIR /app/cmd/

# Build aplikasi Go
RUN go build -o main

# Port yang akan digunakan oleh aplikasi
EXPOSE 8080

# Perintah untuk menjalankan aplikasi ketika container dijalankan
CMD ["./main"]
