package blocklog

import (
	"encoding/binary"
	"io"

	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/util"
)

// EventLookupKey is a globally unique reference to the event:
// block index + index of the request within block + index of the event within the request
type EventLookupKey [8]byte

func NewEventLookupKey(blockIndex uint32, requestIndex, eventIndex uint16) EventLookupKey {
	ret := EventLookupKey{}
	copy(ret[:4], util.Uint32To4Bytes(blockIndex))
	copy(ret[4:6], util.Uint16To2Bytes(requestIndex))
	copy(ret[6:8], util.Uint16To2Bytes(eventIndex))
	return ret
}

func (k EventLookupKey) BlockIndex() uint32 {
	return util.MustUint32From4Bytes(k[:4])
}

func (k EventLookupKey) RequestIndex() uint16 {
	return util.MustUint16From2Bytes(k[4:6])
}

func (k EventLookupKey) RequestEventIndex() uint16 {
	return util.MustUint16From2Bytes(k[6:8])
}

func (k EventLookupKey) Bytes() []byte {
	return k[:]
}

func (k *EventLookupKey) Write(w io.Writer) error {
	_, err := w.Write(k[:])
	return err
}

func EventLookupKeyFromBytes(r io.Reader) (*EventLookupKey, error) {
	k := EventLookupKey{}
	n, err := r.Read(k[:])
	if err != nil || n != 8 {
		return nil, io.EOF
	}
	return &k, nil
}

type EventData struct {
	ContractID isc.Hname
	Payload    []byte
	Topic      string
	Timestamp  uint64
}

func EventDecode(eventData []byte) *EventData {
	hContract := isc.Hname(binary.LittleEndian.Uint32(eventData[:4]))
	eventData = eventData[4:]
	length := binary.LittleEndian.Uint16(eventData[:2])
	eventData = eventData[2:]
	topic := string(eventData[:length])
	eventData = eventData[length:]
	timestamp := binary.LittleEndian.Uint64(eventData[:8])
	return &EventData{
		ContractID: hContract,
		Payload:    eventData[8:],
		Timestamp:  timestamp,
		Topic:      topic,
	}
}
