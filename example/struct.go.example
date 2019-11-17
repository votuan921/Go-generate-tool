package mysql

import "github.com/moneyforward/ca_backend/app/pkg/protoext"

type Item struct {
	// @inject_tag: db:"id"
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id"`
	// @inject_tag: db:"type"
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" db:"type"`
	// @inject_tag: db:"fs"
	Fs int32 `protobuf:"varint,3,opt,name=fs,proto3" json:"fs,omitempty" db:"fs"`
	// @inject_tag: db:"account_group"
	AccountGroup int32 `protobuf:"varint,4,opt,name=account_group,json=accountGroup,proto3" json:"account_group,omitempty" db:"account_group"`
	// @inject_tag: db:"account_property"
	AccountProperty int32 `protobuf:"varint,5,opt,name=account_property,json=accountProperty,proto3" json:"account_property,omitempty" db:"account_property"`
	// @inject_tag: db:"office_id"
	OfficeId uint64 `protobuf:"varint,6,opt,name=office_id,json=officeId,proto3" json:"office_id,omitempty" db:"office_id"`
	// @inject_tag: db:"excise_id"
	ExciseId uint64 `protobuf:"varint,7,opt,name=excise_id,json=exciseId,proto3" json:"excise_id,omitempty" db:"excise_id"`
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty" db:"name"`
	// @inject_tag: db:"yomigana"
	Yomigana *protoext.NullString `protobuf:"bytes,9,opt,name=yomigana,proto3" json:"yomigana,omitempty" db:"yomigana"`
	// @inject_tag: db:"romaji"
	Romaji *protoext.NullString `protobuf:"bytes,10,opt,name=romaji,proto3" json:"romaji,omitempty" db:"romaji"`
	// @inject_tag: db:"code"
	Code *protoext.NullString `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty" db:"code"`
	// @inject_tag: db:"side"
	Side int32 `protobuf:"varint,12,opt,name=side,proto3" json:"side,omitempty" db:"side"`
	// @inject_tag: db:"is_active"
	IsActive bool `protobuf:"varint,13,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" db:"is_active"`
	// @inject_tag: db:"disp_order"
	DispOrder uint64 `protobuf:"varint,14,opt,name=disp_order,json=dispOrder,proto3" json:"disp_order,omitempty" db:"disp_order"`
	// @inject_tag: db:"item_master_id"
	ItemMasterId *protoext.NullInt32 `protobuf:"bytes,15,opt,name=item_master_id,json=itemMasterId,proto3" json:"item_master_id,omitempty" db:"item_master_id"`
	// @inject_tag: db:"created_at"
	CreatedAt *protoext.Timestamp `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: db:"updated_at"
	UpdatedAt            *protoext.Timestamp `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

type Event struct {
	Id uint64
	// &#64;inject_tag: db:"db"
	Db string `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty" db:"db"`
	// &#64;inject_tag: db:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" db:"name"`
	// &#64;inject_tag: db:"body"
	Body []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty" db:"body"`
	// &#64;inject_tag: db:"definer"
	Definer string `protobuf:"bytes,4,opt,name=definer,proto3" json:"definer,omitempty" db:"definer"`
	// &#64;inject_tag: db:"execute_at"
	ExecuteAt *protoext.Timestamp `protobuf:"bytes,5,opt,name=execute_at,json=executeAt,proto3" json:"execute_at,omitempty" db:"execute_at"`
	// &#64;inject_tag: db:"interval_value"
	IntervalValue *protoext.NullInt32 `protobuf:"bytes,6,opt,name=interval_value,json=intervalValue,proto3" json:"interval_value,omitempty" db:"interval_value"`
	// &#64;inject_tag: db:"created"
	Created *protoext.Timestamp `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty" db:"created"`
	// &#64;inject_tag: db:"modified"
	Modified *protoext.Timestamp `protobuf:"bytes,9,opt,name=modified,proto3" json:"modified,omitempty" db:"modified"`
	// &#64;inject_tag: db:"last_executed"
	LastExecuted *protoext.Timestamp `protobuf:"bytes,10,opt,name=last_executed,json=lastExecuted,proto3" json:"last_executed,omitempty" db:"last_executed"`
	// &#64;inject_tag: db:"starts"
	Starts *protoext.Timestamp `protobuf:"bytes,11,opt,name=starts,proto3" json:"starts,omitempty" db:"starts"`
	// &#64;inject_tag: db:"ends"
	Ends *protoext.Timestamp `protobuf:"bytes,12,opt,name=ends,proto3" json:"ends,omitempty" db:"ends"`
	// &#64;inject_tag: db:"comment"
	Comment string `protobuf:"bytes,16,opt,name=comment,proto3" json:"comment,omitempty" db:"comment"`
	// &#64;inject_tag: db:"originator"
	Originator uint32 `protobuf:"varint,17,opt,name=originator,proto3" json:"originator,omitempty" db:"originator"`
	// &#64;inject_tag: db:"time_zone"
	TimeZone string `protobuf:"bytes,18,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty" db:"time_zone"`
	// &#64;inject_tag: db:"character_set_client"
	CharacterSetClient *protoext.NullString `protobuf:"bytes,19,opt,name=character_set_client,json=characterSetClient,proto3" json:"character_set_client,omitempty" db:"character_set_client"`
	// &#64;inject_tag: db:"collation_connection"
	CollationConnection *protoext.NullString `protobuf:"bytes,20,opt,name=collation_connection,json=collationConnection,proto3" json:"collation_connection,omitempty" db:"collation_connection"`
	// &#64;inject_tag: db:"db_collation"
	DbCollation *protoext.NullString `protobuf:"bytes,21,opt,name=db_collation,json=dbCollation,proto3" json:"db_collation,omitempty" db:"db_collation"`
	// &#64;inject_tag: db:"body_utf8"
	BodyUtf8             []byte   `protobuf:"bytes,22,opt,name=body_utf8,json=bodyUtf8,proto3" json:"body_utf8,omitempty" db:"body_utf8"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
