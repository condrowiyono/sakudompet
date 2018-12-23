// Package database contains implementation of database service.
// Any database service should be implemented here.
package database

import (
  "context"
  "errors"
  "fmt"

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
  // gormDb.DropTableIfExists(&saku.Debit{},&pass.Pass{})
  gormDb.AutoMigrate(&saku.Debit{},&pass.Pass{})
  
  return &MySQL{gormDb : gormDb}, nil
}

func (m *MySQL) FindPass(ctx context.Context, id uint) (pass.Pass, error) {
  select {
  case <- ctx.Done():
    return pass.Pass{}, errors.New("Timeout")
  default:
  }
  
  var pass pass.Pass

  result := m.gormDb.First(&pass,id)
  return pass, result.Error
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

//By Primary Key
func (m *MySQL) FindDebit(ctx context.Context, id uint) (saku.Debit, error) {
  select {
  case <- ctx.Done():
    return saku.Debit{}, errors.New("Timeout")
  default:
  }
  
  var debit saku.Debit
  result := m.gormDb.First(&debit,id)
  return debit, result.Error
}

func (m *MySQL) CreateDebit(ctx context.Context, debit saku.Debit) (saku.Debit, error) {
  select {
  case <- ctx.Done():
    return saku.Debit{}, errors.New("Timeout")
  default:
  }
  
  result := m.gormDb.Create(&debit)
  return debit, result.Error
}

func (m *MySQL) PutDebit(ctx context.Context, id uint, debit saku.Debit) (saku.Debit, error) {
  select {
  case <- ctx.Done():
    return saku.Debit{}, errors.New("Timeout")
  default:
  }
  var findDebit saku.Debit
  findResult := m.gormDb.First(&findDebit,id)

  if findResult.Error != nil {
    return saku.Debit{}, findResult.Error
  }

  updates := map[string]interface{} {
    "name":       debit.Name,
    "issued_by":  debit.IssuedBy,
    "number":     debit.Number,
    "balance":    debit.Balance,
    "updated_at": debit.UpdatedAt,
    "pass_id":    debit.PassId,
  }

  result := m.gormDb.Model(&findDebit).Updates(updates)

  return findDebit, result.Error
}

func (m *MySQL) DeleteDebit(ctx context.Context, id uint) (saku.Debit, error) {
  select {
  case <- ctx.Done():
    return saku.Debit{}, errors.New("Timeout")
  default:
  }
  
  var debit saku.Debit
  result := m.gormDb.Where("id = ?", id).Delete(&debit);

  return debit, result.Error
}

func (m *MySQL) CreatePass(ctx context.Context, passRequest pass.Pass) (pass.Pass, error) {
  select {
  case <- ctx.Done():
    return pass.Pass{}, errors.New("Timeout")
  default:
  }
  
  result := m.gormDb.Create(&passRequest)
  return passRequest, result.Error
}

func (m *MySQL) GetPasses(ctx context.Context) ([]pass.Pass, error) {
  select {
  case <- ctx.Done():
    return []pass.Pass{}, errors.New("Timeout")
  default:
  }
  
  var pass []pass.Pass

  result := m.gormDb.Find(&pass)
  return pass, result.Error
}

func (m *MySQL) DeletePass(ctx context.Context, id uint) (pass.Pass, error) {
  select {
  case <- ctx.Done():
    return pass.Pass{}, errors.New("Timeout")
  default:
  }
  
  var pass pass.Pass
  result := m.gormDb.Where("id = ?", id).Delete(&pass);

  return pass, result.Error
}

