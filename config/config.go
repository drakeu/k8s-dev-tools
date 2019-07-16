package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	kubeconfig *string
	k8sConfig  *K8SConfig
}

func NewConfig() *Config {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String(
			"kubeconfig",
			filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file",
		)
	} else {
		kubeconfig = flag.String(
			"kubeconfig",
			"",
			"absolute path to the kubeconfig file",
		)
	}
	flag.Parse()

	return &Config{
		kubeconfig: kubeconfig,
	}
}

func (cfg *Config) LoadConfiguration() {
	data, err := ioutil.ReadFile(*cfg.kubeconfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	k8sConfig := K8SConfig{}

	err = yaml.Unmarshal(data, &k8sConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	cfg.k8sConfig = &k8sConfig
}

func (cfg *Config) GetCurrentContext() string {
	return cfg.k8sConfig.CurrentContext
}

func (cfg *Config) GetAvailableContexts() []string {
	results := []string{}
	for _, container := range cfg.k8sConfig.ContextContainers {
		results = append(results, container.Name)
	}
	return results
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
