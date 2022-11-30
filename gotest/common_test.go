package gotest

import (
	"fmt"
	"go-deployer/utils/log"
	"testing"
)

func Test_Common(t *testing.T) {
	log.Info().Msg(fmt.Sprintf("%s%%", "10"))
}
