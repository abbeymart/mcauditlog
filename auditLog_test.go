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
	Name     string
	Desc     string
	Url      string
	Priority int
	Cost     float64
}

func TestSetCache(t *testing.T) {
	// test-data: db-configuration settings

	tableName := "services"
	userId := "085f48c5-8763-4e22-a1c6-ac1a68ba07de"
	recs := TestParam{Name: "Abi", Desc: "Testing only", Url: "localhost:9000", Priority: 1, Cost: 1000.00}
	tableRecords, _ := json.Marshal(recs)
	fmt.Println("table-records-json", string(tableRecords))
	newRecs := TestParam{Name: "Abi Akindele", Desc: "Testing only - updated", Url: "localhost:9900", Priority: 1, Cost: 2000.00}
	newTableRecords, _ := json.Marshal(newRecs)
	fmt.Println("new-table-records-json", string(newTableRecords))
	//loginParams = tableRecords
	//logoutParams = tableRecords

	myDb := mcdb.DbConfig{
		DbType: "postgres",
	}
	myDb.Host = "localhost"
	myDb.Username = "postgres"
	myDb.Password = "ab12testing"
	myDb.Port = 5432
	myDb.DbName = "mcdev"
	myDb.Filename = "testdb.db"
	myDb.PoolSize = 20
	myDb.Uri = "localhost:5432"
	myDb.Options = mcdb.DbConnectOptions{}

	//var (
	//	dbc mcdb.DbConnectionType
	//	err error
	//)

	// db-connection
	dbc, err := myDb.OpenDb()
	fmt.Printf("*****dbc-info: %v\n", dbc)
	// defer dbClose
	defer myDb.CloseDb()
	// expected db-connection result
	mcLogResult := LogParam{AuditDb: dbc, AuditTable: "audits"}
	// audit-log instance
	mcLog := NewAuditLog(dbc, "audits")

	mctest.McTest(mctest.OptionValue{
		Name: "should connect and return an instance object:",
		TestFunc: func() {
			mctest.AssertEquals(t, err, nil, "error-response should be: nil")
			mctest.AssertEquals(t, mcLog, mcLogResult, "db-connection instance should be: "+mcLogResult.String())
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should store create-transaction log and return success:",
		TestFunc: func() {
			res, err := mcLog.AuditLog(CreateLog, userId, AuditLogOptionsType{
				TableName:  tableName,
				LogRecords: string(tableRecords),
			})
			//fmt.Printf("create-log: %v", res)
			mctest.AssertEquals(t, err, nil, "error-response should be: nil")
			mctest.AssertEquals(t, res.Code, "success", "log-action response-code should be: success")
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should store update-transaction log and return success:",
		TestFunc: func() {
			res, err := mcLog.AuditLog(UpdateLog, userId, AuditLogOptionsType{
				TableName:     tableName,
				LogRecords:    string(tableRecords),
				NewLogRecords: string(newTableRecords),
			})
			//fmt.Printf("update-log: %v", res)
			mctest.AssertEquals(t, err, nil, "error-response should be: nil")
			mctest.AssertEquals(t, res.Code, "success", "log-action response-code should be: success")
		},
	})

	mctest.PostTestResult()
}
