package mqttservice

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/mqttc"
)

/*
	重新连接事件的处理函数
*/
var mosquittoReconnectingHandler mqtt.ReconnectHandler = func(c mqtt.Client, co *mqtt.ClientOptions) {
	app.LLog.Info("MQTT 客户端正在重连", co)
}

/*
	连接成功后的处理函数
*/
var mosquittoConnectHandler mqtt.OnConnectHandler = func(c mqtt.Client) {
	app.LLog.Info("mosquitto client 已经连接成功")
	//订阅Topic，需要在连接成功后订阅Topic，因为断开重连后需要重新订阅
	mosquittoSubTopics(c)
}

/*
	连接丢失时候的处理函数，连接丢失后会自动重新连接
*/
var mosquittoConnnectionLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, e error) {
	app.LLog.Error("mosquitto client 连接丢失 error:", e)
	//连接丢失后取消订阅Topic
	mosquittoUnsubTopics(c)
}

/*
	接收到MQTT信息后的处理函数
*/
var mosquittoMessageHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	app.LLog.Debug(fmt.Sprintf("mosquitto收到信息 topic->%s payload->%s", m.Topic(), m.Payload()))
}

/*
	订阅Topic
*/
func mosquittoSubTopics(c mqtt.Client) {
	//Topic 可以统一订阅到mosquittoMessageHandler处理，也可以针对单个topic订阅自己的Handler
	token := c.Subscribe("topic/test", 1, mosquittoMessageHandler)
	token.Wait()

}

/*
 取消订阅
*/
func mosquittoUnsubTopics(c mqtt.Client) {
	token := c.Unsubscribe("topic/test")
	token.Wait()
}

func mosquittoMQTT() {

	mClient := mqttc.GetInstance("mosquitto")
	//重新连接
	mClient.Opts.SetReconnectingHandler(mosquittoReconnectingHandler)
	//连接成功
	mClient.Opts.SetOnConnectHandler(mosquittoConnectHandler)
	//连接丢失
	mClient.Opts.SetConnectionLostHandler(mosquittoConnnectionLostHandler)

	err := mClient.Connect()
	if err != nil {
		panic(fmt.Sprintf("连接MQTTServer发生错误:%v", err))
	}

}

/*
	发布信息
*/
func MosquittoPub(topic string, qos byte, retained bool, payload interface{}) error {
	mClient := mqttc.GetInstance("mosquitto")
	err := mClient.Connect()
	if err != nil {
		return err
	}
	token := mClient.PahoClient.Publish(topic, qos, retained, payload)
	token.Wait()
	return nil
}
