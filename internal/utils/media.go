package utils

import (
	"github.com/gabriel-vasile/mimetype"
)

const (
	CV                = 0
	Letter            = 1
	VideoPresentation = 2
	UserProfileImage  = 3
)

func IsValidDocument(mType string) bool {
	allowed := []string{"application/msword", "application/pdf"}
	return mimetype.EqualsAny(mType, allowed...)
}

func IsValidImage(mType string) bool {
	allowed := []string{"image/png", "image/jpeg"}
	return mimetype.EqualsAny(mType, allowed...)
}

func IsValidVideo(mType string) bool {
	allowed := []string{"video/mpeg", "video/mp4"}
	return mimetype.EqualsAny(mType, allowed...)
}

func IsValidFile(mType string) bool {
	allowed := []string{"text/plain", "image/png", "image/jpeg", "application/word", "application/pdf", "video/mpeg", "video/mp4"}
	return mimetype.EqualsAny(mType, allowed...)
}
