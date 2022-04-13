package services

import (
	"eden/domain/lands"
	"eden/utils/errors"
)

func GetLand(landId int64) (*lands.Land, *errors.RestErr) {
	result := &lands.Land{Id: landId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateLand(land lands.Land) (*lands.Land, *errors.RestErr) {
	if err := land.Validate(); err != nil {
		return nil, err
	}

	if err := land.Save(); err != nil {
		return nil, err
	}

	return &land, nil
}
