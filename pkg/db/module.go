package db

import "database/sql"

type StatusTable struct {
	Key   string
	Value string
}

type SlaveStatus struct {
	SlaveIoState                  string
	MasterHost                    string
	MasterUser                    string
	MasterPort                    string
	ConnectRetry                  string
	MasterLogFile                 string
	ReadMasterLogPos              string
	RelayLogFile                  string
	RelayLogPos                   string
	RelayMasterLogFile            string
	SlaveIoRunning                string
	SlaveSqlRunning               string
	ReplicateDoDB                 string
	ReplicateIgnoreDB             string
	ReplicateDoTable              string
	ReplicateIgnoreTable          string
	ReplicateWildDoTable          string
	ReplicateWildIgnoreTable      string
	LastErrno                     string
	LastError                     string
	SkipCounter                   string
	Exec_Master_Log_Pos           string
	Relay_Log_Space               string
	Until_Condition               string
	Until_Log_File                string
	Until_Log_Pos                 string
	Master_SSL_Allowed            string
	Master_SSL_CA_File            string
	Master_SSL_CA_Path            string
	Master_SSL_Cert               string
	Master_SSL_Cipher             string
	Master_SSL_Key                string
	Seconds_Behind_Master         string
	Master_SSL_Verify_Server_Cert string
	Last_IO_Errno                 string
	Last_IO_Error                 string
	Last_SQL_Errno                string
	Last_SQL_Error                string
	Replicate_Ignore_Server_Ids   string
	Master_Server_Id              string
	Master_UUID                   string
	Master_Info_File              string
	SQL_Delay                     string
	SqlRemainingDelay             sql.NullString
	Slave_SQL_Running_State       string
	Master_Retry_Count            string
	Master_Bind                   string
	Last_IO_Error_Timestamp       string
	Last_SQL_Error_Timestamp      string
	Master_SSL_Crl                string
	Master_SSL_Crlpath            string
	Retrieved_Gtid_Set            string
	Executed_Gtid_Set             string
	Auto_Position                 string
	Replicate_Rewrite_DB          string
	Channel_Name                  string
	Master_TLS_Version            string
	Master_public_key_path        string
	Get_master_public_key         string
	Network_Namespace             string
}
