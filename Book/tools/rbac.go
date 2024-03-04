package tools

import (
	"github.com/zpatrick/rbac"
)

var roles []rbac.Role

func InitRoles() {
	roles = []rbac.Role{
		{
			RoleID: "root",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("watch", "*"),
			},
		},
		{
			RoleID: "user",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("watch", "/regis"),
				rbac.NewGlobPermission("watch", "/loginByName"),
				rbac.NewGlobPermission("watch", "/loginByEmail"),
				rbac.NewGlobPermission("watch", "/loginByTel"),
				rbac.NewGlobPermission("watch", "/loginByNameAndPRC"),
				rbac.NewGlobPermission("watch", "/captcha"),
				rbac.NewGlobPermission("watch", "/captcha/verify"),
				rbac.NewGlobPermission("watch", "/UserInfo"),
				rbac.NewGlobPermission("watch", "/MyRecord"),
				rbac.NewGlobPermission("watch", "/RootRecord"),
				rbac.NewGlobPermission("watch", "/ReturningBook"),
				rbac.NewGlobPermission("watch", "/BorrowBook"),
				rbac.NewGlobPermission("watch", "/BuyBook"),
			},
		},
	}
}
func FindRole(roleid string) (myrole rbac.Role) {
	for _, role := range roles {
		if role.RoleID == roleid {
			myrole = role
			return
		}
	}
	return rbac.Role{}
}

//for _, role := range roles {
//fmt.Println("Role:", role.RoleID)
//for _, rating := range []string{"g", "pg-13", "r"} {
//canWatch, _ := role.Can("watch", rating)
//fmt.Printf("Can watch %s? %t\n", rating, canWatch)
//}
//}
