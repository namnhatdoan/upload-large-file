package services

import (
	"mime/multipart"
	"uploadLargeFile/payloads"
)

type OHLCDataService interface {
	FilterData(queryFilter payloads.DataQueryFilter, pageFilter payloads.PageFilter) (data interface{}, err error)
	CreateNewUpload(fileName string) (data interface{}, err error)
	ProcessNewChunk(uploadId, partId int64, dataFile *multipart.FileHeader) (data interface{}, err error)
}
