package monitor

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/mritd/certmonitor/alarm"

	"github.com/robfig/cron"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	"github.com/mritd/certmonitor/utils"
)

type Config struct {
	WebSites   []WebSite     `mapstructure:"web_sites"`
	Cron       string        `mapstructure:"cron"`
	BeforeTime time.Duration `mapstructure:"before_time"`
	Timeout    time.Duration `mapstructure:"timeout"`
}

type WebSite struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Address     string `mapstructure:"address"`
}

type WebSiteError struct {
	Message string
}

func (e *WebSiteError) Error() string {
	return e.Message
}

func NewWebSiteError(msg string) *WebSiteError {
	return &WebSiteError{
		Message: msg,
	}
}

func ExampleConfig() Config {
	return Config{
		WebSites: []WebSite{
			{
				"bleem",
				"博客主站点",
				"https://mritd.com",
			},
			{
				"baidu",
				"百度首页",
				"https://baidu.com",
			},
		},
		Cron:       "@every 1h",
		BeforeTime: 7 * 24 * time.Hour,
		Timeout:    10 * time.Second,
	}
}

func check(website WebSite, beforeTime, timeout time.Duration) *WebSiteError {
	logrus.Infof("Check website [%s]...", website.Address)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: timeout,
	}
	resp, err := client.Get(website.Address)
	if !utils.CheckErr(err) {
		return nil
	}
	defer func() { _ = resp.Body.Close() }()

	for _, cert := range resp.TLS.PeerCertificates {
		if !cert.NotAfter.After(time.Now()) {
			return NewWebSiteError(fmt.Sprintf("Website [%s](%s) certificate has expired: %s", website.Name, website.Address, cert.NotAfter.Local().Format("2006-01-02 15:04:05")))
		}

		if cert.NotAfter.Sub(time.Now()) < beforeTime {
			return NewWebSiteError(fmt.Sprintf("Website [%s](%s) certificate will expire, remaining time: %fh", website.Name, website.Address, cert.NotAfter.Sub(time.Now()).Hours()))
		}
	}

	return nil
}

func Start() {
	var config Config
	err := viper.UnmarshalKey("monitor", &config)
	if err != nil {
		logrus.Fatalf("Can't parse server config: %s", err)
	}

	c := cron.New()
	for _, website := range config.WebSites {
		w := website
		logrus.Infof("add new website: %s ^ %s\n", w.Address, config.Cron)
		err := c.AddFunc(config.Cron, func() {
			err := check(w, config.BeforeTime, config.Timeout)
			if err != nil {
				alarm.Alarm(err.Error())
			}
		})
		if err != nil {
			logrus.Fatal(err)
		}
	}
	c.Start()
	logrus.Info("cert monitor started.")
	select {}
}
