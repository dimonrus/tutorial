package core

import (
	"github.com/dimonrus/godb/v2"
	"github.com/dimonrus/porterr"
	"tutorial/app/base"
	"tutorial/app/io/db/models"
)

// GetDatabaseActivity Get all db activities
func GetDatabaseActivity(q godb.Queryer) (models.ActivityList, porterr.IError) {
	collection := models.NewActivityCollection()
	collection.Where().AddExpression("usename = ?", base.App.GetConfig().Db.User)
	collection.AddOrder("pid DESC")
	e := collection.Load(q)
	return collection.Items(), e
}
