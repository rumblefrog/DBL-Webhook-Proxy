package graceful

import (
	"context"

	"github.com/rumblefrog/DBL-Webhook-Proxy/listener"
)

func GracefulExit() {
	if listener.Server != nil {
		listener.Server.Shutdown(context.Background())
	}
}
