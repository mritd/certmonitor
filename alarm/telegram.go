package alarm

import (
	"fmt"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
	"time"
)

type TelegramConfig struct {
	Api   string `yaml:"api"`
	Token string `yaml:"token"`
}

func TelegramConfigExample() *TelegramConfig {
	return &TelegramConfig{
		"https://api.telegram.org",
		"token_example",
	}
}

func (cfg *TelegramConfig) Send(targets []string, message string) {
	bot, err := tb.NewBot(tb.Settings{
		URL:    cfg.Api,
		Token:  cfg.Token,
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		logrus.Errorf("Telegram alarm send failed: %s", err)
		return
	}

	opt := &tb.SendOptions{ParseMode: tb.ModeMarkdown}
	for _, t := range targets {
		to, err := strconv.Atoi(t)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
		_, err = bot.Send(tb.ChatID(to), fmt.Sprintf("```\n%s\n```", message), opt)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
	}
}
