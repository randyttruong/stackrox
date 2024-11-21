// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.1-0.20240409071808-615f978279ca
// source: storage/service_identity.proto

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

func (m *ServiceIdentity) CloneVT() *ServiceIdentity {
	if m == nil {
		return (*ServiceIdentity)(nil)
	}
	r := new(ServiceIdentity)
	r.SerialStr = m.SerialStr
	r.Id = m.Id
	r.Type = m.Type
	r.InitBundleId = m.InitBundleId
	if m.Srl != nil {
		r.Srl = m.Srl.(interface{ CloneVT() isServiceIdentity_Srl }).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ServiceIdentity) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ServiceIdentity_Serial) CloneVT() isServiceIdentity_Srl {
	if m == nil {
		return (*ServiceIdentity_Serial)(nil)
	}
	r := new(ServiceIdentity_Serial)
	r.Serial = m.Serial
	return r
}

func (m *ServiceCertificate) CloneVT() *ServiceCertificate {
	if m == nil {
		return (*ServiceCertificate)(nil)
	}
	r := new(ServiceCertificate)
	if rhs := m.CertPem; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.CertPem = tmpBytes
	}
	if rhs := m.KeyPem; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.KeyPem = tmpBytes
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ServiceCertificate) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *TypedServiceCertificate) CloneVT() *TypedServiceCertificate {
	if m == nil {
		return (*TypedServiceCertificate)(nil)
	}
	r := new(TypedServiceCertificate)
	r.ServiceType = m.ServiceType
	r.Cert = m.Cert.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *TypedServiceCertificate) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *TypedServiceCertificateSet) CloneVT() *TypedServiceCertificateSet {
	if m == nil {
		return (*TypedServiceCertificateSet)(nil)
	}
	r := new(TypedServiceCertificateSet)
	if rhs := m.CaPem; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.CaPem = tmpBytes
	}
	if rhs := m.ServiceCerts; rhs != nil {
		tmpContainer := make([]*TypedServiceCertificate, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ServiceCerts = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *TypedServiceCertificateSet) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ServiceIdentity) EqualVT(that *ServiceIdentity) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Srl == nil && that.Srl != nil {
		return false
	} else if this.Srl != nil {
		if that.Srl == nil {
			return false
		}
		if !this.Srl.(interface {
			EqualVT(isServiceIdentity_Srl) bool
		}).EqualVT(that.Srl) {
			return false
		}
	}
	if this.Id != that.Id {
		return false
	}
	if this.Type != that.Type {
		return false
	}
	if this.SerialStr != that.SerialStr {
		return false
	}
	if this.InitBundleId != that.InitBundleId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ServiceIdentity) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ServiceIdentity)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ServiceIdentity_Serial) EqualVT(thatIface isServiceIdentity_Srl) bool {
	that, ok := thatIface.(*ServiceIdentity_Serial)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Serial != that.Serial {
		return false
	}
	return true
}

func (this *ServiceCertificate) EqualVT(that *ServiceCertificate) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.CertPem) != string(that.CertPem) {
		return false
	}
	if string(this.KeyPem) != string(that.KeyPem) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ServiceCertificate) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ServiceCertificate)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *TypedServiceCertificate) EqualVT(that *TypedServiceCertificate) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ServiceType != that.ServiceType {
		return false
	}
	if !this.Cert.EqualVT(that.Cert) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *TypedServiceCertificate) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*TypedServiceCertificate)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *TypedServiceCertificateSet) EqualVT(that *TypedServiceCertificateSet) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.CaPem) != string(that.CaPem) {
		return false
	}
	if len(this.ServiceCerts) != len(that.ServiceCerts) {
		return false
	}
	for i, vx := range this.ServiceCerts {
		vy := that.ServiceCerts[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &TypedServiceCertificate{}
			}
			if q == nil {
				q = &TypedServiceCertificate{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *TypedServiceCertificateSet) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*TypedServiceCertificateSet)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *ServiceIdentity) MarshalVT() (dAtA []byte, err error) {
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

func (m *ServiceIdentity) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ServiceIdentity) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if vtmsg, ok := m.Srl.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.InitBundleId) > 0 {
		i -= len(m.InitBundleId)
		copy(dAtA[i:], m.InitBundleId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.InitBundleId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SerialStr) > 0 {
		i -= len(m.SerialStr)
		copy(dAtA[i:], m.SerialStr)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.SerialStr)))
		i--
		dAtA[i] = 0x22
	}
	if m.Type != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ServiceIdentity_Serial) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ServiceIdentity_Serial) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Serial))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *ServiceCertificate) MarshalVT() (dAtA []byte, err error) {
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

func (m *ServiceCertificate) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ServiceCertificate) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.KeyPem) > 0 {
		i -= len(m.KeyPem)
		copy(dAtA[i:], m.KeyPem)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.KeyPem)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CertPem) > 0 {
		i -= len(m.CertPem)
		copy(dAtA[i:], m.CertPem)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.CertPem)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TypedServiceCertificate) MarshalVT() (dAtA []byte, err error) {
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

func (m *TypedServiceCertificate) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *TypedServiceCertificate) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if m.Cert != nil {
		size, err := m.Cert.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.ServiceType != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.ServiceType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TypedServiceCertificateSet) MarshalVT() (dAtA []byte, err error) {
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

func (m *TypedServiceCertificateSet) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *TypedServiceCertificateSet) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
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
	if len(m.ServiceCerts) > 0 {
		for iNdEx := len(m.ServiceCerts) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ServiceCerts[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CaPem) > 0 {
		i -= len(m.CaPem)
		copy(dAtA[i:], m.CaPem)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.CaPem)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ServiceIdentity) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Srl.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Type))
	}
	l = len(m.SerialStr)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.InitBundleId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ServiceIdentity_Serial) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + protohelpers.SizeOfVarint(uint64(m.Serial))
	return n
}
func (m *ServiceCertificate) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CertPem)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.KeyPem)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *TypedServiceCertificate) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServiceType != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.ServiceType))
	}
	if m.Cert != nil {
		l = m.Cert.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *TypedServiceCertificateSet) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CaPem)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.ServiceCerts) > 0 {
		for _, e := range m.ServiceCerts {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ServiceIdentity) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: ServiceIdentity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceIdentity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Serial", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Srl = &ServiceIdentity_Serial{Serial: v}
		case 2:
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ServiceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialStr", wireType)
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
			m.SerialStr = stringValue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitBundleId", wireType)
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
			m.InitBundleId = stringValue
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
func (m *ServiceCertificate) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: ServiceCertificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceCertificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertPem", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertPem = dAtA[iNdEx:postIndex]
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyPem", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyPem = dAtA[iNdEx:postIndex]
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
func (m *TypedServiceCertificate) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: TypedServiceCertificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TypedServiceCertificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceType", wireType)
			}
			m.ServiceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceType |= ServiceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
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
			if m.Cert == nil {
				m.Cert = &ServiceCertificate{}
			}
			if err := m.Cert.UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
func (m *TypedServiceCertificateSet) UnmarshalVTUnsafe(dAtA []byte) error {
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
			return fmt.Errorf("proto: TypedServiceCertificateSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TypedServiceCertificateSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaPem", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CaPem = dAtA[iNdEx:postIndex]
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceCerts", wireType)
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
			m.ServiceCerts = append(m.ServiceCerts, &TypedServiceCertificate{})
			if err := m.ServiceCerts[len(m.ServiceCerts)-1].UnmarshalVTUnsafe(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
