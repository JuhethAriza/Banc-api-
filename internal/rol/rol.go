package rol

import "strings"

// Role representa un tipo de usuario dentro del sistema.
// Se usa para asignar permisos y controlar qué operaciones puede hacer cada usuario.
type Role string

// Permission define una acción específica que puede ser permitida o denegada.
type Permission string

const (
	RoleUser    Role = "user"
	RoleManager Role = "manager"
	RoleAdmin   Role = "admin"
	RoleGuest   Role = "guest"
)

const (
	PermissionUserRead        Permission = "user:read"
	PermissionUserList        Permission = "user:list"
	PermissionUserCreate      Permission = "user:create"
	PermissionUserUpdate      Permission = "user:update"
	PermissionUserDelete      Permission = "user:delete"
	PermissionUserManageRoles Permission = "user:manage_roles"
)

var rolePermissions = map[Role]map[Permission]struct{}{
	RoleAdmin: {
		PermissionUserRead:        {},
		PermissionUserList:        {},
		PermissionUserCreate:      {},
		PermissionUserUpdate:      {},
		PermissionUserDelete:      {},
		PermissionUserManageRoles: {},
	},
	RoleManager: {
		PermissionUserRead:   {},
		PermissionUserList:   {},
		PermissionUserCreate: {},
		PermissionUserUpdate: {},
	},
	RoleUser: {
		PermissionUserRead:   {},
		PermissionUserList:   {},
		PermissionUserCreate: {},
	},
	RoleGuest: {
		PermissionUserList: {},
	},
}

// Can devuelve true si este rol incluye el permiso solicitado.
func (r Role) Can(permission Permission) bool {
	perms, ok := rolePermissions[r]
	if !ok {
		return false
	}
	_, allowed := perms[permission]
	return allowed
}

// String convierte el rol a su valor de texto.
func (r Role) String() string {
	return string(r)
}

// ParseRole normaliza un valor de texto a un rol válido.
// Devuelve el rol y un booleano indicando si el valor era válido.
func ParseRole(value string) (Role, bool) {
	normalized := Role(strings.ToLower(strings.TrimSpace(value)))
	switch normalized {
	case RoleAdmin:
		return RoleAdmin, true
	case RoleManager:
		return RoleManager, true
	case RoleUser:
		return RoleUser, true
	case RoleGuest:
		return RoleGuest, true
	default:
		return RoleGuest, false
	}
}

// AllowedPermissions devuelve la lista de permisos asignados al rol.
func (r Role) AllowedPermissions() []Permission {
	perms, ok := rolePermissions[r]
	if !ok {
		return nil
	}

	list := make([]Permission, 0, len(perms))
	for p := range perms {
		list = append(list, p)
	}
	return list
}
