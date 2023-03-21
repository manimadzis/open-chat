package entities

type PermissionValue uint64
type Permission struct {
	Value       PermissionValue
	Name        string
	Description string
}
