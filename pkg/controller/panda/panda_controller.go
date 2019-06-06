/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package panda

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	animalsv1alpha1 "k8s.io/zoo/pkg/apis/animals/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller")

// Add creates a new Panda Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcilePanda{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("panda-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Panda
	err = c.Watch(&source.Kind{Type: &animalsv1alpha1.Panda{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create
	// Uncomment watch a Deployment created by Panda - change this for objects you create
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &animalsv1alpha1.Panda{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcilePanda{}

// ReconcilePanda reconciles a Panda object
type ReconcilePanda struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Panda object and makes changes based on the state read
// and what is in the Panda.Spec
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=animals.zoo.k8s.io,resources=pandas,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=animals.zoo.k8s.io,resources=pandas/status,verbs=get;update;patch
func (r *ReconcilePanda) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the Panda instance
	log.Info("Processing panda instances ......")
	instance := &animalsv1alpha1.Panda{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	var foundPanda PandaRecord
	var found bool

	data := getData()

	for !found {
		pandaRec := findPanda(data.PandaRecords, instance.Name)

		if pandaRec != nil {
			foundPanda = *pandaRec
			found = true
		}
	}

	if !found {
		fmt.Println("Invalid Panda Name given: ", instance.Name)
		fmt.Println()
		return reconcile.Result{}, errors.NewBadRequest("Invalid Panda Name given")
	}

	instance.Status.Age = foundPanda.Age
	instance.Status.Weight = foundPanda.Weight
	instance.Status.BambooConsumption = &(foundPanda.BambooConsumption)
	instance.Status.AppetiteScale = getAppetiteScale(foundPanda.AppetiteScale)
	instance.Status.WeightScale = getWeightScale(foundPanda.WeightScale)

	err = r.Update(context.TODO(), instance)

	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func getAppetiteScale(appetite string) *animalsv1alpha1.AppetiteScale {
	var result animalsv1alpha1.AppetiteScale
	if string(animalsv1alpha1.AppetiteLow) == appetite {
		result = animalsv1alpha1.AppetiteLow
	} else if string(animalsv1alpha1.AppetiteHigh) == appetite {
		result = animalsv1alpha1.AppetiteHigh
	} else {
		result = animalsv1alpha1.AppetiteNormal
	}
	return &result
}

func getWeightScale(weight string) *animalsv1alpha1.WeightScale {
	var result animalsv1alpha1.WeightScale
	if string(animalsv1alpha1.AppetiteLow) == weight {
		result = animalsv1alpha1.WeightLow
	} else if string(animalsv1alpha1.WeightHigh) == weight {
		result = animalsv1alpha1.WeightHigh
	} else {
		result = animalsv1alpha1.WeightNormal
	}
	return &result
}

func getData() Results {
	file, _ := ioutil.ReadFile("config/samples/panda_records.json")

	data := Results{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.PandaRecords); i++ {
		/*
			fmt.Println("Name: ", data.PandaRecords[i].Name)
			fmt.Println("Birth Year: ", data.PandaRecords[i].BirthYear)
			fmt.Println("Birth Month: ", data.PandaRecords[i].BirthMonth)
			fmt.Println("Birth Day: ", data.PandaRecords[i].BirthDay)
			fmt.Println("Birth Place: ", data.PandaRecords[i].BirthPlace)
			fmt.Println("Mom Name: ", data.PandaRecords[i].MomName)
			fmt.Println("Birth Weight: ", data.PandaRecords[i].BirthWeight)
			fmt.Println("Age: ", data.PandaRecords[i].Age)
			fmt.Println("Weight: ", data.PandaRecords[i].Weight)
			fmt.Println("Bamboo Consumption: ", data.PandaRecords[i].BambooConsumption)
			fmt.Println("AppetiteScale: ", data.PandaRecords[i].AppetiteScale)
			fmt.Println("WeightScale: ", data.PandaRecords[i].WeightScale)
		*/
	}

	return data
}

func findPanda(pandaRecords []PandaRecord, pandaname string) *PandaRecord {
	for _, pandaRecord := range pandaRecords {
		if pandaRecord.Name == pandaname {
			return &pandaRecord
		}
	}
	return nil
}

type Results struct {
	PandaRecords []PandaRecord `json:"pandaRecords"`
}

type PandaRecord struct {
	Name              string `json:"name"`
	BirthYear         int32  `json:"birthYear"`
	BirthMonth        int32  `json:"birthMonth"`
	BirthDay          int32  `json:"birthDay"`
	BirthPlace        string `json:"birthPlace"`
	MomName           string `json:"momName"`
	BirthWeight       int32  `json:"birthWeight"`
	Age               int32  `json:"age"`
	Weight            int32  `json:"weight"`
	BambooConsumption int32  `json:"bambooConsumption"`
	AppetiteScale     string `json:"appetiteScale"`
	WeightScale       string `json:"weightScale"`
}
