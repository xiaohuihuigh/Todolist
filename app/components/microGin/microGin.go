package microGin

import (
	"todoList/app/options"
	"github.com/gin-gonic/gin"
)

type MicroGin struct {
	Engine *gin.Engine
	Listen string
}
func NewMicroGin()*MicroGin{
	return &MicroGin{
		Engine:gin.Default(),
		Listen:options.Options.GinService.Listen,
	}
}
func (m *MicroGin)Run(){

	m.Engine.Run(m.Listen)
}
