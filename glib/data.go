package glib

/*
 * GApplicationFlags
 */

type ApplicationFlags int

// UserDirectory is a representation of GLib's GUserDirectory.
type UserDirectory int

type SourceHandle uint

type SignalHandle uint

// Type is a representation of GLib's GType.
type Type uint

type Quark uint32

var (
	USER_DIRECTORY_DESKTOP      UserDirectory
	USER_DIRECTORY_DOCUMENTS    UserDirectory
	USER_DIRECTORY_DOWNLOAD     UserDirectory
	USER_DIRECTORY_MUSIC        UserDirectory
	USER_DIRECTORY_PICTURES     UserDirectory
	USER_DIRECTORY_PUBLIC_SHARE UserDirectory
	USER_DIRECTORY_TEMPLATES    UserDirectory
	USER_DIRECTORY_VIDEOS       UserDirectory
)

var USER_N_DIRECTORIES int

var (
	APPLICATION_FLAGS_NONE           ApplicationFlags
	APPLICATION_IS_SERVICE           ApplicationFlags
	APPLICATION_HANDLES_OPEN         ApplicationFlags
	APPLICATION_HANDLES_COMMAND_LINE ApplicationFlags
	APPLICATION_SEND_ENVIRONMENT     ApplicationFlags
	APPLICATION_NON_UNIQUE           ApplicationFlags
)

var (
	TYPE_INVALID   Type
	TYPE_NONE      Type
	TYPE_INTERFACE Type
	TYPE_CHAR      Type
	TYPE_UCHAR     Type
	TYPE_BOOLEAN   Type
	TYPE_INT       Type
	TYPE_UINT      Type
	TYPE_LONG      Type
	TYPE_ULONG     Type
	TYPE_INT64     Type
	TYPE_UINT64    Type
	TYPE_ENUM      Type
	TYPE_FLAGS     Type
	TYPE_FLOAT     Type
	TYPE_DOUBLE    Type
	TYPE_STRING    Type
	TYPE_POINTER   Type
	TYPE_BOXED     Type
	TYPE_PARAM     Type
	TYPE_OBJECT    Type
	TYPE_VARIANT   Type
)

type VariantClass int

var (
	VARIANT_CLASS_BOOLEAN     VariantClass
	VARIANT_CLASS_BYTE        VariantClass
	VARIANT_CLASS_INT16       VariantClass
	VARIANT_CLASS_UINT16      VariantClass
	VARIANT_CLASS_INT32       VariantClass
	VARIANT_CLASS_UINT32      VariantClass
	VARIANT_CLASS_INT64       VariantClass
	VARIANT_CLASS_UINT64      VariantClass
	VARIANT_CLASS_HANDLE      VariantClass
	VARIANT_CLASS_DOUBLE      VariantClass
	VARIANT_CLASS_STRING      VariantClass
	VARIANT_CLASS_OBJECT_PATH VariantClass
	VARIANT_CLASS_SIGNATURE   VariantClass
	VARIANT_CLASS_VARIANT     VariantClass
	VARIANT_CLASS_MAYBE       VariantClass
	VARIANT_CLASS_ARRAY       VariantClass
	VARIANT_CLASS_TUPLE       VariantClass
	VARIANT_CLASS_DICT_ENTRY  VariantClass
)