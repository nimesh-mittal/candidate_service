package candidate

import (
	"candidate_service/pkg/commons"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type DBContext struct {
	DB *gorm.DB
}

func NewSQLDBContext(dialect string, dbName string) (*DBContext, error) {
	db, err := gorm.Open(dialect, dbName)

	if err != nil {
		logrus.Fatal("failed to connect database", err)
		return nil, err
	}

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	db.SetLogger(logrus.StandardLogger())

	// Migrate the schema
	db.AutoMigrate(&Candidate{}, &Address{})

	defer logrus.Info("sql database setup completed")
	return &DBContext{DB: db}, nil
}

func (ctx *DBContext) SafeClose() {
	ctx.DB.Close()
}

func (ctx *DBContext) ListCandidates(f *commons.FlowContext, limit int, offset int) (*[]Candidate, error) {
	var candidates []Candidate
	res := ctx.DB.Preload("Address").Offset(offset).Limit(limit).Find(&candidates)

	if res.Error != nil {
		logrus.Error(res.Error.Error())
		return nil, res.Error
	}

	return &candidates, nil
}

func (ctx *DBContext) GetCandidate(f *commons.FlowContext, cid string) (*Candidate, error) {
	var candidate Candidate

	response := ctx.DB.Preload("Address").Where(&Candidate{ID: cid}).First(&candidate)
	if response.RecordNotFound() {
		return nil, response.Error
	}
	return &candidate, nil
}

func (ctx *DBContext) CreateCandidate(fCtx *commons.FlowContext, candidate *Candidate) (*Candidate, error) {
	res := ctx.DB.Create(candidate)
	if res.Error != nil {
		logrus.Error(res.Error.Error())
		return nil, res.Error
	}

	return candidate, nil
}

func (ctx *DBContext) UpdateCandidate(f *commons.FlowContext, cid string, entity *Candidate) (string, error) {
	candidate, err := ctx.GetCandidate(f, cid)

	if err != nil {
		return commons.EMPTY, err
	}

	// Hack because Omit does not work for primary id
	entity.ID = cid
	ctx.DB.Model(candidate).Updates(entity)
	return cid, nil
}

func (ctx *DBContext) DeleteCandidate(f *commons.FlowContext, cid string) (*Candidate, error) {
	tx := ctx.DB.Begin()

	var candidate Candidate

	resp := tx.Preload("Address").Where(&Candidate{ID: cid}).First(&candidate)
	if resp.RecordNotFound() {
		return nil, resp.Error
	}

	response := tx.Where(&Candidate{ID: cid}).Delete(Candidate{})
	if response.Error != nil {
		tx.Rollback()
		return nil, response.Error
	}

	response = tx.Where(&Address{CandidateID: cid}).Delete(Address{})
	if response.Error != nil {
		tx.Rollback()
		return nil, response.Error
	}

	tx.Commit()
	return &candidate, nil
}
