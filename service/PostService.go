package service

import (
	"gin_vuePQ/model"
	"gin_vuePQ/repository"
)

type IPostService interface {
	Create(post *model.Post) error
	Update(post *model.Post) error
	Show(postId string) (*model.Post, error)
	Delete(post *model.Post) error
	PageList(pageNum, pageSize int) ([]model.Post, int64, error)
}

type PostService struct {
	Repository repository.PostRepository
}

func NewPostService() IPostService {
	postRepository := repository.NewPostRepository()
	postRepository.DB.AutoMigrate(model.Post{})
	return PostService{Repository: postRepository}
}

func (p PostService) Create(post *model.Post) error {
	err := p.Repository.Create(post)
	if err != nil {
		return err
	}
	return nil
}

func (p PostService) Update(post *model.Post) error {
	err := p.Repository.Update(post)
	if err != nil {
		return err
	}
	return nil
}

func (p PostService) Show(postId string) (*model.Post, error) {
	post, err := p.Repository.Show(postId)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p PostService) Delete(post *model.Post) error {
	err := p.Repository.Delete(post)
	if err != nil {
		return err
	}
	return nil
}

func (p PostService) PageList(pageNum, pageSize int) ([]model.Post, int64, error) {
	pageList, total, err := p.Repository.PageList(pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return pageList, total, nil
}
