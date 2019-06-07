package daos

import (
	"candidate_service/commons"
	"candidate_service/models"

	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"

	"github.com/sirupsen/logrus"
)

type CandidateDBContext struct {
	DB *gorm.DB
}

func NewCandidateSQLDBContext(dialect string, dbName string) (*CandidateDBContext, error) {
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
	db.AutoMigrate(&models.Candidate{}, &models.Address{})

	defer logrus.Info("sql database setup completed")
	return &CandidateDBContext{DB: db}, nil
}

func (ctx *CandidateDBContext) SafeClose() {
	ctx.DB.Close()
}

func (ctx *CandidateDBContext) ListCandidates(limit int, offset int) (*[]models.Candidate, error) {
	var candidates []models.Candidate
	res := ctx.DB.Preload("Address").Offset(offset).Limit(limit).Find(&candidates)

	if res.Error != nil {
		logrus.Error(res.Error.Error())
		return nil, res.Error
	}

	return &candidates, nil
}

func (ctx *CandidateDBContext) GetCandidate(cid string) (*models.Candidate, error) {
	var candidate models.Candidate

	response := ctx.DB.Preload("Address").Where(&models.Candidate{ID: cid}).First(&candidate)
	if response.RecordNotFound() {
		return nil, response.Error
	}
	return &candidate, nil
}

func (ctx *CandidateDBContext) CreateCandidate(candidate *models.Candidate) (*models.Candidate, error) {
	res := ctx.DB.Create(candidate)
	if res.Error != nil {
		logrus.Error(res.Error.Error())
		return nil, res.Error
	}

	return candidate, nil
}

func (ctx *CandidateDBContext) UpdateCandidate(cid string, entity *models.Candidate) (string, error) {
	candidate, err := ctx.GetCandidate(cid)

	if err != nil {
		return commons.EMPTY, err
	}

	// Hack because Omit does not work for primary id
	entity.ID = cid
	ctx.DB.Model(candidate).Updates(entity)
	return cid, nil
}

func (ctx *CandidateDBContext) DeleteCandidate(cid string) (*models.Candidate, error) {
	tx := ctx.DB.Begin()

	var candidate models.Candidate

	resp := tx.Preload("Address").Where(&models.Candidate{ID: cid}).First(&candidate)
	if resp.RecordNotFound() {
		return nil, resp.Error
	}

	response := tx.Where(&models.Candidate{ID: cid}).Delete(models.Candidate{})
	if response.Error != nil {
		tx.Rollback()
		return nil, response.Error
	}

	response = tx.Where(&models.Address{CandidateID: cid}).Delete(models.Address{})
	if response.Error != nil {
		tx.Rollback()
		return nil, response.Error
	}

	tx.Commit()
	return &candidate, nil
}
