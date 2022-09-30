package gorm

// go run gen.go
import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// 参考 https://gitee.com/gorm/gen#%E7%94%9F%E6%88%90model
func TestGenModel(t *testing.T) {

	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "./persistent", // 这个是 crud 方法的目录，model 目录已经固定为 model 了
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		// if you want generate field with unsigned integer type, set FieldSignable true
		FieldSignable: true,
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
		//if you need unit tests for query code, set WithUnitTest true
		//WithUnitTest: true,
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	db, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)

	// 生成model 和 crud 方法，但是生成的crud方法我们一般不用
	//g.ApplyBasic(g.GenerateModel("person"))

	// 只生成 model 就够了
	g.GenerateModel("person")

	// 想自定义某个表生成特性，比如struct的名称/字段类型/tag等，可以指定opt，例如g.GenerateModel("company",gen.FieldIgnore("address")), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address"))
	// g.ApplyBasic(ReconciliationTxnTab{}, g.GenerateModel("reconciliation_txn_tab"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))

	// apply diy interfaces on structs or table models
	// 如果想给某些表或者model生成自定义方法，可以用ApplyInterface，第一个参数是方法接口，可以参考DIY部分文档定义
	// g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	// execute the action of code generation
	g.Execute()
}
