package handlers

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"uploadLargeFile/constants"
	"uploadLargeFile/payloads"
	"uploadLargeFile/services"
	"strconv"
	"strings"
)

type HandlerImpl struct {
	service services.OHLCDataService
}

func InitHandlerImpl(service services.OHLCDataService) Handler {
	return &HandlerImpl{
		service: service,
	}
}

func (h *HandlerImpl) GetData(c *gin.Context) {
	req := initDefaultGetDataReq()
	if success := bindingQueryData(c, &req); !success {
		return
	}

	if data, err := h.service.FilterData(req.DataQueryFilter, req.PageFilter); err != nil {
		responseGeneralFailure(c, err.Error())
	} else {
		responseSuccess(c, data)
	}
}

func initDefaultGetDataReq() payloads.DataQueryRequest {
	return payloads.DataQueryRequest{
		DataQueryFilter: payloads.DataQueryFilter{
			Symbol: "",
		},
		PageFilter: payloads.PageFilter{
			Cursor: 0,
			Limit: constants.DefaultPaginationLimit,
		},
	}
}

func (h *HandlerImpl) UploadData(c *gin.Context) {
	uploadId, partId, filename, success := parseUploadDataForm(c)
	if !success {
		return
	}
	if uploadId <= 0 && filename == "" {
		responseBadRequest(c, "upload_id or filename must be provided")
		return
	}

	var result interface{}
	var err error

	if uploadId > 0 {
		var dataFile *multipart.FileHeader
		dataFile, err = c.FormFile("file")
		if err != nil {
			responseGeneralFailure(c, err.Error())
		}
		result, err = h.service.ProcessNewChunk(uploadId, partId, dataFile)
	} else {
		result, err = h.service.CreateNewUpload(filename)
	}

	if err != nil {
		responseGeneralFailure(c, err.Error())
	} else {
		responseSuccess(c, result)
	}
}

func parseUploadDataForm(c *gin.Context) (uploadId, partId int64, filename string, success bool) {
	var err error
	uploadIdStr, uploadIdExist := c.GetPostForm("upload_id")
	if uploadIdExist {
		uploadId, err = strconv.ParseInt(uploadIdStr, 10, 64)
		if err != nil || uploadId <= 0 {
			// Add log here
			responseBadRequest(c, "Invalid upload id")
			return
		}
	}

	partIdStr, partIdExist := c.GetPostForm("part_id")
	if partIdExist {
		partId, err = strconv.ParseInt(partIdStr, 10, 64)
		if err != nil || partId <= 0 {
			// Add log here
			responseBadRequest(c, "Invalid part id")
			return
		}
	} else if uploadIdExist {
		responseBadRequest(c, "Missing part id for next chunk")
		return
	}
	filename, _ = c.GetPostForm("filename")
	filename = strings.TrimSpace(filename)

	success = true
	return
}
