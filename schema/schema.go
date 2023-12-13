package schema

import (
	"github.com/teerapoom/mini_api003/repository"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreditCardID uint       // เปลี่ยนเป็น uint สำหรับระบุ Foreign Key
	CreditCard   CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // เพิ่มการกำหนด constraint เพื่อป้องกันปัญหาแบบเก็บค่า NULL
}

type CreditCard struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Number   string `json:"number"`
}

// CreditCard
func (card *CreditCard) SaveCard() (*CreditCard, error) {
	if err := repository.Db.Create(&card).Error; err != nil {
		return nil, err
	}
	return card, nil
}

// User
func GetUser(CreditCard *[]CreditCard) (err error) {
	err = repository.Db.Find(&CreditCard).Error
	if err != nil {
		return err
	}
	return nil
}
