// @Author: abbeymart | Abi Akindele | @Created: 2020-12-05 | @Updated: 2020-12-05
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcauditlog

import (
	"encoding/json"
	"fmt"
	"testing"
)
import (
	"github.com/abbeymart/mcdbgo"
	"github.com/abbeymart/mctestgo"
)

type TestParam struct {
	name string
	desc string
	url string
	priority int
	cost float64
}

func TestSetCache(t *testing.T) {
	// test-data: db-configuration settings
	var (
		tableName = "services"
		userId    = "085f48c5-8763-4e22-a1c6-ac1a68ba07de"
		tableRecords, _ = json.Marshal(TestParam{
			name: "Abi",
			desc: "Testing only",
			url: "localhost:9000",
			priority: 1,
			cost: 1000.00,
		})
		tableNewRecords, _ = json.Marshal(TestParam{
			name: "Abi Akindele",
			desc: "Testing only - updated",
			url: "localhost:9900",
			priority: 1,
			cost: 2000.00,
		})
		//loginParams = tableRecords
		//logoutParams = tableRecords
	)

	myDb := mcdb.DbConfig{
		DbType: "postgres",
	}
	myDb.Host = "localhost"
	myDb.Username = "postgres"
	myDb.Password = "ab12trust"
	myDb.Port = 5432
	myDb.DbName = "mcdev"
	myDb.Filename = "testdb.db"
	myDb.PoolSize = 20
	myDb.Uri = "localhost:5432"
	myDb.Options = mcdb.DbConnectOptions{}

	var (
		dbc mcdb.DbConnectionType
		err error
	)

	// db-connection
	dbc, err = myDb.OpenDb()
	// expected db-connection result
	var logInstanceResult = LogParam{AuditDb: dbc, AuditTable: "audits"}
	// audit-log instance
	var mcLog = NewLog(dbc, "audits")

	mctest.McTest(mctest.OptionValue{
		Name: "should successfully connect to the database",
		TestFunc: func() {
			dbc, err := myDb.OpenDb()
			fmt.Println(dbc)
			mctest.AssertEquals(t, err, nil, "response-code should be: nil")
			//mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})

	// TODO: audit-log test cases

	if dbc != nil || err == nil {
		myDb.CloseDb()
	}

	mctest.PostTestResult()
}
