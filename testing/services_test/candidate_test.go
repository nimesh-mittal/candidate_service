package services_test

import (
	"candidate_service/daos"
	"candidate_service/models"
	"candidate_service/services"
	"testing"

	"github.com/stretchr/testify/suite"
	"gotest.tools/assert"
)

type CandidateTestSuite struct {
	suite.Suite
	ctx         services.CandidateServiceContext
	CandidateID string
}

func (suite *CandidateTestSuite) SetupTest() {
	DBCtx, _ := daos.NewCandidateSQLDBContext("sqlite3", "testing.db")
	ctx := services.CandidateServiceContext{DAO: *DBCtx}
	suite.ctx = ctx
	suite.ctx.CreateCandidate(&models.Candidate{ID: "1", Name: "name1"})
	suite.ctx.CreateCandidate(&models.Candidate{ID: "2", Name: "name2"})
	suite.ctx.CreateCandidate(&models.Candidate{ID: "3", Name: "name3"})
	c, _ := suite.ctx.CreateCandidate(&models.Candidate{ID: "4", Name: "name4"})
	suite.CandidateID = c.ID
}

func (suite *CandidateTestSuite) TearDownTest() {
	suite.ctx.DAO.DB.Delete(models.Candidate{})
}

func (suite *CandidateTestSuite) TestListCandidate() {
	candidates, err := suite.ctx.ListCandidates(10, 2)

	if err != nil {
		suite.Error(err)
	}

	assert.Equal(suite.T(), 2, len(*candidates), "list results are not 10")
}

func (suite *CandidateTestSuite) TestGetCandidate() {
	candidate, err := suite.ctx.GetCandidate(suite.CandidateID)

	if err != nil {
		suite.Error(err)
	}

	assert.Equal(suite.T(), "name4", candidate.Name, "candidate does not match")
}

func (suite *CandidateTestSuite) TestUpdateCandidate() {

	_, err := suite.ctx.UpdateCandidate(suite.CandidateID, &models.Candidate{Name: "name_modified"})

	if err != nil {
		suite.Error(err)
	}

	candidate, _ := suite.ctx.GetCandidate(suite.CandidateID)

	assert.Equal(suite.T(), "name_modified", candidate.Name, "update not effected")
}

func (suite *CandidateTestSuite) TestDeleteCandidate() {

	candidate, err := suite.ctx.DeleteCandidate("1")

	if err != nil {
		suite.Error(err)
	}

	assert.Equal(suite.T(), "name1", candidate.Name, "record not deleted")
}

func (suite *CandidateTestSuite) TestCreateCandidate() {

	candidate, err := suite.ctx.CreateCandidate(&models.Candidate{ID: "7", Name: "test1"})

	if err != nil {
		suite.Error(err)
	}

	assert.Equal(suite.T(), "test1", candidate.Name, "record not created")
}

func TestCandidateTestSuite(t *testing.T) {
	suite.Run(t, new(CandidateTestSuite))
}
