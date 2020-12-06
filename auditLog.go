// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcauditlog

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/abbeymart/mcresponsego"
	"strings"
	"time"
)

// interfaces / types
type LogParam struct {
	AuditDb    *sql.DB
	AuditTable string
}

type AuditLogOptionsType struct {
	AuditTable    string
	TableName     string
	LogRecords    interface{}
	newLogRecords interface{}
	queryParams   interface{}
}

type AuditLogger interface {
	AuditLog(logType, userId string, options AuditLogOptionsType) (mcresponse.ResponseMessage, error)
}
type CreateLogger interface {
	CreateLog(table string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error)
}
type UpdateLogger interface {
	UpdateLog(tableName string, logRecords interface{}, newLogRecords interface{}, userId string) (mcresponse.ResponseMessage, error)
}
type ReadLogger interface {
	ReadLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error)
}
type DeleteLogger interface {
	DeleteLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error)
}
type AccessLogger interface {
	LoginLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error)
	LogoutLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error)
}

//type AuditCrudLogger interface {
//	CreateLog(table string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error)
//	UpdateLog(tableName string, logRecords interface{}, newLogRecords interface{}, userId string) (mcresponse.ResponseMessage, error)
//	ReadLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error)
//	DeleteLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error)
//	LoginLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error)
//	LogoutLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error)
//	AuditLog(logType, userId string, options AuditLogOptionsType) (mcresponse.ResponseMessage, error)
//}

// constants
// LogTypes
const (
	CreateLog = "create"
	UpdateLog = "update"
	ReadLog   = "read"
	GetLog    = "get"
	DeleteLog = "delete"
	RemoveLog = "remove"
	LoginLog  = "login"
	LogoutLog = "logout"
)

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

// String() function implementation
func (log LogParam) String() string {
	return fmt.Sprintf(`
	AuditLog DB: %v \n AudiLog Table Name: %v \n
	`,
		log.AuditDb,
		log.AuditTable)
}

func (log LogParam) CreateLog(table string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}

func (log LogParam) UpdateLog(tableName string, logRecords interface{}, newLogRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}

func (log LogParam) ReadLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}

func (log LogParam) DeleteLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.ResponseMessage{}, nil
}

func (log LogParam) LoginLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error) {
	// default-values
	if tableName == "" {
		tableName = "users"
	}

	return mcresponse.ResponseMessage{}, nil
}

func (log LogParam) LogoutLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error) {
	// default-values
	if tableName == "" {
		tableName = "users"
	}

	return mcresponse.ResponseMessage{}, nil
}

func (log LogParam) AuditLog(logType, userId string, options AuditLogOptionsType) (mcresponse.ResponseMessage, error) {
	// variables
	logType = strings.ToLower(logType)
	logBy := userId

	var (
		tableName     = ""
		sqlScript     = ""
		logRecords    interface{}
		newLogRecords interface{}
		logAt         time.Time

		dbResult sql.Result
		err      error
	)

	// log-cases
	switch logType {
	case CreateLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logType = CreateLog
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = errorMessage + " | Table or Collection name is required."
		}
		if logBy == "" {
			errorMessage = errorMessage + " | userId is required."
		}
		if logRecords == nil {
			errorMessage = errorMessage + " | Created record(s) information is required."
		}
		if errorMessage != "" {
			return mcresponse.ResponseMessage{}, errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = "INSERT INTO " + tableName + " (table_name, log_records, log_type, log_by, log_at ) VALUES (?, ?, ?, ?, ?);"

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)
	case UpdateLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		newLogRecords = options.newLogRecords
		logType = CreateLog
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = errorMessage + " | Table or Collection name is required."
		}
		if logBy == "" {
			errorMessage = errorMessage + " | userId is required."
		}
		if logRecords == nil {
			errorMessage = errorMessage + " | Updated record(s) information is required."
		}
		if newLogRecords == nil {
			errorMessage = errorMessage + " | New/Update record(s) information is required."
		}
		if errorMessage != "" {
			return mcresponse.ResponseMessage{}, errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = "INSERT INTO " + tableName + " (table_name, log_records, log_new_records log_type, log_by, log_at ) VALUES (?, ?, ?, ?, ?, ?);"

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, newLogRecords, logType, logBy, logAt)
	case GetLog, ReadLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logType = CreateLog
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = errorMessage + " | Table or Collection name is required."
		}
		if logBy == "" {
			errorMessage = errorMessage + " | userId is required."
		}
		if logRecords == nil {
			errorMessage = errorMessage + " | Read/Get params/keyword information is required."
		}
		if errorMessage != "" {
			return mcresponse.ResponseMessage{}, errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = "INSERT INTO " + tableName + " (table_name, log_records, log_type, log_by, log_at ) VALUES (?, ?, ?, ?, ?);"

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)
	case DeleteLog, RemoveLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logType = CreateLog
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = errorMessage + " | Table or Collection name is required."
		}
		if logBy == "" {
			errorMessage = errorMessage + " | userId is required."
		}
		if logRecords == nil {
			errorMessage = errorMessage + " | Deleted record(s) information is required."
		}
		if errorMessage != "" {
			return mcresponse.ResponseMessage{}, errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = "INSERT INTO " + tableName + " (table_name, log_records, log_type, log_by, log_at ) VALUES (?, ?, ?, ?, ?);"

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)
	case LoginLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logType = CreateLog
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = errorMessage + " | Table or Collection name is required."
		}
		if logBy == "" {
			errorMessage = errorMessage + " | userId is required."
		}
		if logRecords == nil {
			errorMessage = errorMessage + " | Login record(s) information is required."
		}
		if errorMessage != "" {
			return mcresponse.ResponseMessage{}, errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = "INSERT INTO " + tableName + " (table_name, log_records, log_type, log_by, log_at ) VALUES (?, ?, ?, ?, ?);"

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)
	case LogoutLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logType = CreateLog
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = errorMessage + " | Table or Collection name is required."
		}
		if logBy == "" {
			errorMessage = errorMessage + " | userId is required."
		}
		if logRecords == nil {
			errorMessage = errorMessage + " | Logout record(s) information is required."
		}
		if errorMessage != "" {
			return mcresponse.ResponseMessage{}, errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = "INSERT INTO " + tableName + " (table_name, log_records, log_type, log_by, log_at ) VALUES (?, ?, ?, ?, ?);"

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)

	default:
		return mcresponse.ResponseMessage{
			Code:    "insertError",
			Message: "Unknown log type and/or incomplete log information",
		}, nil
	}

	// Handle error
	if err != nil {
		return mcresponse.ResponseMessage{}, errors.New(fmt.Sprintf("%v", err))
	}
	// result
	//type Result interface {
	//	LastInsertId() (int64, error)
	//	RowsAffected() (int64, error)
	//}
	return mcresponse.ResponseMessage{
		Code:    "success",
		Message: "successful create-log action",
		Value:   dbResult,
	}, nil
}
