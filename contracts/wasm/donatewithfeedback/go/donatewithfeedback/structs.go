// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package donatewithfeedback

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type Donation struct {
	Amount    uint64              // amount donated
	Donator   wasmtypes.ScAgentID // who donated
	Error     string              // error to be reported to donator if anything goes wrong
	Feedback  string              // the feedback for the person donated to
	Timestamp uint64              // when the donation took place
}

func NewDonationFromBytes(buf []byte) *Donation {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &Donation{}
	data.Amount = wasmtypes.DecodeUint64(dec)
	data.Donator = wasmtypes.DecodeAgentID(dec)
	data.Error = wasmtypes.DecodeString(dec)
	data.Feedback = wasmtypes.DecodeString(dec)
	data.Timestamp = wasmtypes.DecodeUint64(dec)
	dec.Close()
	return data
}

func (o *Donation) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
	wasmtypes.EncodeUint64(enc, o.Amount)
	wasmtypes.EncodeAgentID(enc, o.Donator)
	wasmtypes.EncodeString(enc, o.Error)
	wasmtypes.EncodeString(enc, o.Feedback)
	wasmtypes.EncodeUint64(enc, o.Timestamp)
	return enc.Buf()
}

type ImmutableDonation struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableDonation) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableDonation) Value() *Donation {
	return NewDonationFromBytes(o.proxy.Get())
}

type MutableDonation struct {
	proxy wasmtypes.Proxy
}

func (o MutableDonation) Delete() {
	o.proxy.Delete()
}

func (o MutableDonation) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableDonation) SetValue(value *Donation) {
	o.proxy.Set(value.Bytes())
}

func (o MutableDonation) Value() *Donation {
	return NewDonationFromBytes(o.proxy.Get())
}
