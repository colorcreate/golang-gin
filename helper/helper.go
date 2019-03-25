package helper

import "strconv"

// CUint is
func CUint(strID string) uint {
	u64, _ := strconv.ParseUint(strID, 10, 32)
	return uint(u64)
}

// CStr is
func CStr(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}
