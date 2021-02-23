// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// MessagesChatAdminsWithInvites represents TL type `messages.chatAdminsWithInvites#b69b72d7`.
//
// See https://core.telegram.org/constructor/messages.chatAdminsWithInvites for reference.
type MessagesChatAdminsWithInvites struct {
	// Admins field of MessagesChatAdminsWithInvites.
	Admins []ChatAdminWithInvites `schemaname:"admins"`
	// Users field of MessagesChatAdminsWithInvites.
	Users []UserClass `schemaname:"users"`
}

// MessagesChatAdminsWithInvitesTypeID is TL type id of MessagesChatAdminsWithInvites.
const MessagesChatAdminsWithInvitesTypeID = 0xb69b72d7

func (c *MessagesChatAdminsWithInvites) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Admins == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesChatAdminsWithInvites) String() string {
	if c == nil {
		return "MessagesChatAdminsWithInvites(nil)"
	}
	type Alias MessagesChatAdminsWithInvites
	return fmt.Sprintf("MessagesChatAdminsWithInvites%+v", Alias(*c))
}

// FillFrom fills MessagesChatAdminsWithInvites from given interface.
func (c *MessagesChatAdminsWithInvites) FillFrom(from interface {
	GetAdmins() (value []ChatAdminWithInvites)
	GetUsers() (value []UserClass)
}) {
	c.Admins = from.GetAdmins()
	c.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *MessagesChatAdminsWithInvites) TypeID() uint32 {
	return MessagesChatAdminsWithInvitesTypeID
}

// SchemaName returns MTProto type name.
func (c *MessagesChatAdminsWithInvites) SchemaName() string {
	return "messages.chatAdminsWithInvites"
}

// Encode implements bin.Encoder.
func (c *MessagesChatAdminsWithInvites) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chatAdminsWithInvites#b69b72d7 as nil")
	}
	b.PutID(MessagesChatAdminsWithInvitesTypeID)
	b.PutVectorHeader(len(c.Admins))
	for idx, v := range c.Admins {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatAdminsWithInvites#b69b72d7: field admins element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.chatAdminsWithInvites#b69b72d7: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatAdminsWithInvites#b69b72d7: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetAdmins returns value of Admins field.
func (c *MessagesChatAdminsWithInvites) GetAdmins() (value []ChatAdminWithInvites) {
	return c.Admins
}

// GetUsers returns value of Users field.
func (c *MessagesChatAdminsWithInvites) GetUsers() (value []UserClass) {
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassSlice helper.
func (c *MessagesChatAdminsWithInvites) MapUsers() (value UserClassSlice) {
	return UserClassSlice(c.Users)
}

// Decode implements bin.Decoder.
func (c *MessagesChatAdminsWithInvites) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chatAdminsWithInvites#b69b72d7 to nil")
	}
	if err := b.ConsumeID(MessagesChatAdminsWithInvitesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.chatAdminsWithInvites#b69b72d7: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatAdminsWithInvites#b69b72d7: field admins: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatAdminWithInvites
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.chatAdminsWithInvites#b69b72d7: field admins: %w", err)
			}
			c.Admins = append(c.Admins, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatAdminsWithInvites#b69b72d7: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.chatAdminsWithInvites#b69b72d7: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesChatAdminsWithInvites.
var (
	_ bin.Encoder = &MessagesChatAdminsWithInvites{}
	_ bin.Decoder = &MessagesChatAdminsWithInvites{}
)
