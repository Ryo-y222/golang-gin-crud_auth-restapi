package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"

	"github.com/davecgh/go-spew/spew"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput, userId uint) (*models.Item, error)
	Update(itemId uint, userId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint, userId uint) error
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint, userId uint) (*models.Item, error) {
	return s.repository.FindById(itemId, userId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput, userId uint) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
		UserID:      userId,
	}
	spew.Dump(newItem)
	// (models.Item) {
	// 	ID: (uint) 0,
	// 	Name: (string) (len=7) "商品6",
	// 	Price: (uint) 5000,
	// 	Description: (string) (len=19) "Postの動作確認",
	// 	SoldOut: (bool) false
	//    }
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, userId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := s.FindById(itemId, userId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return s.repository.Update(*targetItem)
}

func (s *ItemService) Delete(itemId uint, userId uint) error {
	return s.repository.Delete(itemId, userId) //エラーも返す
}
