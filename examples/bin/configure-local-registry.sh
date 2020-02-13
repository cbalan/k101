#!/bin/bash -ex

# Save original docker.json
if [ ! -f /etc/docker/daemon.json.orig ]; then
  cp /etc/docker/daemon.json /etc/docker/daemon.json.orig
fi

# Add my.registry to the list of insecure registries
cat /etc/docker/daemon.json.orig | jq '."insecure-registries" += ["my.registry:31000"]' > /etc/docker/daemon.json

# Restart docker daemon
systemctl restart docker

echo 127.0.0.1 my.registry >> /etc/hosts