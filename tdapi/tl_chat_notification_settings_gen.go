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

// ChatNotificationSettings represents TL type `chatNotificationSettings#38493760`.
type ChatNotificationSettings struct {
	// If true, mute_for is ignored and the value for the relevant type of chat is used
	// instead
	UseDefaultMuteFor bool
	// Time left before notifications will be unmuted, in seconds
	MuteFor int32
	// If true, the value for the relevant type of chat is used instead of sound_id
	UseDefaultSound bool
	// Identifier of the notification sound to be played; 0 if sound is disabled
	SoundID int64
	// If true, show_preview is ignored and the value for the relevant type of chat is used
	// instead
	UseDefaultShowPreview bool
	// True, if message content must be displayed in notifications
	ShowPreview bool
	// If true, disable_pinned_message_notifications is ignored and the value for the
	// relevant type of chat is used instead
	UseDefaultDisablePinnedMessageNotifications bool
	// If true, notifications for incoming pinned messages will be created as for an ordinary
	// unread message
	DisablePinnedMessageNotifications bool
	// If true, disable_mention_notifications is ignored and the value for the relevant type
	// of chat is used instead
	UseDefaultDisableMentionNotifications bool
	// If true, notifications for messages with mentions will be created as for an ordinary
	// unread message
	DisableMentionNotifications bool
}

// ChatNotificationSettingsTypeID is TL type id of ChatNotificationSettings.
const ChatNotificationSettingsTypeID = 0x38493760

// Ensuring interfaces in compile-time for ChatNotificationSettings.
var (
	_ bin.Encoder     = &ChatNotificationSettings{}
	_ bin.Decoder     = &ChatNotificationSettings{}
	_ bin.BareEncoder = &ChatNotificationSettings{}
	_ bin.BareDecoder = &ChatNotificationSettings{}
)

func (c *ChatNotificationSettings) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UseDefaultMuteFor == false) {
		return false
	}
	if !(c.MuteFor == 0) {
		return false
	}
	if !(c.UseDefaultSound == false) {
		return false
	}
	if !(c.SoundID == 0) {
		return false
	}
	if !(c.UseDefaultShowPreview == false) {
		return false
	}
	if !(c.ShowPreview == false) {
		return false
	}
	if !(c.UseDefaultDisablePinnedMessageNotifications == false) {
		return false
	}
	if !(c.DisablePinnedMessageNotifications == false) {
		return false
	}
	if !(c.UseDefaultDisableMentionNotifications == false) {
		return false
	}
	if !(c.DisableMentionNotifications == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatNotificationSettings) String() string {
	if c == nil {
		return "ChatNotificationSettings(nil)"
	}
	type Alias ChatNotificationSettings
	return fmt.Sprintf("ChatNotificationSettings%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatNotificationSettings) TypeID() uint32 {
	return ChatNotificationSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatNotificationSettings) TypeName() string {
	return "chatNotificationSettings"
}

// TypeInfo returns info about TL type.
func (c *ChatNotificationSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatNotificationSettings",
		ID:   ChatNotificationSettingsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UseDefaultMuteFor",
			SchemaName: "use_default_mute_for",
		},
		{
			Name:       "MuteFor",
			SchemaName: "mute_for",
		},
		{
			Name:       "UseDefaultSound",
			SchemaName: "use_default_sound",
		},
		{
			Name:       "SoundID",
			SchemaName: "sound_id",
		},
		{
			Name:       "UseDefaultShowPreview",
			SchemaName: "use_default_show_preview",
		},
		{
			Name:       "ShowPreview",
			SchemaName: "show_preview",
		},
		{
			Name:       "UseDefaultDisablePinnedMessageNotifications",
			SchemaName: "use_default_disable_pinned_message_notifications",
		},
		{
			Name:       "DisablePinnedMessageNotifications",
			SchemaName: "disable_pinned_message_notifications",
		},
		{
			Name:       "UseDefaultDisableMentionNotifications",
			SchemaName: "use_default_disable_mention_notifications",
		},
		{
			Name:       "DisableMentionNotifications",
			SchemaName: "disable_mention_notifications",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatNotificationSettings) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNotificationSettings#38493760 as nil")
	}
	b.PutID(ChatNotificationSettingsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatNotificationSettings) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNotificationSettings#38493760 as nil")
	}
	b.PutBool(c.UseDefaultMuteFor)
	b.PutInt32(c.MuteFor)
	b.PutBool(c.UseDefaultSound)
	b.PutLong(c.SoundID)
	b.PutBool(c.UseDefaultShowPreview)
	b.PutBool(c.ShowPreview)
	b.PutBool(c.UseDefaultDisablePinnedMessageNotifications)
	b.PutBool(c.DisablePinnedMessageNotifications)
	b.PutBool(c.UseDefaultDisableMentionNotifications)
	b.PutBool(c.DisableMentionNotifications)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatNotificationSettings) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNotificationSettings#38493760 to nil")
	}
	if err := b.ConsumeID(ChatNotificationSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatNotificationSettings#38493760: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatNotificationSettings) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNotificationSettings#38493760 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_mute_for: %w", err)
		}
		c.UseDefaultMuteFor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field mute_for: %w", err)
		}
		c.MuteFor = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_sound: %w", err)
		}
		c.UseDefaultSound = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field sound_id: %w", err)
		}
		c.SoundID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_show_preview: %w", err)
		}
		c.UseDefaultShowPreview = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field show_preview: %w", err)
		}
		c.ShowPreview = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_disable_pinned_message_notifications: %w", err)
		}
		c.UseDefaultDisablePinnedMessageNotifications = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field disable_pinned_message_notifications: %w", err)
		}
		c.DisablePinnedMessageNotifications = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_disable_mention_notifications: %w", err)
		}
		c.UseDefaultDisableMentionNotifications = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field disable_mention_notifications: %w", err)
		}
		c.DisableMentionNotifications = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatNotificationSettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNotificationSettings#38493760 as nil")
	}
	b.ObjStart()
	b.PutID("chatNotificationSettings")
	b.Comma()
	b.FieldStart("use_default_mute_for")
	b.PutBool(c.UseDefaultMuteFor)
	b.Comma()
	b.FieldStart("mute_for")
	b.PutInt32(c.MuteFor)
	b.Comma()
	b.FieldStart("use_default_sound")
	b.PutBool(c.UseDefaultSound)
	b.Comma()
	b.FieldStart("sound_id")
	b.PutLong(c.SoundID)
	b.Comma()
	b.FieldStart("use_default_show_preview")
	b.PutBool(c.UseDefaultShowPreview)
	b.Comma()
	b.FieldStart("show_preview")
	b.PutBool(c.ShowPreview)
	b.Comma()
	b.FieldStart("use_default_disable_pinned_message_notifications")
	b.PutBool(c.UseDefaultDisablePinnedMessageNotifications)
	b.Comma()
	b.FieldStart("disable_pinned_message_notifications")
	b.PutBool(c.DisablePinnedMessageNotifications)
	b.Comma()
	b.FieldStart("use_default_disable_mention_notifications")
	b.PutBool(c.UseDefaultDisableMentionNotifications)
	b.Comma()
	b.FieldStart("disable_mention_notifications")
	b.PutBool(c.DisableMentionNotifications)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatNotificationSettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNotificationSettings#38493760 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatNotificationSettings"); err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: %w", err)
			}
		case "use_default_mute_for":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_mute_for: %w", err)
			}
			c.UseDefaultMuteFor = value
		case "mute_for":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field mute_for: %w", err)
			}
			c.MuteFor = value
		case "use_default_sound":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_sound: %w", err)
			}
			c.UseDefaultSound = value
		case "sound_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field sound_id: %w", err)
			}
			c.SoundID = value
		case "use_default_show_preview":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_show_preview: %w", err)
			}
			c.UseDefaultShowPreview = value
		case "show_preview":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field show_preview: %w", err)
			}
			c.ShowPreview = value
		case "use_default_disable_pinned_message_notifications":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_disable_pinned_message_notifications: %w", err)
			}
			c.UseDefaultDisablePinnedMessageNotifications = value
		case "disable_pinned_message_notifications":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field disable_pinned_message_notifications: %w", err)
			}
			c.DisablePinnedMessageNotifications = value
		case "use_default_disable_mention_notifications":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field use_default_disable_mention_notifications: %w", err)
			}
			c.UseDefaultDisableMentionNotifications = value
		case "disable_mention_notifications":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#38493760: field disable_mention_notifications: %w", err)
			}
			c.DisableMentionNotifications = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUseDefaultMuteFor returns value of UseDefaultMuteFor field.
func (c *ChatNotificationSettings) GetUseDefaultMuteFor() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultMuteFor
}

// GetMuteFor returns value of MuteFor field.
func (c *ChatNotificationSettings) GetMuteFor() (value int32) {
	if c == nil {
		return
	}
	return c.MuteFor
}

// GetUseDefaultSound returns value of UseDefaultSound field.
func (c *ChatNotificationSettings) GetUseDefaultSound() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultSound
}

// GetSoundID returns value of SoundID field.
func (c *ChatNotificationSettings) GetSoundID() (value int64) {
	if c == nil {
		return
	}
	return c.SoundID
}

// GetUseDefaultShowPreview returns value of UseDefaultShowPreview field.
func (c *ChatNotificationSettings) GetUseDefaultShowPreview() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultShowPreview
}

// GetShowPreview returns value of ShowPreview field.
func (c *ChatNotificationSettings) GetShowPreview() (value bool) {
	if c == nil {
		return
	}
	return c.ShowPreview
}

// GetUseDefaultDisablePinnedMessageNotifications returns value of UseDefaultDisablePinnedMessageNotifications field.
func (c *ChatNotificationSettings) GetUseDefaultDisablePinnedMessageNotifications() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultDisablePinnedMessageNotifications
}

// GetDisablePinnedMessageNotifications returns value of DisablePinnedMessageNotifications field.
func (c *ChatNotificationSettings) GetDisablePinnedMessageNotifications() (value bool) {
	if c == nil {
		return
	}
	return c.DisablePinnedMessageNotifications
}

// GetUseDefaultDisableMentionNotifications returns value of UseDefaultDisableMentionNotifications field.
func (c *ChatNotificationSettings) GetUseDefaultDisableMentionNotifications() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultDisableMentionNotifications
}

// GetDisableMentionNotifications returns value of DisableMentionNotifications field.
func (c *ChatNotificationSettings) GetDisableMentionNotifications() (value bool) {
	if c == nil {
		return
	}
	return c.DisableMentionNotifications
}
