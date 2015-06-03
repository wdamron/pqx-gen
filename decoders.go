package pgxgen

var Decoders = map[string]OpMap{
	"bool": {
		"bool":          OpAssign,
		"*bool":         OpPtrAssign,
		"pgx.NullBool":  OpCustomScan,
		"*pgx.NullBool": OpCustomScan,
	},
	"int2": {
		"int":            OpAssign | OpCastInt,
		"*int":           OpPtrAssign | OpCastInt,
		"uint":           OpAssign | OpCastUint,
		"*uint":          OpPtrAssign | OpCastUint,
		"int16":          OpAssign,
		"*int16":         OpPtrAssign,
		"uint16":         OpAssign | OpCastUint16 | OpCheckOverflow,
		"*uint16":        OpPtrAssign | OpCastUint16 | OpCheckOverflow,
		"int32":          OpAssign | OpCastUint32,
		"*int32":         OpPtrAssign | OpCastUint32,
		"uint32":         OpAssign | OpCastUint32,
		"*uint32":        OpPtrAssign | OpCastUint32,
		"int64":          OpAssign | OpCastInt64,
		"*int64":         OpPtrAssign | OpCastInt64,
		"uint64":         OpAssign | OpCastUint64,
		"*uint64":        OpPtrAssign | OpCastUint64,
		"pgx.NullInt16":  OpCustomScan,
		"*pgx.NullInt16": OpCustomScan,
	},
	"int4": {
		"int":            OpAssign | OpCastInt,
		"*int":           OpPtrAssign | OpCastInt,
		"uint":           OpAssign | OpCastUint,
		"*uint":          OpPtrAssign | OpCastUint,
		"int16":          OpAssign | OpCastInt16 | OpCheckOverflow,
		"*int16":         OpPtrAssign | OpCastInt16 | OpCheckOverflow,
		"uint16":         OpAssign | OpCastUint16 | OpCheckOverflow,
		"*uint16":        OpPtrAssign | OpCastUint16 | OpCheckOverflow,
		"int32":          OpAssign,
		"*int32":         OpPtrAssign,
		"uint32":         OpAssign | OpCastUint32 | OpCheckOverflow,
		"*uint32":        OpPtrAssign | OpCastUint32 | OpCheckOverflow,
		"int64":          OpAssign | OpCastInt64,
		"*int64":         OpPtrAssign | OpCastInt64,
		"uint64":         OpAssign | OpCastUint64,
		"*uint64":        OpPtrAssign | OpCastUint64,
		"pgx.NullInt32":  OpCustomScan,
		"*pgx.NullInt32": OpCustomScan,
	},
	"int8": {
		"int":            OpAssign | OpCastInt,
		"*int":           OpPtrAssign | OpCastInt,
		"uint":           OpAssign | OpCastUint,
		"*uint":          OpPtrAssign | OpCastUint,
		"int16":          OpAssign | OpCastInt16 | OpCheckOverflow,
		"*int16":         OpPtrAssign | OpCastInt16 | OpCheckOverflow,
		"uint16":         OpAssign | OpCastUint16 | OpCheckOverflow,
		"*uint16":        OpPtrAssign | OpCastUint16 | OpCheckOverflow,
		"int32":          OpAssign | OpCastInt32 | OpCheckOverflow,
		"*int32":         OpPtrAssign | OpCastInt32 | OpCheckOverflow,
		"uint32":         OpAssign | OpCastUint32 | OpCheckOverflow,
		"*uint32":        OpPtrAssign | OpCastUint32 | OpCheckOverflow,
		"int64":          OpAssign,
		"*int64":         OpPtrAssign,
		"uint64":         OpAssign | OpCastUint64 | OpCheckOverflow,
		"*uint64":        OpPtrAssign | OpCastUint64 | OpCheckOverflow,
		"pgx.NullInt64":  OpCustomScan,
		"*pgx.NullInt64": OpCustomScan,
	},
	"real": {
		"float32":          OpAssign,
		"*float32":         OpPtrAssign,
		"float64":          OpAssign | OpCastFloat64,
		"*float64":         OpPtrAssign | OpCastFloat64,
		"pgx.NullFloat32":  OpCustomScan,
		"*pgx.NullFloat32": OpCustomScan,
	},
	"float": {
		"float32":          OpAssign | OpCastFloat32,
		"*float32":         OpPtrAssign | OpCastFloat32,
		"float64":          OpAssign,
		"*float64":         OpPtrAssign,
		"pgx.NullFloat64":  OpCustomScan,
		"*pgx.NullFloat64": OpCustomScan,
	},
	"bytea": {
		"string":  OpAssign | OpCastString,
		"*string": OpAssign | OpCastString,
		"[]byte":  OpAssign,
		"*[]byte": OpPtrAssign,
	},
	"text": {
		"string":          OpAssign,
		"*string":         OpPtrAssign,
		"[]byte":          OpAssign | OpCastBytes,
		"*[]byte":         OpPtrAssign | OpCastBytes,
		"pgx.NullString":  OpCustomScan,
		"*pgx.NullString": OpCustomScan,
	},
	"varchar": {
		"string":          OpAssign,
		"*string":         OpPtrAssign,
		"[]byte":          OpAssign | OpCastBytes,
		"*[]byte":         OpPtrAssign | OpCastBytes,
		"pgx.NullString":  OpCustomScan,
		"*pgx.NullString": OpCustomScan,
	},
	"date": {
		"time.Time":  OpAssign,
		"*time.Time": OpPtrAssign,
	},
	"timestamp": {
		"time.Time":     OpAssign,
		"*time.Time":    OpPtrAssign,
		"pgx.NullTime":  OpCustomScan,
		"*pgx.NullTime": OpCustomScan,
	},
	"timestampTz": {
		"time.Time":     OpAssign,
		"*time.Time":    OpPtrAssign,
		"pgx.NullTime":  OpCustomScan,
		"*pgx.NullTime": OpCustomScan,
	},
	"bool[]": {
		"[]bool":  OpAssign,
		"*[]bool": OpPtrAssign,
	},
	"int2[]": {
		"[]int16":  OpAssign,
		"*[]int16": OpPtrAssign,
	},
	"int4[]": {
		"[]int32":  OpAssign,
		"*[]int32": OpPtrAssign,
	},
	"int8[]": {
		"[]int64":  OpAssign,
		"*[]int64": OpPtrAssign,
	},
	"real[]": {
		"[]float32":  OpAssign,
		"*[]float32": OpPtrAssign,
	},
	"float[]": {
		"[]float64":  OpAssign,
		"*[]float64": OpPtrAssign,
	},
	"text[]": {
		"[]string":  OpAssign,
		"*[]string": OpPtrAssign,
	},
	"varchar[]": {
		"[]string":  OpAssign,
		"*[]string": OpPtrAssign,
	},
	"timestamp[]": {
		"[]time.Time":  OpAssign,
		"*[]time.Time": OpPtrAssign,
	},
	"timestampTz[]": {
		"[]time.Time":  OpAssign,
		"*[]time.Time": OpPtrAssign,
	},
	"hstore": {
		"pgx.Hstore":         OpAssign,
		"*pgx.Hstore":        OpPtrAssign,
		"pgx.NullHstore":     OpCustomScan,
		"*pgx.NullHstore":    OpCustomScan,
		"map[string]string":  OpHstoreMapDecode,
		"*map[string]string": OpPtrAssign | OpHstoreMapDecode,
	},
	"uuid": {
		"string":     OpUuidDecode | OpUuidStringDecode,
		"*string":    OpPtrAssign | OpUuidDecode | OpUuidStringDecode,
		"uuid.UUID":  OpUuidDecode,
		"*uuid.UUID": OpPtrAssign | OpUuidDecode,
	},
}