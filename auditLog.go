// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcauditlog

import (
	"database/sql"
	"github.com/abbeymart/mcresponsego"
)

// types
type LogParam struct {
	AuditDb    *sql.DB
	AuditTable string
}
type AuditLogOptionsType struct {
	AuditTable        string
	TableName         string
	TableDocuments    interface{}
	newTableDocuments interface{}
	recordParams      interface{}
	newRecordParams   interface{}
}

func NewLog(auditDb *sql.DB, auditTable string) LogParam {
	result := LogParam{}
	result.AuditDb = auditDb
	result.AuditTable = auditTable
	// default value
	if result.AuditTable == "" {
		result.AuditTable = "audits"
	}
	return result
}

func CreateLog(log LogParam, table string, tableRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}

func UpdateLog(log LogParam, table string, tableRecords interface{}, newTableRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}

func ReadLog(log LogParam, table string, tableRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}

func DeleteLog(log LogParam, table string, tableRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}

func LoginLog(log LogParam, tableRecords interface{}, userId string, table string, ) (mcresponse.ResponseMessage, error) {
	// default-values
	if table == "" {
		table = "users"
	}

	return mcresponse.ResponseMessage{}, nil
}

func LogoutLog(log LogParam, tableRecords interface{}, userId string, table string, ) (mcresponse.ResponseMessage, error) {
	// default-values
	if table == "" {
		table = "users"
	}

	return mcresponse.ResponseMessage{}, nil
}

func AuditLog(log LogParam, logType, userId string, options AuditLogOptionsType) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}