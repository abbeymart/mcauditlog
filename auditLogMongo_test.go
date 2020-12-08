// @Author: abbeymart | Abi Akindele | @Created: 2020-12-05 | @Updated: 2020-12-05
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcauditlog

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"testing"
)
import (
	"github.com/abbeymart/mcdbgo"
	"github.com/abbeymart/mctestgo"
)

func TestAuditLogMongo(t *testing.T) {
	// test-data: db-configuration settings

	tableName := "services"
	userId := "085f48c5-8763-4e22-a1c6-ac1a68ba07de"
	recs := TestParam{Name: "Abi", Desc: "Testing only", Url: "localhost:9000", Priority: 1, Cost: 1000.00}
	tableRecords, _ := json.Marshal(recs)
	//fmt.Println("table-records-json", string(tableRecords))
	newRecs := TestParam{Name: "Abi Akindele", Desc: "Testing only - updated", Url: "localhost:9900", Priority: 1, Cost: 2000.00}
	newTableRecords, _ := json.Marshal(newRecs)
	//fmt.Println("new-table-records-json", string(newTableRecords))
	readP := map[string][]string{"keywords": {"lagos", "nigeria", "ghana", "accra"}}
	readParams, _ := json.Marshal(readP)

	myDb := mcdb.MongoDbConfig{
		DbType:   "mongodb",
		Host:     "localhost",
		Username: "abbeymart",
		Password: "ab12testing",
		Port:     27017,
		DbName:   "mcdev",
		Filename: "testdb.db",
		PoolSize: 20,
		Url:      "mongodb://localhost:27017",
	}
	myDb.Options = mcdb.MongoDbConnectOptions{}

	// db-connection
	dbc, err := myDb.OpenMongoDb()
	var mgDb *mongo.Database
	// fmt.Printf("*****dbc-info: %v\n", dbc)
	// defer dbClose
	defer myDb.CloseMongoDb()
	// check db-connection-error
	if err == nil {
		mgDb = dbc.Database(myDb.DbName)
	} else {
		fmt.Printf("*****db-connection-error: %v\n", err.Error())
		return
	}
	// expected db-connection result
	mcLogResult := LogParamMongo{AuditDb: mgDb, AuditTable: "audits"}
	// audit-log instance
	mcLog := NewAuditLogMongo(mgDb, "audits")

	mctest.McTest(mctest.OptionValue{
		Name: "should connect to the mongoDB and return an instance object:",
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
			mctest.AssertEquals(t, err, nil, "error-response should be: nil")
			mctest.AssertEquals(t, res.Code, "success", "log-action response-code should be: success")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should store read-transaction log and return success:",
		TestFunc: func() {
			res, err := mcLog.AuditLog(ReadLog, userId, AuditLogOptionsType{
				TableName:  tableName,
				LogRecords: string(readParams),
			})
			mctest.AssertEquals(t, err, nil, "error-response should be: nil")
			mctest.AssertEquals(t, res.Code, "success", "log-action response-code should be: success")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should store delete-transaction log and return success:",
		TestFunc: func() {
			res, err := mcLog.AuditLog(DeleteLog, userId, AuditLogOptionsType{
				TableName:  tableName,
				LogRecords: string(tableRecords),
			})
			mctest.AssertEquals(t, err, nil, "error-response should be: nil")
			mctest.AssertEquals(t, res.Code, "success", "log-action response-code should be: success")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should store login-transaction log and return success:",
		TestFunc: func() {
			res, err := mcLog.AuditLog(LoginLog, userId, AuditLogOptionsType{
				TableName:     tableName,
				LogRecords:    string(tableRecords),
			})
			mctest.AssertEquals(t, err, nil, "error-response should be: nil")
			mctest.AssertEquals(t, res.Code, "success", "log-action response-code should be: success")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should store logout-transaction log and return success:",
		TestFunc: func() {
			res, err := mcLog.AuditLog(LogoutLog, userId, AuditLogOptionsType{
				TableName:     tableName,
				LogRecords:    string(tableRecords),
			})
			mctest.AssertEquals(t, err, nil, "error-response should be: nil")
			mctest.AssertEquals(t, res.Code, "success", "log-action response-code should be: success")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should return paramsError for incomplete/undefined inputs:",
		TestFunc: func() {
			res, err := mcLog.AuditLog(CreateLog, "", AuditLogOptionsType{
				TableName:     tableName,
				LogRecords:    string(tableRecords),
			})
			//fmt.Printf("params-res: %#v", res)
			mctest.AssertNotEquals(t, err, nil, "error-response should not be: nil")
			mctest.AssertEquals(t, res.Code, "paramsError", "log-action response-code should be: paramsError")
			mctest.AssertEquals(t, strings.Contains(res.Message, "userId is required"), true, "log-action response-message should be: true")
			mctest.AssertEquals(t, strings.Contains(err.Error(), "userId is required"), true, "log-action error-message should be: true")
		},
	})

	mctest.PostTestResult()
}
