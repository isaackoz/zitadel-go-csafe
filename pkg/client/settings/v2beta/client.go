package v2beta

import (
	"context"

	"github.com/zitadel/zitadel-go/v3/pkg/client/zitadel"
	settings "github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/settings/v2beta"
)

type Client struct {
	Connection *zitadel.Connection
	settings.SettingsServiceClient
}

func NewClient(ctx context.Context, issuer, api string, scopes []string, options ...zitadel.Option) (*Client, error) {
	conn, err := zitadel.NewConnection(ctx, issuer, api, scopes, options...)
	if err != nil {
		return nil, err
	}

	return &Client{
		Connection:            conn,
		SettingsServiceClient: settings.NewSettingsServiceClient(conn.ClientConn),
	}, nil
}
