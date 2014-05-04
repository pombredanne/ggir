package glib

/*
#include <glib.h>
#include <glib-object.h>
#include <glib-unix.h>
#include <glib/gstdio.h>
#include <stdlib.h>
*/
import "C"

var (
	// BookmarkFileError
	BOOKMARK_FILE_ERROR_INVALID_URI        = C.GBookmarkFileError(C.G_BOOKMARK_FILE_ERROR_INVALID_URI)
	BOOKMARK_FILE_ERROR_INVALID_VALUE      = C.GBookmarkFileError(C.G_BOOKMARK_FILE_ERROR_INVALID_VALUE)
	BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED = C.GBookmarkFileError(C.G_BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED)
	BOOKMARK_FILE_ERROR_URI_NOT_FOUND      = C.GBookmarkFileError(C.G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND)
	BOOKMARK_FILE_ERROR_READ               = C.GBookmarkFileError(C.G_BOOKMARK_FILE_ERROR_READ)
	BOOKMARK_FILE_ERROR_UNKNOWN_ENCODING   = C.GBookmarkFileError(C.G_BOOKMARK_FILE_ERROR_UNKNOWN_ENCODING)
	BOOKMARK_FILE_ERROR_WRITE              = C.GBookmarkFileError(C.G_BOOKMARK_FILE_ERROR_WRITE)
	BOOKMARK_FILE_ERROR_FILE_NOT_FOUND     = C.GBookmarkFileError(C.G_BOOKMARK_FILE_ERROR_FILE_NOT_FOUND)

	// ChecksumType
	CHECKSUM_MD5    = C.GChecksumType(C.G_CHECKSUM_MD5)
	CHECKSUM_SHA1   = C.GChecksumType(C.G_CHECKSUM_SHA1)
	CHECKSUM_SHA256 = C.GChecksumType(C.G_CHECKSUM_SHA256)
	CHECKSUM_SHA512 = C.GChecksumType(C.G_CHECKSUM_SHA512)

	// ConvertError
	CONVERT_ERROR_NO_CONVERSION     = C.GConvertError(C.G_CONVERT_ERROR_NO_CONVERSION)
	CONVERT_ERROR_ILLEGAL_SEQUENCE  = C.GConvertError(C.G_CONVERT_ERROR_ILLEGAL_SEQUENCE)
	CONVERT_ERROR_FAILED            = C.GConvertError(C.G_CONVERT_ERROR_FAILED)
	CONVERT_ERROR_PARTIAL_INPUT     = C.GConvertError(C.G_CONVERT_ERROR_PARTIAL_INPUT)
	CONVERT_ERROR_BAD_URI           = C.GConvertError(C.G_CONVERT_ERROR_BAD_URI)
	CONVERT_ERROR_NOT_ABSOLUTE_PATH = C.GConvertError(C.G_CONVERT_ERROR_NOT_ABSOLUTE_PATH)
	CONVERT_ERROR_NO_MEMORY         = C.GConvertError(C.G_CONVERT_ERROR_NO_MEMORY)

	// DateDMY
	DATE_DAY   = C.GDateDMY(C.G_DATE_DAY)
	DATE_MONTH = C.GDateDMY(C.G_DATE_MONTH)
	DATE_YEAR  = C.GDateDMY(C.G_DATE_YEAR)

	// DateMonth
	DATE_BAD_MONTH = C.GDateMonth(C.G_DATE_BAD_MONTH)
	DATE_JANUARY   = C.GDateMonth(C.G_DATE_JANUARY)
	DATE_FEBRUARY  = C.GDateMonth(C.G_DATE_FEBRUARY)
	DATE_MARCH     = C.GDateMonth(C.G_DATE_MARCH)
	DATE_APRIL     = C.GDateMonth(C.G_DATE_APRIL)
	DATE_MAY       = C.GDateMonth(C.G_DATE_MAY)
	DATE_JUNE      = C.GDateMonth(C.G_DATE_JUNE)
	DATE_JULY      = C.GDateMonth(C.G_DATE_JULY)
	DATE_AUGUST    = C.GDateMonth(C.G_DATE_AUGUST)
	DATE_SEPTEMBER = C.GDateMonth(C.G_DATE_SEPTEMBER)
	DATE_OCTOBER   = C.GDateMonth(C.G_DATE_OCTOBER)
	DATE_NOVEMBER  = C.GDateMonth(C.G_DATE_NOVEMBER)
	DATE_DECEMBER  = C.GDateMonth(C.G_DATE_DECEMBER)

	// DateWeekday
	DATE_BAD_WEEKDAY = C.GDateWeekday(C.G_DATE_BAD_WEEKDAY)
	DATE_MONDAY      = C.GDateWeekday(C.G_DATE_MONDAY)
	DATE_TUESDAY     = C.GDateWeekday(C.G_DATE_TUESDAY)
	DATE_WEDNESDAY   = C.GDateWeekday(C.G_DATE_WEDNESDAY)
	DATE_THURSDAY    = C.GDateWeekday(C.G_DATE_THURSDAY)
	DATE_FRIDAY      = C.GDateWeekday(C.G_DATE_FRIDAY)
	DATE_SATURDAY    = C.GDateWeekday(C.G_DATE_SATURDAY)
	DATE_SUNDAY      = C.GDateWeekday(C.G_DATE_SUNDAY)

	// ErrorType
	ERR_UNKNOWN              = C.GErrorType(C.G_ERR_UNKNOWN)
	ERR_UNEXP_EOF            = C.GErrorType(C.G_ERR_UNEXP_EOF)
	ERR_UNEXP_EOF_IN_STRING  = C.GErrorType(C.G_ERR_UNEXP_EOF_IN_STRING)
	ERR_UNEXP_EOF_IN_COMMENT = C.GErrorType(C.G_ERR_UNEXP_EOF_IN_COMMENT)
	ERR_NON_DIGIT_IN_CONST   = C.GErrorType(C.G_ERR_NON_DIGIT_IN_CONST)
	ERR_DIGIT_RADIX          = C.GErrorType(C.G_ERR_DIGIT_RADIX)
	ERR_FLOAT_RADIX          = C.GErrorType(C.G_ERR_FLOAT_RADIX)
	ERR_FLOAT_MALFORMED      = C.GErrorType(C.G_ERR_FLOAT_MALFORMED)

	// FileError
	FILE_ERROR_EXIST       = C.GFileError(C.G_FILE_ERROR_EXIST)
	FILE_ERROR_ISDIR       = C.GFileError(C.G_FILE_ERROR_ISDIR)
	FILE_ERROR_ACCES       = C.GFileError(C.G_FILE_ERROR_ACCES)
	FILE_ERROR_NAMETOOLONG = C.GFileError(C.G_FILE_ERROR_NAMETOOLONG)
	FILE_ERROR_NOENT       = C.GFileError(C.G_FILE_ERROR_NOENT)
	FILE_ERROR_NOTDIR      = C.GFileError(C.G_FILE_ERROR_NOTDIR)
	FILE_ERROR_NXIO        = C.GFileError(C.G_FILE_ERROR_NXIO)
	FILE_ERROR_NODEV       = C.GFileError(C.G_FILE_ERROR_NODEV)
	FILE_ERROR_ROFS        = C.GFileError(C.G_FILE_ERROR_ROFS)
	FILE_ERROR_TXTBSY      = C.GFileError(C.G_FILE_ERROR_TXTBSY)
	FILE_ERROR_FAULT       = C.GFileError(C.G_FILE_ERROR_FAULT)
	FILE_ERROR_LOOP        = C.GFileError(C.G_FILE_ERROR_LOOP)
	FILE_ERROR_NOSPC       = C.GFileError(C.G_FILE_ERROR_NOSPC)
	FILE_ERROR_NOMEM       = C.GFileError(C.G_FILE_ERROR_NOMEM)
	FILE_ERROR_MFILE       = C.GFileError(C.G_FILE_ERROR_MFILE)
	FILE_ERROR_NFILE       = C.GFileError(C.G_FILE_ERROR_NFILE)
	FILE_ERROR_BADF        = C.GFileError(C.G_FILE_ERROR_BADF)
	FILE_ERROR_INVAL       = C.GFileError(C.G_FILE_ERROR_INVAL)
	FILE_ERROR_PIPE        = C.GFileError(C.G_FILE_ERROR_PIPE)
	FILE_ERROR_AGAIN       = C.GFileError(C.G_FILE_ERROR_AGAIN)
	FILE_ERROR_INTR        = C.GFileError(C.G_FILE_ERROR_INTR)
	FILE_ERROR_IO          = C.GFileError(C.G_FILE_ERROR_IO)
	FILE_ERROR_PERM        = C.GFileError(C.G_FILE_ERROR_PERM)
	FILE_ERROR_NOSYS       = C.GFileError(C.G_FILE_ERROR_NOSYS)
	FILE_ERROR_FAILED      = C.GFileError(C.G_FILE_ERROR_FAILED)

	// IOChannelError
	IO_CHANNEL_ERROR_FBIG     = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_FBIG)
	IO_CHANNEL_ERROR_INVAL    = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_INVAL)
	IO_CHANNEL_ERROR_IO       = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_IO)
	IO_CHANNEL_ERROR_ISDIR    = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_ISDIR)
	IO_CHANNEL_ERROR_NOSPC    = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_NOSPC)
	IO_CHANNEL_ERROR_NXIO     = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_NXIO)
	IO_CHANNEL_ERROR_OVERFLOW = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_OVERFLOW)
	IO_CHANNEL_ERROR_PIPE     = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_PIPE)
	IO_CHANNEL_ERROR_FAILED   = C.GIOChannelError(C.G_IO_CHANNEL_ERROR_FAILED)

	// IOError
	IO_ERROR_NONE    = C.GIOError(C.G_IO_ERROR_NONE)
	IO_ERROR_AGAIN   = C.GIOError(C.G_IO_ERROR_AGAIN)
	IO_ERROR_INVAL   = C.GIOError(C.G_IO_ERROR_INVAL)
	IO_ERROR_UNKNOWN = C.GIOError(C.G_IO_ERROR_UNKNOWN)

	// IOStatus
	IO_STATUS_ERROR  = C.GIOStatus(C.G_IO_STATUS_ERROR)
	IO_STATUS_NORMAL = C.GIOStatus(C.G_IO_STATUS_NORMAL)
	IO_STATUS_EOF    = C.GIOStatus(C.G_IO_STATUS_EOF)
	IO_STATUS_AGAIN  = C.GIOStatus(C.G_IO_STATUS_AGAIN)

	// KeyFileError
	KEY_FILE_ERROR_UNKNOWN_ENCODING = C.GKeyFileError(C.G_KEY_FILE_ERROR_UNKNOWN_ENCODING)
	KEY_FILE_ERROR_PARSE            = C.GKeyFileError(C.G_KEY_FILE_ERROR_PARSE)
	KEY_FILE_ERROR_NOT_FOUND        = C.GKeyFileError(C.G_KEY_FILE_ERROR_NOT_FOUND)
	KEY_FILE_ERROR_KEY_NOT_FOUND    = C.GKeyFileError(C.G_KEY_FILE_ERROR_KEY_NOT_FOUND)
	KEY_FILE_ERROR_GROUP_NOT_FOUND  = C.GKeyFileError(C.G_KEY_FILE_ERROR_GROUP_NOT_FOUND)
	KEY_FILE_ERROR_INVALID_VALUE    = C.GKeyFileError(C.G_KEY_FILE_ERROR_INVALID_VALUE)

	// MarkupError
	MARKUP_ERROR_BAD_UTF8          = C.GMarkupError(C.G_MARKUP_ERROR_BAD_UTF8)
	MARKUP_ERROR_EMPTY             = C.GMarkupError(C.G_MARKUP_ERROR_EMPTY)
	MARKUP_ERROR_PARSE             = C.GMarkupError(C.G_MARKUP_ERROR_PARSE)
	MARKUP_ERROR_UNKNOWN_ELEMENT   = C.GMarkupError(C.G_MARKUP_ERROR_UNKNOWN_ELEMENT)
	MARKUP_ERROR_UNKNOWN_ATTRIBUTE = C.GMarkupError(C.G_MARKUP_ERROR_UNKNOWN_ATTRIBUTE)
	MARKUP_ERROR_INVALID_CONTENT   = C.GMarkupError(C.G_MARKUP_ERROR_INVALID_CONTENT)
	MARKUP_ERROR_MISSING_ATTRIBUTE = C.GMarkupError(C.G_MARKUP_ERROR_MISSING_ATTRIBUTE)

	// NormalizeMode
	NORMALIZE_DEFAULT         = C.GNormalizeMode(C.G_NORMALIZE_DEFAULT)
	NORMALIZE_NFD             = C.GNormalizeMode(C.G_NORMALIZE_NFD)
	NORMALIZE_DEFAULT_COMPOSE = C.GNormalizeMode(C.G_NORMALIZE_DEFAULT_COMPOSE)
	NORMALIZE_NFC             = C.GNormalizeMode(C.G_NORMALIZE_NFC)
	NORMALIZE_ALL             = C.GNormalizeMode(C.G_NORMALIZE_ALL)
	NORMALIZE_NFKD            = C.GNormalizeMode(C.G_NORMALIZE_NFKD)
	NORMALIZE_ALL_COMPOSE     = C.GNormalizeMode(C.G_NORMALIZE_ALL_COMPOSE)
	NORMALIZE_NFKC            = C.GNormalizeMode(C.G_NORMALIZE_NFKC)

	// OnceStatus
	ONCE_STATUS_NOTCALLED = C.GOnceStatus(C.G_ONCE_STATUS_NOTCALLED)
	ONCE_STATUS_PROGRESS  = C.GOnceStatus(C.G_ONCE_STATUS_PROGRESS)
	ONCE_STATUS_READY     = C.GOnceStatus(C.G_ONCE_STATUS_READY)

	// OptionArg
	OPTION_ARG_NONE           = C.GOptionArg(C.G_OPTION_ARG_NONE)
	OPTION_ARG_STRING         = C.GOptionArg(C.G_OPTION_ARG_STRING)
	OPTION_ARG_INT            = C.GOptionArg(C.G_OPTION_ARG_INT)
	OPTION_ARG_CALLBACK       = C.GOptionArg(C.G_OPTION_ARG_CALLBACK)
	OPTION_ARG_FILENAME       = C.GOptionArg(C.G_OPTION_ARG_FILENAME)
	OPTION_ARG_STRING_ARRAY   = C.GOptionArg(C.G_OPTION_ARG_STRING_ARRAY)
	OPTION_ARG_FILENAME_ARRAY = C.GOptionArg(C.G_OPTION_ARG_FILENAME_ARRAY)
	OPTION_ARG_DOUBLE         = C.GOptionArg(C.G_OPTION_ARG_DOUBLE)
	OPTION_ARG_INT64          = C.GOptionArg(C.G_OPTION_ARG_INT64)

	// OptionError
	OPTION_ERROR_UNKNOWN_OPTION = C.GOptionError(C.G_OPTION_ERROR_UNKNOWN_OPTION)
	OPTION_ERROR_BAD_VALUE      = C.GOptionError(C.G_OPTION_ERROR_BAD_VALUE)
	OPTION_ERROR_FAILED         = C.GOptionError(C.G_OPTION_ERROR_FAILED)

	// RegexError
	REGEX_ERROR_COMPILE                                      = C.GRegexError(C.G_REGEX_ERROR_COMPILE)
	REGEX_ERROR_OPTIMIZE                                     = C.GRegexError(C.G_REGEX_ERROR_OPTIMIZE)
	REGEX_ERROR_REPLACE                                      = C.GRegexError(C.G_REGEX_ERROR_REPLACE)
	REGEX_ERROR_MATCH                                        = C.GRegexError(C.G_REGEX_ERROR_MATCH)
	REGEX_ERROR_INTERNAL                                     = C.GRegexError(C.G_REGEX_ERROR_INTERNAL)
	REGEX_ERROR_STRAY_BACKSLASH                              = C.GRegexError(C.G_REGEX_ERROR_STRAY_BACKSLASH)
	REGEX_ERROR_MISSING_CONTROL_CHAR                         = C.GRegexError(C.G_REGEX_ERROR_MISSING_CONTROL_CHAR)
	REGEX_ERROR_UNRECOGNIZED_ESCAPE                          = C.GRegexError(C.G_REGEX_ERROR_UNRECOGNIZED_ESCAPE)
	REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER                     = C.GRegexError(C.G_REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER)
	REGEX_ERROR_QUANTIFIER_TOO_BIG                           = C.GRegexError(C.G_REGEX_ERROR_QUANTIFIER_TOO_BIG)
	REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS                 = C.GRegexError(C.G_REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS)
	REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS            = C.GRegexError(C.G_REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS)
	REGEX_ERROR_RANGE_OUT_OF_ORDER                           = C.GRegexError(C.G_REGEX_ERROR_RANGE_OUT_OF_ORDER)
	REGEX_ERROR_NOTHING_TO_REPEAT                            = C.GRegexError(C.G_REGEX_ERROR_NOTHING_TO_REPEAT)
	REGEX_ERROR_UNRECOGNIZED_CHARACTER                       = C.GRegexError(C.G_REGEX_ERROR_UNRECOGNIZED_CHARACTER)
	REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS              = C.GRegexError(C.G_REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS)
	REGEX_ERROR_UNMATCHED_PARENTHESIS                        = C.GRegexError(C.G_REGEX_ERROR_UNMATCHED_PARENTHESIS)
	REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE              = C.GRegexError(C.G_REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE)
	REGEX_ERROR_UNTERMINATED_COMMENT                         = C.GRegexError(C.G_REGEX_ERROR_UNTERMINATED_COMMENT)
	REGEX_ERROR_EXPRESSION_TOO_LARGE                         = C.GRegexError(C.G_REGEX_ERROR_EXPRESSION_TOO_LARGE)
	REGEX_ERROR_MEMORY_ERROR                                 = C.GRegexError(C.G_REGEX_ERROR_MEMORY_ERROR)
	REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND                   = C.GRegexError(C.G_REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND)
	REGEX_ERROR_MALFORMED_CONDITION                          = C.GRegexError(C.G_REGEX_ERROR_MALFORMED_CONDITION)
	REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES                = C.GRegexError(C.G_REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES)
	REGEX_ERROR_ASSERTION_EXPECTED                           = C.GRegexError(C.G_REGEX_ERROR_ASSERTION_EXPECTED)
	REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME                     = C.GRegexError(C.G_REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME)
	REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED       = C.GRegexError(C.G_REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED)
	REGEX_ERROR_HEX_CODE_TOO_LARGE                           = C.GRegexError(C.G_REGEX_ERROR_HEX_CODE_TOO_LARGE)
	REGEX_ERROR_INVALID_CONDITION                            = C.GRegexError(C.G_REGEX_ERROR_INVALID_CONDITION)
	REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND              = C.GRegexError(C.G_REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND)
	REGEX_ERROR_INFINITE_LOOP                                = C.GRegexError(C.G_REGEX_ERROR_INFINITE_LOOP)
	REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR           = C.GRegexError(C.G_REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR)
	REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME                    = C.GRegexError(C.G_REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME)
	REGEX_ERROR_MALFORMED_PROPERTY                           = C.GRegexError(C.G_REGEX_ERROR_MALFORMED_PROPERTY)
	REGEX_ERROR_UNKNOWN_PROPERTY                             = C.GRegexError(C.G_REGEX_ERROR_UNKNOWN_PROPERTY)
	REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG                     = C.GRegexError(C.G_REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG)
	REGEX_ERROR_TOO_MANY_SUBPATTERNS                         = C.GRegexError(C.G_REGEX_ERROR_TOO_MANY_SUBPATTERNS)
	REGEX_ERROR_INVALID_OCTAL_VALUE                          = C.GRegexError(C.G_REGEX_ERROR_INVALID_OCTAL_VALUE)
	REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE                  = C.GRegexError(C.G_REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE)
	REGEX_ERROR_DEFINE_REPETION                              = C.GRegexError(C.G_REGEX_ERROR_DEFINE_REPETION)
	REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS                 = C.GRegexError(C.G_REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS)
	REGEX_ERROR_MISSING_BACK_REFERENCE                       = C.GRegexError(C.G_REGEX_ERROR_MISSING_BACK_REFERENCE)
	REGEX_ERROR_INVALID_RELATIVE_REFERENCE                   = C.GRegexError(C.G_REGEX_ERROR_INVALID_RELATIVE_REFERENCE)
	REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_FORBIDDEN = C.GRegexError(C.G_REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_FORBIDDEN)
	REGEX_ERROR_UNKNOWN_BACKTRACKING_CONTROL_VERB            = C.GRegexError(C.G_REGEX_ERROR_UNKNOWN_BACKTRACKING_CONTROL_VERB)
	REGEX_ERROR_NUMBER_TOO_BIG                               = C.GRegexError(C.G_REGEX_ERROR_NUMBER_TOO_BIG)
	REGEX_ERROR_MISSING_SUBPATTERN_NAME                      = C.GRegexError(C.G_REGEX_ERROR_MISSING_SUBPATTERN_NAME)
	REGEX_ERROR_MISSING_DIGIT                                = C.GRegexError(C.G_REGEX_ERROR_MISSING_DIGIT)
	REGEX_ERROR_INVALID_DATA_CHARACTER                       = C.GRegexError(C.G_REGEX_ERROR_INVALID_DATA_CHARACTER)
	REGEX_ERROR_EXTRA_SUBPATTERN_NAME                        = C.GRegexError(C.G_REGEX_ERROR_EXTRA_SUBPATTERN_NAME)
	REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_REQUIRED  = C.GRegexError(C.G_REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_REQUIRED)
	REGEX_ERROR_INVALID_CONTROL_CHAR                         = C.GRegexError(C.G_REGEX_ERROR_INVALID_CONTROL_CHAR)
	REGEX_ERROR_MISSING_NAME                                 = C.GRegexError(C.G_REGEX_ERROR_MISSING_NAME)
	REGEX_ERROR_NOT_SUPPORTED_IN_CLASS                       = C.GRegexError(C.G_REGEX_ERROR_NOT_SUPPORTED_IN_CLASS)
	REGEX_ERROR_TOO_MANY_FORWARD_REFERENCES                  = C.GRegexError(C.G_REGEX_ERROR_TOO_MANY_FORWARD_REFERENCES)
	REGEX_ERROR_NAME_TOO_LONG                                = C.GRegexError(C.G_REGEX_ERROR_NAME_TOO_LONG)
	REGEX_ERROR_CHARACTER_VALUE_TOO_LARGE                    = C.GRegexError(C.G_REGEX_ERROR_CHARACTER_VALUE_TOO_LARGE)

	// SeekType
	SEEK_CUR = C.GSeekType(C.G_SEEK_CUR)
	SEEK_SET = C.GSeekType(C.G_SEEK_SET)
	SEEK_END = C.GSeekType(C.G_SEEK_END)

	// ShellError
	SHELL_ERROR_BAD_QUOTING  = C.GShellError(C.G_SHELL_ERROR_BAD_QUOTING)
	SHELL_ERROR_EMPTY_STRING = C.GShellError(C.G_SHELL_ERROR_EMPTY_STRING)
	SHELL_ERROR_FAILED       = C.GShellError(C.G_SHELL_ERROR_FAILED)

	// SliceConfig
	SLICE_CONFIG_ALWAYS_MALLOC      = C.GSliceConfig(C.G_SLICE_CONFIG_ALWAYS_MALLOC)
	SLICE_CONFIG_BYPASS_MAGAZINES   = C.GSliceConfig(C.G_SLICE_CONFIG_BYPASS_MAGAZINES)
	SLICE_CONFIG_WORKING_SET_MSECS  = C.GSliceConfig(C.G_SLICE_CONFIG_WORKING_SET_MSECS)
	SLICE_CONFIG_COLOR_INCREMENT    = C.GSliceConfig(C.G_SLICE_CONFIG_COLOR_INCREMENT)
	SLICE_CONFIG_CHUNK_SIZES        = C.GSliceConfig(C.G_SLICE_CONFIG_CHUNK_SIZES)
	SLICE_CONFIG_CONTENTION_COUNTER = C.GSliceConfig(C.G_SLICE_CONFIG_CONTENTION_COUNTER)

	// SpawnError
	SPAWN_ERROR_FORK        = C.GSpawnError(C.G_SPAWN_ERROR_FORK)
	SPAWN_ERROR_READ        = C.GSpawnError(C.G_SPAWN_ERROR_READ)
	SPAWN_ERROR_CHDIR       = C.GSpawnError(C.G_SPAWN_ERROR_CHDIR)
	SPAWN_ERROR_ACCES       = C.GSpawnError(C.G_SPAWN_ERROR_ACCES)
	SPAWN_ERROR_PERM        = C.GSpawnError(C.G_SPAWN_ERROR_PERM)
	SPAWN_ERROR_TOO_BIG     = C.GSpawnError(C.G_SPAWN_ERROR_TOO_BIG)
	SPAWN_ERROR_2BIG        = C.GSpawnError(C.G_SPAWN_ERROR_2BIG)
	SPAWN_ERROR_NOEXEC      = C.GSpawnError(C.G_SPAWN_ERROR_NOEXEC)
	SPAWN_ERROR_NAMETOOLONG = C.GSpawnError(C.G_SPAWN_ERROR_NAMETOOLONG)
	SPAWN_ERROR_NOENT       = C.GSpawnError(C.G_SPAWN_ERROR_NOENT)
	SPAWN_ERROR_NOMEM       = C.GSpawnError(C.G_SPAWN_ERROR_NOMEM)
	SPAWN_ERROR_NOTDIR      = C.GSpawnError(C.G_SPAWN_ERROR_NOTDIR)
	SPAWN_ERROR_LOOP        = C.GSpawnError(C.G_SPAWN_ERROR_LOOP)
	SPAWN_ERROR_TXTBUSY     = C.GSpawnError(C.G_SPAWN_ERROR_TXTBUSY)
	SPAWN_ERROR_IO          = C.GSpawnError(C.G_SPAWN_ERROR_IO)
	SPAWN_ERROR_NFILE       = C.GSpawnError(C.G_SPAWN_ERROR_NFILE)
	SPAWN_ERROR_MFILE       = C.GSpawnError(C.G_SPAWN_ERROR_MFILE)
	SPAWN_ERROR_INVAL       = C.GSpawnError(C.G_SPAWN_ERROR_INVAL)
	SPAWN_ERROR_ISDIR       = C.GSpawnError(C.G_SPAWN_ERROR_ISDIR)
	SPAWN_ERROR_LIBBAD      = C.GSpawnError(C.G_SPAWN_ERROR_LIBBAD)
	SPAWN_ERROR_FAILED      = C.GSpawnError(C.G_SPAWN_ERROR_FAILED)

	// TestFileType
	TEST_DIST  = C.GTestFileType(C.G_TEST_DIST)
	TEST_BUILT = C.GTestFileType(C.G_TEST_BUILT)

	// TestLogType
	TEST_LOG_NONE         = C.GTestLogType(C.G_TEST_LOG_NONE)
	TEST_LOG_ERROR        = C.GTestLogType(C.G_TEST_LOG_ERROR)
	TEST_LOG_START_BINARY = C.GTestLogType(C.G_TEST_LOG_START_BINARY)
	TEST_LOG_LIST_CASE    = C.GTestLogType(C.G_TEST_LOG_LIST_CASE)
	TEST_LOG_SKIP_CASE    = C.GTestLogType(C.G_TEST_LOG_SKIP_CASE)
	TEST_LOG_START_CASE   = C.GTestLogType(C.G_TEST_LOG_START_CASE)
	TEST_LOG_STOP_CASE    = C.GTestLogType(C.G_TEST_LOG_STOP_CASE)
	TEST_LOG_MIN_RESULT   = C.GTestLogType(C.G_TEST_LOG_MIN_RESULT)
	TEST_LOG_MAX_RESULT   = C.GTestLogType(C.G_TEST_LOG_MAX_RESULT)
	TEST_LOG_MESSAGE      = C.GTestLogType(C.G_TEST_LOG_MESSAGE)
	TEST_LOG_START_SUITE  = C.GTestLogType(C.G_TEST_LOG_START_SUITE)
	TEST_LOG_STOP_SUITE   = C.GTestLogType(C.G_TEST_LOG_STOP_SUITE)

	// ThreadError
	THREAD_ERROR_AGAIN = C.GThreadError(C.G_THREAD_ERROR_AGAIN)

	// TimeType
	TIME_TYPE_STANDARD  = C.GTimeType(C.G_TIME_TYPE_STANDARD)
	TIME_TYPE_DAYLIGHT  = C.GTimeType(C.G_TIME_TYPE_DAYLIGHT)
	TIME_TYPE_UNIVERSAL = C.GTimeType(C.G_TIME_TYPE_UNIVERSAL)

	// TokenType
	TOKEN_EOF             = C.GTokenType(C.G_TOKEN_EOF)
	TOKEN_LEFT_PAREN      = C.GTokenType(C.G_TOKEN_LEFT_PAREN)
	TOKEN_RIGHT_PAREN     = C.GTokenType(C.G_TOKEN_RIGHT_PAREN)
	TOKEN_LEFT_CURLY      = C.GTokenType(C.G_TOKEN_LEFT_CURLY)
	TOKEN_RIGHT_CURLY     = C.GTokenType(C.G_TOKEN_RIGHT_CURLY)
	TOKEN_LEFT_BRACE      = C.GTokenType(C.G_TOKEN_LEFT_BRACE)
	TOKEN_RIGHT_BRACE     = C.GTokenType(C.G_TOKEN_RIGHT_BRACE)
	TOKEN_EQUAL_SIGN      = C.GTokenType(C.G_TOKEN_EQUAL_SIGN)
	TOKEN_COMMA           = C.GTokenType(C.G_TOKEN_COMMA)
	TOKEN_NONE            = C.GTokenType(C.G_TOKEN_NONE)
	TOKEN_ERROR           = C.GTokenType(C.G_TOKEN_ERROR)
	TOKEN_CHAR            = C.GTokenType(C.G_TOKEN_CHAR)
	TOKEN_BINARY          = C.GTokenType(C.G_TOKEN_BINARY)
	TOKEN_OCTAL           = C.GTokenType(C.G_TOKEN_OCTAL)
	TOKEN_INT             = C.GTokenType(C.G_TOKEN_INT)
	TOKEN_HEX             = C.GTokenType(C.G_TOKEN_HEX)
	TOKEN_FLOAT           = C.GTokenType(C.G_TOKEN_FLOAT)
	TOKEN_STRING          = C.GTokenType(C.G_TOKEN_STRING)
	TOKEN_SYMBOL          = C.GTokenType(C.G_TOKEN_SYMBOL)
	TOKEN_IDENTIFIER      = C.GTokenType(C.G_TOKEN_IDENTIFIER)
	TOKEN_IDENTIFIER_NULL = C.GTokenType(C.G_TOKEN_IDENTIFIER_NULL)
	TOKEN_COMMENT_SINGLE  = C.GTokenType(C.G_TOKEN_COMMENT_SINGLE)
	TOKEN_COMMENT_MULTI   = C.GTokenType(C.G_TOKEN_COMMENT_MULTI)

	// TraverseType
	IN_ORDER    = C.GTraverseType(C.G_IN_ORDER)
	PRE_ORDER   = C.GTraverseType(C.G_PRE_ORDER)
	POST_ORDER  = C.GTraverseType(C.G_POST_ORDER)
	LEVEL_ORDER = C.GTraverseType(C.G_LEVEL_ORDER)

	// UnicodeBreakType
	UNICODE_BREAK_MANDATORY                    = C.GUnicodeBreakType(C.G_UNICODE_BREAK_MANDATORY)
	UNICODE_BREAK_CARRIAGE_RETURN              = C.GUnicodeBreakType(C.G_UNICODE_BREAK_CARRIAGE_RETURN)
	UNICODE_BREAK_LINE_FEED                    = C.GUnicodeBreakType(C.G_UNICODE_BREAK_LINE_FEED)
	UNICODE_BREAK_COMBINING_MARK               = C.GUnicodeBreakType(C.G_UNICODE_BREAK_COMBINING_MARK)
	UNICODE_BREAK_SURROGATE                    = C.GUnicodeBreakType(C.G_UNICODE_BREAK_SURROGATE)
	UNICODE_BREAK_ZERO_WIDTH_SPACE             = C.GUnicodeBreakType(C.G_UNICODE_BREAK_ZERO_WIDTH_SPACE)
	UNICODE_BREAK_INSEPARABLE                  = C.GUnicodeBreakType(C.G_UNICODE_BREAK_INSEPARABLE)
	UNICODE_BREAK_NON_BREAKING_GLUE            = C.GUnicodeBreakType(C.G_UNICODE_BREAK_NON_BREAKING_GLUE)
	UNICODE_BREAK_CONTINGENT                   = C.GUnicodeBreakType(C.G_UNICODE_BREAK_CONTINGENT)
	UNICODE_BREAK_SPACE                        = C.GUnicodeBreakType(C.G_UNICODE_BREAK_SPACE)
	UNICODE_BREAK_AFTER                        = C.GUnicodeBreakType(C.G_UNICODE_BREAK_AFTER)
	UNICODE_BREAK_BEFORE                       = C.GUnicodeBreakType(C.G_UNICODE_BREAK_BEFORE)
	UNICODE_BREAK_BEFORE_AND_AFTER             = C.GUnicodeBreakType(C.G_UNICODE_BREAK_BEFORE_AND_AFTER)
	UNICODE_BREAK_HYPHEN                       = C.GUnicodeBreakType(C.G_UNICODE_BREAK_HYPHEN)
	UNICODE_BREAK_NON_STARTER                  = C.GUnicodeBreakType(C.G_UNICODE_BREAK_NON_STARTER)
	UNICODE_BREAK_OPEN_PUNCTUATION             = C.GUnicodeBreakType(C.G_UNICODE_BREAK_OPEN_PUNCTUATION)
	UNICODE_BREAK_CLOSE_PUNCTUATION            = C.GUnicodeBreakType(C.G_UNICODE_BREAK_CLOSE_PUNCTUATION)
	UNICODE_BREAK_QUOTATION                    = C.GUnicodeBreakType(C.G_UNICODE_BREAK_QUOTATION)
	UNICODE_BREAK_EXCLAMATION                  = C.GUnicodeBreakType(C.G_UNICODE_BREAK_EXCLAMATION)
	UNICODE_BREAK_IDEOGRAPHIC                  = C.GUnicodeBreakType(C.G_UNICODE_BREAK_IDEOGRAPHIC)
	UNICODE_BREAK_NUMERIC                      = C.GUnicodeBreakType(C.G_UNICODE_BREAK_NUMERIC)
	UNICODE_BREAK_INFIX_SEPARATOR              = C.GUnicodeBreakType(C.G_UNICODE_BREAK_INFIX_SEPARATOR)
	UNICODE_BREAK_SYMBOL                       = C.GUnicodeBreakType(C.G_UNICODE_BREAK_SYMBOL)
	UNICODE_BREAK_ALPHABETIC                   = C.GUnicodeBreakType(C.G_UNICODE_BREAK_ALPHABETIC)
	UNICODE_BREAK_PREFIX                       = C.GUnicodeBreakType(C.G_UNICODE_BREAK_PREFIX)
	UNICODE_BREAK_POSTFIX                      = C.GUnicodeBreakType(C.G_UNICODE_BREAK_POSTFIX)
	UNICODE_BREAK_COMPLEX_CONTEXT              = C.GUnicodeBreakType(C.G_UNICODE_BREAK_COMPLEX_CONTEXT)
	UNICODE_BREAK_AMBIGUOUS                    = C.GUnicodeBreakType(C.G_UNICODE_BREAK_AMBIGUOUS)
	UNICODE_BREAK_UNKNOWN                      = C.GUnicodeBreakType(C.G_UNICODE_BREAK_UNKNOWN)
	UNICODE_BREAK_NEXT_LINE                    = C.GUnicodeBreakType(C.G_UNICODE_BREAK_NEXT_LINE)
	UNICODE_BREAK_WORD_JOINER                  = C.GUnicodeBreakType(C.G_UNICODE_BREAK_WORD_JOINER)
	UNICODE_BREAK_HANGUL_L_JAMO                = C.GUnicodeBreakType(C.G_UNICODE_BREAK_HANGUL_L_JAMO)
	UNICODE_BREAK_HANGUL_V_JAMO                = C.GUnicodeBreakType(C.G_UNICODE_BREAK_HANGUL_V_JAMO)
	UNICODE_BREAK_HANGUL_T_JAMO                = C.GUnicodeBreakType(C.G_UNICODE_BREAK_HANGUL_T_JAMO)
	UNICODE_BREAK_HANGUL_LV_SYLLABLE           = C.GUnicodeBreakType(C.G_UNICODE_BREAK_HANGUL_LV_SYLLABLE)
	UNICODE_BREAK_HANGUL_LVT_SYLLABLE          = C.GUnicodeBreakType(C.G_UNICODE_BREAK_HANGUL_LVT_SYLLABLE)
	UNICODE_BREAK_CLOSE_PARANTHESIS            = C.GUnicodeBreakType(C.G_UNICODE_BREAK_CLOSE_PARANTHESIS)
	UNICODE_BREAK_CONDITIONAL_JAPANESE_STARTER = C.GUnicodeBreakType(C.G_UNICODE_BREAK_CONDITIONAL_JAPANESE_STARTER)
	UNICODE_BREAK_HEBREW_LETTER                = C.GUnicodeBreakType(C.G_UNICODE_BREAK_HEBREW_LETTER)
	UNICODE_BREAK_REGIONAL_INDICATOR           = C.GUnicodeBreakType(C.G_UNICODE_BREAK_REGIONAL_INDICATOR)

	// UnicodeScript
	UNICODE_SCRIPT_INVALID_CODE           = C.GUnicodeScript(C.G_UNICODE_SCRIPT_INVALID_CODE)
	UNICODE_SCRIPT_COMMON                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_COMMON)
	UNICODE_SCRIPT_INHERITED              = C.GUnicodeScript(C.G_UNICODE_SCRIPT_INHERITED)
	UNICODE_SCRIPT_ARABIC                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_ARABIC)
	UNICODE_SCRIPT_ARMENIAN               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_ARMENIAN)
	UNICODE_SCRIPT_BENGALI                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BENGALI)
	UNICODE_SCRIPT_BOPOMOFO               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BOPOMOFO)
	UNICODE_SCRIPT_CHEROKEE               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_CHEROKEE)
	UNICODE_SCRIPT_COPTIC                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_COPTIC)
	UNICODE_SCRIPT_CYRILLIC               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_CYRILLIC)
	UNICODE_SCRIPT_DESERET                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_DESERET)
	UNICODE_SCRIPT_DEVANAGARI             = C.GUnicodeScript(C.G_UNICODE_SCRIPT_DEVANAGARI)
	UNICODE_SCRIPT_ETHIOPIC               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_ETHIOPIC)
	UNICODE_SCRIPT_GEORGIAN               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_GEORGIAN)
	UNICODE_SCRIPT_GOTHIC                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_GOTHIC)
	UNICODE_SCRIPT_GREEK                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_GREEK)
	UNICODE_SCRIPT_GUJARATI               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_GUJARATI)
	UNICODE_SCRIPT_GURMUKHI               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_GURMUKHI)
	UNICODE_SCRIPT_HAN                    = C.GUnicodeScript(C.G_UNICODE_SCRIPT_HAN)
	UNICODE_SCRIPT_HANGUL                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_HANGUL)
	UNICODE_SCRIPT_HEBREW                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_HEBREW)
	UNICODE_SCRIPT_HIRAGANA               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_HIRAGANA)
	UNICODE_SCRIPT_KANNADA                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_KANNADA)
	UNICODE_SCRIPT_KATAKANA               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_KATAKANA)
	UNICODE_SCRIPT_KHMER                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_KHMER)
	UNICODE_SCRIPT_LAO                    = C.GUnicodeScript(C.G_UNICODE_SCRIPT_LAO)
	UNICODE_SCRIPT_LATIN                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_LATIN)
	UNICODE_SCRIPT_MALAYALAM              = C.GUnicodeScript(C.G_UNICODE_SCRIPT_MALAYALAM)
	UNICODE_SCRIPT_MONGOLIAN              = C.GUnicodeScript(C.G_UNICODE_SCRIPT_MONGOLIAN)
	UNICODE_SCRIPT_MYANMAR                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_MYANMAR)
	UNICODE_SCRIPT_OGHAM                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_OGHAM)
	UNICODE_SCRIPT_OLD_ITALIC             = C.GUnicodeScript(C.G_UNICODE_SCRIPT_OLD_ITALIC)
	UNICODE_SCRIPT_ORIYA                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_ORIYA)
	UNICODE_SCRIPT_RUNIC                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_RUNIC)
	UNICODE_SCRIPT_SINHALA                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SINHALA)
	UNICODE_SCRIPT_SYRIAC                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SYRIAC)
	UNICODE_SCRIPT_TAMIL                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TAMIL)
	UNICODE_SCRIPT_TELUGU                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TELUGU)
	UNICODE_SCRIPT_THAANA                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_THAANA)
	UNICODE_SCRIPT_THAI                   = C.GUnicodeScript(C.G_UNICODE_SCRIPT_THAI)
	UNICODE_SCRIPT_TIBETAN                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TIBETAN)
	UNICODE_SCRIPT_CANADIAN_ABORIGINAL    = C.GUnicodeScript(C.G_UNICODE_SCRIPT_CANADIAN_ABORIGINAL)
	UNICODE_SCRIPT_YI                     = C.GUnicodeScript(C.G_UNICODE_SCRIPT_YI)
	UNICODE_SCRIPT_TAGALOG                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TAGALOG)
	UNICODE_SCRIPT_HANUNOO                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_HANUNOO)
	UNICODE_SCRIPT_BUHID                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BUHID)
	UNICODE_SCRIPT_TAGBANWA               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TAGBANWA)
	UNICODE_SCRIPT_BRAILLE                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BRAILLE)
	UNICODE_SCRIPT_CYPRIOT                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_CYPRIOT)
	UNICODE_SCRIPT_LIMBU                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_LIMBU)
	UNICODE_SCRIPT_OSMANYA                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_OSMANYA)
	UNICODE_SCRIPT_SHAVIAN                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SHAVIAN)
	UNICODE_SCRIPT_LINEAR_B               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_LINEAR_B)
	UNICODE_SCRIPT_TAI_LE                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TAI_LE)
	UNICODE_SCRIPT_UGARITIC               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_UGARITIC)
	UNICODE_SCRIPT_NEW_TAI_LUE            = C.GUnicodeScript(C.G_UNICODE_SCRIPT_NEW_TAI_LUE)
	UNICODE_SCRIPT_BUGINESE               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BUGINESE)
	UNICODE_SCRIPT_GLAGOLITIC             = C.GUnicodeScript(C.G_UNICODE_SCRIPT_GLAGOLITIC)
	UNICODE_SCRIPT_TIFINAGH               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TIFINAGH)
	UNICODE_SCRIPT_SYLOTI_NAGRI           = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SYLOTI_NAGRI)
	UNICODE_SCRIPT_OLD_PERSIAN            = C.GUnicodeScript(C.G_UNICODE_SCRIPT_OLD_PERSIAN)
	UNICODE_SCRIPT_KHAROSHTHI             = C.GUnicodeScript(C.G_UNICODE_SCRIPT_KHAROSHTHI)
	UNICODE_SCRIPT_UNKNOWN                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_UNKNOWN)
	UNICODE_SCRIPT_BALINESE               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BALINESE)
	UNICODE_SCRIPT_CUNEIFORM              = C.GUnicodeScript(C.G_UNICODE_SCRIPT_CUNEIFORM)
	UNICODE_SCRIPT_PHOENICIAN             = C.GUnicodeScript(C.G_UNICODE_SCRIPT_PHOENICIAN)
	UNICODE_SCRIPT_PHAGS_PA               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_PHAGS_PA)
	UNICODE_SCRIPT_NKO                    = C.GUnicodeScript(C.G_UNICODE_SCRIPT_NKO)
	UNICODE_SCRIPT_KAYAH_LI               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_KAYAH_LI)
	UNICODE_SCRIPT_LEPCHA                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_LEPCHA)
	UNICODE_SCRIPT_REJANG                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_REJANG)
	UNICODE_SCRIPT_SUNDANESE              = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SUNDANESE)
	UNICODE_SCRIPT_SAURASHTRA             = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SAURASHTRA)
	UNICODE_SCRIPT_CHAM                   = C.GUnicodeScript(C.G_UNICODE_SCRIPT_CHAM)
	UNICODE_SCRIPT_OL_CHIKI               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_OL_CHIKI)
	UNICODE_SCRIPT_VAI                    = C.GUnicodeScript(C.G_UNICODE_SCRIPT_VAI)
	UNICODE_SCRIPT_CARIAN                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_CARIAN)
	UNICODE_SCRIPT_LYCIAN                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_LYCIAN)
	UNICODE_SCRIPT_LYDIAN                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_LYDIAN)
	UNICODE_SCRIPT_AVESTAN                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_AVESTAN)
	UNICODE_SCRIPT_BAMUM                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BAMUM)
	UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS   = C.GUnicodeScript(C.G_UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS)
	UNICODE_SCRIPT_IMPERIAL_ARAMAIC       = C.GUnicodeScript(C.G_UNICODE_SCRIPT_IMPERIAL_ARAMAIC)
	UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI)
	UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN = C.GUnicodeScript(C.G_UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN)
	UNICODE_SCRIPT_JAVANESE               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_JAVANESE)
	UNICODE_SCRIPT_KAITHI                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_KAITHI)
	UNICODE_SCRIPT_LISU                   = C.GUnicodeScript(C.G_UNICODE_SCRIPT_LISU)
	UNICODE_SCRIPT_MEETEI_MAYEK           = C.GUnicodeScript(C.G_UNICODE_SCRIPT_MEETEI_MAYEK)
	UNICODE_SCRIPT_OLD_SOUTH_ARABIAN      = C.GUnicodeScript(C.G_UNICODE_SCRIPT_OLD_SOUTH_ARABIAN)
	UNICODE_SCRIPT_OLD_TURKIC             = C.GUnicodeScript(C.G_UNICODE_SCRIPT_OLD_TURKIC)
	UNICODE_SCRIPT_SAMARITAN              = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SAMARITAN)
	UNICODE_SCRIPT_TAI_THAM               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TAI_THAM)
	UNICODE_SCRIPT_TAI_VIET               = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TAI_VIET)
	UNICODE_SCRIPT_BATAK                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BATAK)
	UNICODE_SCRIPT_BRAHMI                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_BRAHMI)
	UNICODE_SCRIPT_MANDAIC                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_MANDAIC)
	UNICODE_SCRIPT_CHAKMA                 = C.GUnicodeScript(C.G_UNICODE_SCRIPT_CHAKMA)
	UNICODE_SCRIPT_MEROITIC_CURSIVE       = C.GUnicodeScript(C.G_UNICODE_SCRIPT_MEROITIC_CURSIVE)
	UNICODE_SCRIPT_MEROITIC_HIEROGLYPHS   = C.GUnicodeScript(C.G_UNICODE_SCRIPT_MEROITIC_HIEROGLYPHS)
	UNICODE_SCRIPT_MIAO                   = C.GUnicodeScript(C.G_UNICODE_SCRIPT_MIAO)
	UNICODE_SCRIPT_SHARADA                = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SHARADA)
	UNICODE_SCRIPT_SORA_SOMPENG           = C.GUnicodeScript(C.G_UNICODE_SCRIPT_SORA_SOMPENG)
	UNICODE_SCRIPT_TAKRI                  = C.GUnicodeScript(C.G_UNICODE_SCRIPT_TAKRI)

	// UnicodeType
	UNICODE_CONTROL             = C.GUnicodeType(C.G_UNICODE_CONTROL)
	UNICODE_FORMAT              = C.GUnicodeType(C.G_UNICODE_FORMAT)
	UNICODE_UNASSIGNED          = C.GUnicodeType(C.G_UNICODE_UNASSIGNED)
	UNICODE_PRIVATE_USE         = C.GUnicodeType(C.G_UNICODE_PRIVATE_USE)
	UNICODE_SURROGATE           = C.GUnicodeType(C.G_UNICODE_SURROGATE)
	UNICODE_LOWERCASE_LETTER    = C.GUnicodeType(C.G_UNICODE_LOWERCASE_LETTER)
	UNICODE_MODIFIER_LETTER     = C.GUnicodeType(C.G_UNICODE_MODIFIER_LETTER)
	UNICODE_OTHER_LETTER        = C.GUnicodeType(C.G_UNICODE_OTHER_LETTER)
	UNICODE_TITLECASE_LETTER    = C.GUnicodeType(C.G_UNICODE_TITLECASE_LETTER)
	UNICODE_UPPERCASE_LETTER    = C.GUnicodeType(C.G_UNICODE_UPPERCASE_LETTER)
	UNICODE_SPACING_MARK        = C.GUnicodeType(C.G_UNICODE_SPACING_MARK)
	UNICODE_ENCLOSING_MARK      = C.GUnicodeType(C.G_UNICODE_ENCLOSING_MARK)
	UNICODE_NON_SPACING_MARK    = C.GUnicodeType(C.G_UNICODE_NON_SPACING_MARK)
	UNICODE_DECIMAL_NUMBER      = C.GUnicodeType(C.G_UNICODE_DECIMAL_NUMBER)
	UNICODE_LETTER_NUMBER       = C.GUnicodeType(C.G_UNICODE_LETTER_NUMBER)
	UNICODE_OTHER_NUMBER        = C.GUnicodeType(C.G_UNICODE_OTHER_NUMBER)
	UNICODE_CONNECT_PUNCTUATION = C.GUnicodeType(C.G_UNICODE_CONNECT_PUNCTUATION)
	UNICODE_DASH_PUNCTUATION    = C.GUnicodeType(C.G_UNICODE_DASH_PUNCTUATION)
	UNICODE_CLOSE_PUNCTUATION   = C.GUnicodeType(C.G_UNICODE_CLOSE_PUNCTUATION)
	UNICODE_FINAL_PUNCTUATION   = C.GUnicodeType(C.G_UNICODE_FINAL_PUNCTUATION)
	UNICODE_INITIAL_PUNCTUATION = C.GUnicodeType(C.G_UNICODE_INITIAL_PUNCTUATION)
	UNICODE_OTHER_PUNCTUATION   = C.GUnicodeType(C.G_UNICODE_OTHER_PUNCTUATION)
	UNICODE_OPEN_PUNCTUATION    = C.GUnicodeType(C.G_UNICODE_OPEN_PUNCTUATION)
	UNICODE_CURRENCY_SYMBOL     = C.GUnicodeType(C.G_UNICODE_CURRENCY_SYMBOL)
	UNICODE_MODIFIER_SYMBOL     = C.GUnicodeType(C.G_UNICODE_MODIFIER_SYMBOL)
	UNICODE_MATH_SYMBOL         = C.GUnicodeType(C.G_UNICODE_MATH_SYMBOL)
	UNICODE_OTHER_SYMBOL        = C.GUnicodeType(C.G_UNICODE_OTHER_SYMBOL)
	UNICODE_LINE_SEPARATOR      = C.GUnicodeType(C.G_UNICODE_LINE_SEPARATOR)
	UNICODE_PARAGRAPH_SEPARATOR = C.GUnicodeType(C.G_UNICODE_PARAGRAPH_SEPARATOR)
	UNICODE_SPACE_SEPARATOR     = C.GUnicodeType(C.G_UNICODE_SPACE_SEPARATOR)

	// UserDirectory
	USER_DIRECTORY_DESKTOP      = C.GUserDirectory(C.G_USER_DIRECTORY_DESKTOP)
	USER_DIRECTORY_DOCUMENTS    = C.GUserDirectory(C.G_USER_DIRECTORY_DOCUMENTS)
	USER_DIRECTORY_DOWNLOAD     = C.GUserDirectory(C.G_USER_DIRECTORY_DOWNLOAD)
	USER_DIRECTORY_MUSIC        = C.GUserDirectory(C.G_USER_DIRECTORY_MUSIC)
	USER_DIRECTORY_PICTURES     = C.GUserDirectory(C.G_USER_DIRECTORY_PICTURES)
	USER_DIRECTORY_PUBLIC_SHARE = C.GUserDirectory(C.G_USER_DIRECTORY_PUBLIC_SHARE)
	USER_DIRECTORY_TEMPLATES    = C.GUserDirectory(C.G_USER_DIRECTORY_TEMPLATES)
	USER_DIRECTORY_VIDEOS       = C.GUserDirectory(C.G_USER_DIRECTORY_VIDEOS)
	USER_N_DIRECTORIES          = C.GUserDirectory(C.G_USER_N_DIRECTORIES)

	// VariantClass
	VARIANT_CLASS_BOOLEAN     = C.GVariantClass(C.G_VARIANT_CLASS_BOOLEAN)
	VARIANT_CLASS_BYTE        = C.GVariantClass(C.G_VARIANT_CLASS_BYTE)
	VARIANT_CLASS_INT16       = C.GVariantClass(C.G_VARIANT_CLASS_INT16)
	VARIANT_CLASS_UINT16      = C.GVariantClass(C.G_VARIANT_CLASS_UINT16)
	VARIANT_CLASS_INT32       = C.GVariantClass(C.G_VARIANT_CLASS_INT32)
	VARIANT_CLASS_UINT32      = C.GVariantClass(C.G_VARIANT_CLASS_UINT32)
	VARIANT_CLASS_INT64       = C.GVariantClass(C.G_VARIANT_CLASS_INT64)
	VARIANT_CLASS_UINT64      = C.GVariantClass(C.G_VARIANT_CLASS_UINT64)
	VARIANT_CLASS_HANDLE      = C.GVariantClass(C.G_VARIANT_CLASS_HANDLE)
	VARIANT_CLASS_DOUBLE      = C.GVariantClass(C.G_VARIANT_CLASS_DOUBLE)
	VARIANT_CLASS_STRING      = C.GVariantClass(C.G_VARIANT_CLASS_STRING)
	VARIANT_CLASS_OBJECT_PATH = C.GVariantClass(C.G_VARIANT_CLASS_OBJECT_PATH)
	VARIANT_CLASS_SIGNATURE   = C.GVariantClass(C.G_VARIANT_CLASS_SIGNATURE)
	VARIANT_CLASS_VARIANT     = C.GVariantClass(C.G_VARIANT_CLASS_VARIANT)
	VARIANT_CLASS_MAYBE       = C.GVariantClass(C.G_VARIANT_CLASS_MAYBE)
	VARIANT_CLASS_ARRAY       = C.GVariantClass(C.G_VARIANT_CLASS_ARRAY)
	VARIANT_CLASS_TUPLE       = C.GVariantClass(C.G_VARIANT_CLASS_TUPLE)
	VARIANT_CLASS_DICT_ENTRY  = C.GVariantClass(C.G_VARIANT_CLASS_DICT_ENTRY)

	// VariantParseError
	VARIANT_PARSE_ERROR_FAILED                       = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_FAILED)
	VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED          = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED)
	VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE            = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE)
	VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED       = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED)
	VARIANT_PARSE_ERROR_INPUT_NOT_AT_END             = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_INPUT_NOT_AT_END)
	VARIANT_PARSE_ERROR_INVALID_CHARACTER            = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_INVALID_CHARACTER)
	VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING        = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING)
	VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH          = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH)
	VARIANT_PARSE_ERROR_INVALID_SIGNATURE            = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_INVALID_SIGNATURE)
	VARIANT_PARSE_ERROR_INVALID_TYPE_STRING          = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_INVALID_TYPE_STRING)
	VARIANT_PARSE_ERROR_NO_COMMON_TYPE               = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_NO_COMMON_TYPE)
	VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE          = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE)
	VARIANT_PARSE_ERROR_NUMBER_TOO_BIG               = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_NUMBER_TOO_BIG)
	VARIANT_PARSE_ERROR_TYPE_ERROR                   = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_TYPE_ERROR)
	VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN             = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN)
	VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD              = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD)
	VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT)
	VARIANT_PARSE_ERROR_VALUE_EXPECTED               = C.GVariantParseError(C.G_VARIANT_PARSE_ERROR_VALUE_EXPECTED)
)
