// data Migration file
package data

import (
	"github.com/dimonrus/godb/v2"
	"github.com/dimonrus/gohelp"
	"github.com/dimonrus/gomodel"
	"tutorial/app/base"
)

type m_000000000_init struct{}

func init() {
	base.App.GetMigration().Registry["data"] = append(base.App.GetMigration().Registry["data"], m_000000000_init{})
}

func (m m_000000000_init) GetVersion() string {
	return "m_000000000_init"
}

func (m m_000000000_init) Up(tx *godb.SqlTx) error {
	collection := gomodel.NewCollection[gomodel.DictionaryModel]()
	collection.AddItem(
		&gomodel.DictionaryModel{
			Id:    gohelp.Ptr[int32](1000),
			Type:  gohelp.Ptr("test_category"),
			Code:  gohelp.Ptr("white"),
			Label: gohelp.Ptr("White"),
		},
		&gomodel.DictionaryModel{
			Id:    gohelp.Ptr[int32](1010),
			Type:  gohelp.Ptr("test_category"),
			Code:  gohelp.Ptr("blue"),
			Label: gohelp.Ptr("Blue"),
		},
		&gomodel.DictionaryModel{
			Id:    gohelp.Ptr[int32](1020),
			Type:  gohelp.Ptr("test_category"),
			Code:  gohelp.Ptr("red"),
			Label: gohelp.Ptr("Red"),
		})
	return collection.Save(tx)
}

func (m m_000000000_init) Down(tx *godb.SqlTx) error {
	return nil
}
