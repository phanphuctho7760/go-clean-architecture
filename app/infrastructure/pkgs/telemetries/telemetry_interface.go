package telemetries

import (
	"context"
)

type TelemetryItf interface {
	StartSpan(ctx context.Context, name string) (newCtx context.Context)
	EndSpan()
	GetTraceID() string
	GetSpanID() string
	RecordError(ctx context.Context, err error)
	SetStatusError(ctx context.Context, errStr string)
	SetStatusSuccess(ctx context.Context, message string)
	SetStatusUnset(ctx context.Context, message string)
	SetAttributeString(ctx context.Context, key string, value string)
}
