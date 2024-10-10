package logger

import (
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
)

type RequestLog struct {
	Method  string
	Path    string
	User    string
	Latency string
}

type DbLogger struct {
	Mysql *MySQLLogger
}

type MySQLLogger struct {
	Appli bool
}

func RecordAppliEvent(currentUser aadu.AppliUserLogin, entry RequestLog) {
	mysqlLogger := &MySQLLogger{Appli: true}
	dbLogger := DbLogger{Mysql: mysqlLogger}

	RecordEvent(&currentUser, dbLogger, entry)
}

func RecordEvent(currentUser *aadu.AppliUserLogin, dbLogger DbLogger, entry RequestLog) error {
	err := dbLogger.Mysql.RecordEvent(currentUser, entry)
	if err != nil {
		return err
	}

	return nil
}

func (l *MySQLLogger) RecordEvent(currentUser *aadu.AppliUserLogin, entry RequestLog) error {
	/* TODO: LOG SYSTEM :
	entryRawMessage, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	logEventDesc := entryRawMessage

	logParams := mysqlappli.CreateLogParams{
		LogEventDesc: logEventDesc,
		LogEventCode: mysqlappli.NullString{NullString: sql.NullString{String: "1", Valid: true}},
		LogEventDate: mysqlappli.NullTime{NullTime: sql.NullTime{Time: time.Now(), Valid: true}},
	}
	if currentUser != nil {
		_, err = la.InsertLog(currentUser.ClientId, logParams)
		if err != nil {
			return err
		}
	} */

	return nil
}
