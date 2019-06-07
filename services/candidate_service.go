package services

import (
	"candidate_service/commons"
	"candidate_service/config"
	"candidate_service/daos"
	"candidate_service/models"
	"log"

	"github.com/sirupsen/logrus"
)

type CandidateServiceContext struct {
	DAO daos.CandidateMongoDBContext
}

// You can choose between CandidateDBContext or CandidateMongoDBContext
func NewCandidateServiceContext() *CandidateServiceContext {

	candidateDBContext, err := daos.NewCandidateMongoDBContext(config.GetInstance().Database.MongoURL)
	//candidateDBContext, err := daos.NewCandidateSQLDBContext(config.GetInstance().Database.Dialect,
	//	config.GetInstance().Database.URL)

	if err != nil {
		log.Fatal("unable to create candidate service context", err)
	}

	return &CandidateServiceContext{DAO: *candidateDBContext}
}

func (ctx *CandidateServiceContext) SafeClose() {
	ctx.DAO.SafeClose()
}

func (ctx *CandidateServiceContext) ListCandidates(fCtx *models.FlowCtx, limit int, offset int) (*[]models.Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("listing candidates")

	candidates, err := ctx.DAO.ListCandidates(fCtx, limit, offset)

	if err != nil {
		return nil, err
	}

	return candidates, nil
}

func (ctx *CandidateServiceContext) GetCandidate(fCtx *models.FlowCtx, cid string) (*models.Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("get candidate by id")

	c, err := ctx.DAO.GetCandidate(fCtx, cid)

	if err != nil {
		return nil, err
	}
	return c, nil
}

func (ctx *CandidateServiceContext) CreateCandidate(fCtx *models.FlowCtx, candidate *models.Candidate) (*models.Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("creating candidate")

	c, err := ctx.DAO.CreateCandidate(fCtx, candidate)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (ctx *CandidateServiceContext) UpdateCandidate(fCtx *models.FlowCtx, cid string, entity *models.Candidate) (string, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("updating candidate")

	cid, err := ctx.DAO.UpdateCandidate(fCtx, cid, entity)

	if err != nil {
		return "", err
	}

	return cid, nil
}

func (ctx *CandidateServiceContext) DeleteCandidate(fCtx *models.FlowCtx, cid string) (*models.Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("deleting candidate")

	c, err := ctx.DAO.DeleteCandidate(fCtx, cid)

	if err != nil {
		return nil, err
	}
	return c, nil
}
