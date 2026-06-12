package typeaccount

import "strings"

type TypeAccount string

type Permission string

// estos son los tipos de cuentas
const (
	TypeAccountAhorro    TypeAccount = "Ahorro"
	TypeAccountCorriente TypeAccount = "Corriente"
)

// estos son todos los permisos posibles
const (
	PermissionTypeAccountRead        Permission = "typeaccount:read"
	PermissionTypeAccountList        Permission = "typeaccount:list"
	PermissionTypeAccountCreate      Permission = "typeaccount:create"
	PermissionTypeAccountUpdate      Permission = "typeaccount:update"
	PermissionTypeAccountManageRoles Permission = "typeaccount:manage_roles"
)

// estos son los permisos posibles para cada tipo de cuenta
var role = map[TypeAccount]map[Permission]struct{}{
	TypeAccountAhorro: {
		PermissionTypeAccountRead:        {},
		PermissionTypeAccountList:        {},
		PermissionTypeAccountCreate:      {},
		PermissionTypeAccountUpdate:      {},
		PermissionTypeAccountManageRoles: {},
	},
	TypeAccountCorriente: {
		PermissionTypeAccountRead:        {},
		PermissionTypeAccountList:        {},
		PermissionTypeAccountCreate:      {},
		PermissionTypeAccountUpdate:      {},
		PermissionTypeAccountManageRoles: {},
	},
}

// AllowedPermissions devuelve la lista de permisos asignados al rol.
func (t TypeAccount) AllowedPermissions() []Permission {
	perms, ok := role[t]
	if !ok {
		return nil
	}
	allowed := make([]Permission, 0, len(perms))
	for p := range perms {
		allowed = append(allowed, p)
	}
	return allowed
}

// Can devuelve true si este rol incluye el permiso solicitado.
func (t TypeAccount) can(p Permission) bool {
	perms, ok := role[t]
	if !ok {
		return false
	}
	_, allowed := perms[p]
	return allowed
}

func parse(value string) (TypeAccount, bool) {
	normalized := TypeAccount(strings.ToLower(strings.TrimSpace(value)))
	switch normalized {
	case TypeAccountAhorro:
		return TypeAccountAhorro, true
	case TypeAccountCorriente:
		return TypeAccountCorriente, true
	default:
		return TypeAccountCorriente, false
	}
}

func (t TypeAccount) Can(permission Permission) bool {
	perms, ok := role[t]
	if !ok {
		return false
	}
	_, allowed := perms[permission]
	return allowed
}
