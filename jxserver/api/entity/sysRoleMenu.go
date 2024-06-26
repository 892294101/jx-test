// 角色菜单相关模型
// author xiaoRui

package entity

// SysRoleMenu 角色与菜单关系模型
type SysRoleMenu struct {
	RoleId uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"roleId"` // 角色id
	MenuId uint `gorm:"column:menu_id;comment:'菜单id';NOT NULL" json:"menuId"` // 菜单id
}

func (SysRoleMenu) TableName() string {
	return "ss_basicmanage_roles_perms"
}
