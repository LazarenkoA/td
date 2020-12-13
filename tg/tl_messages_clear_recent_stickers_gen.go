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

// MessagesClearRecentStickersRequest represents TL type `messages.clearRecentStickers#8999602d`.
// Clear recent stickers
//
// See https://core.telegram.org/method/messages.clearRecentStickers for reference.
type MessagesClearRecentStickersRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Set this flag to clear the list of stickers recently attached to photo or video files
	Attached bool
}

// MessagesClearRecentStickersRequestTypeID is TL type id of MessagesClearRecentStickersRequest.
const MessagesClearRecentStickersRequestTypeID = 0x8999602d

// Encode implements bin.Encoder.
func (c *MessagesClearRecentStickersRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.clearRecentStickers#8999602d as nil")
	}
	b.PutID(MessagesClearRecentStickersRequestTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.clearRecentStickers#8999602d: field flags: %w", err)
	}
	return nil
}

// SetAttached sets value of Attached conditional field.
func (c *MessagesClearRecentStickersRequest) SetAttached(value bool) {
	if value {
		c.Flags.Set(0)
	} else {
		c.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (c *MessagesClearRecentStickersRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.clearRecentStickers#8999602d to nil")
	}
	if err := b.ConsumeID(MessagesClearRecentStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.clearRecentStickers#8999602d: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.clearRecentStickers#8999602d: field flags: %w", err)
		}
	}
	c.Attached = c.Flags.Has(0)
	return nil
}

// Ensuring interfaces in compile-time for MessagesClearRecentStickersRequest.
var (
	_ bin.Encoder = &MessagesClearRecentStickersRequest{}
	_ bin.Decoder = &MessagesClearRecentStickersRequest{}
)

// MessagesClearRecentStickers invokes method messages.clearRecentStickers#8999602d returning error if any.
// Clear recent stickers
//
// See https://core.telegram.org/method/messages.clearRecentStickers for reference.
func (c *Client) MessagesClearRecentStickers(ctx context.Context, request *MessagesClearRecentStickersRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
