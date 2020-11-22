build:
	go build -o bin/miniloops-operator

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object paths="./apis/..."
	$(CONTROLLER_GEN) crd:crdVersions=v1 webhook rbac:roleName=manager-role paths="./apis/..." output:crd:artifacts:config=config/crd/bases

# Install CRDs into a cluster
install: generate
	kubectl apply -f config/crd/*

# Uninstall CRDs into a cluster
uninstall: manifests
	kustomize build config/crd | kubectl delete -f -

# Generate clients
client: controller-gen
	go get k8s.io/code-generator/cmd/client-gen@v0.18.8

	client-gen \
		-v 4 \
		--input-base "github.com/clever-telemetry/miniloops/apis" \
		--input "loops/v1" \
        --clientset-name "loops" \
		--output-file-base "loops" \
		--output-package "github.com/clever-telemetry/miniloops/client" \
		--fake-clientset=false

docker-build:
	docker build -t miniloops .

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.0 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(shell which controller-gen)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif
