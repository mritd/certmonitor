package alarm

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type WebHookConfig struct {
	Method  string        `mapstructure:"method"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func WebHookExampleConfig() *WebHookConfig {
	return &WebHookConfig{
		Method:  "get",
		Timeout: 5 * time.Second,
	}
}

func (cfg *WebHookConfig) Send(targets []string, message string) {
	c := http.Client{
		Timeout: cfg.Timeout,
	}
	for _, t := range targets {
		switch strings.ToLower(cfg.Method) {
		case "post":
			cfg.post(c, t, message)
		case "get":
			cfg.get(c, t, message)
		default:
			logrus.Info("Method not support!")
		}
	}
}

func (cfg *WebHookConfig) get(client http.Client, addr string, message string) {
	u, err := url.Parse(addr)
	if err != nil {
		logrus.Errorf("WebHook alarm send failed [%s]: %s", addr, err)
		return
	}

	q := u.Query()
	q.Set("message", message)
	u.RawQuery = q.Encode()
	resp, err := client.Get(u.String())
	if err != nil {
		logrus.Errorf("WebHook alarm send failed [%s]: %s", addr, err)
	} else if http.StatusOK != resp.StatusCode {
		logrus.Errorf("WebHook alarm send failed [%s]: http status code %d", addr, resp.StatusCode)
	}
}

func (cfg *WebHookConfig) post(client http.Client, addr string, message string) {
	resp, err := client.PostForm(addr, url.Values{
		"message": {message},
	})
	if err != nil {
		logrus.Errorf("WebHook alarm send failed [%s]: %s", addr, err)
	} else if http.StatusOK != resp.StatusCode {
		logrus.Errorf("WebHook alarm send failed [%s]: http status code %d", addr, resp.StatusCode)
	}
}
