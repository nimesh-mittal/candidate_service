package handlers

import (
	"candidate_service/commons"
	"candidate_service/config"
	"candidate_service/infra"
	"candidate_service/models"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/shirou/gopsutil/disk"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"

	"github.com/go-chi/chi"
	"github.com/shirou/gopsutil/mem"
)

type AdminResourceContext struct {
	HeartbeatStatus string
}

func NewAdminResource() *AdminResourceContext {
	return &AdminResourceContext{HeartbeatStatus: commons.GREEN_STATUS}
}

func (ctx *AdminResourceContext) SafeClose() {
}

func (ctx *AdminResourceContext) NewAdminRouter() http.Handler {
	r := chi.NewRouter()

	r.Get(infra.WrapNR("/heartbeat", ctx.GetHeartbeat))
	r.Get(infra.WrapNR("/heartbeat/_stop", ctx.StopHeartbeat))
	r.Get(infra.WrapNR("/heartbeat/_start", ctx.StartHeartbeat))
	r.Get(infra.WrapNR("/_info", ctx.GetInfo))
	r.Get(infra.WrapNR("/_health", ctx.GetHealth))

	return r
}

// @Summary Get Heartbeat
// @Tags Admin
// @Success 200 {object} models.HeartbeatResponse
// @Failure 400 {object} models.Response
// @Router /admin/heartbeat [get]
func (ctx *AdminResourceContext) GetHeartbeat(w http.ResponseWriter, r *http.Request) {
	logrus.Info(config.GetInstance().Database.URL)
	heartbeat := models.Heartbeat{Status: ctx.HeartbeatStatus}
	res := commons.MakeResp(heartbeat, commons.EMPTY, nil)
	w.Write(res)
}

// @Summary Disable Heartbeat
// @Tags Admin
// @Success 200 {object} models.HeartbeatResponse
// @Failure 400 {object} models.Response
// @Router /admin/heartbeat/_stop [get]
func (ctx *AdminResourceContext) StopHeartbeat(w http.ResponseWriter, r *http.Request) {
	ctx.HeartbeatStatus = commons.RED_STATUS
	heartbeat := models.Heartbeat{Status: ctx.HeartbeatStatus}
	res := commons.MakeResp(heartbeat, commons.EMPTY, nil)
	w.Write(res)
}

// @Summary Enable Heartbeat
// @Tags Admin
// @Success 200 {object} models.HeartbeatResponse
// @Failure 400 {object} models.Response
// @Router /admin/heartbeat/_start [get]
func (ctx *AdminResourceContext) StartHeartbeat(w http.ResponseWriter, r *http.Request) {
	ctx.HeartbeatStatus = commons.GREEN_STATUS
	heartbeat := models.Heartbeat{Status: ctx.HeartbeatStatus}
	res := commons.MakeResp(heartbeat, commons.EMPTY, nil)
	w.Write(res)
}

// @Summary Get Health
// @Tags Admin
// @Success 200 {object} models.HeartbeatResponse
// @Failure 400 {object} models.Response
// @Router /admin/_health [get]
func (ctx *AdminResourceContext) GetHealth(w http.ResponseWriter, r *http.Request) {

	health := models.Health{KinesisHealth: commons.OK, CacheHealth: commons.OK, DBConnectionHealth: commons.OK,
		CPU: commons.OK, Memory: commons.OK, Host: commons.OK, Disk: commons.OK}

	res := commons.MakeResp(health, commons.EMPTY, nil)
	w.Write(res)
}

// @Summary Get Info
// @Tags Admin
// @Success 200 {object} models.HeartbeatResponse
// @Failure 400 {object} models.Response
// @Router /admin/_info [get]
func (ctx *AdminResourceContext) GetInfo(w http.ResponseWriter, r *http.Request) {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	h, _ := host.Info()
	d, _ := disk.Usage("/")
	health := models.Info{KinesisHealth: commons.OK, CacheHealth: commons.OK, DBConnectionHealth: commons.OK,
		CPU: c, Memory: v, Host: h, Disk: d}

	res := commons.MakeResp(health, commons.EMPTY, nil)
	w.Write(res)
}
