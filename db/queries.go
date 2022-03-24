package db

import (
	"uploadLargeFile/constants"
	"uploadLargeFile/models"
)

func FilterPriceForSymbol(symbol string, cursor, offset int) (*[]models.Prices, error) {
	var prices []models.Prices
	result := db.Offset(offset).Where("symbol = ? AND id > ?", symbol, cursor).Find(&prices)

	return &prices, result.Error
}


func FilterAllPrice(cursor, offset int) (*[]models.Prices, error) {
	var prices []models.Prices
	result := db.Offset(offset).Where("id > ?", cursor).Find(&prices)

	return &prices, result.Error
}


func CreateNewUpload(filename string) (*models.Uploads, error) {
	upload := models.Uploads{
		FileName: filename,
		Finish: false,
		ExpireTimeInSeconds: constants.UploadExpireTimeInSeconds,
	}
	result := db.Create(&upload)
	return &upload, result.Error
}

func GetUploads(pk int64) (*models.Uploads, error) {
	upload := models.Uploads{}
	result := db.Find(&upload, pk)
	return &upload, result.Error
}

func CreateNewChunk(upload *models.Uploads, partId, size int64, hash string) (*models.Chunks, error) {
	chunk := models.Chunks{
		Upload: *upload,
		PartNum: int(partId),
		Size: int(size),
		Hash: hash,
	}
	result := db.Create(&chunk)
	return &chunk, result.Error
}
