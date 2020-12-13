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

// UpdatesGetDifferenceRequest represents TL type `updates.getDifference#25939651`.
// Get new updates.
//
// See https://core.telegram.org/method/updates.getDifference for reference.
type UpdatesGetDifferenceRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// PTS, see updates.
	Pts int
	// For fast updating: if provided and pts + pts_total_limit < remote pts, updates.differenceTooLong will be returned.Simply tells the server to not return the difference if it is bigger than pts_total_limitIf the remote pts is too big (> ~4000000), this field will default to 1000000
	//
	// Use SetPtsTotalLimit and GetPtsTotalLimit helpers.
	PtsTotalLimit int
	// date, see updates.
	Date int
	// QTS, see updates.
	Qts int
}

// UpdatesGetDifferenceRequestTypeID is TL type id of UpdatesGetDifferenceRequest.
const UpdatesGetDifferenceRequestTypeID = 0x25939651

// Encode implements bin.Encoder.
func (g *UpdatesGetDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode updates.getDifference#25939651 as nil")
	}
	b.PutID(UpdatesGetDifferenceRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getDifference#25939651: field flags: %w", err)
	}
	b.PutInt(g.Pts)
	if g.Flags.Has(0) {
		b.PutInt(g.PtsTotalLimit)
	}
	b.PutInt(g.Date)
	b.PutInt(g.Qts)
	return nil
}

// SetPtsTotalLimit sets value of PtsTotalLimit conditional field.
func (g *UpdatesGetDifferenceRequest) SetPtsTotalLimit(value int) {
	g.Flags.Set(0)
	g.PtsTotalLimit = value
}

// GetPtsTotalLimit returns value of PtsTotalLimit conditional field and
// boolean which is true if field was set.
func (g *UpdatesGetDifferenceRequest) GetPtsTotalLimit() (value int, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.PtsTotalLimit, true
}

// Decode implements bin.Decoder.
func (g *UpdatesGetDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode updates.getDifference#25939651 to nil")
	}
	if err := b.ConsumeID(UpdatesGetDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.getDifference#25939651: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field pts: %w", err)
		}
		g.Pts = value
	}
	if g.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field pts_total_limit: %w", err)
		}
		g.PtsTotalLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field date: %w", err)
		}
		g.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field qts: %w", err)
		}
		g.Qts = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UpdatesGetDifferenceRequest.
var (
	_ bin.Encoder = &UpdatesGetDifferenceRequest{}
	_ bin.Decoder = &UpdatesGetDifferenceRequest{}
)

// UpdatesGetDifference invokes method updates.getDifference#25939651 returning error if any.
// Get new updates.
//
// See https://core.telegram.org/method/updates.getDifference for reference.
func (c *Client) UpdatesGetDifference(ctx context.Context, request *UpdatesGetDifferenceRequest) (UpdatesDifferenceClass, error) {
	var result UpdatesDifferenceBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Difference, nil
}
