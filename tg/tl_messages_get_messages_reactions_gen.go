// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessagesGetMessagesReactionsRequest represents TL type `messages.getMessagesReactions#8bba90e6`.
// Get message reactions »¹
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// See https://core.telegram.org/method/messages.getMessagesReactions for reference.
type MessagesGetMessagesReactionsRequest struct {
	// Peer
	Peer InputPeerClass
	// Message IDs
	ID []int
}

// MessagesGetMessagesReactionsRequestTypeID is TL type id of MessagesGetMessagesReactionsRequest.
const MessagesGetMessagesReactionsRequestTypeID = 0x8bba90e6

// Ensuring interfaces in compile-time for MessagesGetMessagesReactionsRequest.
var (
	_ bin.Encoder     = &MessagesGetMessagesReactionsRequest{}
	_ bin.Decoder     = &MessagesGetMessagesReactionsRequest{}
	_ bin.BareEncoder = &MessagesGetMessagesReactionsRequest{}
	_ bin.BareDecoder = &MessagesGetMessagesReactionsRequest{}
)

func (g *MessagesGetMessagesReactionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetMessagesReactionsRequest) String() string {
	if g == nil {
		return "MessagesGetMessagesReactionsRequest(nil)"
	}
	type Alias MessagesGetMessagesReactionsRequest
	return fmt.Sprintf("MessagesGetMessagesReactionsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetMessagesReactionsRequest from given interface.
func (g *MessagesGetMessagesReactionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetMessagesReactionsRequest) TypeID() uint32 {
	return MessagesGetMessagesReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetMessagesReactionsRequest) TypeName() string {
	return "messages.getMessagesReactions"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetMessagesReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getMessagesReactions",
		ID:   MessagesGetMessagesReactionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetMessagesReactionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessagesReactions#8bba90e6 as nil")
	}
	b.PutID(MessagesGetMessagesReactionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetMessagesReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessagesReactions#8bba90e6 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getMessagesReactions#8bba90e6: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessagesReactions#8bba90e6: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessagesReactionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessagesReactions#8bba90e6 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessagesReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessagesReactions#8bba90e6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetMessagesReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessagesReactions#8bba90e6 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessagesReactions#8bba90e6: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessagesReactions#8bba90e6: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getMessagesReactions#8bba90e6: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetMessagesReactionsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *MessagesGetMessagesReactionsRequest) GetID() (value []int) {
	if g == nil {
		return
	}
	return g.ID
}

// MessagesGetMessagesReactions invokes method messages.getMessagesReactions#8bba90e6 returning error if any.
// Get message reactions »¹
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// See https://core.telegram.org/method/messages.getMessagesReactions for reference.
func (c *Client) MessagesGetMessagesReactions(ctx context.Context, request *MessagesGetMessagesReactionsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
