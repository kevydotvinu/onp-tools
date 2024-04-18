# onp-tools

A container image with various network tools pre-installed.

## Create Pod and Service
```bash
apiVersion: v1
kind: Pod
metadata:
  name: onp-tools
  labels:
    app: onp
spec:
  containers:
  - image: quay.io/onp/onp-tools
    imagePullPolicy: Always
    name: onp-tools
    ports:
    - containerPort: 9001
    securityContext:
      runAsNonRoot: true
      allowPrivilegeEscalation: false
      readOnlyRootFilesystem: true
      seccompProfile:
        type: RuntimeDefault
      capabilities:
        drop:
        - ALL
---
apiVersion: v1
kind: Service
metadata:
  name: onp-tools
spec:
  ports:
  - port: 9001
    targetPort: 9001
    protocol: TCP
  selector:
    app: onp
```

## Create a container using `podman`
```bash
sudo podman run --tty --interactive --privileged --net host quay.io/onp/onp-tools
```
