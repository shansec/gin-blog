package model

import "myginblog/utils/errMsg"

type Profile struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(20)" json:"name"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	QQChat    string `gorm:"type:varchar(200)" json:"qqChat"`
	Wechat    string `gorm:"type:varchar(100)" json:"wechat"`
	Weibo     string `gorm:"type:varchar(200)" json:"weibo"`
	Bili      string `gorm:"type:varchar(200)" json:"bili"`
	Email     string `gorm:"type:varchar(200)" json:"email"`
	Img       string `gorm:"type:varchar(200)" json:"img"`
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)" json:"icpRecord"`
}

// GetProfileInfo 获取个人信息设置
func GetProfileInfo(id int) (Profile, int) {
	var profile Profile
	err = db.Where("ID = ?", id).Find(&profile).Error
	if err != nil {
		return profile, errMsg.ERROR
	}
	return profile, errMsg.SUCCESS
}

// UpdateProfileInfo 更新个人设置
func UpdateProfileInfo(id int, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}
