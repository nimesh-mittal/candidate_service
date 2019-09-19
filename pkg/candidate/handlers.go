package candidate

import (
	"candidate_service/infra"
	"candidate_service/pkg/commons"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	uuid2 "github.com/google/uuid"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type ResourceContext struct {
	CandidateService ServiceContext
}

func NewCandidateResource() *ResourceContext {
	candidateServiceContext := NewServiceContext()

	return &ResourceContext{CandidateService: *candidateServiceContext}
}

/**
SafeClose gets called when service gets shutdown
*/
func (ctx *ResourceContext) SafeClose() {
	ctx.CandidateService.SafeClose()
}

/**
New SubRouter for candidates
*/
func (ctx *ResourceContext) NewCandidateRouter() http.Handler {
	r := chi.NewRouter()

	// list candidate
	r.Get(infra.WrapNR("/", ctx.ListCandidate))

	// get candidate
	r.Get(infra.WrapNR("/{CandidateID}", ctx.GetCandidate))

	// create candidate
	r.Post(infra.WrapNR("/", ctx.CreateCandidate))

	// update candidate
	r.Put(infra.WrapNR("/{CandidateID}", ctx.UpdateCandidate))

	// delete candidate
	r.Delete(infra.WrapNR("/{CandidateID}", ctx.DeleteCandidate))

	return r
}

// @Summary List Candidate
// @Tags Candidate
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /candidates [get]
func (ctx *ResourceContext) ListCandidate(w http.ResponseWriter, r *http.Request) {
	tx := infra.StartTx("List Candidate")
	defer infra.EndTx(tx)

	fCtx := &commons.FlowContext{TrackingID: uuid2.New().String()}
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("received list candidates request")

	limitStr := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitStr)

	offsetStr := r.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(offsetStr)

	errs, ok := commons.ValidatePaginationParams(fCtx, limitStr, offsetStr)
	if !ok {
		msg := strings.Join(errs, " ")
		res := commons.MakeResp(nil, commons.InvalidRequestParameter, errors.New(msg))
		w.Write(res)
		return
	}

	candidates, err := ctx.CandidateService.ListCandidates(fCtx, limit, offset)
	res := commons.MakeResp(candidates, commons.Empty, err)
	w.Write(res)
}

// @Summary Get Candidate
// @Tags Candidate
// @Param CandidateID path string true "Candidate ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /candidates/{CandidateID} [get]
func (ctx *ResourceContext) GetCandidate(w http.ResponseWriter, r *http.Request) {
	fCtx := &commons.FlowContext{TrackingID: uuid2.New().String()}
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("received get candidate request")

	cid := chi.URLParam(r, "CandidateID")

	errs, ok := commons.ValidateID(fCtx, cid)
	if !ok {
		msg := strings.Join(errs, " ")
		res := commons.MakeResp(nil, commons.InvalidRequestParameter, errors.New(msg))
		w.Write(res)
		return
	}

	c, err := ctx.CandidateService.GetCandidate(fCtx, cid)

	res := commons.MakeResp(c, commons.Empty, err)
	w.Write(res)
}

// @Summary Create Candidate
// @Tags Candidate
// @Param candidate body models.Candidate true "Create Candidate"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /candidates [post]
func (ctx *ResourceContext) CreateCandidate(w http.ResponseWriter, r *http.Request) {
	fCtx := &commons.FlowContext{TrackingID: uuid2.New().String()}
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("received create candidate request")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var entity Candidate
	err := decoder.Decode(&entity)
	if err != nil {
		res := commons.MakeResp(nil, commons.Empty, err)
		w.Write(res)
		return
	}

	entity.ID = uuid2.New().String()
	for i := 0; i < len(entity.Address); i++ {
		entity.Address[i].ID = uuid2.New().String()
		entity.Address[i].CandidateID = entity.ID
	}

	ok, err := commons.Validate(&entity)

	if !ok {
		b := commons.MakeResp(nil, commons.InvalidRequestBody, err)
		w.Write(b)
		return
	}

	c, err := ctx.CandidateService.CreateCandidate(fCtx, &entity)
	b := commons.MakeResp(c, commons.Empty, err)

	w.Write(b)
}

// @Summary Update Candidate
// @Tags Candidate
// @Param candidate body models.Candidate true "Update Candidate"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /candidates/{CandidateID} [put]
func (ctx *ResourceContext) UpdateCandidate(w http.ResponseWriter, r *http.Request) {
	fCtx := &commons.FlowContext{TrackingID: uuid2.New().String()}
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("received update candidate request")

	cid := chi.URLParam(r, "CandidateID")

	if commons.IsEmpty(cid) {
		res := commons.MakeResp(nil, commons.Empty, errors.New("id is missing"))
		w.Write(res)
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var entity Candidate
	err := decoder.Decode(&entity)

	if err != nil {
		res := commons.MakeResp(nil, commons.Empty, err)
		w.Write(res)
		return
	}

	id, err := ctx.CandidateService.UpdateCandidate(fCtx, cid, &entity)
	res := commons.MakeResp(id, commons.Empty, err)
	w.Write(res)
}

// @Summary Delete Candidate
// @Tags Candidate
// @Param CandidateID path string true "Candidate ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /candidates/{CandidateID} [delete]
func (ctx *ResourceContext) DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	fCtx := &commons.FlowContext{TrackingID: uuid2.New().String()}
	logrus.WithField(commons.TrackingID, fCtx.TrackingID).Info("received delete candidate request")

	cid := chi.URLParam(r, "CandidateID")

	errs, ok := commons.ValidateID(fCtx, cid)

	if !ok {
		msg := strings.Join(errs, " ")
		res := commons.MakeResp(nil, commons.InvalidRequestParameter, errors.New(msg))
		w.Write(res)
		return
	}

	c, err := ctx.CandidateService.DeleteCandidate(fCtx, cid)

	res := commons.MakeResp(c, commons.Empty, err)
	w.Write(res)
}
