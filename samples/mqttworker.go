package main

import (
	log "github.com/Sirupsen/logrus"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/urfave/cli"
)

var usage = `
Usage here
`

var version = "0.1"

func init() {
	log.SetLevel(log.WarnLevel)
	log.SetOutput(colorable.NewColorableStdout())
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

// Connects connect to the MQTT broker with Options.
func (m *MQTTClient) Connect() (MQTT.Client, error) {

	m.Client = MQTT.NewClient(m.Opts)

	log.Info("connecting...")

	if token := m.Client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return m.Client, nil
}

func (m *MQTTClient) Publish(topic string, payload []byte, qos int, retain bool, sync bool) error {
	token := m.Client.Publish(topic, byte(qos), retain, payload)

	if sync == true {
		token.Wait()
	}

	return token.Error()
}

func (m *MQTTClient) Disconnect() error {
	if m.Client.IsConnected() {
		m.Client.Disconnect(20)
		log.Info("client disconnected")
	}
	return nil
}

func (m *MQTTClient) SubscribeOnConnect(client MQTT.Client) {
	log.Infof("client connected")

	if len(m.Subscribed) > 0 {
		token := client.SubscribeMultiple(m.Subscribed, m.onMessageReceived)
		token.Wait()
		if token.Error() != nil {
			log.Error(token.Error())
		}
	}
}

func (m *MQTTClient) ConnectionLost(client MQTT.Client, reason error) {
	log.Errorf("client disconnected: %s", reason)
}

func (m *MQTTClient) onMessageReceived(client MQTT.Client, message MQTT.Message) {
	log.Infof("topic:%s / msg:%s", message.Topic(), message.Payload())
	fmt.Println(string(message.Payload()))
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
	app = cli.NewApp()
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

	cli.VersionPrinter = printVersion

	app.Action = Action
	app.Run(os.Args)
}
