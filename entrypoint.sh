#!/bin/bash

# Start the cron service
service cron start

# List the next scheduled cron jobs
echo "The updater will run every hour, on the hour."

tail -f /dev/null