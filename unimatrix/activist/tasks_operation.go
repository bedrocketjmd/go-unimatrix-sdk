package activist

import "../../unimatrix"

func NewTasksOperation(realm string) *unimatrix.Operation {
	return unimatrix.NewRealmOperation(realm, "tasks")
}
