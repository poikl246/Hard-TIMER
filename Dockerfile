FROM golang:1.16

WORKDIR /app

COPY run.sh /app
COPY main.go /app
RUN chmod +x run.sh
CMD ["./run.sh"]
