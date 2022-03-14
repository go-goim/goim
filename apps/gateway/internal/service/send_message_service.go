package service

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/go-kratos/kratos/v2/log"

	messagev1 "github.com/yusank/goim/api/message/v1"
	"github.com/yusank/goim/apps/gateway/internal/app"
	"github.com/yusank/goim/pkg/mq"
)

type SendMessageService struct{}

var (
	sendMessageService     *SendMessageService
	sendMessageServiceOnce sync.Once
)

func GetSendMessageService() *SendMessageService {
	sendMessageServiceOnce.Do(func() {
		sendMessageService = new(SendMessageService)
	})

	return sendMessageService
}

func (s *SendMessageService) SendMessage(ctx context.Context, msg *messagev1.SendMessageReq) error {
	// check req params

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	rs, err := app.GetApplication().Producer.SendSync(ctx, mq.NewMessage("def_topic", b))
	if err != nil {
		return err
	}

	log.Info(rs.String())
	return nil
}
