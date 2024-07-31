package controllers

import (
	"log"

	"github.com/luisupbeat/go-gorm/db"
	"github.com/luisupbeat/go-gorm/models"
)

func GetHaciendas() ([]models.Hacienda, error) {
	var haciendas []models.Hacienda
	result := db.DB.Find(&haciendas)

	if result.Error != nil {
		return nil, result.Error
	}
	return haciendas, nil
}

func SaveHacienda(nuevaHacienda models.Hacienda) {

	result := db.DB.Table("HACIENDA").Create(&nuevaHacienda)

	if result.Error != nil {
		log.Fatalf("Error al guardar la hacienda: %v", result.Error)
	}
}

func DeleteHacienda() ([]models.Hacienda, error) {
	var haciendas []models.Hacienda
	result := db.DB.Table("HACIENDA").Find(&haciendas)

	if result.Error != nil {
		return nil, result.Error
	}
	return haciendas, nil
}

func UpdateHacienda() ([]models.Hacienda, error) {
	var haciendas []models.Hacienda
	result := db.DB.Table("HACIENDA").Find(&haciendas)

	if result.Error != nil {
		return nil, result.Error
	}
	return haciendas, nil
}

func GetHacienda() ([]models.Hacienda, error) {
	var haciendas []models.Hacienda
	result := db.DB.Table("HACIENDA").Find(&haciendas)

	if result.Error != nil {
		return nil, result.Error
	}
	return haciendas, nil
}
