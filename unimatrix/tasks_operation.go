package unimatrix

func NewTasksOperation(realm string) *Operation {
	return NewRealmOperation(realm, "tasks")
}
