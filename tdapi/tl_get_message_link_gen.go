// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// GetMessageLinkRequest represents TL type `getMessageLink#a0312f6f`.
type GetMessageLinkRequest struct {
	// Identifier of the chat to which the message belongs
	ChatID int64
	// Identifier of the message
	MessageID int64
	// If not 0, timestamp from which the video/audio/video note/voice note playing must
	// start, in seconds. The media can be in the message content or in its web page preview
	MediaTimestamp int32
	// Pass true to create a link for the whole media album
	ForAlbum bool
	// Pass true to create a link to the message as a channel post comment, or from a message
	// thread
	ForComment bool
}

// GetMessageLinkRequestTypeID is TL type id of GetMessageLinkRequest.
const GetMessageLinkRequestTypeID = 0xa0312f6f

// Ensuring interfaces in compile-time for GetMessageLinkRequest.
var (
	_ bin.Encoder     = &GetMessageLinkRequest{}
	_ bin.Decoder     = &GetMessageLinkRequest{}
	_ bin.BareEncoder = &GetMessageLinkRequest{}
	_ bin.BareDecoder = &GetMessageLinkRequest{}
)

func (g *GetMessageLinkRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.MediaTimestamp == 0) {
		return false
	}
	if !(g.ForAlbum == false) {
		return false
	}
	if !(g.ForComment == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessageLinkRequest) String() string {
	if g == nil {
		return "GetMessageLinkRequest(nil)"
	}
	type Alias GetMessageLinkRequest
	return fmt.Sprintf("GetMessageLinkRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageLinkRequest) TypeID() uint32 {
	return GetMessageLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageLinkRequest) TypeName() string {
	return "getMessageLink"
}

// TypeInfo returns info about TL type.
func (g *GetMessageLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageLink",
		ID:   GetMessageLinkRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "MediaTimestamp",
			SchemaName: "media_timestamp",
		},
		{
			Name:       "ForAlbum",
			SchemaName: "for_album",
		},
		{
			Name:       "ForComment",
			SchemaName: "for_comment",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessageLinkRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageLink#a0312f6f as nil")
	}
	b.PutID(GetMessageLinkRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageLinkRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageLink#a0312f6f as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	b.PutInt32(g.MediaTimestamp)
	b.PutBool(g.ForAlbum)
	b.PutBool(g.ForComment)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageLinkRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageLink#a0312f6f to nil")
	}
	if err := b.ConsumeID(GetMessageLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageLink#a0312f6f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageLinkRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageLink#a0312f6f to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field media_timestamp: %w", err)
		}
		g.MediaTimestamp = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field for_album: %w", err)
		}
		g.ForAlbum = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field for_comment: %w", err)
		}
		g.ForComment = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessageLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageLink#a0312f6f as nil")
	}
	b.ObjStart()
	b.PutID("getMessageLink")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.FieldStart("media_timestamp")
	b.PutInt32(g.MediaTimestamp)
	b.Comma()
	b.FieldStart("for_album")
	b.PutBool(g.ForAlbum)
	b.Comma()
	b.FieldStart("for_comment")
	b.PutBool(g.ForComment)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessageLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageLink#a0312f6f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessageLink"); err != nil {
				return fmt.Errorf("unable to decode getMessageLink#a0312f6f: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field message_id: %w", err)
			}
			g.MessageID = value
		case "media_timestamp":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field media_timestamp: %w", err)
			}
			g.MediaTimestamp = value
		case "for_album":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field for_album: %w", err)
			}
			g.ForAlbum = value
		case "for_comment":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageLink#a0312f6f: field for_comment: %w", err)
			}
			g.ForComment = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessageLinkRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetMessageLinkRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetMediaTimestamp returns value of MediaTimestamp field.
func (g *GetMessageLinkRequest) GetMediaTimestamp() (value int32) {
	if g == nil {
		return
	}
	return g.MediaTimestamp
}

// GetForAlbum returns value of ForAlbum field.
func (g *GetMessageLinkRequest) GetForAlbum() (value bool) {
	if g == nil {
		return
	}
	return g.ForAlbum
}

// GetForComment returns value of ForComment field.
func (g *GetMessageLinkRequest) GetForComment() (value bool) {
	if g == nil {
		return
	}
	return g.ForComment
}

// GetMessageLink invokes method getMessageLink#a0312f6f returning error if any.
func (c *Client) GetMessageLink(ctx context.Context, request *GetMessageLinkRequest) (*MessageLink, error) {
	var result MessageLink

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
