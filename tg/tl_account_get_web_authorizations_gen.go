// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// AccountGetWebAuthorizationsRequest represents TL type `account.getWebAuthorizations#182e6d6f`.
// Get web login widget authorizations
//
// See https://core.telegram.org/method/account.getWebAuthorizations for reference.
type AccountGetWebAuthorizationsRequest struct {
}

// AccountGetWebAuthorizationsRequestTypeID is TL type id of AccountGetWebAuthorizationsRequest.
const AccountGetWebAuthorizationsRequestTypeID = 0x182e6d6f

// Encode implements bin.Encoder.
func (g *AccountGetWebAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWebAuthorizations#182e6d6f as nil")
	}
	b.PutID(AccountGetWebAuthorizationsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetWebAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWebAuthorizations#182e6d6f to nil")
	}
	if err := b.ConsumeID(AccountGetWebAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getWebAuthorizations#182e6d6f: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetWebAuthorizationsRequest.
var (
	_ bin.Encoder = &AccountGetWebAuthorizationsRequest{}
	_ bin.Decoder = &AccountGetWebAuthorizationsRequest{}
)

// AccountGetWebAuthorizations invokes method account.getWebAuthorizations#182e6d6f returning error if any.
// Get web login widget authorizations
//
// See https://core.telegram.org/method/account.getWebAuthorizations for reference.
func (c *Client) AccountGetWebAuthorizations(ctx context.Context) (*AccountWebAuthorizations, error) {
	var result AccountWebAuthorizations

	request := &AccountGetWebAuthorizationsRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
