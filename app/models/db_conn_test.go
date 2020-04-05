package models_test

import (
	"git.qutoutiao.net/todoList/app/models"
	"github.com/golib/assert"
	"testing"
)

func TestMain(m *testing.M) {
	err := models.Setup()
	it := assert.Assertions{}
	it.Empty(err)
	m.Run()
}
