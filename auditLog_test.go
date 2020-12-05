// @Author: abbeymart | Abi Akindele | @Created: 2020-12-05 | @Updated: 2020-12-05
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcauditlog

import (
	"fmt"
	"testing"
)
import (
	"github.com/abbeymart/mcdbgo"
	"github.com/abbeymart/mctestgo"
)

func TestSetCache(t *testing.T) {
	// test-data: db-configuration settings
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
