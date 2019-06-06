#!/bin/bash

########################
# include the magic
########################
. demo-magic.sh

DEMO_PROMPT="${GREEN}➜ ${CYAN}\W "

# hide the evidence
clear

# Put your stuff here
export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig

#pe "cd ~/go/src/k8s.io/zoo"
#pe "ls"

pe "cd ~/go/src/k8s.io/zoo/config/samples"

pe "vi panda_qiqi.yaml"

pe "kubectl create -f panda_qiqi.yaml"

pe "kubectl describe panda qiqi"
