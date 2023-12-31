package database

import (
	"github.com/breno5g/chi-rest-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var product []*entity.Product
	var err error

	if sort == "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if limit == 0 {
		limit = 10
	}

	if page != 0 {
		err = p.DB.Offset((page - 1) * limit).Limit(limit).Order("created_at " + sort).Find(&product).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&product).Error
	}

	return product, err
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product

	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
