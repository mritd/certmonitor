package alarm

type Config struct {
	Type    string   `mapstructure:"type"`
	Targets []string `mapstructure:"targets"`
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
