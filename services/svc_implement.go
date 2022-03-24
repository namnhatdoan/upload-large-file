package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"uploadLargeFile/db"
	"uploadLargeFile/payloads"
	"os"
	"strconv"
)

type OHLCServiceImpl struct {}

func (svc *OHLCServiceImpl) FilterData(queryFilter payloads.DataQueryFilter, pageFilter payloads.PageFilter) (data interface{}, err error) {
	if queryFilter.Symbol == "" {
		return db.FilterPriceForSymbol(queryFilter.Symbol, pageFilter.Cursor, pageFilter.Limit)
	}
	return db.FilterAllPrice(pageFilter.Cursor, pageFilter.Limit)
}


func (svc *OHLCServiceImpl) CreateNewUpload(fileName string) (data interface{}, err error) {
	return db.CreateNewUpload(fileName)
}

func (svc *OHLCServiceImpl) ProcessNewChunk(uploadId, partId int64, dataFile *multipart.FileHeader) (data interface{}, err error) {
	f, err := dataFile.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	out, err := createFile(uploadId, partId)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	size, err := io.Copy(out, f)
	if err != nil {
		return nil, err
	}

	uploads, err := db.GetUploads(uploadId)
	if err != nil {
		return nil, err
	}

	result, err := db.CreateNewChunk(uploads, partId, size, "hashing_here")
	return result, err
}

func createFile(uploadId, partId int64) (*os.File, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil ,err
	}
	if err := makeDir(strconv.FormatInt(uploadId, 10)); err != nil && !os.IsExist(err){
		return nil, err
	}

	filePath := fmt.Sprintf("%v/%v/%v", currentDir, uploadId, partId)
	return os.Create(filePath)
}

func makeDir(dir string) error {
	return os.Mkdir(dir, 0755)
}