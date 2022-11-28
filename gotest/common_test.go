package gotest

import (
	"auto-deployer/utils/log"
	"fmt"
	"testing"
)

func Test_Common(t *testing.T) {
	log.Info().Msg(fmt.Sprintf("%s%%", "10"))
}
