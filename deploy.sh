#!/bin/bash

echo "Updating code from Git..."
cd /workspace/duval && sudo git reset --hard origin/master && sudo git fetch origin

echo "Building the application ... "
cd /workspace/duval/cmd/app && go build go build -buildvcs=false

# Check if the build was successful
# shellcheck disable=SC2181
if [ $? -ne 0 ]; then
    echo "Error: Failed to build the application."
    exit 1
fi

echo "Reloading systemd ..."
echo "Reloading systemd..."
sudo systemctl daemon-reload


echo "Restarting the duval service..."
sudo systemctl restart duval

echo "Deployment completed successfully."