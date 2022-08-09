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

// MessagesSetChatAvailableReactionsRequest represents TL type `messages.setChatAvailableReactions#14050ea6`.
// Change the set of message reactions »¹ that can be used in a certain group,
// supergroup or channel
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// See https://core.telegram.org/method/messages.setChatAvailableReactions for reference.
type MessagesSetChatAvailableReactionsRequest struct {
	// Group where to apply changes
	Peer InputPeerClass
	// Allowed reaction emojis
	AvailableReactions []string
}

// MessagesSetChatAvailableReactionsRequestTypeID is TL type id of MessagesSetChatAvailableReactionsRequest.
const MessagesSetChatAvailableReactionsRequestTypeID = 0x14050ea6

// Ensuring interfaces in compile-time for MessagesSetChatAvailableReactionsRequest.
var (
	_ bin.Encoder     = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.Decoder     = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.BareEncoder = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.BareDecoder = &MessagesSetChatAvailableReactionsRequest{}
)

func (s *MessagesSetChatAvailableReactionsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.AvailableReactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetChatAvailableReactionsRequest) String() string {
	if s == nil {
		return "MessagesSetChatAvailableReactionsRequest(nil)"
	}
	type Alias MessagesSetChatAvailableReactionsRequest
	return fmt.Sprintf("MessagesSetChatAvailableReactionsRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetChatAvailableReactionsRequest from given interface.
func (s *MessagesSetChatAvailableReactionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetAvailableReactions() (value []string)
}) {
	s.Peer = from.GetPeer()
	s.AvailableReactions = from.GetAvailableReactions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetChatAvailableReactionsRequest) TypeID() uint32 {
	return MessagesSetChatAvailableReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetChatAvailableReactionsRequest) TypeName() string {
	return "messages.setChatAvailableReactions"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetChatAvailableReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setChatAvailableReactions",
		ID:   MessagesSetChatAvailableReactionsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "AvailableReactions",
			SchemaName: "available_reactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetChatAvailableReactionsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setChatAvailableReactions#14050ea6 as nil")
	}
	b.PutID(MessagesSetChatAvailableReactionsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetChatAvailableReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setChatAvailableReactions#14050ea6 as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#14050ea6: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#14050ea6: field peer: %w", err)
	}
	b.PutVectorHeader(len(s.AvailableReactions))
	for _, v := range s.AvailableReactions {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetChatAvailableReactionsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setChatAvailableReactions#14050ea6 to nil")
	}
	if err := b.ConsumeID(MessagesSetChatAvailableReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setChatAvailableReactions#14050ea6: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetChatAvailableReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setChatAvailableReactions#14050ea6 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatAvailableReactions#14050ea6: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatAvailableReactions#14050ea6: field available_reactions: %w", err)
		}

		if headerLen > 0 {
			s.AvailableReactions = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messages.setChatAvailableReactions#14050ea6: field available_reactions: %w", err)
			}
			s.AvailableReactions = append(s.AvailableReactions, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetChatAvailableReactionsRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetAvailableReactions returns value of AvailableReactions field.
func (s *MessagesSetChatAvailableReactionsRequest) GetAvailableReactions() (value []string) {
	if s == nil {
		return
	}
	return s.AvailableReactions
}

// MessagesSetChatAvailableReactions invokes method messages.setChatAvailableReactions#14050ea6 returning error if any.
// Change the set of message reactions »¹ that can be used in a certain group,
// supergroup or channel
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// Possible errors:
//
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_NOT_MODIFIED: The pinned message wasn't modified.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.setChatAvailableReactions for reference.
func (c *Client) MessagesSetChatAvailableReactions(ctx context.Context, request *MessagesSetChatAvailableReactionsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
