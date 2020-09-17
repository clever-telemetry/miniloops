package controllers

import (
	"time"

	ctrl "sigs.k8s.io/controller-runtime"
)

func NoRequeue() ctrl.Result {
	return ctrl.Result{}
}

func Requeue() ctrl.Result {
	return ctrl.Result{Requeue: true}
}

func RequeueAfter(d time.Duration) ctrl.Result {
	return ctrl.Result{RequeueAfter: d}
}
