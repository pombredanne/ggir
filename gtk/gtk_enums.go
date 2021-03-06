package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <glib.h>
#include <stdlib.h>
*/
import "C"

var (
	// Align
	ALIGN_FILL     = C.GtkAlign(C.GTK_ALIGN_FILL)
	ALIGN_START    = C.GtkAlign(C.GTK_ALIGN_START)
	ALIGN_END      = C.GtkAlign(C.GTK_ALIGN_END)
	ALIGN_CENTER   = C.GtkAlign(C.GTK_ALIGN_CENTER)
	ALIGN_BASELINE = C.GtkAlign(C.GTK_ALIGN_BASELINE)

	// ArrowPlacement
	ARROWS_BOTH  = C.GtkArrowPlacement(C.GTK_ARROWS_BOTH)
	ARROWS_START = C.GtkArrowPlacement(C.GTK_ARROWS_START)
	ARROWS_END   = C.GtkArrowPlacement(C.GTK_ARROWS_END)

	// ArrowType
	ARROW_UP    = C.GtkArrowType(C.GTK_ARROW_UP)
	ARROW_DOWN  = C.GtkArrowType(C.GTK_ARROW_DOWN)
	ARROW_LEFT  = C.GtkArrowType(C.GTK_ARROW_LEFT)
	ARROW_RIGHT = C.GtkArrowType(C.GTK_ARROW_RIGHT)
	ARROW_NONE  = C.GtkArrowType(C.GTK_ARROW_NONE)

	// AssistantPageType
	ASSISTANT_PAGE_CONTENT  = C.GtkAssistantPageType(C.GTK_ASSISTANT_PAGE_CONTENT)
	ASSISTANT_PAGE_INTRO    = C.GtkAssistantPageType(C.GTK_ASSISTANT_PAGE_INTRO)
	ASSISTANT_PAGE_CONFIRM  = C.GtkAssistantPageType(C.GTK_ASSISTANT_PAGE_CONFIRM)
	ASSISTANT_PAGE_SUMMARY  = C.GtkAssistantPageType(C.GTK_ASSISTANT_PAGE_SUMMARY)
	ASSISTANT_PAGE_PROGRESS = C.GtkAssistantPageType(C.GTK_ASSISTANT_PAGE_PROGRESS)
	ASSISTANT_PAGE_CUSTOM   = C.GtkAssistantPageType(C.GTK_ASSISTANT_PAGE_CUSTOM)

	// BaselinePosition
	BASELINE_POSITION_TOP    = C.GtkBaselinePosition(C.GTK_BASELINE_POSITION_TOP)
	BASELINE_POSITION_CENTER = C.GtkBaselinePosition(C.GTK_BASELINE_POSITION_CENTER)
	BASELINE_POSITION_BOTTOM = C.GtkBaselinePosition(C.GTK_BASELINE_POSITION_BOTTOM)

	// BorderStyle
	BORDER_STYLE_NONE   = C.GtkBorderStyle(C.GTK_BORDER_STYLE_NONE)
	BORDER_STYLE_SOLID  = C.GtkBorderStyle(C.GTK_BORDER_STYLE_SOLID)
	BORDER_STYLE_INSET  = C.GtkBorderStyle(C.GTK_BORDER_STYLE_INSET)
	BORDER_STYLE_OUTSET = C.GtkBorderStyle(C.GTK_BORDER_STYLE_OUTSET)
	BORDER_STYLE_HIDDEN = C.GtkBorderStyle(C.GTK_BORDER_STYLE_HIDDEN)
	BORDER_STYLE_DOTTED = C.GtkBorderStyle(C.GTK_BORDER_STYLE_DOTTED)
	BORDER_STYLE_DASHED = C.GtkBorderStyle(C.GTK_BORDER_STYLE_DASHED)
	BORDER_STYLE_DOUBLE = C.GtkBorderStyle(C.GTK_BORDER_STYLE_DOUBLE)
	BORDER_STYLE_GROOVE = C.GtkBorderStyle(C.GTK_BORDER_STYLE_GROOVE)
	BORDER_STYLE_RIDGE  = C.GtkBorderStyle(C.GTK_BORDER_STYLE_RIDGE)

	// BuilderError
	BUILDER_ERROR_INVALID_TYPE_FUNCTION  = C.GtkBuilderError(C.GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION)
	BUILDER_ERROR_UNHANDLED_TAG          = C.GtkBuilderError(C.GTK_BUILDER_ERROR_UNHANDLED_TAG)
	BUILDER_ERROR_MISSING_ATTRIBUTE      = C.GtkBuilderError(C.GTK_BUILDER_ERROR_MISSING_ATTRIBUTE)
	BUILDER_ERROR_INVALID_ATTRIBUTE      = C.GtkBuilderError(C.GTK_BUILDER_ERROR_INVALID_ATTRIBUTE)
	BUILDER_ERROR_INVALID_TAG            = C.GtkBuilderError(C.GTK_BUILDER_ERROR_INVALID_TAG)
	BUILDER_ERROR_MISSING_PROPERTY_VALUE = C.GtkBuilderError(C.GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE)
	BUILDER_ERROR_INVALID_VALUE          = C.GtkBuilderError(C.GTK_BUILDER_ERROR_INVALID_VALUE)
	BUILDER_ERROR_VERSION_MISMATCH       = C.GtkBuilderError(C.GTK_BUILDER_ERROR_VERSION_MISMATCH)
	BUILDER_ERROR_DUPLICATE_ID           = C.GtkBuilderError(C.GTK_BUILDER_ERROR_DUPLICATE_ID)
	BUILDER_ERROR_OBJECT_TYPE_REFUSED    = C.GtkBuilderError(C.GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED)
	BUILDER_ERROR_TEMPLATE_MISMATCH      = C.GtkBuilderError(C.GTK_BUILDER_ERROR_TEMPLATE_MISMATCH)
	BUILDER_ERROR_INVALID_PROPERTY       = C.GtkBuilderError(C.GTK_BUILDER_ERROR_INVALID_PROPERTY)
	BUILDER_ERROR_INVALID_SIGNAL         = C.GtkBuilderError(C.GTK_BUILDER_ERROR_INVALID_SIGNAL)

	// ButtonBoxStyle
	BUTTONBOX_SPREAD = C.GtkButtonBoxStyle(C.GTK_BUTTONBOX_SPREAD)
	BUTTONBOX_EDGE   = C.GtkButtonBoxStyle(C.GTK_BUTTONBOX_EDGE)
	BUTTONBOX_START  = C.GtkButtonBoxStyle(C.GTK_BUTTONBOX_START)
	BUTTONBOX_END    = C.GtkButtonBoxStyle(C.GTK_BUTTONBOX_END)
	BUTTONBOX_CENTER = C.GtkButtonBoxStyle(C.GTK_BUTTONBOX_CENTER)
	BUTTONBOX_EXPAND = C.GtkButtonBoxStyle(C.GTK_BUTTONBOX_EXPAND)

	// ButtonsType
	BUTTONS_NONE      = C.GtkButtonsType(C.GTK_BUTTONS_NONE)
	BUTTONS_OK        = C.GtkButtonsType(C.GTK_BUTTONS_OK)
	BUTTONS_CLOSE     = C.GtkButtonsType(C.GTK_BUTTONS_CLOSE)
	BUTTONS_CANCEL    = C.GtkButtonsType(C.GTK_BUTTONS_CANCEL)
	BUTTONS_YES_NO    = C.GtkButtonsType(C.GTK_BUTTONS_YES_NO)
	BUTTONS_OK_CANCEL = C.GtkButtonsType(C.GTK_BUTTONS_OK_CANCEL)

	// CellRendererAccelMode
	CELL_RENDERER_ACCEL_MODE_GTK   = C.GtkCellRendererAccelMode(C.GTK_CELL_RENDERER_ACCEL_MODE_GTK)
	CELL_RENDERER_ACCEL_MODE_OTHER = C.GtkCellRendererAccelMode(C.GTK_CELL_RENDERER_ACCEL_MODE_OTHER)

	// CellRendererMode
	CELL_RENDERER_MODE_INERT       = C.GtkCellRendererMode(C.GTK_CELL_RENDERER_MODE_INERT)
	CELL_RENDERER_MODE_ACTIVATABLE = C.GtkCellRendererMode(C.GTK_CELL_RENDERER_MODE_ACTIVATABLE)
	CELL_RENDERER_MODE_EDITABLE    = C.GtkCellRendererMode(C.GTK_CELL_RENDERER_MODE_EDITABLE)

	// CornerType
	CORNER_TOP_LEFT     = C.GtkCornerType(C.GTK_CORNER_TOP_LEFT)
	CORNER_BOTTOM_LEFT  = C.GtkCornerType(C.GTK_CORNER_BOTTOM_LEFT)
	CORNER_TOP_RIGHT    = C.GtkCornerType(C.GTK_CORNER_TOP_RIGHT)
	CORNER_BOTTOM_RIGHT = C.GtkCornerType(C.GTK_CORNER_BOTTOM_RIGHT)

	// CssProviderError
	CSS_PROVIDER_ERROR_FAILED        = C.GtkCssProviderError(C.GTK_CSS_PROVIDER_ERROR_FAILED)
	CSS_PROVIDER_ERROR_SYNTAX        = C.GtkCssProviderError(C.GTK_CSS_PROVIDER_ERROR_SYNTAX)
	CSS_PROVIDER_ERROR_IMPORT        = C.GtkCssProviderError(C.GTK_CSS_PROVIDER_ERROR_IMPORT)
	CSS_PROVIDER_ERROR_NAME          = C.GtkCssProviderError(C.GTK_CSS_PROVIDER_ERROR_NAME)
	CSS_PROVIDER_ERROR_DEPRECATED    = C.GtkCssProviderError(C.GTK_CSS_PROVIDER_ERROR_DEPRECATED)
	CSS_PROVIDER_ERROR_UNKNOWN_VALUE = C.GtkCssProviderError(C.GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE)

	// CssSectionType
	CSS_SECTION_DOCUMENT         = C.GtkCssSectionType(C.GTK_CSS_SECTION_DOCUMENT)
	CSS_SECTION_IMPORT           = C.GtkCssSectionType(C.GTK_CSS_SECTION_IMPORT)
	CSS_SECTION_COLOR_DEFINITION = C.GtkCssSectionType(C.GTK_CSS_SECTION_COLOR_DEFINITION)
	CSS_SECTION_BINDING_SET      = C.GtkCssSectionType(C.GTK_CSS_SECTION_BINDING_SET)
	CSS_SECTION_RULESET          = C.GtkCssSectionType(C.GTK_CSS_SECTION_RULESET)
	CSS_SECTION_SELECTOR         = C.GtkCssSectionType(C.GTK_CSS_SECTION_SELECTOR)
	CSS_SECTION_DECLARATION      = C.GtkCssSectionType(C.GTK_CSS_SECTION_DECLARATION)
	CSS_SECTION_VALUE            = C.GtkCssSectionType(C.GTK_CSS_SECTION_VALUE)
	CSS_SECTION_KEYFRAMES        = C.GtkCssSectionType(C.GTK_CSS_SECTION_KEYFRAMES)

	// DeleteType
	DELETE_CHARS             = C.GtkDeleteType(C.GTK_DELETE_CHARS)
	DELETE_WORD_ENDS         = C.GtkDeleteType(C.GTK_DELETE_WORD_ENDS)
	DELETE_WORDS             = C.GtkDeleteType(C.GTK_DELETE_WORDS)
	DELETE_DISPLAY_LINES     = C.GtkDeleteType(C.GTK_DELETE_DISPLAY_LINES)
	DELETE_DISPLAY_LINE_ENDS = C.GtkDeleteType(C.GTK_DELETE_DISPLAY_LINE_ENDS)
	DELETE_PARAGRAPH_ENDS    = C.GtkDeleteType(C.GTK_DELETE_PARAGRAPH_ENDS)
	DELETE_PARAGRAPHS        = C.GtkDeleteType(C.GTK_DELETE_PARAGRAPHS)
	DELETE_WHITESPACE        = C.GtkDeleteType(C.GTK_DELETE_WHITESPACE)

	// DirectionType
	DIR_TAB_FORWARD  = C.GtkDirectionType(C.GTK_DIR_TAB_FORWARD)
	DIR_TAB_BACKWARD = C.GtkDirectionType(C.GTK_DIR_TAB_BACKWARD)
	DIR_UP           = C.GtkDirectionType(C.GTK_DIR_UP)
	DIR_DOWN         = C.GtkDirectionType(C.GTK_DIR_DOWN)
	DIR_LEFT         = C.GtkDirectionType(C.GTK_DIR_LEFT)
	DIR_RIGHT        = C.GtkDirectionType(C.GTK_DIR_RIGHT)

	// DragResult
	DRAG_RESULT_SUCCESS         = C.GtkDragResult(C.GTK_DRAG_RESULT_SUCCESS)
	DRAG_RESULT_NO_TARGET       = C.GtkDragResult(C.GTK_DRAG_RESULT_NO_TARGET)
	DRAG_RESULT_USER_CANCELLED  = C.GtkDragResult(C.GTK_DRAG_RESULT_USER_CANCELLED)
	DRAG_RESULT_TIMEOUT_EXPIRED = C.GtkDragResult(C.GTK_DRAG_RESULT_TIMEOUT_EXPIRED)
	DRAG_RESULT_GRAB_BROKEN     = C.GtkDragResult(C.GTK_DRAG_RESULT_GRAB_BROKEN)
	DRAG_RESULT_ERROR           = C.GtkDragResult(C.GTK_DRAG_RESULT_ERROR)

	// EntryIconPosition
	ENTRY_ICON_PRIMARY   = C.GtkEntryIconPosition(C.GTK_ENTRY_ICON_PRIMARY)
	ENTRY_ICON_SECONDARY = C.GtkEntryIconPosition(C.GTK_ENTRY_ICON_SECONDARY)

	// EventSequenceState
	EVENT_SEQUENCE_NONE    = C.GtkEventSequenceState(C.GTK_EVENT_SEQUENCE_NONE)
	EVENT_SEQUENCE_CLAIMED = C.GtkEventSequenceState(C.GTK_EVENT_SEQUENCE_CLAIMED)
	EVENT_SEQUENCE_DENIED  = C.GtkEventSequenceState(C.GTK_EVENT_SEQUENCE_DENIED)

	// ExpanderStyle
	EXPANDER_COLLAPSED      = C.GtkExpanderStyle(C.GTK_EXPANDER_COLLAPSED)
	EXPANDER_SEMI_COLLAPSED = C.GtkExpanderStyle(C.GTK_EXPANDER_SEMI_COLLAPSED)
	EXPANDER_SEMI_EXPANDED  = C.GtkExpanderStyle(C.GTK_EXPANDER_SEMI_EXPANDED)
	EXPANDER_EXPANDED       = C.GtkExpanderStyle(C.GTK_EXPANDER_EXPANDED)

	// FileChooserAction
	FILE_CHOOSER_ACTION_OPEN          = C.GtkFileChooserAction(C.GTK_FILE_CHOOSER_ACTION_OPEN)
	FILE_CHOOSER_ACTION_SAVE          = C.GtkFileChooserAction(C.GTK_FILE_CHOOSER_ACTION_SAVE)
	FILE_CHOOSER_ACTION_SELECT_FOLDER = C.GtkFileChooserAction(C.GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER)
	FILE_CHOOSER_ACTION_CREATE_FOLDER = C.GtkFileChooserAction(C.GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER)

	// FileChooserConfirmation
	FILE_CHOOSER_CONFIRMATION_CONFIRM         = C.GtkFileChooserConfirmation(C.GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM)
	FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME = C.GtkFileChooserConfirmation(C.GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME)
	FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN    = C.GtkFileChooserConfirmation(C.GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN)

	// FileChooserError
	FILE_CHOOSER_ERROR_NONEXISTENT         = C.GtkFileChooserError(C.GTK_FILE_CHOOSER_ERROR_NONEXISTENT)
	FILE_CHOOSER_ERROR_BAD_FILENAME        = C.GtkFileChooserError(C.GTK_FILE_CHOOSER_ERROR_BAD_FILENAME)
	FILE_CHOOSER_ERROR_ALREADY_EXISTS      = C.GtkFileChooserError(C.GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS)
	FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME = C.GtkFileChooserError(C.GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME)

	// IMPreeditStyle
	IM_PREEDIT_NOTHING  = C.GtkIMPreeditStyle(C.GTK_IM_PREEDIT_NOTHING)
	IM_PREEDIT_CALLBACK = C.GtkIMPreeditStyle(C.GTK_IM_PREEDIT_CALLBACK)
	IM_PREEDIT_NONE     = C.GtkIMPreeditStyle(C.GTK_IM_PREEDIT_NONE)

	// IMStatusStyle
	IM_STATUS_NOTHING  = C.GtkIMStatusStyle(C.GTK_IM_STATUS_NOTHING)
	IM_STATUS_CALLBACK = C.GtkIMStatusStyle(C.GTK_IM_STATUS_CALLBACK)
	IM_STATUS_NONE     = C.GtkIMStatusStyle(C.GTK_IM_STATUS_NONE)

	// IconSize
	ICON_SIZE_INVALID       = C.GtkIconSize(C.GTK_ICON_SIZE_INVALID)
	ICON_SIZE_MENU          = C.GtkIconSize(C.GTK_ICON_SIZE_MENU)
	ICON_SIZE_SMALL_TOOLBAR = C.GtkIconSize(C.GTK_ICON_SIZE_SMALL_TOOLBAR)
	ICON_SIZE_LARGE_TOOLBAR = C.GtkIconSize(C.GTK_ICON_SIZE_LARGE_TOOLBAR)
	ICON_SIZE_BUTTON        = C.GtkIconSize(C.GTK_ICON_SIZE_BUTTON)
	ICON_SIZE_DND           = C.GtkIconSize(C.GTK_ICON_SIZE_DND)
	ICON_SIZE_DIALOG        = C.GtkIconSize(C.GTK_ICON_SIZE_DIALOG)

	// IconThemeError
	ICON_THEME_NOT_FOUND = C.GtkIconThemeError(C.GTK_ICON_THEME_NOT_FOUND)
	ICON_THEME_FAILED    = C.GtkIconThemeError(C.GTK_ICON_THEME_FAILED)

	// IconViewDropPosition
	ICON_VIEW_NO_DROP    = C.GtkIconViewDropPosition(C.GTK_ICON_VIEW_NO_DROP)
	ICON_VIEW_DROP_INTO  = C.GtkIconViewDropPosition(C.GTK_ICON_VIEW_DROP_INTO)
	ICON_VIEW_DROP_LEFT  = C.GtkIconViewDropPosition(C.GTK_ICON_VIEW_DROP_LEFT)
	ICON_VIEW_DROP_RIGHT = C.GtkIconViewDropPosition(C.GTK_ICON_VIEW_DROP_RIGHT)
	ICON_VIEW_DROP_ABOVE = C.GtkIconViewDropPosition(C.GTK_ICON_VIEW_DROP_ABOVE)
	ICON_VIEW_DROP_BELOW = C.GtkIconViewDropPosition(C.GTK_ICON_VIEW_DROP_BELOW)

	// ImageType
	IMAGE_EMPTY     = C.GtkImageType(C.GTK_IMAGE_EMPTY)
	IMAGE_PIXBUF    = C.GtkImageType(C.GTK_IMAGE_PIXBUF)
	IMAGE_STOCK     = C.GtkImageType(C.GTK_IMAGE_STOCK)
	IMAGE_ICON_SET  = C.GtkImageType(C.GTK_IMAGE_ICON_SET)
	IMAGE_ANIMATION = C.GtkImageType(C.GTK_IMAGE_ANIMATION)
	IMAGE_ICON_NAME = C.GtkImageType(C.GTK_IMAGE_ICON_NAME)
	IMAGE_GICON     = C.GtkImageType(C.GTK_IMAGE_GICON)
	IMAGE_SURFACE   = C.GtkImageType(C.GTK_IMAGE_SURFACE)

	// InputPurpose
	INPUT_PURPOSE_FREE_FORM = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_FREE_FORM)
	INPUT_PURPOSE_ALPHA     = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_ALPHA)
	INPUT_PURPOSE_DIGITS    = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_DIGITS)
	INPUT_PURPOSE_NUMBER    = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_NUMBER)
	INPUT_PURPOSE_PHONE     = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_PHONE)
	INPUT_PURPOSE_URL       = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_URL)
	INPUT_PURPOSE_EMAIL     = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_EMAIL)
	INPUT_PURPOSE_NAME      = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_NAME)
	INPUT_PURPOSE_PASSWORD  = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_PASSWORD)
	INPUT_PURPOSE_PIN       = C.GtkInputPurpose(C.GTK_INPUT_PURPOSE_PIN)

	// Justification
	JUSTIFY_LEFT   = C.GtkJustification(C.GTK_JUSTIFY_LEFT)
	JUSTIFY_RIGHT  = C.GtkJustification(C.GTK_JUSTIFY_RIGHT)
	JUSTIFY_CENTER = C.GtkJustification(C.GTK_JUSTIFY_CENTER)
	JUSTIFY_FILL   = C.GtkJustification(C.GTK_JUSTIFY_FILL)

	// LevelBarMode
	LEVEL_BAR_MODE_CONTINUOUS = C.GtkLevelBarMode(C.GTK_LEVEL_BAR_MODE_CONTINUOUS)
	LEVEL_BAR_MODE_DISCRETE   = C.GtkLevelBarMode(C.GTK_LEVEL_BAR_MODE_DISCRETE)

	// License
	LICENSE_UNKNOWN       = C.GtkLicense(C.GTK_LICENSE_UNKNOWN)
	LICENSE_CUSTOM        = C.GtkLicense(C.GTK_LICENSE_CUSTOM)
	LICENSE_GPL_2_0       = C.GtkLicense(C.GTK_LICENSE_GPL_2_0)
	LICENSE_GPL_3_0       = C.GtkLicense(C.GTK_LICENSE_GPL_3_0)
	LICENSE_LGPL_2_1      = C.GtkLicense(C.GTK_LICENSE_LGPL_2_1)
	LICENSE_LGPL_3_0      = C.GtkLicense(C.GTK_LICENSE_LGPL_3_0)
	LICENSE_BSD           = C.GtkLicense(C.GTK_LICENSE_BSD)
	LICENSE_MIT_X11       = C.GtkLicense(C.GTK_LICENSE_MIT_X11)
	LICENSE_ARTISTIC      = C.GtkLicense(C.GTK_LICENSE_ARTISTIC)
	LICENSE_GPL_2_0_ONLY  = C.GtkLicense(C.GTK_LICENSE_GPL_2_0_ONLY)
	LICENSE_GPL_3_0_ONLY  = C.GtkLicense(C.GTK_LICENSE_GPL_3_0_ONLY)
	LICENSE_LGPL_2_1_ONLY = C.GtkLicense(C.GTK_LICENSE_LGPL_2_1_ONLY)
	LICENSE_LGPL_3_0_ONLY = C.GtkLicense(C.GTK_LICENSE_LGPL_3_0_ONLY)

	// MenuDirectionType
	MENU_DIR_PARENT = C.GtkMenuDirectionType(C.GTK_MENU_DIR_PARENT)
	MENU_DIR_CHILD  = C.GtkMenuDirectionType(C.GTK_MENU_DIR_CHILD)
	MENU_DIR_NEXT   = C.GtkMenuDirectionType(C.GTK_MENU_DIR_NEXT)
	MENU_DIR_PREV   = C.GtkMenuDirectionType(C.GTK_MENU_DIR_PREV)

	// MessageType
	MESSAGE_INFO     = C.GtkMessageType(C.GTK_MESSAGE_INFO)
	MESSAGE_WARNING  = C.GtkMessageType(C.GTK_MESSAGE_WARNING)
	MESSAGE_QUESTION = C.GtkMessageType(C.GTK_MESSAGE_QUESTION)
	MESSAGE_ERROR    = C.GtkMessageType(C.GTK_MESSAGE_ERROR)
	MESSAGE_OTHER    = C.GtkMessageType(C.GTK_MESSAGE_OTHER)

	// MovementStep
	MOVEMENT_LOGICAL_POSITIONS = C.GtkMovementStep(C.GTK_MOVEMENT_LOGICAL_POSITIONS)
	MOVEMENT_VISUAL_POSITIONS  = C.GtkMovementStep(C.GTK_MOVEMENT_VISUAL_POSITIONS)
	MOVEMENT_WORDS             = C.GtkMovementStep(C.GTK_MOVEMENT_WORDS)
	MOVEMENT_DISPLAY_LINES     = C.GtkMovementStep(C.GTK_MOVEMENT_DISPLAY_LINES)
	MOVEMENT_DISPLAY_LINE_ENDS = C.GtkMovementStep(C.GTK_MOVEMENT_DISPLAY_LINE_ENDS)
	MOVEMENT_PARAGRAPHS        = C.GtkMovementStep(C.GTK_MOVEMENT_PARAGRAPHS)
	MOVEMENT_PARAGRAPH_ENDS    = C.GtkMovementStep(C.GTK_MOVEMENT_PARAGRAPH_ENDS)
	MOVEMENT_PAGES             = C.GtkMovementStep(C.GTK_MOVEMENT_PAGES)
	MOVEMENT_BUFFER_ENDS       = C.GtkMovementStep(C.GTK_MOVEMENT_BUFFER_ENDS)
	MOVEMENT_HORIZONTAL_PAGES  = C.GtkMovementStep(C.GTK_MOVEMENT_HORIZONTAL_PAGES)

	// NotebookTab
	NOTEBOOK_TAB_FIRST = C.GtkNotebookTab(C.GTK_NOTEBOOK_TAB_FIRST)
	NOTEBOOK_TAB_LAST  = C.GtkNotebookTab(C.GTK_NOTEBOOK_TAB_LAST)

	// NumberUpLayout
	NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM = C.GtkNumberUpLayout(C.GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM)
	NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP = C.GtkNumberUpLayout(C.GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP)
	NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM = C.GtkNumberUpLayout(C.GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM)
	NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP = C.GtkNumberUpLayout(C.GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP)
	NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT = C.GtkNumberUpLayout(C.GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT)
	NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT = C.GtkNumberUpLayout(C.GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT)
	NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT = C.GtkNumberUpLayout(C.GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT)
	NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT = C.GtkNumberUpLayout(C.GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT)

	// Orientation
	ORIENTATION_HORIZONTAL = C.GtkOrientation(C.GTK_ORIENTATION_HORIZONTAL)
	ORIENTATION_VERTICAL   = C.GtkOrientation(C.GTK_ORIENTATION_VERTICAL)

	// PackDirection
	PACK_DIRECTION_LTR = C.GtkPackDirection(C.GTK_PACK_DIRECTION_LTR)
	PACK_DIRECTION_RTL = C.GtkPackDirection(C.GTK_PACK_DIRECTION_RTL)
	PACK_DIRECTION_TTB = C.GtkPackDirection(C.GTK_PACK_DIRECTION_TTB)
	PACK_DIRECTION_BTT = C.GtkPackDirection(C.GTK_PACK_DIRECTION_BTT)

	// PackType
	PACK_START = C.GtkPackType(C.GTK_PACK_START)
	PACK_END   = C.GtkPackType(C.GTK_PACK_END)

	// PageOrientation
	PAGE_ORIENTATION_PORTRAIT          = C.GtkPageOrientation(C.GTK_PAGE_ORIENTATION_PORTRAIT)
	PAGE_ORIENTATION_LANDSCAPE         = C.GtkPageOrientation(C.GTK_PAGE_ORIENTATION_LANDSCAPE)
	PAGE_ORIENTATION_REVERSE_PORTRAIT  = C.GtkPageOrientation(C.GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT)
	PAGE_ORIENTATION_REVERSE_LANDSCAPE = C.GtkPageOrientation(C.GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE)

	// PageSet
	PAGE_SET_ALL  = C.GtkPageSet(C.GTK_PAGE_SET_ALL)
	PAGE_SET_EVEN = C.GtkPageSet(C.GTK_PAGE_SET_EVEN)
	PAGE_SET_ODD  = C.GtkPageSet(C.GTK_PAGE_SET_ODD)

	// PanDirection
	PAN_DIRECTION_LEFT  = C.GtkPanDirection(C.GTK_PAN_DIRECTION_LEFT)
	PAN_DIRECTION_RIGHT = C.GtkPanDirection(C.GTK_PAN_DIRECTION_RIGHT)
	PAN_DIRECTION_UP    = C.GtkPanDirection(C.GTK_PAN_DIRECTION_UP)
	PAN_DIRECTION_DOWN  = C.GtkPanDirection(C.GTK_PAN_DIRECTION_DOWN)

	// PathPriorityType
	PATH_PRIO_LOWEST      = C.GtkPathPriorityType(C.GTK_PATH_PRIO_LOWEST)
	PATH_PRIO_GTK         = C.GtkPathPriorityType(C.GTK_PATH_PRIO_GTK)
	PATH_PRIO_APPLICATION = C.GtkPathPriorityType(C.GTK_PATH_PRIO_APPLICATION)
	PATH_PRIO_THEME       = C.GtkPathPriorityType(C.GTK_PATH_PRIO_THEME)
	PATH_PRIO_RC          = C.GtkPathPriorityType(C.GTK_PATH_PRIO_RC)
	PATH_PRIO_HIGHEST     = C.GtkPathPriorityType(C.GTK_PATH_PRIO_HIGHEST)

	// PathType
	PATH_WIDGET       = C.GtkPathType(C.GTK_PATH_WIDGET)
	PATH_WIDGET_CLASS = C.GtkPathType(C.GTK_PATH_WIDGET_CLASS)
	PATH_CLASS        = C.GtkPathType(C.GTK_PATH_CLASS)

	// PolicyType
	POLICY_ALWAYS    = C.GtkPolicyType(C.GTK_POLICY_ALWAYS)
	POLICY_AUTOMATIC = C.GtkPolicyType(C.GTK_POLICY_AUTOMATIC)
	POLICY_NEVER     = C.GtkPolicyType(C.GTK_POLICY_NEVER)

	// PositionType
	POS_LEFT   = C.GtkPositionType(C.GTK_POS_LEFT)
	POS_RIGHT  = C.GtkPositionType(C.GTK_POS_RIGHT)
	POS_TOP    = C.GtkPositionType(C.GTK_POS_TOP)
	POS_BOTTOM = C.GtkPositionType(C.GTK_POS_BOTTOM)

	// PrintDuplex
	PRINT_DUPLEX_SIMPLEX    = C.GtkPrintDuplex(C.GTK_PRINT_DUPLEX_SIMPLEX)
	PRINT_DUPLEX_HORIZONTAL = C.GtkPrintDuplex(C.GTK_PRINT_DUPLEX_HORIZONTAL)
	PRINT_DUPLEX_VERTICAL   = C.GtkPrintDuplex(C.GTK_PRINT_DUPLEX_VERTICAL)

	// PrintError
	PRINT_ERROR_GENERAL        = C.GtkPrintError(C.GTK_PRINT_ERROR_GENERAL)
	PRINT_ERROR_INTERNAL_ERROR = C.GtkPrintError(C.GTK_PRINT_ERROR_INTERNAL_ERROR)
	PRINT_ERROR_NOMEM          = C.GtkPrintError(C.GTK_PRINT_ERROR_NOMEM)
	PRINT_ERROR_INVALID_FILE   = C.GtkPrintError(C.GTK_PRINT_ERROR_INVALID_FILE)

	// PrintOperationAction
	PRINT_OPERATION_ACTION_PRINT_DIALOG = C.GtkPrintOperationAction(C.GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG)
	PRINT_OPERATION_ACTION_PRINT        = C.GtkPrintOperationAction(C.GTK_PRINT_OPERATION_ACTION_PRINT)
	PRINT_OPERATION_ACTION_PREVIEW      = C.GtkPrintOperationAction(C.GTK_PRINT_OPERATION_ACTION_PREVIEW)
	PRINT_OPERATION_ACTION_EXPORT       = C.GtkPrintOperationAction(C.GTK_PRINT_OPERATION_ACTION_EXPORT)

	// PrintOperationResult
	PRINT_OPERATION_RESULT_ERROR       = C.GtkPrintOperationResult(C.GTK_PRINT_OPERATION_RESULT_ERROR)
	PRINT_OPERATION_RESULT_APPLY       = C.GtkPrintOperationResult(C.GTK_PRINT_OPERATION_RESULT_APPLY)
	PRINT_OPERATION_RESULT_CANCEL      = C.GtkPrintOperationResult(C.GTK_PRINT_OPERATION_RESULT_CANCEL)
	PRINT_OPERATION_RESULT_IN_PROGRESS = C.GtkPrintOperationResult(C.GTK_PRINT_OPERATION_RESULT_IN_PROGRESS)

	// PrintPages
	PRINT_PAGES_ALL       = C.GtkPrintPages(C.GTK_PRINT_PAGES_ALL)
	PRINT_PAGES_CURRENT   = C.GtkPrintPages(C.GTK_PRINT_PAGES_CURRENT)
	PRINT_PAGES_RANGES    = C.GtkPrintPages(C.GTK_PRINT_PAGES_RANGES)
	PRINT_PAGES_SELECTION = C.GtkPrintPages(C.GTK_PRINT_PAGES_SELECTION)

	// PrintQuality
	PRINT_QUALITY_LOW    = C.GtkPrintQuality(C.GTK_PRINT_QUALITY_LOW)
	PRINT_QUALITY_NORMAL = C.GtkPrintQuality(C.GTK_PRINT_QUALITY_NORMAL)
	PRINT_QUALITY_HIGH   = C.GtkPrintQuality(C.GTK_PRINT_QUALITY_HIGH)
	PRINT_QUALITY_DRAFT  = C.GtkPrintQuality(C.GTK_PRINT_QUALITY_DRAFT)

	// PrintStatus
	PRINT_STATUS_INITIAL          = C.GtkPrintStatus(C.GTK_PRINT_STATUS_INITIAL)
	PRINT_STATUS_PREPARING        = C.GtkPrintStatus(C.GTK_PRINT_STATUS_PREPARING)
	PRINT_STATUS_GENERATING_DATA  = C.GtkPrintStatus(C.GTK_PRINT_STATUS_GENERATING_DATA)
	PRINT_STATUS_SENDING_DATA     = C.GtkPrintStatus(C.GTK_PRINT_STATUS_SENDING_DATA)
	PRINT_STATUS_PENDING          = C.GtkPrintStatus(C.GTK_PRINT_STATUS_PENDING)
	PRINT_STATUS_PENDING_ISSUE    = C.GtkPrintStatus(C.GTK_PRINT_STATUS_PENDING_ISSUE)
	PRINT_STATUS_PRINTING         = C.GtkPrintStatus(C.GTK_PRINT_STATUS_PRINTING)
	PRINT_STATUS_FINISHED         = C.GtkPrintStatus(C.GTK_PRINT_STATUS_FINISHED)
	PRINT_STATUS_FINISHED_ABORTED = C.GtkPrintStatus(C.GTK_PRINT_STATUS_FINISHED_ABORTED)

	// PropagationPhase
	PHASE_NONE    = C.GtkPropagationPhase(C.GTK_PHASE_NONE)
	PHASE_CAPTURE = C.GtkPropagationPhase(C.GTK_PHASE_CAPTURE)
	PHASE_BUBBLE  = C.GtkPropagationPhase(C.GTK_PHASE_BUBBLE)
	PHASE_TARGET  = C.GtkPropagationPhase(C.GTK_PHASE_TARGET)

	// RcTokenType
	RC_TOKEN_INVALID        = C.GtkRcTokenType(C.GTK_RC_TOKEN_INVALID)
	RC_TOKEN_INCLUDE        = C.GtkRcTokenType(C.GTK_RC_TOKEN_INCLUDE)
	RC_TOKEN_NORMAL         = C.GtkRcTokenType(C.GTK_RC_TOKEN_NORMAL)
	RC_TOKEN_ACTIVE         = C.GtkRcTokenType(C.GTK_RC_TOKEN_ACTIVE)
	RC_TOKEN_PRELIGHT       = C.GtkRcTokenType(C.GTK_RC_TOKEN_PRELIGHT)
	RC_TOKEN_SELECTED       = C.GtkRcTokenType(C.GTK_RC_TOKEN_SELECTED)
	RC_TOKEN_INSENSITIVE    = C.GtkRcTokenType(C.GTK_RC_TOKEN_INSENSITIVE)
	RC_TOKEN_FG             = C.GtkRcTokenType(C.GTK_RC_TOKEN_FG)
	RC_TOKEN_BG             = C.GtkRcTokenType(C.GTK_RC_TOKEN_BG)
	RC_TOKEN_TEXT           = C.GtkRcTokenType(C.GTK_RC_TOKEN_TEXT)
	RC_TOKEN_BASE           = C.GtkRcTokenType(C.GTK_RC_TOKEN_BASE)
	RC_TOKEN_XTHICKNESS     = C.GtkRcTokenType(C.GTK_RC_TOKEN_XTHICKNESS)
	RC_TOKEN_YTHICKNESS     = C.GtkRcTokenType(C.GTK_RC_TOKEN_YTHICKNESS)
	RC_TOKEN_FONT           = C.GtkRcTokenType(C.GTK_RC_TOKEN_FONT)
	RC_TOKEN_FONTSET        = C.GtkRcTokenType(C.GTK_RC_TOKEN_FONTSET)
	RC_TOKEN_FONT_NAME      = C.GtkRcTokenType(C.GTK_RC_TOKEN_FONT_NAME)
	RC_TOKEN_BG_PIXMAP      = C.GtkRcTokenType(C.GTK_RC_TOKEN_BG_PIXMAP)
	RC_TOKEN_PIXMAP_PATH    = C.GtkRcTokenType(C.GTK_RC_TOKEN_PIXMAP_PATH)
	RC_TOKEN_STYLE          = C.GtkRcTokenType(C.GTK_RC_TOKEN_STYLE)
	RC_TOKEN_BINDING        = C.GtkRcTokenType(C.GTK_RC_TOKEN_BINDING)
	RC_TOKEN_BIND           = C.GtkRcTokenType(C.GTK_RC_TOKEN_BIND)
	RC_TOKEN_WIDGET         = C.GtkRcTokenType(C.GTK_RC_TOKEN_WIDGET)
	RC_TOKEN_WIDGET_CLASS   = C.GtkRcTokenType(C.GTK_RC_TOKEN_WIDGET_CLASS)
	RC_TOKEN_CLASS          = C.GtkRcTokenType(C.GTK_RC_TOKEN_CLASS)
	RC_TOKEN_LOWEST         = C.GtkRcTokenType(C.GTK_RC_TOKEN_LOWEST)
	RC_TOKEN_GTK            = C.GtkRcTokenType(C.GTK_RC_TOKEN_GTK)
	RC_TOKEN_APPLICATION    = C.GtkRcTokenType(C.GTK_RC_TOKEN_APPLICATION)
	RC_TOKEN_THEME          = C.GtkRcTokenType(C.GTK_RC_TOKEN_THEME)
	RC_TOKEN_RC             = C.GtkRcTokenType(C.GTK_RC_TOKEN_RC)
	RC_TOKEN_HIGHEST        = C.GtkRcTokenType(C.GTK_RC_TOKEN_HIGHEST)
	RC_TOKEN_ENGINE         = C.GtkRcTokenType(C.GTK_RC_TOKEN_ENGINE)
	RC_TOKEN_MODULE_PATH    = C.GtkRcTokenType(C.GTK_RC_TOKEN_MODULE_PATH)
	RC_TOKEN_IM_MODULE_PATH = C.GtkRcTokenType(C.GTK_RC_TOKEN_IM_MODULE_PATH)
	RC_TOKEN_IM_MODULE_FILE = C.GtkRcTokenType(C.GTK_RC_TOKEN_IM_MODULE_FILE)
	RC_TOKEN_STOCK          = C.GtkRcTokenType(C.GTK_RC_TOKEN_STOCK)
	RC_TOKEN_LTR            = C.GtkRcTokenType(C.GTK_RC_TOKEN_LTR)
	RC_TOKEN_RTL            = C.GtkRcTokenType(C.GTK_RC_TOKEN_RTL)
	RC_TOKEN_COLOR          = C.GtkRcTokenType(C.GTK_RC_TOKEN_COLOR)
	RC_TOKEN_UNBIND         = C.GtkRcTokenType(C.GTK_RC_TOKEN_UNBIND)
	RC_TOKEN_LAST           = C.GtkRcTokenType(C.GTK_RC_TOKEN_LAST)

	// RecentChooserError
	RECENT_CHOOSER_ERROR_NOT_FOUND   = C.GtkRecentChooserError(C.GTK_RECENT_CHOOSER_ERROR_NOT_FOUND)
	RECENT_CHOOSER_ERROR_INVALID_URI = C.GtkRecentChooserError(C.GTK_RECENT_CHOOSER_ERROR_INVALID_URI)

	// RecentManagerError
	RECENT_MANAGER_ERROR_NOT_FOUND        = C.GtkRecentManagerError(C.GTK_RECENT_MANAGER_ERROR_NOT_FOUND)
	RECENT_MANAGER_ERROR_INVALID_URI      = C.GtkRecentManagerError(C.GTK_RECENT_MANAGER_ERROR_INVALID_URI)
	RECENT_MANAGER_ERROR_INVALID_ENCODING = C.GtkRecentManagerError(C.GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING)
	RECENT_MANAGER_ERROR_NOT_REGISTERED   = C.GtkRecentManagerError(C.GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED)
	RECENT_MANAGER_ERROR_READ             = C.GtkRecentManagerError(C.GTK_RECENT_MANAGER_ERROR_READ)
	RECENT_MANAGER_ERROR_WRITE            = C.GtkRecentManagerError(C.GTK_RECENT_MANAGER_ERROR_WRITE)
	RECENT_MANAGER_ERROR_UNKNOWN          = C.GtkRecentManagerError(C.GTK_RECENT_MANAGER_ERROR_UNKNOWN)

	// RecentSortType
	RECENT_SORT_NONE   = C.GtkRecentSortType(C.GTK_RECENT_SORT_NONE)
	RECENT_SORT_MRU    = C.GtkRecentSortType(C.GTK_RECENT_SORT_MRU)
	RECENT_SORT_LRU    = C.GtkRecentSortType(C.GTK_RECENT_SORT_LRU)
	RECENT_SORT_CUSTOM = C.GtkRecentSortType(C.GTK_RECENT_SORT_CUSTOM)

	// ReliefStyle
	RELIEF_NORMAL = C.GtkReliefStyle(C.GTK_RELIEF_NORMAL)
	RELIEF_HALF   = C.GtkReliefStyle(C.GTK_RELIEF_HALF)
	RELIEF_NONE   = C.GtkReliefStyle(C.GTK_RELIEF_NONE)

	// ResizeMode
	RESIZE_PARENT    = C.GtkResizeMode(C.GTK_RESIZE_PARENT)
	RESIZE_QUEUE     = C.GtkResizeMode(C.GTK_RESIZE_QUEUE)
	RESIZE_IMMEDIATE = C.GtkResizeMode(C.GTK_RESIZE_IMMEDIATE)

	// ResponseType
	RESPONSE_NONE         = C.GtkResponseType(C.GTK_RESPONSE_NONE)
	RESPONSE_REJECT       = C.GtkResponseType(C.GTK_RESPONSE_REJECT)
	RESPONSE_ACCEPT       = C.GtkResponseType(C.GTK_RESPONSE_ACCEPT)
	RESPONSE_DELETE_EVENT = C.GtkResponseType(C.GTK_RESPONSE_DELETE_EVENT)
	RESPONSE_OK           = C.GtkResponseType(C.GTK_RESPONSE_OK)
	RESPONSE_CANCEL       = C.GtkResponseType(C.GTK_RESPONSE_CANCEL)
	RESPONSE_CLOSE        = C.GtkResponseType(C.GTK_RESPONSE_CLOSE)
	RESPONSE_YES          = C.GtkResponseType(C.GTK_RESPONSE_YES)
	RESPONSE_NO           = C.GtkResponseType(C.GTK_RESPONSE_NO)
	RESPONSE_APPLY        = C.GtkResponseType(C.GTK_RESPONSE_APPLY)
	RESPONSE_HELP         = C.GtkResponseType(C.GTK_RESPONSE_HELP)

	// RevealerTransitionType
	REVEALER_TRANSITION_TYPE_NONE        = C.GtkRevealerTransitionType(C.GTK_REVEALER_TRANSITION_TYPE_NONE)
	REVEALER_TRANSITION_TYPE_CROSSFADE   = C.GtkRevealerTransitionType(C.GTK_REVEALER_TRANSITION_TYPE_CROSSFADE)
	REVEALER_TRANSITION_TYPE_SLIDE_RIGHT = C.GtkRevealerTransitionType(C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT)
	REVEALER_TRANSITION_TYPE_SLIDE_LEFT  = C.GtkRevealerTransitionType(C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT)
	REVEALER_TRANSITION_TYPE_SLIDE_UP    = C.GtkRevealerTransitionType(C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP)
	REVEALER_TRANSITION_TYPE_SLIDE_DOWN  = C.GtkRevealerTransitionType(C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN)

	// ScrollStep
	SCROLL_STEPS            = C.GtkScrollStep(C.GTK_SCROLL_STEPS)
	SCROLL_PAGES            = C.GtkScrollStep(C.GTK_SCROLL_PAGES)
	SCROLL_ENDS             = C.GtkScrollStep(C.GTK_SCROLL_ENDS)
	SCROLL_HORIZONTAL_STEPS = C.GtkScrollStep(C.GTK_SCROLL_HORIZONTAL_STEPS)
	SCROLL_HORIZONTAL_PAGES = C.GtkScrollStep(C.GTK_SCROLL_HORIZONTAL_PAGES)
	SCROLL_HORIZONTAL_ENDS  = C.GtkScrollStep(C.GTK_SCROLL_HORIZONTAL_ENDS)

	// ScrollType
	SCROLL_NONE          = C.GtkScrollType(C.GTK_SCROLL_NONE)
	SCROLL_JUMP          = C.GtkScrollType(C.GTK_SCROLL_JUMP)
	SCROLL_STEP_BACKWARD = C.GtkScrollType(C.GTK_SCROLL_STEP_BACKWARD)
	SCROLL_STEP_FORWARD  = C.GtkScrollType(C.GTK_SCROLL_STEP_FORWARD)
	SCROLL_PAGE_BACKWARD = C.GtkScrollType(C.GTK_SCROLL_PAGE_BACKWARD)
	SCROLL_PAGE_FORWARD  = C.GtkScrollType(C.GTK_SCROLL_PAGE_FORWARD)
	SCROLL_STEP_UP       = C.GtkScrollType(C.GTK_SCROLL_STEP_UP)
	SCROLL_STEP_DOWN     = C.GtkScrollType(C.GTK_SCROLL_STEP_DOWN)
	SCROLL_PAGE_UP       = C.GtkScrollType(C.GTK_SCROLL_PAGE_UP)
	SCROLL_PAGE_DOWN     = C.GtkScrollType(C.GTK_SCROLL_PAGE_DOWN)
	SCROLL_STEP_LEFT     = C.GtkScrollType(C.GTK_SCROLL_STEP_LEFT)
	SCROLL_STEP_RIGHT    = C.GtkScrollType(C.GTK_SCROLL_STEP_RIGHT)
	SCROLL_PAGE_LEFT     = C.GtkScrollType(C.GTK_SCROLL_PAGE_LEFT)
	SCROLL_PAGE_RIGHT    = C.GtkScrollType(C.GTK_SCROLL_PAGE_RIGHT)
	SCROLL_START         = C.GtkScrollType(C.GTK_SCROLL_START)
	SCROLL_END           = C.GtkScrollType(C.GTK_SCROLL_END)

	// ScrollablePolicy
	SCROLL_MINIMUM = C.GtkScrollablePolicy(C.GTK_SCROLL_MINIMUM)
	SCROLL_NATURAL = C.GtkScrollablePolicy(C.GTK_SCROLL_NATURAL)

	// SelectionMode
	SELECTION_NONE     = C.GtkSelectionMode(C.GTK_SELECTION_NONE)
	SELECTION_SINGLE   = C.GtkSelectionMode(C.GTK_SELECTION_SINGLE)
	SELECTION_BROWSE   = C.GtkSelectionMode(C.GTK_SELECTION_BROWSE)
	SELECTION_MULTIPLE = C.GtkSelectionMode(C.GTK_SELECTION_MULTIPLE)

	// SensitivityType
	SENSITIVITY_AUTO = C.GtkSensitivityType(C.GTK_SENSITIVITY_AUTO)
	SENSITIVITY_ON   = C.GtkSensitivityType(C.GTK_SENSITIVITY_ON)
	SENSITIVITY_OFF  = C.GtkSensitivityType(C.GTK_SENSITIVITY_OFF)

	// ShadowType
	SHADOW_NONE       = C.GtkShadowType(C.GTK_SHADOW_NONE)
	SHADOW_IN         = C.GtkShadowType(C.GTK_SHADOW_IN)
	SHADOW_OUT        = C.GtkShadowType(C.GTK_SHADOW_OUT)
	SHADOW_ETCHED_IN  = C.GtkShadowType(C.GTK_SHADOW_ETCHED_IN)
	SHADOW_ETCHED_OUT = C.GtkShadowType(C.GTK_SHADOW_ETCHED_OUT)

	// SizeGroupMode
	SIZE_GROUP_NONE       = C.GtkSizeGroupMode(C.GTK_SIZE_GROUP_NONE)
	SIZE_GROUP_HORIZONTAL = C.GtkSizeGroupMode(C.GTK_SIZE_GROUP_HORIZONTAL)
	SIZE_GROUP_VERTICAL   = C.GtkSizeGroupMode(C.GTK_SIZE_GROUP_VERTICAL)
	SIZE_GROUP_BOTH       = C.GtkSizeGroupMode(C.GTK_SIZE_GROUP_BOTH)

	// SizeRequestMode
	SIZE_REQUEST_HEIGHT_FOR_WIDTH = C.GtkSizeRequestMode(C.GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH)
	SIZE_REQUEST_WIDTH_FOR_HEIGHT = C.GtkSizeRequestMode(C.GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT)
	SIZE_REQUEST_CONSTANT_SIZE    = C.GtkSizeRequestMode(C.GTK_SIZE_REQUEST_CONSTANT_SIZE)

	// SortType
	SORT_ASCENDING  = C.GtkSortType(C.GTK_SORT_ASCENDING)
	SORT_DESCENDING = C.GtkSortType(C.GTK_SORT_DESCENDING)

	// SpinButtonUpdatePolicy
	UPDATE_ALWAYS   = C.GtkSpinButtonUpdatePolicy(C.GTK_UPDATE_ALWAYS)
	UPDATE_IF_VALID = C.GtkSpinButtonUpdatePolicy(C.GTK_UPDATE_IF_VALID)

	// SpinType
	SPIN_STEP_FORWARD  = C.GtkSpinType(C.GTK_SPIN_STEP_FORWARD)
	SPIN_STEP_BACKWARD = C.GtkSpinType(C.GTK_SPIN_STEP_BACKWARD)
	SPIN_PAGE_FORWARD  = C.GtkSpinType(C.GTK_SPIN_PAGE_FORWARD)
	SPIN_PAGE_BACKWARD = C.GtkSpinType(C.GTK_SPIN_PAGE_BACKWARD)
	SPIN_HOME          = C.GtkSpinType(C.GTK_SPIN_HOME)
	SPIN_END           = C.GtkSpinType(C.GTK_SPIN_END)
	SPIN_USER_DEFINED  = C.GtkSpinType(C.GTK_SPIN_USER_DEFINED)

	// StackTransitionType
	STACK_TRANSITION_TYPE_NONE             = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_NONE)
	STACK_TRANSITION_TYPE_CROSSFADE        = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_CROSSFADE)
	STACK_TRANSITION_TYPE_SLIDE_RIGHT      = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT)
	STACK_TRANSITION_TYPE_SLIDE_LEFT       = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT)
	STACK_TRANSITION_TYPE_SLIDE_UP         = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP)
	STACK_TRANSITION_TYPE_SLIDE_DOWN       = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN)
	STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT)
	STACK_TRANSITION_TYPE_SLIDE_UP_DOWN    = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN)
	STACK_TRANSITION_TYPE_OVER_UP          = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_OVER_UP)
	STACK_TRANSITION_TYPE_OVER_DOWN        = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_OVER_DOWN)
	STACK_TRANSITION_TYPE_OVER_LEFT        = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_OVER_LEFT)
	STACK_TRANSITION_TYPE_OVER_RIGHT       = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_OVER_RIGHT)
	STACK_TRANSITION_TYPE_UNDER_UP         = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_UNDER_UP)
	STACK_TRANSITION_TYPE_UNDER_DOWN       = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_UNDER_DOWN)
	STACK_TRANSITION_TYPE_UNDER_LEFT       = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_UNDER_LEFT)
	STACK_TRANSITION_TYPE_UNDER_RIGHT      = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT)
	STACK_TRANSITION_TYPE_OVER_UP_DOWN     = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN)
	STACK_TRANSITION_TYPE_OVER_DOWN_UP     = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP)
	STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT  = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT)
	STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT  = C.GtkStackTransitionType(C.GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT)

	// StateType
	STATE_NORMAL       = C.GtkStateType(C.GTK_STATE_NORMAL)
	STATE_ACTIVE       = C.GtkStateType(C.GTK_STATE_ACTIVE)
	STATE_PRELIGHT     = C.GtkStateType(C.GTK_STATE_PRELIGHT)
	STATE_SELECTED     = C.GtkStateType(C.GTK_STATE_SELECTED)
	STATE_INSENSITIVE  = C.GtkStateType(C.GTK_STATE_INSENSITIVE)
	STATE_INCONSISTENT = C.GtkStateType(C.GTK_STATE_INCONSISTENT)
	STATE_FOCUSED      = C.GtkStateType(C.GTK_STATE_FOCUSED)

	// TextBufferTargetInfo
	TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS = C.GtkTextBufferTargetInfo(C.GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS)
	TEXT_BUFFER_TARGET_INFO_RICH_TEXT       = C.GtkTextBufferTargetInfo(C.GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT)
	TEXT_BUFFER_TARGET_INFO_TEXT            = C.GtkTextBufferTargetInfo(C.GTK_TEXT_BUFFER_TARGET_INFO_TEXT)

	// TextDirection
	TEXT_DIR_NONE = C.GtkTextDirection(C.GTK_TEXT_DIR_NONE)
	TEXT_DIR_LTR  = C.GtkTextDirection(C.GTK_TEXT_DIR_LTR)
	TEXT_DIR_RTL  = C.GtkTextDirection(C.GTK_TEXT_DIR_RTL)

	// TextViewLayer
	TEXT_VIEW_LAYER_BELOW = C.GtkTextViewLayer(C.GTK_TEXT_VIEW_LAYER_BELOW)
	TEXT_VIEW_LAYER_ABOVE = C.GtkTextViewLayer(C.GTK_TEXT_VIEW_LAYER_ABOVE)

	// TextWindowType
	TEXT_WINDOW_PRIVATE = C.GtkTextWindowType(C.GTK_TEXT_WINDOW_PRIVATE)
	TEXT_WINDOW_WIDGET  = C.GtkTextWindowType(C.GTK_TEXT_WINDOW_WIDGET)
	TEXT_WINDOW_TEXT    = C.GtkTextWindowType(C.GTK_TEXT_WINDOW_TEXT)
	TEXT_WINDOW_LEFT    = C.GtkTextWindowType(C.GTK_TEXT_WINDOW_LEFT)
	TEXT_WINDOW_RIGHT   = C.GtkTextWindowType(C.GTK_TEXT_WINDOW_RIGHT)
	TEXT_WINDOW_TOP     = C.GtkTextWindowType(C.GTK_TEXT_WINDOW_TOP)
	TEXT_WINDOW_BOTTOM  = C.GtkTextWindowType(C.GTK_TEXT_WINDOW_BOTTOM)

	// ToolbarSpaceStyle
	TOOLBAR_SPACE_EMPTY = C.GtkToolbarSpaceStyle(C.GTK_TOOLBAR_SPACE_EMPTY)
	TOOLBAR_SPACE_LINE  = C.GtkToolbarSpaceStyle(C.GTK_TOOLBAR_SPACE_LINE)

	// ToolbarStyle
	TOOLBAR_ICONS      = C.GtkToolbarStyle(C.GTK_TOOLBAR_ICONS)
	TOOLBAR_TEXT       = C.GtkToolbarStyle(C.GTK_TOOLBAR_TEXT)
	TOOLBAR_BOTH       = C.GtkToolbarStyle(C.GTK_TOOLBAR_BOTH)
	TOOLBAR_BOTH_HORIZ = C.GtkToolbarStyle(C.GTK_TOOLBAR_BOTH_HORIZ)

	// TreeViewColumnSizing
	TREE_VIEW_COLUMN_GROW_ONLY = C.GtkTreeViewColumnSizing(C.GTK_TREE_VIEW_COLUMN_GROW_ONLY)
	TREE_VIEW_COLUMN_AUTOSIZE  = C.GtkTreeViewColumnSizing(C.GTK_TREE_VIEW_COLUMN_AUTOSIZE)
	TREE_VIEW_COLUMN_FIXED     = C.GtkTreeViewColumnSizing(C.GTK_TREE_VIEW_COLUMN_FIXED)

	// TreeViewDropPosition
	TREE_VIEW_DROP_BEFORE         = C.GtkTreeViewDropPosition(C.GTK_TREE_VIEW_DROP_BEFORE)
	TREE_VIEW_DROP_AFTER          = C.GtkTreeViewDropPosition(C.GTK_TREE_VIEW_DROP_AFTER)
	TREE_VIEW_DROP_INTO_OR_BEFORE = C.GtkTreeViewDropPosition(C.GTK_TREE_VIEW_DROP_INTO_OR_BEFORE)
	TREE_VIEW_DROP_INTO_OR_AFTER  = C.GtkTreeViewDropPosition(C.GTK_TREE_VIEW_DROP_INTO_OR_AFTER)

	// TreeViewGridLines
	TREE_VIEW_GRID_LINES_NONE       = C.GtkTreeViewGridLines(C.GTK_TREE_VIEW_GRID_LINES_NONE)
	TREE_VIEW_GRID_LINES_HORIZONTAL = C.GtkTreeViewGridLines(C.GTK_TREE_VIEW_GRID_LINES_HORIZONTAL)
	TREE_VIEW_GRID_LINES_VERTICAL   = C.GtkTreeViewGridLines(C.GTK_TREE_VIEW_GRID_LINES_VERTICAL)
	TREE_VIEW_GRID_LINES_BOTH       = C.GtkTreeViewGridLines(C.GTK_TREE_VIEW_GRID_LINES_BOTH)

	// Unit
	UNIT_NONE   = C.GtkUnit(C.GTK_UNIT_NONE)
	UNIT_POINTS = C.GtkUnit(C.GTK_UNIT_POINTS)
	UNIT_INCH   = C.GtkUnit(C.GTK_UNIT_INCH)
	UNIT_MM     = C.GtkUnit(C.GTK_UNIT_MM)

	// WidgetHelpType
	WIDGET_HELP_TOOLTIP    = C.GtkWidgetHelpType(C.GTK_WIDGET_HELP_TOOLTIP)
	WIDGET_HELP_WHATS_THIS = C.GtkWidgetHelpType(C.GTK_WIDGET_HELP_WHATS_THIS)

	// WindowPosition
	WIN_POS_NONE             = C.GtkWindowPosition(C.GTK_WIN_POS_NONE)
	WIN_POS_CENTER           = C.GtkWindowPosition(C.GTK_WIN_POS_CENTER)
	WIN_POS_MOUSE            = C.GtkWindowPosition(C.GTK_WIN_POS_MOUSE)
	WIN_POS_CENTER_ALWAYS    = C.GtkWindowPosition(C.GTK_WIN_POS_CENTER_ALWAYS)
	WIN_POS_CENTER_ON_PARENT = C.GtkWindowPosition(C.GTK_WIN_POS_CENTER_ON_PARENT)

	// WindowType
	WINDOW_TOPLEVEL = C.GtkWindowType(C.GTK_WINDOW_TOPLEVEL)
	WINDOW_POPUP    = C.GtkWindowType(C.GTK_WINDOW_POPUP)

	// WrapMode
	WRAP_NONE      = C.GtkWrapMode(C.GTK_WRAP_NONE)
	WRAP_CHAR      = C.GtkWrapMode(C.GTK_WRAP_CHAR)
	WRAP_WORD      = C.GtkWrapMode(C.GTK_WRAP_WORD)
	WRAP_WORD_CHAR = C.GtkWrapMode(C.GTK_WRAP_WORD_CHAR)
)
