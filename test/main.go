package main

import "os/exec"

func main() {
	// 创建 kube config
	CreateK8sConfig()

	// 创建 k8s ca
	CreateK8sCertificate()

	// 创建 helmfile 文件
	CreateHelmfile()

	// 执行命令
	cmd := exec.Command("helmfile", "-f", "helmfilepath", "--kubeconfig", "...", "--ca", "...")
	cmd.Run()

	// 删除 kube config
	DeleteK8sConfig()

	// 删除 k8s ca
	DeleteK8sCertificate()

	// 删除 helmfile
	DeleteHelmfile()
}
