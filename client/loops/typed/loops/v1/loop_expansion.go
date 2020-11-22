package v1

import (
	"context"
	"encoding/json"
	"time"

	api "github.com/clever-telemetry/miniloops/apis/loops/v1"
	"github.com/pkg/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/jsonmergepatch"
)

type LoopExpansion interface {
	PatchSpec(ctx context.Context, original, modified *api.Loop) error
	PatchStatus(ctx context.Context, original, modified *api.Loop) error
	PatchMeta(ctx context.Context, original, modified *api.Loop) error
	// Register the last execution
	SetLastExecution(ctx context.Context, name string, executionTime time.Time, isSuccess bool) error
}

// Patch the loop spec only (client side usage)
// original and modified specs will be updated
func (client *loops) PatchSpec(ctx context.Context, original, modified *api.Loop) error {
	current, err := client.Get(ctx, original.GetName(), meta.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "cannot get current loop version")
	}

	ori := api.Loop{Spec: *original.Spec.DeepCopy()}
	mod := api.Loop{Spec: *modified.Spec.DeepCopy()}
	cur := api.Loop{Spec: current.Spec}

	bOri, err := json.Marshal(ori)
	if err != nil {
		return errors.Wrap(err, "cannot marshal origin Loop")
	}

	bMod, err := json.Marshal(mod)
	if err != nil {
		return errors.Wrap(err, "cannot marshal modified Loop")
	}

	bCur, err := json.Marshal(cur)
	if err != nil {
		return errors.Wrap(err, "cannot marshal current Loop")
	}

	patch, err := jsonmergepatch.CreateThreeWayJSONMergePatch(bOri, bMod, bCur)
	if err != nil {
		return errors.Wrap(err, "cannot create merge patch")
	}

	if len(patch) == 0 || string(patch) == "{}" {
		return nil
	}

	next, err := client.Patch(ctx, modified.GetName(), types.MergePatchType, patch, meta.PatchOptions{})
	if err != nil {
		return errors.Wrap(err, "cannot patch Loop")
	}

	original.Spec = *next.Spec.DeepCopy()
	modified.Spec = *next.Spec.DeepCopy()

	return nil
}

// Patch the loop status only (operator side usage)
// original and modified specs will be updated
func (client *loops) PatchStatus(ctx context.Context, original, modified *api.Loop) error {
	current, err := client.Get(ctx, original.GetName(), meta.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "cannot get current loop version")
	}

	ori := api.Loop{Status: *original.Status.DeepCopy()}
	mod := api.Loop{Status: *modified.Status.DeepCopy()}
	cur := api.Loop{Status: current.Status}

	bOri, err := json.Marshal(ori)
	if err != nil {
		return errors.Wrap(err, "cannot marshal origin Loop")
	}

	bMod, err := json.Marshal(mod)
	if err != nil {
		return errors.Wrap(err, "cannot marshal modified Loop")
	}

	bCur, err := json.Marshal(cur)
	if err != nil {
		return errors.Wrap(err, "cannot marshal current Loop")
	}

	patch, err := jsonmergepatch.CreateThreeWayJSONMergePatch(bOri, bMod, bCur)
	if err != nil {
		return errors.Wrap(err, "cannot create merge patch")
	}

	if len(patch) == 0 || string(patch) == "{}" {
		return nil
	}

	next, err := client.Patch(ctx, original.GetName(), types.MergePatchType, patch, meta.PatchOptions{}, "status")
	if err != nil {
		return errors.Wrap(err, "cannot patch Loop")
	}

	original.Status = *next.Status.DeepCopy()
	modified.Status = *next.Status.DeepCopy()

	return nil
}

// Patch the loop spec only (operator/client side usage)
// original and modified specs will be updated
func (client *loops) PatchMeta(ctx context.Context, original, modified *api.Loop) error {
	current, err := client.Get(ctx, original.GetName(), meta.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "cannot get current loop version")
	}

	ori := api.Loop{ObjectMeta: *original.ObjectMeta.DeepCopy()}
	mod := api.Loop{ObjectMeta: *modified.ObjectMeta.DeepCopy()}
	cur := api.Loop{ObjectMeta: current.ObjectMeta}

	bOri, err := json.Marshal(ori)
	if err != nil {
		return errors.Wrap(err, "cannot marshal origin Loop")
	}

	bMod, err := json.Marshal(mod)
	if err != nil {
		return errors.Wrap(err, "cannot marshal modified Loop")
	}

	bCur, err := json.Marshal(cur)
	if err != nil {
		return errors.Wrap(err, "cannot marshal current Loop")
	}

	patch, err := jsonmergepatch.CreateThreeWayJSONMergePatch(bOri, bMod, bCur)
	if err != nil {
		return errors.Wrap(err, "cannot create merge patch")
	}

	if len(patch) == 0 || string(patch) == "{}" {
		return nil
	}

	next, err := client.Patch(ctx, modified.GetName(), types.MergePatchType, patch, meta.PatchOptions{})
	if err != nil {
		return errors.Wrap(err, "cannot patch Loop")
	}

	original.ObjectMeta = *next.ObjectMeta.DeepCopy()
	modified.ObjectMeta = *next.ObjectMeta.DeepCopy()

	return nil
}

// Register an execution
func (client *loops) SetLastExecution(ctx context.Context, name string, executionTime time.Time, isSuccess bool) error {
	t := meta.NewTime(executionTime)

	ori, err := client.Get(ctx, name, meta.GetOptions{})
	if err != nil {
		return err
	}

	mod := ori.DeepCopy()
	mod.Status.LastExecution = &t
	if isSuccess {
		mod.Status.LastExecutionSuccess = &t
	}

	return client.PatchStatus(ctx, ori, mod)
}
