Steps before release of the AKS reflector

- [ ] Secret does not exist but cluster exists - Create a new secret
- [ ] Partial apply of clusters - ensure that subsequent reconciliations work
- [ ] Conditions - ready with count of reflected clusters
- [ ] Events - publish when cluster created or removed
- [ ] Manually triggered reconciliation
- [ ] Provide for authentication via a Secret?
- [ ] CommonLabels support
- [ ] ObservedGeneration support!
- [X] Suspension
- [X] Add managed-by labels!
