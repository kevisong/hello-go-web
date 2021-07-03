package debug

import (
	"fmt"
	"net/http"

	"github.com/KEVISONG/hello-go-web/pkg/config/debug"
	"github.com/sirupsen/logrus"
)

// Run runs debug server
func Run(c debug.Config) {
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", c.Port), nil)
	if err != nil {
		logrus.Error(err)
		return
	}
}
