package helpers

import (
	"context"

	"github.com/phanphuctho7760/go-clean-architecture/app/utils/constants"
)

// GetTrackIDFromContext
// Get track id from context
func GetTrackIDFromContext(ctx context.Context) (trackID string) {
	trackID, ok := ctx.Value(constants.TrackIDKey).(string)
	if !ok {
		return ""
	}
	return
}

// MakeTrackIDIntoNewContext
// Make a context with track id get from old context
func MakeTrackIDIntoNewContext(oldCtx, newCtx context.Context) context.Context {
	return context.WithValue(newCtx, constants.TrackIDKey, GetTrackIDFromContext(oldCtx))
}
