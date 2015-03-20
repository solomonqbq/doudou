package db

import (
	"database/sql"
	"fmt"
	"github.com/dmotylev/goproperties"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	Global_properties properties.Properties
	dataSource        *sql.DB
)

type Meta struct {
	Table_Name     string
	Column_name    string
	Data_type      string
	Column_comment string
}

const (
	QUERY_MATA       = "select `column_name`,`data_type`,`column_comment` from columns where  table_schema=? and table_name=?"
	QUERY_ALL_TABLES = "SELECT `TABLE_NAME` FROM `TABLES` WHERE `TABLE_SCHEMA`=?"
)

func InitAllDS(prop properties.Properties) {
	var db_url_template = "%s:%s@tcp(%s:%s)/%s?charset=%s"
	account := prop.String("db.account", "account")
	password := prop.String("db.password", "password")
	ip := prop.String("db.ip", "127.0.0.1")
	port := prop.String("db.port", "3306")
	//	schema := prop.String("db.schema", "")
	charset := prop.String("db.charset", "utf8")
	fmt.Println("db.ip:", ip)
	db_url := fmt.Sprintf(db_url_template, account, password, ip, port, "information_schema", charset)
	InitDS(db_url, 1, 1)
}

func InitDS(url string, maxIdle int, maxOpen int) (err error) {
	db, err := sql.Open("mysql", url)
	if maxIdle > 0 {
		db.SetMaxIdleConns(maxIdle)
	} else {
		db.SetMaxIdleConns(25)
	}
	if maxOpen > 0 {
		db.SetMaxOpenConns(maxOpen)
	} else {
		db.SetMaxOpenConns(50)
	}

	if err != nil {
		return
	}
	err = db.Ping()

	if err != nil {
		fmt.Println(err)
		return
	}

	dataSource = db
	return nil
}

func QueryAllTables(schema string, filter string) (tables []string, err error) {
	result, err := dataSource.Query(QUERY_ALL_TABLES, schema)
	if err != nil {
		return nil, err
	}
	tables = make([]string, 0)
	for result.Next() {
		tbn := ""
		result.Scan(&tbn)
		if filter == "*" {
			tables = append(tables, tbn)
		} else if strings.Contains(filter, tbn) {
			tables = append(tables, tbn)
		}
	}
	return
}

func QueryMetaInfo(schema string, table string) (table_mata []*Meta, err error) {
	result, err := dataSource.Query(QUERY_MATA, schema, table)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	table_mata = make([]*Meta, 0)
	for result.Next() {
		meta := new(Meta)
		meta.Table_Name = table
		err = result.Scan(&meta.Column_name, &meta.Data_type, &meta.Column_comment)
		if err != nil {
			return nil, err
		}
		table_mata = append(table_mata, meta)
	}
	return
}

func GenModel() error {
	schema := Global_properties.String("db.schema", "")
	filter := Global_properties.String("db.tables", "*")
	tables, err := QueryAllTables(schema, filter)
	if err != nil {
		return err
	}
	all_table_meta := make(map[string][]*Meta, 0)
	for _, tbn := range tables {
		meta, err := QueryMetaInfo(schema, tbn)
		if err != nil {
			return err
		}
		all_table_meta[tbn] = meta
	}
	output_file := Global_properties.String("output_dir", "./output/db_model.go")
	p_name := Global_properties.String("model_package", "model")
	dir, _ := os.Getwd()

	err = WriteFile(filepath.Join(dir, output_file), p_name, all_table_meta)
	if err != nil {
		return err
	}
	return nil
}

func transType(db_type string) string {
	switch strings.ToLower(db_type) {
	case "varchar":
		return "string"
	case "tinyint":
		return "int32"
	case "text":
		return "string"
	case "date":
		return "string"
	case "smallint":
		return "int32"
	case "mediumint":
		return "int64"
	case "int":
		return "int32"
	case "bigint":
		return "int64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "decimal":
		return "float64"
	case "datetime":
		return "string"
	case "timestamp":
		return "string"
	case "time":
		return "string"
	case "year":
		return "string"
	case "char":
		return "string"
	case "tinyblob":
		return "string"
	case "tinytext":
		return "string"
	case "blob":
		return "string"
	case "mediumblob":
		return "string"
	case "mediumtext":
		return "string"
	case "longblob":
		return "string"
	case "longtext":
		return "string"
	case "enum":
		return "string"
	case "set":
		return "string"
	case "bool":
		return "string"
	case "binary":
		return "string"
	case "varbinary":
		return "string"
	default:
		return "string"
	}
}

func WriteFile(desFile string, p_name string, all_table_meta map[string][]*Meta) error {
	dir, _ := filepath.Split(desFile)
	_, err := os.Stat(dir)
	if err != nil {
		fmt.Println("目录不存在，准备创建%s", dir)
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			return err
		}
	}
	str := "package " + p_name + "\n\n\n"
	for table_name, meta := range all_table_meta {
		str = str + "type " + FirstLetterToUpper(table_name) + " struct{\n"
		for _, m := range meta {
			commet := ""
			if m.Column_comment != commet {
				commet = "\t//" + m.Column_comment
			}
			str = str + "\t" + FirstLetterToUpper(m.Column_name) + "\t" + transType(m.Data_type) + commet + "\n"
		}
		str += "}\n"
	}
	fmt.Println(str)
	err = ioutil.WriteFile(desFile, []byte(str), 0666)
	return err
}

func FirstLetterToUpper(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}
