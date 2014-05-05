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

gchar* __G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME = G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME;
gchar* __G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE = G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE;
gchar* __G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE = G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE;
gchar* __G_FILE_ATTRIBUTE_ACCESS_CAN_READ = G_FILE_ATTRIBUTE_ACCESS_CAN_READ;
gchar* __G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME = G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME;
gchar* __G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH = G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH;
gchar* __G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE = G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE;
gchar* __G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE = G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE;
gchar* __G_FILE_ATTRIBUTE_DOS_IS_SYSTEM = G_FILE_ATTRIBUTE_DOS_IS_SYSTEM;
gchar* __G_FILE_ATTRIBUTE_ETAG_VALUE = G_FILE_ATTRIBUTE_ETAG_VALUE;
gchar* __G_FILE_ATTRIBUTE_FILESYSTEM_FREE = G_FILE_ATTRIBUTE_FILESYSTEM_FREE;
gchar* __G_FILE_ATTRIBUTE_FILESYSTEM_READONLY = G_FILE_ATTRIBUTE_FILESYSTEM_READONLY;
gchar* __G_FILE_ATTRIBUTE_FILESYSTEM_SIZE = G_FILE_ATTRIBUTE_FILESYSTEM_SIZE;
gchar* __G_FILE_ATTRIBUTE_FILESYSTEM_TYPE = G_FILE_ATTRIBUTE_FILESYSTEM_TYPE;
gchar* __G_FILE_ATTRIBUTE_FILESYSTEM_USED = G_FILE_ATTRIBUTE_FILESYSTEM_USED;
gchar* __G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW = G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW;
gchar* __G_FILE_ATTRIBUTE_GVFS_BACKEND = G_FILE_ATTRIBUTE_GVFS_BACKEND;
gchar* __G_FILE_ATTRIBUTE_ID_FILE = G_FILE_ATTRIBUTE_ID_FILE;
gchar* __G_FILE_ATTRIBUTE_ID_FILESYSTEM = G_FILE_ATTRIBUTE_ID_FILESYSTEM;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT = G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT = G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL = G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START = G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED = G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP = G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT = G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI = G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC = G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE = G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE = G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE;
gchar* __G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE = G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE;
gchar* __G_FILE_ATTRIBUTE_OWNER_GROUP = G_FILE_ATTRIBUTE_OWNER_GROUP;
gchar* __G_FILE_ATTRIBUTE_OWNER_USER = G_FILE_ATTRIBUTE_OWNER_USER;
gchar* __G_FILE_ATTRIBUTE_OWNER_USER_REAL = G_FILE_ATTRIBUTE_OWNER_USER_REAL;
gchar* __G_FILE_ATTRIBUTE_PREVIEW_ICON = G_FILE_ATTRIBUTE_PREVIEW_ICON;
gchar* __G_FILE_ATTRIBUTE_SELINUX_CONTEXT = G_FILE_ATTRIBUTE_SELINUX_CONTEXT;
gchar* __G_FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE = G_FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE;
gchar* __G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE = G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE;
gchar* __G_FILE_ATTRIBUTE_STANDARD_COPY_NAME = G_FILE_ATTRIBUTE_STANDARD_COPY_NAME;
gchar* __G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION = G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION;
gchar* __G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME = G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME;
gchar* __G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME = G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME;
gchar* __G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE = G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE;
gchar* __G_FILE_ATTRIBUTE_STANDARD_ICON = G_FILE_ATTRIBUTE_STANDARD_ICON;
gchar* __G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP = G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP;
gchar* __G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN = G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN;
gchar* __G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK = G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK;
gchar* __G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL = G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL;
gchar* __G_FILE_ATTRIBUTE_STANDARD_NAME = G_FILE_ATTRIBUTE_STANDARD_NAME;
gchar* __G_FILE_ATTRIBUTE_STANDARD_SIZE = G_FILE_ATTRIBUTE_STANDARD_SIZE;
gchar* __G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER = G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER;
gchar* __G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON = G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON;
gchar* __G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET = G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET;
gchar* __G_FILE_ATTRIBUTE_STANDARD_TARGET_URI = G_FILE_ATTRIBUTE_STANDARD_TARGET_URI;
gchar* __G_FILE_ATTRIBUTE_STANDARD_TYPE = G_FILE_ATTRIBUTE_STANDARD_TYPE;
gchar* __G_FILE_ATTRIBUTE_THUMBNAILING_FAILED = G_FILE_ATTRIBUTE_THUMBNAILING_FAILED;
gchar* __G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID = G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID;
gchar* __G_FILE_ATTRIBUTE_THUMBNAIL_PATH = G_FILE_ATTRIBUTE_THUMBNAIL_PATH;
gchar* __G_FILE_ATTRIBUTE_TIME_ACCESS = G_FILE_ATTRIBUTE_TIME_ACCESS;
gchar* __G_FILE_ATTRIBUTE_TIME_ACCESS_USEC = G_FILE_ATTRIBUTE_TIME_ACCESS_USEC;
gchar* __G_FILE_ATTRIBUTE_TIME_CHANGED = G_FILE_ATTRIBUTE_TIME_CHANGED;
gchar* __G_FILE_ATTRIBUTE_TIME_CHANGED_USEC = G_FILE_ATTRIBUTE_TIME_CHANGED_USEC;
gchar* __G_FILE_ATTRIBUTE_TIME_CREATED = G_FILE_ATTRIBUTE_TIME_CREATED;
gchar* __G_FILE_ATTRIBUTE_TIME_CREATED_USEC = G_FILE_ATTRIBUTE_TIME_CREATED_USEC;
gchar* __G_FILE_ATTRIBUTE_TIME_MODIFIED = G_FILE_ATTRIBUTE_TIME_MODIFIED;
gchar* __G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC = G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC;
gchar* __G_FILE_ATTRIBUTE_TRASH_DELETION_DATE = G_FILE_ATTRIBUTE_TRASH_DELETION_DATE;
gchar* __G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT = G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT;
gchar* __G_FILE_ATTRIBUTE_TRASH_ORIG_PATH = G_FILE_ATTRIBUTE_TRASH_ORIG_PATH;
gchar* __G_FILE_ATTRIBUTE_UNIX_BLOCKS = G_FILE_ATTRIBUTE_UNIX_BLOCKS;
gchar* __G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE = G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE;
gchar* __G_FILE_ATTRIBUTE_UNIX_DEVICE = G_FILE_ATTRIBUTE_UNIX_DEVICE;
gchar* __G_FILE_ATTRIBUTE_UNIX_GID = G_FILE_ATTRIBUTE_UNIX_GID;
gchar* __G_FILE_ATTRIBUTE_UNIX_INODE = G_FILE_ATTRIBUTE_UNIX_INODE;
gchar* __G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT = G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT;
gchar* __G_FILE_ATTRIBUTE_UNIX_MODE = G_FILE_ATTRIBUTE_UNIX_MODE;
gchar* __G_FILE_ATTRIBUTE_UNIX_NLINK = G_FILE_ATTRIBUTE_UNIX_NLINK;
gchar* __G_FILE_ATTRIBUTE_UNIX_RDEV = G_FILE_ATTRIBUTE_UNIX_RDEV;
gchar* __G_FILE_ATTRIBUTE_UNIX_UID = G_FILE_ATTRIBUTE_UNIX_UID;
gchar* __G_MENU_ATTRIBUTE_ACTION = G_MENU_ATTRIBUTE_ACTION;
gchar* __G_MENU_ATTRIBUTE_ACTION_NAMESPACE = G_MENU_ATTRIBUTE_ACTION_NAMESPACE;
gchar* __G_MENU_ATTRIBUTE_ICON = G_MENU_ATTRIBUTE_ICON;
gchar* __G_MENU_ATTRIBUTE_LABEL = G_MENU_ATTRIBUTE_LABEL;
gchar* __G_MENU_ATTRIBUTE_TARGET = G_MENU_ATTRIBUTE_TARGET;
gchar* __G_MENU_LINK_SECTION = G_MENU_LINK_SECTION;
gchar* __G_MENU_LINK_SUBMENU = G_MENU_LINK_SUBMENU;
gchar* __G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME = G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME;
gchar* __G_NETWORK_MONITOR_EXTENSION_POINT_NAME = G_NETWORK_MONITOR_EXTENSION_POINT_NAME;
gchar* __G_PROXY_EXTENSION_POINT_NAME = G_PROXY_EXTENSION_POINT_NAME;
gchar* __G_PROXY_RESOLVER_EXTENSION_POINT_NAME = G_PROXY_RESOLVER_EXTENSION_POINT_NAME;
gchar* __G_TLS_BACKEND_EXTENSION_POINT_NAME = G_TLS_BACKEND_EXTENSION_POINT_NAME;
gchar* __G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT = G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT;
gchar* __G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER = G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER;
gchar* __G_VFS_EXTENSION_POINT_NAME = G_VFS_EXTENSION_POINT_NAME;
gchar* __G_VOLUME_IDENTIFIER_KIND_CLASS = G_VOLUME_IDENTIFIER_KIND_CLASS;
gchar* __G_VOLUME_IDENTIFIER_KIND_HAL_UDI = G_VOLUME_IDENTIFIER_KIND_HAL_UDI;
gchar* __G_VOLUME_IDENTIFIER_KIND_LABEL = G_VOLUME_IDENTIFIER_KIND_LABEL;
gchar* __G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT = G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT;
gchar* __G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE = G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE;
gchar* __G_VOLUME_IDENTIFIER_KIND_UUID = G_VOLUME_IDENTIFIER_KIND_UUID;
gchar* __G_VOLUME_MONITOR_EXTENSION_POINT_NAME = G_VOLUME_MONITOR_EXTENSION_POINT_NAME;
*/
import "C"

var (
	DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME      = C.__G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME
	FILE_ATTRIBUTE_ACCESS_CAN_DELETE                  = C.__G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE
	FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE                 = C.__G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE
	FILE_ATTRIBUTE_ACCESS_CAN_READ                    = C.__G_FILE_ATTRIBUTE_ACCESS_CAN_READ
	FILE_ATTRIBUTE_ACCESS_CAN_RENAME                  = C.__G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME
	FILE_ATTRIBUTE_ACCESS_CAN_TRASH                   = C.__G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH
	FILE_ATTRIBUTE_ACCESS_CAN_WRITE                   = C.__G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE
	FILE_ATTRIBUTE_DOS_IS_ARCHIVE                     = C.__G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE
	FILE_ATTRIBUTE_DOS_IS_SYSTEM                      = C.__G_FILE_ATTRIBUTE_DOS_IS_SYSTEM
	FILE_ATTRIBUTE_ETAG_VALUE                         = C.__G_FILE_ATTRIBUTE_ETAG_VALUE
	FILE_ATTRIBUTE_FILESYSTEM_FREE                    = C.__G_FILE_ATTRIBUTE_FILESYSTEM_FREE
	FILE_ATTRIBUTE_FILESYSTEM_READONLY                = C.__G_FILE_ATTRIBUTE_FILESYSTEM_READONLY
	FILE_ATTRIBUTE_FILESYSTEM_SIZE                    = C.__G_FILE_ATTRIBUTE_FILESYSTEM_SIZE
	FILE_ATTRIBUTE_FILESYSTEM_TYPE                    = C.__G_FILE_ATTRIBUTE_FILESYSTEM_TYPE
	FILE_ATTRIBUTE_FILESYSTEM_USED                    = C.__G_FILE_ATTRIBUTE_FILESYSTEM_USED
	FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW             = C.__G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW
	FILE_ATTRIBUTE_GVFS_BACKEND                       = C.__G_FILE_ATTRIBUTE_GVFS_BACKEND
	FILE_ATTRIBUTE_ID_FILE                            = C.__G_FILE_ATTRIBUTE_ID_FILE
	FILE_ATTRIBUTE_ID_FILESYSTEM                      = C.__G_FILE_ATTRIBUTE_ID_FILESYSTEM
	FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT                = C.__G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT
	FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT                = C.__G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT
	FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL                 = C.__G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL
	FILE_ATTRIBUTE_MOUNTABLE_CAN_START                = C.__G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START
	FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED       = C.__G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED
	FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP                 = C.__G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP
	FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT              = C.__G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT
	FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI                  = C.__G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI
	FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC = C.__G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC
	FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE          = C.__G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE
	FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE              = C.__G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE
	FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE         = C.__G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE
	FILE_ATTRIBUTE_OWNER_GROUP                        = C.__G_FILE_ATTRIBUTE_OWNER_GROUP
	FILE_ATTRIBUTE_OWNER_USER                         = C.__G_FILE_ATTRIBUTE_OWNER_USER
	FILE_ATTRIBUTE_OWNER_USER_REAL                    = C.__G_FILE_ATTRIBUTE_OWNER_USER_REAL
	FILE_ATTRIBUTE_PREVIEW_ICON                       = C.__G_FILE_ATTRIBUTE_PREVIEW_ICON
	FILE_ATTRIBUTE_SELINUX_CONTEXT                    = C.__G_FILE_ATTRIBUTE_SELINUX_CONTEXT
	FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE            = C.__G_FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE
	FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE              = C.__G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE
	FILE_ATTRIBUTE_STANDARD_COPY_NAME                 = C.__G_FILE_ATTRIBUTE_STANDARD_COPY_NAME
	FILE_ATTRIBUTE_STANDARD_DESCRIPTION               = C.__G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION
	FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME              = C.__G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME
	FILE_ATTRIBUTE_STANDARD_EDIT_NAME                 = C.__G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME
	FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE         = C.__G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE
	FILE_ATTRIBUTE_STANDARD_ICON                      = C.__G_FILE_ATTRIBUTE_STANDARD_ICON
	FILE_ATTRIBUTE_STANDARD_IS_BACKUP                 = C.__G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP
	FILE_ATTRIBUTE_STANDARD_IS_HIDDEN                 = C.__G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN
	FILE_ATTRIBUTE_STANDARD_IS_SYMLINK                = C.__G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK
	FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL                = C.__G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL
	FILE_ATTRIBUTE_STANDARD_NAME                      = C.__G_FILE_ATTRIBUTE_STANDARD_NAME
	FILE_ATTRIBUTE_STANDARD_SIZE                      = C.__G_FILE_ATTRIBUTE_STANDARD_SIZE
	FILE_ATTRIBUTE_STANDARD_SORT_ORDER                = C.__G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER
	FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON             = C.__G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON
	FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET            = C.__G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET
	FILE_ATTRIBUTE_STANDARD_TARGET_URI                = C.__G_FILE_ATTRIBUTE_STANDARD_TARGET_URI
	FILE_ATTRIBUTE_STANDARD_TYPE                      = C.__G_FILE_ATTRIBUTE_STANDARD_TYPE
	FILE_ATTRIBUTE_THUMBNAILING_FAILED                = C.__G_FILE_ATTRIBUTE_THUMBNAILING_FAILED
	FILE_ATTRIBUTE_THUMBNAIL_IS_VALID                 = C.__G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID
	FILE_ATTRIBUTE_THUMBNAIL_PATH                     = C.__G_FILE_ATTRIBUTE_THUMBNAIL_PATH
	FILE_ATTRIBUTE_TIME_ACCESS                        = C.__G_FILE_ATTRIBUTE_TIME_ACCESS
	FILE_ATTRIBUTE_TIME_ACCESS_USEC                   = C.__G_FILE_ATTRIBUTE_TIME_ACCESS_USEC
	FILE_ATTRIBUTE_TIME_CHANGED                       = C.__G_FILE_ATTRIBUTE_TIME_CHANGED
	FILE_ATTRIBUTE_TIME_CHANGED_USEC                  = C.__G_FILE_ATTRIBUTE_TIME_CHANGED_USEC
	FILE_ATTRIBUTE_TIME_CREATED                       = C.__G_FILE_ATTRIBUTE_TIME_CREATED
	FILE_ATTRIBUTE_TIME_CREATED_USEC                  = C.__G_FILE_ATTRIBUTE_TIME_CREATED_USEC
	FILE_ATTRIBUTE_TIME_MODIFIED                      = C.__G_FILE_ATTRIBUTE_TIME_MODIFIED
	FILE_ATTRIBUTE_TIME_MODIFIED_USEC                 = C.__G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC
	FILE_ATTRIBUTE_TRASH_DELETION_DATE                = C.__G_FILE_ATTRIBUTE_TRASH_DELETION_DATE
	FILE_ATTRIBUTE_TRASH_ITEM_COUNT                   = C.__G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT
	FILE_ATTRIBUTE_TRASH_ORIG_PATH                    = C.__G_FILE_ATTRIBUTE_TRASH_ORIG_PATH
	FILE_ATTRIBUTE_UNIX_BLOCKS                        = C.__G_FILE_ATTRIBUTE_UNIX_BLOCKS
	FILE_ATTRIBUTE_UNIX_BLOCK_SIZE                    = C.__G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE
	FILE_ATTRIBUTE_UNIX_DEVICE                        = C.__G_FILE_ATTRIBUTE_UNIX_DEVICE
	FILE_ATTRIBUTE_UNIX_GID                           = C.__G_FILE_ATTRIBUTE_UNIX_GID
	FILE_ATTRIBUTE_UNIX_INODE                         = C.__G_FILE_ATTRIBUTE_UNIX_INODE
	FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT                 = C.__G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT
	FILE_ATTRIBUTE_UNIX_MODE                          = C.__G_FILE_ATTRIBUTE_UNIX_MODE
	FILE_ATTRIBUTE_UNIX_NLINK                         = C.__G_FILE_ATTRIBUTE_UNIX_NLINK
	FILE_ATTRIBUTE_UNIX_RDEV                          = C.__G_FILE_ATTRIBUTE_UNIX_RDEV
	FILE_ATTRIBUTE_UNIX_UID                           = C.__G_FILE_ATTRIBUTE_UNIX_UID
	MENU_ATTRIBUTE_ACTION                             = C.__G_MENU_ATTRIBUTE_ACTION
	MENU_ATTRIBUTE_ACTION_NAMESPACE                   = C.__G_MENU_ATTRIBUTE_ACTION_NAMESPACE
	MENU_ATTRIBUTE_ICON                               = C.__G_MENU_ATTRIBUTE_ICON
	MENU_ATTRIBUTE_LABEL                              = C.__G_MENU_ATTRIBUTE_LABEL
	MENU_ATTRIBUTE_TARGET                             = C.__G_MENU_ATTRIBUTE_TARGET
	MENU_LINK_SECTION                                 = C.__G_MENU_LINK_SECTION
	MENU_LINK_SUBMENU                                 = C.__G_MENU_LINK_SUBMENU
	NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME        = C.__G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME
	NETWORK_MONITOR_EXTENSION_POINT_NAME              = C.__G_NETWORK_MONITOR_EXTENSION_POINT_NAME
	PROXY_EXTENSION_POINT_NAME                        = C.__G_PROXY_EXTENSION_POINT_NAME
	PROXY_RESOLVER_EXTENSION_POINT_NAME               = C.__G_PROXY_RESOLVER_EXTENSION_POINT_NAME
	TLS_BACKEND_EXTENSION_POINT_NAME                  = C.__G_TLS_BACKEND_EXTENSION_POINT_NAME
	TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT          = C.__G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT
	TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER          = C.__G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER
	VFS_EXTENSION_POINT_NAME                          = C.__G_VFS_EXTENSION_POINT_NAME
	VOLUME_IDENTIFIER_KIND_CLASS                      = C.__G_VOLUME_IDENTIFIER_KIND_CLASS
	VOLUME_IDENTIFIER_KIND_HAL_UDI                    = C.__G_VOLUME_IDENTIFIER_KIND_HAL_UDI
	VOLUME_IDENTIFIER_KIND_LABEL                      = C.__G_VOLUME_IDENTIFIER_KIND_LABEL
	VOLUME_IDENTIFIER_KIND_NFS_MOUNT                  = C.__G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT
	VOLUME_IDENTIFIER_KIND_UNIX_DEVICE                = C.__G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE
	VOLUME_IDENTIFIER_KIND_UUID                       = C.__G_VOLUME_IDENTIFIER_KIND_UUID
	VOLUME_MONITOR_EXTENSION_POINT_NAME               = C.__G_VOLUME_MONITOR_EXTENSION_POINT_NAME
)
