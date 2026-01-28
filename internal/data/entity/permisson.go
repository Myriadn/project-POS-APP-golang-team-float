package entity

type RolePermisson struct {
	RoleID       uint       `gorm:"column:role_id"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID"`
	PermissionID uint       `gorm:"column:permission_id"`
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID"`
}

type Permission struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Code        string `gorm:"type:text" json:"code"`
	Description string `gorm:"type:text" json:"description"`
}

func (RolePermisson) TableName() string {
	return "role_permissions"
}
func (Permission) TableName() string {
	return "permissions"
}
