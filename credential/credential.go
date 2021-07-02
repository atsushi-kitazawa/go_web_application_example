package credential

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Credential struct {
    SecurityKey string `yaml:"securityKey"`
    ClientId string `yaml:"clientId"`
    Secret string `yaml:"secret"`
}

func Load() Credential {
    buf, err := ioutil.ReadFile("credential.yaml")
    if err != nil {
	log.Fatal("failed to read credential.yaml", err)
	panic(err)
    }

    var c Credential
    err = yaml.Unmarshal(buf, &c)
    if err != nil {
	log.Fatal("Unmarshal credential", err)
	panic(err)
    }

    return c
}
