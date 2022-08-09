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

// AttachMenuBotsBot represents TL type `attachMenuBotsBot#93bf667f`.
//
// See https://core.telegram.org/constructor/attachMenuBotsBot for reference.
type AttachMenuBotsBot struct {
	//
	Bot AttachMenuBot
	//
	Users []UserClass
}

// AttachMenuBotsBotTypeID is TL type id of AttachMenuBotsBot.
const AttachMenuBotsBotTypeID = 0x93bf667f

// Ensuring interfaces in compile-time for AttachMenuBotsBot.
var (
	_ bin.Encoder     = &AttachMenuBotsBot{}
	_ bin.Decoder     = &AttachMenuBotsBot{}
	_ bin.BareEncoder = &AttachMenuBotsBot{}
	_ bin.BareDecoder = &AttachMenuBotsBot{}
)

func (a *AttachMenuBotsBot) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Bot.Zero()) {
		return false
	}
	if !(a.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuBotsBot) String() string {
	if a == nil {
		return "AttachMenuBotsBot(nil)"
	}
	type Alias AttachMenuBotsBot
	return fmt.Sprintf("AttachMenuBotsBot%+v", Alias(*a))
}

// FillFrom fills AttachMenuBotsBot from given interface.
func (a *AttachMenuBotsBot) FillFrom(from interface {
	GetBot() (value AttachMenuBot)
	GetUsers() (value []UserClass)
}) {
	a.Bot = from.GetBot()
	a.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuBotsBot) TypeID() uint32 {
	return AttachMenuBotsBotTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuBotsBot) TypeName() string {
	return "attachMenuBotsBot"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuBotsBot) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuBotsBot",
		ID:   AttachMenuBotsBotTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AttachMenuBotsBot) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBotsBot#93bf667f as nil")
	}
	b.PutID(AttachMenuBotsBotTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuBotsBot) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBotsBot#93bf667f as nil")
	}
	if err := a.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode attachMenuBotsBot#93bf667f: field bot: %w", err)
	}
	b.PutVectorHeader(len(a.Users))
	for idx, v := range a.Users {
		if v == nil {
			return fmt.Errorf("unable to encode attachMenuBotsBot#93bf667f: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode attachMenuBotsBot#93bf667f: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuBotsBot) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBotsBot#93bf667f to nil")
	}
	if err := b.ConsumeID(AttachMenuBotsBotTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuBotsBot#93bf667f: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuBotsBot) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBotsBot#93bf667f to nil")
	}
	{
		if err := a.Bot.Decode(b); err != nil {
			return fmt.Errorf("unable to decode attachMenuBotsBot#93bf667f: field bot: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBotsBot#93bf667f: field users: %w", err)
		}

		if headerLen > 0 {
			a.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode attachMenuBotsBot#93bf667f: field users: %w", err)
			}
			a.Users = append(a.Users, value)
		}
	}
	return nil
}

// GetBot returns value of Bot field.
func (a *AttachMenuBotsBot) GetBot() (value AttachMenuBot) {
	if a == nil {
		return
	}
	return a.Bot
}

// GetUsers returns value of Users field.
func (a *AttachMenuBotsBot) GetUsers() (value []UserClass) {
	if a == nil {
		return
	}
	return a.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (a *AttachMenuBotsBot) MapUsers() (value UserClassArray) {
	return UserClassArray(a.Users)
}
