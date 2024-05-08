#!/bin/bash

echo "Updating code from Git..."
cd /usr/workspace/duval && sudo git fetch origin && sudo git reset --hard origin/master

echo "building..."
cd /usr/workspace/duval/cmd/app && go build -buildvcs=false

echo "Reloading systemd..."
sudo systemctl daemon-reload

echo "Restarting the duval service..."
sudo systemctl restart cend

echo "Deployment completed successfully."