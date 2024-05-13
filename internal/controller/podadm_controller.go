/*
Copyright 2024 Wayne Wang<net_use@bzhy.com>.

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

package controller

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"

	podadmv1beta1 "podadm/api/v1beta1"
)

// PodAdmReconciler reconciles a PodAdm object
type PodAdmReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=podadm.sysadm.cn,resources=podadms,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=podadm.sysadm.cn,resources=podadms/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=podadm.sysadm.cn,resources=podadms/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodAdm object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *PodAdmReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
<<<<<<< HEAD

	logger.Info("process podadm with namespace and object name is ", req.Namespace, req.Name)
=======
>>>>>>> createCRD20240507

	logger.Info("request object information is: ", "object name:", req.Name, "namespace:", req.Namespace, "object string:", req.String())
	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

func (r *PodAdmReconciler) getConfigMapChanged(ctx context.Context, configMap client.Object) []reconcile.Request {
	logger := log.FromContext(ctx)
	logger.Info("changed configMap is:", "configMap name:", configMap.GetName(), "namespaces", configMap.GetNamespace())

	requests := make([]reconcile.Request, 0)

	return requests
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodAdmReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&podadmv1beta1.PodAdm{}).
		Watches(&corev1.ConfigMap{}, handler.EnqueueRequestsFromMapFunc(r.getConfigMapChanged),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Complete(r)
}
