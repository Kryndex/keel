package hipchat

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/daneharrigan/hipchat"
)

type XmppImplementer interface {
	Say(roomID, name, body string)
	Status(s string)
	Join(roomID, resource string)
	KeepAlive()
	Messages() <-chan *hipchat.Message
}

type HipchatClient struct {
	XmppImplementer
}

func connect(username, password string, connAttempts int) *HipchatClient {
	// c := &Client{}
	attempts := connAttempts
	for {
		if attempts == 0 {
			log.Errorf("Can not reach hipchat server after %d attempts", connAttempts)
			return nil
		}
		log.Info("bot.hipchat.connect: try to connect to hipchat")
		client, err := hipchat.NewClient(username, password, "bot", "plain")

		if err != nil {
			log.Errorf("bot.hipchat.connect: Error=%s", err)
			if err.Error() == "could not authenticate" {
				return nil
			}
		}
		if client != nil && err == nil {
			log.Info("Successfully connected to hipchat server")
			return &HipchatClient{client}
			// c.cli = client
			// return c
		}
		log.Debugln("Can not connect to hipcaht now, wait fo 30 seconds")
		time.Sleep(30 * time.Second)
		attempts--
	}
}
