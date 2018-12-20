// Package database contains implementation of database service.
// Any database service should be implemented here.
package database

import (
  "context"
  // "database/sql"
  "errors"
  "fmt"
  // "strings"
  // "time"
  // "strconv"

  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  "github.com/condrowiyono/sakudompet/saku"
  "github.com/condrowiyono/sakudompet/pass"


)

type MySQL struct {
  gormDb    *gorm.DB
}

type Option struct {
  User     string
  Password string
  Host     string
  Port     string
  Database string
  Charset  string
}
// NewMySQL returns a pointer of MySQL instance and error.
func NewMySQL(opt Option) (*MySQL, error) {
  dbUrl := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True", opt.User, opt.Password, opt.Host, opt.Port, opt.Database, opt.Charset)

  gormDb, err := gorm.Open("mysql", dbUrl)

  if err != nil {
    return nil, err
  }

  //Migrate things
  //gormDb.DropTableIfExists(&saku.Debit{},&pass.Pass{})
  gormDb.AutoMigrate(&saku.Debit{},&pass.Pass{})
  
  return &MySQL{gormDb : gormDb}, nil
}

func (m *MySQL) FindDebits(ctx context.Context) ([]saku.Debit, error) {
  select {
  case <- ctx.Done():
    return []saku.Debit{}, errors.New("Timeout")
  default:
  }
  
  var debit []saku.Debit

  result := m.gormDb.Find(&debit)
  
  return debit, result.Error
}