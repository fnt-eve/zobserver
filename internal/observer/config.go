package observer

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

//nolint:revive
type KmFeedConfig struct {
	QueueName        string `default:"" split_words:"true"`
	TTW              string `default:"10" split_words:"true"`
	EsiUserAgent     string `default:"zobserver" split_words:"true"`
	Destinations     string `required:"false" split_words:"true"`
	DestinationsFile string `required:"false" split_words:"true"`
}

type Destination struct {
	Name            string           `yaml:"name"`
	CharacterIDs    []int32          `yaml:"character_ids"`
	CorporationIDs  []int32          `yaml:"corporation_ids"`
	AllianceIDs     []int32          `yaml:"alliance_ids"`
	All             bool             `yaml:"all"`
	DiscordWebhooks []DiscordWebhook `yaml:"discord_webhooks"`
}

type DiscordWebhook struct {
	ID    string `yaml:"id"`
	Token string `yaml:"token"`
}

func GetDestinations(c KmFeedConfig) ([]Destination, error) {
	if c.DestinationsFile != "" {
		confBuf, err := os.ReadFile(c.DestinationsFile)
		if err != nil {
			return nil, err
		}
		return ParseConfig(confBuf)
	}

	if c.Destinations != "" {
		return ParseConfig([]byte(c.Destinations))
	}

	return nil, errors.New("no destinations specified")
}

func ParseConfig(conf []byte) ([]Destination, error) {
	var cfg []Destination
	err := yaml.Unmarshal(conf, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
