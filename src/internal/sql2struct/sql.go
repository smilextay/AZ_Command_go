package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type (
	DBInfo struct {
		DBType   string
		Host     string
		UserName string
		Password string
		Charset  string
	}

	dBModel struct {
		DBEngin *sql.DB
		DBInfo  *DBInfo
	}

	TableColumn struct {
		ColumnName    string
		ColumnType    string
		IsNullable    string
		DataType      string
		ColumnKye     string
		ColumnComment string
	}
)

var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func NewDBModel(info *DBInfo) *dBModel {
	return &dBModel{
		DBInfo: info,
	}
}

func (m *dBModel) Connect() error {
	// switch m.DBInfo.DBType {

	// case "sql":

	// }
	var err error
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=true&loc=Local", m.DBInfo.UserName, m.DBInfo.Password, m.DBInfo.Host, m.DBInfo.Charset)
	m.DBEngin, err = sql.Open(m.DBInfo.DBType, connectionStr)
	return err
}

func (m *dBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {

	queryString := `SELECT 
COLUMN_NAME,DATA_TYPE,COLUMN_KEY,IS_NULLABLE,COLUMN_TYPE,COLUMN_COMMENT
FROM COLUMNS WHERE TABLE_SCHEMA=? AND TABLE_NAME=?
`
	result, err := m.DBEngin.Query(queryString, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("没有找到表数据")
	}

	defer result.Close()
	var cls []*TableColumn

	for result.Next() {
		cl := &TableColumn{}
		err := result.Scan(&cl.ColumnName, &cl.DataType, &cl.ColumnKye, &cl.IsNullable, &cl.ColumnType, &cl.ColumnComment)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)
	}
	return cls, nil
}
