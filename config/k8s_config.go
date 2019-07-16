package config

// K8SConfig is representation of kubernetes configuration file
type K8SConfig struct {
	CurrentContext    string             `yaml:"current-context"`
	ContextContainers []ContextContainer `yaml:"contexts"`
}

type ContextContainer struct {
	Context Context `yaml:"context"`
	Name    string  `yaml:"name"`
}

type Context struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}
