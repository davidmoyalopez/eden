package lands

import (
	"eden/utils/errors"
	"strings"
)

type Land struct {
	Id      int64  `json:"id"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	TTL     int    `json:"ttl"`
}

func (land *Land) Validate() *errors.RestErr {
	land.Name = strings.TrimSpace(strings.ToLower(land.Name))
	if land.Name == "" {
		return errors.NewBadRequestError("invalid name")
	}
	return nil
}
