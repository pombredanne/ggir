package gio

/*
#include <gio/gio.h>
#include <gio/gunixoutputstream.h>
#include <gio/gunixinputstream.h>
#include <gio/gunixcredentialsmessage.h>
#include <gio/gunixmounts.h>
#include <gio/gdesktopappinfo.h>
#include <gio/gunixfdlist.h>
#include <gio/gunixsocketaddress.h>
#include <gio/gunixfdmessage.h>
#include <gio/gnetworking.h>
#include <gio/gunixconnection.h>
#include <stdlib.h>
*/
import "C"

var (
	// BusType
	BUS_TYPE_STARTER = C.GBusType(C.G_BUS_TYPE_STARTER)
	BUS_TYPE_NONE    = C.GBusType(C.G_BUS_TYPE_NONE)
	BUS_TYPE_SYSTEM  = C.GBusType(C.G_BUS_TYPE_SYSTEM)
	BUS_TYPE_SESSION = C.GBusType(C.G_BUS_TYPE_SESSION)

	// ConverterResult
	CONVERTER_ERROR     = C.GConverterResult(C.G_CONVERTER_ERROR)
	CONVERTER_CONVERTED = C.GConverterResult(C.G_CONVERTER_CONVERTED)
	CONVERTER_FINISHED  = C.GConverterResult(C.G_CONVERTER_FINISHED)
	CONVERTER_FLUSHED   = C.GConverterResult(C.G_CONVERTER_FLUSHED)

	// CredentialsType
	CREDENTIALS_TYPE_INVALID              = C.GCredentialsType(C.G_CREDENTIALS_TYPE_INVALID)
	CREDENTIALS_TYPE_LINUX_UCRED          = C.GCredentialsType(C.G_CREDENTIALS_TYPE_LINUX_UCRED)
	CREDENTIALS_TYPE_FREEBSD_CMSGCRED     = C.GCredentialsType(C.G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED)
	CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED = C.GCredentialsType(C.G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED)
	CREDENTIALS_TYPE_SOLARIS_UCRED        = C.GCredentialsType(C.G_CREDENTIALS_TYPE_SOLARIS_UCRED)

	// DBusError
	DBUS_ERROR_FAILED                           = C.GDBusError(C.G_DBUS_ERROR_FAILED)
	DBUS_ERROR_NO_MEMORY                        = C.GDBusError(C.G_DBUS_ERROR_NO_MEMORY)
	DBUS_ERROR_SERVICE_UNKNOWN                  = C.GDBusError(C.G_DBUS_ERROR_SERVICE_UNKNOWN)
	DBUS_ERROR_NAME_HAS_NO_OWNER                = C.GDBusError(C.G_DBUS_ERROR_NAME_HAS_NO_OWNER)
	DBUS_ERROR_NO_REPLY                         = C.GDBusError(C.G_DBUS_ERROR_NO_REPLY)
	DBUS_ERROR_IO_ERROR                         = C.GDBusError(C.G_DBUS_ERROR_IO_ERROR)
	DBUS_ERROR_BAD_ADDRESS                      = C.GDBusError(C.G_DBUS_ERROR_BAD_ADDRESS)
	DBUS_ERROR_NOT_SUPPORTED                    = C.GDBusError(C.G_DBUS_ERROR_NOT_SUPPORTED)
	DBUS_ERROR_LIMITS_EXCEEDED                  = C.GDBusError(C.G_DBUS_ERROR_LIMITS_EXCEEDED)
	DBUS_ERROR_ACCESS_DENIED                    = C.GDBusError(C.G_DBUS_ERROR_ACCESS_DENIED)
	DBUS_ERROR_AUTH_FAILED                      = C.GDBusError(C.G_DBUS_ERROR_AUTH_FAILED)
	DBUS_ERROR_NO_SERVER                        = C.GDBusError(C.G_DBUS_ERROR_NO_SERVER)
	DBUS_ERROR_TIMEOUT                          = C.GDBusError(C.G_DBUS_ERROR_TIMEOUT)
	DBUS_ERROR_NO_NETWORK                       = C.GDBusError(C.G_DBUS_ERROR_NO_NETWORK)
	DBUS_ERROR_ADDRESS_IN_USE                   = C.GDBusError(C.G_DBUS_ERROR_ADDRESS_IN_USE)
	DBUS_ERROR_DISCONNECTED                     = C.GDBusError(C.G_DBUS_ERROR_DISCONNECTED)
	DBUS_ERROR_INVALID_ARGS                     = C.GDBusError(C.G_DBUS_ERROR_INVALID_ARGS)
	DBUS_ERROR_FILE_NOT_FOUND                   = C.GDBusError(C.G_DBUS_ERROR_FILE_NOT_FOUND)
	DBUS_ERROR_FILE_EXISTS                      = C.GDBusError(C.G_DBUS_ERROR_FILE_EXISTS)
	DBUS_ERROR_UNKNOWN_METHOD                   = C.GDBusError(C.G_DBUS_ERROR_UNKNOWN_METHOD)
	DBUS_ERROR_TIMED_OUT                        = C.GDBusError(C.G_DBUS_ERROR_TIMED_OUT)
	DBUS_ERROR_MATCH_RULE_NOT_FOUND             = C.GDBusError(C.G_DBUS_ERROR_MATCH_RULE_NOT_FOUND)
	DBUS_ERROR_MATCH_RULE_INVALID               = C.GDBusError(C.G_DBUS_ERROR_MATCH_RULE_INVALID)
	DBUS_ERROR_SPAWN_EXEC_FAILED                = C.GDBusError(C.G_DBUS_ERROR_SPAWN_EXEC_FAILED)
	DBUS_ERROR_SPAWN_FORK_FAILED                = C.GDBusError(C.G_DBUS_ERROR_SPAWN_FORK_FAILED)
	DBUS_ERROR_SPAWN_CHILD_EXITED               = C.GDBusError(C.G_DBUS_ERROR_SPAWN_CHILD_EXITED)
	DBUS_ERROR_SPAWN_CHILD_SIGNALED             = C.GDBusError(C.G_DBUS_ERROR_SPAWN_CHILD_SIGNALED)
	DBUS_ERROR_SPAWN_FAILED                     = C.GDBusError(C.G_DBUS_ERROR_SPAWN_FAILED)
	DBUS_ERROR_SPAWN_SETUP_FAILED               = C.GDBusError(C.G_DBUS_ERROR_SPAWN_SETUP_FAILED)
	DBUS_ERROR_SPAWN_CONFIG_INVALID             = C.GDBusError(C.G_DBUS_ERROR_SPAWN_CONFIG_INVALID)
	DBUS_ERROR_SPAWN_SERVICE_INVALID            = C.GDBusError(C.G_DBUS_ERROR_SPAWN_SERVICE_INVALID)
	DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND          = C.GDBusError(C.G_DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND)
	DBUS_ERROR_SPAWN_PERMISSIONS_INVALID        = C.GDBusError(C.G_DBUS_ERROR_SPAWN_PERMISSIONS_INVALID)
	DBUS_ERROR_SPAWN_FILE_INVALID               = C.GDBusError(C.G_DBUS_ERROR_SPAWN_FILE_INVALID)
	DBUS_ERROR_SPAWN_NO_MEMORY                  = C.GDBusError(C.G_DBUS_ERROR_SPAWN_NO_MEMORY)
	DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN          = C.GDBusError(C.G_DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN)
	DBUS_ERROR_INVALID_SIGNATURE                = C.GDBusError(C.G_DBUS_ERROR_INVALID_SIGNATURE)
	DBUS_ERROR_INVALID_FILE_CONTENT             = C.GDBusError(C.G_DBUS_ERROR_INVALID_FILE_CONTENT)
	DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN = C.GDBusError(C.G_DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN)
	DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN           = C.GDBusError(C.G_DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN)
	DBUS_ERROR_OBJECT_PATH_IN_USE               = C.GDBusError(C.G_DBUS_ERROR_OBJECT_PATH_IN_USE)
	DBUS_ERROR_UNKNOWN_OBJECT                   = C.GDBusError(C.G_DBUS_ERROR_UNKNOWN_OBJECT)
	DBUS_ERROR_UNKNOWN_INTERFACE                = C.GDBusError(C.G_DBUS_ERROR_UNKNOWN_INTERFACE)
	DBUS_ERROR_UNKNOWN_PROPERTY                 = C.GDBusError(C.G_DBUS_ERROR_UNKNOWN_PROPERTY)
	DBUS_ERROR_PROPERTY_READ_ONLY               = C.GDBusError(C.G_DBUS_ERROR_PROPERTY_READ_ONLY)

	// DBusMessageByteOrder
	DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN    = C.GDBusMessageByteOrder(C.G_DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN)
	DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN = C.GDBusMessageByteOrder(C.G_DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN)

	// DBusMessageHeaderField
	DBUS_MESSAGE_HEADER_FIELD_INVALID      = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_INVALID)
	DBUS_MESSAGE_HEADER_FIELD_PATH         = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_PATH)
	DBUS_MESSAGE_HEADER_FIELD_INTERFACE    = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE)
	DBUS_MESSAGE_HEADER_FIELD_MEMBER       = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_MEMBER)
	DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME   = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME)
	DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL)
	DBUS_MESSAGE_HEADER_FIELD_DESTINATION  = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION)
	DBUS_MESSAGE_HEADER_FIELD_SENDER       = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_SENDER)
	DBUS_MESSAGE_HEADER_FIELD_SIGNATURE    = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE)
	DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS = C.GDBusMessageHeaderField(C.G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS)

	// DBusMessageType
	DBUS_MESSAGE_TYPE_INVALID       = C.GDBusMessageType(C.G_DBUS_MESSAGE_TYPE_INVALID)
	DBUS_MESSAGE_TYPE_METHOD_CALL   = C.GDBusMessageType(C.G_DBUS_MESSAGE_TYPE_METHOD_CALL)
	DBUS_MESSAGE_TYPE_METHOD_RETURN = C.GDBusMessageType(C.G_DBUS_MESSAGE_TYPE_METHOD_RETURN)
	DBUS_MESSAGE_TYPE_ERROR         = C.GDBusMessageType(C.G_DBUS_MESSAGE_TYPE_ERROR)
	DBUS_MESSAGE_TYPE_SIGNAL        = C.GDBusMessageType(C.G_DBUS_MESSAGE_TYPE_SIGNAL)

	// DataStreamByteOrder
	DATA_STREAM_BYTE_ORDER_BIG_ENDIAN    = C.GDataStreamByteOrder(C.G_DATA_STREAM_BYTE_ORDER_BIG_ENDIAN)
	DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN = C.GDataStreamByteOrder(C.G_DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN)
	DATA_STREAM_BYTE_ORDER_HOST_ENDIAN   = C.GDataStreamByteOrder(C.G_DATA_STREAM_BYTE_ORDER_HOST_ENDIAN)

	// DataStreamNewlineType
	DATA_STREAM_NEWLINE_TYPE_LF    = C.GDataStreamNewlineType(C.G_DATA_STREAM_NEWLINE_TYPE_LF)
	DATA_STREAM_NEWLINE_TYPE_CR    = C.GDataStreamNewlineType(C.G_DATA_STREAM_NEWLINE_TYPE_CR)
	DATA_STREAM_NEWLINE_TYPE_CR_LF = C.GDataStreamNewlineType(C.G_DATA_STREAM_NEWLINE_TYPE_CR_LF)
	DATA_STREAM_NEWLINE_TYPE_ANY   = C.GDataStreamNewlineType(C.G_DATA_STREAM_NEWLINE_TYPE_ANY)

	// DriveStartStopType
	DRIVE_START_STOP_TYPE_UNKNOWN   = C.GDriveStartStopType(C.G_DRIVE_START_STOP_TYPE_UNKNOWN)
	DRIVE_START_STOP_TYPE_SHUTDOWN  = C.GDriveStartStopType(C.G_DRIVE_START_STOP_TYPE_SHUTDOWN)
	DRIVE_START_STOP_TYPE_NETWORK   = C.GDriveStartStopType(C.G_DRIVE_START_STOP_TYPE_NETWORK)
	DRIVE_START_STOP_TYPE_MULTIDISK = C.GDriveStartStopType(C.G_DRIVE_START_STOP_TYPE_MULTIDISK)
	DRIVE_START_STOP_TYPE_PASSWORD  = C.GDriveStartStopType(C.G_DRIVE_START_STOP_TYPE_PASSWORD)

	// EmblemOrigin
	EMBLEM_ORIGIN_UNKNOWN      = C.GEmblemOrigin(C.G_EMBLEM_ORIGIN_UNKNOWN)
	EMBLEM_ORIGIN_DEVICE       = C.GEmblemOrigin(C.G_EMBLEM_ORIGIN_DEVICE)
	EMBLEM_ORIGIN_LIVEMETADATA = C.GEmblemOrigin(C.G_EMBLEM_ORIGIN_LIVEMETADATA)
	EMBLEM_ORIGIN_TAG          = C.GEmblemOrigin(C.G_EMBLEM_ORIGIN_TAG)

	// FileAttributeStatus
	FILE_ATTRIBUTE_STATUS_UNSET         = C.GFileAttributeStatus(C.G_FILE_ATTRIBUTE_STATUS_UNSET)
	FILE_ATTRIBUTE_STATUS_SET           = C.GFileAttributeStatus(C.G_FILE_ATTRIBUTE_STATUS_SET)
	FILE_ATTRIBUTE_STATUS_ERROR_SETTING = C.GFileAttributeStatus(C.G_FILE_ATTRIBUTE_STATUS_ERROR_SETTING)

	// FileAttributeType
	FILE_ATTRIBUTE_TYPE_INVALID     = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_INVALID)
	FILE_ATTRIBUTE_TYPE_STRING      = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_STRING)
	FILE_ATTRIBUTE_TYPE_BYTE_STRING = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_BYTE_STRING)
	FILE_ATTRIBUTE_TYPE_BOOLEAN     = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_BOOLEAN)
	FILE_ATTRIBUTE_TYPE_UINT32      = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_UINT32)
	FILE_ATTRIBUTE_TYPE_INT32       = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_INT32)
	FILE_ATTRIBUTE_TYPE_UINT64      = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_UINT64)
	FILE_ATTRIBUTE_TYPE_INT64       = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_INT64)
	FILE_ATTRIBUTE_TYPE_OBJECT      = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_OBJECT)
	FILE_ATTRIBUTE_TYPE_STRINGV     = C.GFileAttributeType(C.G_FILE_ATTRIBUTE_TYPE_STRINGV)

	// FileMonitorEvent
	FILE_MONITOR_EVENT_CHANGED           = C.GFileMonitorEvent(C.G_FILE_MONITOR_EVENT_CHANGED)
	FILE_MONITOR_EVENT_CHANGES_DONE_HINT = C.GFileMonitorEvent(C.G_FILE_MONITOR_EVENT_CHANGES_DONE_HINT)
	FILE_MONITOR_EVENT_DELETED           = C.GFileMonitorEvent(C.G_FILE_MONITOR_EVENT_DELETED)
	FILE_MONITOR_EVENT_CREATED           = C.GFileMonitorEvent(C.G_FILE_MONITOR_EVENT_CREATED)
	FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED = C.GFileMonitorEvent(C.G_FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED)
	FILE_MONITOR_EVENT_PRE_UNMOUNT       = C.GFileMonitorEvent(C.G_FILE_MONITOR_EVENT_PRE_UNMOUNT)
	FILE_MONITOR_EVENT_UNMOUNTED         = C.GFileMonitorEvent(C.G_FILE_MONITOR_EVENT_UNMOUNTED)
	FILE_MONITOR_EVENT_MOVED             = C.GFileMonitorEvent(C.G_FILE_MONITOR_EVENT_MOVED)

	// FileType
	FILE_TYPE_UNKNOWN       = C.GFileType(C.G_FILE_TYPE_UNKNOWN)
	FILE_TYPE_REGULAR       = C.GFileType(C.G_FILE_TYPE_REGULAR)
	FILE_TYPE_DIRECTORY     = C.GFileType(C.G_FILE_TYPE_DIRECTORY)
	FILE_TYPE_SYMBOLIC_LINK = C.GFileType(C.G_FILE_TYPE_SYMBOLIC_LINK)
	FILE_TYPE_SPECIAL       = C.GFileType(C.G_FILE_TYPE_SPECIAL)
	FILE_TYPE_SHORTCUT      = C.GFileType(C.G_FILE_TYPE_SHORTCUT)
	FILE_TYPE_MOUNTABLE     = C.GFileType(C.G_FILE_TYPE_MOUNTABLE)

	// FilesystemPreviewType
	FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS = C.GFilesystemPreviewType(C.G_FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS)
	FILESYSTEM_PREVIEW_TYPE_IF_LOCAL  = C.GFilesystemPreviewType(C.G_FILESYSTEM_PREVIEW_TYPE_IF_LOCAL)
	FILESYSTEM_PREVIEW_TYPE_NEVER     = C.GFilesystemPreviewType(C.G_FILESYSTEM_PREVIEW_TYPE_NEVER)

	// IOErrorEnum
	IO_ERROR_FAILED              = C.GIOErrorEnum(C.G_IO_ERROR_FAILED)
	IO_ERROR_NOT_FOUND           = C.GIOErrorEnum(C.G_IO_ERROR_NOT_FOUND)
	IO_ERROR_EXISTS              = C.GIOErrorEnum(C.G_IO_ERROR_EXISTS)
	IO_ERROR_IS_DIRECTORY        = C.GIOErrorEnum(C.G_IO_ERROR_IS_DIRECTORY)
	IO_ERROR_NOT_DIRECTORY       = C.GIOErrorEnum(C.G_IO_ERROR_NOT_DIRECTORY)
	IO_ERROR_NOT_EMPTY           = C.GIOErrorEnum(C.G_IO_ERROR_NOT_EMPTY)
	IO_ERROR_NOT_REGULAR_FILE    = C.GIOErrorEnum(C.G_IO_ERROR_NOT_REGULAR_FILE)
	IO_ERROR_NOT_SYMBOLIC_LINK   = C.GIOErrorEnum(C.G_IO_ERROR_NOT_SYMBOLIC_LINK)
	IO_ERROR_NOT_MOUNTABLE_FILE  = C.GIOErrorEnum(C.G_IO_ERROR_NOT_MOUNTABLE_FILE)
	IO_ERROR_FILENAME_TOO_LONG   = C.GIOErrorEnum(C.G_IO_ERROR_FILENAME_TOO_LONG)
	IO_ERROR_INVALID_FILENAME    = C.GIOErrorEnum(C.G_IO_ERROR_INVALID_FILENAME)
	IO_ERROR_TOO_MANY_LINKS      = C.GIOErrorEnum(C.G_IO_ERROR_TOO_MANY_LINKS)
	IO_ERROR_NO_SPACE            = C.GIOErrorEnum(C.G_IO_ERROR_NO_SPACE)
	IO_ERROR_INVALID_ARGUMENT    = C.GIOErrorEnum(C.G_IO_ERROR_INVALID_ARGUMENT)
	IO_ERROR_PERMISSION_DENIED   = C.GIOErrorEnum(C.G_IO_ERROR_PERMISSION_DENIED)
	IO_ERROR_NOT_SUPPORTED       = C.GIOErrorEnum(C.G_IO_ERROR_NOT_SUPPORTED)
	IO_ERROR_NOT_MOUNTED         = C.GIOErrorEnum(C.G_IO_ERROR_NOT_MOUNTED)
	IO_ERROR_ALREADY_MOUNTED     = C.GIOErrorEnum(C.G_IO_ERROR_ALREADY_MOUNTED)
	IO_ERROR_CLOSED              = C.GIOErrorEnum(C.G_IO_ERROR_CLOSED)
	IO_ERROR_CANCELLED           = C.GIOErrorEnum(C.G_IO_ERROR_CANCELLED)
	IO_ERROR_PENDING             = C.GIOErrorEnum(C.G_IO_ERROR_PENDING)
	IO_ERROR_READ_ONLY           = C.GIOErrorEnum(C.G_IO_ERROR_READ_ONLY)
	IO_ERROR_CANT_CREATE_BACKUP  = C.GIOErrorEnum(C.G_IO_ERROR_CANT_CREATE_BACKUP)
	IO_ERROR_WRONG_ETAG          = C.GIOErrorEnum(C.G_IO_ERROR_WRONG_ETAG)
	IO_ERROR_TIMED_OUT           = C.GIOErrorEnum(C.G_IO_ERROR_TIMED_OUT)
	IO_ERROR_WOULD_RECURSE       = C.GIOErrorEnum(C.G_IO_ERROR_WOULD_RECURSE)
	IO_ERROR_BUSY                = C.GIOErrorEnum(C.G_IO_ERROR_BUSY)
	IO_ERROR_WOULD_BLOCK         = C.GIOErrorEnum(C.G_IO_ERROR_WOULD_BLOCK)
	IO_ERROR_HOST_NOT_FOUND      = C.GIOErrorEnum(C.G_IO_ERROR_HOST_NOT_FOUND)
	IO_ERROR_WOULD_MERGE         = C.GIOErrorEnum(C.G_IO_ERROR_WOULD_MERGE)
	IO_ERROR_FAILED_HANDLED      = C.GIOErrorEnum(C.G_IO_ERROR_FAILED_HANDLED)
	IO_ERROR_TOO_MANY_OPEN_FILES = C.GIOErrorEnum(C.G_IO_ERROR_TOO_MANY_OPEN_FILES)
	IO_ERROR_NOT_INITIALIZED     = C.GIOErrorEnum(C.G_IO_ERROR_NOT_INITIALIZED)
	IO_ERROR_ADDRESS_IN_USE      = C.GIOErrorEnum(C.G_IO_ERROR_ADDRESS_IN_USE)
	IO_ERROR_PARTIAL_INPUT       = C.GIOErrorEnum(C.G_IO_ERROR_PARTIAL_INPUT)
	IO_ERROR_INVALID_DATA        = C.GIOErrorEnum(C.G_IO_ERROR_INVALID_DATA)
	IO_ERROR_DBUS_ERROR          = C.GIOErrorEnum(C.G_IO_ERROR_DBUS_ERROR)
	IO_ERROR_HOST_UNREACHABLE    = C.GIOErrorEnum(C.G_IO_ERROR_HOST_UNREACHABLE)
	IO_ERROR_NETWORK_UNREACHABLE = C.GIOErrorEnum(C.G_IO_ERROR_NETWORK_UNREACHABLE)
	IO_ERROR_CONNECTION_REFUSED  = C.GIOErrorEnum(C.G_IO_ERROR_CONNECTION_REFUSED)
	IO_ERROR_PROXY_FAILED        = C.GIOErrorEnum(C.G_IO_ERROR_PROXY_FAILED)
	IO_ERROR_PROXY_AUTH_FAILED   = C.GIOErrorEnum(C.G_IO_ERROR_PROXY_AUTH_FAILED)
	IO_ERROR_PROXY_NEED_AUTH     = C.GIOErrorEnum(C.G_IO_ERROR_PROXY_NEED_AUTH)
	IO_ERROR_PROXY_NOT_ALLOWED   = C.GIOErrorEnum(C.G_IO_ERROR_PROXY_NOT_ALLOWED)
	IO_ERROR_BROKEN_PIPE         = C.GIOErrorEnum(C.G_IO_ERROR_BROKEN_PIPE)

	// IOModuleScopeFlags
	IO_MODULE_SCOPE_NONE             = C.GIOModuleScopeFlags(C.G_IO_MODULE_SCOPE_NONE)
	IO_MODULE_SCOPE_BLOCK_DUPLICATES = C.GIOModuleScopeFlags(C.G_IO_MODULE_SCOPE_BLOCK_DUPLICATES)

	// MountOperationResult
	MOUNT_OPERATION_HANDLED   = C.GMountOperationResult(C.G_MOUNT_OPERATION_HANDLED)
	MOUNT_OPERATION_ABORTED   = C.GMountOperationResult(C.G_MOUNT_OPERATION_ABORTED)
	MOUNT_OPERATION_UNHANDLED = C.GMountOperationResult(C.G_MOUNT_OPERATION_UNHANDLED)

	// PasswordSave
	PASSWORD_SAVE_NEVER       = C.GPasswordSave(C.G_PASSWORD_SAVE_NEVER)
	PASSWORD_SAVE_FOR_SESSION = C.GPasswordSave(C.G_PASSWORD_SAVE_FOR_SESSION)
	PASSWORD_SAVE_PERMANENTLY = C.GPasswordSave(C.G_PASSWORD_SAVE_PERMANENTLY)

	// ResolverError
	RESOLVER_ERROR_NOT_FOUND         = C.GResolverError(C.G_RESOLVER_ERROR_NOT_FOUND)
	RESOLVER_ERROR_TEMPORARY_FAILURE = C.GResolverError(C.G_RESOLVER_ERROR_TEMPORARY_FAILURE)
	RESOLVER_ERROR_INTERNAL          = C.GResolverError(C.G_RESOLVER_ERROR_INTERNAL)

	// ResolverRecordType
	RESOLVER_RECORD_SRV = C.GResolverRecordType(C.G_RESOLVER_RECORD_SRV)
	RESOLVER_RECORD_MX  = C.GResolverRecordType(C.G_RESOLVER_RECORD_MX)
	RESOLVER_RECORD_TXT = C.GResolverRecordType(C.G_RESOLVER_RECORD_TXT)
	RESOLVER_RECORD_SOA = C.GResolverRecordType(C.G_RESOLVER_RECORD_SOA)
	RESOLVER_RECORD_NS  = C.GResolverRecordType(C.G_RESOLVER_RECORD_NS)

	// ResourceError
	RESOURCE_ERROR_NOT_FOUND = C.GResourceError(C.G_RESOURCE_ERROR_NOT_FOUND)
	RESOURCE_ERROR_INTERNAL  = C.GResourceError(C.G_RESOURCE_ERROR_INTERNAL)

	// SocketClientEvent
	SOCKET_CLIENT_RESOLVING         = C.GSocketClientEvent(C.G_SOCKET_CLIENT_RESOLVING)
	SOCKET_CLIENT_RESOLVED          = C.GSocketClientEvent(C.G_SOCKET_CLIENT_RESOLVED)
	SOCKET_CLIENT_CONNECTING        = C.GSocketClientEvent(C.G_SOCKET_CLIENT_CONNECTING)
	SOCKET_CLIENT_CONNECTED         = C.GSocketClientEvent(C.G_SOCKET_CLIENT_CONNECTED)
	SOCKET_CLIENT_PROXY_NEGOTIATING = C.GSocketClientEvent(C.G_SOCKET_CLIENT_PROXY_NEGOTIATING)
	SOCKET_CLIENT_PROXY_NEGOTIATED  = C.GSocketClientEvent(C.G_SOCKET_CLIENT_PROXY_NEGOTIATED)
	SOCKET_CLIENT_TLS_HANDSHAKING   = C.GSocketClientEvent(C.G_SOCKET_CLIENT_TLS_HANDSHAKING)
	SOCKET_CLIENT_TLS_HANDSHAKED    = C.GSocketClientEvent(C.G_SOCKET_CLIENT_TLS_HANDSHAKED)
	SOCKET_CLIENT_COMPLETE          = C.GSocketClientEvent(C.G_SOCKET_CLIENT_COMPLETE)

	// SocketFamily
	SOCKET_FAMILY_INVALID = C.GSocketFamily(C.G_SOCKET_FAMILY_INVALID)
	SOCKET_FAMILY_UNIX    = C.GSocketFamily(C.G_SOCKET_FAMILY_UNIX)
	SOCKET_FAMILY_IPV4    = C.GSocketFamily(C.G_SOCKET_FAMILY_IPV4)
	SOCKET_FAMILY_IPV6    = C.GSocketFamily(C.G_SOCKET_FAMILY_IPV6)

	// SocketProtocol
	SOCKET_PROTOCOL_UNKNOWN = C.GSocketProtocol(C.G_SOCKET_PROTOCOL_UNKNOWN)
	SOCKET_PROTOCOL_DEFAULT = C.GSocketProtocol(C.G_SOCKET_PROTOCOL_DEFAULT)
	SOCKET_PROTOCOL_TCP     = C.GSocketProtocol(C.G_SOCKET_PROTOCOL_TCP)
	SOCKET_PROTOCOL_UDP     = C.GSocketProtocol(C.G_SOCKET_PROTOCOL_UDP)
	SOCKET_PROTOCOL_SCTP    = C.GSocketProtocol(C.G_SOCKET_PROTOCOL_SCTP)

	// SocketType
	SOCKET_TYPE_INVALID   = C.GSocketType(C.G_SOCKET_TYPE_INVALID)
	SOCKET_TYPE_STREAM    = C.GSocketType(C.G_SOCKET_TYPE_STREAM)
	SOCKET_TYPE_DATAGRAM  = C.GSocketType(C.G_SOCKET_TYPE_DATAGRAM)
	SOCKET_TYPE_SEQPACKET = C.GSocketType(C.G_SOCKET_TYPE_SEQPACKET)

	// TlsAuthenticationMode
	TLS_AUTHENTICATION_NONE      = C.GTlsAuthenticationMode(C.G_TLS_AUTHENTICATION_NONE)
	TLS_AUTHENTICATION_REQUESTED = C.GTlsAuthenticationMode(C.G_TLS_AUTHENTICATION_REQUESTED)
	TLS_AUTHENTICATION_REQUIRED  = C.GTlsAuthenticationMode(C.G_TLS_AUTHENTICATION_REQUIRED)

	// TlsCertificateRequestFlags
	TLS_CERTIFICATE_REQUEST_NONE = C.GTlsCertificateRequestFlags(C.G_TLS_CERTIFICATE_REQUEST_NONE)

	// TlsDatabaseLookupFlags
	TLS_DATABASE_LOOKUP_NONE    = C.GTlsDatabaseLookupFlags(C.G_TLS_DATABASE_LOOKUP_NONE)
	TLS_DATABASE_LOOKUP_KEYPAIR = C.GTlsDatabaseLookupFlags(C.G_TLS_DATABASE_LOOKUP_KEYPAIR)

	// TlsError
	TLS_ERROR_UNAVAILABLE          = C.GTlsError(C.G_TLS_ERROR_UNAVAILABLE)
	TLS_ERROR_MISC                 = C.GTlsError(C.G_TLS_ERROR_MISC)
	TLS_ERROR_BAD_CERTIFICATE      = C.GTlsError(C.G_TLS_ERROR_BAD_CERTIFICATE)
	TLS_ERROR_NOT_TLS              = C.GTlsError(C.G_TLS_ERROR_NOT_TLS)
	TLS_ERROR_HANDSHAKE            = C.GTlsError(C.G_TLS_ERROR_HANDSHAKE)
	TLS_ERROR_CERTIFICATE_REQUIRED = C.GTlsError(C.G_TLS_ERROR_CERTIFICATE_REQUIRED)
	TLS_ERROR_EOF                  = C.GTlsError(C.G_TLS_ERROR_EOF)

	// TlsInteractionResult
	TLS_INTERACTION_UNHANDLED = C.GTlsInteractionResult(C.G_TLS_INTERACTION_UNHANDLED)
	TLS_INTERACTION_HANDLED   = C.GTlsInteractionResult(C.G_TLS_INTERACTION_HANDLED)
	TLS_INTERACTION_FAILED    = C.GTlsInteractionResult(C.G_TLS_INTERACTION_FAILED)

	// TlsRehandshakeMode
	TLS_REHANDSHAKE_NEVER    = C.GTlsRehandshakeMode(C.G_TLS_REHANDSHAKE_NEVER)
	TLS_REHANDSHAKE_SAFELY   = C.GTlsRehandshakeMode(C.G_TLS_REHANDSHAKE_SAFELY)
	TLS_REHANDSHAKE_UNSAFELY = C.GTlsRehandshakeMode(C.G_TLS_REHANDSHAKE_UNSAFELY)

	// UnixSocketAddressType
	UNIX_SOCKET_ADDRESS_INVALID         = C.GUnixSocketAddressType(C.G_UNIX_SOCKET_ADDRESS_INVALID)
	UNIX_SOCKET_ADDRESS_ANONYMOUS       = C.GUnixSocketAddressType(C.G_UNIX_SOCKET_ADDRESS_ANONYMOUS)
	UNIX_SOCKET_ADDRESS_PATH            = C.GUnixSocketAddressType(C.G_UNIX_SOCKET_ADDRESS_PATH)
	UNIX_SOCKET_ADDRESS_ABSTRACT        = C.GUnixSocketAddressType(C.G_UNIX_SOCKET_ADDRESS_ABSTRACT)
	UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED = C.GUnixSocketAddressType(C.G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED)

	// ZlibCompressorFormat
	ZLIB_COMPRESSOR_FORMAT_ZLIB = C.GZlibCompressorFormat(C.G_ZLIB_COMPRESSOR_FORMAT_ZLIB)
	ZLIB_COMPRESSOR_FORMAT_GZIP = C.GZlibCompressorFormat(C.G_ZLIB_COMPRESSOR_FORMAT_GZIP)
	ZLIB_COMPRESSOR_FORMAT_RAW  = C.GZlibCompressorFormat(C.G_ZLIB_COMPRESSOR_FORMAT_RAW)
)