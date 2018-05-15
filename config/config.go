package config

type DockerConfig struct {
	DockerPath string // default is /usr/bin/docker
	DockerHost string // default is /var/run/docker.sock
}

type KuberConfig struct {
	KubectlPath string // default is /usr/bin/kubectl
	KubectlAPI  string // default is https://localhost:6443
}

type BackendConfig struct {
	Type      string // docker or kubectl (for now)
	Docker    DockerConfig
	Kube      KuberConfig
	ExtraArgs []string // extra args pass to docker or kubectl
}

type ControlConfig struct {
	Enable  bool
	Start   bool
	Stop    bool
	Restart bool
	Kill    bool
}

type Config struct {
	Port     int
	LogLevel string
	Control  ControlConfig
	Backend  BackendConfig
	Servers  []string // for proxy mode
}
