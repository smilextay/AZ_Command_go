package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/smilextay/az_command_go/src/internal/word"
)

const strcutTpl = `type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
		{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
	{{end}}}
	func (model {{.TableName | ToCamelCase}}) TableName() string {
		return "{{.TableName}}"
	}`

type (
	StructTemplate struct {
		strcutTpl string
	}

	StructColumn struct {
		Name    string
		Type    string
		Tag     string
		Comment string
	}

	StructTemplateDB struct {
		TableName string
		Columns   []*StructColumn
	}
)

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}

func (t *StructTemplate) AssemblyColumns(cls []*TableColumn) []*StructColumn {

	scs := []*StructColumn{}
	for _, cl := range cls {
		scs = append(scs, &StructColumn{
			Name:    cl.ColumnName,
			Type:    DBTypeToStructType[cl.DataType],
			Tag:     fmt.Sprintf("`json:%s`", cl.ColumnName),
			Comment: cl.ColumnComment,
		})
	}
	return scs
}

func (t *StructTemplate) Generate(tableName string, columns []*StructColumn) error {

	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{"ToCamelCase": word.UnderScoreToUpperCameCase}).Parse(t.strcutTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   columns,
	}
	return tpl.Execute(os.Stdout, tplDB)
}
