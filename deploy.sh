#!/bin/bash

#sudo apt update
#sudo apt install -y golang
#go version
#sudo apt install git
#sudo git clone https://github.com/cend-org/duval.git
#sudo nano /etc/systemd/system/cend.service
#sudo systemctl daemon-reload
#sudo systemctl enable cend.service
#sudo systemctl start cend.service
#sudo systemctl restart cend.service
#sudo systemctl status cend.service
#to reload the service
#sudo systemctl daemon-reload
#https://medium.com/@olayinkancs/deploying-golang-application-to-aws-ec2-instance-d26891d25b2e

echo "Updating code from Git..."
cd /usr/workspace/duval && sudo git fetch origin && sudo git reset --hard origin/master

echo "building..."
cd /usr/workspace/duval && go build -buildvcs=false

echo "Reloading systemd..."
sudo systemctl daemon-reload

echo "Restarting the duval service..."
sudo systemctl restart cend

echo "Deployment completed successfully."