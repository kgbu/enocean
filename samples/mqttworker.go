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

	"github.com/kgbu/enocean"
)

var usage = `
go run mqttworker.go loop --host hostname --sub subscribetopic --pub publishtopic
`

var version = "sample"
var Subscribed map[string]byte

func init() {
	log.SetLevel(log.DebugLevel)
}

// MQTT operations
func getRandomClientId() string {
	// 0, 1, 6, 9 like characters are removed to avoid mis-reading
	const alphanum = "234578ABCDEFGHJKLMNPQRSTUVWXYZacefghjkrstuvwxyz"
	var bytes = make([]byte, 9)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return "mqttwrk-" + string(bytes)
}


func Publish(m *MQTT.Client, topic string, payload []byte, qos int, retain bool, sync bool) error {
	token := m.Publish(topic, byte(qos), retain, payload)
	token.Wait()

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
	buf := message.Payload()
	log.Infof("topic:%s / msg:%s", message.Topic(), buf)

	// analyze temperture data
	err, c, e := enocean.NewESPData(buf)
	if err != nil {
		log.Errorf("ERROR: %v, parse failed on $v, cosumed %v", err, buf, c)
	}
	temp := (255 - int(e.PayloadData[2])) * 40.0 / 255
	err, m := enocean.GetManufacturerName(e.ManufacturerId)
	if err != nil {
		log.Errorf("ERROR: %v, manufacturer ID is wrong", e.ManufacturerId)
	}
        fmt.Printf("Data: %v : Teach in %v, made by %v, temperature %v\n", e, e.TeachIn, m, temp)
}

// connects MQTT broker
func connect(opts *MQTT.ClientOptions) (*MQTT.Client, error) {


	opts.SetOnConnectHandler(SubscribeOnConnect)
	opts.SetConnectionLostHandler(ConnectionLost)

	m := MQTT.NewClient(opts)

	log.Info("connecting...")

	if token := m.Connect(); token.Wait() && token.Error() != nil {
		return m, token.Error()
	}

	return m, nil
}

// newOption returns ClientOptions via parsing command line options.
func newOption(c *cli.Context) (*MQTT.ClientOptions, error) {
	opts := MQTT.NewClientOptions()

	host := c.String("host")
	port := c.Int("p")

	opts.SetClientID(getRandomClientId())

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

	Subscribed = map[string]byte{
		subtopic: byte(0),
	}

	client, err := connect(opts)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	go func() {
		// Read from Stdin and publish
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			err = Publish(client, pubtopic, []byte(scanner.Text()), qos, retain, false)
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
	app.Usage = "loop -p 9883"
	app.Version = version

	commonFlags := []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 1883,
			Usage: "port number of broker",
		},
		cli.StringFlag{
			Name:  "host",
			Value: "localhost",
			Usage: "broker hostname",
		},
		cli.StringFlag{
			Name:   "u,user",
			Value:  "",
			Usage:  "provide a username",
			EnvVar: "USERNAME",
		},
		cli.StringFlag{
			Name:   "P,password",
			Value:  "",
			Usage:  "password",
			EnvVar: "PASSWORD",
		},
		cli.StringFlag{
			Name:  "sub",
			Value: "enoceantemp/#",
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
			Flags:	commonFlags,
			Action: pubsubloop,
		},
	}

	app.Run(os.Args)
}
