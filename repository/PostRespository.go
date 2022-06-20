package repository

import (
	"gin_vuePQ/common"
	"gin_vuePQ/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository() PostRepository {
	return PostRepository{DB: common.GetDB()}
}

func (p PostRepository) Create(post model.Post) error {

	if err := p.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil

}

func (p PostRepository) Update(post model.Post) error {

	// 更新文章
	if err := p.DB.Save(&post).Error; err != nil {
		return err
	}

	return nil
}

func (p PostRepository) Show(postId string) (model.Post, error) {

	var post model.Post
	if err := p.DB.Preload("Category").Where("id = ?", postId).First(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

func (p PostRepository) Delete(post model.Post) error {

	if err := p.DB.Delete(&post).Error; err != nil {
		return err
	}

	return nil
}

func (p PostRepository) PageList(pageNum, pageSize int) ([]model.Post, int64, error) {

	// 分页
	var posts []model.Post
	err := p.DB.Preload("Category").Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}
	// 前端渲染分页需要知道记录总数
	var total int64
	p.DB.Model(model.Post{}).Count(&total)

	return posts, total, nil
}
