package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	paramNameID    = "id"
	paramNameEmail = "email"
)

func ExtractorRequestParamID(param *gin.Params) (uint64, error) {
	id, ok := param.Get(paramNameID)

	if !ok {
		return 0, errors.New("Param not found")
	}

	parseID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return 0, nil
	}
	return parseID, nil
}

//
//func ExtractorRequestParamEmail(param *gin.Params) (string, error) {
//	id, ok := param.Get(paramNameEmail)
//
//	if !ok {
//		return 0, errors.New("Param not found")
//	}
//
//	parseID, err := strconv.(id, 10, 64)
//
//	if err != nil {
//		return 0, nil
//	}
//	return parseID, nil
//}
