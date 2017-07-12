package models

import (
	// "github.com/jinzhu/gorm"
	// "reflect"
	"sns/util/snslog"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
type BaseModelWithId struct {
	BaseModel
	ID uint `gorm:"primary_key"`
}

func (this BaseModel) InsertOrUpdate(v interface{}) (err error) {
	err = GetDB().Create(v).Error
	if err != nil {
		err = nil
		err = GetDB().Model(v).Where(GetPrimaryKey(v)).Update(v).Error
	}
	return
}
func (this BaseModel) All(v interface{}) (err error) {
	err = GetDB().Find(v).Error
	return
}
func (this BaseModel) DeleteByStruct(user interface{}) (err error) {
	keys := GetPrimaryKey(user)
	query := "delete from " + GetTableName(user) + " where" + map2query(keys)
	snslog.I(query)
	err = GetDB().Exec(query).Error
	return
}
func (this BaseModel) Delete(user interface{}, query string, args ...interface{}) (err error) {
	query = "delete from " + GetTableName(user) + " where" + query
	snslog.I(query)
	err = GetDB().Exec("delete from "+GetTableName(user)+" where"+query, args).Error
	return
}
func (this BaseModel) Query(users interface{}, query interface{}, args ...interface{}) (err error) {
	err = GetDB().Where(query, args).Find(users).Error
	return
}

func (this BaseModel) QueryByKey(out interface{}, query interface{}, args ...interface{}) (err error) {
	keys := GetPrimaryKey(query)
	if len(keys) == 0 {
		err = GetDB().First(out, query, args).Error
	} else {
		err = GetDB().Where(keys).First(out).Error
	}
	return
}
func GetTableName(v interface{}) string {
	scope := GetDB().NewScope(v)
	name := scope.TableName()
	snslog.I(name)
	return name
}
func GetPrimaryKey(v interface{}) map[string]interface{} {
	keys := make(map[string]interface{})
	scope := GetDB().NewScope(v)
	for _, value := range scope.Fields() {
		if value.StructField.IsPrimaryKey {
			keys[value.StructField.DBName] = value.Field.Interface()
		}
	}
	snslog.If("keys:%s,%s", v, keys)
	return keys
}
func map2query(m map[string]interface{}) (query string) {
	for key, value := range m {
		query += " " + key + " = \"" + value.(string) + "\" and"
	}
	if len(query) > 0 {
		query = query[:len(query)-4]
	}
	snslog.I(query)
	return
}

// func (this BaseModel) Exist(v interface{}) bool {
// 	var *value = *v
// 	err2 := GetDB().Where(v).Rows().Next()
// 	if err2 != nil {
// 		return false
// 	}
// 	return true
// }
