package error_tracking

import (
	"context"
)

type ErrorTrackingItf interface {
	CaptureError(ctx context.Context, err error)
	CaptureMessage(ctx context.Context, message string)
	SetContext(ctx context.Context, key, value string)
	SetExtra(ctx context.Context, key string, value any)
	SetExtras(ctx context.Context, extra map[string]any)
}
