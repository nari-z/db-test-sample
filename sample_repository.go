package main

import (
    "github.com/jinzhu/gorm"
)

type SampleRepository struct {
    *gorm.DB
}

func (p *SampleRepository) CreateModel(name string) error {
    model := &SampleModel{ Name: name }
    return p.DB.Create(model).Error
}

func (p *SampleRepository) GetModel(id int64) (*SampleModel, error) {
	var model SampleModel
	
    err := p.DB.Where("id = ?", id).Find(&model).Error

    return &model, err
}
