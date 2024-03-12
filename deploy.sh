#!/bin/bash

echo "Updating code from Git..."
cd /workspace/duval && sudo git reset --hard origin/master && sudo git fetch origin

echo "Building the application ... "
cd /workspace/duval/cmd/app && go build -buildvcs=false

echo "Reloading systemd..."
sudo systemctl daemon-reload

echo "Restarting the duval service..."
sudo systemctl restart duval

echo "Deployment completed successfully."