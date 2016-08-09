package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/urfave/cli"
)

var usage = `
Usage here
`

var version = "sample"
var Subscribed map[string]byte

func init() {
	log.SetLevel(log.WarnLevel)
}

// MQTT operations
func getRandomClientId() string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, 9)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return "mqttwrk-" + string(bytes)
}

func Connect(opts *MQTT.ClientOptions) (*MQTT.Client, error) {
	m := MQTT.NewClient(opts)

	log.Info("connecting...")

	if token := m.Connect(); token.Wait() && token.Error() != nil {
		return m, token.Error()
	}
	return m, nil
}

func Publish(m *MQTT.Client, topic string, payload []byte, qos int, retain bool, sync bool) error {
	token := m.Publish(topic, byte(qos), retain, payload)

	return nil
}

func Disconnect(m *MQTT.Client) error {
	if m.IsConnected() {
		m.Disconnect(20)
		log.Info("client disconnected")
	}
	return nil
}

func SubscribeOnConnect(client *MQTT.Client) {
	log.Infof("client connected")

	if len(Subscribed) > 0 {
		token := client.SubscribeMultiple(Subscribed, OnMessageReceived)
		token.Wait()
		if token.Error() != nil {
			log.Error(token.Error())
		}
	}
}

func ConnectionLost(client *MQTT.Client, reason error) {
	log.Errorf("client disconnected: %s", reason)
}

func OnMessageReceived(client *MQTT.Client, message MQTT.Message) {
	log.Infof("topic:%s / msg:%s", message.Topic(), message.Payload())
	fmt.Println(string(message.Payload()))
}

// connects MQTT broker
func connect(opts *MQTT.ClientOptions, subscribed map[string]byte) (*MQTT.Client, error) {

	client := MQTT.NewClient(opts)
	client.Subscribed = subscribed

	opts.SetOnConnectHandler(SubscribeOnConnect)
	opts.SetConnectionLostHandler(ConnectionLost)

	_, err := Connect(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// newOption returns ClientOptions via parsing command line options.
func newOption(c *cli.Context) (*MQTT.ClientOptions, error) {
	opts := MQTT.NewClientOptions()

	host := c.String("host")
	port := c.Int("p")

	clientId := getRandomClientId()
	opts.SetClientID(clientId)

	scheme := "tcp"
	brokerUri := fmt.Sprintf("%s://%s:%d", scheme, host, port)
	log.Infof("Broker URI: %s", brokerUri)
	opts.AddBroker(brokerUri)

	opts.SetAutoReconnect(true)
	return opts, nil
}

// pubsubloop is a func of pub-sub event loop
func pubsubloop(c *cli.Context) error {
	opts, err := newOption(c)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	qos := 0
	subtopic := c.String("sub")
	if subtopic == "" {
		log.Errorf("Please specify sub topic")
		os.Exit(1)
	}
	log.Infof("Sub Topic: %s", subtopic)
	pubtopic := c.String("pub")
	if pubtopic == "" {
		log.Errorf("Please specify pub topic")
		os.Exit(1)
	}
	log.Infof("Pub Topic: %s", pubtopic)
	retain := c.Bool("r")

	subscribed := map[string]byte{
		subtopic: byte(0),
	}

	client, err := connect(c, opts, subscribed)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	go func() {
		// Read from Stdin and publish
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			err = client.Publish(pubtopic, []byte(scanner.Text()), qos, retain, false)
			if err != nil {
				log.Error(err)
			}
		}
	}()

	// while loop
	for {
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "mqttworkerforenocean-sample"
	app.Usage = "worker -c config-file"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 1883,
			Usage: "port number of broker",
		},
		cli.StringFlag{
			Name:  "host, h",
			Value: "localhost",
			Usage: "broker hostname",
		},
		cli.StringFlag{
			Name:   "u,user",
			Value:  "",
			Usage:  "provide a username",
			EnvVar: "USERNAME"},
		cli.StringFlag{
			Name:   "P,password",
			Value:  "",
			Usage:  "password",
			EnvVar: "PASSWORD"},
		cli.StringFlag{
			Name:  "sub",
			Value: "prefix/gateway/enocean/publish",
			Usage: "subscribe topic",
		},
		cli.StringFlag{
			Name:  "pub",
			Value: "prefix/worker/enocean/publish",
			Usage: "publish parsed data topic",
		},
		cli.IntFlag{
			Name:  "q",
			Value: 0,
			Usage: "Qos level to publish",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "run in verbose mode",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "loop",
			Usage:  "loop",
			Flags:  app.Flags,
			Action: "pubsubloop",
		},
	}

	app.Run(os.Args)
}
