package driver

import (
	"database/sql"
	"fmt"
)

type role struct {
	name string
	id   int
}

type permission struct {
	name  string
	id    int
	roles []*role
}
type RolePermission struct {
	id            int
	role_id       int
	permission_id int
}

var permissionsList [100]*permission
var rolesList [5]*role

func initRolesList() {
	Admin := newRole("Admin")
	Editor := newRole("Editor")
	Author := newRole("Author")
	Contributor := newRole("Contributor")
	Subscriber := newRole("Subscriber")
	/*rolesList = append(rolesList, Admin)
	rolesList = append(rolesList, Editor)
	rolesList = append(rolesList, Author)
	rolesList = append(rolesList, Contributor)
	rolesList = append(rolesList, Subscriber)*/
	rolesList[0] = Admin
	rolesList[1] = Editor
	rolesList[2] = Author
	rolesList[3] = Contributor
	rolesList[4] = Subscriber
}
func newRole(name string) *role {
	return &role{
		name: name,
		id:   0,
	}
}
func newPermission(name string, roles ...*role) *permission {
	return &permission{
		name:  name,
		id:    0,
		roles: roles,
	}
}
func newRolePermission(role_name string, permission_name string, db *sql.DB) error {
	var roleId int
	row := db.QueryRow(`SELECT id FROM tamiat_roles WHERE name=$1`, role_name)
	_ = row.Scan(&roleId)

	var permissionId int
	row = db.QueryRow(`SELECT id FROM tamiat_permissions WHERE name=$1`, permission_name)
	_ = row.Scan(&permissionId)

	_, err := db.Exec(`INSERT INTO tamiat_role_permission(role_id,permission_id) VALUES ($1,$2)`, roleId, permissionId)
	return err
}
func createRoles(db *sql.DB) {
	initRolesList()
	for i := 0; i < 5; i++ {
		_ = db.QueryRow(`INSERT INTO tamiat_roles(name) VALUES($1) RETURNING id`, rolesList[i].name)
	}
}
func createPermissions(db *sql.DB) {

	CreateUsers := newPermission("create_users", rolesList[0])
	ReadUsers := newPermission("list_users", rolesList[0])
	UpdateUsers := newPermission("update_users", rolesList[0])
	DeleteUsers := newPermission("delete_users", rolesList[0])
	CreateSite := newPermission("create_site", rolesList[0])
	DeleteSite := newPermission("delete_site", rolesList[0])
	EditTheme := newPermission("edit_theme", rolesList[0])
	ModerateComments := newPermission("moderate_comments", rolesList[0], rolesList[1])
	ModerateCategories := newPermission("manage_categories", rolesList[0], rolesList[1])

	permissionsList[0] = CreateUsers
	permissionsList[1] = ReadUsers
	permissionsList[2] = UpdateUsers
	permissionsList[3] = DeleteUsers
	permissionsList[4] = CreateSite
	permissionsList[5] = DeleteSite
	permissionsList[6] = EditTheme
	permissionsList[7] = ModerateComments
	permissionsList[8] = ModerateCategories
	for i := 0; i < len(permissionsList); i++ {
		if permissionsList[i] == nil {
			break
		}
		_ = db.QueryRow(`INSERT INTO tamiat_permissions(name) VALUES($1) RETURNING id`, permissionsList[i].name)

	}
}
func createRolesPermissions(db *sql.DB) {
	// admin has all permissions
	for i := 0; i < len(permissionsList); i++ {
		if permissionsList[i] == nil {
			break
		}
		for j := 0; j < len(permissionsList[i].roles); j++ {
			err := newRolePermission(permissionsList[i].roles[j].name, permissionsList[i].name, db)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
func SeedDB(db *sql.DB) {
	createRoles(db)
	createPermissions(db)
	createRolesPermissions(db)
}
