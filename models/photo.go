package models

import "github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/database"

type Photo struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
	UserId   uint   `json:"userId"`
}

func (p *Photo) SetTitle(title string) {
	p.Title = title
}

func (p *Photo) SetCaption(caption string) {
	p.Caption = caption
}

func (p *Photo) SetPhotoUrl(photoUrl string) {
	p.PhotoUrl = photoUrl
}

func (p *Photo) SetUserId(userId uint) {
	p.UserId = userId
}

func CreatePhoto(photo *Photo) (*Photo, error) {
	res := database.Database.Create(photo)
	if res.Error != nil {
		return nil, res.Error
	}
	return photo, nil
}

func FindPhotoById(photoId uint) (*Photo, error) {
	var photo Photo
	res := database.Database.First(&photo, Photo{ID: photoId})
	return &photo, res.Error
}

func FindPhotoByUserId(userId uint) (*Photo, error) {
	var photo Photo
	res := database.Database.First(&photo, Photo{UserId: userId})
	return &photo, res.Error
}

func FindPhotosOfUser(userId uint) (*[]Photo, error) {
	var photos []Photo
	res := database.Database.Find(&photos, Photo{UserId: userId})
	return &photos, res.Error
}

func UpdatePhoto(id uint, update *Photo) (*Photo, error) {
	photo, err := FindPhotoById(id)

	if err != nil {
		return nil, err
	}

	photo.SetCaption(update.Caption) 
	photo.SetTitle(update.Title) 
	photo.SetPhotoUrl(update.PhotoUrl) 
	
	res := database.Database.Save(photo)

	return photo, res.Error
}

func DeletePhoto(id uint) error {
	_, err := FindPhotoById(id)

	if err != nil {
		return err
	}

	res := database.Database.Delete(&Photo{}, id)

	return res.Error
}