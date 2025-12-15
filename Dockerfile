# ------------------------------------------------------------
# 1️⃣  Build the Go binary
# ------------------------------------------------------------
FROM golang:latest AS builder

WORKDIR /app
COPY ./ .
RUN go build -o script main.go

# ------------------------------------------------------------
# 2️⃣  Runtime image – Debian‑slim (lighter than full golang)
# ------------------------------------------------------------
FROM debian:bookworm-slim

# Install cron and any runtime deps you need
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        cron \
        ca-certificates \
        tzdata && \
    rm -rf /var/lib/apt/lists/*

# Copy the compiled binary from the builder stage
COPY --from=builder /app/script /usr/local/bin/script

COPY crontab /etc/cron.d/crontab
RUN chmod 0644 /etc/cron.d/crontab && \
    crontab /etc/cron.d/crontab

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]