package candidate

import (
	"candidate_service/config"
	"candidate_service/pkg/commons"
	"log"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	DAO DBContext
}

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

func (ctx *ServiceContext) SafeClose() {
	ctx.DAO.SafeClose()
}

func (ctx *ServiceContext) ListCandidates(fCtx *commons.FlowContext, limit int, offset int) (*[]Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("listing candidates")

	candidates, err := ctx.DAO.ListCandidates(fCtx, limit, offset)

	if err != nil {
		return nil, err
	}

	return candidates, nil
}

func (ctx *ServiceContext) GetCandidate(fCtx *commons.FlowContext, cid string) (*Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("get candidate by id")

	c, err := ctx.DAO.GetCandidate(fCtx, cid)

	if err != nil {
		return nil, err
	}
	return c, nil
}

func (ctx *ServiceContext) CreateCandidate(f *commons.FlowContext, candidate *Candidate) (*Candidate, error) {
	logrus.WithField(commons.TrackingID, f.TrackingID).Info("creating candidate")

	c, err := ctx.DAO.CreateCandidate(f, candidate)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (ctx *ServiceContext) UpdateCandidate(fCtx *commons.FlowContext, cid string, entity *Candidate) (string, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("updating candidate")

	cid, err := ctx.DAO.UpdateCandidate(fCtx, cid, entity)

	if err != nil {
		return "", err
	}

	return cid, nil
}

func (ctx *ServiceContext) DeleteCandidate(fCtx *commons.FlowContext, cid string) (*Candidate, error) {
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("deleting candidate")

	c, err := ctx.DAO.DeleteCandidate(fCtx, cid)

	if err != nil {
		return nil, err
	}
	return c, nil
}
