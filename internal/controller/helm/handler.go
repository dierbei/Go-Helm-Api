package helmcontroller

import (
	"githup.com/dierbei/go-helm-api/internal/service/helm"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	// UploadChart
	// 上传Chart
	//UploadChart() gin.HandlerFunc

	// PullChart
	// 拉取Chart
	//PullChart() gin.HandlerFunc

	// PackageChart
	// 打包文件夹为 Chart
	//PackageChart() gin.HandlerFunc

	// InstallOrUpgradeRelease
	// 安装或更新 Chart
	InstallOrUpgradeRelease() gin.HandlerFunc

	// CreateChart
	// 创建 Chart
	//CreateChart() gin.HandlerFunc

	// ListRepositoryChart
	// 仓库 Chart 列表
	ListRepositoryChart() gin.HandlerFunc

	// UninstallRelease
	// 卸载 Chart
	//UninstallRelease() gin.HandlerFunc

	// AddRepository
	// 添加 Chart 仓库
	AddRepository() gin.HandlerFunc

	// UpdateRepository
	// 更新 Chart 仓库
	UpdateRepository() gin.HandlerFunc

	// ListRepository
	// Chart 仓库列表
	ListRepository() gin.HandlerFunc

	i()
}

type handler struct {
	Svc helmservice.Service
}

func newHandler(svc helmservice.Service) *handler {
	return &handler{Svc: svc}
}

func New(group *gin.RouterGroup, svc helmservice.Service) {
	h := newHandler(svc)

	repositoryGroup := group.Group("/repositories")
	{
		// helm repo list
		repositoryGroup.GET("", h.ListRepository())

		// helm repo add
		repositoryGroup.POST("", h.AddRepository())

		// helm repo update
		repositoryGroup.PUT("", h.UpdateRepository())
	}

	chartGroup := group.Group("/charts/repository/:repository")
	{
		// helm search repo
		chartGroup.GET("", h.ListRepositoryChart())

		// helm show
		chartGroup.GET("/chart/:chart")

		// upload chart
		chartGroup.POST("")

		// delete chart
		chartGroup.DELETE("")
	}

	releaseGroup := group.Group("/releases/namespace/:namespace")
	{
		// helm install or helm upgrade
		releaseGroup.POST("", h.InstallOrUpgradeRelease())

		// helm uninstall
		releaseGroup.DELETE("/repository/:repository/release/:release")

		// helm get
		releaseGroup.GET("/release/:release")
	}
}

func (h *handler) i() {
}
