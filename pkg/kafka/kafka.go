package kafka

import (
	"github.com/segmentio/kafka-go"
	"github.com/zhf0439/im-server/pkg/config"
)

type Client struct {
	Writer *kafka.Writer
	Config config.KafkaConfig
}

func New(cfg config.KafkaConfig) *Client {
	return &Client{
		Config: cfg,
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(cfg.Brokers...),
			Topic:    cfg.TopicMessage,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (c *Client) Close() error {
	if c == nil || c.Writer == nil {
		return nil
	}
	return c.Writer.Close()
}
