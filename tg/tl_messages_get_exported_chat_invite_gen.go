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

// MessagesGetExportedChatInviteRequest represents TL type `messages.getExportedChatInvite#73746f5c`.
// Get info about a chat invite
//
// See https://core.telegram.org/method/messages.getExportedChatInvite for reference.
type MessagesGetExportedChatInviteRequest struct {
	// Chat
	Peer InputPeerClass
	// Invite link
	Link string
}

// MessagesGetExportedChatInviteRequestTypeID is TL type id of MessagesGetExportedChatInviteRequest.
const MessagesGetExportedChatInviteRequestTypeID = 0x73746f5c

// Ensuring interfaces in compile-time for MessagesGetExportedChatInviteRequest.
var (
	_ bin.Encoder     = &MessagesGetExportedChatInviteRequest{}
	_ bin.Decoder     = &MessagesGetExportedChatInviteRequest{}
	_ bin.BareEncoder = &MessagesGetExportedChatInviteRequest{}
	_ bin.BareDecoder = &MessagesGetExportedChatInviteRequest{}
)

func (g *MessagesGetExportedChatInviteRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetExportedChatInviteRequest) String() string {
	if g == nil {
		return "MessagesGetExportedChatInviteRequest(nil)"
	}
	type Alias MessagesGetExportedChatInviteRequest
	return fmt.Sprintf("MessagesGetExportedChatInviteRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetExportedChatInviteRequest from given interface.
func (g *MessagesGetExportedChatInviteRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetLink() (value string)
}) {
	g.Peer = from.GetPeer()
	g.Link = from.GetLink()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetExportedChatInviteRequest) TypeID() uint32 {
	return MessagesGetExportedChatInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetExportedChatInviteRequest) TypeName() string {
	return "messages.getExportedChatInvite"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetExportedChatInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getExportedChatInvite",
		ID:   MessagesGetExportedChatInviteRequestTypeID,
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
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetExportedChatInviteRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getExportedChatInvite#73746f5c as nil")
	}
	b.PutID(MessagesGetExportedChatInviteRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetExportedChatInviteRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getExportedChatInvite#73746f5c as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getExportedChatInvite#73746f5c: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getExportedChatInvite#73746f5c: field peer: %w", err)
	}
	b.PutString(g.Link)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetExportedChatInviteRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getExportedChatInvite#73746f5c to nil")
	}
	if err := b.ConsumeID(MessagesGetExportedChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getExportedChatInvite#73746f5c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetExportedChatInviteRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getExportedChatInvite#73746f5c to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExportedChatInvite#73746f5c: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExportedChatInvite#73746f5c: field link: %w", err)
		}
		g.Link = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetExportedChatInviteRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetLink returns value of Link field.
func (g *MessagesGetExportedChatInviteRequest) GetLink() (value string) {
	if g == nil {
		return
	}
	return g.Link
}

// MessagesGetExportedChatInvite invokes method messages.getExportedChatInvite#73746f5c returning error if any.
// Get info about a chat invite
//
// Possible errors:
//
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 INVITE_HASH_EXPIRED: The invite link has expired.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.getExportedChatInvite for reference.
func (c *Client) MessagesGetExportedChatInvite(ctx context.Context, request *MessagesGetExportedChatInviteRequest) (MessagesExportedChatInviteClass, error) {
	var result MessagesExportedChatInviteBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ExportedChatInvite, nil
}
