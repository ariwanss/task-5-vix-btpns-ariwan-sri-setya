package models

import (
	"time"

	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/database"
	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/helpers"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	Photo     Photo     `json:"photoId" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *User) setUsername(username string) {
	u.Username = username
}

func (u *User) setEmail(email string) {
	u.Email = email
}

func (u *User) setPassword(password string) {
	u.Password = password
}

func (u *User) setPhoto(photo Photo) {
	u.Photo = photo
}

func CreateUser(user *User) (*User, error) {
	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.setPassword(hashedPassword)

	res := database.Database.Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func FindUserById(id uint) (*User, error) {
	var user User
	res := database.Database.First(&user, id)

	return &user, res.Error
}

func FindUserByUsername(username string) (*User, error) {
	var user User
	res := database.Database.Where("username=?", username).First(&user)

	return &user, res.Error
}

func UpdateUser(id uint, update *User) (*User, error) {
	user, err := FindUserById(id)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := helpers.HashPassword(update.Password)

	if err != nil {
		return nil, err
	}

	user.setUsername(update.Username)
	user.setEmail(update.Email)
	user.setPassword(hashedPassword)
	user.setPhoto(update.Photo)

	res := database.Database.Save(user)
	return user, res.Error
}

func DeleteUser(id uint) error {
	_, err := FindUserById(id)

	if err != nil {
		return err
	}

	res := database.Database.Delete(&User{}, id)

	return res.Error
}
