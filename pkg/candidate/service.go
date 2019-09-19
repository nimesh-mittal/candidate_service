package candidate

import (
	"candidate_service/config"
	"candidate_service/pkg/commons"
	"log"

	"github.com/sirupsen/logrus"
)

// ServiceContext maintains repo instance
type ServiceContext struct {
	DAO DBContext
}

// NewServiceContext creates new ServiceContext
func NewServiceContext() *ServiceContext {

	// candidateDBContext, err := daos.NewCandidateMongoDBContext(config.GetInstance().Database.MongoURL)
	candidateDBContext, err := NewSQLDBContext(
		config.GetInstance().Database.Dialect,
		config.GetInstance().Database.URL)

	if err != nil {
		log.Fatal("unable to create candidate service context", err)
	}

	return &ServiceContext{DAO: *candidateDBContext}
}

// SafeClose gets called when main is about to complete
func (ctx *ServiceContext) SafeClose() {
	ctx.DAO.SafeClose()
}

// ListCandidates to list candidates with pagination
func (ctx *ServiceContext) ListCandidates(fCtx *commons.FlowContext, limit int, offset int) (*[]Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("listing candidates")

	candidates, err := ctx.DAO.ListCandidates(fCtx, limit, offset)

	if err != nil {
		return nil, err
	}

	return candidates, nil
}

// GetCandidate gets candidate by candidate id
func (ctx *ServiceContext) GetCandidate(fCtx *commons.FlowContext, cid string) (*Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("get candidate by id")

	c, err := ctx.DAO.GetCandidate(fCtx, cid)

	if err != nil {
		return nil, err
	}
	return c, nil
}

// CreateCandidate creates candidate
func (ctx *ServiceContext) CreateCandidate(f *commons.FlowContext, candidate *Candidate) (*Candidate, error) {
	logrus.WithField(commons.TrackingID, f.TrackingID).Info("creating candidate")

	c, err := ctx.DAO.CreateCandidate(f, candidate)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// UpdateCandidate updates candidate by candidate id
func (ctx *ServiceContext) UpdateCandidate(fCtx *commons.FlowContext, cid string, entity *Candidate) (string, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("updating candidate")

	cid, err := ctx.DAO.UpdateCandidate(fCtx, cid, entity)

	if err != nil {
		return "", err
	}

	return cid, nil
}

// DeleteCandidate deletes candidate by candidate id
func (ctx *ServiceContext) DeleteCandidate(fCtx *commons.FlowContext, cid string) (*Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("deleting candidate")

	c, err := ctx.DAO.DeleteCandidate(fCtx, cid)

	if err != nil {
		return nil, err
	}
	return c, nil
}
