package sdk

// change_set.go
type changeSetItem struct {
  operation Operation
  change Change
}

type ChangeSet []*changeSetItem
