package sdk

// change_set.go
type changeSetItem struct {
  operation operation
  change Change
}

type ChangeSet []*changeSetItem
