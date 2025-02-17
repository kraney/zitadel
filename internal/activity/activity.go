package activity

import (
	"context"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zitadel/logging"

	"github.com/zitadel/zitadel/internal/api/authz"
	http_utils "github.com/zitadel/zitadel/internal/api/http"
	"github.com/zitadel/zitadel/internal/api/info"
)

const (
	Activity = "activity"

	PathKey          = "zitadel-activity-path"
	RequestMethodKey = "zitadel-activity-request-method"
)

type TriggerMethod int

const (
	Unspecified TriggerMethod = iota
	ResourceAPI
	OIDCAccessToken
	OIDCRefreshToken
	SessionAPI
	SAMLResponse
)

func (t TriggerMethod) String() string {
	switch t {
	case Unspecified:
		return "unspecified"
	case ResourceAPI:
		return "resourceAPI"
	case OIDCRefreshToken:
		return "refreshToken"
	case OIDCAccessToken:
		return "accessToken"
	case SessionAPI:
		return "sessionAPI"
	case SAMLResponse:
		return "samlResponse"
	default:
		return "unknown"
	}
}

// Trigger is used to log a specific events for a user (e.g. session or oidc token creation)
func Trigger(ctx context.Context, orgID, userID string, trigger TriggerMethod) {
	ai := info.ActivityInfoFromContext(ctx)
	triggerLog(
		authz.GetInstance(ctx).InstanceID(),
		orgID,
		userID,
		http_utils.ComposedOrigin(ctx),
		trigger,
		ai.Method,
		ai.Path,
		ai.RequestMethod,
		"",
		"",
		authz.GetCtxData(ctx).SystemMemberships != nil,
	)
}

func TriggerGRPCWithContext(ctx context.Context, trigger TriggerMethod) {
	ai := info.ActivityInfoFromContext(ctx)
	triggerLog(
		authz.GetInstance(ctx).InstanceID(),
		authz.GetCtxData(ctx).OrgID,
		authz.GetCtxData(ctx).UserID,
		http_utils.ComposedOrigin(ctx),
		trigger,
		ai.Method,
		ai.Path,
		ai.RequestMethod,
		strconv.Itoa(int(ai.GRPCStatus)),
		strconv.Itoa(runtime.HTTPStatusFromCode(ai.GRPCStatus)),
		authz.GetCtxData(ctx).SystemMemberships != nil,
	)
}

func triggerLog(instanceID, orgID, userID, domain string, trigger TriggerMethod, method, path, requestMethod, grpcStatus, httpStatus string, isSystemUser bool) {
	logging.WithFields(
		"instance", instanceID,
		"org", orgID,
		"user", userID,
		"domain", domain,
		"trigger", trigger.String(),
		"method", method,
		"path", path,
		"grpcStatus", grpcStatus,
		"httpStatus", httpStatus,
		"requestMethod", requestMethod,
		"isSystemUser", isSystemUser,
	).Info(Activity)
}
