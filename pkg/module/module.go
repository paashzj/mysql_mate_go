package module

import "database/sql"

type StatusTable struct {
	Key   string
	Value string
}

type SlaveStatus struct {
	SlaveIoState              string
	MasterHost                string
	MasterUser                string
	MasterPort                string
	ConnectRetry              string
	MasterLogFile             string
	ReadMasterLogPos          string
	RelayLogFile              string
	RelayLogPos               string
	RelayMasterLogFile        string
	SlaveIoRunning            string
	SlaveSqlRunning           string
	ReplicateDoDB             string
	ReplicateIgnoreDB         string
	ReplicateDoTable          string
	ReplicateIgnoreTable      string
	ReplicateWildDoTable      string
	ReplicateWildIgnoreTable  string
	LastErrno                 string
	LastError                 string
	SkipCounter               string
	ExecMasterLogPos          string
	RelayLogSpace             string
	UntilCondition            string
	UntilLogFile              string
	UntilLogPos               string
	MasterSslAllowed          string
	MasterSslCaFile           string
	MasterSslCaPath           string
	MasterSslCert             string
	MasterSslCipher           string
	MasterSslKey              string
	SecondsBehindMaster       string
	MasterSslVerifyServerCert string
	LastIoErrno               string
	LastIoError               string
	LastSqlErrno              string
	LastSqlError              string
	ReplicateIgnoreServerIds  string
	MasterServerId            string
	MasterUuid                string
	MasterInfoFile            string
	SqlDelay                  string
	SqlRemainingDelay         sql.NullString
	SlaveSqlRunningState      string
	MasterRetryCount          string
	MasterBind                string
	LastIoErrorTimestamp      string
	LastSqlErrorTimestamp     string
	MasterSslCrl              string
	MasterSslCrlpath          string
	RetrievedGtidSet          string
	ExecutedGtidSet           string
	AutoPosition              string
	ReplicateRewriteDb        string
	ChannelName               string
	MasterTlsVersion          string
	MasterPublicKeyPath       string
	GetMasterPublicKey        string
	NetworkNamespace          string
}