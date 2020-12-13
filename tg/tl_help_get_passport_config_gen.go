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

// HelpGetPassportConfigRequest represents TL type `help.getPassportConfig#c661ad08`.
// Get passport configuration
//
// See https://core.telegram.org/method/help.getPassportConfig for reference.
type HelpGetPassportConfigRequest struct {
	// Hash for pagination, for more info click here
	Hash int
}

// HelpGetPassportConfigRequestTypeID is TL type id of HelpGetPassportConfigRequest.
const HelpGetPassportConfigRequestTypeID = 0xc661ad08

// Encode implements bin.Encoder.
func (g *HelpGetPassportConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getPassportConfig#c661ad08 as nil")
	}
	b.PutID(HelpGetPassportConfigRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetPassportConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getPassportConfig#c661ad08 to nil")
	}
	if err := b.ConsumeID(HelpGetPassportConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getPassportConfig#c661ad08: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.getPassportConfig#c661ad08: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetPassportConfigRequest.
var (
	_ bin.Encoder = &HelpGetPassportConfigRequest{}
	_ bin.Decoder = &HelpGetPassportConfigRequest{}
)

// HelpGetPassportConfig invokes method help.getPassportConfig#c661ad08 returning error if any.
// Get passport configuration
//
// See https://core.telegram.org/method/help.getPassportConfig for reference.
func (c *Client) HelpGetPassportConfig(ctx context.Context, hash int) (HelpPassportConfigClass, error) {
	var result HelpPassportConfigBox

	request := &HelpGetPassportConfigRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PassportConfig, nil
}
