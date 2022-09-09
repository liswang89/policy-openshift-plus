# policy-openshift-plus

This repository is a collecton of policies for OpenShift Plus Workflow#2.

To allow policies to apply automatically, it is using GitOps to deploy policies to cluster.

## Example
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: rbac-app
  namespace: openshift-gitops
spec:
  destination:
    namespace: policies
    server: https://kubernetes.default.svc
  project: default
  source:
    directory:
      recurse: true
    path: Usecase-2
    repoURL: https://github.com/mps-interop/policy-openshift-plus.git
    targetRevision: main
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
```
The policies are applied to all managed clusters that have the local-cluster set to true. If policies need to be applies by another condition, update the PlacementRule.spec.clusterSelector.matchExpressions section in the policies.
