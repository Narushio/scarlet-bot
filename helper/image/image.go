package image

import (
	"encoding/base64"
	"fmt"
)

type ContentType int

const (
	Jpeg ContentType = iota
	Png
	Gif
	Jpg
)

func (c *ContentType) String() string {
	switch *c {
	case Jpeg:
		return "jpeg"
	case Png:
		return "png"
	case Gif:
		return "gif"
	case Jpg:
		return "jpg"
	default:
		return "unknown"
	}
}

type Base64Src string

func ToBase64Src(fileBytes []byte, contentType ContentType) Base64Src {
	base64Str := base64.StdEncoding.EncodeToString(fileBytes)
	return Base64Src(fmt.Sprintf("data:image/%s;base64, %s", contentType.String(), base64Str))
}
