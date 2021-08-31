package db

import (
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
	"mysql_mate_go/pkg/module"
)

func GetPerformance() map[string]string {
	return getStatus("SELECT VARIABLE_NAME,VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME in (\"Innodb_buffer_pool_pages_total\", \"Innodb_buffer_pool_pages_free\")")
}

// GetGlobalStatus 获取 mysql status的状态
func GetGlobalStatus() map[string]string {
	return getStatus("SHOW GLOBAL STATUS")
}

func SelectOne() bool {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		logs.Error("Error select one when opening DB", err)
		return false
	}
	defer db.Close()
	_, err = db.Query("SELECT 1")
	if err != nil {
		logs.Error("Error select one when opening DB", err)
		return false
	}
	return true
}

// GetSlaveStatus 获取 slave status的状态
func GetSlaveStatus() module.SlaveStatus {
	var slaveStatus module.SlaveStatus
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		logs.Error("Error %s when opening DB\n", err)
		return slaveStatus
	}
	logs.Debug("open db success show slave status")
	defer db.Close()
	results, err := db.Query("SHOW SLAVE STATUS")
	if err != nil {
		logs.Error("Error %s when querying DB\n", err)
		return slaveStatus
	}
	defer results.Close()
	for results.Next() {
		var slaveStatus module.SlaveStatus
		// for each row, scan the result into our statusDto composite object
		err = results.Scan(&slaveStatus.SlaveIoState, &slaveStatus.MasterHost, &slaveStatus.MasterUser,
			&slaveStatus.MasterPort, &slaveStatus.ConnectRetry, &slaveStatus.MasterLogFile, &slaveStatus.ReadMasterLogPos,
			&slaveStatus.RelayLogFile, &slaveStatus.RelayLogPos, &slaveStatus.RelayMasterLogFile, &slaveStatus.SlaveIoRunning,
			&slaveStatus.SlaveSqlRunning,
			&slaveStatus.ReplicateDoDB,
			&slaveStatus.ReplicateIgnoreDB,
			&slaveStatus.ReplicateDoTable,
			&slaveStatus.ReplicateIgnoreTable,
			&slaveStatus.ReplicateWildDoTable,
			&slaveStatus.ReplicateWildIgnoreTable,
			&slaveStatus.LastErrno,
			&slaveStatus.LastError,
			&slaveStatus.SkipCounter,
			&slaveStatus.ExecMasterLogPos,
			&slaveStatus.RelayLogSpace,
			&slaveStatus.UntilCondition,
			&slaveStatus.UntilLogFile,
			&slaveStatus.UntilLogPos,
			&slaveStatus.MasterSslAllowed,
			&slaveStatus.MasterSslCaFile,
			&slaveStatus.MasterSslCaPath,
			&slaveStatus.MasterSslCert,
			&slaveStatus.MasterSslCipher,
			&slaveStatus.MasterSslKey,
			&slaveStatus.SecondsBehindMaster,
			&slaveStatus.MasterSslVerifyServerCert,
			&slaveStatus.LastIoErrno,
			&slaveStatus.LastIoError,
			&slaveStatus.LastSqlErrno,
			&slaveStatus.LastSqlError,
			&slaveStatus.ReplicateIgnoreServerIds,
			&slaveStatus.MasterServerId,
			&slaveStatus.MasterUuid,
			&slaveStatus.MasterInfoFile,
			&slaveStatus.SqlDelay,
			&slaveStatus.SqlRemainingDelay,
			&slaveStatus.SlaveSqlRunningState,
			&slaveStatus.MasterRetryCount,
			&slaveStatus.MasterBind,
			&slaveStatus.LastIoErrorTimestamp,
			&slaveStatus.LastSqlErrorTimestamp,
			&slaveStatus.MasterSslCrl,
			&slaveStatus.MasterSslCrlpath,
			&slaveStatus.RetrievedGtidSet,
			&slaveStatus.ExecutedGtidSet,
			&slaveStatus.AutoPosition,
			&slaveStatus.ReplicateRewriteDb,
			&slaveStatus.ChannelName,
			&slaveStatus.MasterTlsVersion,
			&slaveStatus.MasterPublicKeyPath,
			&slaveStatus.GetMasterPublicKey,
			&slaveStatus.NetworkNamespace)
		if err != nil {
			logs.Error("Error %s when scanning result ", err)
			return slaveStatus
		}
	}
	return slaveStatus
}

func getStatus(sqlStr string) map[string]string {
	result := make(map[string]string)
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		logs.Error("Error %s when opening DB", sqlStr, err)
		return result
	}
	logs.Debug("open db success sql is ", sqlStr)
	defer db.Close()
	results, err := db.Query(sqlStr)
	if err != nil {
		logs.Error("Error %s when querying DB", sqlStr, err)
		return result
	}
	defer results.Close()
	for results.Next() {
		var statusDto module.StatusTable
		// for each row, scan the result into our statusDto composite object
		err = results.Scan(&statusDto.Key, &statusDto.Value)
		if err != nil {
			logs.Error("sql %s Error %s when scaning result ", sqlStr, err)
			return result
		}
		// and then print out the statusDto's Name attribute
		result[statusDto.Key] = statusDto.Value
	}
	return result
}
