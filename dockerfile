FROM go:1.20.2

WORKDIR /app

EXPOSE 8080

CMD ["go", "run", "main.go"]