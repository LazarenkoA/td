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

// EditInlineMessageCaptionRequest represents TL type `editInlineMessageCaption#d2a446b7`.
type EditInlineMessageCaptionRequest struct {
	// Inline message identifier
	InlineMessageID string
	// The new message reply markup; pass null if none
	ReplyMarkup ReplyMarkupClass
	// New message content caption; pass null to remove caption;
	// 0-GetOption("message_caption_length_max") characters
	Caption FormattedText
}

// EditInlineMessageCaptionRequestTypeID is TL type id of EditInlineMessageCaptionRequest.
const EditInlineMessageCaptionRequestTypeID = 0xd2a446b7

// Ensuring interfaces in compile-time for EditInlineMessageCaptionRequest.
var (
	_ bin.Encoder     = &EditInlineMessageCaptionRequest{}
	_ bin.Decoder     = &EditInlineMessageCaptionRequest{}
	_ bin.BareEncoder = &EditInlineMessageCaptionRequest{}
	_ bin.BareDecoder = &EditInlineMessageCaptionRequest{}
)

func (e *EditInlineMessageCaptionRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.InlineMessageID == "") {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}
	if !(e.Caption.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditInlineMessageCaptionRequest) String() string {
	if e == nil {
		return "EditInlineMessageCaptionRequest(nil)"
	}
	type Alias EditInlineMessageCaptionRequest
	return fmt.Sprintf("EditInlineMessageCaptionRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditInlineMessageCaptionRequest) TypeID() uint32 {
	return EditInlineMessageCaptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditInlineMessageCaptionRequest) TypeName() string {
	return "editInlineMessageCaption"
}

// TypeInfo returns info about TL type.
func (e *EditInlineMessageCaptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editInlineMessageCaption",
		ID:   EditInlineMessageCaptionRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineMessageID",
			SchemaName: "inline_message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
		{
			Name:       "Caption",
			SchemaName: "caption",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditInlineMessageCaptionRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageCaption#d2a446b7 as nil")
	}
	b.PutID(EditInlineMessageCaptionRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditInlineMessageCaptionRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageCaption#d2a446b7 as nil")
	}
	b.PutString(e.InlineMessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageCaption#d2a446b7: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageCaption#d2a446b7: field reply_markup: %w", err)
	}
	if err := e.Caption.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageCaption#d2a446b7: field caption: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditInlineMessageCaptionRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageCaption#d2a446b7 to nil")
	}
	if err := b.ConsumeID(EditInlineMessageCaptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editInlineMessageCaption#d2a446b7: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditInlineMessageCaptionRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageCaption#d2a446b7 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageCaption#d2a446b7: field inline_message_id: %w", err)
		}
		e.InlineMessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageCaption#d2a446b7: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	{
		if err := e.Caption.Decode(b); err != nil {
			return fmt.Errorf("unable to decode editInlineMessageCaption#d2a446b7: field caption: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditInlineMessageCaptionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageCaption#d2a446b7 as nil")
	}
	b.ObjStart()
	b.PutID("editInlineMessageCaption")
	b.Comma()
	b.FieldStart("inline_message_id")
	b.PutString(e.InlineMessageID)
	b.Comma()
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageCaption#d2a446b7: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageCaption#d2a446b7: field reply_markup: %w", err)
	}
	b.Comma()
	b.FieldStart("caption")
	if err := e.Caption.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageCaption#d2a446b7: field caption: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditInlineMessageCaptionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageCaption#d2a446b7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editInlineMessageCaption"); err != nil {
				return fmt.Errorf("unable to decode editInlineMessageCaption#d2a446b7: %w", err)
			}
		case "inline_message_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageCaption#d2a446b7: field inline_message_id: %w", err)
			}
			e.InlineMessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageCaption#d2a446b7: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		case "caption":
			if err := e.Caption.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode editInlineMessageCaption#d2a446b7: field caption: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInlineMessageID returns value of InlineMessageID field.
func (e *EditInlineMessageCaptionRequest) GetInlineMessageID() (value string) {
	if e == nil {
		return
	}
	return e.InlineMessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditInlineMessageCaptionRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	if e == nil {
		return
	}
	return e.ReplyMarkup
}

// GetCaption returns value of Caption field.
func (e *EditInlineMessageCaptionRequest) GetCaption() (value FormattedText) {
	if e == nil {
		return
	}
	return e.Caption
}

// EditInlineMessageCaption invokes method editInlineMessageCaption#d2a446b7 returning error if any.
func (c *Client) EditInlineMessageCaption(ctx context.Context, request *EditInlineMessageCaptionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
