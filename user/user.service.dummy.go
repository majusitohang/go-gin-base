package user

import (
	"context"
	"fmt"
	"time"
)

type UserServiceDummy struct {
	data *[]UserModel
}

func NewUserServiceDummy() *UserServiceDummy {
	return &UserServiceDummy{data: &[]UserModel{
		{
			Id:          1,
			Name:        "maju",
			Gender:      "male",
			DateOfBirth: time.Time{},
		},
		{
			Id:          2,
			Name:        "dewi",
			Gender:      "female",
			DateOfBirth: time.Time{},
		},
		{
			Id:          3,
			Name:        "harry",
			Gender:      "male",
			DateOfBirth: time.Time{},
		},
	}}
}

func (u UserServiceDummy) List(ctx context.Context, page int, limit int) ([]UserModel, error) {
	return *u.data, nil
}

func (u UserServiceDummy) Detail(ctx context.Context, id int64) (*UserModel, error) {
	for _, v := range *u.data {
		if v.Id == id {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("id not found")
}

func (u UserServiceDummy) Create(ctx context.Context, input UserModel) (UserModel, error) {
	input.Id = time.Now().UnixNano()
	models := append(*u.data, input)
	u.data = &models
	return input, nil
}

func (u UserServiceDummy) Update(ctx context.Context, input UserModel) (*UserModel, error) {
	var existingData *UserModel
	for _, v := range *u.data {
		if v.Id == input.Id {
			existingData = &v
			v = input
		}
	}

	if existingData == nil {
		return nil, fmt.Errorf("id not found")
	}

	return existingData, nil
}

func (u UserServiceDummy) Delete(ctx context.Context, id int64) error {
	for _, v := range *u.data {
		if v.Id == id {
			v.Id = 0
		}
	}

	var newData []UserModel
	for _, v := range *u.data {
		if v.Id > 0 {
			newData = append(newData, v)
		}
	}

	u.data = &newData

	return nil
}
