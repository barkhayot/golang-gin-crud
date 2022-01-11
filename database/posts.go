package database

import (
	"editt/models"
	"github.com/jinzhu/gorm"
)

/* Getting all Posts */
func GetPosts(db *gorm.DB) ([]models.Post, error) {
	posts := []models.Post{}
	query := db.Select("posts.*").
			Group("posts.id")
	if err := query.Find(&posts).Error; err != nil {
		return posts, err
	}
	return posts, nil
}

/* Getting Posts by ID */
func GetPostByID(id string, db *gorm.DB) (models.Post, bool, error) {
	p := models.Post{}

	query := db.Select("posts.*")
	query = query.Group("posts.id")
	err := query.Where("posts.id = ?", id).First(&p).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return p, false, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return p, false, nil
	}
	return p, true, nil
}

/* Deleting Posts by ID */
func DeletePost(id string, db *gorm.DB) error {
	var p models.Post
	if err := db.Where("id = ? ", id).Delete(&p).Error; err != nil {
		return err
	}
	return nil
}

/* Update Posts function */
func UpdatePost(db *gorm.DB, b *models.Post) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}