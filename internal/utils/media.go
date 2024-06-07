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
	allowed := []string{
		"image/png",
		"image/jpeg",
		"image/jpg",
		"image/gif",
		"image/bmp",
		"image/tiff",
		"image/webp",
		"image/svg",
		"image/x-ico",
		"image/heic",
	}
	return mimetype.EqualsAny(mType, allowed...)
}

func IsValidVideo(mType string) bool {
	allowed := []string{
		"video/mpeg",
		"video/mp4",
		"video/x-msvideo",  //avi
		"video/x-matroska", //mkv
		"video/quicktime",  //mov
		"video/x-ms-wmv",   //wmv
		"video/x-flv",      //flv
		"video/webm",       //webm
		"video/mpeg",       //mpeg
	}
	return mimetype.EqualsAny(mType, allowed...)
}

func IsValidFile(mType string) bool {
	allowed := []string{
		//image
		"image/png",
		"image/jpeg",
		"image/jpg",
		"image/gif",
		"image/bmp",
		"image/tiff",
		"image/webp",
		"image/svg",
		"image/x-ico",
		"image/heic",
		//document
		"text/plain",
		"application/word",
		"application/pdf",
		//Video
		"video/mpeg",
		"video/mp4",
		"video/x-msvideo",  //avi
		"video/x-matroska", //mkv
		"video/quicktime",  //mov
		"video/x-ms-wmv",   //wmv
		"video/x-flv",      //flv
		"video/webm",       //webm
		"video/mpeg",
	}
	return mimetype.EqualsAny(mType, allowed...)
}
