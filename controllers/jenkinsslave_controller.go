/*


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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	webappv1 "jenkins-slave/api/v1"
)

// JenkinsslaveReconciler reconciles a Jenkinsslave object
type JenkinsslaveReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webapp.jerryyu.org,resources=jenkinsslaves,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webapp.jerryyu.org,resources=jenkinsslaves/status,verbs=get;update;patch

func (r *JenkinsslaveReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	
	ctx := context.Background()
    _ = r.Log.WithValues("jenkinsslave", req.NamespacedName)

  // Getting current CR，then print out
    obj := &webappv1.Guestbook{}
    if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
        log.Println(err, "Unable to fetch object")
    } else {
		log.Println("Starting to register to Jenkins Master:", obj.Spec.JenkinsMasterURL)
		log.Println("Jenkins slave count:", obj.Spec.Replicas)
		// TBD: start the container and register to jenkins master
    }

  // Initialize CR Status to Running
    obj.Status.Status = "Running"
    if err := r.Status().Update(ctx, obj); err != nil {
        log.Println(err, "unable to update status")
    }

	return ctrl.Result{}, nil
}

func (r *JenkinsslaveReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Jenkinsslave{}).
		Complete(r)
}
