package models_test

import (
	"todoList/app/models"
	"github.com/golib/assert"
	"testing"
)

func TestUserInfoExpandModel(t *testing.T) {
	it := assert.New(t)

	task1 := models.NewUserInfoExpandModel()
	err := task1.FindByID(1)
	it.Empty(err)
}