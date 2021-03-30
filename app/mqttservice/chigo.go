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
var chigoReconnectingHandler mqtt.ReconnectHandler = func(c mqtt.Client, co *mqtt.ClientOptions) {
	app.LLog.Info("MQTT 客户端正在重连", co)
}

/*
	连接成功后的处理函数
*/
var chigoConnectHandler mqtt.OnConnectHandler = func(c mqtt.Client) {
	app.LLog.Info("chigo client 已经连接成功")
	//订阅Topic，需要在连接成功后订阅Topic，因为断开重连后需要重新订阅
	chigoSubTopics(c)
}

/*
	连接丢失时候的处理函数，连接丢失后会自动重新连接
*/
var chigoConnnectionLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, e error) {
	app.LLog.Error("chigo client 连接丢失 error:", e)
	//连接丢失后取消订阅Topic
	chigoUnsubTopics(c)
}

/*
	接收到MQTT信息后的处理函数
*/
var chigoMessageHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	app.LLog.Debug(fmt.Sprintf("chigo收到信息 topic->%s payload->%s", m.Topic(), m.Payload()))
}

/*
	订阅Topic
*/
func chigoSubTopics(c mqtt.Client) {
	//Topic 可以统一订阅到chigoMessageHandler处理，也可以针对单个topic订阅自己的Handler
	token := c.Subscribe("+/+/+/internalStatus", 1, chigoMessageHandler)
	token.Wait()

}

/*
 取消订阅
*/
func chigoUnsubTopics(c mqtt.Client) {
	token := c.Unsubscribe("+/+/+/internalStatus")
	token.Wait()
}

func chigoMQTT() {

	mClient := mqttc.GetInstance("chigo")
	//重新连接
	mClient.Opts.SetReconnectingHandler(chigoReconnectingHandler)
	//连接成功
	mClient.Opts.SetOnConnectHandler(chigoConnectHandler)
	//连接丢失
	mClient.Opts.SetConnectionLostHandler(chigoConnnectionLostHandler)

	err := mClient.Connect()
	if err != nil {
		panic(fmt.Sprintf("连接MQTTServer发生错误:%v", err))
	}

}

/*
	发布信息
*/
func ChigoPub(topic string, qos byte, retained bool, payload interface{}) error {
	mClient := mqttc.GetInstance("chigo")
	err := mClient.Connect()
	if err != nil {
		return err
	}
	token := mClient.PahoClient.Publish(topic, qos, retained, payload)
	token.Wait()
	return nil
}
