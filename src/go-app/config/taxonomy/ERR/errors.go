package ERR

import (
	"fmt"

	"go-app/interface/vo"
)

// VOInvalid represents error for invalid ValueObject.
func VOInvalid(vObj vo.InterfaceVO) error {
	return fmt.Errorf(`invalid ValueObject %s`, vObj.GetName())
}

// Queue represents error related to queue actions.
func Queue(action string, queueName string, err error) error {
	return fmt.Errorf(`failed to perform "%s" to %s: %s`, action, queueName, err)
}

// Ping represents error related to ping service.
func Ping(err error) error {
	return fmt.Errorf(`failed to perform ping: %s`, err)
}

// Sys represents generalized internal golang error.
func Sys(err error) error {
	return fmt.Errorf(`got unexpected sys error: %s`, err)
}
