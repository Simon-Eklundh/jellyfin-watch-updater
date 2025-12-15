#!/bin/sh

# Export the three Jellyfin vars so they are guaranteed to be in the
# environment of the cron daemon (PID 1).
export JELLYFIN_USER_ID="${JELLYFIN_USER_ID}"
export JELLYFIN_API_KEY="${JELLYFIN_API_KEY}"
export JELLYFIN_BASE_URL="${JELLYFIN_BASE_URL}"

# Optional sanity‑check – abort if any is missing
: "${JELLYFIN_USER_ID:?missing}"
: "${JELLYFIN_API_KEY:?missing}"
: "${JELLYFIN_BASE_URL:?missing}"

printenv > /etc/environment
echo "✅ Env vars exported – starting cron"
exec cron -f