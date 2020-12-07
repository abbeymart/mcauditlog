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
	NewLogRecords interface{}
	QueryParams   interface{}
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

func NewAuditLog(auditDb *sql.DB, auditTable string) LogParam {
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
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logBy == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecords == nil {
			if errorMessage != "" {
				errorMessage = errorMessage + " | Created record(s) information is required."
			} else {
				errorMessage = "Created record(s) information is required."
			}
		}
		if errorMessage != "" {
			return mcresponse.GetResMessage("paramsError",
				mcresponse.ResponseMessageOptions{
					Message: errorMessage,
					Value:   nil,
				}), errors.New(errorMessage)
		}

		fmt.Println("before log-insert")
		// compose SQL-script
		sqlScript = fmt.Sprintf("INSERT INTO %v(table_name, log_records, log_type, log_by, log_at ) VALUES ($1, $2, $3, $4, $5)", log.AuditTable)

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)
		fmt.Printf("after log-insert: result => %v | err => %v \n", dbResult, err)
	case UpdateLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		newLogRecords = options.NewLogRecords
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logBy == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecords == nil {
			if errorMessage != "" {
				errorMessage = errorMessage + " | Updated record(s) information is required."
			} else {
				errorMessage = "Updated record(s) information is required."
			}
		}
		if newLogRecords == nil {
			if errorMessage != ""{
				errorMessage = errorMessage + " | New/Update record(s) information is required."
			} else {
				errorMessage = "New/Update record(s) information is required."
			}
		}
		if errorMessage != "" {
			return mcresponse.GetResMessage("paramsError",
				mcresponse.ResponseMessageOptions{
					Message: errorMessage,
					Value:   nil,
				}), errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = fmt.Sprintf("INSERT INTO %v(table_name, log_records, new_log_records, log_type, log_by, log_at ) VALUES ($1, $2, $3, $4, $5, $6)", log.AuditTable)

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, newLogRecords, logType, logBy, logAt)
	case GetLog, ReadLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logBy == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecords == nil {
			if errorMessage != "" {
				errorMessage = errorMessage + " | Read/Get Params/Keywords information is required."
			} else {
				errorMessage = "Read/Get Params/Keywords information is required."
			}
		}
		if errorMessage != "" {
			return mcresponse.GetResMessage("paramsError",
				mcresponse.ResponseMessageOptions{
					Message: errorMessage,
					Value:   nil,
				}), errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = fmt.Sprintf("INSERT INTO %v(table_name, log_records, log_type, log_by, log_at ) VALUES ($1, $2, $3, $4, $5)", log.AuditTable)

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)
	case DeleteLog, RemoveLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logBy == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecords == nil {
			if errorMessage != "" {
				errorMessage = errorMessage + " | Deleted record(s) information is required."
			} else {
				errorMessage = "Deleted record(s) information is required."
			}
		}
		if errorMessage != "" {
			return mcresponse.GetResMessage("paramsError",
				mcresponse.ResponseMessageOptions{
					Message: errorMessage,
					Value:   nil,
				}), errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = fmt.Sprintf("INSERT INTO %v(table_name, log_records, log_type, log_by, log_at ) VALUES ($1, $2, $3, $4, $5)", log.AuditTable)

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)
	case LoginLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logBy == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecords == nil {
			if errorMessage != "" {
				errorMessage = errorMessage + " | Login record(s) information is required."
			} else {
				errorMessage = "Login record(s) information is required."
			}
		}
		if errorMessage != "" {
			return mcresponse.GetResMessage("paramsError",
				mcresponse.ResponseMessageOptions{
					Message: errorMessage,
					Value:   nil,
				}), errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = fmt.Sprintf("INSERT INTO %v(table_name, log_records, log_type, log_by, log_at ) VALUES ($1, $2, $3, $4, $5)", log.AuditTable)

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)
	case LogoutLog:
		// set params
		tableName = options.TableName
		logRecords = options.LogRecords
		logAt = time.Now()

		// validate params
		var errorMessage = ""
		if tableName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logBy == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecords == nil {
			if errorMessage != "" {
				errorMessage = errorMessage + " | Logout record(s) information is required."
			} else {
				errorMessage = "Logout record(s) information is required."
			}
		}
		if errorMessage != "" {
			return mcresponse.GetResMessage("paramsError",
				mcresponse.ResponseMessageOptions{
					Message: errorMessage,
					Value:   nil,
				}), errors.New(errorMessage)
		}

		// compose SQL-script
		sqlScript = fmt.Sprintf("INSERT INTO %v(table_name, log_records, log_type, log_by, log_at ) VALUES ($1, $2, $3, $4, $5)", log.AuditTable)

		// perform db-log-insert action
		dbResult, err = log.AuditDb.Exec(sqlScript, tableName, logRecords, logType, logBy, logAt)

	default:
		return mcresponse.GetResMessage("logError",
			mcresponse.ResponseMessageOptions{
				Message: "Unknown log type and/or incomplete log information",
				Value:   nil,
			}), errors.New("unknown log type and/or incomplete log information")
	}

	// Handle error
	if err != nil {
		errMsg := fmt.Sprintf("%v", err)
		return mcresponse.GetResMessage("logError",
			mcresponse.ResponseMessageOptions{
				Message: errMsg,
				Value:   nil,
			}), errors.New(errMsg)
	}

	return mcresponse.GetResMessage("success",
		mcresponse.ResponseMessageOptions{
			Message: "successful audit-log action",
			Value:   dbResult,
		}), nil
}

func (log LogParam) CreateLog(table string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParam) UpdateLog(tableName string, logRecords interface{}, newLogRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParam) ReadLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParam) DeleteLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParam) LoginLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error) {
	// default-values
	if tableName == "" {
		tableName = "users"
	}

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParam) LogoutLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error) {
	// default-values
	if tableName == "" {
		tableName = "users"
	}

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}
