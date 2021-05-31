package alarm

type Config struct {
	Type    string   `yaml:"type"`
	Targets []string `yaml:"targets"`
}

func ExampleConfig() []*Config {
	return []*Config{
		{
			Type: "smtp",
			Targets: []string{
				"mritd1234@gmail.com",
			},
		},
		{
			Type: "webhook",
			Targets: []string{
				"https://google.com",
			},
		},
		{
			Type: "telegram",
			Targets: []string{
				"-124568340456",
			},
		},
	}
}
