# Miniloops

![Docker image](https://github.com/clever-telemetry/miniloops/workflows/Docker%20image/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/clever-telemetry/miniloops)](https://pkg.go.dev/github.com/clever-telemetry/miniloops)
[![Maintainability](https://api.codeclimate.com/v1/badges/1f065d72fb55874e9d87/maintainability)](https://codeclimate.com/github/clever-telemetry/miniloops/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/1f065d72fb55874e9d87/test_coverage)](https://codeclimate.com/github/clever-telemetry/miniloops/test_coverage)

Miniloops goal is to provide an easy to use time series related tasks workflow.

Features:
- Scheduler
- WarpScript support
- Secrets support (mainly for  read/write tokens)
- Basic Loops metrics
- TODO: States

## Getting started

### Installation

To install the operator, run:

```sh
kubectl apply -f https://raw.githubusercontent.com/clever-telemetry/miniloops/master/config/crd/bases/clever-telemetry.io_loops.yaml
kubectl apply -f https://raw.githubusercontent.com/clever-telemetry/miniloops/master/config/operator/deployment.yml
```

### Usage

Run your first loop !

```sh
kubectl apply ./config/examples
```

Debug your first loop !

```sh
kubectl describe loop -n loops
```

## Loop in depth

A Loop is a script which run at regular intervals, its goal is to compute aggregated data from base data.
As an example, we can use IOT temperature devices, each device write its own temperature on a database.
Query a device data is pretty easy, query all data to compute an average data at each time is pretty time and resource consuming.

Then, each Loop has to query data, aggregate it and write it.

### WarpScript

A WarpScript Loop is fully autonomous, in the WarpScript body you have to query, aggregate and persist the data

Ex:
```yaml
---
apiVersion: clever-telemetry.io/v1
kind: Loop
metadata:
  namespace: loops
  name: test
spec:
  endpoint: https://warp10.gra1.metrics.ovh.net/api/v0/exec
  every: 10s
  script: |
    REV
    [ NEWGTS 'c' RENAME ] 
...
```

### Secrets
> This is a secret between us ;-)

You can inject real [Kubernetes secrets](https://kubernetes.io/fr/docs/concepts/configuration/secret/) into your Loop.

Use this syntax:

```yaml
---
apiVersion: clever-telemetry.io/v1
kind: Loop
metadata:
  namespace: loops
  name: test
spec:
  ...
  imports:
  - secret:
      name: mysecret
  ... 
...
```

`imports` is an array of secrets/configs you can import in your Loop.
Each import must be in the same `namespace` of the Loop.
Each item must have a name.

The Runtime will read all secrets and inject them in your WarpScript, when it do it, it iterate over each Secret keys, and wrote a new variable with the key and the value.
The Secret Namespace/Name is not kept, so, 2 imported secrets with the same key will be override by the last one. 

Then you will be able to use them in your script.

Ex:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
  namespace: loops
type: Opaque
string:
  username: dG90bwo= 
  password: bXlwYXNzd29yZAo=
```
```yaml
apiVersion: clever-telemetry.io/v1
kind: Loop
metadata:
  namespace: loops
  name: test
spec:
  endpoint: https://warp10.gra1.metrics.ovh.net/api/v0/exec
  every: 10s
  script: "[ $readToken '~.*' {} ] FIND"
  imports:
  - secret:
      name: mysecret 
...
```

### States

In some use cases, you can have to reuse variables from a previous execution, that's where states came.
States are a map of keys/values where keys are strings and values are whatever you want.

You can enable States for your loop, with a parameter in the spec:

```yaml
apiVersion: clever-telemetry.io/v1
kind: Loop
metadata:
  namespace: loops
  name: test
spec:
  ...
  useState: true
  ...
```

By enabling this feature, each Loop execution will append a map to your stack.
Then, you need put a map on your stack at the ends of your loops

```warpscript
// { 'count' 1 } // Injected state
'previous' STORE

// TODO: handle the case of first state, where 'count' is not defined
$previous 'count' GET
$count 1 +
{ 'count' $count } // push a map (the next state) at the end of the stack
```



### Runner metrics

Each loop produce serveral metrics, this is the list.
There are exposed with all controllers metrics on this [endpoint](http://127.0.0.1:9100/metrics)

```
miniloops_runner_execution_count{ loop="LOOP_NAME", namespace="LOOP_NAMESPACE" } int
miniloops_runner_execution_duration{ loop="LOOP_NAME", namespace="LOOP_NAMESPACE",unit="ms" } int
miniloops_runner_execution_error_count{ loop="LOOP_NAME", namespace="LOOP_NAMESPACE" } int
```
