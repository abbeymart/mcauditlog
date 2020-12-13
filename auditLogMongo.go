// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcauditlog

import (
	"context"
	"errors"
	"fmt"
	"github.com/abbeymart/mcresponse"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

// interfaces / types
type LogParamMongo struct {
	AuditDb    *mongo.Database
	AuditTable string
}

type AuditRecord struct {
	TableName     string      `json:"table_name,omitempty"`
	LogType       string      `json:"log_type,omitempty"`
	LogBy         string      `json:"log_by,omitempty"`
	LogAt         time.Time   `json:"log_at,omitempty"`
	LogRecords    interface{} `json:"log_records,omitempty"`
	NewLogRecords interface{} `json:"new_log_records,omitempty"`
}

func NewAuditLogMongo(auditDb *mongo.Database, auditTable string) LogParamMongo {
	result := LogParamMongo{}
	result.AuditDb = auditDb
	result.AuditTable = auditTable
	// default value
	if result.AuditTable == "" {
		result.AuditTable = "audits"
	}
	return result
}

// String() function implementation
func (log LogParamMongo) String() string {
	return fmt.Sprintf(`
	AuditLog DB: %v \n AudiLog Table Name: %v \n
	`,
		log.AuditDb,
		log.AuditTable)
}

func (log LogParamMongo) AuditLog(logType, userId string, options AuditLogOptionsType) (mcresponse.ResponseMessage, error) {
	// variables
	loggingType := strings.ToLower(logType)
	logUserId := userId

	var (
		tabName     = ""
		logRecs     interface{}
		newLogRecs  interface{}
		loggingAt   time.Time
		auditRecord AuditRecord
	)

	// log-cases
	switch loggingType {
	case CreateLog:
		// set params
		tabName = options.TableName
		logRecs = options.LogRecords
		loggingAt = time.Now()

		// validate params
		var errorMessage = ""
		if tabName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logUserId == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecs == nil {
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

		// define audit-record
		auditRecord = AuditRecord{
			TableName:  tabName,
			LogRecords: logRecs,
			LogType:    loggingType,
			LogBy:      logUserId,
			LogAt:      loggingAt,
		}
	case UpdateLog:
		// set params
		tabName = options.TableName
		logRecs = options.LogRecords
		newLogRecs = options.NewLogRecords
		loggingAt = time.Now()

		// validate params
		var errorMessage = ""
		if tabName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logUserId == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecs == nil {
			if errorMessage != "" {
				errorMessage = errorMessage + " | Updated record(s) information is required."
			} else {
				errorMessage = "Updated record(s) information is required."
			}
		}
		if newLogRecs == nil {
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

		// define audit-record
		auditRecord = AuditRecord{
			TableName:     tabName,
			LogRecords:    logRecs,
			NewLogRecords: newLogRecs,
			LogType:       loggingType,
			LogBy:         logUserId,
			LogAt:         loggingAt,
		}
	case GetLog, ReadLog:
		// set params
		tabName = options.TableName
		logRecs = options.LogRecords
		loggingAt = time.Now()

		// validate params
		var errorMessage = ""
		if tabName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logUserId == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecs == nil {
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

		// define audit-record
		auditRecord = AuditRecord{
			TableName:  tabName,
			LogRecords: logRecs,
			LogType:    logType,
			LogBy:      logUserId,
			LogAt:      loggingAt,
		}
	case DeleteLog, RemoveLog:
		// set params
		tabName = options.TableName
		logRecs = options.LogRecords
		loggingAt = time.Now()

		// validate params
		var errorMessage = ""
		if tabName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logUserId == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecs == nil {
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

		// define audit-record
		auditRecord = AuditRecord{
			TableName:  tabName,
			LogRecords: logRecs,
			LogType:    logType,
			LogBy:      logUserId,
			LogAt:      loggingAt,
		}
	case LoginLog:
		// set params
		tabName = options.TableName
		logRecs = options.LogRecords
		loggingAt = time.Now()

		// validate params
		var errorMessage = ""
		if tabName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logUserId == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecs == nil {
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

		// define audit-record
		auditRecord = AuditRecord{
			TableName:  tabName,
			LogRecords: logRecs,
			LogType:    logType,
			LogBy:      logUserId,
			LogAt:      loggingAt,
		}
	case LogoutLog:
		// set params
		tabName = options.TableName
		logRecs = options.LogRecords
		loggingAt = time.Now()

		// validate params
		var errorMessage = ""
		if tabName == "" {
			errorMessage = "Table or Collection name is required."
		}
		if logUserId == "" {
			if errorMessage != "" {
				errorMessage = errorMessage + " | userId is required."
			} else {
				errorMessage = "userId is required."
			}
		}
		if logRecs == nil {
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

		// define audit-record
		auditRecord = AuditRecord{
			TableName:  tabName,
			LogRecords: logRecs,
			LogType:    logType,
			LogBy:      logUserId,
			LogAt:      loggingAt,
		}
	default:
		return mcresponse.GetResMessage("logError",
			mcresponse.ResponseMessageOptions{
				Message: "Unknown log type and/or incomplete log information",
				Value:   nil,
			}), errors.New("unknown log type and/or incomplete log information")
	}

	// perform db-log-insert action
	dbColl := log.AuditDb.Collection(log.AuditTable)
	dbResult, err := dbColl.InsertOne(context.TODO(), auditRecord)

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

func (log LogParamMongo) CreateLog(table string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {
	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParamMongo) UpdateLog(tableName string, logRecords interface{}, newLogRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParamMongo) ReadLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParamMongo) DeleteLog(tableName string, logRecords interface{}, userId string) (mcresponse.ResponseMessage, error) {

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParamMongo) LoginLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error) {
	// default-values
	if tableName == "" {
		tableName = "users"
	}

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}

func (log LogParamMongo) LogoutLog(logRecords interface{}, userId string, tableName string) (mcresponse.ResponseMessage, error) {
	// default-values
	if tableName == "" {
		tableName = "users"
	}

	return mcresponse.GetResMessage("success", mcresponse.ResponseMessageOptions{}), nil
}
