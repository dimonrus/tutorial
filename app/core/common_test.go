package core

import (
	"testing"
	"tutorial/app/base"
)

func TestCommon(t *testing.T) {
	base.App.SuccessMessage("Tests are work's fine")
}

func TestDatabaseVersion(t *testing.T) {
	var version string
	db := base.App.GetDB()
	err := db.QueryRow("SELECT VERSION();").Scan(&version)
	if err != nil {
		t.Fatal(err)
	}
	base.App.SuccessMessage("Postgres version is: " + version)
}
