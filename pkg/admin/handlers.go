package admin

import (
	"candidate_service/config"
	"candidate_service/infra"
	"candidate_service/pkg/commons"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
)

// ResourceContext holds service instance
type ResourceContext struct {
	HeartbeatStatus string
}

// NewAdminResource creates new resource context
func NewAdminResource() *ResourceContext {
	return &ResourceContext{HeartbeatStatus: commons.GreenStatus}
}

// SafeClose gets called when program terminates
func (ctx *ResourceContext) SafeClose() {
}

// NewAdminRouter creates new admin router
func (ctx *ResourceContext) NewAdminRouter() http.Handler {
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
// @Success 200 {object} HeartbeatResponse
// @Failure 400 {object} Response
// @Router /admin/heartbeat [get]
func (ctx *ResourceContext) GetHeartbeat(w http.ResponseWriter, r *http.Request) {
	logrus.Info(config.GetInstance().Database.URL)
	heartbeat := Heartbeat{Status: ctx.HeartbeatStatus}
	res := commons.MakeResp(heartbeat, commons.Empty, nil)
	w.Write(res)
}

// @Summary Disable Heartbeat
// @Tags Admin
// @Success 200 {object} HeartbeatResponse
// @Failure 400 {object} Response
// @Router /admin/heartbeat/_stop [get]
func (ctx *ResourceContext) StopHeartbeat(w http.ResponseWriter, r *http.Request) {
	ctx.HeartbeatStatus = commons.RedStatus
	heartbeat := Heartbeat{Status: ctx.HeartbeatStatus}
	res := commons.MakeResp(heartbeat, commons.Empty, nil)
	w.Write(res)
}

// @Summary Enable Heartbeat
// @Tags Admin
// @Success 200 {object} HeartbeatResponse
// @Failure 400 {object} Response
// @Router /admin/heartbeat/_start [get]
func (ctx *ResourceContext) StartHeartbeat(w http.ResponseWriter, r *http.Request) {
	ctx.HeartbeatStatus = commons.GreenStatus
	heartbeat := Heartbeat{Status: ctx.HeartbeatStatus}
	res := commons.MakeResp(heartbeat, commons.Empty, nil)
	w.Write(res)
}

// @Summary Get Health
// @Tags Admin
// @Success 200 {object} HeartbeatResponse
// @Failure 400 {object} Response
// @Router /admin/_health [get]
func (ctx *ResourceContext) GetHealth(w http.ResponseWriter, r *http.Request) {

	health := Health{KinesisHealth: commons.Ok, CacheHealth: commons.Ok, DBConnectionHealth: commons.Ok,
		CPU: commons.Ok, Memory: commons.Ok, Host: commons.Ok, Disk: commons.Ok}

	res := commons.MakeResp(health, commons.Empty, nil)
	w.Write(res)
}

// @Summary Get Info
// @Tags Admin
// @Success 200 {object} HeartbeatResponse
// @Failure 400 {object} Response
// @Router /admin/_info [get]
func (ctx *ResourceContext) GetInfo(w http.ResponseWriter, r *http.Request) {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	h, _ := host.Info()
	d, _ := disk.Usage("/")
	health := Info{KinesisHealth: commons.Ok, CacheHealth: commons.Ok, DBConnectionHealth: commons.Ok,
		CPU: c, Memory: v, Host: h, Disk: d}

	res := commons.MakeResp(health, commons.Empty, nil)
	w.Write(res)
}
