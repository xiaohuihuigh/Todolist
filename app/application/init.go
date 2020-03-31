package application

import (
	"fmt"
	"git.qutoutiao.net/todoList/app/components/microGin"
	"git.qutoutiao.net/todoList/app/options"
	"os"
)
var ApplicationInif *_applicationInit
type _applicationInit struct {

}
func(i *_applicationInit)Init()*microGin.MicroGin {
	err := options.NewAppOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to created a new app option, %v\n", err)
		return nil
	}
	return microGin.NewMicroGin()
}