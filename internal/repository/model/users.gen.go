// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	UserID    int32  `gorm:"column:user_id" json:"user_id"`
	Account   string `gorm:"column:account" json:"account"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	Email     string `gorm:"column:email" json:"email"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
