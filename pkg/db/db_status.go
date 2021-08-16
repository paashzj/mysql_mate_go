package db

import (
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
)

func GetPerformance() map[string]string {
	return getStatus("SELECT VARIABLE_NAME,VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME in (\"Innodb_buffer_pool_pages_total\", \"Innodb_buffer_pool_pages_free\")")
}

/**
获取 mysql status的状态
*/
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

/**
获取 slave status的状态
*/
func GetSlaveStatus() SlaveStatus {
	var slaveStatus SlaveStatus
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
		var slaveStatus SlaveStatus
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
			&slaveStatus.Exec_Master_Log_Pos,
			&slaveStatus.Relay_Log_Space,
			&slaveStatus.Until_Condition,
			&slaveStatus.Until_Log_File,
			&slaveStatus.Until_Log_Pos,
			&slaveStatus.Master_SSL_Allowed,
			&slaveStatus.Master_SSL_CA_File,
			&slaveStatus.Master_SSL_CA_Path,
			&slaveStatus.Master_SSL_Cert,
			&slaveStatus.Master_SSL_Cipher,
			&slaveStatus.Master_SSL_Key,
			&slaveStatus.Seconds_Behind_Master,
			&slaveStatus.Master_SSL_Verify_Server_Cert,
			&slaveStatus.Last_IO_Errno,
			&slaveStatus.Last_IO_Error,
			&slaveStatus.Last_SQL_Errno,
			&slaveStatus.Last_SQL_Error,
			&slaveStatus.Replicate_Ignore_Server_Ids,
			&slaveStatus.Master_Server_Id,
			&slaveStatus.Master_UUID,
			&slaveStatus.Master_Info_File,
			&slaveStatus.SQL_Delay,
			&slaveStatus.SqlRemainingDelay,
			&slaveStatus.Slave_SQL_Running_State,
			&slaveStatus.Master_Retry_Count,
			&slaveStatus.Master_Bind,
			&slaveStatus.Last_IO_Error_Timestamp,
			&slaveStatus.Last_SQL_Error_Timestamp,
			&slaveStatus.Master_SSL_Crl,
			&slaveStatus.Master_SSL_Crlpath,
			&slaveStatus.Retrieved_Gtid_Set,
			&slaveStatus.Executed_Gtid_Set,
			&slaveStatus.Auto_Position,
			&slaveStatus.Replicate_Rewrite_DB,
			&slaveStatus.Channel_Name,
			&slaveStatus.Master_TLS_Version,
			&slaveStatus.Master_public_key_path,
			&slaveStatus.Get_master_public_key,
			&slaveStatus.Network_Namespace)
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
		var statusDto StatusTable
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
