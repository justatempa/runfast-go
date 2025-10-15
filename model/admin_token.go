package model

import (
	"github.com/justatempa/runfast-go/pkg/database"
	"time"
)

// AdminToken 管理员Token模型
type AdminToken struct {
	Model

	Token     string     `json:"token" gorm:"type:varchar(255);not null;unique"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Role      string     `json:"role" gorm:"type:varchar(255);not null"` // admin 或 user
	ExpiredAt *time.Time `json:"expired_at" gorm:"type:datetime"`
}

func CreateAdminToken(adminToken AdminToken) (int, error) {
	result := database.GetDB().Model(&adminToken).Create(&adminToken).Error
	return adminToken.ID, result
}

// GetAdminTokenByID 根据ID获取AdminToken
func GetAdminTokenByID(id uint) (*AdminToken, error) {
	var adminToken AdminToken
	err := database.GetDB().First(&adminToken, id).Error
	if err != nil {
		return nil, err
	}
	return &adminToken, nil
}

// GetAdminTokenByToken 根据Token获取AdminToken
func GetAdminTokenByToken(token string) (*AdminToken, error) {
	var adminToken AdminToken
	err := database.GetDB().Where("`token` = ?", token).First(&adminToken).Error
	if err != nil {
		return nil, err
	}
	return &adminToken, nil
}

// UpdateAdminToken 更新AdminToken
func UpdateAdminToken(id int, info map[string]any) error {
	return database.GetDB().Model(&AdminToken{}).Where("id = ?", id).Updates(&info).Error
}

// DeleteAdminToken 删除AdminToken
func DeleteAdminToken(id uint) error {
	return database.GetDB().Delete(&AdminToken{}, id).Error
}

// ListAdminTokens 获取AdminToken列表
func ListAdminTokens(limit, offset int) ([]*AdminToken, error) {
	var adminTokens []*AdminToken
	err := database.GetDB().Limit(limit).Offset(offset).Find(&adminTokens).Error
	if err != nil {
		return nil, err
	}
	return adminTokens, nil
}
