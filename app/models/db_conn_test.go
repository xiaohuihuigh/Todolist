package models_test

import (
	"todoList/app/models"
	"github.com/golib/assert"
	"testing"
)

func TestMain(m *testing.M) {
	err := models.Setup()
	it := assert.Assertions{}
	it.Empty(err)
	m.Run()
}
