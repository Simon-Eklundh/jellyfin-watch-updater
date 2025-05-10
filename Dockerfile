FROM golang:latest
WORKDIR /app  
COPY ./ .  
RUN go build -o script main.go  

# Install cron  
RUN apt-get update && apt-get install -y cron  

# Copy crontab file to cron.d directory  
COPY crontab /etc/cron.d/crontab  
RUN chmod 0644 /etc/cron.d/crontab && crontab /etc/cron.d/crontab  

CMD ["cron", "-f"]
