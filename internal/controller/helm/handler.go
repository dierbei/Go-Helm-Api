package helmcontroller

import (
	v1 "githup.com/dierbei/go-helm-api/internal/controller/helm/v1"
	"githup.com/dierbei/go-helm-api/internal/service/helm"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	// UploadChart
	// 上传Chart
	UploadChart() gin.HandlerFunc

	// DeleteChart
	// 删除 Chart
	DeleteChart() gin.HandlerFunc

	// CreateAndUploadChart
	// 创建 Chart 包并上传
	CreateAndUploadChart() gin.HandlerFunc

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

	// UninstallRelease
	// 卸载 Chart
	UninstallRelease() gin.HandlerFunc

	// ShowReleaseInfo
	// 显示 Release 信息
	ShowReleaseInfo() gin.HandlerFunc

	// GetChartInfo
	// 获取 Chart 信息
	GetChartInfo() gin.HandlerFunc

	// RollBack
	// 回滚版本
	RollBack() gin.HandlerFunc

	// ListReleaseHistories
	// Release 部署历史
	ListReleaseHistories() gin.HandlerFunc
}

func New(group *gin.RouterGroup, svc helmservice.Service) {
	h := v1.NewHandler(svc)

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
		chartGroup.GET("/chart", h.GetChartInfo())

		// upload chart
		chartGroup.POST("", h.UploadChart())

		// delete chart
		chartGroup.DELETE("", h.DeleteChart())

		// create chart
		chartGroup.POST("/upload", h.CreateAndUploadChart())
	}

	releaseGroup := group.Group("/releases/namespace/:namespace")
	{
		// helm install or helm upgrade
		releaseGroup.POST("", h.InstallOrUpgradeRelease())

		// helm uninstall
		releaseGroup.DELETE("/release/:release", h.UninstallRelease())

		// helm get
		releaseGroup.GET("/release/:release", h.ShowReleaseInfo())

		// helm rollback
		releaseGroup.POST("/rollback", h.RollBack())

		// helm history
		releaseGroup.POST("/history", h.ListReleaseHistories())
	}
}
