package sakudompet

import (
	"context"
	"github.com/condrowiyono/sakudompet/saku"
	"github.com/condrowiyono/sakudompet/pass"
)

type SakuDompet struct {
	database                  DatabaseInterface
}

type DatabaseInterface interface {
	GetPasses(context.Context) ([]pass.Pass, error)
	FindPass(context.Context, uint) (pass.Pass, error)
	DeletePass(context.Context, uint) (pass.Pass, error)
	CreatePass(context.Context, pass.Pass) (pass.Pass, error)
	FindDebits(context.Context) ([]saku.Debit, error)
	CreateDebit(context.Context, saku.Debit) (saku.Debit, error)
	FindDebit(context.Context, uint) (saku.Debit, error)
	PutDebit(context.Context, uint, saku.Debit) (saku.Debit, error)
	DeleteDebit(context.Context, uint) (saku.Debit, error)
}

func NewSakuDompet(db DatabaseInterface) *SakuDompet {
	return &SakuDompet{
		database:                   db,
	}
}

func (s *SakuDompet) FindPass(ctx context.Context, id uint) (pass.Pass, error) {
	return s.database.FindPass(ctx,id)
}

func (s *SakuDompet) DeletePass(ctx context.Context, id uint) (pass.Pass, error) {
	return s.database.DeletePass(ctx,id)
}

func (s *SakuDompet) CreatePass(ctx context.Context, pass pass.Pass) (pass.Pass, error) {
	return s.database.CreatePass(ctx, pass)
}

func (s *SakuDompet) GetDebits(ctx context.Context) ([]saku.Debit, error) {
	return s.database.FindDebits(ctx)
}

func (s *SakuDompet) FindDebit(ctx context.Context, id uint) (saku.Debit, error) {
	return s.database.FindDebit(ctx,id)
}

func (s *SakuDompet) CreateDebit(ctx context.Context, debit saku.Debit) (saku.Debit, error) {
	return s.database.CreateDebit(ctx, debit)
}

func (s *SakuDompet) PutDebit(ctx context.Context, id uint,  debit saku.Debit) (saku.Debit, error) {
	return s.database.PutDebit(ctx, id, debit)
}

func (s *SakuDompet) DeleteDebit(ctx context.Context, id uint) (saku.Debit, error) {
	return s.database.DeleteDebit(ctx,id)
}

func (s *SakuDompet) GetPasses(ctx context.Context) ([]pass.Pass, error) {
	return s.database.GetPasses(ctx)
}