// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.1-0.20240409071808-615f978279ca
// source: storage/node_integration.proto

package storage

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *NodeIntegration) CloneVT() *NodeIntegration {
	if m == nil {
		return (*NodeIntegration)(nil)
	}
	r := new(NodeIntegration)
	r.Id = m.Id
	r.Name = m.Name
	r.Type = m.Type
	if m.IntegrationConfig != nil {
		r.IntegrationConfig = m.IntegrationConfig.(interface {
			CloneVT() isNodeIntegration_IntegrationConfig
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *NodeIntegration) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *NodeIntegration_Clairify) CloneVT() isNodeIntegration_IntegrationConfig {
	if m == nil {
		return (*NodeIntegration_Clairify)(nil)
	}
	r := new(NodeIntegration_Clairify)
	r.Clairify = m.Clairify.CloneVT()
	return r
}

func (m *NodeIntegration_Scannerv4) CloneVT() isNodeIntegration_IntegrationConfig {
	if m == nil {
		return (*NodeIntegration_Scannerv4)(nil)
	}
	r := new(NodeIntegration_Scannerv4)
	r.Scannerv4 = m.Scannerv4.CloneVT()
	return r
}

func (this *NodeIntegration) EqualVT(that *NodeIntegration) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.IntegrationConfig == nil && that.IntegrationConfig != nil {
		return false
	} else if this.IntegrationConfig != nil {
		if that.IntegrationConfig == nil {
			return false
		}
		if !this.IntegrationConfig.(interface {
			EqualVT(isNodeIntegration_IntegrationConfig) bool
		}).EqualVT(that.IntegrationConfig) {
			return false
		}
	}
	if this.Id != that.Id {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.Type != that.Type {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *NodeIntegration) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*NodeIntegration)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *NodeIntegration_Clairify) EqualVT(thatIface isNodeIntegration_IntegrationConfig) bool {
	that, ok := thatIface.(*NodeIntegration_Clairify)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Clairify, that.Clairify; p != q {
		if p == nil {
			p = &ClairifyConfig{}
		}
		if q == nil {
			q = &ClairifyConfig{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *NodeIntegration_Scannerv4) EqualVT(thatIface isNodeIntegration_IntegrationConfig) bool {
	that, ok := thatIface.(*NodeIntegration_Scannerv4)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Scannerv4, that.Scannerv4; p != q {
		if p == nil {
			p = &ScannerV4Config{}
		}
		if q == nil {
			q = &ScannerV4Config{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (m *NodeIntegration) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeIntegration) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *NodeIntegration) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.IntegrationConfig.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NodeIntegration_Clairify) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *NodeIntegration_Clairify) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Clairify != nil {
		size, err := m.Clairify.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *NodeIntegration_Scannerv4) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *NodeIntegration_Scannerv4) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Scannerv4 != nil {
		size, err := m.Scannerv4.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *NodeIntegration) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if vtmsg, ok := m.IntegrationConfig.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *NodeIntegration_Clairify) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Clairify != nil {
		l = m.Clairify.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 2
	}
	return n
}
func (m *NodeIntegration_Scannerv4) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Scannerv4 != nil {
		l = m.Scannerv4.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 2
	}
	return n
}
func (m *NodeIntegration) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeIntegration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeIntegration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.Id = stringValue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.Name = stringValue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.Type = stringValue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clairify", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.IntegrationConfig.(*NodeIntegration_Clairify); ok {
				if err := oneof.Clairify.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &ClairifyConfig{}
				if err := v.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.IntegrationConfig = &NodeIntegration_Clairify{Clairify: v}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scannerv4", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.IntegrationConfig.(*NodeIntegration_Scannerv4); ok {
				if err := oneof.Scannerv4.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &ScannerV4Config{}
				if err := v.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.IntegrationConfig = &NodeIntegration_Scannerv4{Scannerv4: v}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
