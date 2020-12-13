// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcauditlog

import (
	"context"
	"errors"
	"fmt"
	"github.com/abbeymart/mcresponse"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"time"
)

// interfaces / types
type PgxLogParam struct {
	AuditDb    *pgxpool.Pool
	AuditTable string
}

type PgxAuditLogOptionsType struct {
	AuditTable    string
	TableName     string
	LogRecords    interface{}
	NewLogRecords interface{}
	QueryParams   interface{}
}

func NewAuditLogPgx(auditDb *pgxpool.Pool, auditTable string) PgxLogParam {
	result := PgxLogParam{}
	result.AuditDb = auditDb
	result.AuditTable = auditTable
	// default value
	if result.AuditTable == "" {
		result.AuditTable = "audits"
	}
	return result
}

// String() function implementation
func (log PgxLogParam) String() string {
	return fmt.Sprintf(`
	AuditLog DB: %v \n AudiLog Table Name: %v \n
	`,
		log.AuditDb,
		log.AuditTable)
}

func (log PgxLogParam) AuditLog(logType, userId string, options PgxAuditLogOptionsType) (mcresponse.ResponseMessage, error) {
	// variables
	logType = strings.ToLower(logType)
	logBy := userId

	var (
		tableName     = ""
		sqlScript     = ""
		logRecords    interface{}
		newLogRecords interface{}
		logAt         time.Time

		dbResult pgconn.CommandTag
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
		dbResult, err = log.AuditDb.Exec(context.Background(), sqlScript, tableName, logRecords, logType, logBy, logAt)
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
			if errorMessage != "" {
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
		dbResult, err = log.AuditDb.Exec(context.Background(), sqlScript, tableName, logRecords, newLogRecords, logType, logBy, logAt)
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
		dbResult, err = log.AuditDb.Exec(context.Background(), sqlScript, tableName, logRecords, logType, logBy, logAt)
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
		dbResult, err = log.AuditDb.Exec(context.Background(), sqlScript, tableName, logRecords, logType, logBy, logAt)
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
		dbResult, err = log.AuditDb.Exec(context.Background(), sqlScript, tableName, logRecords, logType, logBy, logAt)
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
		dbResult, err = log.AuditDb.Exec(context.Background(), sqlScript, tableName, logRecords, logType, logBy, logAt)

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

