package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/validator.v2"
)

type VxlanPolicyAgent struct {
	PollInterval      int    `json:"poll_interval" validate:"nonzero"`
	Datastore         string `json:"cni_datastore_path" validate:"nonzero"`
	PolicyServerURL   string `json:"policy_server_url" validate:"nonzero"`
	VNI               int    `json:"vni" validate:"nonzero"`
	FlannelSubnetFile string `json:"flannel_subnet_file" validate:"nonzero"`
	MetronAddress     string `json:"metron_address" validate:"nonzero"`
	ServerCACertFile  string `json:"ca_cert_file" validate:"nonzero"`
	ClientCertFile    string `json:"client_cert_file" validate:"nonzero"`
	ClientKeyFile     string `json:"client_key_file" validate:"nonzero"`
}

func New(configFilePath string) (*VxlanPolicyAgent, error) {
	cfg := &VxlanPolicyAgent{}
	if _, err := os.Stat(configFilePath); err != nil {
		return cfg, fmt.Errorf("file does not exist: %s", err)
	}

	configBytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return cfg, fmt.Errorf("reading config file: %s", err)
	}

	err = json.Unmarshal(configBytes, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("parsing config (%s): %s", configFilePath, err)
	}

	if errs := validator.Validate(cfg); errs != nil {
		return cfg, fmt.Errorf("invalid config: %s", errs)
	}

	return cfg, nil
}
