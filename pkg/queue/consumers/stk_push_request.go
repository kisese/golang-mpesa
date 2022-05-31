package consumers

import (
	"encoding/json"
	"fmt"
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/forms"
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/service"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
	"github.com/kisese/golang_mpesa/pkg/queue/utils"
	"github.com/streadway/amqp"
	"os"
)

func ProcessSTKPushRequest(ch *amqp.Channel) {

	queue := os.Getenv("STK_PUSH_REQUESTS_QUEUE")
	msgs := utils.Consume(ch, queue)

	infrastructure.Log.Debugw("Queue Listening ", "queue", queue)

	go func() {
		for d := range msgs {
			var payload forms.STKRequest
			message := fmt.Sprintf("%s", d.Body)

			infrastructure.Log.Debugw("RabbitMQ Consumer Received Message " + message)

			err := json.Unmarshal([]byte(message), &payload)
			if err != nil {
				infrastructure.Log.Errorw("Payload unmarshall error ", "error", err, "queue", queue)
			}

			stkPushService := service.NewStkRequestService(payload)
			stkPushService.ProcessSTKPush()
		}
	}()
}
