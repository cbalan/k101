#!/bin/bash -ex
docker run --rm -it \
  --net host \
  -v $(dirname $(realpath $0))/kind-kubeconfig.yaml:/root/.kube/config \
  -v $(pwd):/work \
  -w /work \
  k8s.gcr.io/hyperkube:v1.17.3 kubectl $@