package main

import (
	CLIENT "github.com/shirou/mqttcli"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/urfave/cli"
	)

var usage = `
Usage here
`

func init() {
}

// connects MQTT broker
func connect(opts *MQTT.ClientOptions, subscribed map[string]byte) (*MQTTClient, error) {

	client := &MQTTClient{Opts: opts}
	client.lock = new(sync.Mutex)
	client.Subscribed = subscribed

	opts.SetOnConnectHandler(client.SubscribeOnConnect)
	opts.SetConnectionLostHandler(client.ConnectionLost)

	_, err := client.Connect()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func pubsub(c *cli.Context) error {
	setDebugLevel(c)
	opts, err := NewOption(c)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	qos := c.Int("q")
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
	app = cli.NewApp()
	app.Name = "mqttworkerforenocean-sample"
	app.Usage = "worker -c config-file"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:	"port, p",
			Value:	1883,
			Usage:	"port number of broker",
		},
		cli.StringFlag{
			Name:   "host, h",
			Value:  "localhost",
			Usage:  "broker hostname",
		},
		cli.StringFlag{
			Name:	"sub, s",
			Value:	"prefix/gateway/enocean/publish",
			Usage:	"subscribe topic",
		}
		cli.BoolFlag{
			Name:  "d",
			Usage: "run in verbose mode",
		},
	}

	cli.VersionPrinter = printVersion

	app.Action = Action
	app.Run(os.Args)
}

