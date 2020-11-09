package platform

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

type ConnectionAmqp interface {
	Open() *amqp.Connection
	Reconnect(client *amqp.Connection) *amqp.Connection
}

type connectionStringAmqp struct {
	client string
	domain string
}

func InitializeAmqp(client, domain string) ConnectionAmqp {
	return &connectionStringAmqp{
		client: client,
		domain: domain,
	}
}

func (c *connectionStringAmqp) Open() *amqp.Connection {
	var conn *amqp.Connection
	var err error

	logFields := logrus.Fields{
		"platform": "amqp",
		"domain":   c.domain,
	}

	for i := 0; i < 5; i++ {
		logrus.WithFields(logFields).Info(fmt.Sprintf("Connecting to: %s", c.client))
		conn, err = amqp.Dial(c.client)

		if err != nil {
			logrus.WithFields(logFields).Fatal(err)

			// Wait 1 second
			time.Sleep(5 * time.Second)
		}

		break
	}

	go func() {
		conn = c.Reconnect(conn)
	}()

	return conn
}

// reconnect reconnects to server if the connection or a channel
// is closed unexpectedly. Normal shutdown is ignored. It tries
// maximum of 7200 times and sleeps half a second in between
// each try which equals to 1 hour.
func (c *connectionStringAmqp) Reconnect(client *amqp.Connection) *amqp.Connection {
	var conn *amqp.Connection
	var mux sync.RWMutex
WATCH:
	connErr := <-client.NotifyClose(make(chan *amqp.Error))
	if connErr != nil {
		log.Println("CRITICAL: Connection dropped, reconnecting")

		var err error

		for i := 0; i < 7200; i++ {
			mux.RLock()
			conn, err = amqp.Dial(c.client)
			mux.RUnlock()
			if err == nil {
				log.Println("INFO: Reconnected")

				goto WATCH
			}

			time.Sleep(500 * time.Millisecond)
		}

		log.Println(errors.Wrap(err, "CRITICAL: Failed to reconnect"))
		return nil
	} else {
		log.Println("INFO: Connection dropped normally, will not reconnect")
		return conn
	}
}
