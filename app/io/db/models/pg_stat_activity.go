package models

import "github.com/dimonrus/gomodel"

// Activity data
type Activity struct {
	// Process id
	PID *int32 `json:"pid"`
	// Database name
	DataBaseName *string `json:"dataBaseName"`
	// User name
	UserName *string `json:"userName"`
	// Application name
	ApplicationName *string `json:"applicationName"`
	// Client adders
	ClientAddress *string `json:"clientAddress"`
	// State of query
	State *string `json:"state"`
	// Executed query
	Query *string `json:"query"`
}

// Columns get columns
func (a *Activity) Columns() []string {
	return []string{"pid", "datname", "usename", "application_name", "client_addr", "state", "query"}
}

// Values get values
func (a *Activity) Values() []any {
	return []any{&a.PID, &a.DataBaseName, &a.UserName, &a.ApplicationName, &a.ClientAddress, &a.State, &a.Query}
}

// Table get table name
func (a *Activity) Table() string {
	return "pg_stat_activity"
}

// ActivityList List
type ActivityList []*Activity

// NewActivity New Activity method
func NewActivity() *Activity {
	return &Activity{}
}

// NewActivityCollection init Activity collection
func NewActivityCollection() *gomodel.Collection[Activity] {
	return gomodel.NewCollection[Activity]()
}
