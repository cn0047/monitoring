package ERR

import (
	"fmt"

	"go-app/common/vo"
)

func VOInvalid(vObj vo.InterfaceVO) error {
	return fmt.Errorf(`invalid ValueObject %s`, vObj.GetName())
}

func Queue(action string, queueName string, err error) error {
	return fmt.Errorf(`failed to perform "%s" to %s: %s`, action, queueName, err)
}

func Ping(err error) error {
	return fmt.Errorf(`failed to perform ping: %s`, err)
}
