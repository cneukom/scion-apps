// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type CertChainReq struct{ capnp.Struct }

// CertChainReq_TypeID is the unique identifier for the type CertChainReq.
const CertChainReq_TypeID = 0xc464d1e0777e54d3

func NewCertChainReq(s *capnp.Segment) (CertChainReq, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 0})
	return CertChainReq{st}, err
}

func NewRootCertChainReq(s *capnp.Segment) (CertChainReq, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 0})
	return CertChainReq{st}, err
}

func ReadRootCertChainReq(msg *capnp.Message) (CertChainReq, error) {
	root, err := msg.RootPtr()
	return CertChainReq{root.Struct()}, err
}

func (s CertChainReq) String() string {
	str, _ := text.Marshal(0xc464d1e0777e54d3, s.Struct)
	return str
}

func (s CertChainReq) Isdas() uint64 {
	return s.Struct.Uint64(0)
}

func (s CertChainReq) SetIsdas(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s CertChainReq) Version() uint64 {
	return s.Struct.Uint64(8)
}

func (s CertChainReq) SetVersion(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s CertChainReq) CacheOnly() bool {
	return s.Struct.Bit(128)
}

func (s CertChainReq) SetCacheOnly(v bool) {
	s.Struct.SetBit(128, v)
}

// CertChainReq_List is a list of CertChainReq.
type CertChainReq_List struct{ capnp.List }

// NewCertChainReq creates a new list of CertChainReq.
func NewCertChainReq_List(s *capnp.Segment, sz int32) (CertChainReq_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 0}, sz)
	return CertChainReq_List{l}, err
}

func (s CertChainReq_List) At(i int) CertChainReq { return CertChainReq{s.List.Struct(i)} }

func (s CertChainReq_List) Set(i int, v CertChainReq) error { return s.List.SetStruct(i, v.Struct) }

func (s CertChainReq_List) String() string {
	str, _ := text.MarshalList(0xc464d1e0777e54d3, s.List)
	return str
}

// CertChainReq_Promise is a wrapper for a CertChainReq promised by a client call.
type CertChainReq_Promise struct{ *capnp.Pipeline }

func (p CertChainReq_Promise) Struct() (CertChainReq, error) {
	s, err := p.Pipeline.Struct()
	return CertChainReq{s}, err
}

type CertChain struct{ capnp.Struct }

// CertChain_TypeID is the unique identifier for the type CertChain.
const CertChain_TypeID = 0xadadc71f7e190917

func NewCertChain(s *capnp.Segment) (CertChain, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CertChain{st}, err
}

func NewRootCertChain(s *capnp.Segment) (CertChain, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CertChain{st}, err
}

func ReadRootCertChain(msg *capnp.Message) (CertChain, error) {
	root, err := msg.RootPtr()
	return CertChain{root.Struct()}, err
}

func (s CertChain) String() string {
	str, _ := text.Marshal(0xadadc71f7e190917, s.Struct)
	return str
}

func (s CertChain) Chain() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s CertChain) HasChain() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertChain) SetChain(v []byte) error {
	return s.Struct.SetData(0, v)
}

// CertChain_List is a list of CertChain.
type CertChain_List struct{ capnp.List }

// NewCertChain creates a new list of CertChain.
func NewCertChain_List(s *capnp.Segment, sz int32) (CertChain_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return CertChain_List{l}, err
}

func (s CertChain_List) At(i int) CertChain { return CertChain{s.List.Struct(i)} }

func (s CertChain_List) Set(i int, v CertChain) error { return s.List.SetStruct(i, v.Struct) }

func (s CertChain_List) String() string {
	str, _ := text.MarshalList(0xadadc71f7e190917, s.List)
	return str
}

// CertChain_Promise is a wrapper for a CertChain promised by a client call.
type CertChain_Promise struct{ *capnp.Pipeline }

func (p CertChain_Promise) Struct() (CertChain, error) {
	s, err := p.Pipeline.Struct()
	return CertChain{s}, err
}

type CertChainIssReq struct{ capnp.Struct }

// CertChainIssReq_TypeID is the unique identifier for the type CertChainIssReq.
const CertChainIssReq_TypeID = 0xb2de94224c009676

func NewCertChainIssReq(s *capnp.Segment) (CertChainIssReq, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CertChainIssReq{st}, err
}

func NewRootCertChainIssReq(s *capnp.Segment) (CertChainIssReq, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CertChainIssReq{st}, err
}

func ReadRootCertChainIssReq(msg *capnp.Message) (CertChainIssReq, error) {
	root, err := msg.RootPtr()
	return CertChainIssReq{root.Struct()}, err
}

func (s CertChainIssReq) String() string {
	str, _ := text.Marshal(0xb2de94224c009676, s.Struct)
	return str
}

func (s CertChainIssReq) Cert() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s CertChainIssReq) HasCert() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertChainIssReq) SetCert(v []byte) error {
	return s.Struct.SetData(0, v)
}

// CertChainIssReq_List is a list of CertChainIssReq.
type CertChainIssReq_List struct{ capnp.List }

// NewCertChainIssReq creates a new list of CertChainIssReq.
func NewCertChainIssReq_List(s *capnp.Segment, sz int32) (CertChainIssReq_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return CertChainIssReq_List{l}, err
}

func (s CertChainIssReq_List) At(i int) CertChainIssReq { return CertChainIssReq{s.List.Struct(i)} }

func (s CertChainIssReq_List) Set(i int, v CertChainIssReq) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s CertChainIssReq_List) String() string {
	str, _ := text.MarshalList(0xb2de94224c009676, s.List)
	return str
}

// CertChainIssReq_Promise is a wrapper for a CertChainIssReq promised by a client call.
type CertChainIssReq_Promise struct{ *capnp.Pipeline }

func (p CertChainIssReq_Promise) Struct() (CertChainIssReq, error) {
	s, err := p.Pipeline.Struct()
	return CertChainIssReq{s}, err
}

type CertChainIssRep struct{ capnp.Struct }

// CertChainIssRep_TypeID is the unique identifier for the type CertChainIssRep.
const CertChainIssRep_TypeID = 0xc95b16276878cfc1

func NewCertChainIssRep(s *capnp.Segment) (CertChainIssRep, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CertChainIssRep{st}, err
}

func NewRootCertChainIssRep(s *capnp.Segment) (CertChainIssRep, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CertChainIssRep{st}, err
}

func ReadRootCertChainIssRep(msg *capnp.Message) (CertChainIssRep, error) {
	root, err := msg.RootPtr()
	return CertChainIssRep{root.Struct()}, err
}

func (s CertChainIssRep) String() string {
	str, _ := text.Marshal(0xc95b16276878cfc1, s.Struct)
	return str
}

func (s CertChainIssRep) Chain() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s CertChainIssRep) HasChain() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertChainIssRep) SetChain(v []byte) error {
	return s.Struct.SetData(0, v)
}

// CertChainIssRep_List is a list of CertChainIssRep.
type CertChainIssRep_List struct{ capnp.List }

// NewCertChainIssRep creates a new list of CertChainIssRep.
func NewCertChainIssRep_List(s *capnp.Segment, sz int32) (CertChainIssRep_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return CertChainIssRep_List{l}, err
}

func (s CertChainIssRep_List) At(i int) CertChainIssRep { return CertChainIssRep{s.List.Struct(i)} }

func (s CertChainIssRep_List) Set(i int, v CertChainIssRep) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s CertChainIssRep_List) String() string {
	str, _ := text.MarshalList(0xc95b16276878cfc1, s.List)
	return str
}

// CertChainIssRep_Promise is a wrapper for a CertChainIssRep promised by a client call.
type CertChainIssRep_Promise struct{ *capnp.Pipeline }

func (p CertChainIssRep_Promise) Struct() (CertChainIssRep, error) {
	s, err := p.Pipeline.Struct()
	return CertChainIssRep{s}, err
}

type TRCReq struct{ capnp.Struct }

// TRCReq_TypeID is the unique identifier for the type TRCReq.
const TRCReq_TypeID = 0xd4c43f7ac10a9dbc

func NewTRCReq(s *capnp.Segment) (TRCReq, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return TRCReq{st}, err
}

func NewRootTRCReq(s *capnp.Segment) (TRCReq, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return TRCReq{st}, err
}

func ReadRootTRCReq(msg *capnp.Message) (TRCReq, error) {
	root, err := msg.RootPtr()
	return TRCReq{root.Struct()}, err
}

func (s TRCReq) String() string {
	str, _ := text.Marshal(0xd4c43f7ac10a9dbc, s.Struct)
	return str
}

func (s TRCReq) Isd() uint16 {
	return s.Struct.Uint16(0)
}

func (s TRCReq) SetIsd(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s TRCReq) Version() uint64 {
	return s.Struct.Uint64(8)
}

func (s TRCReq) SetVersion(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s TRCReq) CacheOnly() bool {
	return s.Struct.Bit(16)
}

func (s TRCReq) SetCacheOnly(v bool) {
	s.Struct.SetBit(16, v)
}

// TRCReq_List is a list of TRCReq.
type TRCReq_List struct{ capnp.List }

// NewTRCReq creates a new list of TRCReq.
func NewTRCReq_List(s *capnp.Segment, sz int32) (TRCReq_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return TRCReq_List{l}, err
}

func (s TRCReq_List) At(i int) TRCReq { return TRCReq{s.List.Struct(i)} }

func (s TRCReq_List) Set(i int, v TRCReq) error { return s.List.SetStruct(i, v.Struct) }

func (s TRCReq_List) String() string {
	str, _ := text.MarshalList(0xd4c43f7ac10a9dbc, s.List)
	return str
}

// TRCReq_Promise is a wrapper for a TRCReq promised by a client call.
type TRCReq_Promise struct{ *capnp.Pipeline }

func (p TRCReq_Promise) Struct() (TRCReq, error) {
	s, err := p.Pipeline.Struct()
	return TRCReq{s}, err
}

type TRC struct{ capnp.Struct }

// TRC_TypeID is the unique identifier for the type TRC.
const TRC_TypeID = 0x9aee2af152a5f7d7

func NewTRC(s *capnp.Segment) (TRC, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TRC{st}, err
}

func NewRootTRC(s *capnp.Segment) (TRC, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TRC{st}, err
}

func ReadRootTRC(msg *capnp.Message) (TRC, error) {
	root, err := msg.RootPtr()
	return TRC{root.Struct()}, err
}

func (s TRC) String() string {
	str, _ := text.Marshal(0x9aee2af152a5f7d7, s.Struct)
	return str
}

func (s TRC) Trc() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s TRC) HasTrc() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TRC) SetTrc(v []byte) error {
	return s.Struct.SetData(0, v)
}

// TRC_List is a list of TRC.
type TRC_List struct{ capnp.List }

// NewTRC creates a new list of TRC.
func NewTRC_List(s *capnp.Segment, sz int32) (TRC_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return TRC_List{l}, err
}

func (s TRC_List) At(i int) TRC { return TRC{s.List.Struct(i)} }

func (s TRC_List) Set(i int, v TRC) error { return s.List.SetStruct(i, v.Struct) }

func (s TRC_List) String() string {
	str, _ := text.MarshalList(0x9aee2af152a5f7d7, s.List)
	return str
}

// TRC_Promise is a wrapper for a TRC promised by a client call.
type TRC_Promise struct{ *capnp.Pipeline }

func (p TRC_Promise) Struct() (TRC, error) {
	s, err := p.Pipeline.Struct()
	return TRC{s}, err
}

type CertMgmt struct{ capnp.Struct }
type CertMgmt_Which uint16

const (
	CertMgmt_Which_unset           CertMgmt_Which = 0
	CertMgmt_Which_certChainReq    CertMgmt_Which = 1
	CertMgmt_Which_certChain       CertMgmt_Which = 2
	CertMgmt_Which_trcReq          CertMgmt_Which = 3
	CertMgmt_Which_trc             CertMgmt_Which = 4
	CertMgmt_Which_certChainIssReq CertMgmt_Which = 5
	CertMgmt_Which_certChainIssRep CertMgmt_Which = 6
)

func (w CertMgmt_Which) String() string {
	const s = "unsetcertChainReqcertChaintrcReqtrccertChainIssReqcertChainIssRep"
	switch w {
	case CertMgmt_Which_unset:
		return s[0:5]
	case CertMgmt_Which_certChainReq:
		return s[5:17]
	case CertMgmt_Which_certChain:
		return s[17:26]
	case CertMgmt_Which_trcReq:
		return s[26:32]
	case CertMgmt_Which_trc:
		return s[32:35]
	case CertMgmt_Which_certChainIssReq:
		return s[35:50]
	case CertMgmt_Which_certChainIssRep:
		return s[50:65]

	}
	return "CertMgmt_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// CertMgmt_TypeID is the unique identifier for the type CertMgmt.
const CertMgmt_TypeID = 0xa19070b486ecd839

func NewCertMgmt(s *capnp.Segment) (CertMgmt, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CertMgmt{st}, err
}

func NewRootCertMgmt(s *capnp.Segment) (CertMgmt, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CertMgmt{st}, err
}

func ReadRootCertMgmt(msg *capnp.Message) (CertMgmt, error) {
	root, err := msg.RootPtr()
	return CertMgmt{root.Struct()}, err
}

func (s CertMgmt) String() string {
	str, _ := text.Marshal(0xa19070b486ecd839, s.Struct)
	return str
}

func (s CertMgmt) Which() CertMgmt_Which {
	return CertMgmt_Which(s.Struct.Uint16(0))
}
func (s CertMgmt) SetUnset() {
	s.Struct.SetUint16(0, 0)

}

func (s CertMgmt) CertChainReq() (CertChainReq, error) {
	p, err := s.Struct.Ptr(0)
	return CertChainReq{Struct: p.Struct()}, err
}

func (s CertMgmt) HasCertChainReq() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertMgmt) SetCertChainReq(v CertChainReq) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCertChainReq sets the certChainReq field to a newly
// allocated CertChainReq struct, preferring placement in s's segment.
func (s CertMgmt) NewCertChainReq() (CertChainReq, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewCertChainReq(s.Struct.Segment())
	if err != nil {
		return CertChainReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s CertMgmt) CertChain() (CertChain, error) {
	p, err := s.Struct.Ptr(0)
	return CertChain{Struct: p.Struct()}, err
}

func (s CertMgmt) HasCertChain() bool {
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertMgmt) SetCertChain(v CertChain) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCertChain sets the certChain field to a newly
// allocated CertChain struct, preferring placement in s's segment.
func (s CertMgmt) NewCertChain() (CertChain, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewCertChain(s.Struct.Segment())
	if err != nil {
		return CertChain{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s CertMgmt) TrcReq() (TRCReq, error) {
	p, err := s.Struct.Ptr(0)
	return TRCReq{Struct: p.Struct()}, err
}

func (s CertMgmt) HasTrcReq() bool {
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertMgmt) SetTrcReq(v TRCReq) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTrcReq sets the trcReq field to a newly
// allocated TRCReq struct, preferring placement in s's segment.
func (s CertMgmt) NewTrcReq() (TRCReq, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewTRCReq(s.Struct.Segment())
	if err != nil {
		return TRCReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s CertMgmt) Trc() (TRC, error) {
	p, err := s.Struct.Ptr(0)
	return TRC{Struct: p.Struct()}, err
}

func (s CertMgmt) HasTrc() bool {
	if s.Struct.Uint16(0) != 4 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertMgmt) SetTrc(v TRC) error {
	s.Struct.SetUint16(0, 4)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTrc sets the trc field to a newly
// allocated TRC struct, preferring placement in s's segment.
func (s CertMgmt) NewTrc() (TRC, error) {
	s.Struct.SetUint16(0, 4)
	ss, err := NewTRC(s.Struct.Segment())
	if err != nil {
		return TRC{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s CertMgmt) CertChainIssReq() (CertChainIssReq, error) {
	p, err := s.Struct.Ptr(0)
	return CertChainIssReq{Struct: p.Struct()}, err
}

func (s CertMgmt) HasCertChainIssReq() bool {
	if s.Struct.Uint16(0) != 5 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertMgmt) SetCertChainIssReq(v CertChainIssReq) error {
	s.Struct.SetUint16(0, 5)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCertChainIssReq sets the certChainIssReq field to a newly
// allocated CertChainIssReq struct, preferring placement in s's segment.
func (s CertMgmt) NewCertChainIssReq() (CertChainIssReq, error) {
	s.Struct.SetUint16(0, 5)
	ss, err := NewCertChainIssReq(s.Struct.Segment())
	if err != nil {
		return CertChainIssReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s CertMgmt) CertChainIssRep() (CertChainIssRep, error) {
	p, err := s.Struct.Ptr(0)
	return CertChainIssRep{Struct: p.Struct()}, err
}

func (s CertMgmt) HasCertChainIssRep() bool {
	if s.Struct.Uint16(0) != 6 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CertMgmt) SetCertChainIssRep(v CertChainIssRep) error {
	s.Struct.SetUint16(0, 6)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCertChainIssRep sets the certChainIssRep field to a newly
// allocated CertChainIssRep struct, preferring placement in s's segment.
func (s CertMgmt) NewCertChainIssRep() (CertChainIssRep, error) {
	s.Struct.SetUint16(0, 6)
	ss, err := NewCertChainIssRep(s.Struct.Segment())
	if err != nil {
		return CertChainIssRep{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// CertMgmt_List is a list of CertMgmt.
type CertMgmt_List struct{ capnp.List }

// NewCertMgmt creates a new list of CertMgmt.
func NewCertMgmt_List(s *capnp.Segment, sz int32) (CertMgmt_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return CertMgmt_List{l}, err
}

func (s CertMgmt_List) At(i int) CertMgmt { return CertMgmt{s.List.Struct(i)} }

func (s CertMgmt_List) Set(i int, v CertMgmt) error { return s.List.SetStruct(i, v.Struct) }

func (s CertMgmt_List) String() string {
	str, _ := text.MarshalList(0xa19070b486ecd839, s.List)
	return str
}

// CertMgmt_Promise is a wrapper for a CertMgmt promised by a client call.
type CertMgmt_Promise struct{ *capnp.Pipeline }

func (p CertMgmt_Promise) Struct() (CertMgmt, error) {
	s, err := p.Pipeline.Struct()
	return CertMgmt{s}, err
}

func (p CertMgmt_Promise) CertChainReq() CertChainReq_Promise {
	return CertChainReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p CertMgmt_Promise) CertChain() CertChain_Promise {
	return CertChain_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p CertMgmt_Promise) TrcReq() TRCReq_Promise {
	return TRCReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p CertMgmt_Promise) Trc() TRC_Promise {
	return TRC_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p CertMgmt_Promise) CertChainIssReq() CertChainIssReq_Promise {
	return CertChainIssReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p CertMgmt_Promise) CertChainIssRep() CertChainIssRep_Promise {
	return CertChainIssRep_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_ec3b2b10a5e23975 = "x\xda\x94\x93OHT]\x18\xc6\xdf\xe7\x9c;3\x9f" +
	"\x1f^\xbc\x97;|\x1fA0\xe0&\xb0?\x94\xae\xb4" +
	"\xc5H\x93\x0bA\xc93\xb8\x08Z\xc40\x0e\xceD\xde" +
	"\xc6\xb9W\xfb\xb3\xb0\x08j\x17%DP\xe0\xa2\xc0M" +
	" V$d\xb8h`\xa2\x8cD-#\x83\x88\x8a6" +
	"\x85\x05\xae\xc2\xccN\x9c;\xcd\\\x99\x91\xac\xdd\xb9\xe7" +
	"<\xef\x8f\xf7y\xdf\xe7\xee\xbe\x89V\xb6'0\xcb\x89" +
	"\xc4\x8e@P\xbe\xfc:\x1a_n\xf8r\x8dL\x03r" +
	"\xa0\xf9\xfd\xa8\xb1}\xef\x12\x05\x10\"\xb2fp\xc3z" +
	"\xe1\x9d\xe6\x11%\xc8\xe6\xc5\xa5\xf3w\xb3\x97\xae\x930" +
	"\xb0N\xdc\x86\x90\xd2,\xe3\x89\xb5\xa6\xd4M+\xb8\x08" +
	"\x82\xfc\xbff\xcbP\xe4\xd1\xd8\xd8F\xe8a\xfe\xcc\x1a" +
	"\xe1\xeat\x95G\x09\xab\x83W:\xea/\xbf\xb9S\xad" +
	"l\x9a\xe4\x0c\xd6CO\x9aWR\xf9\xbc{\xe8\xf8\xdb" +
	"\xf9\x9e\x82\xea\x82\xfbjM)\xde\xf1O\xd6gO\xfb" +
	"\x91\x8f\x13d~\xf6Dz\xdb\x7f\x87\xa67h\xa1\xe9" +
	"\x9c\xc6`\x0d{e\x174\x05\x9e\x1a\xf97\x7f*Z" +
	"XP`V\x01\x1e\xd3\xee[\x13\xde\xe9\xb66N;" +
	"e2\x95s\x0f\xf7\xf5\xf6\xc1\xdd\x95Ld\xedlK" +
	"\xa8;\x1e\xeb\x02\x84\xc65\"\x0dD\xa6^O$\xfe" +
	"\xe1\x10a\x86\x90\x9bKB'\x06\x9dP\xaee\xa5\xda" +
	"X*\xe7v\xf6\xf6\xb9D\x8a\xb0\x95k\xb5Rz\x88" +
	"\x89F\"q\x8bCL1\xe8\xf8!\xc3P\xb7\x93G" +
	"\x88\xc4=\x0eQ`\xd0\xd9\x9a\x0c\x83\x11\x99\xf98\x91" +
	"x\xc0!\x9e2\xe8\xfc\xbb\x0c\x83\x13\x99\xd3-D\xa2" +
	"\xc0!\xe6\x18tmU\x86\xa1\x11\x993\xaa\xb5\xc7\x1c" +
	"b\x81A\x0f|\x93a\x04\x88\xcc\xf9\xb3Db\x8eC" +
	"\xbcf\xd0\x83+2\x8c \x91\xf9J\xdd.r\x88\x0f" +
	"\x0c\x91\x01\xdbI\xb9\x14\xf4,\xc4\xd2\x89\x0c\xd5\xd9\xf1" +
	"T?\x0c\x7f)\x04\x18\xbf<z\x02\xd80\xfc$\x14" +
	"_\xa3n.Y,+\x8f\xbc\xf8\xe0\x8d\xc9\xf03Y" +
	"\x01\x83\xdd\xee8\xf1T?\xc1(\x05fcE\x96`" +
	"\xf8\xbb_\xa7\xa9\x1a{\xb1\xa8bq\x8d\xfe\xe2\"\xc9" +
	"t\"cW\xad\x8eW2J\x9dU\x90\x1a|R\x9d" +
	"\xaa\xfd}\x06\x14\xc8\x1b\xa8\xa2\xd4\x96)m\xaa\x9fV" +
	"\x0e\xd1\xc1`\x02\xc5\x10\xb4\xef#\x12\xfb9D\x17\x83" +
	"\xc9\xce\x143\xd0\xa92\xd0\xc1!\x0e2D2NO" +
	"\xc2A\x0d1\xd4\x10N\x0f\xa6rN\xe6\x98]\xfa\x96" +
	"\xc9D2\x9d:`\x1f%\x9c\x04\x88\x01\x9b\xbb\xcb\xd2" +
	"_\xcf\xa9\xfc{D\xbb\xe3\xb1jc\xf5\x9b\x1a3\xaa" +
	"\x8d\x852N\x0fB\xc4\x10\xfa3[?\x03\x00\x00\xff" +
	"\xffJ\x0aF\xc7"

func init() {
	schemas.Register(schema_ec3b2b10a5e23975,
		0x9aee2af152a5f7d7,
		0xa19070b486ecd839,
		0xadadc71f7e190917,
		0xb2de94224c009676,
		0xc464d1e0777e54d3,
		0xc95b16276878cfc1,
		0xd4c43f7ac10a9dbc)
}
