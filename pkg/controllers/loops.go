package controllers

import (
	"context"

	loops "github.com/clever-telemetry/miniloops/apis/loops/v1"
	"github.com/clever-telemetry/miniloops/pkg/client"
	"github.com/clever-telemetry/miniloops/pkg/runner"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type LoopsController struct {
	runner runner.Runner
}

func Loops(mgr ctrl.Manager) error {
	loopsCtrl := &LoopsController{
		runner: runner.Local(),
	}

	return ctrl.
		NewControllerManagedBy(mgr).
		For(&loops.Loop{}).
		Complete(loopsCtrl)
}

func (lCtrl *LoopsController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()

	logrus.Infof("Reconcile %s", req.NamespacedName)

	loop, err := client.
		LoopsFor(req.Namespace).
		Get(ctx, req.Name, meta.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return NoRequeue(), nil
		}
		return Requeue(), err
	}

	modified := loop.DeepCopy()

	if err := lCtrl.runner.UpsertLoop(modified); err != nil {
		return Requeue(), err
	}

	if !loop.Status.Equal(modified.Status) {
		err := client.LoopsFor(loop.GetNamespace()).PatchStatus(ctx, loop, modified)
		if err != nil {
			return Requeue(), err
		}
	}

	if !loop.EqualMeta(modified.ObjectMeta) {
		err := client.LoopsFor(loop.GetNamespace()).PatchMeta(ctx, loop, modified)
		if err != nil {
			return Requeue(), err
		}
	}

	return NoRequeue(), nil
}
