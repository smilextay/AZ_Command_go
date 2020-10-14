package cmd

import (
	"log"

	"github.com/smilextay/az_command_go/src/internal/sql2struct"
	"github.com/spf13/cobra"
)

var (
	dbType, host, username, password, charset, dbname, tablename string

	sqlCmd = &cobra.Command{
		Use:   "sql",
		Short: "sql转换和处理",
		Long:  "sql转换和处理",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	sql2StructCmd = &cobra.Command{
		Use:   "struct",
		Short: "sql 生成 go 结构体",
		Long:  "根据mysql 表结构自动生成go结构体",
		Run: func(cmd *cobra.Command, arg []string) {

			dbinfo := &sql2struct.DBInfo{
				DBType:   dbType,
				Host:     host,
				UserName: username,
				Password: password,
				Charset:  charset,
			}
			dbModel := sql2struct.NewDBModel(dbinfo)

			err := dbModel.Connect()
			if err != nil {
				log.Fatalf("dbModel.Connect err:%v", err)
			}
			cls, err := dbModel.GetColumns(dbname, tablename)
			if err != nil {
				log.Fatalf("GetClolumns err:%v", err)
			}
			template := sql2struct.NewStructTemplate()

			tcls := template.AssemblyColumns(cls)
			err = template.Generate(tablename, tcls)
			if err != nil {
				log.Fatalf("template.Generate err:%v", err)
			}
		},
	}
)

func init() {
	sqlCmd.AddCommand(sql2StructCmd)
	sql2StructCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据的账号")
	sql2StructCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据的密码")
	sql2StructCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库服务器")
	sql2StructCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库字符编码")
	sql2StructCmd.Flags().StringVarP(&dbType, "dbType", "", "mysql", "请输入数据库类型")
	sql2StructCmd.Flags().StringVarP(&dbname, "db", "", "", "请输入数据库名称")
	sql2StructCmd.Flags().StringVarP(&tablename, "table", "", "", "请输入表名")
}
