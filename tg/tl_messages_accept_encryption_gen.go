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

// MessagesAcceptEncryptionRequest represents TL type `messages.acceptEncryption#3dbc0415`.
// Confirms creation of a secret chat
//
// See https://core.telegram.org/method/messages.acceptEncryption for reference.
type MessagesAcceptEncryptionRequest struct {
	// Secret chat ID
	Peer InputEncryptedChat
	// B = g ^ b mod p, see Wikipedia
	GB []byte
	// 64-bit fingerprint of the received key
	KeyFingerprint int64
}

// MessagesAcceptEncryptionRequestTypeID is TL type id of MessagesAcceptEncryptionRequest.
const MessagesAcceptEncryptionRequestTypeID = 0x3dbc0415

// Encode implements bin.Encoder.
func (a *MessagesAcceptEncryptionRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.acceptEncryption#3dbc0415 as nil")
	}
	b.PutID(MessagesAcceptEncryptionRequestTypeID)
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.acceptEncryption#3dbc0415: field peer: %w", err)
	}
	b.PutBytes(a.GB)
	b.PutLong(a.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAcceptEncryptionRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.acceptEncryption#3dbc0415 to nil")
	}
	if err := b.ConsumeID(MessagesAcceptEncryptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.acceptEncryption#3dbc0415: %w", err)
	}
	{
		if err := a.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.acceptEncryption#3dbc0415: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.acceptEncryption#3dbc0415: field g_b: %w", err)
		}
		a.GB = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.acceptEncryption#3dbc0415: field key_fingerprint: %w", err)
		}
		a.KeyFingerprint = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAcceptEncryptionRequest.
var (
	_ bin.Encoder = &MessagesAcceptEncryptionRequest{}
	_ bin.Decoder = &MessagesAcceptEncryptionRequest{}
)

// MessagesAcceptEncryption invokes method messages.acceptEncryption#3dbc0415 returning error if any.
// Confirms creation of a secret chat
//
// See https://core.telegram.org/method/messages.acceptEncryption for reference.
func (c *Client) MessagesAcceptEncryption(ctx context.Context, request *MessagesAcceptEncryptionRequest) (EncryptedChatClass, error) {
	var result EncryptedChatBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EncryptedChat, nil
}
