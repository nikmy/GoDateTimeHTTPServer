FROM golang
ADD . .
EXPOSE 80
CMD ["go", "run", "main.go"]