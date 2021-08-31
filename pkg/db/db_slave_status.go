package db

import (
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
	"mysql_mate_go/pkg/module"
)

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
		err = results.Scan(&slaveStatus.SlaveIoState,
			&slaveStatus.MasterHost,
			&slaveStatus.MasterUser,
			&slaveStatus.MasterPort,
			&slaveStatus.ConnectRetry,
			&slaveStatus.MasterLogFile,
			&slaveStatus.ReadMasterLogPos,
			&slaveStatus.RelayLogFile,
			&slaveStatus.RelayLogPos,
			&slaveStatus.RelayMasterLogFile,
			&slaveStatus.SlaveIoRunning,
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
