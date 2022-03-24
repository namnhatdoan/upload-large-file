package payloads

import (
	"uploadLargeFile/models"
	"time"
)

type UploadFileRequest struct{
	UploadId string `json:"upload_id"`
	FileName string `json:"file_name"`
	Ctime time.Time `json:"ctime"`
}

type UploadFileResponse struct {
	CommonResponse
	data string
}

type DataQueryRequest struct {
	DataQueryFilter
	PageFilter
}

type DataQueryFilter struct {
	Symbol string `json:"symbol"`
}

type DataQueryResponse struct {
	CommonResponse
	Data *[]models.Prices `json:"data"`
}
