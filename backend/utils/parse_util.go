package utils

import "strconv"

func ParseStrParamsToUint(paramStr string) (uint, error){
	paramUint64, err := strconv.ParseUint(paramStr, 10, 32)
	paramUint := uint(paramUint64)
	if err != nil {
		return 0, err
	}
	return paramUint, nil
	
}