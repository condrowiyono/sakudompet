package sakudompet

import (
  "context"
  "github.com/condrowiyono/sakudompet/saku"
)

type SakuDompet struct {
  database                  DatabaseInterface
}

type DatabaseInterface interface {
  FindDebits(context.Context) ([]saku.Debit, error)
}

func NewSakuDompet(db DatabaseInterface) *SakuDompet {
  return &SakuDompet{
    database:                   db,
  }
}

func (s *SakuDompet) GetDebits(ctx context.Context) ([]saku.Debit, error) {
  return s.database.FindDebits(ctx)
}