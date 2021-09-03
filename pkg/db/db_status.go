package db

import (
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
	"mysql_mate_go/pkg/module"
)

func GetPerformance() map[string]string {
	return getKvStr("SELECT VARIABLE_NAME,VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME in (\"Innodb_buffer_pool_pages_total\", \"Innodb_buffer_pool_pages_free\")")
}

// GetGlobalStatus 获取 mysql status的状态
func GetGlobalStatus() map[string]string {
	return getKvStr("SHOW GLOBAL STATUS")
}

func SelectOne() bool {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		logs.Error("Error select one when opening DB", err)
		return false
	}
	defer db.Close()
	result, err := db.Query("SELECT 1")
	defer result.Close()
	if err != nil {
		logs.Error("Error select one when opening DB", err)
		return false
	}
	return true
}

func getKvStr(sqlStr string) map[string]string {
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
		var kv module.KvStr
		// for each row, scan the result into our kv composite object
		err = results.Scan(&kv.Key, &kv.Value)
		if err != nil {
			logs.Error("sql %s Error %s when scanning result ", sqlStr, err)
			return result
		}
		// and then print out the kv's Name attribute
		result[kv.Key] = kv.Value
	}
	return result
}
