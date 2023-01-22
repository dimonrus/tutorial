package core

import (
	"tutorial/app/base"
	"testing"
)

func TestCommon(t *testing.T) {
	base.App.SuccessMessage("Tests are work's fine")
}
