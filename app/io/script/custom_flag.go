// script file
package script

import (
	"github.com/dimonrus/gocli"
	"tutorial/app/base"
)

func init() {
	base.App.GetScripts()["custom_flag"] = func(args gocli.Arguments) {
		color := args["color"]
		base.App.GetLogger().Infof("Your color is: %s", color.GetString())
	}
}
