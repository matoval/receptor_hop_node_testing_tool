package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

type deploymentVars struct {
	DeploymentName    string
	Namespace         string
	ExternalHostName  string
	ExternalIPAddress string
	ConfigMap         string
}

func KubeLauncher(socketPath string, tlsCert string, tlsKey string) {
	vars := deploymentVars{
		DeploymentName:    "hop-node-test-deployment",
		Namespace:         "receptor-hop-node-test",
		ExternalHostName:  "example.com",
		ExternalIPAddress: "0.0.0.0",
		ConfigMap:         "hop-node-deployment-receptor-config",
	}

	tmpl, err := template.ParseFiles("template/deployment.tmpl")
	if err != nil {
		fmt.Printf("error templating: %v", err)
	}
	file, err := os.Create("/tmp/deployment.yml")
	if err != nil {
		fmt.Printf("error creating /tmp/deployment.yml file,: %v", err)
	}

	tmpl.Execute(file, vars)

	cmd := exec.Command("receptorctl", "--cert", tlsCert, "--key", tlsKey, "--socket", socketPath, "work", "submit", "launch_deployment", "--param", "secret_kube_config=@$HOME/.kube/config", "--param", "secret_kube_deployment=@/tmp/deployment.yml", "--no-payload")

	fmt.Printf("CMD: %v\n", cmd)

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}

	fmt.Println(string(stdout))
}
