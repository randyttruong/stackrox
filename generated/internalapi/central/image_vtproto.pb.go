// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.1-0.20240409071808-615f978279ca
// source: internalapi/central/image.proto

package central

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	storage "github.com/stackrox/rox/generated/storage"
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

func (m *ScanImage) CloneVT() *ScanImage {
	if m == nil {
		return (*ScanImage)(nil)
	}
	r := new(ScanImage)
	r.RequestId = m.RequestId
	r.ImageName = m.ImageName
	r.Force = m.Force
	r.Namespace = m.Namespace
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ScanImage) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ImageIntegrations) CloneVT() *ImageIntegrations {
	if m == nil {
		return (*ImageIntegrations)(nil)
	}
	r := new(ImageIntegrations)
	if rhs := m.UpdatedIntegrations; rhs != nil {
		tmpContainer := make([]*storage.ImageIntegration, len(rhs))
		for k, v := range rhs {
			if vtpb, ok := interface{}(v).(interface {
				CloneVT() *storage.ImageIntegration
			}); ok {
				tmpContainer[k] = vtpb.CloneVT()
			} else {
				tmpContainer[k] = proto.Clone(v).(*storage.ImageIntegration)
			}
		}
		r.UpdatedIntegrations = tmpContainer
	}
	if rhs := m.DeletedIntegrationIds; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.DeletedIntegrationIds = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ImageIntegrations) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ScanImage) EqualVT(that *ScanImage) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RequestId != that.RequestId {
		return false
	}
	if this.ImageName != that.ImageName {
		return false
	}
	if this.Force != that.Force {
		return false
	}
	if this.Namespace != that.Namespace {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ScanImage) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ScanImage)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ImageIntegrations) EqualVT(that *ImageIntegrations) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.UpdatedIntegrations) != len(that.UpdatedIntegrations) {
		return false
	}
	for i, vx := range this.UpdatedIntegrations {
		vy := that.UpdatedIntegrations[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &storage.ImageIntegration{}
			}
			if q == nil {
				q = &storage.ImageIntegration{}
			}
			if equal, ok := interface{}(p).(interface {
				EqualVT(*storage.ImageIntegration) bool
			}); ok {
				if !equal.EqualVT(q) {
					return false
				}
			} else if !proto.Equal(p, q) {
				return false
			}
		}
	}
	if len(this.DeletedIntegrationIds) != len(that.DeletedIntegrationIds) {
		return false
	}
	for i, vx := range this.DeletedIntegrationIds {
		vy := that.DeletedIntegrationIds[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ImageIntegrations) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ImageIntegrations)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *ScanImage) MarshalVT() (dAtA []byte, err error) {
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

func (m *ScanImage) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ScanImage) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x22
	}
	if m.Force {
		i--
		if m.Force {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ImageName) > 0 {
		i -= len(m.ImageName)
		copy(dAtA[i:], m.ImageName)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ImageName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ImageIntegrations) MarshalVT() (dAtA []byte, err error) {
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

func (m *ImageIntegrations) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ImageIntegrations) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.DeletedIntegrationIds) > 0 {
		for iNdEx := len(m.DeletedIntegrationIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DeletedIntegrationIds[iNdEx])
			copy(dAtA[i:], m.DeletedIntegrationIds[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DeletedIntegrationIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.UpdatedIntegrations) > 0 {
		for iNdEx := len(m.UpdatedIntegrations) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.UpdatedIntegrations[iNdEx]).(interface {
				MarshalToSizedBufferVT([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.UpdatedIntegrations[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ScanImage) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.ImageName)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Force {
		n += 2
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ImageIntegrations) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.UpdatedIntegrations) > 0 {
		for _, e := range m.UpdatedIntegrations {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.DeletedIntegrationIds) > 0 {
		for _, s := range m.DeletedIntegrationIds {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ScanImage) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: ScanImage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScanImage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
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
			m.RequestId = stringValue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageName", wireType)
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
			m.ImageName = stringValue
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Force", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Force = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
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
			m.Namespace = stringValue
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
func (m *ImageIntegrations) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: ImageIntegrations: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageIntegrations: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedIntegrations", wireType)
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
			m.UpdatedIntegrations = append(m.UpdatedIntegrations, &storage.ImageIntegration{})
			if unmarshal, ok := interface{}(m.UpdatedIntegrations[len(m.UpdatedIntegrations)-1]).(interface {
				UnmarshalVTUnsafe([]byte) error
			}); ok {
				if err := unmarshal.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				if err := proto.Unmarshal(dAtA[iNdEx:postIndex], m.UpdatedIntegrations[len(m.UpdatedIntegrations)-1]); err != nil {
					return err
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedIntegrationIds", wireType)
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
			m.DeletedIntegrationIds = append(m.DeletedIntegrationIds, stringValue)
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
