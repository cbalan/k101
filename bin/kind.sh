#!/bin/bash -e
docker run --rm -it \
  --net host \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(dirname $(realpath $0))/kind-kubeconfig.yaml:/root/.kube/config \
  -v $(pwd):/work \
  -w /work \
  docker.io/catalinbalan/k101-kind:v0.7.0 $@