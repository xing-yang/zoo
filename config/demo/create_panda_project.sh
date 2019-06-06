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

pe "cd ~/go/src/k8s.io/zoo"

echo "==> Step 1: Create a new project"
p "kubebuilder init --domain zoo.k8s.io --license apache2 --owner \"The Kubernetes Authors\""

echo "==> Step 2: Define a new API"
p "kubebuilder create api --group animals --version v1alpha1 --kind Panda"
pe "ls"

echo "==> Step 3: Add CRD definitions"
pe "vi ~/go/src/k8s.io/zoo/pkg/apis/animals/v1alpha1/panda_types.go"

echo "==> Step 4: Add controller logic"
pe "vi ~/go/src/k8s.io/zoo/pkg/controller/panda/panda_controller.go"

#pe "make"

#echo "=> CRD manifest generated"
#pe "vi ~/go/src/k8s.io/zoo/config/crds/animals_v1alpha1_panda.yaml"

echo "==> Step 5: Install the CRDs into Kubernetes Cluster"
pe "make install"

#echo "=> CRD manifest generated"
pe "vi ~/go/src/k8s.io/zoo/config/crds/animals_v1alpha1_panda.yaml"

echo "==> Step 6: Build and run the controller manager"
pe "make run"

#echo "==> cd ~/go/src/k8s.io/zoo/config/samples"
#cd ~/go/src/k8s.io/zoo/config/samples

#vi panda_lingling.yaml
#sleep 1

#echo "==> kubectl create -f panda_lingling.yaml"
#kubectl create -f panda_lingling.yaml
#sleep 1

#echo "==> kubectl describe pandas"
#kubectl describe pandas
#sleep 1
