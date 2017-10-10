// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	Group
	Entity
	Alias
	GroupAlias
*/
package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Group represents an identity group.
type Group struct {
	// ID is the unique identifier for this group
	ID string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Name is the unique name for this group
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Policies are the vault policies to be granted to members of this group
	Policies []string `protobuf:"bytes,3,rep,name=policies" json:"policies,omitempty"`
	// ParentGroupIDs are the identifiers of those groups to which this group is a
	// member of. These will serve as references to the parent group in the
	// hierarchy.
	ParentGroupIDs []string `protobuf:"bytes,4,rep,name=parent_group_ids,json=parentGroupIds" json:"parent_group_ids,omitempty"`
	// MemberEntityIDs are the identifiers of entities which are members of this
	// group
	MemberEntityIDs []string `protobuf:"bytes,5,rep,name=member_entity_ids,json=memberEntityIDs" json:"member_entity_ids,omitempty"`
	// Metadata represents the custom data tied with this group
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// CreationTime is the time at which this group was created
	CreationTime *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the time at which this group was last modified
	LastUpdateTime *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// ModifyIndex tracks the number of updates to the group. It is useful to detect
	// updates to the groups.
	ModifyIndex uint64 `protobuf:"varint,9,opt,name=modify_index,json=modifyIndex" json:"modify_index,omitempty"`
	// BucketKeyHash is the MD5 hash of the storage bucket key into which this
	// group is stored in the underlying storage. This is useful to find all
	// the groups belonging to a particular bucket during invalidation of the
	// storage key.
	BucketKeyHash string      `protobuf:"bytes,10,opt,name=bucket_key_hash,json=bucketKeyHash" json:"bucket_key_hash,omitempty"`
	Alias         *GroupAlias `protobuf:"bytes,11,opt,name=alias" json:"alias,omitempty"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Group) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetPolicies() []string {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *Group) GetParentGroupIDs() []string {
	if m != nil {
		return m.ParentGroupIDs
	}
	return nil
}

func (m *Group) GetMemberEntityIDs() []string {
	if m != nil {
		return m.MemberEntityIDs
	}
	return nil
}

func (m *Group) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Group) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Group) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Group) GetModifyIndex() uint64 {
	if m != nil {
		return m.ModifyIndex
	}
	return 0
}

func (m *Group) GetBucketKeyHash() string {
	if m != nil {
		return m.BucketKeyHash
	}
	return ""
}

func (m *Group) GetAlias() *GroupAlias {
	if m != nil {
		return m.Alias
	}
	return nil
}

// Entity represents an entity that gets persisted and indexed.
// Entity is fundamentally composed of zero or many aliases.
type Entity struct {
	// Aliases are the identities that this entity is made of. This can be
	// empty as well to favor being able to create the entity first and then
	// incrementally adding aliases.
	Aliases []*Alias `protobuf:"bytes,1,rep,name=aliases" json:"aliases,omitempty"`
	// ID is the unique identifier of the entity which always be a UUID. This
	// should never be allowed to be updated.
	ID string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// Name is a unique identifier of the entity which is intended to be
	// human-friendly. The default name might not be human friendly since it
	// gets suffixed by a UUID, but it can optionally be updated, unlike the ID
	// field.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Metadata represents the explicit metadata which is set by the
	// clients.  This is useful to tie any information pertaining to the
	// aliases. This is a non-unique field of entity, meaning multiple
	// entities can have the same metadata set. Entities will be indexed based
	// on this explicit metadata. This enables virtual groupings of entities
	// based on its metadata.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// CreationTime is the time at which this entity is first created.
	CreationTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the most recent time at which the properties of this
	// entity got modified. This is helpful in filtering out entities based on
	// its age and to take action on them, if desired.
	LastUpdateTime *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// MergedEntityIDs are the entities which got merged to this one. Entities
	// will be indexed based on all the entities that got merged into it. This
	// helps to apply the actions on this entity on the tokens that are merged
	// to the merged entities. Merged entities will be deleted entirely and
	// this is the only trackable trail of its earlier presence.
	MergedEntityIDs []string `protobuf:"bytes,7,rep,name=merged_entity_ids,json=mergedEntityIDs" json:"merged_entity_ids,omitempty"`
	// Policies the entity is entitled to
	Policies []string `protobuf:"bytes,8,rep,name=policies" json:"policies,omitempty"`
	// BucketKeyHash is the MD5 hash of the storage bucket key into which this
	// entity is stored in the underlying storage. This is useful to find all
	// the entities belonging to a particular bucket during invalidation of the
	// storage key.
	BucketKeyHash string `protobuf:"bytes,9,opt,name=bucket_key_hash,json=bucketKeyHash" json:"bucket_key_hash,omitempty"`
}

func (m *Entity) Reset()                    { *m = Entity{} }
func (m *Entity) String() string            { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()               {}
func (*Entity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Entity) GetAliases() []*Alias {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Entity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Entity) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Entity) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Entity) GetMergedEntityIDs() []string {
	if m != nil {
		return m.MergedEntityIDs
	}
	return nil
}

func (m *Entity) GetPolicies() []string {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *Entity) GetBucketKeyHash() string {
	if m != nil {
		return m.BucketKeyHash
	}
	return ""
}

// Alias represents the alias that gets stored inside of the
// entity object in storage and also represents in an in-memory index of an
// alias object.
type Alias struct {
	// ID is the unique identifier that represents this alias
	ID string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// EntityID is the entity identifier to which this alias belongs to
	EntityID string `protobuf:"bytes,2,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	// MountType is the backend mount's type to which this alias belongs to.
	// This enables categorically querying aliases of specific backend types.
	MountType string `protobuf:"bytes,3,opt,name=mount_type,json=mountType" json:"mount_type,omitempty"`
	// MountAccessor is the backend mount's accessor to which this alias
	// belongs to.
	MountAccessor string `protobuf:"bytes,4,opt,name=mount_accessor,json=mountAccessor" json:"mount_accessor,omitempty"`
	// MountPath is the backend mount's path to which the Maccessor belongs to. This
	// field is not used for any operational purposes. This is only returned when
	// alias is read, only as a nicety.
	MountPath string `protobuf:"bytes,5,opt,name=mount_path,json=mountPath" json:"mount_path,omitempty"`
	// Metadata is the explicit metadata that clients set against an entity
	// which enables virtual grouping of aliases. Aliases will be indexed
	// against their metadata.
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Name is the identifier of this alias in its authentication source.
	// This does not uniquely identify an alias in Vault. This in conjunction
	// with MountAccessor form to be the factors that represent an alias in a
	// unique way. Aliases will be indexed based on this combined uniqueness
	// factor.
	Name string `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	// CreationTime is the time at which this alias was first created
	CreationTime *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the most recent time at which the properties of this
	// alias got modified. This is helpful in filtering out aliases based
	// on its age and to take action on them, if desired.
	LastUpdateTime *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// MergedFromEntityIDs is the FIFO history of merging activity by entity IDs from
	// which this alias is transfered over to the entity to which it
	// currently belongs to.
	MergedFromEntityIDs []string `protobuf:"bytes,10,rep,name=merged_from_entity_ids,json=mergedFromEntityIDs" json:"merged_from_entity_ids,omitempty"`
}

func (m *Alias) Reset()                    { *m = Alias{} }
func (m *Alias) String() string            { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()               {}
func (*Alias) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Alias) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Alias) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *Alias) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *Alias) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *Alias) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *Alias) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Alias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Alias) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Alias) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Alias) GetMergedFromEntityIDs() []string {
	if m != nil {
		return m.MergedFromEntityIDs
	}
	return nil
}

type GroupAlias struct {
	ID             string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name           string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	GroupID        string                     `protobuf:"bytes,3,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	MountType      string                     `protobuf:"bytes,4,opt,name=mount_type,json=mountType" json:"mount_type,omitempty"`
	MountAccessor  string                     `protobuf:"bytes,5,opt,name=mount_accessor,json=mountAccessor" json:"mount_accessor,omitempty"`
	CreationTime   *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	LastUpdateTime *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
}

func (m *GroupAlias) Reset()                    { *m = GroupAlias{} }
func (m *GroupAlias) String() string            { return proto.CompactTextString(m) }
func (*GroupAlias) ProtoMessage()               {}
func (*GroupAlias) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GroupAlias) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GroupAlias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupAlias) GetGroupID() string {
	if m != nil {
		return m.GroupID
	}
	return ""
}

func (m *GroupAlias) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *GroupAlias) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *GroupAlias) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *GroupAlias) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Group)(nil), "identity.Group")
	proto.RegisterType((*Entity)(nil), "identity.Entity")
	proto.RegisterType((*Alias)(nil), "identity.Alias")
	proto.RegisterType((*GroupAlias)(nil), "identity.GroupAlias")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0xd5, 0x26, 0x69, 0x92, 0xd3, 0xad, 0xdb, 0xe7, 0x6f, 0x42, 0xa1, 0x68, 0x50, 0x26,
	0x81, 0xc2, 0x2e, 0x32, 0x69, 0xbb, 0x81, 0x71, 0x81, 0x26, 0x31, 0x60, 0x42, 0x48, 0x28, 0x1a,
	0xd7, 0x91, 0xdb, 0x78, 0xad, 0xb5, 0x26, 0x8e, 0x62, 0x07, 0x91, 0xe7, 0xe0, 0x86, 0x37, 0xe2,
	0x39, 0x78, 0x13, 0x94, 0xe3, 0xa4, 0x0d, 0xeb, 0x60, 0x9b, 0xd6, 0xbb, 0xe4, 0x7f, 0x8e, 0x8f,
	0x8f, 0xcf, 0xef, 0x6f, 0x43, 0x5f, 0x95, 0x19, 0x93, 0x41, 0x96, 0x0b, 0x25, 0x88, 0xc3, 0x63,
	0x96, 0x2a, 0xae, 0xca, 0xe1, 0x93, 0xa9, 0x10, 0xd3, 0x39, 0x3b, 0x40, 0x7d, 0x5c, 0x5c, 0x1c,
	0x28, 0x9e, 0x30, 0xa9, 0x68, 0x92, 0xe9, 0xd4, 0xbd, 0xef, 0x26, 0x58, 0xef, 0x73, 0x51, 0x64,
	0x64, 0x00, 0x5d, 0x1e, 0x7b, 0x9d, 0x51, 0xc7, 0x77, 0xc3, 0x2e, 0x8f, 0x09, 0x01, 0x33, 0xa5,
	0x09, 0xf3, 0xba, 0xa8, 0xe0, 0x37, 0x19, 0x82, 0x93, 0x89, 0x39, 0x9f, 0x70, 0x26, 0x3d, 0x63,
	0x64, 0xf8, 0x6e, 0xb8, 0xf8, 0x27, 0x3e, 0x6c, 0x67, 0x34, 0x67, 0xa9, 0x8a, 0xa6, 0x55, 0xbd,
	0x88, 0xc7, 0xd2, 0x33, 0x31, 0x67, 0xa0, 0x75, 0xdc, 0xe6, 0x2c, 0x96, 0x64, 0x1f, 0xfe, 0x4b,
	0x58, 0x32, 0x66, 0x79, 0xa4, 0xbb, 0xc4, 0x54, 0x0b, 0x53, 0xb7, 0x74, 0xe0, 0x14, 0xf5, 0x2a,
	0xf7, 0x15, 0x38, 0x09, 0x53, 0x34, 0xa6, 0x8a, 0x7a, 0xbd, 0x91, 0xe1, 0xf7, 0x0f, 0x77, 0x83,
	0xe6, 0x74, 0x01, 0x56, 0x0c, 0x3e, 0xd5, 0xf1, 0xd3, 0x54, 0xe5, 0x65, 0xb8, 0x48, 0x27, 0x6f,
	0x60, 0x73, 0x92, 0x33, 0xaa, 0xb8, 0x48, 0xa3, 0xea, 0xd8, 0x9e, 0x3d, 0xea, 0xf8, 0xfd, 0xc3,
	0x61, 0xa0, 0x67, 0x12, 0x34, 0x33, 0x09, 0xce, 0x9b, 0x99, 0x84, 0x1b, 0xcd, 0x82, 0x4a, 0x22,
	0x6f, 0x61, 0x7b, 0x4e, 0xa5, 0x8a, 0x8a, 0x2c, 0xa6, 0x8a, 0xe9, 0x1a, 0xce, 0x8d, 0x35, 0x06,
	0xd5, 0x9a, 0x2f, 0xb8, 0x04, 0xab, 0x3c, 0x85, 0x8d, 0x44, 0xc4, 0xfc, 0xa2, 0x8c, 0x78, 0x1a,
	0xb3, 0x6f, 0x9e, 0x3b, 0xea, 0xf8, 0x66, 0xd8, 0xd7, 0xda, 0x59, 0x25, 0x91, 0xe7, 0xb0, 0x35,
	0x2e, 0x26, 0x97, 0x4c, 0x45, 0x97, 0xac, 0x8c, 0x66, 0x54, 0xce, 0x3c, 0xc0, 0xa9, 0x6f, 0x6a,
	0xf9, 0x23, 0x2b, 0x3f, 0x50, 0x39, 0x23, 0xfb, 0x60, 0xd1, 0x39, 0xa7, 0xd2, 0xeb, 0x63, 0x17,
	0x3b, 0x57, 0x26, 0x71, 0x52, 0xc5, 0x42, 0x9d, 0x32, 0x7c, 0x0d, 0x9b, 0x7f, 0x0c, 0x86, 0x6c,
	0x83, 0x71, 0xc9, 0xca, 0x1a, 0x70, 0xf5, 0x49, 0x76, 0xc0, 0xfa, 0x4a, 0xe7, 0x45, 0x83, 0x58,
	0xff, 0x1c, 0x77, 0x5f, 0x76, 0xf6, 0x7e, 0x1a, 0xd0, 0xd3, 0x0c, 0xc8, 0x0b, 0xb0, 0xb1, 0x20,
	0x93, 0x5e, 0x07, 0xe7, 0xbf, 0xb5, 0xdc, 0x55, 0x6f, 0xd8, 0xc4, 0x6b, 0x07, 0x75, 0x57, 0x1c,
	0x64, 0xb4, 0x1c, 0x74, 0xdc, 0xe2, 0x69, 0x62, 0xbd, 0xc7, 0xcb, 0x7a, 0x7a, 0xcb, 0xdb, 0x03,
	0xb5, 0xd6, 0x00, 0xb4, 0x77, 0x67, 0xa0, 0x68, 0xdf, 0x7c, 0xca, 0xe2, 0xb6, 0x7d, 0xed, 0xc6,
	0xbe, 0x55, 0x60, 0x69, 0xdf, 0xf6, 0x85, 0x71, 0xae, 0x5c, 0x98, 0x6b, 0xa8, 0xbb, 0xd7, 0x50,
	0xbf, 0x1f, 0xc9, 0x5f, 0x06, 0x58, 0x88, 0x69, 0xe5, 0x7e, 0x3f, 0x02, 0x77, 0xd1, 0x7f, 0xbd,
	0xce, 0x61, 0x75, 0xe3, 0x64, 0x17, 0x20, 0x11, 0x45, 0xaa, 0xa2, 0xea, 0x59, 0xa9, 0x01, 0xba,
	0xa8, 0x9c, 0x97, 0x19, 0x23, 0xcf, 0x60, 0xa0, 0xc3, 0x74, 0x32, 0x61, 0x52, 0x8a, 0xdc, 0x33,
	0x75, 0xe7, 0xa8, 0x9e, 0xd4, 0xe2, 0xb2, 0x4a, 0x46, 0xd5, 0x0c, 0x69, 0x35, 0x55, 0x3e, 0x53,
	0x35, 0xfb, 0xf7, 0xdd, 0xc6, 0xa6, 0xff, 0x6a, 0x85, 0xc6, 0x5a, 0x76, 0xcb, 0x5a, 0x2b, 0xf6,
	0x70, 0xd6, 0x60, 0x0f, 0xf7, 0xce, 0xf6, 0x38, 0x82, 0x07, 0xb5, 0x3d, 0x2e, 0x72, 0x91, 0xb4,
	0x3d, 0x02, 0x68, 0x80, 0xff, 0x75, 0xf4, 0x5d, 0x2e, 0x92, 0x85, 0x4f, 0xee, 0xc7, 0xf8, 0x47,
	0x17, 0x60, 0xf9, 0x00, 0xdc, 0xea, 0x21, 0x7f, 0x08, 0x4e, 0xf3, 0x4a, 0xd7, 0x74, 0xed, 0xa9,
	0x7e, 0x9e, 0xaf, 0xa0, 0x37, 0x6f, 0x46, 0x6f, 0x5d, 0x87, 0x7e, 0x05, 0x46, 0x6f, 0x0d, 0x30,
	0xec, 0xbb, 0xc2, 0x18, 0xf7, 0x30, 0xe7, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x8a,
	0xe0, 0xf0, 0x1f, 0x07, 0x00, 0x00,
}
