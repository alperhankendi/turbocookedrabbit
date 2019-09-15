package models

import (
	"fmt"

	"github.com/streadway/amqp"
)

// ErrorMessage allow for you to replay a message that was returned.
type ErrorMessage struct {
	Code    int
	Reason  string
	Server  bool
	Recover bool
}

// NewErrorMessage creates a new ErrorMessage.
func NewErrorMessage(amqpError *amqp.Error) *ErrorMessage {

	return &ErrorMessage{
		Code:    amqpError.Code,
		Reason:  amqpError.Reason,
		Server:  amqpError.Server,
		Recover: amqpError.Recover,
	}
}

// Error allows you to quickly log the ErrorMessage struct as a string.
func (em *ErrorMessage) Error() string {
	return fmt.Sprintf("[ErrorCode: %d] Reason: %s \r\n[Server Initiated: %v]\r\n[Recoverable: %v]\r\n", em.Code, em.Reason, em.Server, em.Recover)
}
