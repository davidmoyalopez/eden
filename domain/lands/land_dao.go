package lands

import (
	"eden/utils/errors"
	"fmt"
)

var (
	landDB = make(map[int64]*Land)
)

func (land *Land) Get() *errors.RestErr {
	result := landDB[land.Id]
	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d not found", land.Id))
	}

	land.Id = result.Id
	land.Name = result.Name
	land.TTL = result.TTL
	land.Version = result.Version
	land.Type = result.Type

	return nil
}

func (land *Land) Save() *errors.RestErr {
	current := landDB[land.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("land %d already exists", land.Id))
	}
	landDB[land.Id] = land
	return nil
}
