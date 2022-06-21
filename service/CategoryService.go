package service

import (
	"gin_vuePQ/model"
	"gin_vuePQ/repository"
)

type ICategoryService interface {
	Create(name string) (*model.Category, error)
	Update(category model.Category, name string) (*model.Category, error)
	Show(id int) (*model.Category, error)
	Delete(id int) error
}

func NewCategoryService() ICategoryService {
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.DB.AutoMigrate(model.Category{})
	return CategoryService{Repository: categoryRepository}
}

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (c CategoryService) Create(name string) (*model.Category, error) {
	category, err := c.Repository.Create(name)
	if err != nil {
		return nil, err
	}
	return category, err
}

func (c CategoryService) Update(category model.Category, name string) (*model.Category, error) {
	data, err := c.Repository.Update(category, name)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (c CategoryService) Show(id int) (*model.Category, error) {
	category, err := c.Repository.SelectById(id)
	if err != nil {
		return nil, err
	}
	return category, err
}

func (c CategoryService) Delete(id int) error {
	err := c.Repository.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
