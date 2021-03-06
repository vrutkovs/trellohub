package settings

import (
	"github.com/vrutkovs/trellohub/pkg/github"
	"github.com/vrutkovs/trellohub/pkg/trello"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// DefaultSyncTimeoutMinutes sets default sync period
const DefaultSyncTimeoutMinutes = 5

// Settings holds app-level settings
type Settings struct {
	Trello      trello.Settings `yaml:"trello"`
	Github      github.Settings `yaml:"github"`
	SyncTimeout uint64          `yaml:"sync_timeout"`
}

// LoadSettings creates Settings object from yaml
func LoadSettings(path string) (*Settings, error) {
	s := Settings{
		SyncTimeout: DefaultSyncTimeoutMinutes,
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
