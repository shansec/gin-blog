package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"myginblog/utils/errMsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20); not null" json:"username" label:"用户名"`
	Password string `gorm:"type:varchar(100); not null" json:"password" label:"密码"`
	// 1表示管理员，2表示阅读者
	Role int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=1" label:"用户权限"`
}

// CheckUser 检查用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errMsg.ERROR_USERNAME_USED //1001
	}
	return errMsg.SUCCESS
}

// ChangeUserPass 修改密码
func ChangeUserPass(id int, data *User) int {
	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errMsg.ERROR_USERNAME_USED //1001
	}
	return errMsg.SUCCESS
}

// CreateUser 添加用户
func CreateUser(data *User) int {
	err := db.Create(data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// GetUser 获取单个用户
func GetUser(id int) (User, int) {
	var user User
	err = db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errMsg.ERROR
	}
	return user, errMsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64, int) {
	var users []User
	var user User
	var total int64
	db.Model(&user).Count(&total)
	if username == "" {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	} else {
		err = db.Where("username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errMsg.ERROR
	}
	return users, total, errMsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// EditUser 编辑用户
func EditUser(id int, data *User) int {
	var user User
	userMap := make(map[string]interface{})
	userMap["username"] = data.Username
	userMap["role"] = data.Role
	err = db.Model(&user).Where(" id = ?", id).Updates(userMap).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// EncryptPwd 密码加密
func EncryptPwd(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

// CheckLogin 登录验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).Find(&user)
	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return errMsg.ERROR_USER_NOT_EXIST
	}
	if passwordErr != nil {
		return errMsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errMsg.ERROR_USER_NO_RIGHT
	}
	return errMsg.SUCCESS
}
