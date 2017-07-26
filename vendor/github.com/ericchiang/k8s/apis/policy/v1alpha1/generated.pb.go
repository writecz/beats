// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/policy/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/policy/v1alpha1/generated.proto

	It has these top-level messages:
		Eviction
		PodDisruptionBudget
		PodDisruptionBudgetList
		PodDisruptionBudgetSpec
		PodDisruptionBudgetStatus
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/ericchiang/k8s/api/resource"
import k8s_io_kubernetes_pkg_api_unversioned "github.com/ericchiang/k8s/api/unversioned"
import k8s_io_kubernetes_pkg_api_v1 "github.com/ericchiang/k8s/api/v1"
import _ "github.com/ericchiang/k8s/runtime"
import k8s_io_kubernetes_pkg_util_intstr "github.com/ericchiang/k8s/util/intstr"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Eviction evicts a pod from its node subject to certain policies and safety constraints.
// This is a subresource of Pod.  A request to cause such an eviction is
// created by POSTing to .../pods/<pod name>/evictions.
type Eviction struct {
	// ObjectMeta describes the pod that is being evicted.
	Metadata *k8s_io_kubernetes_pkg_api_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// DeleteOptions may be provided
	DeleteOptions    *k8s_io_kubernetes_pkg_api_v1.DeleteOptions `protobuf:"bytes,2,opt,name=deleteOptions" json:"deleteOptions,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *Eviction) Reset()                    { *m = Eviction{} }
func (m *Eviction) String() string            { return proto.CompactTextString(m) }
func (*Eviction) ProtoMessage()               {}
func (*Eviction) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *Eviction) GetMetadata() *k8s_io_kubernetes_pkg_api_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Eviction) GetDeleteOptions() *k8s_io_kubernetes_pkg_api_v1.DeleteOptions {
	if m != nil {
		return m.DeleteOptions
	}
	return nil
}

// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
type PodDisruptionBudget struct {
	Metadata *k8s_io_kubernetes_pkg_api_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec *PodDisruptionBudgetSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Most recently observed status of the PodDisruptionBudget.
	Status           *PodDisruptionBudgetStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *PodDisruptionBudget) Reset()                    { *m = PodDisruptionBudget{} }
func (m *PodDisruptionBudget) String() string            { return proto.CompactTextString(m) }
func (*PodDisruptionBudget) ProtoMessage()               {}
func (*PodDisruptionBudget) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *PodDisruptionBudget) GetMetadata() *k8s_io_kubernetes_pkg_api_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PodDisruptionBudget) GetSpec() *PodDisruptionBudgetSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PodDisruptionBudget) GetStatus() *PodDisruptionBudgetStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
type PodDisruptionBudgetList struct {
	Metadata         *k8s_io_kubernetes_pkg_api_unversioned.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Items            []*PodDisruptionBudget                          `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *PodDisruptionBudgetList) Reset()                    { *m = PodDisruptionBudgetList{} }
func (m *PodDisruptionBudgetList) String() string            { return proto.CompactTextString(m) }
func (*PodDisruptionBudgetList) ProtoMessage()               {}
func (*PodDisruptionBudgetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *PodDisruptionBudgetList) GetMetadata() *k8s_io_kubernetes_pkg_api_unversioned.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PodDisruptionBudgetList) GetItems() []*PodDisruptionBudget {
	if m != nil {
		return m.Items
	}
	return nil
}

// PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
type PodDisruptionBudgetSpec struct {
	// The minimum number of pods that must be available simultaneously.  This
	// can be either an integer or a string specifying a percentage, e.g. "28%".
	MinAvailable *k8s_io_kubernetes_pkg_util_intstr.IntOrString `protobuf:"bytes,1,opt,name=minAvailable" json:"minAvailable,omitempty"`
	// Label query over pods whose evictions are managed by the disruption
	// budget.
	Selector         *k8s_io_kubernetes_pkg_api_unversioned.LabelSelector `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
	XXX_unrecognized []byte                                               `json:"-"`
}

func (m *PodDisruptionBudgetSpec) Reset()                    { *m = PodDisruptionBudgetSpec{} }
func (m *PodDisruptionBudgetSpec) String() string            { return proto.CompactTextString(m) }
func (*PodDisruptionBudgetSpec) ProtoMessage()               {}
func (*PodDisruptionBudgetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *PodDisruptionBudgetSpec) GetMinAvailable() *k8s_io_kubernetes_pkg_util_intstr.IntOrString {
	if m != nil {
		return m.MinAvailable
	}
	return nil
}

func (m *PodDisruptionBudgetSpec) GetSelector() *k8s_io_kubernetes_pkg_api_unversioned.LabelSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

// PodDisruptionBudgetStatus represents information about the status of a
// PodDisruptionBudget. Status may trail the actual state of a system.
type PodDisruptionBudgetStatus struct {
	// Whether or not a disruption is currently allowed.
	DisruptionAllowed *bool `protobuf:"varint,1,opt,name=disruptionAllowed" json:"disruptionAllowed,omitempty"`
	// current number of healthy pods
	CurrentHealthy *int32 `protobuf:"varint,2,opt,name=currentHealthy" json:"currentHealthy,omitempty"`
	// minimum desired number of healthy pods
	DesiredHealthy *int32 `protobuf:"varint,3,opt,name=desiredHealthy" json:"desiredHealthy,omitempty"`
	// total number of pods counted by this disruption budget
	ExpectedPods     *int32 `protobuf:"varint,4,opt,name=expectedPods" json:"expectedPods,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PodDisruptionBudgetStatus) Reset()         { *m = PodDisruptionBudgetStatus{} }
func (m *PodDisruptionBudgetStatus) String() string { return proto.CompactTextString(m) }
func (*PodDisruptionBudgetStatus) ProtoMessage()    {}
func (*PodDisruptionBudgetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{4}
}

func (m *PodDisruptionBudgetStatus) GetDisruptionAllowed() bool {
	if m != nil && m.DisruptionAllowed != nil {
		return *m.DisruptionAllowed
	}
	return false
}

func (m *PodDisruptionBudgetStatus) GetCurrentHealthy() int32 {
	if m != nil && m.CurrentHealthy != nil {
		return *m.CurrentHealthy
	}
	return 0
}

func (m *PodDisruptionBudgetStatus) GetDesiredHealthy() int32 {
	if m != nil && m.DesiredHealthy != nil {
		return *m.DesiredHealthy
	}
	return 0
}

func (m *PodDisruptionBudgetStatus) GetExpectedPods() int32 {
	if m != nil && m.ExpectedPods != nil {
		return *m.ExpectedPods
	}
	return 0
}

func init() {
	proto.RegisterType((*Eviction)(nil), "github.com/ericchiang.k8s.apis.policy.v1alpha1.Eviction")
	proto.RegisterType((*PodDisruptionBudget)(nil), "github.com/ericchiang.k8s.apis.policy.v1alpha1.PodDisruptionBudget")
	proto.RegisterType((*PodDisruptionBudgetList)(nil), "github.com/ericchiang.k8s.apis.policy.v1alpha1.PodDisruptionBudgetList")
	proto.RegisterType((*PodDisruptionBudgetSpec)(nil), "github.com/ericchiang.k8s.apis.policy.v1alpha1.PodDisruptionBudgetSpec")
	proto.RegisterType((*PodDisruptionBudgetStatus)(nil), "github.com/ericchiang.k8s.apis.policy.v1alpha1.PodDisruptionBudgetStatus")
}
func (m *Eviction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Eviction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.DeleteOptions != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.DeleteOptions.Size()))
		n2, err := m.DeleteOptions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PodDisruptionBudget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudget) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n3, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Spec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
		n4, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Status != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
		n5, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PodDisruptionBudgetList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n6, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PodDisruptionBudgetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MinAvailable != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.MinAvailable.Size()))
		n7, err := m.MinAvailable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Selector != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Selector.Size()))
		n8, err := m.Selector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PodDisruptionBudgetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodDisruptionBudgetStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DisruptionAllowed != nil {
		dAtA[i] = 0x8
		i++
		if *m.DisruptionAllowed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CurrentHealthy != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.CurrentHealthy))
	}
	if m.DesiredHealthy != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.DesiredHealthy))
	}
	if m.ExpectedPods != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.ExpectedPods))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Eviction) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.DeleteOptions != nil {
		l = m.DeleteOptions.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PodDisruptionBudget) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PodDisruptionBudgetList) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PodDisruptionBudgetSpec) Size() (n int) {
	var l int
	_ = l
	if m.MinAvailable != nil {
		l = m.MinAvailable.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PodDisruptionBudgetStatus) Size() (n int) {
	var l int
	_ = l
	if m.DisruptionAllowed != nil {
		n += 2
	}
	if m.CurrentHealthy != nil {
		n += 1 + sovGenerated(uint64(*m.CurrentHealthy))
	}
	if m.DesiredHealthy != nil {
		n += 1 + sovGenerated(uint64(*m.DesiredHealthy))
	}
	if m.ExpectedPods != nil {
		n += 1 + sovGenerated(uint64(*m.ExpectedPods))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Eviction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Eviction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Eviction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_kubernetes_pkg_api_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeleteOptions == nil {
				m.DeleteOptions = &k8s_io_kubernetes_pkg_api_v1.DeleteOptions{}
			}
			if err := m.DeleteOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_kubernetes_pkg_api_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &PodDisruptionBudgetSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &PodDisruptionBudgetStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_kubernetes_pkg_api_unversioned.ListMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &PodDisruptionBudget{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAvailable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinAvailable == nil {
				m.MinAvailable = &k8s_io_kubernetes_pkg_util_intstr.IntOrString{}
			}
			if err := m.MinAvailable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_kubernetes_pkg_api_unversioned.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodDisruptionBudgetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodDisruptionBudgetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodDisruptionBudgetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisruptionAllowed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.DisruptionAllowed = &b
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentHealthy", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CurrentHealthy = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredHealthy", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DesiredHealthy = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedPods", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ExpectedPods = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/ericchiang/k8s/apis/policy/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x8a, 0x53, 0x31,
	0x14, 0x80, 0xbd, 0xf3, 0x23, 0x25, 0x33, 0x0a, 0xc6, 0x85, 0xb5, 0x8b, 0x22, 0x5d, 0x48, 0xd1,
	0x31, 0x97, 0x16, 0x05, 0x71, 0x23, 0x33, 0x76, 0x40, 0x51, 0x69, 0x4d, 0x11, 0x41, 0x70, 0x91,
	0xde, 0x1c, 0x3a, 0xb1, 0x69, 0x12, 0x92, 0x73, 0xaf, 0xce, 0x73, 0xb8, 0x71, 0xef, 0x73, 0x88,
	0x5b, 0x97, 0x3e, 0x82, 0xd4, 0x17, 0x91, 0x7b, 0xfb, 0xe3, 0xf4, 0xe7, 0x0e, 0x03, 0xe3, 0xf6,
	0xe4, 0x7c, 0x5f, 0xce, 0x1f, 0x79, 0x32, 0x7a, 0x1c, 0x98, 0xb2, 0xf1, 0x28, 0x1d, 0x80, 0x37,
	0x80, 0x10, 0x62, 0x37, 0x1a, 0xc6, 0xc2, 0xa9, 0x10, 0x3b, 0xab, 0x55, 0x72, 0x1a, 0x67, 0x2d,
	0xa1, 0xdd, 0x89, 0x68, 0xc5, 0x43, 0x30, 0xe0, 0x05, 0x82, 0x64, 0xce, 0x5b, 0xb4, 0xf4, 0xde,
	0x94, 0x65, 0xff, 0x58, 0xe6, 0x46, 0x43, 0x96, 0xb3, 0x6c, 0xca, 0xb2, 0x39, 0x5b, 0x6b, 0x97,
	0xfe, 0x13, 0x7b, 0x08, 0x36, 0xf5, 0x09, 0xac, 0xfa, 0x6b, 0x8f, 0xca, 0x99, 0xd4, 0x64, 0xe0,
	0x83, 0xb2, 0x06, 0xe4, 0x1a, 0x76, 0x50, 0x8e, 0x65, 0x6b, 0x4d, 0xd4, 0x1e, 0x6c, 0xce, 0xf6,
	0xa9, 0x41, 0x35, 0x5e, 0xaf, 0xa9, 0xb5, 0x39, 0x3d, 0x45, 0xa5, 0x63, 0x65, 0x30, 0xa0, 0x5f,
	0x45, 0x1a, 0xdf, 0x22, 0x52, 0x39, 0xce, 0x54, 0x82, 0xca, 0x1a, 0xda, 0x21, 0x95, 0x31, 0xa0,
	0x90, 0x02, 0x45, 0x35, 0xba, 0x13, 0x35, 0xf7, 0xda, 0x4d, 0x56, 0x3a, 0x46, 0x96, 0xb5, 0x58,
	0x77, 0xf0, 0x11, 0x12, 0x7c, 0x0d, 0x28, 0xf8, 0x82, 0xa4, 0x6f, 0xc8, 0x35, 0x09, 0x1a, 0x10,
	0xba, 0x2e, 0xb7, 0x86, 0xea, 0x56, 0xa1, 0xba, 0x7f, 0xbe, 0xaa, 0x73, 0x16, 0xe1, 0xcb, 0x86,
	0xc6, 0x97, 0x2d, 0x72, 0xb3, 0x67, 0x65, 0x47, 0x05, 0x9f, 0x16, 0xa1, 0xa3, 0x54, 0x0e, 0x01,
	0xff, 0x53, 0xc1, 0xef, 0xc8, 0x4e, 0x70, 0x90, 0xcc, 0xea, 0x7c, 0xc6, 0x2e, 0x7e, 0x39, 0x6c,
	0x43, 0x51, 0x7d, 0x07, 0x09, 0x2f, 0x84, 0xf4, 0x03, 0xb9, 0x1a, 0x50, 0x60, 0x1a, 0xaa, 0xdb,
	0x85, 0xfa, 0xf8, 0xb2, 0xea, 0x42, 0xc6, 0x67, 0xd2, 0xc6, 0xf7, 0x88, 0xdc, 0xda, 0x90, 0xf5,
	0x4a, 0x05, 0xa4, 0x2f, 0xd7, 0x26, 0x13, 0x9f, 0x33, 0x99, 0x33, 0x17, 0xcb, 0x72, 0x7c, 0x65,
	0x40, 0x6f, 0xc9, 0xae, 0x42, 0x18, 0xe7, 0x9b, 0xdc, 0x6e, 0xee, 0xb5, 0x9f, 0x5e, 0xb2, 0x0d,
	0x3e, 0xb5, 0x35, 0x7e, 0x6c, 0xae, 0x3f, 0x1f, 0x20, 0xe5, 0x64, 0x7f, 0xac, 0xcc, 0x61, 0x26,
	0x94, 0x16, 0x03, 0x0d, 0xb3, 0x1e, 0x58, 0xc9, 0xcf, 0xf9, 0x85, 0xb3, 0xe9, 0x85, 0xb3, 0x17,
	0x06, 0xbb, 0xbe, 0x8f, 0x5e, 0x99, 0x21, 0x5f, 0x72, 0xd0, 0x1e, 0xa9, 0x04, 0xd0, 0x90, 0xa0,
	0xf5, 0xb3, 0x5d, 0x3f, 0xbc, 0xe8, 0x4c, 0xc4, 0x00, 0x74, 0x7f, 0xc6, 0xf2, 0x85, 0x25, 0xdf,
	0xc0, 0xed, 0xd2, 0x3d, 0xd1, 0x03, 0x72, 0x43, 0x2e, 0x5e, 0x0e, 0xb5, 0xb6, 0x9f, 0x40, 0x16,
	0x8d, 0x54, 0xf8, 0xfa, 0x03, 0xbd, 0x4b, 0xae, 0x27, 0xa9, 0xf7, 0x60, 0xf0, 0x39, 0x08, 0x8d,
	0x27, 0xa7, 0x45, 0x8d, 0xbb, 0x7c, 0x25, 0x9a, 0xe7, 0x49, 0x08, 0xca, 0x83, 0x9c, 0xe7, 0x6d,
	0x4f, 0xf3, 0x96, 0xa3, 0xb4, 0x41, 0xf6, 0xe1, 0xb3, 0x83, 0x04, 0x41, 0xf6, 0xac, 0x0c, 0xd5,
	0x9d, 0x22, 0x6b, 0x29, 0x76, 0x54, 0xfb, 0x39, 0xa9, 0x47, 0xbf, 0x26, 0xf5, 0xe8, 0xf7, 0xa4,
	0x1e, 0x7d, 0xfd, 0x53, 0xbf, 0xf2, 0xbe, 0x32, 0x5f, 0xdc, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0xc2, 0x95, 0x04, 0x7d, 0x05, 0x00, 0x00,
}