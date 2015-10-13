package dh

import (
	"database/sql"
	"github.com/astaxie/beego/validation"
	"reflect"
	"appworkz/base/api"
	conn "appworkz/base/configdb"
	"strings"
	"upper.io/db"
	// "upper.io/db/mysql"
	"upper.io/db/util/sqlutil"
	// "fmt"
)

// @ descripcion
// DataHandler efectua todas las operaciones CRUD a nuestra base de datos
// Definimos los campos de nuestra tabla en una estructura
// los atributos deben ser publicos es por eso que estan en mayusculas
// tambien tenemos tags db que se refiere al campo de nuestra tabla en
// base de datos y el atributo json se utiliza para nombrar cada llave
// al ser convertida el struct a un objeto tipo json

type Dh struct {
	Model interface{}
}

// La variable Collection contiene la conexion a nuestra tabla
var Collection db.Collection

// NewDh
// inicializamos nuestra conexiona nuestra tabla
// obtenemos el nombre de la tabla por medio de reflections
func NewDh(model interface{}) Dh {
	self := Dh{Model: model}
	val := reflect.ValueOf(model)
	table := self.GetTableName(val)
	var err error
	Collection, err = conn.Collection(table)
	if err != nil {
		panic(err.Error())
	}
	return self
}

func (self *Dh) SqlDriver() *sql.DB {
	return conn.Sess().Driver().(*sql.DB)
}

// Nombre: Save()
// Este metodo Guarda una estructura de datos en nuestra base de datos
func (self *Dh) Save() (int64, api.Err) {
	var errs api.Err
	if exist := self.Exist(); exist {
		 _, errup := self.Update()
		return 0, errup
	}

	id, err := Collection.Append(self.Model)

	if err != nil {
		errs.Push(api.Msg{
			Field: "DataBase insert error:",
			Error: err.Error(),
		})
		return 0, errs
	}

	idInsert := id.(int64)
	return idInsert, errs
}

// FindById
// Podemos hacer busquedas por PK Id este metodo
// recibe como parametro la pk de la tabla y trae todos sus
// elementos.
func (self *Dh) FindById(id int64) (interface{}, api.Err) {
	var errs api.Err
	var res db.Result
	res = Collection.Find(db.Cond{"id": id})
	count, _ := res.Count()

	if count == 0 {
		errs.Push(api.Msg{
			Field: "Not found",
			Error: api.ErrNoSuchRow,
		})
		return nil, errs
	}

	res.One(self.Model)

	return self.Model, errs
}

func (self *Dh) FindByFieldValue(field string, value interface{}) (interface{}, api.Err) {
	var errs api.Err
	var res db.Result
	res = Collection.Find(db.Cond{field: value})
	count, _ := res.Count()

	if count == 0 {
		errs.Push(api.Msg{
			Field: "Not found",
			Error: api.ErrNoSuchRow,
		})
		return nil, errs
	}
	res.One(self.Model)
	return self.Model, errs
}

func (self *Dh) FindByConditions(conditions interface{}) (uint64, api.Err) {
	var errs api.Err
	var res db.Result
	res = Collection.Find(conditions)
	count, _ := res.Count()
	if count == 0 {
		errs.Push(api.Msg{
			Field: "Not found",
			Error: api.ErrNoSuchRow,
		})
		return count, errs
	}
	res.One(self.Model)
	return count, errs
}

func (self *Dh) FindAllByConditions(bar interface{}, conditions interface{}) api.Err {
	var res db.Result
	var errs api.Err
	res = Collection.Find(conditions)
	err := res.All(bar)
	if err != nil {
		errs.Push(api.Msg{
			Field: "Not found",
			Error: api.ErrNoSuchRow,
		})
	}
	return errs
}

func (self *Dh) FindAllByConditionsSortBy(bar interface{}, conditions interface{},sort string) api.Err {
	var res db.Result
	var errs api.Err
	res = Collection.Find(conditions).Sort(sort)
	err := res.All(bar)
	if err != nil {
		errs.Push(api.Msg{
			Field: "Not found",
			Error: api.ErrNoSuchRow,
		})
	}
	return errs
}

// func (self *Dh) FindAll(bar interface{}) interface{} {
// 	var res db.Result
// 	t := reflect.TypeOf(bar)
// 	data := reflect.New(t).Interface()

// 	res = Collection.Find()
// 	count, _ := res.Count()

// 	if count == 0 {
// 		return nil
// 	}
// 	res.All(data)
// 	return data
// }

func (self *Dh) FindAll(bar interface{}) api.Err {
	var res db.Result
	var errs api.Err
	res = Collection.Find()
	count, _ := res.Count()
	if count == 0 {
		errs.Push(api.Msg{
			Field: "Select All",
			Error: api.ErrSelectAll,
		})
	}
	res.All(bar)
	return errs
}

func (self *Dh) FindAllSortBy(bar interface{},sort string) api.Err {
	var res db.Result
	var errs api.Err
	res = Collection.Find().Sort(sort)
	count, _ := res.Count()
	if count == 0 {
		errs.Push(api.Msg{
			Field: "Select All",
			Error: api.ErrSelectAll,
		})
	}
	res.All(bar)
	return errs
}

func (self *Dh) FindAllCountFieldValue(field string, value interface{}) uint64 {
	var res db.Result
	res = Collection.Find(db.Cond{field: value})
	count, _ := res.Count()
	return count
}

func (self *Dh) CountByConditions(conditions interface{}) uint64 {
	var res db.Result
	res = Collection.Find(conditions)
	count, _ := res.Count()
	return count
}

func (self *Dh) FindAllCount() uint64 {
	var res db.Result
	res = Collection.Find()
	count, _ := res.Count()
	return count
}

func (self *Dh) Count(query string) int64 {
	var errs api.Err
	drv := self.SqlDriver()
	rows, err := drv.Query(query)
	if err != nil {
		errs.Push(api.Msg{
			Field: "Database select All  error: ",
			Error: err.Error(),
		})
		return 0
	}

	type Total struct {
		Count int64 `db:"count"`
	}

	total := []Total{}
	// Mapping to an array.
	if err = sqlutil.FetchRows(rows, &total); err != nil {
		errs.Push(api.Msg{
			Field: "Database select All  error: ",
			Error: err.Error(),
		})
	}
	
	defer rows.Close()
	return total[0].Count
}

// Nombre: Exist
// Verifica que exista un elemento en nuestra base de datos
// La busca por su PK
func (self *Dh) Exist() bool {


	s := reflect.ValueOf(self.Model).Elem().Field(0)
	id := s.Interface().(int64)
	// change the field id 
	// fmt.Println(s.Interface().(int64))
	// fmt.Println(reflect.ValueOf(self.Model).Elem().FieldByName("Id"),"data")
	// id := reflect.ValueOf(self.Model).Elem().FieldByName("Id").Interface().(int64)
	var res db.Result
	res = Collection.Find(db.Cond{"id": id})
	count, _ := res.Count()
	if count == 0 {
		return false
	}
	return true
}

// Nombre: Update
// Este metodo se encarga de actualizar un registro
// si ya fue creado en nuestra base de datos es llamado por el
// metodo save
func (self *Dh) Update() (interface{}, api.Err) {

	var errs api.Err
	var res db.Result

	id := reflect.ValueOf(self.Model).Elem().FieldByName("Id").Interface().(int64)
	res = Collection.Find(db.Cond{"id": id})
	err := res.Update(self.Model)

	if err != nil {
		errs.Push(api.Msg{
			Field: "Save",
			Error: err.Error(),
		})
		return nil, errs
	}

	dataUpdate, errr := self.FindById(id)
	return dataUpdate, errr
}

// Nombre: Delete
// Elimina para siempre de la base de datos

func (self *Dh) Delete() api.Err {

	var res db.Result
	var errs api.Err

	id := reflect.ValueOf(self.Model).Elem().FieldByName("Id").Interface().(int64)
	res = Collection.Find(db.Cond{"id": id})
	exist, err := res.Count()

	if exist == 0 {
		errs.Push(api.Msg{
			Field: "Delete: ",
			Error: api.ErrNoSuchRow,
		})
		return errs
	}

	err = res.Remove()

	if err == nil {
		return errs
	} else {
		errs.Push(api.Msg{
			Field: "Database delete  error: ",
			Error: err.Error(),
		})
	}

	return errs
}

func (self *Dh) DeleteLogic() api.Err {

	var errs api.Err
	var res db.Result

	id := reflect.ValueOf(self.Model).Elem().FieldByName("Id").Interface().(int64)
	res = Collection.Find(db.Cond{"id": id})
	res.One(self.Model)
	reflect.ValueOf(self.Model).Elem().FieldByName("Status").SetBool(false)
	// Estatus := reflect.ValueOf(self.Model).Elem().FieldByName("Estatus").Interface().(bool)
	err := res.Update(self.Model)

	if err != nil {
		errs.Push(api.Msg{
			Field: "Save",
			Error: err.Error(),
		})
		return errs
	}

	return errs
}

func (self *Dh) RemoveForeverByConditions(conditions interface{}) api.Err {
	var res db.Result
	var errs api.Err
	res = Collection.Find(conditions)
	exist, err := res.Count()

	if exist == 0 {
		errs.Push(api.Msg{
			Field: "Delete: ",
			Error: api.ErrNoSuchRow,
		})
		return errs
	}

	err = res.Remove()

	if err == nil {
		return errs
	} else {
		errs.Push(api.Msg{
			Field: "Database delete  error: ",
			Error: err.Error(),
		})
	}
	return errs
}

func (self *Dh) Validate() *api.Err {
	errors := api.NewError()
	valid := validation.Validation{}
	b, _ := valid.Valid(self.Model)
	if !b {
		for _, err := range valid.Errors {
			errors.SetErr(err.Key, err.Message, 0)
		}
	}
	return errors
}

// GetTableName
// Obtiene el nombre de la tabla por medio del modelo
// registrado el metodo TableName existe declarado en
// el modelo y nos retorna un string con el nombre de
// la tabla

func (self *Dh) GetTableName(val reflect.Value) string {
	ind := reflect.Indirect(val)
	fun := val.MethodByName("TableName")
	if fun.IsValid() {
		vals := fun.Call([]reflect.Value{})
		if len(vals) > 0 {
			val := vals[0]
			if val.Kind() == reflect.String {
				return val.String()
			}
		}
	}
	return self.snakeString(ind.Type().Name())
}

func (self *Dh) snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	for i, d := range s {
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		j = (d != '_')
		data = append(data, byte(d))
	}
	return strings.ToLower(string(data[:len(data)]))
}
