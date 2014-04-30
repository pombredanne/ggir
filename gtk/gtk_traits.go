package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"
import "errors"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
	_ = errors.New("")
}

type _TraitAboutDialog struct{ CPointer *C.GtkAboutDialog }

/*
Creates a new section in the Credits page.
*/
func (self *_TraitAboutDialog) AddCreditSection(section_name string, people []string) {
	__cgo__section_name := (*C.gchar)(unsafe.Pointer(C.CString(section_name)))
	__header__people := (*reflect.SliceHeader)(unsafe.Pointer(&people))
	C.gtk_about_dialog_add_credit_section(self.CPointer, __cgo__section_name, (**C.gchar)(unsafe.Pointer(__header__people.Data)))
	C.free(unsafe.Pointer(__cgo__section_name))
	return
}

/*
Returns the string which are displayed in the artists tab
of the secondary credits dialog.
*/
func (self *_TraitAboutDialog) GetArtists() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_artists(self.CPointer)
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Returns the string which are displayed in the authors tab
of the secondary credits dialog.
*/
func (self *_TraitAboutDialog) GetAuthors() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_authors(self.CPointer)
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Returns the comments string.
*/
func (self *_TraitAboutDialog) GetComments() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_comments(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the copyright string.
*/
func (self *_TraitAboutDialog) GetCopyright() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_copyright(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the string which are displayed in the documenters
tab of the secondary credits dialog.
*/
func (self *_TraitAboutDialog) GetDocumenters() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_documenters(self.CPointer)
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Returns the license information.
*/
func (self *_TraitAboutDialog) GetLicense() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_license(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the license set using gtk_about_dialog_set_license_type()
*/
func (self *_TraitAboutDialog) GetLicenseType() (return__ C.GtkLicense) {
	return__ = C.gtk_about_dialog_get_license_type(self.CPointer)
	return
}

/*
Returns the pixbuf displayed as logo in the about dialog.
*/
func (self *_TraitAboutDialog) GetLogo() (return__ *C.GdkPixbuf) {
	return__ = C.gtk_about_dialog_get_logo(self.CPointer)
	return
}

/*
Returns the icon name displayed as logo in the about dialog.
*/
func (self *_TraitAboutDialog) GetLogoIconName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_logo_icon_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the program name displayed in the about dialog.
*/
func (self *_TraitAboutDialog) GetProgramName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_program_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the translator credits string which is displayed
in the translators tab of the secondary credits dialog.
*/
func (self *_TraitAboutDialog) GetTranslatorCredits() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_translator_credits(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the version string.
*/
func (self *_TraitAboutDialog) GetVersion() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_version(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the website URL.
*/
func (self *_TraitAboutDialog) GetWebsite() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_website(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the label used for the website link.
*/
func (self *_TraitAboutDialog) GetWebsiteLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_about_dialog_get_website_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns whether the license text in @about is
automatically wrapped.
*/
func (self *_TraitAboutDialog) GetWrapLicense() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_about_dialog_get_wrap_license(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the strings which are displayed in the artists tab
of the secondary credits dialog.
*/
func (self *_TraitAboutDialog) SetArtists(artists []string) {
	__header__artists := (*reflect.SliceHeader)(unsafe.Pointer(&artists))
	C.gtk_about_dialog_set_artists(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__artists.Data)))
	return
}

/*
Sets the strings which are displayed in the authors tab
of the secondary credits dialog.
*/
func (self *_TraitAboutDialog) SetAuthors(authors []string) {
	__header__authors := (*reflect.SliceHeader)(unsafe.Pointer(&authors))
	C.gtk_about_dialog_set_authors(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__authors.Data)))
	return
}

/*
Sets the comments string to display in the about dialog.
This should be a short string of one or two lines.
*/
func (self *_TraitAboutDialog) SetComments(comments string) {
	__cgo__comments := (*C.gchar)(unsafe.Pointer(C.CString(comments)))
	C.gtk_about_dialog_set_comments(self.CPointer, __cgo__comments)
	C.free(unsafe.Pointer(__cgo__comments))
	return
}

/*
Sets the copyright string to display in the about dialog.
This should be a short string of one or two lines.
*/
func (self *_TraitAboutDialog) SetCopyright(copyright string) {
	__cgo__copyright := (*C.gchar)(unsafe.Pointer(C.CString(copyright)))
	C.gtk_about_dialog_set_copyright(self.CPointer, __cgo__copyright)
	C.free(unsafe.Pointer(__cgo__copyright))
	return
}

/*
Sets the strings which are displayed in the documenters tab
of the secondary credits dialog.
*/
func (self *_TraitAboutDialog) SetDocumenters(documenters []string) {
	__header__documenters := (*reflect.SliceHeader)(unsafe.Pointer(&documenters))
	C.gtk_about_dialog_set_documenters(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__documenters.Data)))
	return
}

/*
Sets the license information to be displayed in the secondary
license dialog. If @license is %NULL, the license button is
hidden.
*/
func (self *_TraitAboutDialog) SetLicense(license string) {
	__cgo__license := (*C.gchar)(unsafe.Pointer(C.CString(license)))
	C.gtk_about_dialog_set_license(self.CPointer, __cgo__license)
	C.free(unsafe.Pointer(__cgo__license))
	return
}

/*
Sets the license of the application showing the @about dialog from a
list of known licenses.

This function overrides the license set using
gtk_about_dialog_set_license().
*/
func (self *_TraitAboutDialog) SetLicenseType(license_type C.GtkLicense) {
	C.gtk_about_dialog_set_license_type(self.CPointer, license_type)
	return
}

/*
Sets the pixbuf to be displayed as logo in the about dialog.
If it is %NULL, the default window icon set with
gtk_window_set_default_icon() will be used.
*/
func (self *_TraitAboutDialog) SetLogo(logo *C.GdkPixbuf) {
	C.gtk_about_dialog_set_logo(self.CPointer, logo)
	return
}

/*
Sets the pixbuf to be displayed as logo in the about dialog.
If it is %NULL, the default window icon set with
gtk_window_set_default_icon() will be used.
*/
func (self *_TraitAboutDialog) SetLogoIconName(icon_name string) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_about_dialog_set_logo_icon_name(self.CPointer, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

/*
Sets the name to display in the about dialog.
If this is not set, it defaults to g_get_application_name().
*/
func (self *_TraitAboutDialog) SetProgramName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_about_dialog_set_program_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Sets the translator credits string which is displayed in
the translators tab of the secondary credits dialog.

The intended use for this string is to display the translator
of the language which is currently used in the user interface.
Using gettext(), a simple way to achieve that is to mark the
string for translation:
|[<!-- language="C" -->
 gtk_about_dialog_set_translator_credits (about,
                                          _("translator-credits"));
]|
It is a good idea to use the customary msgid “translator-credits” for this
purpose, since translators will already know the purpose of that msgid, and
since #GtkAboutDialog will detect if “translator-credits” is untranslated
and hide the tab.
*/
func (self *_TraitAboutDialog) SetTranslatorCredits(translator_credits string) {
	__cgo__translator_credits := (*C.gchar)(unsafe.Pointer(C.CString(translator_credits)))
	C.gtk_about_dialog_set_translator_credits(self.CPointer, __cgo__translator_credits)
	C.free(unsafe.Pointer(__cgo__translator_credits))
	return
}

/*
Sets the version string to display in the about dialog.
*/
func (self *_TraitAboutDialog) SetVersion(version string) {
	__cgo__version := (*C.gchar)(unsafe.Pointer(C.CString(version)))
	C.gtk_about_dialog_set_version(self.CPointer, __cgo__version)
	C.free(unsafe.Pointer(__cgo__version))
	return
}

/*
Sets the URL to use for the website link.
*/
func (self *_TraitAboutDialog) SetWebsite(website string) {
	__cgo__website := (*C.gchar)(unsafe.Pointer(C.CString(website)))
	C.gtk_about_dialog_set_website(self.CPointer, __cgo__website)
	C.free(unsafe.Pointer(__cgo__website))
	return
}

/*
Sets the label to be used for the website link.
*/
func (self *_TraitAboutDialog) SetWebsiteLabel(website_label string) {
	__cgo__website_label := (*C.gchar)(unsafe.Pointer(C.CString(website_label)))
	C.gtk_about_dialog_set_website_label(self.CPointer, __cgo__website_label)
	C.free(unsafe.Pointer(__cgo__website_label))
	return
}

/*
Sets whether the license text in @about is
automatically wrapped.
*/
func (self *_TraitAboutDialog) SetWrapLicense(wrap_license bool) {
	__cgo__wrap_license := C.gboolean(0)
	if wrap_license {
		__cgo__wrap_license = C.gboolean(1)
	}
	C.gtk_about_dialog_set_wrap_license(self.CPointer, __cgo__wrap_license)
	return
}

type _TraitAccelGroup struct{ CPointer *C.GtkAccelGroup }

/*
Finds the first accelerator in @accel_group that matches
@accel_key and @accel_mods, and activates it.
*/
func (self *_TraitAccelGroup) Activate(accel_quark C.GQuark, acceleratable *C.GObject, accel_key uint, accel_mods C.GdkModifierType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_accel_group_activate(self.CPointer, accel_quark, acceleratable, C.guint(accel_key), accel_mods)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Installs an accelerator in this group. When @accel_group is being
activated in response to a call to gtk_accel_groups_activate(),
@closure will be invoked if the @accel_key and @accel_mods from
gtk_accel_groups_activate() match those of this connection.

The signature used for the @closure is that of #GtkAccelGroupActivate.

Note that, due to implementation details, a single closure can
only be connected to one accelerator group.
*/
func (self *_TraitAccelGroup) Connect(accel_key uint, accel_mods C.GdkModifierType, accel_flags C.GtkAccelFlags, closure *C.GClosure) {
	C.gtk_accel_group_connect(self.CPointer, C.guint(accel_key), accel_mods, accel_flags, closure)
	return
}

/*
Installs an accelerator in this group, using an accelerator path
to look up the appropriate key and modifiers (see
gtk_accel_map_add_entry()). When @accel_group is being activated
in response to a call to gtk_accel_groups_activate(), @closure will
be invoked if the @accel_key and @accel_mods from
gtk_accel_groups_activate() match the key and modifiers for the path.

The signature used for the @closure is that of #GtkAccelGroupActivate.

Note that @accel_path string will be stored in a #GQuark. Therefore,
if you pass a static string, you can save some memory by interning it
first with g_intern_static_string().
*/
func (self *_TraitAccelGroup) ConnectByPath(accel_path string, closure *C.GClosure) {
	__cgo__accel_path := (*C.gchar)(unsafe.Pointer(C.CString(accel_path)))
	C.gtk_accel_group_connect_by_path(self.CPointer, __cgo__accel_path, closure)
	C.free(unsafe.Pointer(__cgo__accel_path))
	return
}

/*
Removes an accelerator previously installed through
gtk_accel_group_connect().

Since 2.20 @closure can be %NULL.
*/
func (self *_TraitAccelGroup) Disconnect(closure *C.GClosure) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_accel_group_disconnect(self.CPointer, closure)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes an accelerator previously installed through
gtk_accel_group_connect().
*/
func (self *_TraitAccelGroup) DisconnectKey(accel_key uint, accel_mods C.GdkModifierType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_accel_group_disconnect_key(self.CPointer, C.guint(accel_key), accel_mods)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Finds the first entry in an accelerator group for which
@find_func returns %TRUE and returns its #GtkAccelKey.
*/
func (self *_TraitAccelGroup) Find(find_func C.GtkAccelGroupFindFunc, data unsafe.Pointer) (return__ *AccelKey) {
	var __cgo__return__ *C.GtkAccelKey
	__cgo__return__ = C.gtk_accel_group_find(self.CPointer, find_func, (C.gpointer)(data))
	return__ = (*AccelKey)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Locks are added and removed using gtk_accel_group_lock() and
gtk_accel_group_unlock().
*/
func (self *_TraitAccelGroup) GetIsLocked() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_accel_group_get_is_locked(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a #GdkModifierType representing the mask for this
@accel_group. For example, #GDK_CONTROL_MASK, #GDK_SHIFT_MASK, etc.
*/
func (self *_TraitAccelGroup) GetModifierMask() (return__ C.GdkModifierType) {
	return__ = C.gtk_accel_group_get_modifier_mask(self.CPointer)
	return
}

/*
Locks the given accelerator group.

Locking an acelerator group prevents the accelerators contained
within it to be changed during runtime. Refer to
gtk_accel_map_change_entry() about runtime accelerator changes.

If called more than once, @accel_group remains locked until
gtk_accel_group_unlock() has been called an equivalent number
of times.
*/
func (self *_TraitAccelGroup) Lock() {
	C.gtk_accel_group_lock(self.CPointer)
	return
}

/*
Queries an accelerator group for all entries matching @accel_key
and @accel_mods.
*/
func (self *_TraitAccelGroup) Query(accel_key uint, accel_mods C.GdkModifierType) (n_entries uint, return__ *AccelGroupEntry) {
	var __cgo__n_entries C.guint
	var __cgo__return__ *C.GtkAccelGroupEntry
	__cgo__return__ = C.gtk_accel_group_query(self.CPointer, C.guint(accel_key), accel_mods, &__cgo__n_entries)
	n_entries = uint(__cgo__n_entries)
	return__ = (*AccelGroupEntry)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Undoes the last call to gtk_accel_group_lock() on this @accel_group.
*/
func (self *_TraitAccelGroup) Unlock() {
	C.gtk_accel_group_unlock(self.CPointer)
	return
}

type _TraitAccelLabel struct{ CPointer *C.GtkAccelLabel }

/*
Gets the keyval and modifier mask set with
gtk_accel_label_set_accel().
*/
func (self *_TraitAccelLabel) GetAccel() (accelerator_key uint, accelerator_mods C.GdkModifierType) {
	var __cgo__accelerator_key C.guint
	C.gtk_accel_label_get_accel(self.CPointer, &__cgo__accelerator_key, &accelerator_mods)
	accelerator_key = uint(__cgo__accelerator_key)
	return
}

/*
Fetches the widget monitored by this accelerator label. See
gtk_accel_label_set_accel_widget().
*/
func (self *_TraitAccelLabel) GetAccelWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_accel_label_get_accel_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the width needed to display the accelerator key(s).
This is used by menus to align all of the #GtkMenuItem widgets, and shouldn't
be needed by applications.
*/
func (self *_TraitAccelLabel) GetAccelWidth() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_accel_label_get_accel_width(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Recreates the string representing the accelerator keys.
This should not be needed since the string is automatically updated whenever
accelerators are added or removed from the associated widget.
*/
func (self *_TraitAccelLabel) Refetch() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_accel_label_refetch(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Manually sets a keyval and modifier mask as the accelerator rendered
by @accel_label.

If a keyval and modifier are explicitly set then these values are
used regardless of any associated accel closure or widget.

Providing an @accelerator_key of 0 removes the manual setting.
*/
func (self *_TraitAccelLabel) SetAccel(accelerator_key uint, accelerator_mods C.GdkModifierType) {
	C.gtk_accel_label_set_accel(self.CPointer, C.guint(accelerator_key), accelerator_mods)
	return
}

/*
Sets the closure to be monitored by this accelerator label. The closure
must be connected to an accelerator group; see gtk_accel_group_connect().
*/
func (self *_TraitAccelLabel) SetAccelClosure(accel_closure *C.GClosure) {
	C.gtk_accel_label_set_accel_closure(self.CPointer, accel_closure)
	return
}

/*
Sets the widget to be monitored by this accelerator label.
*/
func (self *_TraitAccelLabel) SetAccelWidget(accel_widget *Widget) {
	C.gtk_accel_label_set_accel_widget(self.CPointer, (*C.GtkWidget)(accel_widget.CPointer))
	return
}

type _TraitAccelMap struct{ CPointer *C.GtkAccelMap }

type _TraitAccessible struct{ CPointer *C.GtkAccessible }

// gtk_accessible_connect_widget_destroyed is not generated due to deprecation attr

/*
Gets the #GtkWidget corresponding to the #GtkAccessible.
The returned widget does not have a reference added, so
you do not need to unref it.
*/
func (self *_TraitAccessible) GetWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_accessible_get_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Sets the #GtkWidget corresponding to the #GtkAccessible.

@accessible will not hold a reference to @widget.
It is the caller’s responsibility to ensure that when @widget
is destroyed, the widget is unset by calling this function
again with @widget set to %NULL.
*/
func (self *_TraitAccessible) SetWidget(widget *Widget) {
	C.gtk_accessible_set_widget(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

type _TraitAction struct{ CPointer *C.GtkAction }

// gtk_action_activate is not generated due to deprecation attr

// gtk_action_block_activate is not generated due to deprecation attr

// gtk_action_connect_accelerator is not generated due to deprecation attr

// gtk_action_create_icon is not generated due to deprecation attr

// gtk_action_create_menu is not generated due to deprecation attr

// gtk_action_create_menu_item is not generated due to deprecation attr

// gtk_action_create_tool_item is not generated due to deprecation attr

// gtk_action_disconnect_accelerator is not generated due to deprecation attr

// gtk_action_get_accel_closure is not generated due to deprecation attr

// gtk_action_get_accel_path is not generated due to deprecation attr

// gtk_action_get_always_show_image is not generated due to deprecation attr

// gtk_action_get_gicon is not generated due to deprecation attr

// gtk_action_get_icon_name is not generated due to deprecation attr

// gtk_action_get_is_important is not generated due to deprecation attr

// gtk_action_get_label is not generated due to deprecation attr

// gtk_action_get_name is not generated due to deprecation attr

// gtk_action_get_proxies is not generated due to deprecation attr

// gtk_action_get_sensitive is not generated due to deprecation attr

// gtk_action_get_short_label is not generated due to deprecation attr

// gtk_action_get_stock_id is not generated due to deprecation attr

// gtk_action_get_tooltip is not generated due to deprecation attr

// gtk_action_get_visible is not generated due to deprecation attr

// gtk_action_get_visible_horizontal is not generated due to deprecation attr

// gtk_action_get_visible_vertical is not generated due to deprecation attr

// gtk_action_is_sensitive is not generated due to deprecation attr

// gtk_action_is_visible is not generated due to deprecation attr

// gtk_action_set_accel_group is not generated due to deprecation attr

// gtk_action_set_accel_path is not generated due to deprecation attr

// gtk_action_set_always_show_image is not generated due to deprecation attr

// gtk_action_set_gicon is not generated due to deprecation attr

// gtk_action_set_icon_name is not generated due to deprecation attr

// gtk_action_set_is_important is not generated due to deprecation attr

// gtk_action_set_label is not generated due to deprecation attr

// gtk_action_set_sensitive is not generated due to deprecation attr

// gtk_action_set_short_label is not generated due to deprecation attr

// gtk_action_set_stock_id is not generated due to deprecation attr

// gtk_action_set_tooltip is not generated due to deprecation attr

// gtk_action_set_visible is not generated due to deprecation attr

// gtk_action_set_visible_horizontal is not generated due to deprecation attr

// gtk_action_set_visible_vertical is not generated due to deprecation attr

// gtk_action_unblock_activate is not generated due to deprecation attr

type _TraitActionBar struct{ CPointer *C.GtkActionBar }

/*
Retrieves the center bar widget of the bar.
*/
func (self *_TraitActionBar) GetCenterWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_action_bar_get_center_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Adds @child to @action_bar, packed with reference to the
end of the @action_bar.
*/
func (self *_TraitActionBar) PackEnd(child *Widget) {
	C.gtk_action_bar_pack_end(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Adds @child to @action_bar, packed with reference to the
start of the @action_bar.
*/
func (self *_TraitActionBar) PackStart(child *Widget) {
	C.gtk_action_bar_pack_start(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Sets the center widget for the #GtkActionBar.
*/
func (self *_TraitActionBar) SetCenterWidget(center_widget *Widget) {
	C.gtk_action_bar_set_center_widget(self.CPointer, (*C.GtkWidget)(center_widget.CPointer))
	return
}

type _TraitActionGroup struct{ CPointer *C.GtkActionGroup }

// gtk_action_group_add_action is not generated due to deprecation attr

// gtk_action_group_add_action_with_accel is not generated due to deprecation attr

// gtk_action_group_add_actions is not generated due to deprecation attr

// gtk_action_group_add_actions_full is not generated due to deprecation attr

// gtk_action_group_add_radio_actions is not generated due to deprecation attr

// gtk_action_group_add_radio_actions_full is not generated due to deprecation attr

// gtk_action_group_add_toggle_actions is not generated due to deprecation attr

// gtk_action_group_add_toggle_actions_full is not generated due to deprecation attr

// gtk_action_group_get_accel_group is not generated due to deprecation attr

// gtk_action_group_get_action is not generated due to deprecation attr

// gtk_action_group_get_name is not generated due to deprecation attr

// gtk_action_group_get_sensitive is not generated due to deprecation attr

// gtk_action_group_get_visible is not generated due to deprecation attr

// gtk_action_group_list_actions is not generated due to deprecation attr

// gtk_action_group_remove_action is not generated due to deprecation attr

// gtk_action_group_set_accel_group is not generated due to deprecation attr

// gtk_action_group_set_sensitive is not generated due to deprecation attr

// gtk_action_group_set_translate_func is not generated due to deprecation attr

// gtk_action_group_set_translation_domain is not generated due to deprecation attr

// gtk_action_group_set_visible is not generated due to deprecation attr

// gtk_action_group_translate_string is not generated due to deprecation attr

type _TraitAdjustment struct{ CPointer *C.GtkAdjustment }

/*
Emits a #GtkAdjustment::changed signal from the #GtkAdjustment.
This is typically called by the owner of the #GtkAdjustment after it has
changed any of the #GtkAdjustment properties other than the value.
*/
func (self *_TraitAdjustment) Changed() {
	C.gtk_adjustment_changed(self.CPointer)
	return
}

/*
Updates the #GtkAdjustment:value property to ensure that the range
between @lower and @upper is in the current page (i.e. between
#GtkAdjustment:value and #GtkAdjustment:value + #GtkAdjustment:page_size).
If the range is larger than the page size, then only the start of it will
be in the current page.
A #GtkAdjustment::changed signal will be emitted if the value is changed.
*/
func (self *_TraitAdjustment) ClampPage(lower float64, upper float64) {
	C.gtk_adjustment_clamp_page(self.CPointer, C.gdouble(lower), C.gdouble(upper))
	return
}

/*
Sets all properties of the adjustment at once.

Use this function to avoid multiple emissions of the #GtkAdjustment::changed
signal. See gtk_adjustment_set_lower() for an alternative way
of compressing multiple emissions of #GtkAdjustment::changed into one.
*/
func (self *_TraitAdjustment) Configure(value float64, lower float64, upper float64, step_increment float64, page_increment float64, page_size float64) {
	C.gtk_adjustment_configure(self.CPointer, C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size))
	return
}

/*
Retrieves the minimum value of the adjustment.
*/
func (self *_TraitAdjustment) GetLower() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_adjustment_get_lower(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the smaller of step increment and page increment.
*/
func (self *_TraitAdjustment) GetMinimumIncrement() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_adjustment_get_minimum_increment(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Retrieves the page increment of the adjustment.
*/
func (self *_TraitAdjustment) GetPageIncrement() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_adjustment_get_page_increment(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Retrieves the page size of the adjustment.
*/
func (self *_TraitAdjustment) GetPageSize() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_adjustment_get_page_size(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Retrieves the step increment of the adjustment.
*/
func (self *_TraitAdjustment) GetStepIncrement() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_adjustment_get_step_increment(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Retrieves the maximum value of the adjustment.
*/
func (self *_TraitAdjustment) GetUpper() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_adjustment_get_upper(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the current value of the adjustment. See
gtk_adjustment_set_value ().
*/
func (self *_TraitAdjustment) GetValue() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_adjustment_get_value(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Sets the minimum value of the adjustment.

When setting multiple adjustment properties via their individual
setters, multiple #GtkAdjustment::changed signals will be emitted. However, since
the emission of the #GtkAdjustment::changed signal is tied to the emission of the
#GObject::notify signals of the changed properties, it’s possible
to compress the #GtkAdjustment::changed signals into one by calling
g_object_freeze_notify() and g_object_thaw_notify() around the
calls to the individual setters.

Alternatively, using a single g_object_set() for all the properties
to change, or using gtk_adjustment_configure() has the same effect
of compressing #GtkAdjustment::changed emissions.
*/
func (self *_TraitAdjustment) SetLower(lower float64) {
	C.gtk_adjustment_set_lower(self.CPointer, C.gdouble(lower))
	return
}

/*
Sets the page increment of the adjustment.

See gtk_adjustment_set_lower() about how to compress multiple
emissions of the #GtkAdjustment::changed signal when setting multiple adjustment
properties.
*/
func (self *_TraitAdjustment) SetPageIncrement(page_increment float64) {
	C.gtk_adjustment_set_page_increment(self.CPointer, C.gdouble(page_increment))
	return
}

/*
Sets the page size of the adjustment.

See gtk_adjustment_set_lower() about how to compress multiple
emissions of the GtkAdjustment::changed signal when setting multiple adjustment
properties.
*/
func (self *_TraitAdjustment) SetPageSize(page_size float64) {
	C.gtk_adjustment_set_page_size(self.CPointer, C.gdouble(page_size))
	return
}

/*
Sets the step increment of the adjustment.

See gtk_adjustment_set_lower() about how to compress multiple
emissions of the #GtkAdjustment::changed signal when setting multiple adjustment
properties.
*/
func (self *_TraitAdjustment) SetStepIncrement(step_increment float64) {
	C.gtk_adjustment_set_step_increment(self.CPointer, C.gdouble(step_increment))
	return
}

/*
Sets the maximum value of the adjustment.

Note that values will be restricted by
`upper - page-size` if the page-size
property is nonzero.

See gtk_adjustment_set_lower() about how to compress multiple
emissions of the #GtkAdjustment::changed signal when setting multiple adjustment
properties.
*/
func (self *_TraitAdjustment) SetUpper(upper float64) {
	C.gtk_adjustment_set_upper(self.CPointer, C.gdouble(upper))
	return
}

/*
Sets the #GtkAdjustment value. The value is clamped to lie between
#GtkAdjustment:lower and #GtkAdjustment:upper.

Note that for adjustments which are used in a #GtkScrollbar, the effective
range of allowed values goes from #GtkAdjustment:lower to
#GtkAdjustment:upper - #GtkAdjustment:page_size.
*/
func (self *_TraitAdjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(self.CPointer, C.gdouble(value))
	return
}

/*
Emits a #GtkAdjustment::value_changed signal from the #GtkAdjustment.
This is typically called by the owner of the #GtkAdjustment after it has
changed the #GtkAdjustment:value property.
*/
func (self *_TraitAdjustment) ValueChanged() {
	C.gtk_adjustment_value_changed(self.CPointer)
	return
}

type _TraitAlignment struct{ CPointer *C.GtkAlignment }

/*
Gets the padding on the different sides of the widget.
See gtk_alignment_set_padding ().
*/
func (self *_TraitAlignment) GetPadding() (padding_top uint, padding_bottom uint, padding_left uint, padding_right uint) {
	var __cgo__padding_top C.guint
	var __cgo__padding_bottom C.guint
	var __cgo__padding_left C.guint
	var __cgo__padding_right C.guint
	C.gtk_alignment_get_padding(self.CPointer, &__cgo__padding_top, &__cgo__padding_bottom, &__cgo__padding_left, &__cgo__padding_right)
	padding_top = uint(__cgo__padding_top)
	padding_bottom = uint(__cgo__padding_bottom)
	padding_left = uint(__cgo__padding_left)
	padding_right = uint(__cgo__padding_right)
	return
}

/*
Sets the #GtkAlignment values.
*/
func (self *_TraitAlignment) Set(xalign float32, yalign float32, xscale float32, yscale float32) {
	C.gtk_alignment_set(self.CPointer, C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale), C.gfloat(yscale))
	return
}

/*
Sets the padding on the different sides of the widget.
The padding adds blank space to the sides of the widget. For instance,
this can be used to indent the child widget towards the right by adding
padding on the left.
*/
func (self *_TraitAlignment) SetPadding(padding_top uint, padding_bottom uint, padding_left uint, padding_right uint) {
	C.gtk_alignment_set_padding(self.CPointer, C.guint(padding_top), C.guint(padding_bottom), C.guint(padding_left), C.guint(padding_right))
	return
}

type _TraitAppChooserButton struct{ CPointer *C.GtkAppChooserButton }

/*
Appends a custom item to the list of applications that is shown
in the popup; the item name must be unique per-widget.
Clients can use the provided name as a detail for the
#GtkAppChooserButton::custom-item-activated signal, to add a
callback for the activation of a particular custom item in the list.
See also gtk_app_chooser_button_append_separator().
*/
func (self *_TraitAppChooserButton) AppendCustomItem(name string, label string, icon *C.GIcon) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_app_chooser_button_append_custom_item(self.CPointer, __cgo__name, __cgo__label, icon)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Appends a separator to the list of applications that is shown
in the popup.
*/
func (self *_TraitAppChooserButton) AppendSeparator() {
	C.gtk_app_chooser_button_append_separator(self.CPointer)
	return
}

/*
Returns the text to display at the top of the dialog.
*/
func (self *_TraitAppChooserButton) GetHeading() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_app_chooser_button_get_heading(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the current value of the #GtkAppChooserButton:show-default-item
property.
*/
func (self *_TraitAppChooserButton) GetShowDefaultItem() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_app_chooser_button_get_show_default_item(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current value of the #GtkAppChooserButton:show-dialog-item
property.
*/
func (self *_TraitAppChooserButton) GetShowDialogItem() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_app_chooser_button_get_show_dialog_item(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Selects a custom item previously added with
gtk_app_chooser_button_append_custom_item().

Use gtk_app_chooser_refresh() to bring the selection
to its initial state.
*/
func (self *_TraitAppChooserButton) SetActiveCustomItem(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_app_chooser_button_set_active_custom_item(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Sets the text to display at the top of the dialog.
If the heading is not set, the dialog displays a default text.
*/
func (self *_TraitAppChooserButton) SetHeading(heading string) {
	__cgo__heading := (*C.gchar)(unsafe.Pointer(C.CString(heading)))
	C.gtk_app_chooser_button_set_heading(self.CPointer, __cgo__heading)
	C.free(unsafe.Pointer(__cgo__heading))
	return
}

/*
Sets whether the dropdown menu of this button should show the
default application for the given content type at top.
*/
func (self *_TraitAppChooserButton) SetShowDefaultItem(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_app_chooser_button_set_show_default_item(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether the dropdown menu of this button should show an
entry to trigger a #GtkAppChooserDialog.
*/
func (self *_TraitAppChooserButton) SetShowDialogItem(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_app_chooser_button_set_show_dialog_item(self.CPointer, __cgo__setting)
	return
}

type _TraitAppChooserDialog struct{ CPointer *C.GtkAppChooserDialog }

/*
Returns the text to display at the top of the dialog.
*/
func (self *_TraitAppChooserDialog) GetHeading() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_app_chooser_dialog_get_heading(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the #GtkAppChooserWidget of this dialog.
*/
func (self *_TraitAppChooserDialog) GetWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_app_chooser_dialog_get_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Sets the text to display at the top of the dialog.
If the heading is not set, the dialog displays a default text.
*/
func (self *_TraitAppChooserDialog) SetHeading(heading string) {
	__cgo__heading := (*C.gchar)(unsafe.Pointer(C.CString(heading)))
	C.gtk_app_chooser_dialog_set_heading(self.CPointer, __cgo__heading)
	C.free(unsafe.Pointer(__cgo__heading))
	return
}

type _TraitAppChooserWidget struct{ CPointer *C.GtkAppChooserWidget }

/*
Returns the text that is shown if there are not applications
that can handle the content type.
*/
func (self *_TraitAppChooserWidget) GetDefaultText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_app_chooser_widget_get_default_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the current value of the #GtkAppChooserWidget:show-all
property.
*/
func (self *_TraitAppChooserWidget) GetShowAll() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_app_chooser_widget_get_show_all(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current value of the #GtkAppChooserWidget:show-default
property.
*/
func (self *_TraitAppChooserWidget) GetShowDefault() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_app_chooser_widget_get_show_default(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current value of the #GtkAppChooserWidget:show-fallback
property.
*/
func (self *_TraitAppChooserWidget) GetShowFallback() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_app_chooser_widget_get_show_fallback(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current value of the #GtkAppChooserWidget:show-other
property.
*/
func (self *_TraitAppChooserWidget) GetShowOther() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_app_chooser_widget_get_show_other(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current value of the #GtkAppChooserWidget:show-recommended
property.
*/
func (self *_TraitAppChooserWidget) GetShowRecommended() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_app_chooser_widget_get_show_recommended(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the text that is shown if there are not applications
that can handle the content type.
*/
func (self *_TraitAppChooserWidget) SetDefaultText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_app_chooser_widget_set_default_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Sets whether the app chooser should show all applications
in a flat list.
*/
func (self *_TraitAppChooserWidget) SetShowAll(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_app_chooser_widget_set_show_all(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether the app chooser should show the default handler
for the content type in a separate section.
*/
func (self *_TraitAppChooserWidget) SetShowDefault(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_app_chooser_widget_set_show_default(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether the app chooser should show related applications
for the content type in a separate section.
*/
func (self *_TraitAppChooserWidget) SetShowFallback(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_app_chooser_widget_set_show_fallback(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether the app chooser should show applications
which are unrelated to the content type.
*/
func (self *_TraitAppChooserWidget) SetShowOther(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_app_chooser_widget_set_show_other(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether the app chooser should show recommended applications
for the content type in a separate section.
*/
func (self *_TraitAppChooserWidget) SetShowRecommended(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_app_chooser_widget_set_show_recommended(self.CPointer, __cgo__setting)
	return
}

type _TraitApplication struct{ CPointer *C.GtkApplication }

/*
Installs an accelerator that will cause the named action
to be activated when the key combination specificed by @accelerator
is pressed.

@accelerator must be a string that can be parsed by gtk_accelerator_parse(),
e.g. "<Primary>q" or “<Control><Alt>p”.

@action_name must be the name of an action as it would be used
in the app menu, i.e. actions that have been added to the application
are referred to with an “app.” prefix, and window-specific actions
with a “win.” prefix.

GtkApplication also extracts accelerators out of “accel” attributes
in the #GMenuModels passed to gtk_application_set_app_menu() and
gtk_application_set_menubar(), which is usually more convenient
than calling this function for each accelerator.
*/
func (self *_TraitApplication) AddAccelerator(accelerator string, action_name string, parameter *C.GVariant) {
	__cgo__accelerator := (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	__cgo__action_name := (*C.gchar)(unsafe.Pointer(C.CString(action_name)))
	C.gtk_application_add_accelerator(self.CPointer, __cgo__accelerator, __cgo__action_name, parameter)
	C.free(unsafe.Pointer(__cgo__accelerator))
	C.free(unsafe.Pointer(__cgo__action_name))
	return
}

/*
Adds a window to @application.

This call is equivalent to setting the #GtkWindow:application
property of @window to @application.

Normally, the connection between the application and the window
will remain until the window is destroyed, but you can explicitly
remove it with gtk_application_remove_window().

GTK+ will keep the application running as long as it has
any windows.
*/
func (self *_TraitApplication) AddWindow(window *Window) {
	C.gtk_application_add_window(self.CPointer, (*C.GtkWindow)(window.CPointer))
	return
}

/*
Gets the accelerators that are currently associated with
the given action.
*/
func (self *_TraitApplication) GetAccelsForAction(detailed_action_name string) (return__ []string) {
	__cgo__detailed_action_name := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action_name)))
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.gtk_application_get_accels_for_action(self.CPointer, __cgo__detailed_action_name)
	C.free(unsafe.Pointer(__cgo__detailed_action_name))
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Gets the “active” window for the application.

The active window is the one that was most recently focused (within
the application).  This window may not have the focus at the moment
if another application has it -- this is just the most
recently-focused window within this application.
*/
func (self *_TraitApplication) GetActiveWindow() (return__ *Window) {
	var __cgo__return__ *C.GtkWindow
	__cgo__return__ = C.gtk_application_get_active_window(self.CPointer)
	return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the menu model that has been set with
gtk_application_set_app_menu().
*/
func (self *_TraitApplication) GetAppMenu() (return__ *C.GMenuModel) {
	return__ = C.gtk_application_get_app_menu(self.CPointer)
	return
}

/*
Returns the menu model that has been set with
gtk_application_set_menubar().
*/
func (self *_TraitApplication) GetMenubar() (return__ *C.GMenuModel) {
	return__ = C.gtk_application_get_menubar(self.CPointer)
	return
}

/*
Returns the #GtkApplicationWindow with the given ID.
*/
func (self *_TraitApplication) GetWindowById(id uint) (return__ *Window) {
	var __cgo__return__ *C.GtkWindow
	__cgo__return__ = C.gtk_application_get_window_by_id(self.CPointer, C.guint(id))
	return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets a list of the #GtkWindows associated with @application.

The list is sorted by most recently focused window, such that the first
element is the currently focused window. (Useful for choosing a parent
for a transient window.)

The list that is returned should not be modified in any way. It will
only remain valid until the next focus change or window creation or
deletion.
*/
func (self *_TraitApplication) GetWindows() (return__ *C.GList) {
	return__ = C.gtk_application_get_windows(self.CPointer)
	return
}

/*
Inform the session manager that certain types of actions should be
inhibited. This is not guaranteed to work on all platforms and for
all types of actions.

Applications should invoke this method when they begin an operation
that should not be interrupted, such as creating a CD or DVD. The
types of actions that may be blocked are specified by the @flags
parameter. When the application completes the operation it should
call gtk_application_uninhibit() to remove the inhibitor. Note that
an application can have multiple inhibitors, and all of the must
be individually removed. Inhibitors are also cleared when the
application exits.

Applications should not expect that they will always be able to block
the action. In most cases, users will be given the option to force
the action to take place.

Reasons should be short and to the point.

If @window is given, the session manager may point the user to
this window to find out more about why the action is inhibited.
*/
func (self *_TraitApplication) Inhibit(window *Window, flags C.GtkApplicationInhibitFlags, reason string) (return__ uint) {
	__cgo__reason := (*C.gchar)(unsafe.Pointer(C.CString(reason)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_application_inhibit(self.CPointer, (*C.GtkWindow)(window.CPointer), flags, __cgo__reason)
	C.free(unsafe.Pointer(__cgo__reason))
	return__ = uint(__cgo__return__)
	return
}

/*
Determines if any of the actions specified in @flags are
currently inhibited (possibly by another application).
*/
func (self *_TraitApplication) IsInhibited(flags C.GtkApplicationInhibitFlags) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_application_is_inhibited(self.CPointer, flags)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Lists the detailed action names which have associated accelerators.
See gtk_application_set_accels_for_action().
*/
func (self *_TraitApplication) ListActionDescriptions() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.gtk_application_list_action_descriptions(self.CPointer)
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Removes an accelerator that has been previously added
with gtk_application_add_accelerator().
*/
func (self *_TraitApplication) RemoveAccelerator(action_name string, parameter *C.GVariant) {
	__cgo__action_name := (*C.gchar)(unsafe.Pointer(C.CString(action_name)))
	C.gtk_application_remove_accelerator(self.CPointer, __cgo__action_name, parameter)
	C.free(unsafe.Pointer(__cgo__action_name))
	return
}

/*
Remove a window from @application.

If @window belongs to @application then this call is equivalent to
setting the #GtkWindow:application property of @window to
%NULL.

The application may stop running as a result of a call to this
function.
*/
func (self *_TraitApplication) RemoveWindow(window *Window) {
	C.gtk_application_remove_window(self.CPointer, (*C.GtkWindow)(window.CPointer))
	return
}

/*
Sets one or more keyboard accelerator that will trigger the
given action. The first item in @accels will be the primary
accelerator, which may be displayed in the UI.
*/
func (self *_TraitApplication) SetAccelsForAction(detailed_action_name string, accels []string) {
	__cgo__detailed_action_name := (*C.gchar)(unsafe.Pointer(C.CString(detailed_action_name)))
	__header__accels := (*reflect.SliceHeader)(unsafe.Pointer(&accels))
	C.gtk_application_set_accels_for_action(self.CPointer, __cgo__detailed_action_name, (**C.gchar)(unsafe.Pointer(__header__accels.Data)))
	C.free(unsafe.Pointer(__cgo__detailed_action_name))
	return
}

/*
Sets or unsets the application menu for @application.

This can only be done in the primary instance of the application,
after it has been registered.  #GApplication::startup is a good place
to call this.

The application menu is a single menu containing items that typically
impact the application as a whole, rather than acting on a specific
window or document.  For example, you would expect to see
“Preferences” or “Quit” in an application menu, but not “Save” or
“Print”.

If supported, the application menu will be rendered by the desktop
environment.

Use the base #GActionMap interface to add actions, to respond to the user
selecting these menu items.
*/
func (self *_TraitApplication) SetAppMenu(app_menu *C.GMenuModel) {
	C.gtk_application_set_app_menu(self.CPointer, app_menu)
	return
}

/*
Sets or unsets the menubar for windows of @application.

This is a menubar in the traditional sense.

This can only be done in the primary instance of the application,
after it has been registered.  #GApplication::startup is a good place
to call this.

Depending on the desktop environment, this may appear at the top of
each window, or at the top of the screen.  In some environments, if
both the application menu and the menubar are set, the application
menu will be presented as if it were the first item of the menubar.
Other environments treat the two as completely separate -- for
example, the application menu may be rendered by the desktop shell
while the menubar (if set) remains in each individual window.

Use the base #GActionMap interface to add actions, to respond to the user
selecting these menu items.
*/
func (self *_TraitApplication) SetMenubar(menubar *C.GMenuModel) {
	C.gtk_application_set_menubar(self.CPointer, menubar)
	return
}

/*
Removes an inhibitor that has been established with gtk_application_inhibit().
Inhibitors are also cleared when the application exits.
*/
func (self *_TraitApplication) Uninhibit(cookie uint) {
	C.gtk_application_uninhibit(self.CPointer, C.guint(cookie))
	return
}

type _TraitApplicationWindow struct{ CPointer *C.GtkApplicationWindow }

/*
Returns the unique ID of the window. If the window has not yet been added to
a #GtkApplication, returns `0`.
*/
func (self *_TraitApplicationWindow) GetId() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_application_window_get_id(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Returns whether the window will display a menubar for the app menu
and menubar as needed.
*/
func (self *_TraitApplicationWindow) GetShowMenubar() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_application_window_get_show_menubar(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets whether the window will display a menubar for the app menu
and menubar as needed.
*/
func (self *_TraitApplicationWindow) SetShowMenubar(show_menubar bool) {
	__cgo__show_menubar := C.gboolean(0)
	if show_menubar {
		__cgo__show_menubar = C.gboolean(1)
	}
	C.gtk_application_window_set_show_menubar(self.CPointer, __cgo__show_menubar)
	return
}

type _TraitArrow struct{ CPointer *C.GtkArrow }

/*
Sets the direction and style of the #GtkArrow, @arrow.
*/
func (self *_TraitArrow) Set(arrow_type C.GtkArrowType, shadow_type C.GtkShadowType) {
	C.gtk_arrow_set(self.CPointer, arrow_type, shadow_type)
	return
}

type _TraitAspectFrame struct{ CPointer *C.GtkAspectFrame }

/*
Set parameters for an existing #GtkAspectFrame.
*/
func (self *_TraitAspectFrame) Set(xalign float32, yalign float32, ratio float32, obey_child bool) {
	__cgo__obey_child := C.gboolean(0)
	if obey_child {
		__cgo__obey_child = C.gboolean(1)
	}
	C.gtk_aspect_frame_set(self.CPointer, C.gfloat(xalign), C.gfloat(yalign), C.gfloat(ratio), __cgo__obey_child)
	return
}

type _TraitAssistant struct{ CPointer *C.GtkAssistant }

/*
Adds a widget to the action area of a #GtkAssistant.
*/
func (self *_TraitAssistant) AddActionWidget(child *Widget) {
	C.gtk_assistant_add_action_widget(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Appends a page to the @assistant.
*/
func (self *_TraitAssistant) AppendPage(page *Widget) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_assistant_append_page(self.CPointer, (*C.GtkWidget)(page.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Erases the visited page history so the back button is not
shown on the current page, and removes the cancel button
from subsequent pages.

Use this when the information provided up to the current
page is hereafter deemed permanent and cannot be modified
or undone. For example, showing a progress page to track
a long-running, unreversible operation after the user has
clicked apply on a confirmation page.
*/
func (self *_TraitAssistant) Commit() {
	C.gtk_assistant_commit(self.CPointer)
	return
}

/*
Returns the page number of the current page.
*/
func (self *_TraitAssistant) GetCurrentPage() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_assistant_get_current_page(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the number of pages in the @assistant
*/
func (self *_TraitAssistant) GetNPages() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_assistant_get_n_pages(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the child widget contained in page number @page_num.
*/
func (self *_TraitAssistant) GetNthPage(page_num int) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_assistant_get_nth_page(self.CPointer, C.gint(page_num))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets whether @page is complete.
*/
func (self *_TraitAssistant) GetPageComplete(page *Widget) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_assistant_get_page_complete(self.CPointer, (*C.GtkWidget)(page.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_assistant_get_page_header_image is not generated due to deprecation attr

// gtk_assistant_get_page_side_image is not generated due to deprecation attr

/*
Gets the title for @page.
*/
func (self *_TraitAssistant) GetPageTitle(page *Widget) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_assistant_get_page_title(self.CPointer, (*C.GtkWidget)(page.CPointer))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the page type of @page.
*/
func (self *_TraitAssistant) GetPageType(page *Widget) (return__ C.GtkAssistantPageType) {
	return__ = C.gtk_assistant_get_page_type(self.CPointer, (*C.GtkWidget)(page.CPointer))
	return
}

/*
Inserts a page in the @assistant at a given position.
*/
func (self *_TraitAssistant) InsertPage(page *Widget, position int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_assistant_insert_page(self.CPointer, (*C.GtkWidget)(page.CPointer), C.gint(position))
	return__ = int(__cgo__return__)
	return
}

/*
Navigate to the next page.

It is a programming error to call this function when
there is no next page.

This function is for use when creating pages of the
#GTK_ASSISTANT_PAGE_CUSTOM type.
*/
func (self *_TraitAssistant) NextPage() {
	C.gtk_assistant_next_page(self.CPointer)
	return
}

/*
Prepends a page to the @assistant.
*/
func (self *_TraitAssistant) PrependPage(page *Widget) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_assistant_prepend_page(self.CPointer, (*C.GtkWidget)(page.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Navigate to the previous visited page.

It is a programming error to call this function when
no previous page is available.

This function is for use when creating pages of the
#GTK_ASSISTANT_PAGE_CUSTOM type.
*/
func (self *_TraitAssistant) PreviousPage() {
	C.gtk_assistant_previous_page(self.CPointer)
	return
}

/*
Removes a widget from the action area of a #GtkAssistant.
*/
func (self *_TraitAssistant) RemoveActionWidget(child *Widget) {
	C.gtk_assistant_remove_action_widget(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Removes the @page_num’s page from @assistant.
*/
func (self *_TraitAssistant) RemovePage(page_num int) {
	C.gtk_assistant_remove_page(self.CPointer, C.gint(page_num))
	return
}

/*
Switches the page to @page_num.

Note that this will only be necessary in custom buttons,
as the @assistant flow can be set with
gtk_assistant_set_forward_page_func().
*/
func (self *_TraitAssistant) SetCurrentPage(page_num int) {
	C.gtk_assistant_set_current_page(self.CPointer, C.gint(page_num))
	return
}

/*
Sets the page forwarding function to be @page_func.

This function will be used to determine what will be
the next page when the user presses the forward button.
Setting @page_func to %NULL will make the assistant to
use the default forward function, which just goes to the
next visible page.
*/
func (self *_TraitAssistant) SetForwardPageFunc(page_func C.GtkAssistantPageFunc, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_assistant_set_forward_page_func(self.CPointer, page_func, (C.gpointer)(data), destroy)
	return
}

/*
Sets whether @page contents are complete.

This will make @assistant update the buttons state
to be able to continue the task.
*/
func (self *_TraitAssistant) SetPageComplete(page *Widget, complete bool) {
	__cgo__complete := C.gboolean(0)
	if complete {
		__cgo__complete = C.gboolean(1)
	}
	C.gtk_assistant_set_page_complete(self.CPointer, (*C.GtkWidget)(page.CPointer), __cgo__complete)
	return
}

// gtk_assistant_set_page_header_image is not generated due to deprecation attr

// gtk_assistant_set_page_side_image is not generated due to deprecation attr

/*
Sets a title for @page.

The title is displayed in the header area of the assistant
when @page is the current page.
*/
func (self *_TraitAssistant) SetPageTitle(page *Widget, title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_assistant_set_page_title(self.CPointer, (*C.GtkWidget)(page.CPointer), __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
Sets the page type for @page.

The page type determines the page behavior in the @assistant.
*/
func (self *_TraitAssistant) SetPageType(page *Widget, type_ C.GtkAssistantPageType) {
	C.gtk_assistant_set_page_type(self.CPointer, (*C.GtkWidget)(page.CPointer), type_)
	return
}

/*
Forces @assistant to recompute the buttons state.

GTK+ automatically takes care of this in most situations,
e.g. when the user goes to a different page, or when the
visibility or completeness of a page changes.

One situation where it can be necessary to call this
function is when changing a value on the current page
affects the future page flow of the assistant.
*/
func (self *_TraitAssistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(self.CPointer)
	return
}

type _TraitBin struct{ CPointer *C.GtkBin }

/*
Gets the child of the #GtkBin, or %NULL if the bin contains
no child widget. The returned widget does not have a reference
added, so you do not need to unref it.
*/
func (self *_TraitBin) GetChild() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_bin_get_child(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

type _TraitBox struct{ CPointer *C.GtkBox }

/*
Gets the value set by gtk_box_set_baseline_position().
*/
func (self *_TraitBox) GetBaselinePosition() (return__ C.GtkBaselinePosition) {
	return__ = C.gtk_box_get_baseline_position(self.CPointer)
	return
}

/*
Retrieves the center widget of the box.
*/
func (self *_TraitBox) GetCenterWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_box_get_center_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether the box is homogeneous (all children are the
same size). See gtk_box_set_homogeneous().
*/
func (self *_TraitBox) GetHomogeneous() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_box_get_homogeneous(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_box_set_spacing().
*/
func (self *_TraitBox) GetSpacing() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_box_get_spacing(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Adds @child to @box, packed with reference to the end of @box.
The @child is packed after (away from end of) any other child
packed with reference to the end of @box.
*/
func (self *_TraitBox) PackEnd(child *Widget, expand bool, fill bool, padding uint) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	__cgo__fill := C.gboolean(0)
	if fill {
		__cgo__fill = C.gboolean(1)
	}
	C.gtk_box_pack_end(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__expand, __cgo__fill, C.guint(padding))
	return
}

/*
Adds @child to @box, packed with reference to the start of @box.
The @child is packed after any other child packed with reference
to the start of @box.
*/
func (self *_TraitBox) PackStart(child *Widget, expand bool, fill bool, padding uint) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	__cgo__fill := C.gboolean(0)
	if fill {
		__cgo__fill = C.gboolean(1)
	}
	C.gtk_box_pack_start(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__expand, __cgo__fill, C.guint(padding))
	return
}

/*
Obtains information about how @child is packed into @box.
*/
func (self *_TraitBox) QueryChildPacking(child *Widget) (expand bool, fill bool, padding uint, pack_type C.GtkPackType) {
	var __cgo__expand C.gboolean
	var __cgo__fill C.gboolean
	var __cgo__padding C.guint
	C.gtk_box_query_child_packing(self.CPointer, (*C.GtkWidget)(child.CPointer), &__cgo__expand, &__cgo__fill, &__cgo__padding, &pack_type)
	expand = __cgo__expand == C.gboolean(1)
	fill = __cgo__fill == C.gboolean(1)
	padding = uint(__cgo__padding)
	return
}

/*
Moves @child to a new @position in the list of @box children.
The list contains widgets packed #GTK_PACK_START
as well as widgets packed #GTK_PACK_END, in the order that these
widgets were added to @box.

A widget’s position in the @box children list determines where
the widget is packed into @box.  A child widget at some position
in the list will be packed just after all other widgets of the
same packing type that appear earlier in the list.
*/
func (self *_TraitBox) ReorderChild(child *Widget, position int) {
	C.gtk_box_reorder_child(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(position))
	return
}

/*
Sets the baseline position of a box. This affects
only horizontal boxes with at least one baseline aligned
child. If there is more vertical space available than requested,
and the baseline is not allocated by the parent then
@position is used to allocate the baseline wrt the
extra space available.
*/
func (self *_TraitBox) SetBaselinePosition(position C.GtkBaselinePosition) {
	C.gtk_box_set_baseline_position(self.CPointer, position)
	return
}

/*
Sets a center widget; that is a child widget that will be
centered with respect to the full width of the box, even
if the children at either side take up different amounts
of space.
*/
func (self *_TraitBox) SetCenterWidget(widget *Widget) {
	C.gtk_box_set_center_widget(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

/*
Sets the way @child is packed into @box.
*/
func (self *_TraitBox) SetChildPacking(child *Widget, expand bool, fill bool, padding uint, pack_type C.GtkPackType) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	__cgo__fill := C.gboolean(0)
	if fill {
		__cgo__fill = C.gboolean(1)
	}
	C.gtk_box_set_child_packing(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__expand, __cgo__fill, C.guint(padding), pack_type)
	return
}

/*
Sets the #GtkBox:homogeneous property of @box, controlling
whether or not all children of @box are given equal space
in the box.
*/
func (self *_TraitBox) SetHomogeneous(homogeneous bool) {
	__cgo__homogeneous := C.gboolean(0)
	if homogeneous {
		__cgo__homogeneous = C.gboolean(1)
	}
	C.gtk_box_set_homogeneous(self.CPointer, __cgo__homogeneous)
	return
}

/*
Sets the #GtkBox:spacing property of @box, which is the
number of pixels to place between children of @box.
*/
func (self *_TraitBox) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(self.CPointer, C.gint(spacing))
	return
}

type _TraitBuilder struct{ CPointer *C.GtkBuilder }

/*
Adds the @callback_symbol to the scope of @builder under the given @callback_name.

Using this function overrides the behavior of gtk_builder_connect_signals()
for any callback symbols that are added. Using this method allows for better
encapsulation as it does not require that callback symbols be declared in
the global namespace.
*/
func (self *_TraitBuilder) AddCallbackSymbol(callback_name string, callback_symbol C.GCallback) {
	__cgo__callback_name := (*C.gchar)(unsafe.Pointer(C.CString(callback_name)))
	C.gtk_builder_add_callback_symbol(self.CPointer, __cgo__callback_name, callback_symbol)
	C.free(unsafe.Pointer(__cgo__callback_name))
	return
}

// gtk_builder_add_callback_symbols is not generated due to varargs

/*
Parses a file containing a [GtkBuilder UI definition][BUILDER-UI]
and merges it with the current contents of @builder.

Most users will probably want to use gtk_builder_new_from_file().

Upon errors 0 will be returned and @error will be assigned a
#GError from the #GTK_BUILDER_ERROR, #G_MARKUP_ERROR or #G_FILE_ERROR
domain.

It’s not really reasonable to attempt to handle failures of this
call.  You should not use this function with untrusted files (ie:
files that are not part of your application).  Broken #GtkBuilder
files can easily crash your program, and it’s possible that memory
was leaked leading up to the reported failure.  The only reasonable
thing to do when an error is detected is to call g_error().
*/
func (self *_TraitBuilder) AddFromFile(filename string) (return__ uint, __err__ error) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_builder_add_from_file(self.CPointer, __cgo__filename, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Parses a resource file containing a [GtkBuilder UI definition][BUILDER-UI]
and merges it with the current contents of @builder.

Most users will probably want to use gtk_builder_new_from_resource().

Upon errors 0 will be returned and @error will be assigned a
#GError from the #GTK_BUILDER_ERROR, #G_MARKUP_ERROR or #G_RESOURCE_ERROR
domain.

It’s not really reasonable to attempt to handle failures of this
call.  The only reasonable thing to do when an error is detected is
to call g_error().
*/
func (self *_TraitBuilder) AddFromResource(resource_path string) (return__ uint, __err__ error) {
	__cgo__resource_path := (*C.gchar)(unsafe.Pointer(C.CString(resource_path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_builder_add_from_resource(self.CPointer, __cgo__resource_path, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__resource_path))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Parses a string containing a [GtkBuilder UI definition][BUILDER-UI]
and merges it with the current contents of @builder.

Most users will probably want to use gtk_builder_new_from_string().

Upon errors 0 will be returned and @error will be assigned a
#GError from the #GTK_BUILDER_ERROR or #G_MARKUP_ERROR domain.

It’s not really reasonable to attempt to handle failures of this
call.  The only reasonable thing to do when an error is detected is
to call g_error().
*/
func (self *_TraitBuilder) AddFromString(buffer string, length int64) (return__ uint, __err__ error) {
	__cgo__buffer := (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_builder_add_from_string(self.CPointer, __cgo__buffer, C.gsize(length), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__buffer))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Parses a file containing a [GtkBuilder UI definition][BUILDER-UI]
building only the requested objects and merges
them with the current contents of @builder.

Upon errors 0 will be returned and @error will be assigned a
#GError from the #GTK_BUILDER_ERROR, #G_MARKUP_ERROR or #G_FILE_ERROR
domain.

If you are adding an object that depends on an object that is not
its child (for instance a #GtkTreeView that depends on its
#GtkTreeModel), you have to explicitly list all of them in @object_ids.
*/
func (self *_TraitBuilder) AddObjectsFromFile(filename string, object_ids []string) (return__ uint, __err__ error) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	__header__object_ids := (*reflect.SliceHeader)(unsafe.Pointer(&object_ids))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_builder_add_objects_from_file(self.CPointer, __cgo__filename, (**C.gchar)(unsafe.Pointer(__header__object_ids.Data)), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Parses a resource file containing a [GtkBuilder UI definition][BUILDER-UI]
building only the requested objects and merges
them with the current contents of @builder.

Upon errors 0 will be returned and @error will be assigned a
#GError from the #GTK_BUILDER_ERROR, #G_MARKUP_ERROR or #G_RESOURCE_ERROR
domain.

If you are adding an object that depends on an object that is not
its child (for instance a #GtkTreeView that depends on its
#GtkTreeModel), you have to explicitly list all of them in @object_ids.
*/
func (self *_TraitBuilder) AddObjectsFromResource(resource_path string, object_ids []string) (return__ uint, __err__ error) {
	__cgo__resource_path := (*C.gchar)(unsafe.Pointer(C.CString(resource_path)))
	__header__object_ids := (*reflect.SliceHeader)(unsafe.Pointer(&object_ids))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_builder_add_objects_from_resource(self.CPointer, __cgo__resource_path, (**C.gchar)(unsafe.Pointer(__header__object_ids.Data)), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__resource_path))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Parses a string containing a [GtkBuilder UI definition][BUILDER-UI]
building only the requested objects and merges
them with the current contents of @builder.

Upon errors 0 will be returned and @error will be assigned a
#GError from the #GTK_BUILDER_ERROR or #G_MARKUP_ERROR domain.

If you are adding an object that depends on an object that is not
its child (for instance a #GtkTreeView that depends on its
#GtkTreeModel), you have to explicitly list all of them in @object_ids.
*/
func (self *_TraitBuilder) AddObjectsFromString(buffer string, length int64, object_ids []string) (return__ uint, __err__ error) {
	__cgo__buffer := (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	__header__object_ids := (*reflect.SliceHeader)(unsafe.Pointer(&object_ids))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_builder_add_objects_from_string(self.CPointer, __cgo__buffer, C.gsize(length), (**C.gchar)(unsafe.Pointer(__header__object_ids.Data)), &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__buffer))
	return__ = uint(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This method is a simpler variation of gtk_builder_connect_signals_full().
It uses symbols explicitly added to @builder with prior calls to
gtk_builder_add_callback_symbol(). In the case that symbols are not
explicitly added; it uses #GModule’s introspective features (by opening the module %NULL)
to look at the application’s symbol table. From here it tries to match
the signal handler names given in the interface description with
symbols in the application and connects the signals. Note that this
function can only be called once, subsequent calls will do nothing.

Note that unless gtk_builder_add_callback_symbol() is called for
all signal callbacks which are referenced by the loaded XML, this
function will require that #GModule be supported on the platform.

If you rely on #GModule support to lookup callbacks in the symbol table,
the following details should be noted:

When compiling applications for Windows, you must declare signal callbacks
with #G_MODULE_EXPORT, or they will not be put in the symbol table.
On Linux and Unices, this is not necessary; applications should instead
be compiled with the -Wl,--export-dynamic CFLAGS, and linked against
gmodule-export-2.0.
*/
func (self *_TraitBuilder) ConnectSignals(user_data unsafe.Pointer) {
	C.gtk_builder_connect_signals(self.CPointer, (C.gpointer)(user_data))
	return
}

/*
This function can be thought of the interpreted language binding
version of gtk_builder_connect_signals(), except that it does not
require GModule to function correctly.
*/
func (self *_TraitBuilder) ConnectSignalsFull(func_ C.GtkBuilderConnectFunc, user_data unsafe.Pointer) {
	C.gtk_builder_connect_signals_full(self.CPointer, func_, (C.gpointer)(user_data))
	return
}

/*
Add @object to the @builder object pool so it can be referenced just like any
other object built by builder.
*/
func (self *_TraitBuilder) ExposeObject(name string, object *C.GObject) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_builder_expose_object(self.CPointer, __cgo__name, object)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Gets the #GtkApplication associated with the builder.

The #GtkApplication is used for creating action proxies as requested
from XML that the builder is loading.

By default, the builder uses the default application: the one from
g_application_get_default().  If you want to use another application
for constructing proxies, use gtk_builder_set_application().
*/
func (self *_TraitBuilder) GetApplication() (return__ *Application) {
	var __cgo__return__ *C.GtkApplication
	__cgo__return__ = C.gtk_builder_get_application(self.CPointer)
	return__ = NewApplicationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the object named @name. Note that this function does not
increment the reference count of the returned object.
*/
func (self *_TraitBuilder) GetObject(name string) (return__ *C.GObject) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.gtk_builder_get_object(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Gets all objects that have been constructed by @builder. Note that
this function does not increment the reference counts of the returned
objects.
*/
func (self *_TraitBuilder) GetObjects() (return__ *C.GSList) {
	return__ = C.gtk_builder_get_objects(self.CPointer)
	return
}

/*
Gets the translation domain of @builder.
*/
func (self *_TraitBuilder) GetTranslationDomain() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_builder_get_translation_domain(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Looks up a type by name, using the virtual function that
#GtkBuilder has for that purpose. This is mainly used when
implementing the #GtkBuildable interface on a type.
*/
func (self *_TraitBuilder) GetTypeFromName(type_name string) (return__ C.GType) {
	__cgo__type_name := C.CString(type_name)
	return__ = C.gtk_builder_get_type_from_name(self.CPointer, __cgo__type_name)
	C.free(unsafe.Pointer(__cgo__type_name))
	return
}

/*
Fetches a symbol previously added to @builder
with gtk_builder_add_callback_symbols()

This function is intended for possible use in language bindings
or for any case that one might be cusomizing signal connections
using gtk_builder_connect_signals_full()
*/
func (self *_TraitBuilder) LookupCallbackSymbol(callback_name string) (return__ C.GCallback) {
	__cgo__callback_name := (*C.gchar)(unsafe.Pointer(C.CString(callback_name)))
	return__ = C.gtk_builder_lookup_callback_symbol(self.CPointer, __cgo__callback_name)
	C.free(unsafe.Pointer(__cgo__callback_name))
	return
}

/*
Sets the application associated with @builder.

You only need this function if there is more than one #GApplication
in your process.  @application cannot be %NULL.
*/
func (self *_TraitBuilder) SetApplication(application *Application) {
	C.gtk_builder_set_application(self.CPointer, (*C.GtkApplication)(application.CPointer))
	return
}

/*
Sets the translation domain of @builder.
See #GtkBuilder:translation-domain.
*/
func (self *_TraitBuilder) SetTranslationDomain(domain string) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	C.gtk_builder_set_translation_domain(self.CPointer, __cgo__domain)
	C.free(unsafe.Pointer(__cgo__domain))
	return
}

/*
This function demarshals a value from a string. This function
calls g_value_init() on the @value argument, so it need not be
initialised beforehand.

This function can handle char, uchar, boolean, int, uint, long,
ulong, enum, flags, float, double, string, #GdkColor, #GdkRGBA and
#GtkAdjustment type values. Support for #GtkWidget type values is
still to come.

Upon errors %FALSE will be returned and @error will be assigned a
#GError from the #GTK_BUILDER_ERROR domain.
*/
func (self *_TraitBuilder) ValueFromString(pspec *C.GParamSpec, string_ string) (value C.GValue, return__ bool, __err__ error) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_builder_value_from_string(self.CPointer, pspec, __cgo__string_, &value, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Like gtk_builder_value_from_string(), this function demarshals
a value from a string, but takes a #GType instead of #GParamSpec.
This function calls g_value_init() on the @value argument, so it
need not be initialised beforehand.

Upon errors %FALSE will be returned and @error will be assigned a
#GError from the #GTK_BUILDER_ERROR domain.
*/
func (self *_TraitBuilder) ValueFromStringType(type_ C.GType, string_ string) (value C.GValue, return__ bool, __err__ error) {
	__cgo__string_ := (*C.gchar)(unsafe.Pointer(C.CString(string_)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_builder_value_from_string_type(self.CPointer, type_, __cgo__string_, &value, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__string_))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type _TraitButton struct{ CPointer *C.GtkButton }

/*
Emits a #GtkButton::clicked signal to the given #GtkButton.
*/
func (self *_TraitButton) Clicked() {
	C.gtk_button_clicked(self.CPointer)
	return
}

// gtk_button_enter is not generated due to deprecation attr

/*
Gets the alignment of the child in the button.
*/
func (self *_TraitButton) GetAlignment() (xalign float32, yalign float32) {
	var __cgo__xalign C.gfloat
	var __cgo__yalign C.gfloat
	C.gtk_button_get_alignment(self.CPointer, &__cgo__xalign, &__cgo__yalign)
	xalign = float32(__cgo__xalign)
	yalign = float32(__cgo__yalign)
	return
}

/*
Returns whether the button will ignore the #GtkSettings:gtk-button-images
setting and always show the image, if available.
*/
func (self *_TraitButton) GetAlwaysShowImage() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_button_get_always_show_image(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the button’s event window if it is realized, %NULL otherwise.
This function should be rarely needed.
*/
func (self *_TraitButton) GetEventWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_button_get_event_window(self.CPointer)
	return
}

/*
Returns whether the button grabs focus when it is clicked with the mouse.
See gtk_button_set_focus_on_click().
*/
func (self *_TraitButton) GetFocusOnClick() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_button_get_focus_on_click(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the widget that is currenty set as the image of @button.
This may have been explicitly set by gtk_button_set_image()
or constructed by gtk_button_new_from_stock().
*/
func (self *_TraitButton) GetImage() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_button_get_image(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the position of the image relative to the text
inside the button.
*/
func (self *_TraitButton) GetImagePosition() (return__ C.GtkPositionType) {
	return__ = C.gtk_button_get_image_position(self.CPointer)
	return
}

/*
Fetches the text from the label of the button, as set by
gtk_button_set_label(). If the label text has not
been set the return value will be %NULL. This will be the
case if you create an empty button with gtk_button_new() to
use as a container.
*/
func (self *_TraitButton) GetLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_button_get_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the current relief style of the given #GtkButton.
*/
func (self *_TraitButton) GetRelief() (return__ C.GtkReliefStyle) {
	return__ = C.gtk_button_get_relief(self.CPointer)
	return
}

// gtk_button_get_use_stock is not generated due to deprecation attr

/*
Returns whether an embedded underline in the button label indicates a
mnemonic. See gtk_button_set_use_underline ().
*/
func (self *_TraitButton) GetUseUnderline() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_button_get_use_underline(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_button_leave is not generated due to deprecation attr

// gtk_button_pressed is not generated due to deprecation attr

// gtk_button_released is not generated due to deprecation attr

/*
Sets the alignment of the child. This property has no effect unless
the child is a #GtkMisc or a #GtkAlignment.
*/
func (self *_TraitButton) SetAlignment(xalign float32, yalign float32) {
	C.gtk_button_set_alignment(self.CPointer, C.gfloat(xalign), C.gfloat(yalign))
	return
}

/*
If %TRUE, the button will ignore the #GtkSettings:gtk-button-images
setting and always show the image, if available.

Use this property if the button  would be useless or hard to use
without the image.
*/
func (self *_TraitButton) SetAlwaysShowImage(always_show bool) {
	__cgo__always_show := C.gboolean(0)
	if always_show {
		__cgo__always_show = C.gboolean(1)
	}
	C.gtk_button_set_always_show_image(self.CPointer, __cgo__always_show)
	return
}

/*
Sets whether the button will grab focus when it is clicked with the mouse.
Making mouse clicks not grab focus is useful in places like toolbars where
you don’t want the keyboard focus removed from the main area of the
application.
*/
func (self *_TraitButton) SetFocusOnClick(focus_on_click bool) {
	__cgo__focus_on_click := C.gboolean(0)
	if focus_on_click {
		__cgo__focus_on_click = C.gboolean(1)
	}
	C.gtk_button_set_focus_on_click(self.CPointer, __cgo__focus_on_click)
	return
}

/*
Set the image of @button to the given widget. The image will be
displayed if the label text is %NULL or if
#GtkButton:always-show-image is %TRUE. You don’t have to call
gtk_widget_show() on @image yourself.
*/
func (self *_TraitButton) SetImage(image *Widget) {
	C.gtk_button_set_image(self.CPointer, (*C.GtkWidget)(image.CPointer))
	return
}

/*
Sets the position of the image relative to the text
inside the button.
*/
func (self *_TraitButton) SetImagePosition(position C.GtkPositionType) {
	C.gtk_button_set_image_position(self.CPointer, position)
	return
}

/*
Sets the text of the label of the button to @str. This text is
also used to select the stock item if gtk_button_set_use_stock()
is used.

This will also clear any previously set labels.
*/
func (self *_TraitButton) SetLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_button_set_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Sets the relief style of the edges of the given #GtkButton widget.
Three styles exist, GTK_RELIEF_NORMAL, GTK_RELIEF_HALF, GTK_RELIEF_NONE.
The default style is, as one can guess, GTK_RELIEF_NORMAL.
*/
func (self *_TraitButton) SetRelief(newstyle C.GtkReliefStyle) {
	C.gtk_button_set_relief(self.CPointer, newstyle)
	return
}

// gtk_button_set_use_stock is not generated due to deprecation attr

/*
If true, an underline in the text of the button label indicates
the next character should be used for the mnemonic accelerator key.
*/
func (self *_TraitButton) SetUseUnderline(use_underline bool) {
	__cgo__use_underline := C.gboolean(0)
	if use_underline {
		__cgo__use_underline = C.gboolean(1)
	}
	C.gtk_button_set_use_underline(self.CPointer, __cgo__use_underline)
	return
}

type _TraitButtonBox struct{ CPointer *C.GtkButtonBox }

/*
Returns whether the child is exempted from homogenous
sizing.
*/
func (self *_TraitButtonBox) GetChildNonHomogeneous(child *Widget) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_button_box_get_child_non_homogeneous(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether @child should appear in a secondary group of children.
*/
func (self *_TraitButtonBox) GetChildSecondary(child *Widget) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_button_box_get_child_secondary(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the method being used to arrange the buttons in a button box.
*/
func (self *_TraitButtonBox) GetLayout() (return__ C.GtkButtonBoxStyle) {
	return__ = C.gtk_button_box_get_layout(self.CPointer)
	return
}

/*
Sets whether the child is exempted from homogeous sizing.
*/
func (self *_TraitButtonBox) SetChildNonHomogeneous(child *Widget, non_homogeneous bool) {
	__cgo__non_homogeneous := C.gboolean(0)
	if non_homogeneous {
		__cgo__non_homogeneous = C.gboolean(1)
	}
	C.gtk_button_box_set_child_non_homogeneous(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__non_homogeneous)
	return
}

/*
Sets whether @child should appear in a secondary group of children.
A typical use of a secondary child is the help button in a dialog.

This group appears after the other children if the style
is %GTK_BUTTONBOX_START, %GTK_BUTTONBOX_SPREAD or
%GTK_BUTTONBOX_EDGE, and before the other children if the style
is %GTK_BUTTONBOX_END. For horizontal button boxes, the definition
of before/after depends on direction of the widget (see
gtk_widget_set_direction()). If the style is %GTK_BUTTONBOX_START
or %GTK_BUTTONBOX_END, then the secondary children are aligned at
the other end of the button box from the main children. For the
other styles, they appear immediately next to the main children.
*/
func (self *_TraitButtonBox) SetChildSecondary(child *Widget, is_secondary bool) {
	__cgo__is_secondary := C.gboolean(0)
	if is_secondary {
		__cgo__is_secondary = C.gboolean(1)
	}
	C.gtk_button_box_set_child_secondary(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__is_secondary)
	return
}

/*
Changes the way buttons are arranged in their container.
*/
func (self *_TraitButtonBox) SetLayout(layout_style C.GtkButtonBoxStyle) {
	C.gtk_button_box_set_layout(self.CPointer, layout_style)
	return
}

type _TraitCalendar struct{ CPointer *C.GtkCalendar }

/*
Remove all visual markers.
*/
func (self *_TraitCalendar) ClearMarks() {
	C.gtk_calendar_clear_marks(self.CPointer)
	return
}

/*
Obtains the selected date from a #GtkCalendar.
*/
func (self *_TraitCalendar) GetDate() (year uint, month uint, day uint) {
	var __cgo__year C.guint
	var __cgo__month C.guint
	var __cgo__day C.guint
	C.gtk_calendar_get_date(self.CPointer, &__cgo__year, &__cgo__month, &__cgo__day)
	year = uint(__cgo__year)
	month = uint(__cgo__month)
	day = uint(__cgo__day)
	return
}

/*
Returns if the @day of the @calendar is already marked.
*/
func (self *_TraitCalendar) GetDayIsMarked(day uint) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_calendar_get_day_is_marked(self.CPointer, C.guint(day))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Queries the height of detail cells, in rows.
See #GtkCalendar:detail-width-chars.
*/
func (self *_TraitCalendar) GetDetailHeightRows() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_calendar_get_detail_height_rows(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Queries the width of detail cells, in characters.
See #GtkCalendar:detail-width-chars.
*/
func (self *_TraitCalendar) GetDetailWidthChars() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_calendar_get_detail_width_chars(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the current display options of @calendar.
*/
func (self *_TraitCalendar) GetDisplayOptions() (return__ C.GtkCalendarDisplayOptions) {
	return__ = C.gtk_calendar_get_display_options(self.CPointer)
	return
}

/*
Places a visual marker on a particular day.
*/
func (self *_TraitCalendar) MarkDay(day uint) {
	C.gtk_calendar_mark_day(self.CPointer, C.guint(day))
	return
}

/*
Selects a day from the current month.
*/
func (self *_TraitCalendar) SelectDay(day uint) {
	C.gtk_calendar_select_day(self.CPointer, C.guint(day))
	return
}

/*
Shifts the calendar to a different month.
*/
func (self *_TraitCalendar) SelectMonth(month uint, year uint) {
	C.gtk_calendar_select_month(self.CPointer, C.guint(month), C.guint(year))
	return
}

/*
Installs a function which provides Pango markup with detail information
for each day. Examples for such details are holidays or appointments. That
information is shown below each day when #GtkCalendar:show-details is set.
A tooltip containing with full detail information is provided, if the entire
text should not fit into the details area, or if #GtkCalendar:show-details
is not set.

The size of the details area can be restricted by setting the
#GtkCalendar:detail-width-chars and #GtkCalendar:detail-height-rows
properties.
*/
func (self *_TraitCalendar) SetDetailFunc(func_ C.GtkCalendarDetailFunc, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_calendar_set_detail_func(self.CPointer, func_, (C.gpointer)(data), destroy)
	return
}

/*
Updates the height of detail cells.
See #GtkCalendar:detail-height-rows.
*/
func (self *_TraitCalendar) SetDetailHeightRows(rows int) {
	C.gtk_calendar_set_detail_height_rows(self.CPointer, C.gint(rows))
	return
}

/*
Updates the width of detail cells.
See #GtkCalendar:detail-width-chars.
*/
func (self *_TraitCalendar) SetDetailWidthChars(chars int) {
	C.gtk_calendar_set_detail_width_chars(self.CPointer, C.gint(chars))
	return
}

/*
Sets display options (whether to display the heading and the month
headings).
*/
func (self *_TraitCalendar) SetDisplayOptions(flags C.GtkCalendarDisplayOptions) {
	C.gtk_calendar_set_display_options(self.CPointer, flags)
	return
}

/*
Removes the visual marker from a particular day.
*/
func (self *_TraitCalendar) UnmarkDay(day uint) {
	C.gtk_calendar_unmark_day(self.CPointer, C.guint(day))
	return
}

type _TraitCellArea struct{ CPointer *C.GtkCellArea }

/*
Activates @area, usually by activating the currently focused
cell, however some subclasses which embed widgets in the area
can also activate a widget if it currently has the focus.
*/
func (self *_TraitCellArea) Activate(context *CellAreaContext, widget *Widget, cell_area *C.GdkRectangle, flags C.GtkCellRendererState, edit_only bool) (return__ bool) {
	__cgo__edit_only := C.gboolean(0)
	if edit_only {
		__cgo__edit_only = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_area_activate(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), cell_area, flags, __cgo__edit_only)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This is used by #GtkCellArea subclasses when handling events
to activate cells, the base #GtkCellArea class activates cells
for keyboard events for free in its own GtkCellArea->activate()
implementation.
*/
func (self *_TraitCellArea) ActivateCell(widget *Widget, renderer *CellRenderer, event *C.GdkEvent, cell_area *C.GdkRectangle, flags C.GtkCellRendererState) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_area_activate_cell(self.CPointer, (*C.GtkWidget)(widget.CPointer), (*C.GtkCellRenderer)(renderer.CPointer), event, cell_area, flags)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Adds @renderer to @area with the default child cell properties.
*/
func (self *_TraitCellArea) Add(renderer *CellRenderer) {
	C.gtk_cell_area_add(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer))
	return
}

/*
Adds @sibling to @renderer’s focusable area, focus will be drawn
around @renderer and all of its siblings if @renderer can
focus for a given row.

Events handled by focus siblings can also activate the given
focusable @renderer.
*/
func (self *_TraitCellArea) AddFocusSibling(renderer *CellRenderer, sibling *CellRenderer) {
	C.gtk_cell_area_add_focus_sibling(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), (*C.GtkCellRenderer)(sibling.CPointer))
	return
}

// gtk_cell_area_add_with_properties is not generated due to varargs

/*
Applies any connected attributes to the renderers in
@area by pulling the values from @tree_model.
*/
func (self *_TraitCellArea) ApplyAttributes(tree_model *C.GtkTreeModel, iter *TreeIter, is_expander bool, is_expanded bool) {
	__cgo__is_expander := C.gboolean(0)
	if is_expander {
		__cgo__is_expander = C.gboolean(1)
	}
	__cgo__is_expanded := C.gboolean(0)
	if is_expanded {
		__cgo__is_expanded = C.gboolean(1)
	}
	C.gtk_cell_area_apply_attributes(self.CPointer, tree_model, (*C.GtkTreeIter)(unsafe.Pointer(iter)), __cgo__is_expander, __cgo__is_expanded)
	return
}

/*
Connects an @attribute to apply values from @column for the
#GtkTreeModel in use.
*/
func (self *_TraitCellArea) AttributeConnect(renderer *CellRenderer, attribute string, column int) {
	__cgo__attribute := (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	C.gtk_cell_area_attribute_connect(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), __cgo__attribute, C.gint(column))
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Disconnects @attribute for the @renderer in @area so that
attribute will no longer be updated with values from the
model.
*/
func (self *_TraitCellArea) AttributeDisconnect(renderer *CellRenderer, attribute string) {
	__cgo__attribute := (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	C.gtk_cell_area_attribute_disconnect(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), __cgo__attribute)
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

// gtk_cell_area_cell_get is not generated due to varargs

/*
Gets the value of a cell property for @renderer in @area.
*/
func (self *_TraitCellArea) CellGetProperty(renderer *CellRenderer, property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.gtk_cell_area_cell_get_property(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

// gtk_cell_area_cell_get_valist is not generated due to varargs

// gtk_cell_area_cell_set is not generated due to varargs

/*
Sets a cell property for @renderer in @area.
*/
func (self *_TraitCellArea) CellSetProperty(renderer *CellRenderer, property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.gtk_cell_area_cell_set_property(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

// gtk_cell_area_cell_set_valist is not generated due to varargs

/*
This is sometimes needed for cases where rows need to share
alignments in one orientation but may be separately grouped
in the opposing orientation.

For instance, #GtkIconView creates all icons (rows) to have
the same width and the cells theirin to have the same
horizontal alignments. However each row of icons may have
a separate collective height. #GtkIconView uses this to
request the heights of each row based on a context which
was already used to request all the row widths that are
to be displayed.
*/
func (self *_TraitCellArea) CopyContext(context *CellAreaContext) (return__ *CellAreaContext) {
	var __cgo__return__ *C.GtkCellAreaContext
	__cgo__return__ = C.gtk_cell_area_copy_context(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer))
	return__ = NewCellAreaContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Creates a #GtkCellAreaContext to be used with @area for
all purposes. #GtkCellAreaContext stores geometry information
for rows for which it was operated on, it is important to use
the same context for the same row of data at all times (i.e.
one should render and handle events with the same #GtkCellAreaContext
which was used to request the size of those rows of data).
*/
func (self *_TraitCellArea) CreateContext() (return__ *CellAreaContext) {
	var __cgo__return__ *C.GtkCellAreaContext
	__cgo__return__ = C.gtk_cell_area_create_context(self.CPointer)
	return__ = NewCellAreaContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Delegates event handling to a #GtkCellArea.
*/
func (self *_TraitCellArea) Event(context *CellAreaContext, widget *Widget, event *C.GdkEvent, cell_area *C.GdkRectangle, flags C.GtkCellRendererState) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_cell_area_event(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), event, cell_area, flags)
	return__ = int(__cgo__return__)
	return
}

/*
This should be called by the @area’s owning layout widget
when focus is to be passed to @area, or moved within @area
for a given @direction and row data.

Implementing #GtkCellArea classes should implement this
method to receive and navigate focus in its own way particular
to how it lays out cells.
*/
func (self *_TraitCellArea) Focus(direction C.GtkDirectionType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_area_focus(self.CPointer, direction)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Calls @callback for every #GtkCellRenderer in @area.
*/
func (self *_TraitCellArea) Foreach(callback C.GtkCellCallback, callback_data unsafe.Pointer) {
	C.gtk_cell_area_foreach(self.CPointer, callback, (C.gpointer)(callback_data))
	return
}

/*
Calls @callback for every #GtkCellRenderer in @area with the
allocated rectangle inside @cell_area.
*/
func (self *_TraitCellArea) ForeachAlloc(context *CellAreaContext, widget *Widget, cell_area *C.GdkRectangle, background_area *C.GdkRectangle, callback C.GtkCellAllocCallback, callback_data unsafe.Pointer) {
	C.gtk_cell_area_foreach_alloc(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), cell_area, background_area, callback, (C.gpointer)(callback_data))
	return
}

/*
Derives the allocation of @renderer inside @area if @area
were to be renderered in @cell_area.
*/
func (self *_TraitCellArea) GetCellAllocation(context *CellAreaContext, widget *Widget, renderer *CellRenderer, cell_area *C.GdkRectangle) (allocation C.GdkRectangle) {
	C.gtk_cell_area_get_cell_allocation(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), (*C.GtkCellRenderer)(renderer.CPointer), cell_area, &allocation)
	return
}

/*
Gets the #GtkCellRenderer at @x and @y coordinates inside @area and optionally
returns the full cell allocation for it inside @cell_area.
*/
func (self *_TraitCellArea) GetCellAtPosition(context *CellAreaContext, widget *Widget, cell_area *C.GdkRectangle, x int, y int) (alloc_area C.GdkRectangle, return__ *CellRenderer) {
	var __cgo__return__ *C.GtkCellRenderer
	__cgo__return__ = C.gtk_cell_area_get_cell_at_position(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), cell_area, C.gint(x), C.gint(y), &alloc_area)
	return__ = NewCellRendererFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the current #GtkTreePath string for the currently
applied #GtkTreeIter, this is implicitly updated when
gtk_cell_area_apply_attributes() is called and can be
used to interact with renderers from #GtkCellArea
subclasses.
*/
func (self *_TraitCellArea) GetCurrentPathString() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_cell_area_get_current_path_string(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #GtkCellEditable widget currently used
to edit the currently edited cell.
*/
func (self *_TraitCellArea) GetEditWidget() (return__ *C.GtkCellEditable) {
	return__ = C.gtk_cell_area_get_edit_widget(self.CPointer)
	return
}

/*
Gets the #GtkCellRenderer in @area that is currently
being edited.
*/
func (self *_TraitCellArea) GetEditedCell() (return__ *CellRenderer) {
	var __cgo__return__ *C.GtkCellRenderer
	__cgo__return__ = C.gtk_cell_area_get_edited_cell(self.CPointer)
	return__ = NewCellRendererFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the currently focused cell for @area
*/
func (self *_TraitCellArea) GetFocusCell() (return__ *CellRenderer) {
	var __cgo__return__ *C.GtkCellRenderer
	__cgo__return__ = C.gtk_cell_area_get_focus_cell(self.CPointer)
	return__ = NewCellRendererFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the #GtkCellRenderer which is expected to be focusable
for which @renderer is, or may be a sibling.

This is handy for #GtkCellArea subclasses when handling events,
after determining the renderer at the event location it can
then chose to activate the focus cell for which the event
cell may have been a sibling.
*/
func (self *_TraitCellArea) GetFocusFromSibling(renderer *CellRenderer) (return__ *CellRenderer) {
	var __cgo__return__ *C.GtkCellRenderer
	__cgo__return__ = C.gtk_cell_area_get_focus_from_sibling(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer))
	return__ = NewCellRendererFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the focus sibling cell renderers for @renderer.
*/
func (self *_TraitCellArea) GetFocusSiblings(renderer *CellRenderer) (return__ *C.GList) {
	return__ = C.gtk_cell_area_get_focus_siblings(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer))
	return
}

/*
Retrieves a cell area’s initial minimum and natural height.

@area will store some geometrical information in @context along the way;
when requesting sizes over an arbitrary number of rows, it’s not important
to check the @minimum_height and @natural_height of this call but rather to
consult gtk_cell_area_context_get_preferred_height() after a series of
requests.
*/
func (self *_TraitCellArea) GetPreferredHeight(context *CellAreaContext, widget *Widget) (minimum_height int, natural_height int) {
	var __cgo__minimum_height C.gint
	var __cgo__natural_height C.gint
	C.gtk_cell_area_get_preferred_height(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), &__cgo__minimum_height, &__cgo__natural_height)
	minimum_height = int(__cgo__minimum_height)
	natural_height = int(__cgo__natural_height)
	return
}

/*
Retrieves a cell area’s minimum and natural height if it would be given
the specified @width.

@area stores some geometrical information in @context along the way
while calling gtk_cell_area_get_preferred_width(). It’s important to
perform a series of gtk_cell_area_get_preferred_width() requests with
@context first and then call gtk_cell_area_get_preferred_height_for_width()
on each cell area individually to get the height for width of each
fully requested row.

If at some point, the width of a single row changes, it should be
requested with gtk_cell_area_get_preferred_width() again and then
the full width of the requested rows checked again with
gtk_cell_area_context_get_preferred_width().
*/
func (self *_TraitCellArea) GetPreferredHeightForWidth(context *CellAreaContext, widget *Widget, width int) (minimum_height int, natural_height int) {
	var __cgo__minimum_height C.gint
	var __cgo__natural_height C.gint
	C.gtk_cell_area_get_preferred_height_for_width(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), C.gint(width), &__cgo__minimum_height, &__cgo__natural_height)
	minimum_height = int(__cgo__minimum_height)
	natural_height = int(__cgo__natural_height)
	return
}

/*
Retrieves a cell area’s initial minimum and natural width.

@area will store some geometrical information in @context along the way;
when requesting sizes over an arbitrary number of rows, it’s not important
to check the @minimum_width and @natural_width of this call but rather to
consult gtk_cell_area_context_get_preferred_width() after a series of
requests.
*/
func (self *_TraitCellArea) GetPreferredWidth(context *CellAreaContext, widget *Widget) (minimum_width int, natural_width int) {
	var __cgo__minimum_width C.gint
	var __cgo__natural_width C.gint
	C.gtk_cell_area_get_preferred_width(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), &__cgo__minimum_width, &__cgo__natural_width)
	minimum_width = int(__cgo__minimum_width)
	natural_width = int(__cgo__natural_width)
	return
}

/*
Retrieves a cell area’s minimum and natural width if it would be given
the specified @height.

@area stores some geometrical information in @context along the way
while calling gtk_cell_area_get_preferred_height(). It’s important to
perform a series of gtk_cell_area_get_preferred_height() requests with
@context first and then call gtk_cell_area_get_preferred_width_for_height()
on each cell area individually to get the height for width of each
fully requested row.

If at some point, the height of a single row changes, it should be
requested with gtk_cell_area_get_preferred_height() again and then
the full height of the requested rows checked again with
gtk_cell_area_context_get_preferred_height().
*/
func (self *_TraitCellArea) GetPreferredWidthForHeight(context *CellAreaContext, widget *Widget, height int) (minimum_width int, natural_width int) {
	var __cgo__minimum_width C.gint
	var __cgo__natural_width C.gint
	C.gtk_cell_area_get_preferred_width_for_height(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), C.gint(height), &__cgo__minimum_width, &__cgo__natural_width)
	minimum_width = int(__cgo__minimum_width)
	natural_width = int(__cgo__natural_width)
	return
}

/*
Gets whether the area prefers a height-for-width layout
or a width-for-height layout.
*/
func (self *_TraitCellArea) GetRequestMode() (return__ C.GtkSizeRequestMode) {
	return__ = C.gtk_cell_area_get_request_mode(self.CPointer)
	return
}

/*
Checks if @area contains @renderer.
*/
func (self *_TraitCellArea) HasRenderer(renderer *CellRenderer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_area_has_renderer(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This is a convenience function for #GtkCellArea implementations
to get the inner area where a given #GtkCellRenderer will be
rendered. It removes any padding previously added by gtk_cell_area_request_renderer().
*/
func (self *_TraitCellArea) InnerCellArea(widget *Widget, cell_area *C.GdkRectangle) (inner_area C.GdkRectangle) {
	C.gtk_cell_area_inner_cell_area(self.CPointer, (*C.GtkWidget)(widget.CPointer), cell_area, &inner_area)
	return
}

/*
Returns whether the area can do anything when activated,
after applying new attributes to @area.
*/
func (self *_TraitCellArea) IsActivatable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_area_is_activatable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether @sibling is one of @renderer’s focus siblings
(see gtk_cell_area_add_focus_sibling()).
*/
func (self *_TraitCellArea) IsFocusSibling(renderer *CellRenderer, sibling *CellRenderer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_area_is_focus_sibling(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), (*C.GtkCellRenderer)(sibling.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes @renderer from @area.
*/
func (self *_TraitCellArea) Remove(renderer *CellRenderer) {
	C.gtk_cell_area_remove(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer))
	return
}

/*
Removes @sibling from @renderer’s focus sibling list
(see gtk_cell_area_add_focus_sibling()).
*/
func (self *_TraitCellArea) RemoveFocusSibling(renderer *CellRenderer, sibling *CellRenderer) {
	C.gtk_cell_area_remove_focus_sibling(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), (*C.GtkCellRenderer)(sibling.CPointer))
	return
}

/*
Renders @area’s cells according to @area’s layout onto @widget at
the given coordinates.
*/
func (self *_TraitCellArea) Render(context *CellAreaContext, widget *Widget, cr *C.cairo_t, background_area *C.GdkRectangle, cell_area *C.GdkRectangle, flags C.GtkCellRendererState, paint_focus bool) {
	__cgo__paint_focus := C.gboolean(0)
	if paint_focus {
		__cgo__paint_focus = C.gboolean(1)
	}
	C.gtk_cell_area_render(self.CPointer, (*C.GtkCellAreaContext)(context.CPointer), (*C.GtkWidget)(widget.CPointer), cr, background_area, cell_area, flags, __cgo__paint_focus)
	return
}

/*
This is a convenience function for #GtkCellArea implementations
to request size for cell renderers. It’s important to use this
function to request size and then use gtk_cell_area_inner_cell_area()
at render and event time since this function will add padding
around the cell for focus painting.
*/
func (self *_TraitCellArea) RequestRenderer(renderer *CellRenderer, orientation C.GtkOrientation, widget *Widget, for_size int) (minimum_size int, natural_size int) {
	var __cgo__minimum_size C.gint
	var __cgo__natural_size C.gint
	C.gtk_cell_area_request_renderer(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), orientation, (*C.GtkWidget)(widget.CPointer), C.gint(for_size), &__cgo__minimum_size, &__cgo__natural_size)
	minimum_size = int(__cgo__minimum_size)
	natural_size = int(__cgo__natural_size)
	return
}

/*
Explicitly sets the currently focused cell to @renderer.

This is generally called by implementations of
#GtkCellAreaClass.focus() or #GtkCellAreaClass.event(),
however it can also be used to implement functions such
as gtk_tree_view_set_cursor_on_cell().
*/
func (self *_TraitCellArea) SetFocusCell(renderer *CellRenderer) {
	C.gtk_cell_area_set_focus_cell(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer))
	return
}

/*
Explicitly stops the editing of the currently edited cell.

If @canceled is %TRUE, the currently edited cell renderer
will emit the ::editing-canceled signal, otherwise the
the ::editing-done signal will be emitted on the current
edit widget.

See gtk_cell_area_get_edited_cell() and gtk_cell_area_get_edit_widget().
*/
func (self *_TraitCellArea) StopEditing(canceled bool) {
	__cgo__canceled := C.gboolean(0)
	if canceled {
		__cgo__canceled = C.gboolean(1)
	}
	C.gtk_cell_area_stop_editing(self.CPointer, __cgo__canceled)
	return
}

type _TraitCellAreaBox struct{ CPointer *C.GtkCellAreaBox }

/*
Gets the spacing added between cell renderers.
*/
func (self *_TraitCellAreaBox) GetSpacing() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_cell_area_box_get_spacing(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Adds @renderer to @box, packed with reference to the end of @box.

The @renderer is packed after (away from end of) any other
#GtkCellRenderer packed with reference to the end of @box.
*/
func (self *_TraitCellAreaBox) PackEnd(renderer *CellRenderer, expand bool, align bool, fixed bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	__cgo__align := C.gboolean(0)
	if align {
		__cgo__align = C.gboolean(1)
	}
	__cgo__fixed := C.gboolean(0)
	if fixed {
		__cgo__fixed = C.gboolean(1)
	}
	C.gtk_cell_area_box_pack_end(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), __cgo__expand, __cgo__align, __cgo__fixed)
	return
}

/*
Adds @renderer to @box, packed with reference to the start of @box.

The @renderer is packed after any other #GtkCellRenderer packed
with reference to the start of @box.
*/
func (self *_TraitCellAreaBox) PackStart(renderer *CellRenderer, expand bool, align bool, fixed bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	__cgo__align := C.gboolean(0)
	if align {
		__cgo__align = C.gboolean(1)
	}
	__cgo__fixed := C.gboolean(0)
	if fixed {
		__cgo__fixed = C.gboolean(1)
	}
	C.gtk_cell_area_box_pack_start(self.CPointer, (*C.GtkCellRenderer)(renderer.CPointer), __cgo__expand, __cgo__align, __cgo__fixed)
	return
}

/*
Sets the spacing to add between cell renderers in @box.
*/
func (self *_TraitCellAreaBox) SetSpacing(spacing int) {
	C.gtk_cell_area_box_set_spacing(self.CPointer, C.gint(spacing))
	return
}

type _TraitCellAreaContext struct{ CPointer *C.GtkCellAreaContext }

/*
Allocates a width and/or a height for all rows which are to be
rendered with @context.

Usually allocation is performed only horizontally or sometimes
vertically since a group of rows are usually rendered side by
side vertically or horizontally and share either the same width
or the same height. Sometimes they are allocated in both horizontal
and vertical orientations producing a homogeneous effect of the
rows. This is generally the case for #GtkTreeView when
#GtkTreeView:fixed-height-mode is enabled.

Since 3.0
*/
func (self *_TraitCellAreaContext) Allocate(width int, height int) {
	C.gtk_cell_area_context_allocate(self.CPointer, C.gint(width), C.gint(height))
	return
}

/*
Fetches the current allocation size for @context.

If the context was not allocated in width or height, or if the
context was recently reset with gtk_cell_area_context_reset(),
the returned value will be -1.
*/
func (self *_TraitCellAreaContext) GetAllocation() (width int, height int) {
	var __cgo__width C.gint
	var __cgo__height C.gint
	C.gtk_cell_area_context_get_allocation(self.CPointer, &__cgo__width, &__cgo__height)
	width = int(__cgo__width)
	height = int(__cgo__height)
	return
}

/*
Fetches the #GtkCellArea this @context was created by.

This is generally unneeded by layouting widgets; however,
it is important for the context implementation itself to
fetch information about the area it is being used for.

For instance at #GtkCellAreaContextClass.allocate() time
it’s important to know details about any cell spacing
that the #GtkCellArea is configured with in order to
compute a proper allocation.
*/
func (self *_TraitCellAreaContext) GetArea() (return__ *CellArea) {
	var __cgo__return__ *C.GtkCellArea
	__cgo__return__ = C.gtk_cell_area_context_get_area(self.CPointer)
	return__ = NewCellAreaFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the accumulative preferred height for all rows which have been
requested with this context.

After gtk_cell_area_context_reset() is called and/or before ever
requesting the size of a #GtkCellArea, the returned values are 0.
*/
func (self *_TraitCellAreaContext) GetPreferredHeight() (minimum_height int, natural_height int) {
	var __cgo__minimum_height C.gint
	var __cgo__natural_height C.gint
	C.gtk_cell_area_context_get_preferred_height(self.CPointer, &__cgo__minimum_height, &__cgo__natural_height)
	minimum_height = int(__cgo__minimum_height)
	natural_height = int(__cgo__natural_height)
	return
}

/*
Gets the accumulative preferred height for @width for all rows
which have been requested for the same said @width with this context.

After gtk_cell_area_context_reset() is called and/or before ever
requesting the size of a #GtkCellArea, the returned values are -1.
*/
func (self *_TraitCellAreaContext) GetPreferredHeightForWidth(width int) (minimum_height int, natural_height int) {
	var __cgo__minimum_height C.gint
	var __cgo__natural_height C.gint
	C.gtk_cell_area_context_get_preferred_height_for_width(self.CPointer, C.gint(width), &__cgo__minimum_height, &__cgo__natural_height)
	minimum_height = int(__cgo__minimum_height)
	natural_height = int(__cgo__natural_height)
	return
}

/*
Gets the accumulative preferred width for all rows which have been
requested with this context.

After gtk_cell_area_context_reset() is called and/or before ever
requesting the size of a #GtkCellArea, the returned values are 0.
*/
func (self *_TraitCellAreaContext) GetPreferredWidth() (minimum_width int, natural_width int) {
	var __cgo__minimum_width C.gint
	var __cgo__natural_width C.gint
	C.gtk_cell_area_context_get_preferred_width(self.CPointer, &__cgo__minimum_width, &__cgo__natural_width)
	minimum_width = int(__cgo__minimum_width)
	natural_width = int(__cgo__natural_width)
	return
}

/*
Gets the accumulative preferred width for @height for all rows which
have been requested for the same said @height with this context.

After gtk_cell_area_context_reset() is called and/or before ever
requesting the size of a #GtkCellArea, the returned values are -1.
*/
func (self *_TraitCellAreaContext) GetPreferredWidthForHeight(height int) (minimum_width int, natural_width int) {
	var __cgo__minimum_width C.gint
	var __cgo__natural_width C.gint
	C.gtk_cell_area_context_get_preferred_width_for_height(self.CPointer, C.gint(height), &__cgo__minimum_width, &__cgo__natural_width)
	minimum_width = int(__cgo__minimum_width)
	natural_width = int(__cgo__natural_width)
	return
}

/*
Causes the minimum and/or natural height to grow if the new
proposed sizes exceed the current minimum and natural height.

This is used by #GtkCellAreaContext implementations during
the request process over a series of #GtkTreeModel rows to
progressively push the requested height over a series of
gtk_cell_area_get_preferred_height() requests.
*/
func (self *_TraitCellAreaContext) PushPreferredHeight(minimum_height int, natural_height int) {
	C.gtk_cell_area_context_push_preferred_height(self.CPointer, C.gint(minimum_height), C.gint(natural_height))
	return
}

/*
Causes the minimum and/or natural width to grow if the new
proposed sizes exceed the current minimum and natural width.

This is used by #GtkCellAreaContext implementations during
the request process over a series of #GtkTreeModel rows to
progressively push the requested width over a series of
gtk_cell_area_get_preferred_width() requests.
*/
func (self *_TraitCellAreaContext) PushPreferredWidth(minimum_width int, natural_width int) {
	C.gtk_cell_area_context_push_preferred_width(self.CPointer, C.gint(minimum_width), C.gint(natural_width))
	return
}

/*
Resets any previously cached request and allocation
data.

When underlying #GtkTreeModel data changes its
important to reset the context if the content
size is allowed to shrink. If the content size
is only allowed to grow (this is usually an option
for views rendering large data stores as a measure
of optimization), then only the row that changed
or was inserted needs to be (re)requested with
gtk_cell_area_get_preferred_width().

When the new overall size of the context requires
that the allocated size changes (or whenever this
allocation changes at all), the variable row
sizes need to be re-requested for every row.

For instance, if the rows are displayed all with
the same width from top to bottom then a change
in the allocated width necessitates a recalculation
of all the displayed row heights using
gtk_cell_area_get_preferred_height_for_width().

Since 3.0
*/
func (self *_TraitCellAreaContext) Reset() {
	C.gtk_cell_area_context_reset(self.CPointer)
	return
}

type _TraitCellRenderer struct{ CPointer *C.GtkCellRenderer }

/*
Passes an activate event to the cell renderer for possible processing.
Some cell renderers may use events; for example, #GtkCellRendererToggle
toggles when it gets a mouse click.
*/
func (self *_TraitCellRenderer) Activate(event *C.GdkEvent, widget *Widget, path string, background_area *C.GdkRectangle, cell_area *C.GdkRectangle, flags C.GtkCellRendererState) (return__ bool) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_renderer_activate(self.CPointer, event, (*C.GtkWidget)(widget.CPointer), __cgo__path, background_area, cell_area, flags)
	C.free(unsafe.Pointer(__cgo__path))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the aligned area used by @cell inside @cell_area. Used for finding
the appropriate edit and focus rectangle.
*/
func (self *_TraitCellRenderer) GetAlignedArea(widget *Widget, flags C.GtkCellRendererState, cell_area *C.GdkRectangle) (aligned_area C.GdkRectangle) {
	C.gtk_cell_renderer_get_aligned_area(self.CPointer, (*C.GtkWidget)(widget.CPointer), flags, cell_area, &aligned_area)
	return
}

/*
Fills in @xalign and @yalign with the appropriate values of @cell.
*/
func (self *_TraitCellRenderer) GetAlignment() (xalign float32, yalign float32) {
	var __cgo__xalign C.gfloat
	var __cgo__yalign C.gfloat
	C.gtk_cell_renderer_get_alignment(self.CPointer, &__cgo__xalign, &__cgo__yalign)
	xalign = float32(__cgo__xalign)
	yalign = float32(__cgo__yalign)
	return
}

/*
Fills in @width and @height with the appropriate size of @cell.
*/
func (self *_TraitCellRenderer) GetFixedSize() (width int, height int) {
	var __cgo__width C.gint
	var __cgo__height C.gint
	C.gtk_cell_renderer_get_fixed_size(self.CPointer, &__cgo__width, &__cgo__height)
	width = int(__cgo__width)
	height = int(__cgo__height)
	return
}

/*
Fills in @xpad and @ypad with the appropriate values of @cell.
*/
func (self *_TraitCellRenderer) GetPadding() (xpad int, ypad int) {
	var __cgo__xpad C.gint
	var __cgo__ypad C.gint
	C.gtk_cell_renderer_get_padding(self.CPointer, &__cgo__xpad, &__cgo__ypad)
	xpad = int(__cgo__xpad)
	ypad = int(__cgo__ypad)
	return
}

/*
Retreives a renderer’s natural size when rendered to @widget.
*/
func (self *_TraitCellRenderer) GetPreferredHeight(widget *Widget) (minimum_size int, natural_size int) {
	var __cgo__minimum_size C.gint
	var __cgo__natural_size C.gint
	C.gtk_cell_renderer_get_preferred_height(self.CPointer, (*C.GtkWidget)(widget.CPointer), &__cgo__minimum_size, &__cgo__natural_size)
	minimum_size = int(__cgo__minimum_size)
	natural_size = int(__cgo__natural_size)
	return
}

/*
Retreives a cell renderers’s minimum and natural height if it were rendered to
@widget with the specified @width.
*/
func (self *_TraitCellRenderer) GetPreferredHeightForWidth(widget *Widget, width int) (minimum_height int, natural_height int) {
	var __cgo__minimum_height C.gint
	var __cgo__natural_height C.gint
	C.gtk_cell_renderer_get_preferred_height_for_width(self.CPointer, (*C.GtkWidget)(widget.CPointer), C.gint(width), &__cgo__minimum_height, &__cgo__natural_height)
	minimum_height = int(__cgo__minimum_height)
	natural_height = int(__cgo__natural_height)
	return
}

/*
Retrieves the minimum and natural size of a cell taking
into account the widget’s preference for height-for-width management.
*/
func (self *_TraitCellRenderer) GetPreferredSize(widget *Widget) (minimum_size C.GtkRequisition, natural_size C.GtkRequisition) {
	C.gtk_cell_renderer_get_preferred_size(self.CPointer, (*C.GtkWidget)(widget.CPointer), &minimum_size, &natural_size)
	return
}

/*
Retreives a renderer’s natural size when rendered to @widget.
*/
func (self *_TraitCellRenderer) GetPreferredWidth(widget *Widget) (minimum_size int, natural_size int) {
	var __cgo__minimum_size C.gint
	var __cgo__natural_size C.gint
	C.gtk_cell_renderer_get_preferred_width(self.CPointer, (*C.GtkWidget)(widget.CPointer), &__cgo__minimum_size, &__cgo__natural_size)
	minimum_size = int(__cgo__minimum_size)
	natural_size = int(__cgo__natural_size)
	return
}

/*
Retreives a cell renderers’s minimum and natural width if it were rendered to
@widget with the specified @height.
*/
func (self *_TraitCellRenderer) GetPreferredWidthForHeight(widget *Widget, height int) (minimum_width int, natural_width int) {
	var __cgo__minimum_width C.gint
	var __cgo__natural_width C.gint
	C.gtk_cell_renderer_get_preferred_width_for_height(self.CPointer, (*C.GtkWidget)(widget.CPointer), C.gint(height), &__cgo__minimum_width, &__cgo__natural_width)
	minimum_width = int(__cgo__minimum_width)
	natural_width = int(__cgo__natural_width)
	return
}

/*
Gets whether the cell renderer prefers a height-for-width layout
or a width-for-height layout.
*/
func (self *_TraitCellRenderer) GetRequestMode() (return__ C.GtkSizeRequestMode) {
	return__ = C.gtk_cell_renderer_get_request_mode(self.CPointer)
	return
}

/*
Returns the cell renderer’s sensitivity.
*/
func (self *_TraitCellRenderer) GetSensitive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_renderer_get_sensitive(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_cell_renderer_get_size is not generated due to deprecation attr

/*
Translates the cell renderer state to #GtkStateFlags,
based on the cell renderer and widget sensitivity, and
the given #GtkCellRendererState.
*/
func (self *_TraitCellRenderer) GetState(widget *Widget, cell_state C.GtkCellRendererState) (return__ C.GtkStateFlags) {
	return__ = C.gtk_cell_renderer_get_state(self.CPointer, (*C.GtkWidget)(widget.CPointer), cell_state)
	return
}

/*
Returns the cell renderer’s visibility.
*/
func (self *_TraitCellRenderer) GetVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_renderer_get_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks whether the cell renderer can do something when activated.
*/
func (self *_TraitCellRenderer) IsActivatable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_renderer_is_activatable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Invokes the virtual render function of the #GtkCellRenderer. The three
passed-in rectangles are areas in @cr. Most renderers will draw within
@cell_area; the xalign, yalign, xpad, and ypad fields of the #GtkCellRenderer
should be honored with respect to @cell_area. @background_area includes the
blank space around the cell, and also the area containing the tree expander;
so the @background_area rectangles for all cells tile to cover the entire
@window.
*/
func (self *_TraitCellRenderer) Render(cr *C.cairo_t, widget *Widget, background_area *C.GdkRectangle, cell_area *C.GdkRectangle, flags C.GtkCellRendererState) {
	C.gtk_cell_renderer_render(self.CPointer, cr, (*C.GtkWidget)(widget.CPointer), background_area, cell_area, flags)
	return
}

/*
Sets the renderer’s alignment within its available space.
*/
func (self *_TraitCellRenderer) SetAlignment(xalign float32, yalign float32) {
	C.gtk_cell_renderer_set_alignment(self.CPointer, C.gfloat(xalign), C.gfloat(yalign))
	return
}

/*
Sets the renderer size to be explicit, independent of the properties set.
*/
func (self *_TraitCellRenderer) SetFixedSize(width int, height int) {
	C.gtk_cell_renderer_set_fixed_size(self.CPointer, C.gint(width), C.gint(height))
	return
}

/*
Sets the renderer’s padding.
*/
func (self *_TraitCellRenderer) SetPadding(xpad int, ypad int) {
	C.gtk_cell_renderer_set_padding(self.CPointer, C.gint(xpad), C.gint(ypad))
	return
}

/*
Sets the cell renderer’s sensitivity.
*/
func (self *_TraitCellRenderer) SetSensitive(sensitive bool) {
	__cgo__sensitive := C.gboolean(0)
	if sensitive {
		__cgo__sensitive = C.gboolean(1)
	}
	C.gtk_cell_renderer_set_sensitive(self.CPointer, __cgo__sensitive)
	return
}

/*
Sets the cell renderer’s visibility.
*/
func (self *_TraitCellRenderer) SetVisible(visible bool) {
	__cgo__visible := C.gboolean(0)
	if visible {
		__cgo__visible = C.gboolean(1)
	}
	C.gtk_cell_renderer_set_visible(self.CPointer, __cgo__visible)
	return
}

/*
Passes an activate event to the cell renderer for possible processing.
*/
func (self *_TraitCellRenderer) StartEditing(event *C.GdkEvent, widget *Widget, path string, background_area *C.GdkRectangle, cell_area *C.GdkRectangle, flags C.GtkCellRendererState) (return__ *C.GtkCellEditable) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	return__ = C.gtk_cell_renderer_start_editing(self.CPointer, event, (*C.GtkWidget)(widget.CPointer), __cgo__path, background_area, cell_area, flags)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Informs the cell renderer that the editing is stopped.
If @canceled is %TRUE, the cell renderer will emit the
#GtkCellRenderer::editing-canceled signal.

This function should be called by cell renderer implementations
in response to the #GtkCellEditable::editing-done signal of
#GtkCellEditable.
*/
func (self *_TraitCellRenderer) StopEditing(canceled bool) {
	__cgo__canceled := C.gboolean(0)
	if canceled {
		__cgo__canceled = C.gboolean(1)
	}
	C.gtk_cell_renderer_stop_editing(self.CPointer, __cgo__canceled)
	return
}

type _TraitCellRendererAccel struct{ CPointer *C.GtkCellRendererAccel }

type _TraitCellRendererCombo struct{ CPointer *C.GtkCellRendererCombo }

type _TraitCellRendererPixbuf struct{ CPointer *C.GtkCellRendererPixbuf }

type _TraitCellRendererProgress struct{ CPointer *C.GtkCellRendererProgress }

type _TraitCellRendererSpin struct{ CPointer *C.GtkCellRendererSpin }

type _TraitCellRendererSpinner struct{ CPointer *C.GtkCellRendererSpinner }

type _TraitCellRendererText struct{ CPointer *C.GtkCellRendererText }

/*
Sets the height of a renderer to explicitly be determined by the “font” and
“y_pad” property set on it.  Further changes in these properties do not
affect the height, so they must be accompanied by a subsequent call to this
function.  Using this function is unflexible, and should really only be used
if calculating the size of a cell is too slow (ie, a massive number of cells
displayed).  If @number_of_rows is -1, then the fixed height is unset, and
the height is determined by the properties again.
*/
func (self *_TraitCellRendererText) SetFixedHeightFromFont(number_of_rows int) {
	C.gtk_cell_renderer_text_set_fixed_height_from_font(self.CPointer, C.gint(number_of_rows))
	return
}

type _TraitCellRendererToggle struct{ CPointer *C.GtkCellRendererToggle }

/*
Returns whether the cell renderer is activatable. See
gtk_cell_renderer_toggle_set_activatable().
*/
func (self *_TraitCellRendererToggle) GetActivatable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_renderer_toggle_get_activatable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the cell renderer is active. See
gtk_cell_renderer_toggle_set_active().
*/
func (self *_TraitCellRendererToggle) GetActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_renderer_toggle_get_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether we’re rendering radio toggles rather than checkboxes.
*/
func (self *_TraitCellRendererToggle) GetRadio() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_renderer_toggle_get_radio(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Makes the cell renderer activatable.
*/
func (self *_TraitCellRendererToggle) SetActivatable(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_cell_renderer_toggle_set_activatable(self.CPointer, __cgo__setting)
	return
}

/*
Activates or deactivates a cell renderer.
*/
func (self *_TraitCellRendererToggle) SetActive(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_cell_renderer_toggle_set_active(self.CPointer, __cgo__setting)
	return
}

/*
If @radio is %TRUE, the cell renderer renders a radio toggle
(i.e. a toggle in a group of mutually-exclusive toggles).
If %FALSE, it renders a check toggle (a standalone boolean option).
This can be set globally for the cell renderer, or changed just
before rendering each cell in the model (for #GtkTreeView, you set
up a per-row setting using #GtkTreeViewColumn to associate model
columns with cell renderer properties).
*/
func (self *_TraitCellRendererToggle) SetRadio(radio bool) {
	__cgo__radio := C.gboolean(0)
	if radio {
		__cgo__radio = C.gboolean(1)
	}
	C.gtk_cell_renderer_toggle_set_radio(self.CPointer, __cgo__radio)
	return
}

type _TraitCellView struct{ CPointer *C.GtkCellView }

/*
Returns a #GtkTreePath referring to the currently
displayed row. If no row is currently displayed,
%NULL is returned.
*/
func (self *_TraitCellView) GetDisplayedRow() (return__ *TreePath) {
	var __cgo__return__ *C.GtkTreePath
	__cgo__return__ = C.gtk_cell_view_get_displayed_row(self.CPointer)
	return__ = (*TreePath)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Gets whether @cell_view is configured to draw all of its
cells in a sensitive state.
*/
func (self *_TraitCellView) GetDrawSensitive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_view_get_draw_sensitive(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether @cell_view is configured to request space
to fit the entire #GtkTreeModel.
*/
func (self *_TraitCellView) GetFitModel() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_cell_view_get_fit_model(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the model for @cell_view. If no model is used %NULL is
returned.
*/
func (self *_TraitCellView) GetModel() (return__ *C.GtkTreeModel) {
	return__ = C.gtk_cell_view_get_model(self.CPointer)
	return
}

// gtk_cell_view_get_size_of_row is not generated due to deprecation attr

// gtk_cell_view_set_background_color is not generated due to deprecation attr

/*
Sets the background color of @cell_view.
*/
func (self *_TraitCellView) SetBackgroundRgba(rgba *C.GdkRGBA) {
	C.gtk_cell_view_set_background_rgba(self.CPointer, rgba)
	return
}

/*
Sets the row of the model that is currently displayed
by the #GtkCellView. If the path is unset, then the
contents of the cellview “stick” at their last value;
this is not normally a desired result, but may be
a needed intermediate state if say, the model for
the #GtkCellView becomes temporarily empty.
*/
func (self *_TraitCellView) SetDisplayedRow(path *TreePath) {
	C.gtk_cell_view_set_displayed_row(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Sets whether @cell_view should draw all of its
cells in a sensitive state, this is used by #GtkComboBox menus
to ensure that rows with insensitive cells that contain
children appear sensitive in the parent menu item.
*/
func (self *_TraitCellView) SetDrawSensitive(draw_sensitive bool) {
	__cgo__draw_sensitive := C.gboolean(0)
	if draw_sensitive {
		__cgo__draw_sensitive = C.gboolean(1)
	}
	C.gtk_cell_view_set_draw_sensitive(self.CPointer, __cgo__draw_sensitive)
	return
}

/*
Sets whether @cell_view should request space to fit the entire #GtkTreeModel.

This is used by #GtkComboBox to ensure that the cell view displayed on
the combo box’s button always gets enough space and does not resize
when selection changes.
*/
func (self *_TraitCellView) SetFitModel(fit_model bool) {
	__cgo__fit_model := C.gboolean(0)
	if fit_model {
		__cgo__fit_model = C.gboolean(1)
	}
	C.gtk_cell_view_set_fit_model(self.CPointer, __cgo__fit_model)
	return
}

/*
Sets the model for @cell_view.  If @cell_view already has a model
set, it will remove it before setting the new model.  If @model is
%NULL, then it will unset the old model.
*/
func (self *_TraitCellView) SetModel(model *C.GtkTreeModel) {
	C.gtk_cell_view_set_model(self.CPointer, model)
	return
}

type _TraitCheckButton struct{ CPointer *C.GtkCheckButton }

type _TraitCheckMenuItem struct{ CPointer *C.GtkCheckMenuItem }

/*
Returns whether the check menu item is active. See
gtk_check_menu_item_set_active ().
*/
func (self *_TraitCheckMenuItem) GetActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_check_menu_item_get_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether @check_menu_item looks like a #GtkRadioMenuItem
*/
func (self *_TraitCheckMenuItem) GetDrawAsRadio() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_check_menu_item_get_draw_as_radio(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the value set by gtk_check_menu_item_set_inconsistent().
*/
func (self *_TraitCheckMenuItem) GetInconsistent() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_check_menu_item_get_inconsistent(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the active state of the menu item’s check box.
*/
func (self *_TraitCheckMenuItem) SetActive(is_active bool) {
	__cgo__is_active := C.gboolean(0)
	if is_active {
		__cgo__is_active = C.gboolean(1)
	}
	C.gtk_check_menu_item_set_active(self.CPointer, __cgo__is_active)
	return
}

/*
Sets whether @check_menu_item is drawn like a #GtkRadioMenuItem
*/
func (self *_TraitCheckMenuItem) SetDrawAsRadio(draw_as_radio bool) {
	__cgo__draw_as_radio := C.gboolean(0)
	if draw_as_radio {
		__cgo__draw_as_radio = C.gboolean(1)
	}
	C.gtk_check_menu_item_set_draw_as_radio(self.CPointer, __cgo__draw_as_radio)
	return
}

/*
If the user has selected a range of elements (such as some text or
spreadsheet cells) that are affected by a boolean setting, and the
current values in that range are inconsistent, you may want to
display the check in an “in between” state. This function turns on
“in between” display.  Normally you would turn off the inconsistent
state again if the user explicitly selects a setting. This has to be
done manually, gtk_check_menu_item_set_inconsistent() only affects
visual appearance, it doesn’t affect the semantics of the widget.
*/
func (self *_TraitCheckMenuItem) SetInconsistent(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_check_menu_item_set_inconsistent(self.CPointer, __cgo__setting)
	return
}

/*
Emits the #GtkCheckMenuItem::toggled signal.
*/
func (self *_TraitCheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(self.CPointer)
	return
}

type _TraitClipboard struct{ CPointer *C.GtkClipboard }

/*
Clears the contents of the clipboard. Generally this should only
be called between the time you call gtk_clipboard_set_with_owner()
or gtk_clipboard_set_with_data(),
and when the @clear_func you supplied is called. Otherwise, the
clipboard may be owned by someone else.
*/
func (self *_TraitClipboard) Clear() {
	C.gtk_clipboard_clear(self.CPointer)
	return
}

/*
Gets the #GdkDisplay associated with @clipboard
*/
func (self *_TraitClipboard) GetDisplay() (return__ *C.GdkDisplay) {
	return__ = C.gtk_clipboard_get_display(self.CPointer)
	return
}

/*
If the clipboard contents callbacks were set with
gtk_clipboard_set_with_owner(), and the gtk_clipboard_set_with_data() or
gtk_clipboard_clear() has not subsequently called, returns the owner set
by gtk_clipboard_set_with_owner().
*/
func (self *_TraitClipboard) GetOwner() (return__ *C.GObject) {
	return__ = C.gtk_clipboard_get_owner(self.CPointer)
	return
}

/*
Requests the contents of clipboard as the given target.
When the results of the result are later received the supplied callback
will be called.
*/
func (self *_TraitClipboard) RequestContents(target C.GdkAtom, callback C.GtkClipboardReceivedFunc, user_data unsafe.Pointer) {
	C.gtk_clipboard_request_contents(self.CPointer, target, callback, (C.gpointer)(user_data))
	return
}

/*
Requests the contents of the clipboard as image. When the image is
later received, it will be converted to a #GdkPixbuf, and
@callback will be called.

The @pixbuf parameter to @callback will contain the resulting
#GdkPixbuf if the request succeeded, or %NULL if it failed. This
could happen for various reasons, in particular if the clipboard
was empty or if the contents of the clipboard could not be
converted into an image.
*/
func (self *_TraitClipboard) RequestImage(callback C.GtkClipboardImageReceivedFunc, user_data unsafe.Pointer) {
	C.gtk_clipboard_request_image(self.CPointer, callback, (C.gpointer)(user_data))
	return
}

/*
Requests the contents of the clipboard as rich text. When the rich
text is later received, @callback will be called.

The @text parameter to @callback will contain the resulting rich
text if the request succeeded, or %NULL if it failed. The @length
parameter will contain @text’s length. This function can fail for
various reasons, in particular if the clipboard was empty or if the
contents of the clipboard could not be converted into rich text form.
*/
func (self *_TraitClipboard) RequestRichText(buffer *TextBuffer, callback C.GtkClipboardRichTextReceivedFunc, user_data unsafe.Pointer) {
	C.gtk_clipboard_request_rich_text(self.CPointer, (*C.GtkTextBuffer)(buffer.CPointer), callback, (C.gpointer)(user_data))
	return
}

/*
Requests the contents of the clipboard as list of supported targets.
When the list is later received, @callback will be called.

The @targets parameter to @callback will contain the resulting targets if
the request succeeded, or %NULL if it failed.
*/
func (self *_TraitClipboard) RequestTargets(callback C.GtkClipboardTargetsReceivedFunc, user_data unsafe.Pointer) {
	C.gtk_clipboard_request_targets(self.CPointer, callback, (C.gpointer)(user_data))
	return
}

/*
Requests the contents of the clipboard as text. When the text is
later received, it will be converted to UTF-8 if necessary, and
@callback will be called.

The @text parameter to @callback will contain the resulting text if
the request succeeded, or %NULL if it failed. This could happen for
various reasons, in particular if the clipboard was empty or if the
contents of the clipboard could not be converted into text form.
*/
func (self *_TraitClipboard) RequestText(callback C.GtkClipboardTextReceivedFunc, user_data unsafe.Pointer) {
	C.gtk_clipboard_request_text(self.CPointer, callback, (C.gpointer)(user_data))
	return
}

/*
Requests the contents of the clipboard as URIs. When the URIs are
later received @callback will be called.

The @uris parameter to @callback will contain the resulting array of
URIs if the request succeeded, or %NULL if it failed. This could happen
for various reasons, in particular if the clipboard was empty or if the
contents of the clipboard could not be converted into URI form.
*/
func (self *_TraitClipboard) RequestUris(callback C.GtkClipboardURIReceivedFunc, user_data unsafe.Pointer) {
	C.gtk_clipboard_request_uris(self.CPointer, callback, (C.gpointer)(user_data))
	return
}

/*
Hints that the clipboard data should be stored somewhere when the
application exits or when gtk_clipboard_store () is called.

This value is reset when the clipboard owner changes.
Where the clipboard data is stored is platform dependent,
see gdk_display_store_clipboard () for more information.
*/
func (self *_TraitClipboard) SetCanStore(targets *TargetEntry, n_targets int) {
	C.gtk_clipboard_set_can_store(self.CPointer, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.gint(n_targets))
	return
}

/*
Sets the contents of the clipboard to the given #GdkPixbuf.
GTK+ will take responsibility for responding for requests
for the image, and for converting the image into the
requested format.
*/
func (self *_TraitClipboard) SetImage(pixbuf *C.GdkPixbuf) {
	C.gtk_clipboard_set_image(self.CPointer, pixbuf)
	return
}

/*
Sets the contents of the clipboard to the given UTF-8 string. GTK+ will
make a copy of the text and take responsibility for responding
for requests for the text, and for converting the text into
the requested format.
*/
func (self *_TraitClipboard) SetText(text string, len_ int) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_clipboard_set_text(self.CPointer, __cgo__text, C.gint(len_))
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Virtually sets the contents of the specified clipboard by providing
a list of supported formats for the clipboard data and a function
to call to get the actual data when it is requested.
*/
func (self *_TraitClipboard) SetWithData(targets *TargetEntry, n_targets uint, get_func C.GtkClipboardGetFunc, clear_func C.GtkClipboardClearFunc, user_data unsafe.Pointer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_clipboard_set_with_data(self.CPointer, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.guint(n_targets), get_func, clear_func, (C.gpointer)(user_data))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Virtually sets the contents of the specified clipboard by providing
a list of supported formats for the clipboard data and a function
to call to get the actual data when it is requested.

The difference between this function and gtk_clipboard_set_with_data()
is that instead of an generic @user_data pointer, a #GObject is passed
in.
*/
func (self *_TraitClipboard) SetWithOwner(targets *TargetEntry, n_targets uint, get_func C.GtkClipboardGetFunc, clear_func C.GtkClipboardClearFunc, owner *C.GObject) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_clipboard_set_with_owner(self.CPointer, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.guint(n_targets), get_func, clear_func, owner)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Stores the current clipboard data somewhere so that it will stay
around after the application has quit.
*/
func (self *_TraitClipboard) Store() {
	C.gtk_clipboard_store(self.CPointer)
	return
}

/*
Requests the contents of the clipboard using the given target.
This function waits for the data to be received using the main
loop, so events, timeouts, etc, may be dispatched during the wait.
*/
func (self *_TraitClipboard) WaitForContents(target C.GdkAtom) (return__ *SelectionData) {
	var __cgo__return__ *C.GtkSelectionData
	__cgo__return__ = C.gtk_clipboard_wait_for_contents(self.CPointer, target)
	return__ = (*SelectionData)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Requests the contents of the clipboard as image and converts
the result to a #GdkPixbuf. This function waits for
the data to be received using the main loop, so events,
timeouts, etc, may be dispatched during the wait.
*/
func (self *_TraitClipboard) WaitForImage() (return__ *C.GdkPixbuf) {
	return__ = C.gtk_clipboard_wait_for_image(self.CPointer)
	return
}

/*
Requests the contents of the clipboard as rich text.  This function
waits for the data to be received using the main loop, so events,
timeouts, etc, may be dispatched during the wait.
*/
func (self *_TraitClipboard) WaitForRichText(buffer *TextBuffer) (format C.GdkAtom, length int64, return__ []byte) {
	var __cgo__length C.gsize
	var __cgo__return__ *C.guint8
	__cgo__return__ = C.gtk_clipboard_wait_for_rich_text(self.CPointer, (*C.GtkTextBuffer)(buffer.CPointer), &format, &__cgo__length)
	length = int64(__cgo__length)
	defer func() { return__ = C.GoBytes(unsafe.Pointer(__cgo__return__), C.int(length)) }()
	return
}

/*
Returns a list of targets that are present on the clipboard, or %NULL
if there aren’t any targets available. The returned list must be
freed with g_free().
This function waits for the data to be received using the main
loop, so events, timeouts, etc, may be dispatched during the wait.
*/
func (self *_TraitClipboard) WaitForTargets() (targets *C.GdkAtom, n_targets int, return__ bool) {
	var __cgo__n_targets C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_clipboard_wait_for_targets(self.CPointer, &targets, &__cgo__n_targets)
	n_targets = int(__cgo__n_targets)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Requests the contents of the clipboard as text and converts
the result to UTF-8 if necessary. This function waits for
the data to be received using the main loop, so events,
timeouts, etc, may be dispatched during the wait.
*/
func (self *_TraitClipboard) WaitForText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_clipboard_wait_for_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Requests the contents of the clipboard as URIs. This function waits
for the data to be received using the main loop, so events,
timeouts, etc, may be dispatched during the wait.
*/
func (self *_TraitClipboard) WaitForUris() (return__ []string) {
	var __cgo__return__ **C.gchar
	__cgo__return__ = C.gtk_clipboard_wait_for_uris(self.CPointer)
	var __slice__return__ []*C.gchar
	__header__return__ := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__return__))
	__header__return__.Len = 4294967296
	__header__return__.Cap = 4294967296
	__header__return__.Data = uintptr(unsafe.Pointer(__cgo__return__))
	for _, p := range __slice__return__ {
		if p == nil {
			break
		}
		return__ = append(return__, C.GoString((*C.char)(unsafe.Pointer(p))))
	}
	return
}

/*
Test to see if there is an image available to be pasted
This is done by requesting the TARGETS atom and checking
if it contains any of the supported image targets. This function
waits for the data to be received using the main loop, so events,
timeouts, etc, may be dispatched during the wait.

This function is a little faster than calling
gtk_clipboard_wait_for_image() since it doesn’t need to retrieve
the actual image data.
*/
func (self *_TraitClipboard) WaitIsImageAvailable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_clipboard_wait_is_image_available(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Test to see if there is rich text available to be pasted
This is done by requesting the TARGETS atom and checking
if it contains any of the supported rich text targets. This function
waits for the data to be received using the main loop, so events,
timeouts, etc, may be dispatched during the wait.

This function is a little faster than calling
gtk_clipboard_wait_for_rich_text() since it doesn’t need to retrieve
the actual text.
*/
func (self *_TraitClipboard) WaitIsRichTextAvailable(buffer *TextBuffer) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_clipboard_wait_is_rich_text_available(self.CPointer, (*C.GtkTextBuffer)(buffer.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Checks if a clipboard supports pasting data of a given type. This
function can be used to determine if a “Paste” menu item should be
insensitive or not.

If you want to see if there’s text available on the clipboard, use
gtk_clipboard_wait_is_text_available () instead.
*/
func (self *_TraitClipboard) WaitIsTargetAvailable(target C.GdkAtom) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_clipboard_wait_is_target_available(self.CPointer, target)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Test to see if there is text available to be pasted
This is done by requesting the TARGETS atom and checking
if it contains any of the supported text targets. This function
waits for the data to be received using the main loop, so events,
timeouts, etc, may be dispatched during the wait.

This function is a little faster than calling
gtk_clipboard_wait_for_text() since it doesn’t need to retrieve
the actual text.
*/
func (self *_TraitClipboard) WaitIsTextAvailable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_clipboard_wait_is_text_available(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Test to see if there is a list of URIs available to be pasted
This is done by requesting the TARGETS atom and checking
if it contains the URI targets. This function
waits for the data to be received using the main loop, so events,
timeouts, etc, may be dispatched during the wait.

This function is a little faster than calling
gtk_clipboard_wait_for_uris() since it doesn’t need to retrieve
the actual URI data.
*/
func (self *_TraitClipboard) WaitIsUrisAvailable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_clipboard_wait_is_uris_available(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type _TraitColorButton struct{ CPointer *C.GtkColorButton }

// gtk_color_button_get_alpha is not generated due to deprecation attr

// gtk_color_button_get_color is not generated due to deprecation attr

// gtk_color_button_get_rgba is not generated due to deprecation attr

/*
Gets the title of the color selection dialog.
*/
func (self *_TraitColorButton) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_color_button_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// gtk_color_button_get_use_alpha is not generated due to deprecation attr

// gtk_color_button_set_alpha is not generated due to deprecation attr

// gtk_color_button_set_color is not generated due to deprecation attr

// gtk_color_button_set_rgba is not generated due to deprecation attr

/*
Sets the title for the color selection dialog.
*/
func (self *_TraitColorButton) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_color_button_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

// gtk_color_button_set_use_alpha is not generated due to deprecation attr

type _TraitColorChooserDialog struct{ CPointer *C.GtkColorChooserDialog }

type _TraitColorChooserWidget struct{ CPointer *C.GtkColorChooserWidget }

type _TraitColorSelection struct{ CPointer *C.GtkColorSelection }

// gtk_color_selection_get_current_alpha is not generated due to explicit ignore

// gtk_color_selection_get_current_color is not generated due to deprecation attr

// gtk_color_selection_get_current_rgba is not generated due to explicit ignore

// gtk_color_selection_get_has_opacity_control is not generated due to explicit ignore

// gtk_color_selection_get_has_palette is not generated due to explicit ignore

// gtk_color_selection_get_previous_alpha is not generated due to explicit ignore

// gtk_color_selection_get_previous_color is not generated due to deprecation attr

// gtk_color_selection_get_previous_rgba is not generated due to explicit ignore

// gtk_color_selection_is_adjusting is not generated due to explicit ignore

// gtk_color_selection_set_current_alpha is not generated due to explicit ignore

// gtk_color_selection_set_current_color is not generated due to deprecation attr

// gtk_color_selection_set_current_rgba is not generated due to explicit ignore

// gtk_color_selection_set_has_opacity_control is not generated due to explicit ignore

// gtk_color_selection_set_has_palette is not generated due to explicit ignore

// gtk_color_selection_set_previous_alpha is not generated due to explicit ignore

// gtk_color_selection_set_previous_color is not generated due to deprecation attr

// gtk_color_selection_set_previous_rgba is not generated due to explicit ignore

type _TraitColorSelectionDialog struct{ CPointer *C.GtkColorSelectionDialog }

// gtk_color_selection_dialog_get_color_selection is not generated due to explicit ignore

type _TraitComboBox struct{ CPointer *C.GtkComboBox }

/*
Returns the index of the currently active item, or -1 if there’s no
active item. If the model is a non-flat treemodel, and the active item
is not an immediate child of the root of the tree, this function returns
`gtk_tree_path_get_indices (path)[0]`, where
`path` is the #GtkTreePath of the active item.
*/
func (self *_TraitComboBox) GetActive() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_combo_box_get_active(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the ID of the active row of @combo_box.  This value is taken
from the active row and the column specified by the #GtkComboBox:id-column
property of @combo_box (see gtk_combo_box_set_id_column()).

The returned value is an interned string which means that you can
compare the pointer by value to other interned strings and that you
must not free it.

If the #GtkComboBox:id-column property of @combo_box is not set, or if
no row is active, or if the active row has a %NULL ID value, then %NULL
is returned.
*/
func (self *_TraitComboBox) GetActiveId() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_combo_box_get_active_id(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Sets @iter to point to the current active item, if it exists.
*/
func (self *_TraitComboBox) GetActiveIter() (iter C.GtkTreeIter, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_combo_box_get_active_iter(self.CPointer, &iter)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_combo_box_get_add_tearoffs is not generated due to deprecation attr

/*
Returns whether the combo box sets the dropdown button
sensitive or not when there are no items in the model.
*/
func (self *_TraitComboBox) GetButtonSensitivity() (return__ C.GtkSensitivityType) {
	return__ = C.gtk_combo_box_get_button_sensitivity(self.CPointer)
	return
}

/*
Returns the column with column span information for @combo_box.
*/
func (self *_TraitComboBox) GetColumnSpanColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_combo_box_get_column_span_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the column which @combo_box is using to get the strings
from to display in the internal entry.
*/
func (self *_TraitComboBox) GetEntryTextColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_combo_box_get_entry_text_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns whether the combo box grabs focus when it is clicked
with the mouse. See gtk_combo_box_set_focus_on_click().
*/
func (self *_TraitComboBox) GetFocusOnClick() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_combo_box_get_focus_on_click(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the combo box has an entry.
*/
func (self *_TraitComboBox) GetHasEntry() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_combo_box_get_has_entry(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the column which @combo_box is using to get string IDs
for values from.
*/
func (self *_TraitComboBox) GetIdColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_combo_box_get_id_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the #GtkTreeModel which is acting as data source for @combo_box.
*/
func (self *_TraitComboBox) GetModel() (return__ *C.GtkTreeModel) {
	return__ = C.gtk_combo_box_get_model(self.CPointer)
	return
}

/*
Gets the accessible object corresponding to the combo box’s popup.

This function is mostly intended for use by accessibility technologies;
applications should have little use for it.
*/
func (self *_TraitComboBox) GetPopupAccessible() (return__ *C.AtkObject) {
	return__ = C.gtk_combo_box_get_popup_accessible(self.CPointer)
	return
}

/*
Gets whether the popup uses a fixed width matching
the allocated width of the combo box.
*/
func (self *_TraitComboBox) GetPopupFixedWidth() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_combo_box_get_popup_fixed_width(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current row separator function.
*/
func (self *_TraitComboBox) GetRowSeparatorFunc() (return__ C.GtkTreeViewRowSeparatorFunc) {
	return__ = C.gtk_combo_box_get_row_separator_func(self.CPointer)
	return
}

/*
Returns the column with row span information for @combo_box.
*/
func (self *_TraitComboBox) GetRowSpanColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_combo_box_get_row_span_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

// gtk_combo_box_get_title is not generated due to deprecation attr

/*
Returns the wrap width which is used to determine the number of columns
for the popup menu. If the wrap width is larger than 1, the combo box
is in table mode.
*/
func (self *_TraitComboBox) GetWrapWidth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_combo_box_get_wrap_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Hides the menu or dropdown list of @combo_box.

This function is mostly intended for use by accessibility technologies;
applications should have little use for it.
*/
func (self *_TraitComboBox) Popdown() {
	C.gtk_combo_box_popdown(self.CPointer)
	return
}

/*
Pops up the menu or dropdown list of @combo_box.

This function is mostly intended for use by accessibility technologies;
applications should have little use for it.
*/
func (self *_TraitComboBox) Popup() {
	C.gtk_combo_box_popup(self.CPointer)
	return
}

/*
Pops up the menu or dropdown list of @combo_box, the popup window
will be grabbed so only @device and its associated pointer/keyboard
are the only #GdkDevices able to send events to it.
*/
func (self *_TraitComboBox) PopupForDevice(device *C.GdkDevice) {
	C.gtk_combo_box_popup_for_device(self.CPointer, device)
	return
}

/*
Sets the active item of @combo_box to be the item at @index.
*/
func (self *_TraitComboBox) SetActive(index_ int) {
	C.gtk_combo_box_set_active(self.CPointer, C.gint(index_))
	return
}

/*
Changes the active row of @combo_box to the one that has an ID equal to
@active_id, or unsets the active row if @active_id is %NULL.  Rows having
a %NULL ID string cannot be made active by this function.

If the #GtkComboBox:id-column property of @combo_box is unset or if no
row has the given ID then the function does nothing and returns %FALSE.
*/
func (self *_TraitComboBox) SetActiveId(active_id string) (return__ bool) {
	__cgo__active_id := (*C.gchar)(unsafe.Pointer(C.CString(active_id)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_combo_box_set_active_id(self.CPointer, __cgo__active_id)
	C.free(unsafe.Pointer(__cgo__active_id))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the current active item to be the one referenced by @iter, or
unsets the active item if @iter is %NULL.
*/
func (self *_TraitComboBox) SetActiveIter(iter *TreeIter) {
	C.gtk_combo_box_set_active_iter(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return
}

// gtk_combo_box_set_add_tearoffs is not generated due to deprecation attr

/*
Sets whether the dropdown button of the combo box should be
always sensitive (%GTK_SENSITIVITY_ON), never sensitive (%GTK_SENSITIVITY_OFF)
or only if there is at least one item to display (%GTK_SENSITIVITY_AUTO).
*/
func (self *_TraitComboBox) SetButtonSensitivity(sensitivity C.GtkSensitivityType) {
	C.gtk_combo_box_set_button_sensitivity(self.CPointer, sensitivity)
	return
}

/*
Sets the column with column span information for @combo_box to be
@column_span. The column span column contains integers which indicate
how many columns an item should span.
*/
func (self *_TraitComboBox) SetColumnSpanColumn(column_span int) {
	C.gtk_combo_box_set_column_span_column(self.CPointer, C.gint(column_span))
	return
}

/*
Sets the model column which @combo_box should use to get strings from
to be @text_column. The column @text_column in the model of @combo_box
must be of type %G_TYPE_STRING.

This is only relevant if @combo_box has been created with
#GtkComboBox:has-entry as %TRUE.
*/
func (self *_TraitComboBox) SetEntryTextColumn(text_column int) {
	C.gtk_combo_box_set_entry_text_column(self.CPointer, C.gint(text_column))
	return
}

/*
Sets whether the combo box will grab focus when it is clicked with
the mouse. Making mouse clicks not grab focus is useful in places
like toolbars where you don’t want the keyboard focus removed from
the main area of the application.
*/
func (self *_TraitComboBox) SetFocusOnClick(focus_on_click bool) {
	__cgo__focus_on_click := C.gboolean(0)
	if focus_on_click {
		__cgo__focus_on_click = C.gboolean(1)
	}
	C.gtk_combo_box_set_focus_on_click(self.CPointer, __cgo__focus_on_click)
	return
}

/*
Sets the model column which @combo_box should use to get string IDs
for values from. The column @id_column in the model of @combo_box
must be of type %G_TYPE_STRING.
*/
func (self *_TraitComboBox) SetIdColumn(id_column int) {
	C.gtk_combo_box_set_id_column(self.CPointer, C.gint(id_column))
	return
}

/*
Sets the model used by @combo_box to be @model. Will unset a previously set
model (if applicable). If model is %NULL, then it will unset the model.

Note that this function does not clear the cell renderers, you have to
call gtk_cell_layout_clear() yourself if you need to set up different
cell renderers for the new model.
*/
func (self *_TraitComboBox) SetModel(model *C.GtkTreeModel) {
	C.gtk_combo_box_set_model(self.CPointer, model)
	return
}

/*
Specifies whether the popup’s width should be a fixed width
matching the allocated width of the combo box.
*/
func (self *_TraitComboBox) SetPopupFixedWidth(fixed bool) {
	__cgo__fixed := C.gboolean(0)
	if fixed {
		__cgo__fixed = C.gboolean(1)
	}
	C.gtk_combo_box_set_popup_fixed_width(self.CPointer, __cgo__fixed)
	return
}

/*
Sets the row separator function, which is used to determine
whether a row should be drawn as a separator. If the row separator
function is %NULL, no separators are drawn. This is the default value.
*/
func (self *_TraitComboBox) SetRowSeparatorFunc(func_ C.GtkTreeViewRowSeparatorFunc, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_combo_box_set_row_separator_func(self.CPointer, func_, (C.gpointer)(data), destroy)
	return
}

/*
Sets the column with row span information for @combo_box to be @row_span.
The row span column contains integers which indicate how many rows
an item should span.
*/
func (self *_TraitComboBox) SetRowSpanColumn(row_span int) {
	C.gtk_combo_box_set_row_span_column(self.CPointer, C.gint(row_span))
	return
}

// gtk_combo_box_set_title is not generated due to deprecation attr

/*
Sets the wrap width of @combo_box to be @width. The wrap width is basically
the preferred number of columns when you want the popup to be layed out
in a table.
*/
func (self *_TraitComboBox) SetWrapWidth(width int) {
	C.gtk_combo_box_set_wrap_width(self.CPointer, C.gint(width))
	return
}

type _TraitComboBoxText struct{ CPointer *C.GtkComboBoxText }

/*
Appends @text to the list of strings stored in @combo_box.
If @id is non-%NULL then it is used as the ID of the row.

This is the same as calling gtk_combo_box_text_insert() with a
position of -1.
*/
func (self *_TraitComboBoxText) Append(id string, text string) {
	__cgo__id := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_combo_box_text_append(self.CPointer, __cgo__id, __cgo__text)
	C.free(unsafe.Pointer(__cgo__id))
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Appends @text to the list of strings stored in @combo_box.

This is the same as calling gtk_combo_box_text_insert_text() with a
position of -1.
*/
func (self *_TraitComboBoxText) AppendText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_combo_box_text_append_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Returns the currently active string in @combo_box, or %NULL
if none is selected. If @combo_box contains an entry, this
function will return its contents (which will not necessarily
be an item from the list).
*/
func (self *_TraitComboBoxText) GetActiveText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_combo_box_text_get_active_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Inserts @text at @position in the list of strings stored in @combo_box.
If @id is non-%NULL then it is used as the ID of the row.  See
#GtkComboBox:id-column.

If @position is negative then @text is appended.
*/
func (self *_TraitComboBoxText) Insert(position int, id string, text string) {
	__cgo__id := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_combo_box_text_insert(self.CPointer, C.gint(position), __cgo__id, __cgo__text)
	C.free(unsafe.Pointer(__cgo__id))
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Inserts @text at @position in the list of strings stored in @combo_box.

If @position is negative then @text is appended.

This is the same as calling gtk_combo_box_text_insert() with a %NULL
ID string.
*/
func (self *_TraitComboBoxText) InsertText(position int, text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_combo_box_text_insert_text(self.CPointer, C.gint(position), __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Prepends @text to the list of strings stored in @combo_box.
If @id is non-%NULL then it is used as the ID of the row.

This is the same as calling gtk_combo_box_text_insert() with a
position of 0.
*/
func (self *_TraitComboBoxText) Prepend(id string, text string) {
	__cgo__id := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_combo_box_text_prepend(self.CPointer, __cgo__id, __cgo__text)
	C.free(unsafe.Pointer(__cgo__id))
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Prepends @text to the list of strings stored in @combo_box.

This is the same as calling gtk_combo_box_text_insert_text() with a
position of 0.
*/
func (self *_TraitComboBoxText) PrependText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_combo_box_text_prepend_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Removes the string at @position from @combo_box.
*/
func (self *_TraitComboBoxText) Remove(position int) {
	C.gtk_combo_box_text_remove(self.CPointer, C.gint(position))
	return
}

/*
Removes all the text entries from the combo box.
*/
func (self *_TraitComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all(self.CPointer)
	return
}

type _TraitContainer struct{ CPointer *C.GtkContainer }

/*
Adds @widget to @container. Typically used for simple containers
such as #GtkWindow, #GtkFrame, or #GtkButton; for more complicated
layout containers such as #GtkBox or #GtkGrid, this function will
pick default packing parameters that may not be correct.  So
consider functions such as gtk_box_pack_start() and
gtk_grid_attach() as an alternative to gtk_container_add() in
those cases. A widget may be added to only one container at a time;
you can’t place the same widget inside two different containers.

Note that some containers, such as #GtkScrolledWindow or #GtkListBox,
may add intermediate children between the added widget and the
container.
*/
func (self *_TraitContainer) Add(widget *Widget) {
	C.gtk_container_add(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

// gtk_container_add_with_properties is not generated due to varargs

func (self *_TraitContainer) CheckResize() {
	C.gtk_container_check_resize(self.CPointer)
	return
}

// gtk_container_child_get is not generated due to varargs

/*
Gets the value of a child property for @child and @container.
*/
func (self *_TraitContainer) ChildGetProperty(child *Widget, property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.gtk_container_child_get_property(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

// gtk_container_child_get_valist is not generated due to varargs

/*
Emits a #GtkWidget::child-notify signal for the
[child property][child-properties]
@child_property on widget.

This is an analogue of g_object_notify() for child properties.

Also see gtk_widget_child_notify().
*/
func (self *_TraitContainer) ChildNotify(child *Widget, child_property string) {
	__cgo__child_property := (*C.gchar)(unsafe.Pointer(C.CString(child_property)))
	C.gtk_container_child_notify(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__child_property)
	C.free(unsafe.Pointer(__cgo__child_property))
	return
}

// gtk_container_child_set is not generated due to varargs

/*
Sets a child property for @child and @container.
*/
func (self *_TraitContainer) ChildSetProperty(child *Widget, property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.gtk_container_child_set_property(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

// gtk_container_child_set_valist is not generated due to varargs

/*
Returns the type of the children supported by the container.

Note that this may return %G_TYPE_NONE to indicate that no more
children can be added, e.g. for a #GtkPaned which already has two
children.
*/
func (self *_TraitContainer) ChildType() (return__ C.GType) {
	return__ = C.gtk_container_child_type(self.CPointer)
	return
}

/*
Invokes @callback on each child of @container, including children
that are considered “internal” (implementation details of the
container). “Internal” children generally weren’t added by the user
of the container, but were added by the container implementation
itself.  Most applications should use gtk_container_foreach(),
rather than gtk_container_forall().
*/
func (self *_TraitContainer) Forall(callback C.GtkCallback, callback_data unsafe.Pointer) {
	C.gtk_container_forall(self.CPointer, callback, (C.gpointer)(callback_data))
	return
}

/*
Invokes @callback on each non-internal child of @container. See
gtk_container_forall() for details on what constitutes an
“internal” child.  Most applications should use
gtk_container_foreach(), rather than gtk_container_forall().
*/
func (self *_TraitContainer) Foreach(callback C.GtkCallback, callback_data unsafe.Pointer) {
	C.gtk_container_foreach(self.CPointer, callback, (C.gpointer)(callback_data))
	return
}

/*
Retrieves the border width of the container. See
gtk_container_set_border_width().
*/
func (self *_TraitContainer) GetBorderWidth() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_container_get_border_width(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Returns the container’s non-internal children. See
gtk_container_forall() for details on what constitutes an "internal" child.
*/
func (self *_TraitContainer) GetChildren() (return__ *C.GList) {
	return__ = C.gtk_container_get_children(self.CPointer)
	return
}

/*
Retrieves the focus chain of the container, if one has been
set explicitly. If no focus chain has been explicitly
set, GTK+ computes the focus chain based on the positions
of the children. In that case, GTK+ stores %NULL in
@focusable_widgets and returns %FALSE.
*/
func (self *_TraitContainer) GetFocusChain() (focusable_widgets *C.GList, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_container_get_focus_chain(self.CPointer, &focusable_widgets)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current focus child widget inside @container. This is not the
currently focused widget. That can be obtained by calling
gtk_window_get_focus().
*/
func (self *_TraitContainer) GetFocusChild() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_container_get_focus_child(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the horizontal focus adjustment for the container. See
gtk_container_set_focus_hadjustment ().
*/
func (self *_TraitContainer) GetFocusHadjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_container_get_focus_hadjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the vertical focus adjustment for the container. See
gtk_container_set_focus_vadjustment().
*/
func (self *_TraitContainer) GetFocusVadjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_container_get_focus_vadjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns a newly created widget path representing all the widget hierarchy
from the toplevel down to and including @child.
*/
func (self *_TraitContainer) GetPathForChild(child *Widget) (return__ *WidgetPath) {
	var __cgo__return__ *C.GtkWidgetPath
	__cgo__return__ = C.gtk_container_get_path_for_child(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = (*WidgetPath)(unsafe.Pointer(__cgo__return__))
	return
}

// gtk_container_get_resize_mode is not generated due to deprecation attr

/*
When a container receives a call to the draw function, it must send
synthetic #GtkWidget::draw calls to all children that don’t have their
own #GdkWindows. This function provides a convenient way of doing this.
A container, when it receives a call to its #GtkWidget::draw function,
calls gtk_container_propagate_draw() once for each child, passing in
the @cr the container received.

gtk_container_propagate_draw() takes care of translating the origin of @cr,
and deciding whether the draw needs to be sent to the child. It is a
convenient and optimized way of getting the same effect as calling
gtk_widget_draw() on the child directly.

In most cases, a container can simply either inherit the
#GtkWidget::draw implementation from #GtkContainer, or do some drawing
and then chain to the ::draw implementation from #GtkContainer.
*/
func (self *_TraitContainer) PropagateDraw(child *Widget, cr *C.cairo_t) {
	C.gtk_container_propagate_draw(self.CPointer, (*C.GtkWidget)(child.CPointer), cr)
	return
}

/*
Removes @widget from @container. @widget must be inside @container.
Note that @container will own a reference to @widget, and that this
may be the last reference held; so removing a widget from its
container can destroy that widget. If you want to use @widget
again, you need to add a reference to it while it’s not inside
a container, using g_object_ref(). If you don’t want to use @widget
again it’s usually more efficient to simply destroy it directly
using gtk_widget_destroy() since this will remove it from the
container and help break any circular reference count cycles.
*/
func (self *_TraitContainer) Remove(widget *Widget) {
	C.gtk_container_remove(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

// gtk_container_resize_children is not generated due to deprecation attr

/*
Sets the border width of the container.

The border width of a container is the amount of space to leave
around the outside of the container. The only exception to this is
#GtkWindow; because toplevel windows can’t leave space outside,
they leave the space inside. The border is added on all sides of
the container. To add space to only one side, one approach is to
create a #GtkAlignment widget, call gtk_widget_set_size_request()
to give it a size, and place it on the side of the container as
a spacer.
*/
func (self *_TraitContainer) SetBorderWidth(border_width uint) {
	C.gtk_container_set_border_width(self.CPointer, C.guint(border_width))
	return
}

/*
Sets a focus chain, overriding the one computed automatically by GTK+.

In principle each widget in the chain should be a descendant of the
container, but this is not enforced by this method, since it’s allowed
to set the focus chain before you pack the widgets, or have a widget
in the chain that isn’t always packed. The necessary checks are done
when the focus chain is actually traversed.
*/
func (self *_TraitContainer) SetFocusChain(focusable_widgets *C.GList) {
	C.gtk_container_set_focus_chain(self.CPointer, focusable_widgets)
	return
}

/*
Sets, or unsets if @child is %NULL, the focused child of @container.

This function emits the GtkContainer::set_focus_child signal of
@container. Implementations of #GtkContainer can override the
default behaviour by overriding the class closure of this signal.

This is function is mostly meant to be used by widgets. Applications can use
gtk_widget_grab_focus() to manualy set the focus to a specific widget.
*/
func (self *_TraitContainer) SetFocusChild(child *Widget) {
	C.gtk_container_set_focus_child(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Hooks up an adjustment to focus handling in a container, so when a child
of the container is focused, the adjustment is scrolled to show that
widget. This function sets the horizontal alignment.
See gtk_scrolled_window_get_hadjustment() for a typical way of obtaining
the adjustment and gtk_container_set_focus_vadjustment() for setting
the vertical adjustment.

The adjustments have to be in pixel units and in the same coordinate
system as the allocation for immediate children of the container.
*/
func (self *_TraitContainer) SetFocusHadjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_hadjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
Hooks up an adjustment to focus handling in a container, so when a
child of the container is focused, the adjustment is scrolled to
show that widget. This function sets the vertical alignment. See
gtk_scrolled_window_get_vadjustment() for a typical way of obtaining
the adjustment and gtk_container_set_focus_hadjustment() for setting
the horizontal adjustment.

The adjustments have to be in pixel units and in the same coordinate
system as the allocation for immediate children of the container.
*/
func (self *_TraitContainer) SetFocusVadjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_vadjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
Sets the @reallocate_redraws flag of the container to the given value.

Containers requesting reallocation redraws get automatically
redrawn if any of their children changed allocation.
*/
func (self *_TraitContainer) SetReallocateRedraws(needs_redraws bool) {
	__cgo__needs_redraws := C.gboolean(0)
	if needs_redraws {
		__cgo__needs_redraws = C.gboolean(1)
	}
	C.gtk_container_set_reallocate_redraws(self.CPointer, __cgo__needs_redraws)
	return
}

// gtk_container_set_resize_mode is not generated due to deprecation attr

/*
Removes a focus chain explicitly set with gtk_container_set_focus_chain().
*/
func (self *_TraitContainer) UnsetFocusChain() {
	C.gtk_container_unset_focus_chain(self.CPointer)
	return
}

type _TraitCssProvider struct{ CPointer *C.GtkCssProvider }

/*
Loads @data into @css_provider, making it clear any previously loaded
information.
*/
func (self *_TraitCssProvider) LoadFromData(data []byte, length int64) (return__ bool, __err__ error) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_css_provider_load_from_data(self.CPointer, (*C.gchar)(unsafe.Pointer(__header__data.Data)), C.gssize(length), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Loads the data contained in @file into @css_provider, making it
clear any previously loaded information.
*/
func (self *_TraitCssProvider) LoadFromFile(file *C.GFile) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_css_provider_load_from_file(self.CPointer, file, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Loads the data contained in @path into @css_provider, making it clear
any previously loaded information.
*/
func (self *_TraitCssProvider) LoadFromPath(path string) (return__ bool, __err__ error) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_css_provider_load_from_path(self.CPointer, __cgo__path, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__path))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Converts the @provider into a string representation in CSS
format.

Using gtk_css_provider_load_from_data() with the return value
from this function on a new provider created with
gtk_css_provider_new() will basically create a duplicate of
this @provider.
*/
func (self *_TraitCssProvider) ToString() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.gtk_css_provider_to_string(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

type _TraitDialog struct{ CPointer *C.GtkDialog }

/*
Adds an activatable widget to the action area of a #GtkDialog,
connecting a signal handler that will emit the #GtkDialog::response
signal on the dialog when the widget is activated. The widget is
appended to the end of the dialog’s action area. If you want to add a
non-activatable widget, simply pack it into the @action_area field
of the #GtkDialog struct.
*/
func (self *_TraitDialog) AddActionWidget(child *Widget, response_id int) {
	C.gtk_dialog_add_action_widget(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(response_id))
	return
}

/*
Adds a button with the given text and sets things up so that
clicking the button will emit the #GtkDialog::response signal with
the given @response_id. The button is appended to the end of the
dialog’s action area. The button widget is returned, but usually
you don’t need it.
*/
func (self *_TraitDialog) AddButton(button_text string, response_id int) (return__ *Widget) {
	__cgo__button_text := (*C.gchar)(unsafe.Pointer(C.CString(button_text)))
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_dialog_add_button(self.CPointer, __cgo__button_text, C.gint(response_id))
	C.free(unsafe.Pointer(__cgo__button_text))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

// gtk_dialog_add_buttons is not generated due to varargs

// gtk_dialog_get_action_area is not generated due to deprecation attr

/*
Returns the content area of @dialog.
*/
func (self *_TraitDialog) GetContentArea() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_dialog_get_content_area(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the header bar of @dialog. Note that the
headerbar is only used by the dialog if the
#GtkDialog:use-header-bar property is %TRUE.
*/
func (self *_TraitDialog) GetHeaderBar() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_dialog_get_header_bar(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the response id of a widget in the action area
of a dialog.
*/
func (self *_TraitDialog) GetResponseForWidget(widget *Widget) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_dialog_get_response_for_widget(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Gets the widget button that uses the given response ID in the action area
of a dialog.
*/
func (self *_TraitDialog) GetWidgetForResponse(response_id int) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_dialog_get_widget_for_response(self.CPointer, C.gint(response_id))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Emits the #GtkDialog::response signal with the given response ID.
Used to indicate that the user has responded to the dialog in some way;
typically either you or gtk_dialog_run() will be monitoring the
::response signal and take appropriate action.
*/
func (self *_TraitDialog) Response(response_id int) {
	C.gtk_dialog_response(self.CPointer, C.gint(response_id))
	return
}

/*
Blocks in a recursive main loop until the @dialog either emits the
#GtkDialog::response signal, or is destroyed. If the dialog is
destroyed during the call to gtk_dialog_run(), gtk_dialog_run() returns
#GTK_RESPONSE_NONE. Otherwise, it returns the response ID from the
::response signal emission.

Before entering the recursive main loop, gtk_dialog_run() calls
gtk_widget_show() on the dialog for you. Note that you still
need to show any children of the dialog yourself.

During gtk_dialog_run(), the default behavior of #GtkWidget::delete-event
is disabled; if the dialog receives ::delete_event, it will not be
destroyed as windows usually are, and gtk_dialog_run() will return
#GTK_RESPONSE_DELETE_EVENT. Also, during gtk_dialog_run() the dialog
will be modal. You can force gtk_dialog_run() to return at any time by
calling gtk_dialog_response() to emit the ::response signal. Destroying
the dialog during gtk_dialog_run() is a very bad idea, because your
post-run code won’t know whether the dialog was destroyed or not.

After gtk_dialog_run() returns, you are responsible for hiding or
destroying the dialog if you wish to do so.

Typical usage of this function might be:
|[<!-- language="C" -->
  gint result = gtk_dialog_run (GTK_DIALOG (dialog));
  switch (result)
    {
      case GTK_RESPONSE_ACCEPT:
         do_application_specific_something ();
         break;
      default:
         do_nothing_since_dialog_was_cancelled ();
         break;
    }
  gtk_widget_destroy (dialog);
]|

Note that even though the recursive main loop gives the effect of a
modal dialog (it prevents the user from interacting with other
windows in the same window group while the dialog is run), callbacks
such as timeouts, IO channel watches, DND drops, etc, will
be triggered during a gtk_dialog_run() call.
*/
func (self *_TraitDialog) Run() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_dialog_run(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

// gtk_dialog_set_alternative_button_order is not generated due to deprecation attr

// gtk_dialog_set_alternative_button_order_from_array is not generated due to deprecation attr

/*
Sets the last widget in the dialog’s action area with the given @response_id
as the default widget for the dialog. Pressing “Enter” normally activates
the default widget.
*/
func (self *_TraitDialog) SetDefaultResponse(response_id int) {
	C.gtk_dialog_set_default_response(self.CPointer, C.gint(response_id))
	return
}

/*
Calls `gtk_widget_set_sensitive (widget, @setting)`
for each widget in the dialog’s action area with the given @response_id.
A convenient way to sensitize/desensitize dialog buttons.
*/
func (self *_TraitDialog) SetResponseSensitive(response_id int, setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_dialog_set_response_sensitive(self.CPointer, C.gint(response_id), __cgo__setting)
	return
}

type _TraitDrawingArea struct{ CPointer *C.GtkDrawingArea }

type _TraitEntry struct{ CPointer *C.GtkEntry }

/*
Retrieves the value set by gtk_entry_set_activates_default().
*/
func (self *_TraitEntry) GetActivatesDefault() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_get_activates_default(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_entry_set_alignment().
*/
func (self *_TraitEntry) GetAlignment() (return__ float32) {
	var __cgo__return__ C.gfloat
	__cgo__return__ = C.gtk_entry_get_alignment(self.CPointer)
	return__ = float32(__cgo__return__)
	return
}

/*
Gets the attribute list that was set on the entry using
gtk_entry_set_attributes(), if any.
*/
func (self *_TraitEntry) GetAttributes() (return__ *C.PangoAttrList) {
	return__ = C.gtk_entry_get_attributes(self.CPointer)
	return
}

/*
Get the #GtkEntryBuffer object which holds the text for
this widget.
*/
func (self *_TraitEntry) GetBuffer() (return__ *EntryBuffer) {
	var __cgo__return__ *C.GtkEntryBuffer
	__cgo__return__ = C.gtk_entry_get_buffer(self.CPointer)
	return__ = NewEntryBufferFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the auxiliary completion object currently in use by @entry.
*/
func (self *_TraitEntry) GetCompletion() (return__ *EntryCompletion) {
	var __cgo__return__ *C.GtkEntryCompletion
	__cgo__return__ = C.gtk_entry_get_completion(self.CPointer)
	return__ = NewEntryCompletionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the index of the icon which is the source of the current
DND operation, or -1.

This function is meant to be used in a #GtkWidget::drag-data-get
callback.
*/
func (self *_TraitEntry) GetCurrentIconDragSource() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_get_current_icon_drag_source(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the horizontal cursor adjustment for the entry.
See gtk_entry_set_cursor_hadjustment().
*/
func (self *_TraitEntry) GetCursorHadjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_entry_get_cursor_hadjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the value set by gtk_entry_set_has_frame().
*/
func (self *_TraitEntry) GetHasFrame() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_get_has_frame(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the icon is activatable.
*/
func (self *_TraitEntry) GetIconActivatable(icon_pos C.GtkEntryIconPosition) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_get_icon_activatable(self.CPointer, icon_pos)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the area where entry’s icon at @icon_pos is drawn.
This function is useful when drawing something to the
entry in a draw callback.

If the entry is not realized or has no icon at the given position,
@icon_area is filled with zeros.

See also gtk_entry_get_text_area()
*/
func (self *_TraitEntry) GetIconArea(icon_pos C.GtkEntryIconPosition) (icon_area C.GdkRectangle) {
	C.gtk_entry_get_icon_area(self.CPointer, icon_pos, &icon_area)
	return
}

/*
Finds the icon at the given position and return its index. The
position’s coordinates are relative to the @entry’s top left corner.
If @x, @y doesn’t lie inside an icon, -1 is returned.
This function is intended for use in a #GtkWidget::query-tooltip
signal handler.
*/
func (self *_TraitEntry) GetIconAtPos(x int, y int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_get_icon_at_pos(self.CPointer, C.gint(x), C.gint(y))
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the #GIcon used for the icon, or %NULL if there is
no icon or if the icon was set by some other method (e.g., by
stock, pixbuf, or icon name).
*/
func (self *_TraitEntry) GetIconGicon(icon_pos C.GtkEntryIconPosition) (return__ *C.GIcon) {
	return__ = C.gtk_entry_get_icon_gicon(self.CPointer, icon_pos)
	return
}

/*
Retrieves the icon name used for the icon, or %NULL if there is
no icon or if the icon was set by some other method (e.g., by
pixbuf, stock or gicon).
*/
func (self *_TraitEntry) GetIconName(icon_pos C.GtkEntryIconPosition) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_entry_get_icon_name(self.CPointer, icon_pos)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the image used for the icon.

Unlike the other methods of setting and getting icon data, this
method will work regardless of whether the icon was set using a
#GdkPixbuf, a #GIcon, a stock item, or an icon name.
*/
func (self *_TraitEntry) GetIconPixbuf(icon_pos C.GtkEntryIconPosition) (return__ *C.GdkPixbuf) {
	return__ = C.gtk_entry_get_icon_pixbuf(self.CPointer, icon_pos)
	return
}

/*
Returns whether the icon appears sensitive or insensitive.
*/
func (self *_TraitEntry) GetIconSensitive(icon_pos C.GtkEntryIconPosition) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_get_icon_sensitive(self.CPointer, icon_pos)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_entry_get_icon_stock is not generated due to deprecation attr

/*
Gets the type of representation being used by the icon
to store image data. If the icon has no image data,
the return value will be %GTK_IMAGE_EMPTY.
*/
func (self *_TraitEntry) GetIconStorageType(icon_pos C.GtkEntryIconPosition) (return__ C.GtkImageType) {
	return__ = C.gtk_entry_get_icon_storage_type(self.CPointer, icon_pos)
	return
}

/*
Gets the contents of the tooltip on the icon at the specified
position in @entry.
*/
func (self *_TraitEntry) GetIconTooltipMarkup(icon_pos C.GtkEntryIconPosition) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_entry_get_icon_tooltip_markup(self.CPointer, icon_pos)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the contents of the tooltip on the icon at the specified
position in @entry.
*/
func (self *_TraitEntry) GetIconTooltipText(icon_pos C.GtkEntryIconPosition) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_entry_get_icon_tooltip_text(self.CPointer, icon_pos)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

// gtk_entry_get_inner_border is not generated due to deprecation attr

/*
Gets the value of the #GtkEntry:input-hints property.
*/
func (self *_TraitEntry) GetInputHints() (return__ C.GtkInputHints) {
	return__ = C.gtk_entry_get_input_hints(self.CPointer)
	return
}

/*
Gets the value of the #GtkEntry:input-purpose property.
*/
func (self *_TraitEntry) GetInputPurpose() (return__ C.GtkInputPurpose) {
	return__ = C.gtk_entry_get_input_purpose(self.CPointer)
	return
}

/*
Retrieves the character displayed in place of the real characters
for entries with visibility set to false. See gtk_entry_set_invisible_char().
*/
func (self *_TraitEntry) GetInvisibleChar() (return__ rune) {
	var __cgo__return__ C.gunichar
	__cgo__return__ = C.gtk_entry_get_invisible_char(self.CPointer)
	return__ = rune(__cgo__return__)
	return
}

/*
Gets the #PangoLayout used to display the entry.
The layout is useful to e.g. convert text positions to
pixel positions, in combination with gtk_entry_get_layout_offsets().
The returned layout is owned by the entry and must not be
modified or freed by the caller.

Keep in mind that the layout text may contain a preedit string, so
gtk_entry_layout_index_to_text_index() and
gtk_entry_text_index_to_layout_index() are needed to convert byte
indices in the layout to byte indices in the entry contents.
*/
func (self *_TraitEntry) GetLayout() (return__ *C.PangoLayout) {
	return__ = C.gtk_entry_get_layout(self.CPointer)
	return
}

/*
Obtains the position of the #PangoLayout used to render text
in the entry, in widget coordinates. Useful if you want to line
up the text in an entry with some other text, e.g. when using the
entry to implement editable cells in a sheet widget.

Also useful to convert mouse events into coordinates inside the
#PangoLayout, e.g. to take some action if some part of the entry text
is clicked.

Note that as the user scrolls around in the entry the offsets will
change; you’ll need to connect to the “notify::scroll-offset”
signal to track this. Remember when using the #PangoLayout
functions you need to convert to and from pixels using
PANGO_PIXELS() or #PANGO_SCALE.

Keep in mind that the layout text may contain a preedit string, so
gtk_entry_layout_index_to_text_index() and
gtk_entry_text_index_to_layout_index() are needed to convert byte
indices in the layout to byte indices in the entry contents.
*/
func (self *_TraitEntry) GetLayoutOffsets() (x int, y int) {
	var __cgo__x C.gint
	var __cgo__y C.gint
	C.gtk_entry_get_layout_offsets(self.CPointer, &__cgo__x, &__cgo__y)
	x = int(__cgo__x)
	y = int(__cgo__y)
	return
}

/*
Retrieves the maximum allowed length of the text in
@entry. See gtk_entry_set_max_length().

This is equivalent to:

|[<!-- language="C" -->
GtkEntryBuffer *buffer;
buffer = gtk_entry_get_buffer (entry);
gtk_entry_buffer_get_max_length (buffer);
]|
*/
func (self *_TraitEntry) GetMaxLength() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_get_max_length(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the desired maximum width of @entry, in characters.
See gtk_entry_set_max_width_chars().
*/
func (self *_TraitEntry) GetMaxWidthChars() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_get_max_width_chars(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value set by gtk_entry_set_overwrite_mode().
*/
func (self *_TraitEntry) GetOverwriteMode() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_get_overwrite_mode(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the text that will be displayed when @entry is empty and unfocused
*/
func (self *_TraitEntry) GetPlaceholderText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_entry_get_placeholder_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the current fraction of the task that’s been completed.
See gtk_entry_set_progress_fraction().
*/
func (self *_TraitEntry) GetProgressFraction() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_entry_get_progress_fraction(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Retrieves the pulse step set with gtk_entry_set_progress_pulse_step().
*/
func (self *_TraitEntry) GetProgressPulseStep() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_entry_get_progress_pulse_step(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the tabstops that were set on the entry using gtk_entry_set_tabs(), if
any.
*/
func (self *_TraitEntry) GetTabs() (return__ *C.PangoTabArray) {
	return__ = C.gtk_entry_get_tabs(self.CPointer)
	return
}

/*
Retrieves the contents of the entry widget.
See also gtk_editable_get_chars().

This is equivalent to:

|[<!-- language="C" -->
GtkEntryBuffer *buffer;
buffer = gtk_entry_get_buffer (entry);
gtk_entry_buffer_get_text (buffer);
]|
*/
func (self *_TraitEntry) GetText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_entry_get_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the area where the entry’s text is drawn. This function is
useful when drawing something to the entry in a draw callback.

If the entry is not realized, @text_area is filled with zeros.

See also gtk_entry_get_icon_area().
*/
func (self *_TraitEntry) GetTextArea() (text_area C.GdkRectangle) {
	C.gtk_entry_get_text_area(self.CPointer, &text_area)
	return
}

/*
Retrieves the current length of the text in
@entry.

This is equivalent to:

|[<!-- language="C" -->
GtkEntryBuffer *buffer;
buffer = gtk_entry_get_buffer (entry);
gtk_entry_buffer_get_length (buffer);
]|
*/
func (self *_TraitEntry) GetTextLength() (return__ uint16) {
	var __cgo__return__ C.guint16
	__cgo__return__ = C.gtk_entry_get_text_length(self.CPointer)
	return__ = uint16(__cgo__return__)
	return
}

/*
Retrieves whether the text in @entry is visible. See
gtk_entry_set_visibility().
*/
func (self *_TraitEntry) GetVisibility() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_get_visibility(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_entry_set_width_chars().
*/
func (self *_TraitEntry) GetWidthChars() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_get_width_chars(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Allow the #GtkEntry input method to internally handle key press
and release events. If this function returns %TRUE, then no further
processing should be done for this key event. See
gtk_im_context_filter_keypress().

Note that you are expected to call this function from your handler
when overriding key event handling. This is needed in the case when
you need to insert your own key handling between the input method
and the default key event handling of the #GtkEntry.
See gtk_text_view_reset_im_context() for an example of use.
*/
func (self *_TraitEntry) ImContextFilterKeypress(event *C.GdkEventKey) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_im_context_filter_keypress(self.CPointer, event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts from a position in the entry contents (returned
by gtk_entry_get_text()) to a position in the
entry’s #PangoLayout (returned by gtk_entry_get_layout(),
with text retrieved via pango_layout_get_text()).
*/
func (self *_TraitEntry) LayoutIndexToTextIndex(layout_index int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_layout_index_to_text_index(self.CPointer, C.gint(layout_index))
	return__ = int(__cgo__return__)
	return
}

/*
Indicates that some progress is made, but you don’t know how much.
Causes the entry’s progress indicator to enter “activity mode,”
where a block bounces back and forth. Each call to
gtk_entry_progress_pulse() causes the block to move by a little bit
(the amount of movement per pulse is determined by
gtk_entry_set_progress_pulse_step()).
*/
func (self *_TraitEntry) ProgressPulse() {
	C.gtk_entry_progress_pulse(self.CPointer)
	return
}

/*
Reset the input method context of the entry if needed.

This can be necessary in the case where modifying the buffer
would confuse on-going input method behavior.
*/
func (self *_TraitEntry) ResetImContext() {
	C.gtk_entry_reset_im_context(self.CPointer)
	return
}

/*
If @setting is %TRUE, pressing Enter in the @entry will activate the default
widget for the window containing the entry. This usually means that
the dialog box containing the entry will be closed, since the default
widget is usually one of the dialog buttons.

(For experts: if @setting is %TRUE, the entry calls
gtk_window_activate_default() on the window containing the entry, in
the default handler for the #GtkEntry::activate signal.)
*/
func (self *_TraitEntry) SetActivatesDefault(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_entry_set_activates_default(self.CPointer, __cgo__setting)
	return
}

/*
Sets the alignment for the contents of the entry. This controls
the horizontal positioning of the contents when the displayed
text is shorter than the width of the entry.
*/
func (self *_TraitEntry) SetAlignment(xalign float32) {
	C.gtk_entry_set_alignment(self.CPointer, C.gfloat(xalign))
	return
}

/*
Sets a #PangoAttrList; the attributes in the list are applied to the
entry text.
*/
func (self *_TraitEntry) SetAttributes(attrs *C.PangoAttrList) {
	C.gtk_entry_set_attributes(self.CPointer, attrs)
	return
}

/*
Set the #GtkEntryBuffer object which holds the text for
this widget.
*/
func (self *_TraitEntry) SetBuffer(buffer *EntryBuffer) {
	C.gtk_entry_set_buffer(self.CPointer, (*C.GtkEntryBuffer)(buffer.CPointer))
	return
}

/*
Sets @completion to be the auxiliary completion object to use with @entry.
All further configuration of the completion mechanism is done on
@completion using the #GtkEntryCompletion API. Completion is disabled if
@completion is set to %NULL.
*/
func (self *_TraitEntry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(self.CPointer, (*C.GtkEntryCompletion)(completion.CPointer))
	return
}

/*
Hooks up an adjustment to the cursor position in an entry, so that when
the cursor is moved, the adjustment is scrolled to show that position.
See gtk_scrolled_window_get_hadjustment() for a typical way of obtaining
the adjustment.

The adjustment has to be in pixel units and in the same coordinate system
as the entry.
*/
func (self *_TraitEntry) SetCursorHadjustment(adjustment *Adjustment) {
	C.gtk_entry_set_cursor_hadjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
Sets whether the entry has a beveled frame around it.
*/
func (self *_TraitEntry) SetHasFrame(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_entry_set_has_frame(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether the icon is activatable.
*/
func (self *_TraitEntry) SetIconActivatable(icon_pos C.GtkEntryIconPosition, activatable bool) {
	__cgo__activatable := C.gboolean(0)
	if activatable {
		__cgo__activatable = C.gboolean(1)
	}
	C.gtk_entry_set_icon_activatable(self.CPointer, icon_pos, __cgo__activatable)
	return
}

/*
Sets up the icon at the given position so that GTK+ will start a drag
operation when the user clicks and drags the icon.

To handle the drag operation, you need to connect to the usual
#GtkWidget::drag-data-get (or possibly #GtkWidget::drag-data-delete)
signal, and use gtk_entry_get_current_icon_drag_source() in
your signal handler to find out if the drag was started from
an icon.

By default, GTK+ uses the icon as the drag icon. You can use the
#GtkWidget::drag-begin signal to set a different icon. Note that you
have to use g_signal_connect_after() to ensure that your signal handler
gets executed after the default handler.
*/
func (self *_TraitEntry) SetIconDragSource(icon_pos C.GtkEntryIconPosition, target_list *TargetList, actions C.GdkDragAction) {
	C.gtk_entry_set_icon_drag_source(self.CPointer, icon_pos, (*C.GtkTargetList)(unsafe.Pointer(target_list)), actions)
	return
}

/*
Sets the icon shown in the entry at the specified position
from the current icon theme.
If the icon isn’t known, a “broken image” icon will be displayed
instead.

If @icon is %NULL, no icon will be shown in the specified position.
*/
func (self *_TraitEntry) SetIconFromGicon(icon_pos C.GtkEntryIconPosition, icon *C.GIcon) {
	C.gtk_entry_set_icon_from_gicon(self.CPointer, icon_pos, icon)
	return
}

/*
Sets the icon shown in the entry at the specified position
from the current icon theme.

If the icon name isn’t known, a “broken image” icon will be displayed
instead.

If @icon_name is %NULL, no icon will be shown in the specified position.
*/
func (self *_TraitEntry) SetIconFromIconName(icon_pos C.GtkEntryIconPosition, icon_name string) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_entry_set_icon_from_icon_name(self.CPointer, icon_pos, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

/*
Sets the icon shown in the specified position using a pixbuf.

If @pixbuf is %NULL, no icon will be shown in the specified position.
*/
func (self *_TraitEntry) SetIconFromPixbuf(icon_pos C.GtkEntryIconPosition, pixbuf *C.GdkPixbuf) {
	C.gtk_entry_set_icon_from_pixbuf(self.CPointer, icon_pos, pixbuf)
	return
}

// gtk_entry_set_icon_from_stock is not generated due to deprecation attr

/*
Sets the sensitivity for the specified icon.
*/
func (self *_TraitEntry) SetIconSensitive(icon_pos C.GtkEntryIconPosition, sensitive bool) {
	__cgo__sensitive := C.gboolean(0)
	if sensitive {
		__cgo__sensitive = C.gboolean(1)
	}
	C.gtk_entry_set_icon_sensitive(self.CPointer, icon_pos, __cgo__sensitive)
	return
}

/*
Sets @tooltip as the contents of the tooltip for the icon at
the specified position. @tooltip is assumed to be marked up with
the [Pango text markup language][PangoMarkupFormat].

Use %NULL for @tooltip to remove an existing tooltip.

See also gtk_widget_set_tooltip_markup() and
gtk_entry_set_icon_tooltip_text().
*/
func (self *_TraitEntry) SetIconTooltipMarkup(icon_pos C.GtkEntryIconPosition, tooltip string) {
	__cgo__tooltip := (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
	C.gtk_entry_set_icon_tooltip_markup(self.CPointer, icon_pos, __cgo__tooltip)
	C.free(unsafe.Pointer(__cgo__tooltip))
	return
}

/*
Sets @tooltip as the contents of the tooltip for the icon
at the specified position.

Use %NULL for @tooltip to remove an existing tooltip.

See also gtk_widget_set_tooltip_text() and
gtk_entry_set_icon_tooltip_markup().
*/
func (self *_TraitEntry) SetIconTooltipText(icon_pos C.GtkEntryIconPosition, tooltip string) {
	__cgo__tooltip := (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
	C.gtk_entry_set_icon_tooltip_text(self.CPointer, icon_pos, __cgo__tooltip)
	C.free(unsafe.Pointer(__cgo__tooltip))
	return
}

// gtk_entry_set_inner_border is not generated due to deprecation attr

/*
Sets the #GtkEntry:input-hints property, which
allows input methods to fine-tune their behaviour.
*/
func (self *_TraitEntry) SetInputHints(hints C.GtkInputHints) {
	C.gtk_entry_set_input_hints(self.CPointer, hints)
	return
}

/*
Sets the #GtkEntry:input-purpose property which
can be used by on-screen keyboards and other input
methods to adjust their behaviour.
*/
func (self *_TraitEntry) SetInputPurpose(purpose C.GtkInputPurpose) {
	C.gtk_entry_set_input_purpose(self.CPointer, purpose)
	return
}

/*
Sets the character to use in place of the actual text when
gtk_entry_set_visibility() has been called to set text visibility
to %FALSE. i.e. this is the character used in “password mode” to
show the user how many characters have been typed. By default, GTK+
picks the best invisible char available in the current font. If you
set the invisible char to 0, then the user will get no feedback
at all; there will be no text on the screen as they type.
*/
func (self *_TraitEntry) SetInvisibleChar(ch rune) {
	C.gtk_entry_set_invisible_char(self.CPointer, C.gunichar(ch))
	return
}

/*
Sets the maximum allowed length of the contents of the widget. If
the current contents are longer than the given length, then they
will be truncated to fit.

This is equivalent to:

|[<!-- language="C" -->
GtkEntryBuffer *buffer;
buffer = gtk_entry_get_buffer (entry);
gtk_entry_buffer_set_max_length (buffer, max);
]|
*/
func (self *_TraitEntry) SetMaxLength(max int) {
	C.gtk_entry_set_max_length(self.CPointer, C.gint(max))
	return
}

/*
Sets the desired maximum width in characters of @entry.
*/
func (self *_TraitEntry) SetMaxWidthChars(n_chars int) {
	C.gtk_entry_set_max_width_chars(self.CPointer, C.gint(n_chars))
	return
}

/*
Sets whether the text is overwritten when typing in the #GtkEntry.
*/
func (self *_TraitEntry) SetOverwriteMode(overwrite bool) {
	__cgo__overwrite := C.gboolean(0)
	if overwrite {
		__cgo__overwrite = C.gboolean(1)
	}
	C.gtk_entry_set_overwrite_mode(self.CPointer, __cgo__overwrite)
	return
}

/*
Sets text to be displayed in @entry when it is empty and unfocused.
This can be used to give a visual hint of the expected contents of
the #GtkEntry.

Note that since the placeholder text gets removed when the entry
received focus, using this feature is a bit problematic if the entry
is given the initial focus in a window. Sometimes this can be
worked around by delaying the initial focus setting until the
first key event arrives.
*/
func (self *_TraitEntry) SetPlaceholderText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_entry_set_placeholder_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Causes the entry’s progress indicator to “fill in” the given
fraction of the bar. The fraction should be between 0.0 and 1.0,
inclusive.
*/
func (self *_TraitEntry) SetProgressFraction(fraction float64) {
	C.gtk_entry_set_progress_fraction(self.CPointer, C.gdouble(fraction))
	return
}

/*
Sets the fraction of total entry width to move the progress
bouncing block for each call to gtk_entry_progress_pulse().
*/
func (self *_TraitEntry) SetProgressPulseStep(fraction float64) {
	C.gtk_entry_set_progress_pulse_step(self.CPointer, C.gdouble(fraction))
	return
}

/*
Sets a #PangoTabArray; the tabstops in the array are applied to the entry
text.
*/
func (self *_TraitEntry) SetTabs(tabs *C.PangoTabArray) {
	C.gtk_entry_set_tabs(self.CPointer, tabs)
	return
}

/*
Sets the text in the widget to the given
value, replacing the current contents.

See gtk_entry_buffer_set_text().
*/
func (self *_TraitEntry) SetText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_entry_set_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Sets whether the contents of the entry are visible or not.
When visibility is set to %FALSE, characters are displayed
as the invisible char, and will also appear that way when
the text in the entry widget is copied elsewhere.

By default, GTK+ picks the best invisible character available
in the current font, but it can be changed with
gtk_entry_set_invisible_char().

Note that you probably want to set #GtkEntry:input-purpose
to %GTK_INPUT_PURPOSE_PASSWORD or %GTK_INPUT_PURPOSE_PIN to
inform input methods about the purpose of this entry,
in addition to setting visibility to %FALSE.
*/
func (self *_TraitEntry) SetVisibility(visible bool) {
	__cgo__visible := C.gboolean(0)
	if visible {
		__cgo__visible = C.gboolean(1)
	}
	C.gtk_entry_set_visibility(self.CPointer, __cgo__visible)
	return
}

/*
Changes the size request of the entry to be about the right size
for @n_chars characters. Note that it changes the size
request, the size can still be affected by
how you pack the widget into containers. If @n_chars is -1, the
size reverts to the default entry size.
*/
func (self *_TraitEntry) SetWidthChars(n_chars int) {
	C.gtk_entry_set_width_chars(self.CPointer, C.gint(n_chars))
	return
}

/*
Converts from a position in the entry’s #PangoLayout (returned by
gtk_entry_get_layout()) to a position in the entry contents
(returned by gtk_entry_get_text()).
*/
func (self *_TraitEntry) TextIndexToLayoutIndex(text_index int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_text_index_to_layout_index(self.CPointer, C.gint(text_index))
	return__ = int(__cgo__return__)
	return
}

/*
Unsets the invisible char previously set with
gtk_entry_set_invisible_char(). So that the
default invisible char is used again.
*/
func (self *_TraitEntry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(self.CPointer)
	return
}

type _TraitEntryBuffer struct{ CPointer *C.GtkEntryBuffer }

/*
Deletes a sequence of characters from the buffer. @n_chars characters are
deleted starting at @position. If @n_chars is negative, then all characters
until the end of the text are deleted.

If @position or @n_chars are out of bounds, then they are coerced to sane
values.

Note that the positions are specified in characters, not bytes.
*/
func (self *_TraitEntryBuffer) DeleteText(position uint, n_chars int) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_entry_buffer_delete_text(self.CPointer, C.guint(position), C.gint(n_chars))
	return__ = uint(__cgo__return__)
	return
}

/*
Used when subclassing #GtkEntryBuffer
*/
func (self *_TraitEntryBuffer) EmitDeletedText(position uint, n_chars uint) {
	C.gtk_entry_buffer_emit_deleted_text(self.CPointer, C.guint(position), C.guint(n_chars))
	return
}

/*
Used when subclassing #GtkEntryBuffer
*/
func (self *_TraitEntryBuffer) EmitInsertedText(position uint, chars string, n_chars uint) {
	__cgo__chars := (*C.gchar)(unsafe.Pointer(C.CString(chars)))
	C.gtk_entry_buffer_emit_inserted_text(self.CPointer, C.guint(position), __cgo__chars, C.guint(n_chars))
	C.free(unsafe.Pointer(__cgo__chars))
	return
}

/*
Retrieves the length in bytes of the buffer.
See gtk_entry_buffer_get_length().
*/
func (self *_TraitEntryBuffer) GetBytes() (return__ int64) {
	var __cgo__return__ C.gsize
	__cgo__return__ = C.gtk_entry_buffer_get_bytes(self.CPointer)
	return__ = int64(__cgo__return__)
	return
}

/*
Retrieves the length in characters of the buffer.
*/
func (self *_TraitEntryBuffer) GetLength() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_entry_buffer_get_length(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Retrieves the maximum allowed length of the text in
@buffer. See gtk_entry_buffer_set_max_length().
*/
func (self *_TraitEntryBuffer) GetMaxLength() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_buffer_get_max_length(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the contents of the buffer.

The memory pointer returned by this call will not change
unless this object emits a signal, or is finalized.
*/
func (self *_TraitEntryBuffer) GetText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_entry_buffer_get_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Inserts @n_chars characters of @chars into the contents of the
buffer, at position @position.

If @n_chars is negative, then characters from chars will be inserted
until a null-terminator is found. If @position or @n_chars are out of
bounds, or the maximum buffer text length is exceeded, then they are
coerced to sane values.

Note that the position and length are in characters, not in bytes.
*/
func (self *_TraitEntryBuffer) InsertText(position uint, chars string, n_chars int) (return__ uint) {
	__cgo__chars := (*C.gchar)(unsafe.Pointer(C.CString(chars)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_entry_buffer_insert_text(self.CPointer, C.guint(position), __cgo__chars, C.gint(n_chars))
	C.free(unsafe.Pointer(__cgo__chars))
	return__ = uint(__cgo__return__)
	return
}

/*
Sets the maximum allowed length of the contents of the buffer. If
the current contents are longer than the given length, then they
will be truncated to fit.
*/
func (self *_TraitEntryBuffer) SetMaxLength(max_length int) {
	C.gtk_entry_buffer_set_max_length(self.CPointer, C.gint(max_length))
	return
}

/*
Sets the text in the buffer.

This is roughly equivalent to calling gtk_entry_buffer_delete_text()
and gtk_entry_buffer_insert_text().

Note that @n_chars is in characters, not in bytes.
*/
func (self *_TraitEntryBuffer) SetText(chars string, n_chars int) {
	__cgo__chars := (*C.gchar)(unsafe.Pointer(C.CString(chars)))
	C.gtk_entry_buffer_set_text(self.CPointer, __cgo__chars, C.gint(n_chars))
	C.free(unsafe.Pointer(__cgo__chars))
	return
}

type _TraitEntryCompletion struct{ CPointer *C.GtkEntryCompletion }

/*
Requests a completion operation, or in other words a refiltering of the
current list with completions, using the current key. The completion list
view will be updated accordingly.
*/
func (self *_TraitEntryCompletion) Complete() {
	C.gtk_entry_completion_complete(self.CPointer)
	return
}

/*
Computes the common prefix that is shared by all rows in @completion
that start with @key. If no row matches @key, %NULL will be returned.
Note that a text column must have been set for this function to work,
see gtk_entry_completion_set_text_column() for details.
*/
func (self *_TraitEntryCompletion) ComputePrefix(key string) (return__ string) {
	__cgo__key := C.CString(key)
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_entry_completion_compute_prefix(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Deletes the action at @index_ from @completion’s action list.
*/
func (self *_TraitEntryCompletion) DeleteAction(index_ int) {
	C.gtk_entry_completion_delete_action(self.CPointer, C.gint(index_))
	return
}

/*
Get the original text entered by the user that triggered
the completion or %NULL if there’s no completion ongoing.
*/
func (self *_TraitEntryCompletion) GetCompletionPrefix() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_entry_completion_get_completion_prefix(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the entry @completion has been attached to.
*/
func (self *_TraitEntryCompletion) GetEntry() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_entry_completion_get_entry(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether the common prefix of the possible completions should
be automatically inserted in the entry.
*/
func (self *_TraitEntryCompletion) GetInlineCompletion() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_completion_get_inline_completion(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if inline-selection mode is turned on.
*/
func (self *_TraitEntryCompletion) GetInlineSelection() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_completion_get_inline_selection(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the minimum key length as set for @completion.
*/
func (self *_TraitEntryCompletion) GetMinimumKeyLength() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_completion_get_minimum_key_length(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the model the #GtkEntryCompletion is using as data source.
Returns %NULL if the model is unset.
*/
func (self *_TraitEntryCompletion) GetModel() (return__ *C.GtkTreeModel) {
	return__ = C.gtk_entry_completion_get_model(self.CPointer)
	return
}

/*
Returns whether the completions should be presented in a popup window.
*/
func (self *_TraitEntryCompletion) GetPopupCompletion() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_completion_get_popup_completion(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the  completion popup window will be resized to the
width of the entry.
*/
func (self *_TraitEntryCompletion) GetPopupSetWidth() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_completion_get_popup_set_width(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the completion popup window will appear even if there is
only a single match.
*/
func (self *_TraitEntryCompletion) GetPopupSingleMatch() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_entry_completion_get_popup_single_match(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the column in the model of @completion to get strings from.
*/
func (self *_TraitEntryCompletion) GetTextColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_entry_completion_get_text_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Inserts an action in @completion’s action item list at position @index_
with markup @markup.
*/
func (self *_TraitEntryCompletion) InsertActionMarkup(index_ int, markup string) {
	__cgo__markup := (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	C.gtk_entry_completion_insert_action_markup(self.CPointer, C.gint(index_), __cgo__markup)
	C.free(unsafe.Pointer(__cgo__markup))
	return
}

/*
Inserts an action in @completion’s action item list at position @index_
with text @text. If you want the action item to have markup, use
gtk_entry_completion_insert_action_markup().
*/
func (self *_TraitEntryCompletion) InsertActionText(index_ int, text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_entry_completion_insert_action_text(self.CPointer, C.gint(index_), __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Requests a prefix insertion.
*/
func (self *_TraitEntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix(self.CPointer)
	return
}

/*
Sets whether the common prefix of the possible completions should
be automatically inserted in the entry.
*/
func (self *_TraitEntryCompletion) SetInlineCompletion(inline_completion bool) {
	__cgo__inline_completion := C.gboolean(0)
	if inline_completion {
		__cgo__inline_completion = C.gboolean(1)
	}
	C.gtk_entry_completion_set_inline_completion(self.CPointer, __cgo__inline_completion)
	return
}

/*
Sets whether it is possible to cycle through the possible completions
inside the entry.
*/
func (self *_TraitEntryCompletion) SetInlineSelection(inline_selection bool) {
	__cgo__inline_selection := C.gboolean(0)
	if inline_selection {
		__cgo__inline_selection = C.gboolean(1)
	}
	C.gtk_entry_completion_set_inline_selection(self.CPointer, __cgo__inline_selection)
	return
}

/*
Sets the match function for @completion to be @func. The match function
is used to determine if a row should or should not be in the completion
list.
*/
func (self *_TraitEntryCompletion) SetMatchFunc(func_ C.GtkEntryCompletionMatchFunc, func_data unsafe.Pointer, func_notify C.GDestroyNotify) {
	C.gtk_entry_completion_set_match_func(self.CPointer, func_, (C.gpointer)(func_data), func_notify)
	return
}

/*
Requires the length of the search key for @completion to be at least
@length. This is useful for long lists, where completing using a small
key takes a lot of time and will come up with meaningless results anyway
(ie, a too large dataset).
*/
func (self *_TraitEntryCompletion) SetMinimumKeyLength(length int) {
	C.gtk_entry_completion_set_minimum_key_length(self.CPointer, C.gint(length))
	return
}

/*
Sets the model for a #GtkEntryCompletion. If @completion already has
a model set, it will remove it before setting the new model.
If model is %NULL, then it will unset the model.
*/
func (self *_TraitEntryCompletion) SetModel(model *C.GtkTreeModel) {
	C.gtk_entry_completion_set_model(self.CPointer, model)
	return
}

/*
Sets whether the completions should be presented in a popup window.
*/
func (self *_TraitEntryCompletion) SetPopupCompletion(popup_completion bool) {
	__cgo__popup_completion := C.gboolean(0)
	if popup_completion {
		__cgo__popup_completion = C.gboolean(1)
	}
	C.gtk_entry_completion_set_popup_completion(self.CPointer, __cgo__popup_completion)
	return
}

/*
Sets whether the completion popup window will be resized to be the same
width as the entry.
*/
func (self *_TraitEntryCompletion) SetPopupSetWidth(popup_set_width bool) {
	__cgo__popup_set_width := C.gboolean(0)
	if popup_set_width {
		__cgo__popup_set_width = C.gboolean(1)
	}
	C.gtk_entry_completion_set_popup_set_width(self.CPointer, __cgo__popup_set_width)
	return
}

/*
Sets whether the completion popup window will appear even if there is
only a single match. You may want to set this to %FALSE if you
are using [inline completion][GtkEntryCompletion--inline-completion].
*/
func (self *_TraitEntryCompletion) SetPopupSingleMatch(popup_single_match bool) {
	__cgo__popup_single_match := C.gboolean(0)
	if popup_single_match {
		__cgo__popup_single_match = C.gboolean(1)
	}
	C.gtk_entry_completion_set_popup_single_match(self.CPointer, __cgo__popup_single_match)
	return
}

/*
Convenience function for setting up the most used case of this code: a
completion list with just strings. This function will set up @completion
to have a list displaying all (and just) strings in the completion list,
and to get those strings from @column in the model of @completion.

This functions creates and adds a #GtkCellRendererText for the selected
column. If you need to set the text column, but don't want the cell
renderer, use g_object_set() to set the #GtkEntryCompletion:text-column
property directly.
*/
func (self *_TraitEntryCompletion) SetTextColumn(column int) {
	C.gtk_entry_completion_set_text_column(self.CPointer, C.gint(column))
	return
}

type _TraitEventBox struct{ CPointer *C.GtkEventBox }

/*
Returns whether the event box window is above or below the
windows of its child. See gtk_event_box_set_above_child()
for details.
*/
func (self *_TraitEventBox) GetAboveChild() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_event_box_get_above_child(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the event box has a visible window.
See gtk_event_box_set_visible_window() for details.
*/
func (self *_TraitEventBox) GetVisibleWindow() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_event_box_get_visible_window(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Set whether the event box window is positioned above the windows
of its child, as opposed to below it. If the window is above, all
events inside the event box will go to the event box. If the window
is below, events in windows of child widgets will first got to that
widget, and then to its parents.

The default is to keep the window below the child.
*/
func (self *_TraitEventBox) SetAboveChild(above_child bool) {
	__cgo__above_child := C.gboolean(0)
	if above_child {
		__cgo__above_child = C.gboolean(1)
	}
	C.gtk_event_box_set_above_child(self.CPointer, __cgo__above_child)
	return
}

/*
Set whether the event box uses a visible or invisible child
window. The default is to use visible windows.

In an invisible window event box, the window that the
event box creates is a %GDK_INPUT_ONLY window, which
means that it is invisible and only serves to receive
events.

A visible window event box creates a visible (%GDK_INPUT_OUTPUT)
window that acts as the parent window for all the widgets
contained in the event box.

You should generally make your event box invisible if
you just want to trap events. Creating a visible window
may cause artifacts that are visible to the user, especially
if the user is using a theme with gradients or pixmaps.

The main reason to create a non input-only event box is if
you want to set the background to a different color or
draw on it.

There is one unexpected issue for an invisible event box that has its
window below the child. (See gtk_event_box_set_above_child().)
Since the input-only window is not an ancestor window of any windows
that descendent widgets of the event box create, events on these
windows aren’t propagated up by the windowing system, but only by GTK+.
The practical effect of this is if an event isn’t in the event
mask for the descendant window (see gtk_widget_add_events()),
it won’t be received by the event box.

This problem doesn’t occur for visible event boxes, because in
that case, the event box window is actually the ancestor of the
descendant windows, not just at the same place on the screen.
*/
func (self *_TraitEventBox) SetVisibleWindow(visible_window bool) {
	__cgo__visible_window := C.gboolean(0)
	if visible_window {
		__cgo__visible_window = C.gboolean(1)
	}
	C.gtk_event_box_set_visible_window(self.CPointer, __cgo__visible_window)
	return
}

type _TraitExpander struct{ CPointer *C.GtkExpander }

/*
Queries a #GtkExpander and returns its current state. Returns %TRUE
if the child widget is revealed.

See gtk_expander_set_expanded().
*/
func (self *_TraitExpander) GetExpanded() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_expander_get_expanded(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Fetches the text from a label widget including any embedded
underlines indicating mnemonics and Pango markup, as set by
gtk_expander_set_label(). If the label text has not been set the
return value will be %NULL. This will be the case if you create an
empty button with gtk_button_new() to use as a container.

Note that this function behaved differently in versions prior to
2.14 and used to return the label text stripped of embedded
underlines indicating mnemonics and Pango markup. This problem can
be avoided by fetching the label text directly from the label
widget.
*/
func (self *_TraitExpander) GetLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_expander_get_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns whether the label widget will fill all available
horizontal space allocated to @expander.
*/
func (self *_TraitExpander) GetLabelFill() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_expander_get_label_fill(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the label widget for the frame. See
gtk_expander_set_label_widget().
*/
func (self *_TraitExpander) GetLabelWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_expander_get_label_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether the expander will resize the toplevel widget
containing the expander upon resizing and collpasing.
*/
func (self *_TraitExpander) GetResizeToplevel() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_expander_get_resize_toplevel(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_expander_set_spacing().
*/
func (self *_TraitExpander) GetSpacing() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_expander_get_spacing(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns whether the label’s text is interpreted as marked up with
the [Pango text markup language][PangoMarkupFormat].
See gtk_expander_set_use_markup().
*/
func (self *_TraitExpander) GetUseMarkup() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_expander_get_use_markup(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether an embedded underline in the expander label
indicates a mnemonic. See gtk_expander_set_use_underline().
*/
func (self *_TraitExpander) GetUseUnderline() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_expander_get_use_underline(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the state of the expander. Set to %TRUE, if you want
the child widget to be revealed, and %FALSE if you want the
child widget to be hidden.
*/
func (self *_TraitExpander) SetExpanded(expanded bool) {
	__cgo__expanded := C.gboolean(0)
	if expanded {
		__cgo__expanded = C.gboolean(1)
	}
	C.gtk_expander_set_expanded(self.CPointer, __cgo__expanded)
	return
}

/*
Sets the text of the label of the expander to @label.

This will also clear any previously set labels.
*/
func (self *_TraitExpander) SetLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_expander_set_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Sets whether the label widget should fill all available
horizontal space allocated to @expander.
*/
func (self *_TraitExpander) SetLabelFill(label_fill bool) {
	__cgo__label_fill := C.gboolean(0)
	if label_fill {
		__cgo__label_fill = C.gboolean(1)
	}
	C.gtk_expander_set_label_fill(self.CPointer, __cgo__label_fill)
	return
}

/*
Set the label widget for the expander. This is the widget
that will appear embedded alongside the expander arrow.
*/
func (self *_TraitExpander) SetLabelWidget(label_widget *Widget) {
	C.gtk_expander_set_label_widget(self.CPointer, (*C.GtkWidget)(label_widget.CPointer))
	return
}

/*
Sets whether the expander will resize the toplevel widget
containing the expander upon resizing and collpasing.
*/
func (self *_TraitExpander) SetResizeToplevel(resize_toplevel bool) {
	__cgo__resize_toplevel := C.gboolean(0)
	if resize_toplevel {
		__cgo__resize_toplevel = C.gboolean(1)
	}
	C.gtk_expander_set_resize_toplevel(self.CPointer, __cgo__resize_toplevel)
	return
}

/*
Sets the spacing field of @expander, which is the number of
pixels to place between expander and the child.
*/
func (self *_TraitExpander) SetSpacing(spacing int) {
	C.gtk_expander_set_spacing(self.CPointer, C.gint(spacing))
	return
}

/*
Sets whether the text of the label contains markup in
[Pango’s text markup language][PangoMarkupFormat].
See gtk_label_set_markup().
*/
func (self *_TraitExpander) SetUseMarkup(use_markup bool) {
	__cgo__use_markup := C.gboolean(0)
	if use_markup {
		__cgo__use_markup = C.gboolean(1)
	}
	C.gtk_expander_set_use_markup(self.CPointer, __cgo__use_markup)
	return
}

/*
If true, an underline in the text of the expander label indicates
the next character should be used for the mnemonic accelerator key.
*/
func (self *_TraitExpander) SetUseUnderline(use_underline bool) {
	__cgo__use_underline := C.gboolean(0)
	if use_underline {
		__cgo__use_underline = C.gboolean(1)
	}
	C.gtk_expander_set_use_underline(self.CPointer, __cgo__use_underline)
	return
}

type _TraitFileChooserButton struct{ CPointer *C.GtkFileChooserButton }

/*
Returns whether the button grabs focus when it is clicked with the mouse.
See gtk_file_chooser_button_set_focus_on_click().
*/
func (self *_TraitFileChooserButton) GetFocusOnClick() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_file_chooser_button_get_focus_on_click(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the title of the browse dialog used by @button. The returned value
should not be modified or freed.
*/
func (self *_TraitFileChooserButton) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_file_chooser_button_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the width in characters of the @button widget’s entry and/or label.
*/
func (self *_TraitFileChooserButton) GetWidthChars() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_file_chooser_button_get_width_chars(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Sets whether the button will grab focus when it is clicked with the mouse.
Making mouse clicks not grab focus is useful in places like toolbars where
you don’t want the keyboard focus removed from the main area of the
application.
*/
func (self *_TraitFileChooserButton) SetFocusOnClick(focus_on_click bool) {
	__cgo__focus_on_click := C.gboolean(0)
	if focus_on_click {
		__cgo__focus_on_click = C.gboolean(1)
	}
	C.gtk_file_chooser_button_set_focus_on_click(self.CPointer, __cgo__focus_on_click)
	return
}

/*
Modifies the @title of the browse dialog used by @button.
*/
func (self *_TraitFileChooserButton) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_file_chooser_button_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
Sets the width (in characters) that @button will use to @n_chars.
*/
func (self *_TraitFileChooserButton) SetWidthChars(n_chars int) {
	C.gtk_file_chooser_button_set_width_chars(self.CPointer, C.gint(n_chars))
	return
}

type _TraitFileChooserDialog struct{ CPointer *C.GtkFileChooserDialog }

type _TraitFileChooserWidget struct{ CPointer *C.GtkFileChooserWidget }

type _TraitFileFilter struct{ CPointer *C.GtkFileFilter }

/*
Adds rule to a filter that allows files based on a custom callback
function. The bitfield @needed which is passed in provides information
about what sorts of information that the filter function needs;
this allows GTK+ to avoid retrieving expensive information when
it isn’t needed by the filter.
*/
func (self *_TraitFileFilter) AddCustom(needed C.GtkFileFilterFlags, func_ C.GtkFileFilterFunc, data unsafe.Pointer, notify C.GDestroyNotify) {
	C.gtk_file_filter_add_custom(self.CPointer, needed, func_, (C.gpointer)(data), notify)
	return
}

/*
Adds a rule allowing a given mime type to @filter.
*/
func (self *_TraitFileFilter) AddMimeType(mime_type string) {
	__cgo__mime_type := (*C.gchar)(unsafe.Pointer(C.CString(mime_type)))
	C.gtk_file_filter_add_mime_type(self.CPointer, __cgo__mime_type)
	C.free(unsafe.Pointer(__cgo__mime_type))
	return
}

/*
Adds a rule allowing a shell style glob to a filter.
*/
func (self *_TraitFileFilter) AddPattern(pattern string) {
	__cgo__pattern := (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	C.gtk_file_filter_add_pattern(self.CPointer, __cgo__pattern)
	C.free(unsafe.Pointer(__cgo__pattern))
	return
}

/*
Adds a rule allowing image files in the formats supported
by GdkPixbuf.
*/
func (self *_TraitFileFilter) AddPixbufFormats() {
	C.gtk_file_filter_add_pixbuf_formats(self.CPointer)
	return
}

/*
Tests whether a file should be displayed according to @filter.
The #GtkFileFilterInfo @filter_info should include
the fields returned from gtk_file_filter_get_needed().

This function will not typically be used by applications; it
is intended principally for use in the implementation of
#GtkFileChooser.
*/
func (self *_TraitFileFilter) Filter(filter_info *FileFilterInfo) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_file_filter_filter(self.CPointer, (*C.GtkFileFilterInfo)(unsafe.Pointer(filter_info)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the human-readable name for the filter. See gtk_file_filter_set_name().
*/
func (self *_TraitFileFilter) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_file_filter_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the fields that need to be filled in for the #GtkFileFilterInfo
passed to gtk_file_filter_filter()

This function will not typically be used by applications; it
is intended principally for use in the implementation of
#GtkFileChooser.
*/
func (self *_TraitFileFilter) GetNeeded() (return__ C.GtkFileFilterFlags) {
	return__ = C.gtk_file_filter_get_needed(self.CPointer)
	return
}

/*
Sets the human-readable name of the filter; this is the string
that will be displayed in the file selector user interface if
there is a selectable list of filters.
*/
func (self *_TraitFileFilter) SetName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_file_filter_set_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

type _TraitFixed struct{ CPointer *C.GtkFixed }

/*
Moves a child of a #GtkFixed container to the given position.
*/
func (self *_TraitFixed) Move(widget *Widget, x int, y int) {
	C.gtk_fixed_move(self.CPointer, (*C.GtkWidget)(widget.CPointer), C.gint(x), C.gint(y))
	return
}

/*
Adds a widget to a #GtkFixed container at the given position.
*/
func (self *_TraitFixed) Put(widget *Widget, x int, y int) {
	C.gtk_fixed_put(self.CPointer, (*C.GtkWidget)(widget.CPointer), C.gint(x), C.gint(y))
	return
}

type _TraitFlowBox struct{ CPointer *C.GtkFlowBox }

/*
Returns whether children activate on single clicks.
*/
func (self *_TraitFlowBox) GetActivateOnSingleClick() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_flow_box_get_activate_on_single_click(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the nth child in the @box.
*/
func (self *_TraitFlowBox) GetChildAtIndex(idx int) (return__ *FlowBoxChild) {
	var __cgo__return__ *C.GtkFlowBoxChild
	__cgo__return__ = C.gtk_flow_box_get_child_at_index(self.CPointer, C.gint(idx))
	return__ = NewFlowBoxChildFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the horizontal spacing.
*/
func (self *_TraitFlowBox) GetColumnSpacing() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_flow_box_get_column_spacing(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Returns whether the box is homogeneous (all children are the
same size). See gtk_box_set_homogeneous().
*/
func (self *_TraitFlowBox) GetHomogeneous() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_flow_box_get_homogeneous(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the maximum number of children per line.
*/
func (self *_TraitFlowBox) GetMaxChildrenPerLine() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_flow_box_get_max_children_per_line(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the minimum number of children per line.
*/
func (self *_TraitFlowBox) GetMinChildrenPerLine() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_flow_box_get_min_children_per_line(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the vertical spacing.
*/
func (self *_TraitFlowBox) GetRowSpacing() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_flow_box_get_row_spacing(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Creates a list of all selected children.
*/
func (self *_TraitFlowBox) GetSelectedChildren() (return__ *C.GList) {
	return__ = C.gtk_flow_box_get_selected_children(self.CPointer)
	return
}

/*
Gets the selection mode of @box.
*/
func (self *_TraitFlowBox) GetSelectionMode() (return__ C.GtkSelectionMode) {
	return__ = C.gtk_flow_box_get_selection_mode(self.CPointer)
	return
}

/*
Inserts the @widget into @box at @position.

If a sort function is set, the widget will actually be inserted
at the calculated position and this function has the same effect
as gtk_container_add().

If @position is -1, or larger than the total number of children
in the @box, then the @widget will be appended to the end.
*/
func (self *_TraitFlowBox) Insert(widget *Widget, position int) {
	C.gtk_flow_box_insert(self.CPointer, (*C.GtkWidget)(widget.CPointer), C.gint(position))
	return
}

/*
Updates the filtering for all children.

Call this function when the result of the filter
function on the @box is changed due ot an external
factor. For instance, this would be used if the
filter function just looked for a specific search
term, and the entry with the string has changed.
*/
func (self *_TraitFlowBox) InvalidateFilter() {
	C.gtk_flow_box_invalidate_filter(self.CPointer)
	return
}

/*
Updates the sorting for all children.

Call this when the result of the sort function on
@box is changed due to an external factor.
*/
func (self *_TraitFlowBox) InvalidateSort() {
	C.gtk_flow_box_invalidate_sort(self.CPointer)
	return
}

/*
Select all children of @box, if the selection
mode allows it.
*/
func (self *_TraitFlowBox) SelectAll() {
	C.gtk_flow_box_select_all(self.CPointer)
	return
}

/*
Selects a single child of @box, if the selection
mode allows it.
*/
func (self *_TraitFlowBox) SelectChild(child *FlowBoxChild) {
	C.gtk_flow_box_select_child(self.CPointer, (*C.GtkFlowBoxChild)(child.CPointer))
	return
}

/*
Calls a function for each selected child.

Note that the selection cannot be modified from within
this function.
*/
func (self *_TraitFlowBox) SelectedForeach(func_ C.GtkFlowBoxForeachFunc, data unsafe.Pointer) {
	C.gtk_flow_box_selected_foreach(self.CPointer, func_, (C.gpointer)(data))
	return
}

/*
If @single is %TRUE, children will be activated when you click
on them, otherwise you need to double-click.
*/
func (self *_TraitFlowBox) SetActivateOnSingleClick(single bool) {
	__cgo__single := C.gboolean(0)
	if single {
		__cgo__single = C.gboolean(1)
	}
	C.gtk_flow_box_set_activate_on_single_click(self.CPointer, __cgo__single)
	return
}

/*
Sets the horizontal space to add between children.
See the #GtkFlowBox:column-spacing property.
*/
func (self *_TraitFlowBox) SetColumnSpacing(spacing uint) {
	C.gtk_flow_box_set_column_spacing(self.CPointer, C.guint(spacing))
	return
}

/*
By setting a filter function on the @box one can decide dynamically
which of the children to show. For instance, to implement a search
function that only shows the children matching the search terms.

The @filter_func will be called for each child after the call, and
it will continue to be called each time a child changes (via
gtk_flow_box_child_changed()) or when gtk_flow_box_invalidate_filter()
is called.
*/
func (self *_TraitFlowBox) SetFilterFunc(filter_func C.GtkFlowBoxFilterFunc, user_data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_flow_box_set_filter_func(self.CPointer, filter_func, (C.gpointer)(user_data), destroy)
	return
}

/*
Hooks up an adjustment to focus handling in @box.
The adjustment is also used for autoscrolling during
rubberband selection. See gtk_scrolled_window_get_hadjustment()
for a typical way of obtaining the adjustment, and
gtk_flow_box_set_vadjustment()for setting the vertical
adjustment.

The adjustments have to be in pixel units and in the same
coordinate system as the allocation for immediate children
of the box.
*/
func (self *_TraitFlowBox) SetHadjustment(adjustment *Adjustment) {
	C.gtk_flow_box_set_hadjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
Sets the #GtkFlowBox:homogeneous property of @box, controlling
whether or not all children of @box are given equal space
in the box.
*/
func (self *_TraitFlowBox) SetHomogeneous(homogeneous bool) {
	__cgo__homogeneous := C.gboolean(0)
	if homogeneous {
		__cgo__homogeneous = C.gboolean(1)
	}
	C.gtk_flow_box_set_homogeneous(self.CPointer, __cgo__homogeneous)
	return
}

/*
Sets the maximum number of children to request and
allocate space for in @box’s orientation.

Setting the maximum number of children per line
limits the overall natural size request to be no more
than @n_children children long in the given orientation.
*/
func (self *_TraitFlowBox) SetMaxChildrenPerLine(n_children uint) {
	C.gtk_flow_box_set_max_children_per_line(self.CPointer, C.guint(n_children))
	return
}

/*
Sets the minimum number of children to line up
in @box’s orientation before flowing.
*/
func (self *_TraitFlowBox) SetMinChildrenPerLine(n_children uint) {
	C.gtk_flow_box_set_min_children_per_line(self.CPointer, C.guint(n_children))
	return
}

/*
Sets the vertical space to add between children.
See the #GtkFlowBox:row-spacing property.
*/
func (self *_TraitFlowBox) SetRowSpacing(spacing uint) {
	C.gtk_flow_box_set_row_spacing(self.CPointer, C.guint(spacing))
	return
}

/*
Sets how selection works in @box.
See #GtkSelectionMode for details.
*/
func (self *_TraitFlowBox) SetSelectionMode(mode C.GtkSelectionMode) {
	C.gtk_flow_box_set_selection_mode(self.CPointer, mode)
	return
}

/*
By setting a sort function on the @box, one can dynamically
reorder the children of the box, based on the contents of
the children.

The @sort_func will be called for each child after the call,
and will continue to be called each time a child changes (via
gtk_flow_box_child_changed()) and when gtk_flow_box_invalidate_sort()
is called.
*/
func (self *_TraitFlowBox) SetSortFunc(sort_func C.GtkFlowBoxSortFunc, user_data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_flow_box_set_sort_func(self.CPointer, sort_func, (C.gpointer)(user_data), destroy)
	return
}

/*
Hooks up an adjustment to focus handling in @box.
The adjustment is also used for autoscrolling during
rubberband selection. See gtk_scrolled_window_get_vadjustment()
for a typical way of obtaining the adjustment, and
gtk_flow_box_set_hadjustment()for setting the horizontal
adjustment.

The adjustments have to be in pixel units and in the same
coordinate system as the allocation for immediate children
of the box.
*/
func (self *_TraitFlowBox) SetVadjustment(adjustment *Adjustment) {
	C.gtk_flow_box_set_vadjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
Unselect all children of @box, if the selection
mode allows it.
*/
func (self *_TraitFlowBox) UnselectAll() {
	C.gtk_flow_box_unselect_all(self.CPointer)
	return
}

/*
Unselects a single child of @box, if the selection
mode allows it.
*/
func (self *_TraitFlowBox) UnselectChild(child *FlowBoxChild) {
	C.gtk_flow_box_unselect_child(self.CPointer, (*C.GtkFlowBoxChild)(child.CPointer))
	return
}

type _TraitFlowBoxChild struct{ CPointer *C.GtkFlowBoxChild }

/*
Marks @child as changed, causing any state that depends on this
to be updated. This affects sorting and filtering.

Note that calls to this method must be in sync with the data
used for the sorting and filtering functions. For instance, if
the list is mirroring some external data set, and *two* children
changed in the external data set when you call
gtk_flow_box_child_changed() on the first child, the sort function
must only read the new data for the first of the two changed
children, otherwise the resorting of the children will be wrong.

This generally means that if you don’t fully control the data
model, you have to duplicate the data that affects the sorting
and filtering functions into the widgets themselves. Another
alternative is to call gtk_flow_box_invalidate_sort() on any
model change, but that is more expensive.
*/
func (self *_TraitFlowBoxChild) Changed() {
	C.gtk_flow_box_child_changed(self.CPointer)
	return
}

/*
Gets the current index of the @child in its #GtkFlowBox container.
*/
func (self *_TraitFlowBoxChild) GetIndex() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_flow_box_child_get_index(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns whether the @child is currently selected in its
#GtkFlowBox container.
*/
func (self *_TraitFlowBoxChild) IsSelected() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_flow_box_child_is_selected(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type _TraitFontButton struct{ CPointer *C.GtkFontButton }

/*
Retrieves the name of the currently selected font. This name includes
style and size information as well. If you want to render something
with the font, use this string with pango_font_description_from_string() .
If you’re interested in peeking certain values (family name,
style, size, weight) just query these properties from the
#PangoFontDescription object.
*/
func (self *_TraitFontButton) GetFontName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_font_button_get_font_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns whether the font size will be shown in the label.
*/
func (self *_TraitFontButton) GetShowSize() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_font_button_get_show_size(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the name of the font style will be shown in the label.
*/
func (self *_TraitFontButton) GetShowStyle() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_font_button_get_show_style(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the title of the font chooser dialog.
*/
func (self *_TraitFontButton) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_font_button_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns whether the selected font is used in the label.
*/
func (self *_TraitFontButton) GetUseFont() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_font_button_get_use_font(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the selected size is used in the label.
*/
func (self *_TraitFontButton) GetUseSize() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_font_button_get_use_size(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets or updates the currently-displayed font in font picker dialog.
*/
func (self *_TraitFontButton) SetFontName(fontname string) (return__ bool) {
	__cgo__fontname := (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_font_button_set_font_name(self.CPointer, __cgo__fontname)
	C.free(unsafe.Pointer(__cgo__fontname))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
If @show_size is %TRUE, the font size will be displayed along with the name of the selected font.
*/
func (self *_TraitFontButton) SetShowSize(show_size bool) {
	__cgo__show_size := C.gboolean(0)
	if show_size {
		__cgo__show_size = C.gboolean(1)
	}
	C.gtk_font_button_set_show_size(self.CPointer, __cgo__show_size)
	return
}

/*
If @show_style is %TRUE, the font style will be displayed along with name of the selected font.
*/
func (self *_TraitFontButton) SetShowStyle(show_style bool) {
	__cgo__show_style := C.gboolean(0)
	if show_style {
		__cgo__show_style = C.gboolean(1)
	}
	C.gtk_font_button_set_show_style(self.CPointer, __cgo__show_style)
	return
}

/*
Sets the title for the font chooser dialog.
*/
func (self *_TraitFontButton) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_font_button_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
If @use_font is %TRUE, the font name will be written using the selected font.
*/
func (self *_TraitFontButton) SetUseFont(use_font bool) {
	__cgo__use_font := C.gboolean(0)
	if use_font {
		__cgo__use_font = C.gboolean(1)
	}
	C.gtk_font_button_set_use_font(self.CPointer, __cgo__use_font)
	return
}

/*
If @use_size is %TRUE, the font name will be written using the selected size.
*/
func (self *_TraitFontButton) SetUseSize(use_size bool) {
	__cgo__use_size := C.gboolean(0)
	if use_size {
		__cgo__use_size = C.gboolean(1)
	}
	C.gtk_font_button_set_use_size(self.CPointer, __cgo__use_size)
	return
}

type _TraitFontChooserDialog struct{ CPointer *C.GtkFontChooserDialog }

type _TraitFontChooserWidget struct{ CPointer *C.GtkFontChooserWidget }

type _TraitFontSelection struct{ CPointer *C.GtkFontSelection }

// gtk_font_selection_get_face is not generated due to deprecation attr

// gtk_font_selection_get_face_list is not generated due to deprecation attr

// gtk_font_selection_get_family is not generated due to deprecation attr

// gtk_font_selection_get_family_list is not generated due to deprecation attr

// gtk_font_selection_get_font_name is not generated due to deprecation attr

// gtk_font_selection_get_preview_entry is not generated due to deprecation attr

// gtk_font_selection_get_preview_text is not generated due to deprecation attr

// gtk_font_selection_get_size is not generated due to deprecation attr

// gtk_font_selection_get_size_entry is not generated due to deprecation attr

// gtk_font_selection_get_size_list is not generated due to deprecation attr

// gtk_font_selection_set_font_name is not generated due to deprecation attr

// gtk_font_selection_set_preview_text is not generated due to deprecation attr

type _TraitFontSelectionDialog struct{ CPointer *C.GtkFontSelectionDialog }

// gtk_font_selection_dialog_get_cancel_button is not generated due to deprecation attr

// gtk_font_selection_dialog_get_font_name is not generated due to deprecation attr

// gtk_font_selection_dialog_get_font_selection is not generated due to deprecation attr

// gtk_font_selection_dialog_get_ok_button is not generated due to deprecation attr

// gtk_font_selection_dialog_get_preview_text is not generated due to deprecation attr

// gtk_font_selection_dialog_set_font_name is not generated due to deprecation attr

// gtk_font_selection_dialog_set_preview_text is not generated due to deprecation attr

type _TraitFrame struct{ CPointer *C.GtkFrame }

/*
If the frame’s label widget is a #GtkLabel, returns the
text in the label widget. (The frame will have a #GtkLabel
for the label widget if a non-%NULL argument was passed
to gtk_frame_new().)
*/
func (self *_TraitFrame) GetLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_frame_get_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the X and Y alignment of the frame’s label. See
gtk_frame_set_label_align().
*/
func (self *_TraitFrame) GetLabelAlign() (xalign float32, yalign float32) {
	var __cgo__xalign C.gfloat
	var __cgo__yalign C.gfloat
	C.gtk_frame_get_label_align(self.CPointer, &__cgo__xalign, &__cgo__yalign)
	xalign = float32(__cgo__xalign)
	yalign = float32(__cgo__yalign)
	return
}

/*
Retrieves the label widget for the frame. See
gtk_frame_set_label_widget().
*/
func (self *_TraitFrame) GetLabelWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_frame_get_label_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the shadow type of the frame. See
gtk_frame_set_shadow_type().
*/
func (self *_TraitFrame) GetShadowType() (return__ C.GtkShadowType) {
	return__ = C.gtk_frame_get_shadow_type(self.CPointer)
	return
}

/*
Sets the text of the label. If @label is %NULL,
the current label is removed.
*/
func (self *_TraitFrame) SetLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_frame_set_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Sets the alignment of the frame widget’s label. The
default values for a newly created frame are 0.0 and 0.5.
*/
func (self *_TraitFrame) SetLabelAlign(xalign float32, yalign float32) {
	C.gtk_frame_set_label_align(self.CPointer, C.gfloat(xalign), C.gfloat(yalign))
	return
}

/*
Sets the label widget for the frame. This is the widget that
will appear embedded in the top edge of the frame as a
title.
*/
func (self *_TraitFrame) SetLabelWidget(label_widget *Widget) {
	C.gtk_frame_set_label_widget(self.CPointer, (*C.GtkWidget)(label_widget.CPointer))
	return
}

/*
Sets the shadow type for @frame.
*/
func (self *_TraitFrame) SetShadowType(type_ C.GtkShadowType) {
	C.gtk_frame_set_shadow_type(self.CPointer, type_)
	return
}

type _TraitGrid struct{ CPointer *C.GtkGrid }

/*
Adds a widget to the grid.

The position of @child is determined by @left and @top. The
number of “cells” that @child will occupy is determined by
@width and @height.
*/
func (self *_TraitGrid) Attach(child *Widget, left int, top int, width int, height int) {
	C.gtk_grid_attach(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(left), C.gint(top), C.gint(width), C.gint(height))
	return
}

/*
Adds a widget to the grid.

The widget is placed next to @sibling, on the side determined by
@side. When @sibling is %NULL, the widget is placed in row (for
left or right placement) or column 0 (for top or bottom placement),
at the end indicated by @side.

Attaching widgets labeled [1], [2], [3] with @sibling == %NULL and
@side == %GTK_POS_LEFT yields a layout of [3][2][1].
*/
func (self *_TraitGrid) AttachNextTo(child *Widget, sibling *Widget, side C.GtkPositionType, width int, height int) {
	C.gtk_grid_attach_next_to(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(sibling.CPointer), side, C.gint(width), C.gint(height))
	return
}

/*
Returns which row defines the global baseline of @grid.
*/
func (self *_TraitGrid) GetBaselineRow() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_grid_get_baseline_row(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the child of @grid whose area covers the grid
cell whose upper left corner is at @left, @top.
*/
func (self *_TraitGrid) GetChildAt(left int, top int) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_grid_get_child_at(self.CPointer, C.gint(left), C.gint(top))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether all columns of @grid have the same width.
*/
func (self *_TraitGrid) GetColumnHomogeneous() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_grid_get_column_homogeneous(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the amount of space between the columns of @grid.
*/
func (self *_TraitGrid) GetColumnSpacing() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_grid_get_column_spacing(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Returns the baseline position of @row as set
by gtk_grid_set_row_baseline_position() or the default value
%GTK_BASELINE_POSITION_CENTER.
*/
func (self *_TraitGrid) GetRowBaselinePosition(row int) (return__ C.GtkBaselinePosition) {
	return__ = C.gtk_grid_get_row_baseline_position(self.CPointer, C.gint(row))
	return
}

/*
Returns whether all rows of @grid have the same height.
*/
func (self *_TraitGrid) GetRowHomogeneous() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_grid_get_row_homogeneous(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the amount of space between the rows of @grid.
*/
func (self *_TraitGrid) GetRowSpacing() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_grid_get_row_spacing(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Inserts a column at the specified position.

Children which are attached at or to the right of this position
are moved one column to the right. Children which span across this
position are grown to span the new column.
*/
func (self *_TraitGrid) InsertColumn(position int) {
	C.gtk_grid_insert_column(self.CPointer, C.gint(position))
	return
}

/*
Inserts a row or column at the specified position.

The new row or column is placed next to @sibling, on the side
determined by @side. If @side is %GTK_POS_TOP or %GTK_POS_BOTTOM,
a row is inserted. If @side is %GTK_POS_LEFT of %GTK_POS_RIGHT,
a column is inserted.
*/
func (self *_TraitGrid) InsertNextTo(sibling *Widget, side C.GtkPositionType) {
	C.gtk_grid_insert_next_to(self.CPointer, (*C.GtkWidget)(sibling.CPointer), side)
	return
}

/*
Inserts a row at the specified position.

Children which are attached at or below this position
are moved one row down. Children which span across this
position are grown to span the new row.
*/
func (self *_TraitGrid) InsertRow(position int) {
	C.gtk_grid_insert_row(self.CPointer, C.gint(position))
	return
}

/*
Removes a column from the grid.

Children that are placed in this column are removed,
spanning children that overlap this column have their
width reduced by one, and children after the column
are moved to the left.
*/
func (self *_TraitGrid) RemoveColumn(position int) {
	C.gtk_grid_remove_column(self.CPointer, C.gint(position))
	return
}

/*
Removes a row from the grid.

Children that are placed in this row are removed,
spanning children that overlap this row have their
height reduced by one, and children below the row
are moved up.
*/
func (self *_TraitGrid) RemoveRow(position int) {
	C.gtk_grid_remove_row(self.CPointer, C.gint(position))
	return
}

/*
Sets which row defines the global baseline for the entire grid.
Each row in the grid can have its own local baseline, but only
one of those is global, meaning it will be the baseline in the
parent of the @grid.
*/
func (self *_TraitGrid) SetBaselineRow(row int) {
	C.gtk_grid_set_baseline_row(self.CPointer, C.gint(row))
	return
}

/*
Sets whether all columns of @grid will have the same width.
*/
func (self *_TraitGrid) SetColumnHomogeneous(homogeneous bool) {
	__cgo__homogeneous := C.gboolean(0)
	if homogeneous {
		__cgo__homogeneous = C.gboolean(1)
	}
	C.gtk_grid_set_column_homogeneous(self.CPointer, __cgo__homogeneous)
	return
}

/*
Sets the amount of space between columns of @grid.
*/
func (self *_TraitGrid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(self.CPointer, C.guint(spacing))
	return
}

/*
Sets how the baseline should be positioned on @row of the
grid, in case that row is assigned more space than is requested.
*/
func (self *_TraitGrid) SetRowBaselinePosition(row int, pos C.GtkBaselinePosition) {
	C.gtk_grid_set_row_baseline_position(self.CPointer, C.gint(row), pos)
	return
}

/*
Sets whether all rows of @grid will have the same height.
*/
func (self *_TraitGrid) SetRowHomogeneous(homogeneous bool) {
	__cgo__homogeneous := C.gboolean(0)
	if homogeneous {
		__cgo__homogeneous = C.gboolean(1)
	}
	C.gtk_grid_set_row_homogeneous(self.CPointer, __cgo__homogeneous)
	return
}

/*
Sets the amount of space between rows of @grid.
*/
func (self *_TraitGrid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(self.CPointer, C.guint(spacing))
	return
}

type _TraitHBox struct{ CPointer *C.GtkHBox }

type _TraitHButtonBox struct{ CPointer *C.GtkHButtonBox }

type _TraitHPaned struct{ CPointer *C.GtkHPaned }

type _TraitHSV struct{ CPointer *C.GtkHSV }

// gtk_hsv_get_color is not generated due to explicit ignore

// gtk_hsv_get_metrics is not generated due to explicit ignore

// gtk_hsv_is_adjusting is not generated due to explicit ignore

// gtk_hsv_set_color is not generated due to explicit ignore

// gtk_hsv_set_metrics is not generated due to explicit ignore

type _TraitHScale struct{ CPointer *C.GtkHScale }

type _TraitHScrollbar struct{ CPointer *C.GtkHScrollbar }

type _TraitHSeparator struct{ CPointer *C.GtkHSeparator }

type _TraitHandleBox struct{ CPointer *C.GtkHandleBox }

// gtk_handle_box_get_child_detached is not generated due to deprecation attr

// gtk_handle_box_get_handle_position is not generated due to deprecation attr

// gtk_handle_box_get_shadow_type is not generated due to deprecation attr

// gtk_handle_box_get_snap_edge is not generated due to deprecation attr

// gtk_handle_box_set_handle_position is not generated due to deprecation attr

// gtk_handle_box_set_shadow_type is not generated due to deprecation attr

// gtk_handle_box_set_snap_edge is not generated due to deprecation attr

type _TraitHeaderBar struct{ CPointer *C.GtkHeaderBar }

/*
Retrieves the custom title widget of the header. See
gtk_header_bar_set_custom_title().
*/
func (self *_TraitHeaderBar) GetCustomTitle() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_header_bar_get_custom_title(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the decoration layout set with
gtk_header_bar_set_decoration_layout().
*/
func (self *_TraitHeaderBar) GetDecorationLayout() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_header_bar_get_decoration_layout(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves whether the header bar reserves space for
a subtitle, regardless if one is currently set or not.
*/
func (self *_TraitHeaderBar) GetHasSubtitle() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_header_bar_get_has_subtitle(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether this header bar shows the standard window
decorations.
*/
func (self *_TraitHeaderBar) GetShowCloseButton() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_header_bar_get_show_close_button(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the subtitle of the header. See gtk_header_bar_set_subtitle().
*/
func (self *_TraitHeaderBar) GetSubtitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_header_bar_get_subtitle(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the title of the header. See gtk_header_bar_set_title().
*/
func (self *_TraitHeaderBar) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_header_bar_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Adds @child to @box, packed with reference to the
end of the @box.
*/
func (self *_TraitHeaderBar) PackEnd(child *Widget) {
	C.gtk_header_bar_pack_end(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Adds @child to @box, packed with reference to the
start of the @box.
*/
func (self *_TraitHeaderBar) PackStart(child *Widget) {
	C.gtk_header_bar_pack_start(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Sets a custom title for the #GtkHeaderBar.

The title should help a user identify the current view. This
supersedes any title set by gtk_header_bar_set_title() or
gtk_header_bar_set_subtitle(). To achieve the same style as
the builtin title and subtitle, use the “title” and “subtitle”
style classes.

You should set the custom title to %NULL, for the header title
label to be visible again.
*/
func (self *_TraitHeaderBar) SetCustomTitle(title_widget *Widget) {
	C.gtk_header_bar_set_custom_title(self.CPointer, (*C.GtkWidget)(title_widget.CPointer))
	return
}

/*
Sets the decoration layout for this header bar, overriding
the #GtkSettings:gtk-decoration-layout setting.

There can be valid reasons for overriding the setting, such
as a header bar design that does not allow for buttons to take
room on the right, or only offers room for a single close button.
Split header bars are another example for overriding the
setting.

The format of the string is button names, separated by commas.
A colon separates the buttons that should appear on the left
from those on the right. Recognized button names are minimize,
maximize, close, icon (the window icon) and menu (a menu button
for the fallback app menu).

For example, “menu:minimize,maximize,close” specifies a menu
on the left, and minimize, maximize and close buttons on the right.
*/
func (self *_TraitHeaderBar) SetDecorationLayout(layout string) {
	__cgo__layout := (*C.gchar)(unsafe.Pointer(C.CString(layout)))
	C.gtk_header_bar_set_decoration_layout(self.CPointer, __cgo__layout)
	C.free(unsafe.Pointer(__cgo__layout))
	return
}

/*
Sets whether the header bar should reserve space
for a subtitle, even if none is currently set.
*/
func (self *_TraitHeaderBar) SetHasSubtitle(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_header_bar_set_has_subtitle(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether this header bar shows the standard window decorations,
including close, maximize, and minimize.
*/
func (self *_TraitHeaderBar) SetShowCloseButton(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_header_bar_set_show_close_button(self.CPointer, __cgo__setting)
	return
}

/*
Sets the subtitle of the #GtkHeaderBar. The title should give a user
an additional detail to help him identify the current view.

Note that GtkHeaderBar by default reserves room for the subtitle,
even if none is currently set. If this is not desired, set the
#GtkHeaderBar:has-subtitle property to %FALSE.
*/
func (self *_TraitHeaderBar) SetSubtitle(subtitle string) {
	__cgo__subtitle := (*C.gchar)(unsafe.Pointer(C.CString(subtitle)))
	C.gtk_header_bar_set_subtitle(self.CPointer, __cgo__subtitle)
	C.free(unsafe.Pointer(__cgo__subtitle))
	return
}

/*
Sets the title of the #GtkHeaderBar. The title should help a user
identify the current view. A good title should not include the
application name.
*/
func (self *_TraitHeaderBar) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_header_bar_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

type _TraitIMContext struct{ CPointer *C.GtkIMContext }

/*
Asks the widget that the input context is attached to to delete
characters around the cursor position by emitting the
GtkIMContext::delete_surrounding signal. Note that @offset and @n_chars
are in characters not in bytes which differs from the usage other
places in #GtkIMContext.

In order to use this function, you should first call
gtk_im_context_get_surrounding() to get the current context, and
call this function immediately afterwards to make sure that you
know what you are deleting. You should also account for the fact
that even if the signal was handled, the input context might not
have deleted all the characters that were requested to be deleted.

This function is used by an input method that wants to make
subsitutions in the existing text in response to new input. It is
not useful for applications.
*/
func (self *_TraitIMContext) DeleteSurrounding(offset int, n_chars int) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_im_context_delete_surrounding(self.CPointer, C.gint(offset), C.gint(n_chars))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Allow an input method to internally handle key press and release
events. If this function returns %TRUE, then no further processing
should be done for this key event.
*/
func (self *_TraitIMContext) FilterKeypress(event *C.GdkEventKey) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_im_context_filter_keypress(self.CPointer, event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Notify the input method that the widget to which this
input context corresponds has gained focus. The input method
may, for example, change the displayed feedback to reflect
this change.
*/
func (self *_TraitIMContext) FocusIn() {
	C.gtk_im_context_focus_in(self.CPointer)
	return
}

/*
Notify the input method that the widget to which this
input context corresponds has lost focus. The input method
may, for example, change the displayed feedback or reset the contexts
state to reflect this change.
*/
func (self *_TraitIMContext) FocusOut() {
	C.gtk_im_context_focus_out(self.CPointer)
	return
}

/*
Retrieve the current preedit string for the input context,
and a list of attributes to apply to the string.
This string should be displayed inserted at the insertion
point.
*/
func (self *_TraitIMContext) GetPreeditString() (str string, attrs *C.PangoAttrList, cursor_pos int) {
	var __cgo__str *C.gchar
	var __cgo__cursor_pos C.gint
	C.gtk_im_context_get_preedit_string(self.CPointer, &__cgo__str, &attrs, &__cgo__cursor_pos)
	str = C.GoString((*C.char)(unsafe.Pointer(__cgo__str)))
	cursor_pos = int(__cgo__cursor_pos)
	return
}

/*
Retrieves context around the insertion point. Input methods
typically want context in order to constrain input text based on
existing text; this is important for languages such as Thai where
only some sequences of characters are allowed.

This function is implemented by emitting the
GtkIMContext::retrieve_surrounding signal on the input method; in
response to this signal, a widget should provide as much context as
is available, up to an entire paragraph, by calling
gtk_im_context_set_surrounding(). Note that there is no obligation
for a widget to respond to the ::retrieve_surrounding signal, so input
methods must be prepared to function without context.
*/
func (self *_TraitIMContext) GetSurrounding() (text string, cursor_index int, return__ bool) {
	var __cgo__text *C.gchar
	var __cgo__cursor_index C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_im_context_get_surrounding(self.CPointer, &__cgo__text, &__cgo__cursor_index)
	text = C.GoString((*C.char)(unsafe.Pointer(__cgo__text)))
	cursor_index = int(__cgo__cursor_index)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Notify the input method that a change such as a change in cursor
position has been made. This will typically cause the input
method to clear the preedit state.
*/
func (self *_TraitIMContext) Reset() {
	C.gtk_im_context_reset(self.CPointer)
	return
}

/*
Set the client window for the input context; this is the
#GdkWindow in which the input appears. This window is
used in order to correctly position status windows, and may
also be used for purposes internal to the input method.
*/
func (self *_TraitIMContext) SetClientWindow(window *C.GdkWindow) {
	C.gtk_im_context_set_client_window(self.CPointer, window)
	return
}

/*
Notify the input method that a change in cursor
position has been made. The location is relative to the client
window.
*/
func (self *_TraitIMContext) SetCursorLocation(area *C.GdkRectangle) {
	C.gtk_im_context_set_cursor_location(self.CPointer, area)
	return
}

/*
Sets surrounding context around the insertion point and preedit
string. This function is expected to be called in response to the
GtkIMContext::retrieve_surrounding signal, and will likely have no
effect if called at other times.
*/
func (self *_TraitIMContext) SetSurrounding(text string, len_ int, cursor_index int) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_im_context_set_surrounding(self.CPointer, __cgo__text, C.gint(len_), C.gint(cursor_index))
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Sets whether the IM context should use the preedit string
to display feedback. If @use_preedit is FALSE (default
is TRUE), then the IM context may use some other method to display
feedback, such as displaying it in a child of the root window.
*/
func (self *_TraitIMContext) SetUsePreedit(use_preedit bool) {
	__cgo__use_preedit := C.gboolean(0)
	if use_preedit {
		__cgo__use_preedit = C.gboolean(1)
	}
	C.gtk_im_context_set_use_preedit(self.CPointer, __cgo__use_preedit)
	return
}

type _TraitIMContextSimple struct{ CPointer *C.GtkIMContextSimple }

/*
Adds an additional table to search to the input context.
Each row of the table consists of @max_seq_len key symbols
followed by two #guint16 interpreted as the high and low
words of a #gunicode value. Tables are searched starting
from the last added.

The table must be sorted in dictionary order on the
numeric value of the key symbol fields. (Values beyond
the length of the sequence should be zero.)
*/
func (self *_TraitIMContextSimple) AddTable(data *C.guint16, max_seq_len int, n_seqs int) {
	C.gtk_im_context_simple_add_table(self.CPointer, data, C.gint(max_seq_len), C.gint(n_seqs))
	return
}

type _TraitIMMulticontext struct{ CPointer *C.GtkIMMulticontext }

// gtk_im_multicontext_append_menuitems is not generated due to deprecation attr

/*
Gets the id of the currently active slave of the @context.
*/
func (self *_TraitIMMulticontext) GetContextId() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.gtk_im_multicontext_get_context_id(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Sets the context id for @context.

This causes the currently active slave of @context to be
replaced by the slave corresponding to the new context id.
*/
func (self *_TraitIMMulticontext) SetContextId(context_id string) {
	__cgo__context_id := C.CString(context_id)
	C.gtk_im_multicontext_set_context_id(self.CPointer, __cgo__context_id)
	C.free(unsafe.Pointer(__cgo__context_id))
	return
}

type _TraitIconFactory struct{ CPointer *C.GtkIconFactory }

// gtk_icon_factory_add is not generated due to deprecation attr

// gtk_icon_factory_add_default is not generated due to deprecation attr

// gtk_icon_factory_lookup is not generated due to deprecation attr

// gtk_icon_factory_remove_default is not generated due to deprecation attr

type _TraitIconInfo struct{ CPointer *C.GtkIconInfo }

// gtk_icon_info_copy is not generated due to deprecation attr

// gtk_icon_info_free is not generated due to deprecation attr

/*
Fetches the set of attach points for an icon. An attach point
is a location in the icon that can be used as anchor points for attaching
emblems or overlays to the icon.
*/
func (self *_TraitIconInfo) GetAttachPoints() (points *C.GdkPoint, n_points int, return__ bool) {
	var __cgo__n_points C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_info_get_attach_points(self.CPointer, &points, &__cgo__n_points)
	n_points = int(__cgo__n_points)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the base scale for the icon. The base scale is a scale for the
icon that was specified by the icon theme creator. For instance an
icon drawn for a high-dpi screen with window-scale 2 for a base
size of 32 will be 64 pixels tall and have a base_scale of 2.
*/
func (self *_TraitIconInfo) GetBaseScale() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_info_get_base_scale(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the base size for the icon. The base size
is a size for the icon that was specified by
the icon theme creator. This may be different
than the actual size of image; an example of
this is small emblem icons that can be attached
to a larger icon. These icons will be given
the same base size as the larger icons to which
they are attached.

Note that for scaled icons the base size does
not include the base scale.
*/
func (self *_TraitIconInfo) GetBaseSize() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_info_get_base_size(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the built-in image for this icon, if any. To allow
GTK+ to use built in icon images, you must pass the
%GTK_ICON_LOOKUP_USE_BUILTIN to
gtk_icon_theme_lookup_icon().
*/
func (self *_TraitIconInfo) GetBuiltinPixbuf() (return__ *C.GdkPixbuf) {
	return__ = C.gtk_icon_info_get_builtin_pixbuf(self.CPointer)
	return
}

/*
Gets the display name for an icon. A display name is a
string to be used in place of the icon name in a user
visible context like a list of icons.
*/
func (self *_TraitIconInfo) GetDisplayName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_icon_info_get_display_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the coordinates of a rectangle within the icon
that can be used for display of information such
as a preview of the contents of a text file.
See gtk_icon_info_set_raw_coordinates() for further
information about the coordinate system.
*/
func (self *_TraitIconInfo) GetEmbeddedRect() (rectangle C.GdkRectangle, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_info_get_embedded_rect(self.CPointer, &rectangle)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the filename for the icon. If the
%GTK_ICON_LOOKUP_USE_BUILTIN flag was passed
to gtk_icon_theme_lookup_icon(), there may be
no filename if a builtin icon is returned; in this
case, you should use gtk_icon_info_get_builtin_pixbuf().
*/
func (self *_TraitIconInfo) GetFilename() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_icon_info_get_filename(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Checks if the icon is symbolic or not. This currently uses only
the file name and not the file contents for determining this.
This behaviour may change in the future.
*/
func (self *_TraitIconInfo) IsSymbolic() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_info_is_symbolic(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Renders an icon previously looked up in an icon theme using
gtk_icon_theme_lookup_icon(); the size will be based on the size
passed to gtk_icon_theme_lookup_icon(). Note that the resulting
pixbuf may not be exactly this size; an icon theme may have icons
that differ slightly from their nominal sizes, and in addition GTK+
will avoid scaling icons that it considers sufficiently close to the
requested size or for which the source image would have to be scaled
up too far. (This maintains sharpness.). This behaviour can be changed
by passing the %GTK_ICON_LOOKUP_FORCE_SIZE flag when obtaining
the #GtkIconInfo. If this flag has been specified, the pixbuf
returned by this function will be scaled to the exact size.
*/
func (self *_TraitIconInfo) LoadIcon() (return__ *C.GdkPixbuf, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_info_load_icon(self.CPointer, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously load, render and scale an icon previously looked up
from the icon theme using gtk_icon_theme_lookup_icon().

For more details, see gtk_icon_info_load_icon() which is the synchronous
version of this call.
*/
func (self *_TraitIconInfo) LoadIconAsync(cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.gtk_icon_info_load_icon_async(self.CPointer, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an async icon load, see gtk_icon_info_load_icon_async().
*/
func (self *_TraitIconInfo) LoadIconFinish(res *C.GAsyncResult) (return__ *C.GdkPixbuf, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_info_load_icon_finish(self.CPointer, res, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Renders an icon previously looked up in an icon theme using
gtk_icon_theme_lookup_icon(); the size will be based on the size
passed to gtk_icon_theme_lookup_icon(). Note that the resulting
surface may not be exactly this size; an icon theme may have icons
that differ slightly from their nominal sizes, and in addition GTK+
will avoid scaling icons that it considers sufficiently close to the
requested size or for which the source image would have to be scaled
up too far. (This maintains sharpness.). This behaviour can be changed
by passing the %GTK_ICON_LOOKUP_FORCE_SIZE flag when obtaining
the #GtkIconInfo. If this flag has been specified, the pixbuf
returned by this function will be scaled to the exact size.
*/
func (self *_TraitIconInfo) LoadSurface(for_window *C.GdkWindow) (return__ *C.cairo_surface_t, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_info_load_surface(self.CPointer, for_window, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Loads an icon, modifying it to match the system colours for the foreground,
success, warning and error colors provided. If the icon is not a symbolic
one, the function will return the result from gtk_icon_info_load_icon().

This allows loading symbolic icons that will match the system theme.

Unless you are implementing a widget, you will want to use
g_themed_icon_new_with_default_fallbacks() to load the icon.

As implementation details, the icon loaded needs to be of SVG type,
contain the “symbolic” term as the last component of the icon name,
and use the “fg”, “success”, “warning” and “error” CSS styles in the
SVG file itself.

See the [Symbolic Icons Specification](http://www.freedesktop.org/wiki/SymbolicIcons)
for more information about symbolic icons.
*/
func (self *_TraitIconInfo) LoadSymbolic(fg *C.GdkRGBA, success_color *C.GdkRGBA, warning_color *C.GdkRGBA, error_color *C.GdkRGBA) (was_symbolic bool, return__ *C.GdkPixbuf, __err__ error) {
	var __cgo__was_symbolic C.gboolean
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_info_load_symbolic(self.CPointer, fg, success_color, warning_color, error_color, &__cgo__was_symbolic, &__cgo_error__)
	was_symbolic = __cgo__was_symbolic == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously load, render and scale a symbolic icon previously looked up
from the icon theme using gtk_icon_theme_lookup_icon().

For more details, see gtk_icon_info_load_symbolic() which is the synchronous
version of this call.
*/
func (self *_TraitIconInfo) LoadSymbolicAsync(fg *C.GdkRGBA, success_color *C.GdkRGBA, warning_color *C.GdkRGBA, error_color *C.GdkRGBA, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.gtk_icon_info_load_symbolic_async(self.CPointer, fg, success_color, warning_color, error_color, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an async icon load, see gtk_icon_info_load_symbolic_async().
*/
func (self *_TraitIconInfo) LoadSymbolicFinish(res *C.GAsyncResult) (was_symbolic bool, return__ *C.GdkPixbuf, __err__ error) {
	var __cgo__was_symbolic C.gboolean
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_info_load_symbolic_finish(self.CPointer, res, &__cgo__was_symbolic, &__cgo_error__)
	was_symbolic = __cgo__was_symbolic == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Loads an icon, modifying it to match the system colors for the foreground,
success, warning and error colors provided. If the icon is not a symbolic
one, the function will return the result from gtk_icon_info_load_icon().
This function uses the regular foreground color and the symbolic colors
with the names “success_color”, “warning_color” and “error_color” from
the context.

This allows loading symbolic icons that will match the system theme.

See gtk_icon_info_load_symbolic() for more details.
*/
func (self *_TraitIconInfo) LoadSymbolicForContext(context *StyleContext) (was_symbolic bool, return__ *C.GdkPixbuf, __err__ error) {
	var __cgo__was_symbolic C.gboolean
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_info_load_symbolic_for_context(self.CPointer, (*C.GtkStyleContext)(context.CPointer), &__cgo__was_symbolic, &__cgo_error__)
	was_symbolic = __cgo__was_symbolic == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously load, render and scale a symbolic icon previously looked up
from the icon theme using gtk_icon_theme_lookup_icon().

For more details, see gtk_icon_info_load_symbolic_for_context() which is the synchronous
version of this call.
*/
func (self *_TraitIconInfo) LoadSymbolicForContextAsync(context *StyleContext, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.gtk_icon_info_load_symbolic_for_context_async(self.CPointer, (*C.GtkStyleContext)(context.CPointer), cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an async icon load, see gtk_icon_info_load_symbolic_for_context_async().
*/
func (self *_TraitIconInfo) LoadSymbolicForContextFinish(res *C.GAsyncResult) (was_symbolic bool, return__ *C.GdkPixbuf, __err__ error) {
	var __cgo__was_symbolic C.gboolean
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_info_load_symbolic_for_context_finish(self.CPointer, res, &__cgo__was_symbolic, &__cgo_error__)
	was_symbolic = __cgo__was_symbolic == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

// gtk_icon_info_load_symbolic_for_style is not generated due to deprecation attr

/*
Sets whether the coordinates returned by gtk_icon_info_get_embedded_rect()
and gtk_icon_info_get_attach_points() should be returned in their
original form as specified in the icon theme, instead of scaled
appropriately for the pixbuf returned by gtk_icon_info_load_icon().

Raw coordinates are somewhat strange; they are specified to be with
respect to the unscaled pixmap for PNG and XPM icons, but for SVG
icons, they are in a 1000x1000 coordinate space that is scaled
to the final size of the icon.  You can determine if the icon is an SVG
icon by using gtk_icon_info_get_filename(), and seeing if it is non-%NULL
and ends in “.svg”.

This function is provided primarily to allow compatibility wrappers
for older API's, and is not expected to be useful for applications.
*/
func (self *_TraitIconInfo) SetRawCoordinates(raw_coordinates bool) {
	__cgo__raw_coordinates := C.gboolean(0)
	if raw_coordinates {
		__cgo__raw_coordinates = C.gboolean(1)
	}
	C.gtk_icon_info_set_raw_coordinates(self.CPointer, __cgo__raw_coordinates)
	return
}

type _TraitIconTheme struct{ CPointer *C.GtkIconTheme }

/*
Appends a directory to the search path.
See gtk_icon_theme_set_search_path().
*/
func (self *_TraitIconTheme) AppendSearchPath(path string) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	C.gtk_icon_theme_append_search_path(self.CPointer, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

// gtk_icon_theme_choose_icon is not generated due to explicit ignore

// gtk_icon_theme_choose_icon_for_scale is not generated due to explicit ignore

/*
Gets the name of an icon that is representative of the
current theme (for instance, to use when presenting
a list of themes to the user.)
*/
func (self *_TraitIconTheme) GetExampleIconName() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.gtk_icon_theme_get_example_icon_name(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Returns an array of integers describing the sizes at which
the icon is available without scaling. A size of -1 means
that the icon is available in a scalable format. The array
is zero-terminated.
*/
func (self *_TraitIconTheme) GetIconSizes(icon_name string) (return__ *C.gint) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	return__ = C.gtk_icon_theme_get_icon_sizes(self.CPointer, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

// gtk_icon_theme_get_search_path is not generated due to explicit ignore

/*
Checks whether an icon theme includes an icon
for a particular name.
*/
func (self *_TraitIconTheme) HasIcon(icon_name string) (return__ bool) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_theme_has_icon(self.CPointer, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the list of contexts available within the current
hierarchy of icon themes
*/
func (self *_TraitIconTheme) ListContexts() (return__ *C.GList) {
	return__ = C.gtk_icon_theme_list_contexts(self.CPointer)
	return
}

/*
Lists the icons in the current icon theme. Only a subset
of the icons can be listed by providing a context string.
The set of values for the context string is system dependent,
but will typically include such values as “Applications” and
“MimeTypes”.
*/
func (self *_TraitIconTheme) ListIcons(context string) (return__ *C.GList) {
	__cgo__context := (*C.gchar)(unsafe.Pointer(C.CString(context)))
	return__ = C.gtk_icon_theme_list_icons(self.CPointer, __cgo__context)
	C.free(unsafe.Pointer(__cgo__context))
	return
}

/*
Looks up an icon in an icon theme, scales it to the given size
and renders it into a pixbuf. This is a convenience function;
if more details about the icon are needed, use
gtk_icon_theme_lookup_icon() followed by gtk_icon_info_load_icon().

Note that you probably want to listen for icon theme changes and
update the icon. This is usually done by connecting to the
GtkWidget::style-set signal. If for some reason you do not want to
update the icon when the icon theme changes, you should consider
using gdk_pixbuf_copy() to make a private copy of the pixbuf
returned by this function. Otherwise GTK+ may need to keep the old
icon theme loaded, which would be a waste of memory.
*/
func (self *_TraitIconTheme) LoadIcon(icon_name string, size int, flags C.GtkIconLookupFlags) (return__ *C.GdkPixbuf, __err__ error) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_theme_load_icon(self.CPointer, __cgo__icon_name, C.gint(size), flags, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__icon_name))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Looks up an icon in an icon theme for a particular window scale,
scales it to the given size and renders it into a pixbuf. This is a
convenience function; if more details about the icon are needed,
use gtk_icon_theme_lookup_icon() followed by
gtk_icon_info_load_icon().

Note that you probably want to listen for icon theme changes and
update the icon. This is usually done by connecting to the
GtkWidget::style-set signal. If for some reason you do not want to
update the icon when the icon theme changes, you should consider
using gdk_pixbuf_copy() to make a private copy of the pixbuf
returned by this function. Otherwise GTK+ may need to keep the old
icon theme loaded, which would be a waste of memory.
*/
func (self *_TraitIconTheme) LoadIconForScale(icon_name string, size int, scale int, flags C.GtkIconLookupFlags) (return__ *C.GdkPixbuf, __err__ error) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_theme_load_icon_for_scale(self.CPointer, __cgo__icon_name, C.gint(size), C.gint(scale), flags, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__icon_name))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Looks up an icon in an icon theme for a particular window scale,
scales it to the given size and renders it into a cairo surface. This is a
convenience function; if more details about the icon are needed,
use gtk_icon_theme_lookup_icon() followed by
gtk_icon_info_load_surface().

Note that you probably want to listen for icon theme changes and
update the icon. This is usually done by connecting to the
GtkWidget::style-set signal.
*/
func (self *_TraitIconTheme) LoadSurface(icon_name string, size int, scale int, for_window *C.GdkWindow, flags C.GtkIconLookupFlags) (return__ *C.cairo_surface_t, __err__ error) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	var __cgo_error__ *C.GError
	return__ = C.gtk_icon_theme_load_surface(self.CPointer, __cgo__icon_name, C.gint(size), C.gint(scale), for_window, flags, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__icon_name))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Looks up an icon and returns a #GtkIconInfo containing
information such as the filename of the icon.
The icon can then be rendered into a pixbuf using
gtk_icon_info_load_icon().
*/
func (self *_TraitIconTheme) LookupByGicon(icon *C.GIcon, size int, flags C.GtkIconLookupFlags) (return__ *IconInfo) {
	var __cgo__return__ *C.GtkIconInfo
	__cgo__return__ = C.gtk_icon_theme_lookup_by_gicon(self.CPointer, icon, C.gint(size), flags)
	return__ = NewIconInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Looks up an icon and returns a #GtkIconInfo containing
information such as the filename of the icon.
The icon can then be rendered into a pixbuf using
gtk_icon_info_load_icon().
*/
func (self *_TraitIconTheme) LookupByGiconForScale(icon *C.GIcon, size int, scale int, flags C.GtkIconLookupFlags) (return__ *IconInfo) {
	var __cgo__return__ *C.GtkIconInfo
	__cgo__return__ = C.gtk_icon_theme_lookup_by_gicon_for_scale(self.CPointer, icon, C.gint(size), C.gint(scale), flags)
	return__ = NewIconInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Looks up a named icon and returns a #GtkIconInfo containing
information such as the filename of the icon. The icon
can then be rendered into a pixbuf using
gtk_icon_info_load_icon(). (gtk_icon_theme_load_icon()
combines these two steps if all you need is the pixbuf.)
*/
func (self *_TraitIconTheme) LookupIcon(icon_name string, size int, flags C.GtkIconLookupFlags) (return__ *IconInfo) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	var __cgo__return__ *C.GtkIconInfo
	__cgo__return__ = C.gtk_icon_theme_lookup_icon(self.CPointer, __cgo__icon_name, C.gint(size), flags)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return__ = NewIconInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Looks up a named icon for a particular window scale and returns a
#GtkIconInfo containing information such as the filename of the
icon. The icon can then be rendered into a pixbuf using
gtk_icon_info_load_icon(). (gtk_icon_theme_load_icon() combines
these two steps if all you need is the pixbuf.)
*/
func (self *_TraitIconTheme) LookupIconForScale(icon_name string, size int, scale int, flags C.GtkIconLookupFlags) (return__ *IconInfo) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	var __cgo__return__ *C.GtkIconInfo
	__cgo__return__ = C.gtk_icon_theme_lookup_icon_for_scale(self.CPointer, __cgo__icon_name, C.gint(size), C.gint(scale), flags)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return__ = NewIconInfoFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Prepends a directory to the search path.
See gtk_icon_theme_set_search_path().
*/
func (self *_TraitIconTheme) PrependSearchPath(path string) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	C.gtk_icon_theme_prepend_search_path(self.CPointer, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Checks to see if the icon theme has changed; if it has, any
currently cached information is discarded and will be reloaded
next time @icon_theme is accessed.
*/
func (self *_TraitIconTheme) RescanIfNeeded() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_theme_rescan_if_needed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the name of the icon theme that the #GtkIconTheme object uses
overriding system configuration. This function cannot be called
on the icon theme objects returned from gtk_icon_theme_get_default()
and gtk_icon_theme_get_for_screen().
*/
func (self *_TraitIconTheme) SetCustomTheme(theme_name string) {
	__cgo__theme_name := (*C.gchar)(unsafe.Pointer(C.CString(theme_name)))
	C.gtk_icon_theme_set_custom_theme(self.CPointer, __cgo__theme_name)
	C.free(unsafe.Pointer(__cgo__theme_name))
	return
}

/*
Sets the screen for an icon theme; the screen is used
to track the user’s currently configured icon theme,
which might be different for different screens.
*/
func (self *_TraitIconTheme) SetScreen(screen *C.GdkScreen) {
	C.gtk_icon_theme_set_screen(self.CPointer, screen)
	return
}

// gtk_icon_theme_set_search_path is not generated due to explicit ignore

type _TraitIconView struct{ CPointer *C.GtkIconView }

/*
Converts widget coordinates to coordinates for the bin_window,
as expected by e.g. gtk_icon_view_get_path_at_pos().
*/
func (self *_TraitIconView) ConvertWidgetToBinWindowCoords(wx int, wy int) (bx int, by int) {
	var __cgo__bx C.gint
	var __cgo__by C.gint
	C.gtk_icon_view_convert_widget_to_bin_window_coords(self.CPointer, C.gint(wx), C.gint(wy), &__cgo__bx, &__cgo__by)
	bx = int(__cgo__bx)
	by = int(__cgo__by)
	return
}

/*
Creates a #cairo_surface_t representation of the item at @path.
This image is used for a drag icon.
*/
func (self *_TraitIconView) CreateDragIcon(path *TreePath) (return__ *C.cairo_surface_t) {
	return__ = C.gtk_icon_view_create_drag_icon(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Turns @icon_view into a drop destination for automatic DND. Calling this
method sets #GtkIconView:reorderable to %FALSE.
*/
func (self *_TraitIconView) EnableModelDragDest(targets *TargetEntry, n_targets int, actions C.GdkDragAction) {
	C.gtk_icon_view_enable_model_drag_dest(self.CPointer, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.gint(n_targets), actions)
	return
}

/*
Turns @icon_view into a drag source for automatic DND. Calling this
method sets #GtkIconView:reorderable to %FALSE.
*/
func (self *_TraitIconView) EnableModelDragSource(start_button_mask C.GdkModifierType, targets *TargetEntry, n_targets int, actions C.GdkDragAction) {
	C.gtk_icon_view_enable_model_drag_source(self.CPointer, start_button_mask, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.gint(n_targets), actions)
	return
}

/*
Gets the setting set by gtk_icon_view_set_activate_on_single_click().
*/
func (self *_TraitIconView) GetActivateOnSingleClick() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_view_get_activate_on_single_click(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Fills the bounding rectangle in widget coordinates for the cell specified by
@path and @cell. If @cell is %NULL the main cell area is used.

This function is only valid if @icon_view is realized.
*/
func (self *_TraitIconView) GetCellRect(path *TreePath, cell *CellRenderer) (rect C.GdkRectangle, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_view_get_cell_rect(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkCellRenderer)(cell.CPointer), &rect)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the value of the ::column-spacing property.
*/
func (self *_TraitIconView) GetColumnSpacing() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_column_spacing(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the value of the ::columns property.
*/
func (self *_TraitIconView) GetColumns() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_columns(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Fills in @path and @cell with the current cursor path and cell.
If the cursor isn’t currently set, then *@path will be %NULL.
If no cell currently has focus, then *@cell will be %NULL.

The returned #GtkTreePath must be freed with gtk_tree_path_free().
*/
func (self *_TraitIconView) GetCursor() (path *TreePath, cell *CellRenderer, return__ bool) {
	var __cgo__path *C.GtkTreePath
	var __cgo__cell *C.GtkCellRenderer
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_view_get_cursor(self.CPointer, &__cgo__path, &__cgo__cell)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	cell = NewCellRendererFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__cell).Pointer()))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines the destination item for a given position.
*/
func (self *_TraitIconView) GetDestItemAtPos(drag_x int, drag_y int) (path *TreePath, pos C.GtkIconViewDropPosition, return__ bool) {
	var __cgo__path *C.GtkTreePath
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_view_get_dest_item_at_pos(self.CPointer, C.gint(drag_x), C.gint(drag_y), &__cgo__path, &pos)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets information about the item that is highlighted for feedback.
*/
func (self *_TraitIconView) GetDragDestItem() (path *TreePath, pos C.GtkIconViewDropPosition) {
	var __cgo__path *C.GtkTreePath
	C.gtk_icon_view_get_drag_dest_item(self.CPointer, &__cgo__path, &pos)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	return
}

/*
Finds the path at the point (@x, @y), relative to bin_window coordinates.
In contrast to gtk_icon_view_get_path_at_pos(), this function also
obtains the cell at the specified position. The returned path should
be freed with gtk_tree_path_free().
See gtk_icon_view_convert_widget_to_bin_window_coords() for converting
widget coordinates to bin_window coordinates.
*/
func (self *_TraitIconView) GetItemAtPos(x int, y int) (path *TreePath, cell *CellRenderer, return__ bool) {
	var __cgo__path *C.GtkTreePath
	var __cgo__cell *C.GtkCellRenderer
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_view_get_item_at_pos(self.CPointer, C.gint(x), C.gint(y), &__cgo__path, &__cgo__cell)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	cell = NewCellRendererFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__cell).Pointer()))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the column in which the item @path is currently
displayed. Column numbers start at 0.
*/
func (self *_TraitIconView) GetItemColumn(path *TreePath) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_item_column(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return__ = int(__cgo__return__)
	return
}

/*
Returns the value of the ::item-orientation property which determines
whether the labels are drawn beside the icons instead of below.
*/
func (self *_TraitIconView) GetItemOrientation() (return__ C.GtkOrientation) {
	return__ = C.gtk_icon_view_get_item_orientation(self.CPointer)
	return
}

/*
Returns the value of the ::item-padding property.
*/
func (self *_TraitIconView) GetItemPadding() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_item_padding(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the row in which the item @path is currently
displayed. Row numbers start at 0.
*/
func (self *_TraitIconView) GetItemRow(path *TreePath) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_item_row(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return__ = int(__cgo__return__)
	return
}

/*
Returns the value of the ::item-width property.
*/
func (self *_TraitIconView) GetItemWidth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_item_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the value of the ::margin property.
*/
func (self *_TraitIconView) GetMargin() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_margin(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the column with markup text for @icon_view.
*/
func (self *_TraitIconView) GetMarkupColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_markup_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the model the #GtkIconView is based on.  Returns %NULL if the
model is unset.
*/
func (self *_TraitIconView) GetModel() (return__ *C.GtkTreeModel) {
	return__ = C.gtk_icon_view_get_model(self.CPointer)
	return
}

/*
Finds the path at the point (@x, @y), relative to bin_window coordinates.
See gtk_icon_view_get_item_at_pos(), if you are also interested in
the cell at the specified position.
See gtk_icon_view_convert_widget_to_bin_window_coords() for converting
widget coordinates to bin_window coordinates.
*/
func (self *_TraitIconView) GetPathAtPos(x int, y int) (return__ *TreePath) {
	var __cgo__return__ *C.GtkTreePath
	__cgo__return__ = C.gtk_icon_view_get_path_at_pos(self.CPointer, C.gint(x), C.gint(y))
	return__ = (*TreePath)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the column with pixbufs for @icon_view.
*/
func (self *_TraitIconView) GetPixbufColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_pixbuf_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves whether the user can reorder the list via drag-and-drop.
See gtk_icon_view_set_reorderable().
*/
func (self *_TraitIconView) GetReorderable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_view_get_reorderable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the value of the ::row-spacing property.
*/
func (self *_TraitIconView) GetRowSpacing() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_row_spacing(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Creates a list of paths of all selected items. Additionally, if you are
planning on modifying the model after calling this function, you may
want to convert the returned list into a list of #GtkTreeRowReferences.
To do this, you can use gtk_tree_row_reference_new().

To free the return value, use:
|[<!-- language="C" -->
g_list_free_full (list, (GDestroyNotify) gtk_tree_path_free);
]|
*/
func (self *_TraitIconView) GetSelectedItems() (return__ *C.GList) {
	return__ = C.gtk_icon_view_get_selected_items(self.CPointer)
	return
}

/*
Gets the selection mode of the @icon_view.
*/
func (self *_TraitIconView) GetSelectionMode() (return__ C.GtkSelectionMode) {
	return__ = C.gtk_icon_view_get_selection_mode(self.CPointer)
	return
}

/*
Returns the value of the ::spacing property.
*/
func (self *_TraitIconView) GetSpacing() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_spacing(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the column with text for @icon_view.
*/
func (self *_TraitIconView) GetTextColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_text_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the column of @icon_view’s model which is being used for
displaying tooltips on @icon_view’s rows.
*/
func (self *_TraitIconView) GetTooltipColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_icon_view_get_tooltip_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

// gtk_icon_view_get_tooltip_context is not generated due to inout param

/*
Sets @start_path and @end_path to be the first and last visible path.
Note that there may be invisible paths in between.

Both paths should be freed with gtk_tree_path_free() after use.
*/
func (self *_TraitIconView) GetVisibleRange() (start_path *TreePath, end_path *TreePath, return__ bool) {
	var __cgo__start_path *C.GtkTreePath
	var __cgo__end_path *C.GtkTreePath
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_view_get_visible_range(self.CPointer, &__cgo__start_path, &__cgo__end_path)
	start_path = (*TreePath)(unsafe.Pointer(__cgo__start_path))
	end_path = (*TreePath)(unsafe.Pointer(__cgo__end_path))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Activates the item determined by @path.
*/
func (self *_TraitIconView) ItemActivated(path *TreePath) {
	C.gtk_icon_view_item_activated(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Returns %TRUE if the icon pointed to by @path is currently
selected. If @path does not point to a valid location, %FALSE is returned.
*/
func (self *_TraitIconView) PathIsSelected(path *TreePath) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_icon_view_path_is_selected(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Moves the alignments of @icon_view to the position specified by @path.
@row_align determines where the row is placed, and @col_align determines
where @column is placed.  Both are expected to be between 0.0 and 1.0.
0.0 means left/top alignment, 1.0 means right/bottom alignment, 0.5 means
center.

If @use_align is %FALSE, then the alignment arguments are ignored, and the
tree does the minimum amount of work to scroll the item onto the screen.
This means that the item will be scrolled to the edge closest to its current
position.  If the item is currently visible on the screen, nothing is done.

This function only works if the model is set, and @path is a valid row on
the model. If the model changes before the @icon_view is realized, the
centered path will be modified to reflect this change.
*/
func (self *_TraitIconView) ScrollToPath(path *TreePath, use_align bool, row_align float32, col_align float32) {
	__cgo__use_align := C.gboolean(0)
	if use_align {
		__cgo__use_align = C.gboolean(1)
	}
	C.gtk_icon_view_scroll_to_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), __cgo__use_align, C.gfloat(row_align), C.gfloat(col_align))
	return
}

/*
Selects all the icons. @icon_view must has its selection mode set
to #GTK_SELECTION_MULTIPLE.
*/
func (self *_TraitIconView) SelectAll() {
	C.gtk_icon_view_select_all(self.CPointer)
	return
}

/*
Selects the row at @path.
*/
func (self *_TraitIconView) SelectPath(path *TreePath) {
	C.gtk_icon_view_select_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Calls a function for each selected icon. Note that the model or
selection cannot be modified from within this function.
*/
func (self *_TraitIconView) SelectedForeach(func_ C.GtkIconViewForeachFunc, data unsafe.Pointer) {
	C.gtk_icon_view_selected_foreach(self.CPointer, func_, (C.gpointer)(data))
	return
}

/*
Causes the #GtkIconView::item-activated signal to be emitted on
a single click instead of a double click.
*/
func (self *_TraitIconView) SetActivateOnSingleClick(single bool) {
	__cgo__single := C.gboolean(0)
	if single {
		__cgo__single = C.gboolean(1)
	}
	C.gtk_icon_view_set_activate_on_single_click(self.CPointer, __cgo__single)
	return
}

/*
Sets the ::column-spacing property which specifies the space
which is inserted between the columns of the icon view.
*/
func (self *_TraitIconView) SetColumnSpacing(column_spacing int) {
	C.gtk_icon_view_set_column_spacing(self.CPointer, C.gint(column_spacing))
	return
}

/*
Sets the ::columns property which determines in how
many columns the icons are arranged. If @columns is
-1, the number of columns will be chosen automatically
to fill the available area.
*/
func (self *_TraitIconView) SetColumns(columns int) {
	C.gtk_icon_view_set_columns(self.CPointer, C.gint(columns))
	return
}

/*
Sets the current keyboard focus to be at @path, and selects it.  This is
useful when you want to focus the user’s attention on a particular item.
If @cell is not %NULL, then focus is given to the cell specified by
it. Additionally, if @start_editing is %TRUE, then editing should be
started in the specified cell.

This function is often followed by `gtk_widget_grab_focus
(icon_view)` in order to give keyboard focus to the widget.
Please note that editing can only happen when the widget is realized.
*/
func (self *_TraitIconView) SetCursor(path *TreePath, cell *CellRenderer, start_editing bool) {
	__cgo__start_editing := C.gboolean(0)
	if start_editing {
		__cgo__start_editing = C.gboolean(1)
	}
	C.gtk_icon_view_set_cursor(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkCellRenderer)(cell.CPointer), __cgo__start_editing)
	return
}

/*
Sets the item that is highlighted for feedback.
*/
func (self *_TraitIconView) SetDragDestItem(path *TreePath, pos C.GtkIconViewDropPosition) {
	C.gtk_icon_view_set_drag_dest_item(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), pos)
	return
}

/*
Sets the ::item-orientation property which determines whether the labels
are drawn beside the icons instead of below.
*/
func (self *_TraitIconView) SetItemOrientation(orientation C.GtkOrientation) {
	C.gtk_icon_view_set_item_orientation(self.CPointer, orientation)
	return
}

/*
Sets the #GtkIconView:item-padding property which specifies the padding
around each of the icon view’s items.
*/
func (self *_TraitIconView) SetItemPadding(item_padding int) {
	C.gtk_icon_view_set_item_padding(self.CPointer, C.gint(item_padding))
	return
}

/*
Sets the ::item-width property which specifies the width
to use for each item. If it is set to -1, the icon view will
automatically determine a suitable item size.
*/
func (self *_TraitIconView) SetItemWidth(item_width int) {
	C.gtk_icon_view_set_item_width(self.CPointer, C.gint(item_width))
	return
}

/*
Sets the ::margin property which specifies the space
which is inserted at the top, bottom, left and right
of the icon view.
*/
func (self *_TraitIconView) SetMargin(margin int) {
	C.gtk_icon_view_set_margin(self.CPointer, C.gint(margin))
	return
}

/*
Sets the column with markup information for @icon_view to be
@column. The markup column must be of type #G_TYPE_STRING.
If the markup column is set to something, it overrides
the text column set by gtk_icon_view_set_text_column().
*/
func (self *_TraitIconView) SetMarkupColumn(column int) {
	C.gtk_icon_view_set_markup_column(self.CPointer, C.gint(column))
	return
}

/*
Sets the model for a #GtkIconView.
If the @icon_view already has a model set, it will remove
it before setting the new model.  If @model is %NULL, then
it will unset the old model.
*/
func (self *_TraitIconView) SetModel(model *C.GtkTreeModel) {
	C.gtk_icon_view_set_model(self.CPointer, model)
	return
}

/*
Sets the column with pixbufs for @icon_view to be @column. The pixbuf
column must be of type #GDK_TYPE_PIXBUF
*/
func (self *_TraitIconView) SetPixbufColumn(column int) {
	C.gtk_icon_view_set_pixbuf_column(self.CPointer, C.gint(column))
	return
}

/*
This function is a convenience function to allow you to reorder models that
support the #GtkTreeDragSourceIface and the #GtkTreeDragDestIface.  Both
#GtkTreeStore and #GtkListStore support these.  If @reorderable is %TRUE, then
the user can reorder the model by dragging and dropping rows.  The
developer can listen to these changes by connecting to the model's
row_inserted and row_deleted signals. The reordering is implemented by setting up
the icon view as a drag source and destination. Therefore, drag and
drop can not be used in a reorderable view for any other purpose.

This function does not give you any degree of control over the order -- any
reordering is allowed.  If more control is needed, you should probably
handle drag and drop manually.
*/
func (self *_TraitIconView) SetReorderable(reorderable bool) {
	__cgo__reorderable := C.gboolean(0)
	if reorderable {
		__cgo__reorderable = C.gboolean(1)
	}
	C.gtk_icon_view_set_reorderable(self.CPointer, __cgo__reorderable)
	return
}

/*
Sets the ::row-spacing property which specifies the space
which is inserted between the rows of the icon view.
*/
func (self *_TraitIconView) SetRowSpacing(row_spacing int) {
	C.gtk_icon_view_set_row_spacing(self.CPointer, C.gint(row_spacing))
	return
}

/*
Sets the selection mode of the @icon_view.
*/
func (self *_TraitIconView) SetSelectionMode(mode C.GtkSelectionMode) {
	C.gtk_icon_view_set_selection_mode(self.CPointer, mode)
	return
}

/*
Sets the ::spacing property which specifies the space
which is inserted between the cells (i.e. the icon and
the text) of an item.
*/
func (self *_TraitIconView) SetSpacing(spacing int) {
	C.gtk_icon_view_set_spacing(self.CPointer, C.gint(spacing))
	return
}

/*
Sets the column with text for @icon_view to be @column. The text
column must be of type #G_TYPE_STRING.
*/
func (self *_TraitIconView) SetTextColumn(column int) {
	C.gtk_icon_view_set_text_column(self.CPointer, C.gint(column))
	return
}

/*
Sets the tip area of @tooltip to the area which @cell occupies in
the item pointed to by @path. See also gtk_tooltip_set_tip_area().

See also gtk_icon_view_set_tooltip_column() for a simpler alternative.
*/
func (self *_TraitIconView) SetTooltipCell(tooltip *Tooltip, path *TreePath, cell *CellRenderer) {
	C.gtk_icon_view_set_tooltip_cell(self.CPointer, (*C.GtkTooltip)(tooltip.CPointer), (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkCellRenderer)(cell.CPointer))
	return
}

/*
If you only plan to have simple (text-only) tooltips on full items, you
can use this function to have #GtkIconView handle these automatically
for you. @column should be set to the column in @icon_view’s model
containing the tooltip texts, or -1 to disable this feature.

When enabled, #GtkWidget:has-tooltip will be set to %TRUE and
@icon_view will connect a #GtkWidget::query-tooltip signal handler.

Note that the signal handler sets the text with gtk_tooltip_set_markup(),
so &, <, etc have to be escaped in the text.
*/
func (self *_TraitIconView) SetTooltipColumn(column int) {
	C.gtk_icon_view_set_tooltip_column(self.CPointer, C.gint(column))
	return
}

/*
Sets the tip area of @tooltip to be the area covered by the item at @path.
See also gtk_icon_view_set_tooltip_column() for a simpler alternative.
See also gtk_tooltip_set_tip_area().
*/
func (self *_TraitIconView) SetTooltipItem(tooltip *Tooltip, path *TreePath) {
	C.gtk_icon_view_set_tooltip_item(self.CPointer, (*C.GtkTooltip)(tooltip.CPointer), (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Unselects all the icons.
*/
func (self *_TraitIconView) UnselectAll() {
	C.gtk_icon_view_unselect_all(self.CPointer)
	return
}

/*
Unselects the row at @path.
*/
func (self *_TraitIconView) UnselectPath(path *TreePath) {
	C.gtk_icon_view_unselect_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Undoes the effect of gtk_icon_view_enable_model_drag_dest(). Calling this
method sets #GtkIconView:reorderable to %FALSE.
*/
func (self *_TraitIconView) UnsetModelDragDest() {
	C.gtk_icon_view_unset_model_drag_dest(self.CPointer)
	return
}

/*
Undoes the effect of gtk_icon_view_enable_model_drag_source(). Calling this
method sets #GtkIconView:reorderable to %FALSE.
*/
func (self *_TraitIconView) UnsetModelDragSource() {
	C.gtk_icon_view_unset_model_drag_source(self.CPointer)
	return
}

type _TraitImage struct{ CPointer *C.GtkImage }

/*
Resets the image to be empty.
*/
func (self *_TraitImage) Clear() {
	C.gtk_image_clear(self.CPointer)
	return
}

/*
Gets the #GdkPixbufAnimation being displayed by the #GtkImage.
The storage type of the image must be %GTK_IMAGE_EMPTY or
%GTK_IMAGE_ANIMATION (see gtk_image_get_storage_type()).
The caller of this function does not own a reference to the
returned animation.
*/
func (self *_TraitImage) GetAnimation() (return__ *C.GdkPixbufAnimation) {
	return__ = C.gtk_image_get_animation(self.CPointer)
	return
}

/*
Gets the #GIcon and size being displayed by the #GtkImage.
The storage type of the image must be %GTK_IMAGE_EMPTY or
%GTK_IMAGE_GICON (see gtk_image_get_storage_type()).
The caller of this function does not own a reference to the
returned #GIcon.
*/
func (self *_TraitImage) GetGicon() (gicon *C.GIcon, size C.GtkIconSize) {
	C.gtk_image_get_gicon(self.CPointer, &gicon, &size)
	return
}

/*
Gets the icon name and size being displayed by the #GtkImage.
The storage type of the image must be %GTK_IMAGE_EMPTY or
%GTK_IMAGE_ICON_NAME (see gtk_image_get_storage_type()).
The returned string is owned by the #GtkImage and should not
be freed.
*/
func (self *_TraitImage) GetIconName() (icon_name string, size C.GtkIconSize) {
	var __cgo__icon_name *C.gchar
	C.gtk_image_get_icon_name(self.CPointer, &__cgo__icon_name, &size)
	icon_name = C.GoString((*C.char)(unsafe.Pointer(__cgo__icon_name)))
	return
}

// gtk_image_get_icon_set is not generated due to deprecation attr

/*
Gets the #GdkPixbuf being displayed by the #GtkImage.
The storage type of the image must be %GTK_IMAGE_EMPTY or
%GTK_IMAGE_PIXBUF (see gtk_image_get_storage_type()).
The caller of this function does not own a reference to the
returned pixbuf.
*/
func (self *_TraitImage) GetPixbuf() (return__ *C.GdkPixbuf) {
	return__ = C.gtk_image_get_pixbuf(self.CPointer)
	return
}

/*
Gets the pixel size used for named icons.
*/
func (self *_TraitImage) GetPixelSize() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_image_get_pixel_size(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

// gtk_image_get_stock is not generated due to deprecation attr

/*
Gets the type of representation being used by the #GtkImage
to store image data. If the #GtkImage has no image data,
the return value will be %GTK_IMAGE_EMPTY.
*/
func (self *_TraitImage) GetStorageType() (return__ C.GtkImageType) {
	return__ = C.gtk_image_get_storage_type(self.CPointer)
	return
}

/*
Causes the #GtkImage to display the given animation (or display
nothing, if you set the animation to %NULL).
*/
func (self *_TraitImage) SetFromAnimation(animation *C.GdkPixbufAnimation) {
	C.gtk_image_set_from_animation(self.CPointer, animation)
	return
}

/*
See gtk_image_new_from_file() for details.
*/
func (self *_TraitImage) SetFromFile(filename string) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	C.gtk_image_set_from_file(self.CPointer, __cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	return
}

/*
See gtk_image_new_from_gicon() for details.
*/
func (self *_TraitImage) SetFromGicon(icon *C.GIcon, size C.GtkIconSize) {
	C.gtk_image_set_from_gicon(self.CPointer, icon, size)
	return
}

/*
See gtk_image_new_from_icon_name() for details.
*/
func (self *_TraitImage) SetFromIconName(icon_name string, size C.GtkIconSize) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_image_set_from_icon_name(self.CPointer, __cgo__icon_name, size)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

// gtk_image_set_from_icon_set is not generated due to deprecation attr

/*
See gtk_image_new_from_pixbuf() for details.
*/
func (self *_TraitImage) SetFromPixbuf(pixbuf *C.GdkPixbuf) {
	C.gtk_image_set_from_pixbuf(self.CPointer, pixbuf)
	return
}

/*
See gtk_image_new_from_resource() for details.
*/
func (self *_TraitImage) SetFromResource(resource_path string) {
	__cgo__resource_path := (*C.gchar)(unsafe.Pointer(C.CString(resource_path)))
	C.gtk_image_set_from_resource(self.CPointer, __cgo__resource_path)
	C.free(unsafe.Pointer(__cgo__resource_path))
	return
}

// gtk_image_set_from_stock is not generated due to deprecation attr

/*
See gtk_image_new_from_surface() for details.
*/
func (self *_TraitImage) SetFromSurface(surface *C.cairo_surface_t) {
	C.gtk_image_set_from_surface(self.CPointer, surface)
	return
}

/*
Sets the pixel size to use for named icons. If the pixel size is set
to a value != -1, it is used instead of the icon size set by
gtk_image_set_from_icon_name().
*/
func (self *_TraitImage) SetPixelSize(pixel_size int) {
	C.gtk_image_set_pixel_size(self.CPointer, C.gint(pixel_size))
	return
}

type _TraitImageMenuItem struct{ CPointer *C.GtkImageMenuItem }

// gtk_image_menu_item_get_always_show_image is not generated due to deprecation attr

// gtk_image_menu_item_get_image is not generated due to deprecation attr

// gtk_image_menu_item_get_use_stock is not generated due to deprecation attr

// gtk_image_menu_item_set_accel_group is not generated due to deprecation attr

// gtk_image_menu_item_set_always_show_image is not generated due to deprecation attr

// gtk_image_menu_item_set_image is not generated due to deprecation attr

// gtk_image_menu_item_set_use_stock is not generated due to deprecation attr

type _TraitInfoBar struct{ CPointer *C.GtkInfoBar }

/*
Add an activatable widget to the action area of a #GtkInfoBar,
connecting a signal handler that will emit the #GtkInfoBar::response
signal on the message area when the widget is activated. The widget
is appended to the end of the message areas action area.
*/
func (self *_TraitInfoBar) AddActionWidget(child *Widget, response_id int) {
	C.gtk_info_bar_add_action_widget(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(response_id))
	return
}

/*
Adds a button with the given text and sets things up so that
clicking the button will emit the “response” signal with the given
response_id. The button is appended to the end of the info bars's
action area. The button widget is returned, but usually you don't
need it.
*/
func (self *_TraitInfoBar) AddButton(button_text string, response_id int) (return__ *Widget) {
	__cgo__button_text := (*C.gchar)(unsafe.Pointer(C.CString(button_text)))
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_info_bar_add_button(self.CPointer, __cgo__button_text, C.gint(response_id))
	C.free(unsafe.Pointer(__cgo__button_text))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

// gtk_info_bar_add_buttons is not generated due to varargs

/*
Returns the action area of @info_bar.
*/
func (self *_TraitInfoBar) GetActionArea() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_info_bar_get_action_area(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the content area of @info_bar.
*/
func (self *_TraitInfoBar) GetContentArea() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_info_bar_get_content_area(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the message type of the message area.
*/
func (self *_TraitInfoBar) GetMessageType() (return__ C.GtkMessageType) {
	return__ = C.gtk_info_bar_get_message_type(self.CPointer)
	return
}

/*
Returns whether the widget will display a standard close button.
*/
func (self *_TraitInfoBar) GetShowCloseButton() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_info_bar_get_show_close_button(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Emits the “response” signal with the given @response_id.
*/
func (self *_TraitInfoBar) Response(response_id int) {
	C.gtk_info_bar_response(self.CPointer, C.gint(response_id))
	return
}

/*
Sets the last widget in the info bar’s action area with
the given response_id as the default widget for the dialog.
Pressing “Enter” normally activates the default widget.

Note that this function currently requires @info_bar to
be added to a widget hierarchy.
*/
func (self *_TraitInfoBar) SetDefaultResponse(response_id int) {
	C.gtk_info_bar_set_default_response(self.CPointer, C.gint(response_id))
	return
}

/*
Sets the message type of the message area.
GTK+ uses this type to determine what color to use
when drawing the message area.
*/
func (self *_TraitInfoBar) SetMessageType(message_type C.GtkMessageType) {
	C.gtk_info_bar_set_message_type(self.CPointer, message_type)
	return
}

/*
Calls gtk_widget_set_sensitive (widget, setting) for each
widget in the info bars’s action area with the given response_id.
A convenient way to sensitize/desensitize dialog buttons.
*/
func (self *_TraitInfoBar) SetResponseSensitive(response_id int, setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_info_bar_set_response_sensitive(self.CPointer, C.gint(response_id), __cgo__setting)
	return
}

/*
If true, a standard close button is shown. When clicked it emits
the response %GTK_RESPONSE_CLOSE.
*/
func (self *_TraitInfoBar) SetShowCloseButton(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_info_bar_set_show_close_button(self.CPointer, __cgo__setting)
	return
}

type _TraitInvisible struct{ CPointer *C.GtkInvisible }

/*
Returns the #GdkScreen object associated with @invisible
*/
func (self *_TraitInvisible) GetScreen() (return__ *C.GdkScreen) {
	return__ = C.gtk_invisible_get_screen(self.CPointer)
	return
}

/*
Sets the #GdkScreen where the #GtkInvisible object will be displayed.
*/
func (self *_TraitInvisible) SetScreen(screen *C.GdkScreen) {
	C.gtk_invisible_set_screen(self.CPointer, screen)
	return
}

type _TraitLabel struct{ CPointer *C.GtkLabel }

/*
Gets the angle of rotation for the label. See
gtk_label_set_angle().
*/
func (self *_TraitLabel) GetAngle() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_label_get_angle(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the attribute list that was set on the label using
gtk_label_set_attributes(), if any. This function does
not reflect attributes that come from the labels markup
(see gtk_label_set_markup()). If you want to get the
effective attributes for the label, use
pango_layout_get_attribute (gtk_label_get_layout (label)).
*/
func (self *_TraitLabel) GetAttributes() (return__ *C.PangoAttrList) {
	return__ = C.gtk_label_get_attributes(self.CPointer)
	return
}

/*
Returns the URI for the currently active link in the label.
The active link is the one under the mouse pointer or, in a
selectable label, the link in which the text cursor is currently
positioned.

This function is intended for use in a #GtkLabel::activate-link handler
or for use in a #GtkWidget::query-tooltip handler.
*/
func (self *_TraitLabel) GetCurrentUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_label_get_current_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the ellipsizing position of the label. See gtk_label_set_ellipsize().
*/
func (self *_TraitLabel) GetEllipsize() (return__ C.PangoEllipsizeMode) {
	return__ = C.gtk_label_get_ellipsize(self.CPointer)
	return
}

/*
Returns the justification of the label. See gtk_label_set_justify().
*/
func (self *_TraitLabel) GetJustify() (return__ C.GtkJustification) {
	return__ = C.gtk_label_get_justify(self.CPointer)
	return
}

/*
Fetches the text from a label widget including any embedded
underlines indicating mnemonics and Pango markup. (See
gtk_label_get_text()).
*/
func (self *_TraitLabel) GetLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_label_get_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #PangoLayout used to display the label.
The layout is useful to e.g. convert text positions to
pixel positions, in combination with gtk_label_get_layout_offsets().
The returned layout is owned by the @label so need not be
freed by the caller. The @label is free to recreate its layout at
any time, so it should be considered read-only.
*/
func (self *_TraitLabel) GetLayout() (return__ *C.PangoLayout) {
	return__ = C.gtk_label_get_layout(self.CPointer)
	return
}

/*
Obtains the coordinates where the label will draw the #PangoLayout
representing the text in the label; useful to convert mouse events
into coordinates inside the #PangoLayout, e.g. to take some action
if some part of the label is clicked. Of course you will need to
create a #GtkEventBox to receive the events, and pack the label
inside it, since labels are a #GTK_NO_WINDOW widget. Remember
when using the #PangoLayout functions you need to convert to
and from pixels using PANGO_PIXELS() or #PANGO_SCALE.
*/
func (self *_TraitLabel) GetLayoutOffsets() (x int, y int) {
	var __cgo__x C.gint
	var __cgo__y C.gint
	C.gtk_label_get_layout_offsets(self.CPointer, &__cgo__x, &__cgo__y)
	x = int(__cgo__x)
	y = int(__cgo__y)
	return
}

/*
Returns whether lines in the label are automatically wrapped.
See gtk_label_set_line_wrap().
*/
func (self *_TraitLabel) GetLineWrap() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_label_get_line_wrap(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns line wrap mode used by the label. See gtk_label_set_line_wrap_mode().
*/
func (self *_TraitLabel) GetLineWrapMode() (return__ C.PangoWrapMode) {
	return__ = C.gtk_label_get_line_wrap_mode(self.CPointer)
	return
}

/*
Gets the number of lines to which an ellipsized, wrapping
label should be limited. See gtk_label_set_lines().
*/
func (self *_TraitLabel) GetLines() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_label_get_lines(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the desired maximum width of @label, in characters. See
gtk_label_set_width_chars().
*/
func (self *_TraitLabel) GetMaxWidthChars() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_label_get_max_width_chars(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
If the label has been set so that it has an mnemonic key this function
returns the keyval used for the mnemonic accelerator. If there is no
mnemonic set up it returns #GDK_KEY_VoidSymbol.
*/
func (self *_TraitLabel) GetMnemonicKeyval() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_label_get_mnemonic_keyval(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Retrieves the target of the mnemonic (keyboard shortcut) of this
label. See gtk_label_set_mnemonic_widget().
*/
func (self *_TraitLabel) GetMnemonicWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_label_get_mnemonic_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the value set by gtk_label_set_selectable().
*/
func (self *_TraitLabel) GetSelectable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_label_get_selectable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the selected range of characters in the label, returning %TRUE
if there’s a selection.
*/
func (self *_TraitLabel) GetSelectionBounds() (start int, end int, return__ bool) {
	var __cgo__start C.gint
	var __cgo__end C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_label_get_selection_bounds(self.CPointer, &__cgo__start, &__cgo__end)
	start = int(__cgo__start)
	end = int(__cgo__end)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the label is in single line mode.
*/
func (self *_TraitLabel) GetSingleLineMode() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_label_get_single_line_mode(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Fetches the text from a label widget, as displayed on the
screen. This does not include any embedded underlines
indicating mnemonics or Pango markup. (See gtk_label_get_label())
*/
func (self *_TraitLabel) GetText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_label_get_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns whether the label is currently keeping track
of clicked links.
*/
func (self *_TraitLabel) GetTrackVisitedLinks() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_label_get_track_visited_links(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the label’s text is interpreted as marked up with
the [Pango text markup language][PangoMarkupFormat].
See gtk_label_set_use_markup ().
*/
func (self *_TraitLabel) GetUseMarkup() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_label_get_use_markup(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether an embedded underline in the label indicates a
mnemonic. See gtk_label_set_use_underline().
*/
func (self *_TraitLabel) GetUseUnderline() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_label_get_use_underline(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the desired width of @label, in characters. See
gtk_label_set_width_chars().
*/
func (self *_TraitLabel) GetWidthChars() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_label_get_width_chars(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Selects a range of characters in the label, if the label is selectable.
See gtk_label_set_selectable(). If the label is not selectable,
this function has no effect. If @start_offset or
@end_offset are -1, then the end of the label will be substituted.
*/
func (self *_TraitLabel) SelectRegion(start_offset int, end_offset int) {
	C.gtk_label_select_region(self.CPointer, C.gint(start_offset), C.gint(end_offset))
	return
}

/*
Sets the angle of rotation for the label. An angle of 90 reads from
from bottom to top, an angle of 270, from top to bottom. The angle
setting for the label is ignored if the label is selectable,
wrapped, or ellipsized.
*/
func (self *_TraitLabel) SetAngle(angle float64) {
	C.gtk_label_set_angle(self.CPointer, C.gdouble(angle))
	return
}

/*
Sets a #PangoAttrList; the attributes in the list are applied to the
label text.

The attributes set with this function will be applied
and merged with any other attributes previously effected by way
of the #GtkLabel:use-underline or #GtkLabel:use-markup properties.
While it is not recommended to mix markup strings with manually set
attributes, if you must; know that the attributes will be applied
to the label after the markup string is parsed.
*/
func (self *_TraitLabel) SetAttributes(attrs *C.PangoAttrList) {
	C.gtk_label_set_attributes(self.CPointer, attrs)
	return
}

/*
Sets the mode used to ellipsize (add an ellipsis: "...") to the text
if there is not enough space to render the entire string.
*/
func (self *_TraitLabel) SetEllipsize(mode C.PangoEllipsizeMode) {
	C.gtk_label_set_ellipsize(self.CPointer, mode)
	return
}

/*
Sets the alignment of the lines in the text of the label relative to
each other. %GTK_JUSTIFY_LEFT is the default value when the
widget is first created with gtk_label_new(). If you instead want
to set the alignment of the label as a whole, use
gtk_misc_set_alignment() instead. gtk_label_set_justify() has no
effect on labels containing only a single line.
*/
func (self *_TraitLabel) SetJustify(jtype C.GtkJustification) {
	C.gtk_label_set_justify(self.CPointer, jtype)
	return
}

/*
Sets the text of the label. The label is interpreted as
including embedded underlines and/or Pango markup depending
on the values of the #GtkLabel:use-underline" and
#GtkLabel:use-markup properties.
*/
func (self *_TraitLabel) SetLabel(str string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	C.gtk_label_set_label(self.CPointer, __cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return
}

/*
Toggles line wrapping within the #GtkLabel widget. %TRUE makes it break
lines if text exceeds the widget’s size. %FALSE lets the text get cut off
by the edge of the widget if it exceeds the widget size.

Note that setting line wrapping to %TRUE does not make the label
wrap at its parent container’s width, because GTK+ widgets
conceptually can’t make their requisition depend on the parent
container’s size. For a label that wraps at a specific position,
set the label’s width using gtk_widget_set_size_request().
*/
func (self *_TraitLabel) SetLineWrap(wrap bool) {
	__cgo__wrap := C.gboolean(0)
	if wrap {
		__cgo__wrap = C.gboolean(1)
	}
	C.gtk_label_set_line_wrap(self.CPointer, __cgo__wrap)
	return
}

/*
If line wrapping is on (see gtk_label_set_line_wrap()) this controls how
the line wrapping is done. The default is %PANGO_WRAP_WORD which means
wrap on word boundaries.
*/
func (self *_TraitLabel) SetLineWrapMode(wrap_mode C.PangoWrapMode) {
	C.gtk_label_set_line_wrap_mode(self.CPointer, wrap_mode)
	return
}

/*
Sets the number of lines to which an ellipsized, wrapping label
should be limited. This has no effect if the label is not wrapping
or ellipsized. Set this to -1 if you don’t want to limit the
number of lines.
*/
func (self *_TraitLabel) SetLines(lines int) {
	C.gtk_label_set_lines(self.CPointer, C.gint(lines))
	return
}

/*
Parses @str which is marked up with the
[Pango text markup language][PangoMarkupFormat], setting the
label’s text and attribute list based on the parse results. If the @str is
external data, you may need to escape it with g_markup_escape_text() or
g_markup_printf_escaped():
|[<!-- language="C" -->
const char *format = "<span style=\"italic\">\%s</span>";
char *markup;

markup = g_markup_printf_escaped (format, str);
gtk_label_set_markup (GTK_LABEL (label), markup);
g_free (markup);
]|
*/
func (self *_TraitLabel) SetMarkup(str string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	C.gtk_label_set_markup(self.CPointer, __cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return
}

/*
Parses @str which is marked up with the
[Pango text markup language][PangoMarkupFormat],
setting the label’s text and attribute list based on the parse results.
If characters in @str are preceded by an underscore, they are underlined
indicating that they represent a keyboard accelerator called a mnemonic.

The mnemonic key can be used to activate another widget, chosen
automatically, or explicitly using gtk_label_set_mnemonic_widget().
*/
func (self *_TraitLabel) SetMarkupWithMnemonic(str string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	C.gtk_label_set_markup_with_mnemonic(self.CPointer, __cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return
}

/*
Sets the desired maximum width in characters of @label to @n_chars.
*/
func (self *_TraitLabel) SetMaxWidthChars(n_chars int) {
	C.gtk_label_set_max_width_chars(self.CPointer, C.gint(n_chars))
	return
}

/*
If the label has been set so that it has an mnemonic key (using
i.e. gtk_label_set_markup_with_mnemonic(),
gtk_label_set_text_with_mnemonic(), gtk_label_new_with_mnemonic()
or the “use_underline” property) the label can be associated with a
widget that is the target of the mnemonic. When the label is inside
a widget (like a #GtkButton or a #GtkNotebook tab) it is
automatically associated with the correct widget, but sometimes
(i.e. when the target is a #GtkEntry next to the label) you need to
set it explicitly using this function.

The target widget will be accelerated by emitting the
GtkWidget::mnemonic-activate signal on it. The default handler for
this signal will activate the widget if there are no mnemonic collisions
and toggle focus between the colliding widgets otherwise.
*/
func (self *_TraitLabel) SetMnemonicWidget(widget *Widget) {
	C.gtk_label_set_mnemonic_widget(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

/*
The pattern of underlines you want under the existing text within the
#GtkLabel widget.  For example if the current text of the label says
“FooBarBaz” passing a pattern of “___   ___” will underline
“Foo” and “Baz” but not “Bar”.
*/
func (self *_TraitLabel) SetPattern(pattern string) {
	__cgo__pattern := (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	C.gtk_label_set_pattern(self.CPointer, __cgo__pattern)
	C.free(unsafe.Pointer(__cgo__pattern))
	return
}

/*
Selectable labels allow the user to select text from the label, for
copy-and-paste.
*/
func (self *_TraitLabel) SetSelectable(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_label_set_selectable(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether the label is in single line mode.
*/
func (self *_TraitLabel) SetSingleLineMode(single_line_mode bool) {
	__cgo__single_line_mode := C.gboolean(0)
	if single_line_mode {
		__cgo__single_line_mode = C.gboolean(1)
	}
	C.gtk_label_set_single_line_mode(self.CPointer, __cgo__single_line_mode)
	return
}

/*
Sets the text within the #GtkLabel widget. It overwrites any text that
was there before.

This will also clear any previously set mnemonic accelerators.
*/
func (self *_TraitLabel) SetText(str string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	C.gtk_label_set_text(self.CPointer, __cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return
}

/*
Sets the label’s text from the string @str.
If characters in @str are preceded by an underscore, they are underlined
indicating that they represent a keyboard accelerator called a mnemonic.
The mnemonic key can be used to activate another widget, chosen
automatically, or explicitly using gtk_label_set_mnemonic_widget().
*/
func (self *_TraitLabel) SetTextWithMnemonic(str string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	C.gtk_label_set_text_with_mnemonic(self.CPointer, __cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return
}

/*
Sets whether the label should keep track of clicked
links (and use a different color for them).
*/
func (self *_TraitLabel) SetTrackVisitedLinks(track_links bool) {
	__cgo__track_links := C.gboolean(0)
	if track_links {
		__cgo__track_links = C.gboolean(1)
	}
	C.gtk_label_set_track_visited_links(self.CPointer, __cgo__track_links)
	return
}

/*
Sets whether the text of the label contains markup in
[Pango’s text markup language][PangoMarkupFormat].
See gtk_label_set_markup().
*/
func (self *_TraitLabel) SetUseMarkup(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_label_set_use_markup(self.CPointer, __cgo__setting)
	return
}

/*
If true, an underline in the text indicates the next character should be
used for the mnemonic accelerator key.
*/
func (self *_TraitLabel) SetUseUnderline(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_label_set_use_underline(self.CPointer, __cgo__setting)
	return
}

/*
Sets the desired width in characters of @label to @n_chars.
*/
func (self *_TraitLabel) SetWidthChars(n_chars int) {
	C.gtk_label_set_width_chars(self.CPointer, C.gint(n_chars))
	return
}

type _TraitLayout struct{ CPointer *C.GtkLayout }

/*
Retrieve the bin window of the layout used for drawing operations.
*/
func (self *_TraitLayout) GetBinWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_layout_get_bin_window(self.CPointer)
	return
}

// gtk_layout_get_hadjustment is not generated due to deprecation attr

/*
Gets the size that has been set on the layout, and that determines
the total extents of the layout’s scrollbar area. See
gtk_layout_set_size ().
*/
func (self *_TraitLayout) GetSize() (width uint, height uint) {
	var __cgo__width C.guint
	var __cgo__height C.guint
	C.gtk_layout_get_size(self.CPointer, &__cgo__width, &__cgo__height)
	width = uint(__cgo__width)
	height = uint(__cgo__height)
	return
}

// gtk_layout_get_vadjustment is not generated due to deprecation attr

/*
Moves a current child of @layout to a new position.
*/
func (self *_TraitLayout) Move(child_widget *Widget, x int, y int) {
	C.gtk_layout_move(self.CPointer, (*C.GtkWidget)(child_widget.CPointer), C.gint(x), C.gint(y))
	return
}

/*
Adds @child_widget to @layout, at position (@x,@y).
@layout becomes the new parent container of @child_widget.
*/
func (self *_TraitLayout) Put(child_widget *Widget, x int, y int) {
	C.gtk_layout_put(self.CPointer, (*C.GtkWidget)(child_widget.CPointer), C.gint(x), C.gint(y))
	return
}

// gtk_layout_set_hadjustment is not generated due to deprecation attr

/*
Sets the size of the scrollable area of the layout.
*/
func (self *_TraitLayout) SetSize(width uint, height uint) {
	C.gtk_layout_set_size(self.CPointer, C.guint(width), C.guint(height))
	return
}

// gtk_layout_set_vadjustment is not generated due to deprecation attr

type _TraitLevelBar struct{ CPointer *C.GtkLevelBar }

/*
Adds a new offset marker on @self at the position specified by @value.
When the bar value is in the interval topped by @value (or between @value
and #GtkLevelBar:max-value in case the offset is the last one on the bar)
a style class named `level-`@name will be applied
when rendering the level bar fill.
If another offset marker named @name exists, its value will be
replaced by @value.
*/
func (self *_TraitLevelBar) AddOffsetValue(name string, value float64) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_level_bar_add_offset_value(self.CPointer, __cgo__name, C.gdouble(value))
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Return the value of the #GtkLevelBar:inverted property.
*/
func (self *_TraitLevelBar) GetInverted() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_level_bar_get_inverted(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the value of the #GtkLevelBar:max-value property.
*/
func (self *_TraitLevelBar) GetMaxValue() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_level_bar_get_max_value(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Returns the value of the #GtkLevelBar:min-value property.
*/
func (self *_TraitLevelBar) GetMinValue() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_level_bar_get_min_value(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Returns the value of the #GtkLevelBar:mode property.
*/
func (self *_TraitLevelBar) GetMode() (return__ C.GtkLevelBarMode) {
	return__ = C.gtk_level_bar_get_mode(self.CPointer)
	return
}

/*
Fetches the value specified for the offset marker @name in @self,
returning %TRUE in case an offset named @name was found.
*/
func (self *_TraitLevelBar) GetOffsetValue(name string) (value float64, return__ bool) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__value C.gdouble
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_level_bar_get_offset_value(self.CPointer, __cgo__name, &__cgo__value)
	C.free(unsafe.Pointer(__cgo__name))
	value = float64(__cgo__value)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the value of the #GtkLevelBar:value property.
*/
func (self *_TraitLevelBar) GetValue() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_level_bar_get_value(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Removes an offset marker previously added with
gtk_level_bar_add_offset_value().
*/
func (self *_TraitLevelBar) RemoveOffsetValue(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_level_bar_remove_offset_value(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Sets the value of the #GtkLevelBar:inverted property.
*/
func (self *_TraitLevelBar) SetInverted(inverted bool) {
	__cgo__inverted := C.gboolean(0)
	if inverted {
		__cgo__inverted = C.gboolean(1)
	}
	C.gtk_level_bar_set_inverted(self.CPointer, __cgo__inverted)
	return
}

/*
Sets the value of the #GtkLevelBar:max-value property.
*/
func (self *_TraitLevelBar) SetMaxValue(value float64) {
	C.gtk_level_bar_set_max_value(self.CPointer, C.gdouble(value))
	return
}

/*
Sets the value of the #GtkLevelBar:min-value property.
*/
func (self *_TraitLevelBar) SetMinValue(value float64) {
	C.gtk_level_bar_set_min_value(self.CPointer, C.gdouble(value))
	return
}

/*
Sets the value of the #GtkLevelBar:mode property.
*/
func (self *_TraitLevelBar) SetMode(mode C.GtkLevelBarMode) {
	C.gtk_level_bar_set_mode(self.CPointer, mode)
	return
}

/*
Sets the value of the #GtkLevelBar:value property.
*/
func (self *_TraitLevelBar) SetValue(value float64) {
	C.gtk_level_bar_set_value(self.CPointer, C.gdouble(value))
	return
}

type _TraitLinkButton struct{ CPointer *C.GtkLinkButton }

/*
Retrieves the URI set using gtk_link_button_set_uri().
*/
func (self *_TraitLinkButton) GetUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_link_button_get_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the “visited” state of the URI where the #GtkLinkButton
points. The button becomes visited when it is clicked. If the URI
is changed on the button, the “visited” state is unset again.

The state may also be changed using gtk_link_button_set_visited().
*/
func (self *_TraitLinkButton) GetVisited() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_link_button_get_visited(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets @uri as the URI where the #GtkLinkButton points. As a side-effect
this unsets the “visited” state of the button.
*/
func (self *_TraitLinkButton) SetUri(uri string) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	C.gtk_link_button_set_uri(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return
}

/*
Sets the “visited” state of the URI where the #GtkLinkButton
points.  See gtk_link_button_get_visited() for more details.
*/
func (self *_TraitLinkButton) SetVisited(visited bool) {
	__cgo__visited := C.gboolean(0)
	if visited {
		__cgo__visited = C.gboolean(1)
	}
	C.gtk_link_button_set_visited(self.CPointer, __cgo__visited)
	return
}

type _TraitListBox struct{ CPointer *C.GtkListBox }

/*
This is a helper function for implementing DnD onto a #GtkListBox.
The passed in @row will be highlighted via gtk_drag_highlight(),
and any previously highlighted row will be unhighlighted.

The row will also be unhighlighted when the widget gets
a drag leave event.
*/
func (self *_TraitListBox) DragHighlightRow(row *ListBoxRow) {
	C.gtk_list_box_drag_highlight_row(self.CPointer, (*C.GtkListBoxRow)(row.CPointer))
	return
}

/*
If a row has previously been highlighted via gtk_list_box_drag_highlight_row()
it will have the highlight removed.
*/
func (self *_TraitListBox) DragUnhighlightRow() {
	C.gtk_list_box_drag_unhighlight_row(self.CPointer)
	return
}

/*
Returns whether rows activate on single clicks.
*/
func (self *_TraitListBox) GetActivateOnSingleClick() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_list_box_get_activate_on_single_click(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the adjustment (if any) that the widget uses to
for vertical scrolling.
*/
func (self *_TraitListBox) GetAdjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_list_box_get_adjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the n:th child in the list (not counting headers).
*/
func (self *_TraitListBox) GetRowAtIndex(index_ int) (return__ *ListBoxRow) {
	var __cgo__return__ *C.GtkListBoxRow
	__cgo__return__ = C.gtk_list_box_get_row_at_index(self.CPointer, C.gint(index_))
	return__ = NewListBoxRowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the row at the @y position.
*/
func (self *_TraitListBox) GetRowAtY(y int) (return__ *ListBoxRow) {
	var __cgo__return__ *C.GtkListBoxRow
	__cgo__return__ = C.gtk_list_box_get_row_at_y(self.CPointer, C.gint(y))
	return__ = NewListBoxRowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the selected row.
*/
func (self *_TraitListBox) GetSelectedRow() (return__ *ListBoxRow) {
	var __cgo__return__ *C.GtkListBoxRow
	__cgo__return__ = C.gtk_list_box_get_selected_row(self.CPointer)
	return__ = NewListBoxRowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the selection mode of the listbox.
*/
func (self *_TraitListBox) GetSelectionMode() (return__ C.GtkSelectionMode) {
	return__ = C.gtk_list_box_get_selection_mode(self.CPointer)
	return
}

/*
Insert the @child into the @list_box at @position. If a sort function is
set, the widget will actually be inserted at the calculated position and
this function has the same effect of gtk_container_add().

If @position is -1, or larger than the total number of items in the
@list_box, then the @child will be appended to the end.
*/
func (self *_TraitListBox) Insert(child *Widget, position int) {
	C.gtk_list_box_insert(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(position))
	return
}

/*
Update the filtering for all rows. Call this when result
of the filter function on the @list_box is changed due
to an external factor. For instance, this would be used
if the filter function just looked for a specific search
string and the entry with the search string has changed.
*/
func (self *_TraitListBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter(self.CPointer)
	return
}

/*
Update the separators for all rows. Call this when result
of the header function on the @list_box is changed due
to an external factor.
*/
func (self *_TraitListBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers(self.CPointer)
	return
}

/*
Update the sorting for all rows. Call this when result
of the sort function on the @list_box is changed due
to an external factor.
*/
func (self *_TraitListBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort(self.CPointer)
	return
}

/*
Prepend a widget to the list. If a sort function is set, the widget will
actually be inserted at the calculated position and this function has the
same effect of gtk_container_add().
*/
func (self *_TraitListBox) Prepend(child *Widget) {
	C.gtk_list_box_prepend(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Make @row the currently selected row.
*/
func (self *_TraitListBox) SelectRow(row *ListBoxRow) {
	C.gtk_list_box_select_row(self.CPointer, (*C.GtkListBoxRow)(row.CPointer))
	return
}

/*
If @single is %TRUE, rows will be activated when you click on them,
otherwise you need to double-click.
*/
func (self *_TraitListBox) SetActivateOnSingleClick(single bool) {
	__cgo__single := C.gboolean(0)
	if single {
		__cgo__single = C.gboolean(1)
	}
	C.gtk_list_box_set_activate_on_single_click(self.CPointer, __cgo__single)
	return
}

/*
Sets the adjustment (if any) that the widget uses to
for vertical scrolling. For instance, this is used
to get the page size for PageUp/Down key handling.

In the normal case when the @list_box is packed inside
a #GtkScrolledWindow the adjustment from that will
be picked up automatically, so there is no need
to manually do that.
*/
func (self *_TraitListBox) SetAdjustment(adjustment *Adjustment) {
	C.gtk_list_box_set_adjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
By setting a filter function on the @list_box one can decide dynamically which
of the rows to show. For instance, to implement a search function on a list that
filters the original list to only show the matching rows.

The @filter_func will be called for each row after the call, and it will
continue to be called each time a row changes (via gtk_list_box_row_changed()) or
when gtk_list_box_invalidate_filter() is called.
*/
func (self *_TraitListBox) SetFilterFunc(filter_func C.GtkListBoxFilterFunc, user_data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_list_box_set_filter_func(self.CPointer, filter_func, (C.gpointer)(user_data), destroy)
	return
}

/*
By setting a header function on the @list_box one can dynamically add headers
in front of rows, depending on the contents of the row and its position in the list.
For instance, one could use it to add headers in front of the first item of a
new kind, in a list sorted by the kind.

The @update_header can look at the current header widget using gtk_list_box_row_get_header()
and either update the state of the widget as needed, or set a new one using
gtk_list_box_row_set_header(). If no header is needed, set the header to %NULL.

Note that you may get many calls @update_header to this for a particular row when e.g.
changing things that don’t affect the header. In this case it is important for performance
to not blindly replace an exisiting header widh an identical one.

The @update_header function will be called for each row after the call, and it will
continue to be called each time a row changes (via gtk_list_box_row_changed()) and when
the row before changes (either by gtk_list_box_row_changed() on the previous row, or when
the previous row becomes a different row). It is also called for all rows when
gtk_list_box_invalidate_headers() is called.
*/
func (self *_TraitListBox) SetHeaderFunc(update_header C.GtkListBoxUpdateHeaderFunc, user_data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_list_box_set_header_func(self.CPointer, update_header, (C.gpointer)(user_data), destroy)
	return
}

/*
Sets the placeholder widget that is shown in the list when
it doesn’t display any visible children.
*/
func (self *_TraitListBox) SetPlaceholder(placeholder *Widget) {
	C.gtk_list_box_set_placeholder(self.CPointer, (*C.GtkWidget)(placeholder.CPointer))
	return
}

/*
Sets how selection works in the listbox.
See #GtkSelectionMode for details.

Note: #GtkListBox does not support @GTK_SELECTION_MULTIPLE.
*/
func (self *_TraitListBox) SetSelectionMode(mode C.GtkSelectionMode) {
	C.gtk_list_box_set_selection_mode(self.CPointer, mode)
	return
}

/*
By setting a sort function on the @list_box one can dynamically reorder the rows
of the list, based on the contents of the rows.

The @sort_func will be called for each row after the call, and will continue to
be called each time a row changes (via gtk_list_box_row_changed()) and when
gtk_list_box_invalidate_sort() is called.
*/
func (self *_TraitListBox) SetSortFunc(sort_func C.GtkListBoxSortFunc, user_data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_list_box_set_sort_func(self.CPointer, sort_func, (C.gpointer)(user_data), destroy)
	return
}

type _TraitListBoxRow struct{ CPointer *C.GtkListBoxRow }

/*
Marks @row as changed, causing any state that depends on this
to be updated. This affects sorting, filtering and headers.

Note that calls to this method must be in sync with the data
used for the row functions. For instance, if the list is
mirroring some external data set, and *two* rows changed in the
external data set then when you call gtk_list_box_row_changed()
on the first row the sort function must only read the new data
for the first of the two changed rows, otherwise the resorting
of the rows will be wrong.

This generally means that if you don’t fully control the data
model you have to duplicate the data that affects the listbox
row functions into the row widgets themselves. Another alternative
is to call gtk_list_box_invalidate_sort() on any model change,
but that is more expensive.
*/
func (self *_TraitListBoxRow) Changed() {
	C.gtk_list_box_row_changed(self.CPointer)
	return
}

/*
Returns the current header of the @row. This can be used
in a #GtkListBoxUpdateHeaderFunc to see if there is a header
set already, and if so to update the state of it.
*/
func (self *_TraitListBoxRow) GetHeader() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_list_box_row_get_header(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the current index of the @row in its #GtkListBox container.
*/
func (self *_TraitListBoxRow) GetIndex() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_list_box_row_get_index(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Sets the current header of the @row. This is only allowed to be called
from a #GtkListBoxUpdateHeaderFunc. It will replace any existing
header in the row, and be shown in front of the row in the listbox.
*/
func (self *_TraitListBoxRow) SetHeader(header *Widget) {
	C.gtk_list_box_row_set_header(self.CPointer, (*C.GtkWidget)(header.CPointer))
	return
}

type _TraitListStore struct{ CPointer *C.GtkListStore }

/*
Appends a new row to @list_store.  @iter will be changed to point to this new
row.  The row will be empty after this function is called.  To fill in
values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
*/
func (self *_TraitListStore) Append() (iter C.GtkTreeIter) {
	C.gtk_list_store_append(self.CPointer, &iter)
	return
}

/*
Removes all rows from the list store.
*/
func (self *_TraitListStore) Clear() {
	C.gtk_list_store_clear(self.CPointer)
	return
}

/*
Creates a new row at @position.  @iter will be changed to point to this new
row.  If @position is -1 or is larger than the number of rows on the list,
then the new row will be appended to the list. The row will be empty after
this function is called.  To fill in values, you need to call
gtk_list_store_set() or gtk_list_store_set_value().
*/
func (self *_TraitListStore) Insert(position int) (iter C.GtkTreeIter) {
	C.gtk_list_store_insert(self.CPointer, &iter, C.gint(position))
	return
}

/*
Inserts a new row after @sibling. If @sibling is %NULL, then the row will be
prepended to the beginning of the list. @iter will be changed to point to
this new row. The row will be empty after this function is called. To fill
in values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
*/
func (self *_TraitListStore) InsertAfter(sibling *TreeIter) (iter C.GtkTreeIter) {
	C.gtk_list_store_insert_after(self.CPointer, &iter, (*C.GtkTreeIter)(unsafe.Pointer(sibling)))
	return
}

/*
Inserts a new row before @sibling. If @sibling is %NULL, then the row will
be appended to the end of the list. @iter will be changed to point to this
new row. The row will be empty after this function is called. To fill in
values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
*/
func (self *_TraitListStore) InsertBefore(sibling *TreeIter) (iter C.GtkTreeIter) {
	C.gtk_list_store_insert_before(self.CPointer, &iter, (*C.GtkTreeIter)(unsafe.Pointer(sibling)))
	return
}

// gtk_list_store_insert_with_values is not generated due to varargs

/*
A variant of gtk_list_store_insert_with_values() which
takes the columns and values as two arrays, instead of
varargs. This function is mainly intended for
language-bindings.
*/
func (self *_TraitListStore) InsertWithValuesv(position int, columns []int, values *C.GValue, n_values int) (iter C.GtkTreeIter) {
	__header__columns := (*reflect.SliceHeader)(unsafe.Pointer(&columns))
	C.gtk_list_store_insert_with_valuesv(self.CPointer, &iter, C.gint(position), (*C.gint)(unsafe.Pointer(__header__columns.Data)), values, C.gint(n_values))
	return
}

/*
> This function is slow. Only use it for debugging and/or testing
> purposes.

Checks if the given iter is a valid iter for this #GtkListStore.
*/
func (self *_TraitListStore) IterIsValid(iter *TreeIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_list_store_iter_is_valid(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Moves @iter in @store to the position after @position. Note that this
function only works with unsorted stores. If @position is %NULL, @iter
will be moved to the start of the list.
*/
func (self *_TraitListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	C.gtk_list_store_move_after(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), (*C.GtkTreeIter)(unsafe.Pointer(position)))
	return
}

/*
Moves @iter in @store to the position before @position. Note that this
function only works with unsorted stores. If @position is %NULL, @iter
will be moved to the end of the list.
*/
func (self *_TraitListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	C.gtk_list_store_move_before(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), (*C.GtkTreeIter)(unsafe.Pointer(position)))
	return
}

/*
Prepends a new row to @list_store. @iter will be changed to point to this new
row. The row will be empty after this function is called. To fill in
values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
*/
func (self *_TraitListStore) Prepend() (iter C.GtkTreeIter) {
	C.gtk_list_store_prepend(self.CPointer, &iter)
	return
}

/*
Removes the given row from the list store.  After being removed,
@iter is set to be the next valid row, or invalidated if it pointed
to the last row in @list_store.
*/
func (self *_TraitListStore) Remove(iter *TreeIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_list_store_remove(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Reorders @store to follow the order indicated by @new_order. Note that
this function only works with unsorted stores.
*/
func (self *_TraitListStore) Reorder(new_order *C.gint) {
	C.gtk_list_store_reorder(self.CPointer, new_order)
	return
}

// gtk_list_store_set is not generated due to varargs

/*
This function is meant primarily for #GObjects that inherit from #GtkListStore,
and should only be used when constructing a new #GtkListStore.  It will not
function after a row has been added, or a method on the #GtkTreeModel
interface is called.
*/
func (self *_TraitListStore) SetColumnTypes(n_columns int, types *C.GType) {
	C.gtk_list_store_set_column_types(self.CPointer, C.gint(n_columns), types)
	return
}

// gtk_list_store_set_valist is not generated due to varargs

/*
Sets the data in the cell specified by @iter and @column.
The type of @value must be convertible to the type of the
column.
*/
func (self *_TraitListStore) SetValue(iter *TreeIter, column int, value *C.GValue) {
	C.gtk_list_store_set_value(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), C.gint(column), value)
	return
}

/*
A variant of gtk_list_store_set_valist() which
takes the columns and values as two arrays, instead of
varargs. This function is mainly intended for
language-bindings and in case the number of columns to
change is not known until run-time.
*/
func (self *_TraitListStore) SetValuesv(iter *TreeIter, columns []int, values *C.GValue, n_values int) {
	__header__columns := (*reflect.SliceHeader)(unsafe.Pointer(&columns))
	C.gtk_list_store_set_valuesv(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), (*C.gint)(unsafe.Pointer(__header__columns.Data)), values, C.gint(n_values))
	return
}

/*
Swaps @a and @b in @store. Note that this function only works with
unsorted stores.
*/
func (self *_TraitListStore) Swap(a *TreeIter, b *TreeIter) {
	C.gtk_list_store_swap(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(a)), (*C.GtkTreeIter)(unsafe.Pointer(b)))
	return
}

type _TraitLockButton struct{ CPointer *C.GtkLockButton }

/*
Obtains the #GPermission object that controls @button.
*/
func (self *_TraitLockButton) GetPermission() (return__ *C.GPermission) {
	return__ = C.gtk_lock_button_get_permission(self.CPointer)
	return
}

/*
Sets the #GPermission object that controls @button.
*/
func (self *_TraitLockButton) SetPermission(permission *C.GPermission) {
	C.gtk_lock_button_set_permission(self.CPointer, permission)
	return
}

type _TraitMenu struct{ CPointer *C.GtkMenu }

/*
Adds a new #GtkMenuItem to a (table) menu. The number of “cells” that
an item will occupy is specified by @left_attach, @right_attach,
@top_attach and @bottom_attach. These each represent the leftmost,
rightmost, uppermost and lower column and row numbers of the table.
(Columns and rows are indexed from zero).

Note that this function is not related to gtk_menu_detach().
*/
func (self *_TraitMenu) Attach(child *Widget, left_attach uint, right_attach uint, top_attach uint, bottom_attach uint) {
	C.gtk_menu_attach(self.CPointer, (*C.GtkWidget)(child.CPointer), C.guint(left_attach), C.guint(right_attach), C.guint(top_attach), C.guint(bottom_attach))
	return
}

/*
Attaches the menu to the widget and provides a callback function
that will be invoked when the menu calls gtk_menu_detach() during
its destruction.

If the menu is attached to the widget then it will be destroyed
when the widget is destroyed, as if it was a child widget.
An attached menu will also move between screens correctly if the
widgets moves between screens.
*/
func (self *_TraitMenu) AttachToWidget(attach_widget *Widget, detacher C.GtkMenuDetachFunc) {
	C.gtk_menu_attach_to_widget(self.CPointer, (*C.GtkWidget)(attach_widget.CPointer), detacher)
	return
}

/*
Detaches the menu from the widget to which it had been attached.
This function will call the callback function, @detacher, provided
when the gtk_menu_attach_to_widget() function was called.
*/
func (self *_TraitMenu) Detach() {
	C.gtk_menu_detach(self.CPointer)
	return
}

/*
Gets the #GtkAccelGroup which holds global accelerators for the
menu. See gtk_menu_set_accel_group().
*/
func (self *_TraitMenu) GetAccelGroup() (return__ *AccelGroup) {
	var __cgo__return__ *C.GtkAccelGroup
	__cgo__return__ = C.gtk_menu_get_accel_group(self.CPointer)
	return__ = NewAccelGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the accelerator path set on the menu.
*/
func (self *_TraitMenu) GetAccelPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_menu_get_accel_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the selected menu item from the menu.  This is used by the
#GtkComboBox.
*/
func (self *_TraitMenu) GetActive() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_menu_get_active(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the #GtkWidget that the menu is attached to.
*/
func (self *_TraitMenu) GetAttachWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_menu_get_attach_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the number of the monitor on which to show the menu.
*/
func (self *_TraitMenu) GetMonitor() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_menu_get_monitor(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns whether the menu reserves space for toggles and
icons, regardless of their actual presence.
*/
func (self *_TraitMenu) GetReserveToggleSize() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_menu_get_reserve_toggle_size(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_menu_get_tearoff_state is not generated due to deprecation attr

// gtk_menu_get_title is not generated due to deprecation attr

/*
Removes the menu from the screen.
*/
func (self *_TraitMenu) Popdown() {
	C.gtk_menu_popdown(self.CPointer)
	return
}

/*
Displays a menu and makes it available for selection.

Applications can use this function to display context-sensitive
menus, and will typically supply %NULL for the @parent_menu_shell,
@parent_menu_item, @func and @data parameters. The default menu
positioning function will position the menu at the current mouse
cursor position.

The @button parameter should be the mouse button pressed to initiate
the menu popup. If the menu popup was initiated by something other
than a mouse button press, such as a mouse button release or a keypress,
@button should be 0.

The @activate_time parameter is used to conflict-resolve initiation
of concurrent requests for mouse/keyboard grab requests. To function
properly, this needs to be the timestamp of the user event (such as
a mouse click or key press) that caused the initiation of the popup.
Only if no such event is available, gtk_get_current_event_time() can
be used instead.
*/
func (self *_TraitMenu) Popup(parent_menu_shell *Widget, parent_menu_item *Widget, func_ C.GtkMenuPositionFunc, data unsafe.Pointer, button uint, activate_time uint32) {
	C.gtk_menu_popup(self.CPointer, (*C.GtkWidget)(parent_menu_shell.CPointer), (*C.GtkWidget)(parent_menu_item.CPointer), func_, (C.gpointer)(data), C.guint(button), C.guint32(activate_time))
	return
}

/*
Displays a menu and makes it available for selection.

Applications can use this function to display context-sensitive menus,
and will typically supply %NULL for the @parent_menu_shell,
@parent_menu_item, @func, @data and @destroy parameters. The default
menu positioning function will position the menu at the current position
of @device (or its corresponding pointer).

The @button parameter should be the mouse button pressed to initiate
the menu popup. If the menu popup was initiated by something other than
a mouse button press, such as a mouse button release or a keypress,
@button should be 0.

The @activate_time parameter is used to conflict-resolve initiation of
concurrent requests for mouse/keyboard grab requests. To function
properly, this needs to be the time stamp of the user event (such as
a mouse click or key press) that caused the initiation of the popup.
Only if no such event is available, gtk_get_current_event_time() can
be used instead.
*/
func (self *_TraitMenu) PopupForDevice(device *C.GdkDevice, parent_menu_shell *Widget, parent_menu_item *Widget, func_ C.GtkMenuPositionFunc, data unsafe.Pointer, destroy C.GDestroyNotify, button uint, activate_time uint32) {
	C.gtk_menu_popup_for_device(self.CPointer, device, (*C.GtkWidget)(parent_menu_shell.CPointer), (*C.GtkWidget)(parent_menu_item.CPointer), func_, (C.gpointer)(data), destroy, C.guint(button), C.guint32(activate_time))
	return
}

/*
Moves @child to a new @position in the list of @menu
children.
*/
func (self *_TraitMenu) ReorderChild(child *Widget, position int) {
	C.gtk_menu_reorder_child(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(position))
	return
}

/*
Repositions the menu according to its position function.
*/
func (self *_TraitMenu) Reposition() {
	C.gtk_menu_reposition(self.CPointer)
	return
}

/*
Set the #GtkAccelGroup which holds global accelerators for the
menu.  This accelerator group needs to also be added to all windows
that this menu is being used in with gtk_window_add_accel_group(),
in order for those windows to support all the accelerators
contained in this group.
*/
func (self *_TraitMenu) SetAccelGroup(accel_group *AccelGroup) {
	C.gtk_menu_set_accel_group(self.CPointer, (*C.GtkAccelGroup)(accel_group.CPointer))
	return
}

/*
Sets an accelerator path for this menu from which accelerator paths
for its immediate children, its menu items, can be constructed.
The main purpose of this function is to spare the programmer the
inconvenience of having to call gtk_menu_item_set_accel_path() on
each menu item that should support runtime user changable accelerators.
Instead, by just calling gtk_menu_set_accel_path() on their parent,
each menu item of this menu, that contains a label describing its
purpose, automatically gets an accel path assigned.

For example, a menu containing menu items “New” and “Exit”, will, after
`gtk_menu_set_accel_path (menu, "<Gnumeric-Sheet>/File");` has been
called, assign its items the accel paths: `"<Gnumeric-Sheet>/File/New"`
and `"<Gnumeric-Sheet>/File/Exit"`.

Assigning accel paths to menu items then enables the user to change
their accelerators at runtime. More details about accelerator paths
and their default setups can be found at gtk_accel_map_add_entry().

Note that @accel_path string will be stored in a #GQuark. Therefore,
if you pass a static string, you can save some memory by interning
it first with g_intern_static_string().
*/
func (self *_TraitMenu) SetAccelPath(accel_path string) {
	__cgo__accel_path := (*C.gchar)(unsafe.Pointer(C.CString(accel_path)))
	C.gtk_menu_set_accel_path(self.CPointer, __cgo__accel_path)
	C.free(unsafe.Pointer(__cgo__accel_path))
	return
}

/*
Selects the specified menu item within the menu.  This is used by
the #GtkComboBox and should not be used by anyone else.
*/
func (self *_TraitMenu) SetActive(index uint) {
	C.gtk_menu_set_active(self.CPointer, C.guint(index))
	return
}

/*
Informs GTK+ on which monitor a menu should be popped up.
See gdk_screen_get_monitor_geometry().

This function should be called from a #GtkMenuPositionFunc
if the menu should not appear on the same monitor as the pointer.
This information can’t be reliably inferred from the coordinates
returned by a #GtkMenuPositionFunc, since, for very long menus,
these coordinates may extend beyond the monitor boundaries or even
the screen boundaries.
*/
func (self *_TraitMenu) SetMonitor(monitor_num int) {
	C.gtk_menu_set_monitor(self.CPointer, C.gint(monitor_num))
	return
}

/*
Sets whether the menu should reserve space for drawing toggles
or icons, regardless of their actual presence.
*/
func (self *_TraitMenu) SetReserveToggleSize(reserve_toggle_size bool) {
	__cgo__reserve_toggle_size := C.gboolean(0)
	if reserve_toggle_size {
		__cgo__reserve_toggle_size = C.gboolean(1)
	}
	C.gtk_menu_set_reserve_toggle_size(self.CPointer, __cgo__reserve_toggle_size)
	return
}

/*
Sets the #GdkScreen on which the menu will be displayed.
*/
func (self *_TraitMenu) SetScreen(screen *C.GdkScreen) {
	C.gtk_menu_set_screen(self.CPointer, screen)
	return
}

// gtk_menu_set_tearoff_state is not generated due to deprecation attr

// gtk_menu_set_title is not generated due to deprecation attr

type _TraitMenuBar struct{ CPointer *C.GtkMenuBar }

/*
Retrieves the current child pack direction of the menubar.
See gtk_menu_bar_set_child_pack_direction().
*/
func (self *_TraitMenuBar) GetChildPackDirection() (return__ C.GtkPackDirection) {
	return__ = C.gtk_menu_bar_get_child_pack_direction(self.CPointer)
	return
}

/*
Retrieves the current pack direction of the menubar.
See gtk_menu_bar_set_pack_direction().
*/
func (self *_TraitMenuBar) GetPackDirection() (return__ C.GtkPackDirection) {
	return__ = C.gtk_menu_bar_get_pack_direction(self.CPointer)
	return
}

/*
Sets how widgets should be packed inside the children of a menubar.
*/
func (self *_TraitMenuBar) SetChildPackDirection(child_pack_dir C.GtkPackDirection) {
	C.gtk_menu_bar_set_child_pack_direction(self.CPointer, child_pack_dir)
	return
}

/*
Sets how items should be packed inside a menubar.
*/
func (self *_TraitMenuBar) SetPackDirection(pack_dir C.GtkPackDirection) {
	C.gtk_menu_bar_set_pack_direction(self.CPointer, pack_dir)
	return
}

type _TraitMenuButton struct{ CPointer *C.GtkMenuButton }

/*
Returns the parent #GtkWidget to use to line up with menu.
*/
func (self *_TraitMenuButton) GetAlignWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_menu_button_get_align_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the direction the popup will be pointing at when popped up.
*/
func (self *_TraitMenuButton) GetDirection() (return__ C.GtkArrowType) {
	return__ = C.gtk_menu_button_get_direction(self.CPointer)
	return
}

/*
Returns the #GMenuModel used to generate the popup.
*/
func (self *_TraitMenuButton) GetMenuModel() (return__ *C.GMenuModel) {
	return__ = C.gtk_menu_button_get_menu_model(self.CPointer)
	return
}

/*
Returns the #GtkPopover that pops out of the button.
If the button is not using a #GtkPopover, this function
returns %NULL.
*/
func (self *_TraitMenuButton) GetPopover() (return__ *Popover) {
	var __cgo__return__ *C.GtkPopover
	__cgo__return__ = C.gtk_menu_button_get_popover(self.CPointer)
	return__ = NewPopoverFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the #GtkMenu that pops out of the button.
If the button does not use a #GtkMenu, this function
returns %NULL.
*/
func (self *_TraitMenuButton) GetPopup() (return__ *Menu) {
	var __cgo__return__ *C.GtkMenu
	__cgo__return__ = C.gtk_menu_button_get_popup(self.CPointer)
	return__ = NewMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether a #GtkPopover or a #GtkMenu will be constructed
from the menu model.
*/
func (self *_TraitMenuButton) GetUsePopover() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_menu_button_get_use_popover(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the #GtkWidget to use to line the menu with when popped up.
Note that the @align_widget must contain the #GtkMenuButton itself.

Setting it to %NULL means that the menu will be aligned with the
button itself.

Note that this property is only used with menus currently,
and not for popovers.
*/
func (self *_TraitMenuButton) SetAlignWidget(align_widget *Widget) {
	C.gtk_menu_button_set_align_widget(self.CPointer, (*C.GtkWidget)(align_widget.CPointer))
	return
}

/*
Sets the direction in which the popup will be popped up, as
well as changing the arrow’s direction. The child will not
be changed to an arrow if it was customized.

If the does not fit in the available space in the given direction,
GTK+ will its best to keep it inside the screen and fully visible.

If you pass %GTK_ARROW_NONE for a @direction, the popup will behave
as if you passed %GTK_ARROW_DOWN (although you won’t see any arrows).
*/
func (self *_TraitMenuButton) SetDirection(direction C.GtkArrowType) {
	C.gtk_menu_button_set_direction(self.CPointer, direction)
	return
}

/*
Sets the #GMenuModel from which the popup will be constructed,
or %NULL to disable the button.

Depending on the value of #GtkMenuButton:use-popover, either a
#GtkMenu will be created with gtk_menu_new_from_model(), or a
#GtkPopover with gtk_popover_new_from_model(). In either case,
actions will be connected as documented for these functions.

If #GtkMenuButton:popup or #GtkMenuButton:popover are already set,
their content will be lost and replaced by the newly created popup.
*/
func (self *_TraitMenuButton) SetMenuModel(menu_model *C.GMenuModel) {
	C.gtk_menu_button_set_menu_model(self.CPointer, menu_model)
	return
}

/*
Sets the #GtkPopover that will be popped up when the button is
clicked, or %NULL to disable the button. If #GtkMenuButton:menu-model
or #GtkMenuButton:popup are set, they will be set to %NULL.
*/
func (self *_TraitMenuButton) SetPopover(popover *Widget) {
	C.gtk_menu_button_set_popover(self.CPointer, (*C.GtkWidget)(popover.CPointer))
	return
}

/*
Sets the #GtkMenu that will be popped up when the button is clicked,
or %NULL to disable the button. If #GtkMenuButton:menu-model or
#GtkMenuButton:popover are set, they will be set to %NULL.
*/
func (self *_TraitMenuButton) SetPopup(menu *Widget) {
	C.gtk_menu_button_set_popup(self.CPointer, (*C.GtkWidget)(menu.CPointer))
	return
}

/*
Sets whether to construct a #GtkPopover instead of #GtkMenu
when gtk_menu_button_set_menu_model() is called. Note that
this property is only consulted when a new menu model is set.
*/
func (self *_TraitMenuButton) SetUsePopover(use_popover bool) {
	__cgo__use_popover := C.gboolean(0)
	if use_popover {
		__cgo__use_popover = C.gboolean(1)
	}
	C.gtk_menu_button_set_use_popover(self.CPointer, __cgo__use_popover)
	return
}

type _TraitMenuItem struct{ CPointer *C.GtkMenuItem }

/*
Emits the #GtkMenuItem::activate signal on the given item
*/
func (self *_TraitMenuItem) Activate() {
	C.gtk_menu_item_activate(self.CPointer)
	return
}

/*
Emits the #GtkMenuItem::deselect signal on the given item.
*/
func (self *_TraitMenuItem) Deselect() {
	C.gtk_menu_item_deselect(self.CPointer)
	return
}

/*
Retrieve the accelerator path that was previously set on @menu_item.

See gtk_menu_item_set_accel_path() for details.
*/
func (self *_TraitMenuItem) GetAccelPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_menu_item_get_accel_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Sets @text on the @menu_item label
*/
func (self *_TraitMenuItem) GetLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_menu_item_get_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns whether the @menu_item reserves space for
the submenu indicator, regardless if it has a submenu
or not.
*/
func (self *_TraitMenuItem) GetReserveIndicator() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_menu_item_get_reserve_indicator(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_menu_item_get_right_justified is not generated due to deprecation attr

/*
Gets the submenu underneath this menu item, if any.
See gtk_menu_item_set_submenu().
*/
func (self *_TraitMenuItem) GetSubmenu() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_menu_item_get_submenu(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Checks if an underline in the text indicates the next character
should be used for the mnemonic accelerator key.
*/
func (self *_TraitMenuItem) GetUseUnderline() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_menu_item_get_use_underline(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Emits the #GtkMenuItem::select signal on the given item.
*/
func (self *_TraitMenuItem) Select() {
	C.gtk_menu_item_select(self.CPointer)
	return
}

/*
Set the accelerator path on @menu_item, through which runtime
changes of the menu item’s accelerator caused by the user can be
identified and saved to persistent storage (see gtk_accel_map_save()
on this). To set up a default accelerator for this menu item, call
gtk_accel_map_add_entry() with the same @accel_path. See also
gtk_accel_map_add_entry() on the specifics of accelerator paths,
and gtk_menu_set_accel_path() for a more convenient variant of
this function.

This function is basically a convenience wrapper that handles
calling gtk_widget_set_accel_path() with the appropriate accelerator
group for the menu item.

Note that you do need to set an accelerator on the parent menu with
gtk_menu_set_accel_group() for this to work.

Note that @accel_path string will be stored in a #GQuark.
Therefore, if you pass a static string, you can save some memory
by interning it first with g_intern_static_string().
*/
func (self *_TraitMenuItem) SetAccelPath(accel_path string) {
	__cgo__accel_path := (*C.gchar)(unsafe.Pointer(C.CString(accel_path)))
	C.gtk_menu_item_set_accel_path(self.CPointer, __cgo__accel_path)
	C.free(unsafe.Pointer(__cgo__accel_path))
	return
}

/*
Sets @text on the @menu_item label
*/
func (self *_TraitMenuItem) SetLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_menu_item_set_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Sets whether the @menu_item should reserve space for
the submenu indicator, regardless if it actually has
a submenu or not.

There should be little need for applications to call
this functions.
*/
func (self *_TraitMenuItem) SetReserveIndicator(reserve bool) {
	__cgo__reserve := C.gboolean(0)
	if reserve {
		__cgo__reserve = C.gboolean(1)
	}
	C.gtk_menu_item_set_reserve_indicator(self.CPointer, __cgo__reserve)
	return
}

// gtk_menu_item_set_right_justified is not generated due to deprecation attr

/*
Sets or replaces the menu item’s submenu, or removes it when a %NULL
submenu is passed.
*/
func (self *_TraitMenuItem) SetSubmenu(submenu *Widget) {
	C.gtk_menu_item_set_submenu(self.CPointer, (*C.GtkWidget)(submenu.CPointer))
	return
}

/*
If true, an underline in the text indicates the next character
should be used for the mnemonic accelerator key.
*/
func (self *_TraitMenuItem) SetUseUnderline(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_menu_item_set_use_underline(self.CPointer, __cgo__setting)
	return
}

/*
Emits the #GtkMenuItem::toggle-size-allocate signal on the given item.
*/
func (self *_TraitMenuItem) ToggleSizeAllocate(allocation int) {
	C.gtk_menu_item_toggle_size_allocate(self.CPointer, C.gint(allocation))
	return
}

// gtk_menu_item_toggle_size_request is not generated due to inout param

type _TraitMenuShell struct{ CPointer *C.GtkMenuShell }

/*
Activates the menu item within the menu shell.
*/
func (self *_TraitMenuShell) ActivateItem(menu_item *Widget, force_deactivate bool) {
	__cgo__force_deactivate := C.gboolean(0)
	if force_deactivate {
		__cgo__force_deactivate = C.gboolean(1)
	}
	C.gtk_menu_shell_activate_item(self.CPointer, (*C.GtkWidget)(menu_item.CPointer), __cgo__force_deactivate)
	return
}

/*
Adds a new #GtkMenuItem to the end of the menu shell's
item list.
*/
func (self *_TraitMenuShell) Append(child *Widget) {
	C.gtk_menu_shell_append(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Establishes a binding between a #GtkMenuShell and a #GMenuModel.

The contents of @shell are removed and then refilled with menu items
according to @model.  When @model changes, @shell is updated.
Calling this function twice on @shell with different @model will
cause the first binding to be replaced with a binding to the new
model. If @model is %NULL then any previous binding is undone and
all children are removed.

@with_separators determines if toplevel items (eg: sections) have
separators inserted between them.  This is typically desired for
menus but doesn’t make sense for menubars.

If @action_namespace is non-%NULL then the effect is as if all
actions mentioned in the @model have their names prefixed with the
namespace, plus a dot.  For example, if the action “quit” is
mentioned and @action_namespace is “app” then the effective action
name is “app.quit”.

This function uses #GtkActionable to define the action name and
target values on the created menu items.  If you want to use an
action group other than “app” and “win”, or if you want to use a
#GtkMenuShell outside of a #GtkApplicationWindow, then you will need
to attach your own action group to the widget hierarchy using
gtk_widget_insert_action_group().  As an example, if you created a
group with a “quit” action and inserted it with the name “mygroup”
then you would use the action name “mygroup.quit” in your
#GMenuModel.

For most cases you are probably better off using
gtk_menu_new_from_model() or gtk_menu_bar_new_from_model() or just
directly passing the #GMenuModel to gtk_application_set_app_menu() or
gtk_application_set_menubar().
*/
func (self *_TraitMenuShell) BindModel(model *C.GMenuModel, action_namespace string, with_separators bool) {
	__cgo__action_namespace := (*C.gchar)(unsafe.Pointer(C.CString(action_namespace)))
	__cgo__with_separators := C.gboolean(0)
	if with_separators {
		__cgo__with_separators = C.gboolean(1)
	}
	C.gtk_menu_shell_bind_model(self.CPointer, model, __cgo__action_namespace, __cgo__with_separators)
	C.free(unsafe.Pointer(__cgo__action_namespace))
	return
}

/*
Cancels the selection within the menu shell.
*/
func (self *_TraitMenuShell) Cancel() {
	C.gtk_menu_shell_cancel(self.CPointer)
	return
}

/*
Deactivates the menu shell.

Typically this results in the menu shell being erased
from the screen.
*/
func (self *_TraitMenuShell) Deactivate() {
	C.gtk_menu_shell_deactivate(self.CPointer)
	return
}

/*
Deselects the currently selected item from the menu shell,
if any.
*/
func (self *_TraitMenuShell) Deselect() {
	C.gtk_menu_shell_deselect(self.CPointer)
	return
}

/*
Gets the parent menu shell.

The parent menu shell of a submenu is the #GtkMenu or #GtkMenuBar
from which it was opened up.
*/
func (self *_TraitMenuShell) GetParentShell() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_menu_shell_get_parent_shell(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the currently selected item.
*/
func (self *_TraitMenuShell) GetSelectedItem() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_menu_shell_get_selected_item(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns %TRUE if the menu shell will take the keyboard focus on popup.
*/
func (self *_TraitMenuShell) GetTakeFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_menu_shell_get_take_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Adds a new #GtkMenuItem to the menu shell’s item list
at the position indicated by @position.
*/
func (self *_TraitMenuShell) Insert(child *Widget, position int) {
	C.gtk_menu_shell_insert(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(position))
	return
}

/*
Adds a new #GtkMenuItem to the beginning of the menu shell's
item list.
*/
func (self *_TraitMenuShell) Prepend(child *Widget) {
	C.gtk_menu_shell_prepend(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Select the first visible or selectable child of the menu shell;
don’t select tearoff items unless the only item is a tearoff
item.
*/
func (self *_TraitMenuShell) SelectFirst(search_sensitive bool) {
	__cgo__search_sensitive := C.gboolean(0)
	if search_sensitive {
		__cgo__search_sensitive = C.gboolean(1)
	}
	C.gtk_menu_shell_select_first(self.CPointer, __cgo__search_sensitive)
	return
}

/*
Selects the menu item from the menu shell.
*/
func (self *_TraitMenuShell) SelectItem(menu_item *Widget) {
	C.gtk_menu_shell_select_item(self.CPointer, (*C.GtkWidget)(menu_item.CPointer))
	return
}

/*
If @take_focus is %TRUE (the default) the menu shell will take
the keyboard focus so that it will receive all keyboard events
which is needed to enable keyboard navigation in menus.

Setting @take_focus to %FALSE is useful only for special applications
like virtual keyboard implementations which should not take keyboard
focus.

The @take_focus state of a menu or menu bar is automatically
propagated to submenus whenever a submenu is popped up, so you
don’t have to worry about recursively setting it for your entire
menu hierarchy. Only when programmatically picking a submenu and
popping it up manually, the @take_focus property of the submenu
needs to be set explicitly.

Note that setting it to %FALSE has side-effects:

If the focus is in some other app, it keeps the focus and keynav in
the menu doesn’t work. Consequently, keynav on the menu will only
work if the focus is on some toplevel owned by the onscreen keyboard.

To avoid confusing the user, menus with @take_focus set to %FALSE
should not display mnemonics or accelerators, since it cannot be
guaranteed that they will work.

See also gdk_keyboard_grab()
*/
func (self *_TraitMenuShell) SetTakeFocus(take_focus bool) {
	__cgo__take_focus := C.gboolean(0)
	if take_focus {
		__cgo__take_focus = C.gboolean(1)
	}
	C.gtk_menu_shell_set_take_focus(self.CPointer, __cgo__take_focus)
	return
}

type _TraitMenuToolButton struct{ CPointer *C.GtkMenuToolButton }

/*
Gets the #GtkMenu associated with #GtkMenuToolButton.
*/
func (self *_TraitMenuToolButton) GetMenu() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_menu_tool_button_get_menu(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Sets the tooltip markup text to be used as tooltip for the arrow button
which pops up the menu.  See gtk_tool_item_set_tooltip_text() for setting
a tooltip on the whole #GtkMenuToolButton.
*/
func (self *_TraitMenuToolButton) SetArrowTooltipMarkup(markup string) {
	__cgo__markup := (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	C.gtk_menu_tool_button_set_arrow_tooltip_markup(self.CPointer, __cgo__markup)
	C.free(unsafe.Pointer(__cgo__markup))
	return
}

/*
Sets the tooltip text to be used as tooltip for the arrow button which
pops up the menu.  See gtk_tool_item_set_tooltip_text() for setting a tooltip
on the whole #GtkMenuToolButton.
*/
func (self *_TraitMenuToolButton) SetArrowTooltipText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_menu_tool_button_set_arrow_tooltip_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Sets the #GtkMenu that is popped up when the user clicks on the arrow.
If @menu is NULL, the arrow button becomes insensitive.
*/
func (self *_TraitMenuToolButton) SetMenu(menu *Widget) {
	C.gtk_menu_tool_button_set_menu(self.CPointer, (*C.GtkWidget)(menu.CPointer))
	return
}

type _TraitMessageDialog struct{ CPointer *C.GtkMessageDialog }

// gtk_message_dialog_format_secondary_markup is not generated due to varargs

// gtk_message_dialog_format_secondary_text is not generated due to varargs

// gtk_message_dialog_get_image is not generated due to deprecation attr

/*
Returns the message area of the dialog. This is the box where the
dialog’s primary and secondary labels are packed. You can add your
own extra content to that box and it will appear below those labels.
See gtk_dialog_get_content_area() for the corresponding
function in the parent #GtkDialog.
*/
func (self *_TraitMessageDialog) GetMessageArea() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_message_dialog_get_message_area(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

// gtk_message_dialog_set_image is not generated due to deprecation attr

/*
Sets the text of the message dialog to be @str, which is marked
up with the [Pango text markup language][PangoMarkupFormat].
*/
func (self *_TraitMessageDialog) SetMarkup(str string) {
	__cgo__str := (*C.gchar)(unsafe.Pointer(C.CString(str)))
	C.gtk_message_dialog_set_markup(self.CPointer, __cgo__str)
	C.free(unsafe.Pointer(__cgo__str))
	return
}

type _TraitMisc struct{ CPointer *C.GtkMisc }

/*
Gets the X and Y alignment of the widget within its allocation.
See gtk_misc_set_alignment().
*/
func (self *_TraitMisc) GetAlignment() (xalign float32, yalign float32) {
	var __cgo__xalign C.gfloat
	var __cgo__yalign C.gfloat
	C.gtk_misc_get_alignment(self.CPointer, &__cgo__xalign, &__cgo__yalign)
	xalign = float32(__cgo__xalign)
	yalign = float32(__cgo__yalign)
	return
}

/*
Gets the padding in the X and Y directions of the widget.
See gtk_misc_set_padding().
*/
func (self *_TraitMisc) GetPadding() (xpad int, ypad int) {
	var __cgo__xpad C.gint
	var __cgo__ypad C.gint
	C.gtk_misc_get_padding(self.CPointer, &__cgo__xpad, &__cgo__ypad)
	xpad = int(__cgo__xpad)
	ypad = int(__cgo__ypad)
	return
}

/*
Sets the alignment of the widget.
*/
func (self *_TraitMisc) SetAlignment(xalign float32, yalign float32) {
	C.gtk_misc_set_alignment(self.CPointer, C.gfloat(xalign), C.gfloat(yalign))
	return
}

/*
Sets the amount of space to add around the widget.
*/
func (self *_TraitMisc) SetPadding(xpad int, ypad int) {
	C.gtk_misc_set_padding(self.CPointer, C.gint(xpad), C.gint(ypad))
	return
}

type _TraitMountOperation struct{ CPointer *C.GtkMountOperation }

/*
Gets the transient parent used by the #GtkMountOperation
*/
func (self *_TraitMountOperation) GetParent() (return__ *Window) {
	var __cgo__return__ *C.GtkWindow
	__cgo__return__ = C.gtk_mount_operation_get_parent(self.CPointer)
	return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the screen on which windows of the #GtkMountOperation
will be shown.
*/
func (self *_TraitMountOperation) GetScreen() (return__ *C.GdkScreen) {
	return__ = C.gtk_mount_operation_get_screen(self.CPointer)
	return
}

/*
Returns whether the #GtkMountOperation is currently displaying
a window.
*/
func (self *_TraitMountOperation) IsShowing() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_mount_operation_is_showing(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the transient parent for windows shown by the
#GtkMountOperation.
*/
func (self *_TraitMountOperation) SetParent(parent *Window) {
	C.gtk_mount_operation_set_parent(self.CPointer, (*C.GtkWindow)(parent.CPointer))
	return
}

/*
Sets the screen to show windows of the #GtkMountOperation on.
*/
func (self *_TraitMountOperation) SetScreen(screen *C.GdkScreen) {
	C.gtk_mount_operation_set_screen(self.CPointer, screen)
	return
}

type _TraitNotebook struct{ CPointer *C.GtkNotebook }

/*
Appends a page to @notebook.
*/
func (self *_TraitNotebook) AppendPage(child *Widget, tab_label *Widget) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_append_page(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(tab_label.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Appends a page to @notebook, specifying the widget to use as the
label in the popup menu.
*/
func (self *_TraitNotebook) AppendPageMenu(child *Widget, tab_label *Widget, menu_label *Widget) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_append_page_menu(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(tab_label.CPointer), (*C.GtkWidget)(menu_label.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Gets one of the action widgets. See gtk_notebook_set_action_widget().
*/
func (self *_TraitNotebook) GetActionWidget(pack_type C.GtkPackType) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_notebook_get_action_widget(self.CPointer, pack_type)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the page number of the current page.
*/
func (self *_TraitNotebook) GetCurrentPage() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_get_current_page(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the current group name for @notebook.
*/
func (self *_TraitNotebook) GetGroupName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_notebook_get_group_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the menu label widget of the page containing @child.
*/
func (self *_TraitNotebook) GetMenuLabel(child *Widget) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_notebook_get_menu_label(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the text of the menu label for the page containing
@child.
*/
func (self *_TraitNotebook) GetMenuLabelText(child *Widget) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_notebook_get_menu_label_text(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the number of pages in a notebook.
*/
func (self *_TraitNotebook) GetNPages() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_get_n_pages(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the child widget contained in page number @page_num.
*/
func (self *_TraitNotebook) GetNthPage(page_num int) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_notebook_get_nth_page(self.CPointer, C.gint(page_num))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether the tab label area has arrows for scrolling.
See gtk_notebook_set_scrollable().
*/
func (self *_TraitNotebook) GetScrollable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_notebook_get_scrollable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether a bevel will be drawn around the notebook pages.
See gtk_notebook_set_show_border().
*/
func (self *_TraitNotebook) GetShowBorder() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_notebook_get_show_border(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the tabs of the notebook are shown.
See gtk_notebook_set_show_tabs().
*/
func (self *_TraitNotebook) GetShowTabs() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_notebook_get_show_tabs(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the tab contents can be detached from @notebook.
*/
func (self *_TraitNotebook) GetTabDetachable(child *Widget) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_notebook_get_tab_detachable(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_notebook_get_tab_hborder is not generated due to deprecation attr

/*
Returns the tab label widget for the page @child.
%NULL is returned if @child is not in @notebook or
if no tab label has specifically been set for @child.
*/
func (self *_TraitNotebook) GetTabLabel(child *Widget) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_notebook_get_tab_label(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the text of the tab label for the page containing
@child.
*/
func (self *_TraitNotebook) GetTabLabelText(child *Widget) (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_notebook_get_tab_label_text(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the edge at which the tabs for switching pages in the
notebook are drawn.
*/
func (self *_TraitNotebook) GetTabPos() (return__ C.GtkPositionType) {
	return__ = C.gtk_notebook_get_tab_pos(self.CPointer)
	return
}

/*
Gets whether the tab can be reordered via drag and drop or not.
*/
func (self *_TraitNotebook) GetTabReorderable(child *Widget) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_notebook_get_tab_reorderable(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_notebook_get_tab_vborder is not generated due to deprecation attr

/*
Insert a page into @notebook at the given position.
*/
func (self *_TraitNotebook) InsertPage(child *Widget, tab_label *Widget, position int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_insert_page(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(tab_label.CPointer), C.gint(position))
	return__ = int(__cgo__return__)
	return
}

/*
Insert a page into @notebook at the given position, specifying
the widget to use as the label in the popup menu.
*/
func (self *_TraitNotebook) InsertPageMenu(child *Widget, tab_label *Widget, menu_label *Widget, position int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_insert_page_menu(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(tab_label.CPointer), (*C.GtkWidget)(menu_label.CPointer), C.gint(position))
	return__ = int(__cgo__return__)
	return
}

/*
Switches to the next page. Nothing happens if the current page is
the last page.
*/
func (self *_TraitNotebook) NextPage() {
	C.gtk_notebook_next_page(self.CPointer)
	return
}

/*
Finds the index of the page which contains the given child
widget.
*/
func (self *_TraitNotebook) PageNum(child *Widget) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_page_num(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Disables the popup menu.
*/
func (self *_TraitNotebook) PopupDisable() {
	C.gtk_notebook_popup_disable(self.CPointer)
	return
}

/*
Enables the popup menu: if the user clicks with the right
mouse button on the tab labels, a menu with all the pages
will be popped up.
*/
func (self *_TraitNotebook) PopupEnable() {
	C.gtk_notebook_popup_enable(self.CPointer)
	return
}

/*
Prepends a page to @notebook.
*/
func (self *_TraitNotebook) PrependPage(child *Widget, tab_label *Widget) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_prepend_page(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(tab_label.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Prepends a page to @notebook, specifying the widget to use as the
label in the popup menu.
*/
func (self *_TraitNotebook) PrependPageMenu(child *Widget, tab_label *Widget, menu_label *Widget) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_notebook_prepend_page_menu(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(tab_label.CPointer), (*C.GtkWidget)(menu_label.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Switches to the previous page. Nothing happens if the current page
is the first page.
*/
func (self *_TraitNotebook) PrevPage() {
	C.gtk_notebook_prev_page(self.CPointer)
	return
}

/*
Removes a page from the notebook given its index
in the notebook.
*/
func (self *_TraitNotebook) RemovePage(page_num int) {
	C.gtk_notebook_remove_page(self.CPointer, C.gint(page_num))
	return
}

/*
Reorders the page containing @child, so that it appears in position
@position. If @position is greater than or equal to the number of
children in the list or negative, @child will be moved to the end
of the list.
*/
func (self *_TraitNotebook) ReorderChild(child *Widget, position int) {
	C.gtk_notebook_reorder_child(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(position))
	return
}

/*
Sets @widget as one of the action widgets. Depending on the pack type
the widget will be placed before or after the tabs. You can use
a #GtkBox if you need to pack more than one widget on the same side.

Note that action widgets are “internal” children of the notebook and thus
not included in the list returned from gtk_container_foreach().
*/
func (self *_TraitNotebook) SetActionWidget(widget *Widget, pack_type C.GtkPackType) {
	C.gtk_notebook_set_action_widget(self.CPointer, (*C.GtkWidget)(widget.CPointer), pack_type)
	return
}

/*
Switches to the page number @page_num.

Note that due to historical reasons, GtkNotebook refuses
to switch to a page unless the child widget is visible.
Therefore, it is recommended to show child widgets before
adding them to a notebook.
*/
func (self *_TraitNotebook) SetCurrentPage(page_num int) {
	C.gtk_notebook_set_current_page(self.CPointer, C.gint(page_num))
	return
}

/*
Sets a group name for @notebook.

Notebooks with the same name will be able to exchange tabs
via drag and drop. A notebook with a %NULL group name will
not be able to exchange tabs with any other notebook.
*/
func (self *_TraitNotebook) SetGroupName(group_name string) {
	__cgo__group_name := (*C.gchar)(unsafe.Pointer(C.CString(group_name)))
	C.gtk_notebook_set_group_name(self.CPointer, __cgo__group_name)
	C.free(unsafe.Pointer(__cgo__group_name))
	return
}

/*
Changes the menu label for the page containing @child.
*/
func (self *_TraitNotebook) SetMenuLabel(child *Widget, menu_label *Widget) {
	C.gtk_notebook_set_menu_label(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(menu_label.CPointer))
	return
}

/*
Creates a new label and sets it as the menu label of @child.
*/
func (self *_TraitNotebook) SetMenuLabelText(child *Widget, menu_text string) {
	__cgo__menu_text := (*C.gchar)(unsafe.Pointer(C.CString(menu_text)))
	C.gtk_notebook_set_menu_label_text(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__menu_text)
	C.free(unsafe.Pointer(__cgo__menu_text))
	return
}

/*
Sets whether the tab label area will have arrows for
scrolling if there are too many tabs to fit in the area.
*/
func (self *_TraitNotebook) SetScrollable(scrollable bool) {
	__cgo__scrollable := C.gboolean(0)
	if scrollable {
		__cgo__scrollable = C.gboolean(1)
	}
	C.gtk_notebook_set_scrollable(self.CPointer, __cgo__scrollable)
	return
}

/*
Sets whether a bevel will be drawn around the notebook pages.
This only has a visual effect when the tabs are not shown.
See gtk_notebook_set_show_tabs().
*/
func (self *_TraitNotebook) SetShowBorder(show_border bool) {
	__cgo__show_border := C.gboolean(0)
	if show_border {
		__cgo__show_border = C.gboolean(1)
	}
	C.gtk_notebook_set_show_border(self.CPointer, __cgo__show_border)
	return
}

/*
Sets whether to show the tabs for the notebook or not.
*/
func (self *_TraitNotebook) SetShowTabs(show_tabs bool) {
	__cgo__show_tabs := C.gboolean(0)
	if show_tabs {
		__cgo__show_tabs = C.gboolean(1)
	}
	C.gtk_notebook_set_show_tabs(self.CPointer, __cgo__show_tabs)
	return
}

/*
Sets whether the tab can be detached from @notebook to another
notebook or widget.

Note that 2 notebooks must share a common group identificator
(see gtk_notebook_set_group_name()) to allow automatic tabs
interchange between them.

If you want a widget to interact with a notebook through DnD
(i.e.: accept dragged tabs from it) it must be set as a drop
destination and accept the target “GTK_NOTEBOOK_TAB”. The notebook
will fill the selection with a GtkWidget** pointing to the child
widget that corresponds to the dropped tab.
|[<!-- language="C" -->
 static void
 on_drag_data_received (GtkWidget        *widget,
                        GdkDragContext   *context,
                        gint              x,
                        gint              y,
                        GtkSelectionData *data,
                        guint             info,
                        guint             time,
                        gpointer          user_data)
 {
   GtkWidget *notebook;
   GtkWidget **child;
   GtkContainer *container;

   notebook = gtk_drag_get_source_widget (context);
   child = (void*) gtk_selection_data_get_data (data);

   process_widget (*child);
   container = GTK_CONTAINER (notebook);
   gtk_container_remove (container, *child);
 }
]|

If you want a notebook to accept drags from other widgets,
you will have to set your own DnD code to do it.
*/
func (self *_TraitNotebook) SetTabDetachable(child *Widget, detachable bool) {
	__cgo__detachable := C.gboolean(0)
	if detachable {
		__cgo__detachable = C.gboolean(1)
	}
	C.gtk_notebook_set_tab_detachable(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__detachable)
	return
}

/*
Changes the tab label for @child.
If %NULL is specified for @tab_label, then the page will
have the label “page N”.
*/
func (self *_TraitNotebook) SetTabLabel(child *Widget, tab_label *Widget) {
	C.gtk_notebook_set_tab_label(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkWidget)(tab_label.CPointer))
	return
}

/*
Creates a new label and sets it as the tab label for the page
containing @child.
*/
func (self *_TraitNotebook) SetTabLabelText(child *Widget, tab_text string) {
	__cgo__tab_text := (*C.gchar)(unsafe.Pointer(C.CString(tab_text)))
	C.gtk_notebook_set_tab_label_text(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__tab_text)
	C.free(unsafe.Pointer(__cgo__tab_text))
	return
}

/*
Sets the edge at which the tabs for switching pages in the
notebook are drawn.
*/
func (self *_TraitNotebook) SetTabPos(pos C.GtkPositionType) {
	C.gtk_notebook_set_tab_pos(self.CPointer, pos)
	return
}

/*
Sets whether the notebook tab can be reordered
via drag and drop or not.
*/
func (self *_TraitNotebook) SetTabReorderable(child *Widget, reorderable bool) {
	__cgo__reorderable := C.gboolean(0)
	if reorderable {
		__cgo__reorderable = C.gboolean(1)
	}
	C.gtk_notebook_set_tab_reorderable(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__reorderable)
	return
}

type _TraitNumerableIcon struct{ CPointer *C.GtkNumerableIcon }

/*
Returns the #GIcon that was set as the base background image, or
%NULL if there’s none. The caller of this function does not own
a reference to the returned #GIcon.
*/
func (self *_TraitNumerableIcon) GetBackgroundGicon() (return__ *C.GIcon) {
	return__ = C.gtk_numerable_icon_get_background_gicon(self.CPointer)
	return
}

/*
Returns the icon name used as the base background image,
or %NULL if there’s none.
*/
func (self *_TraitNumerableIcon) GetBackgroundIconName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_numerable_icon_get_background_icon_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the value currently displayed by @self.
*/
func (self *_TraitNumerableIcon) GetCount() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_numerable_icon_get_count(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the currently displayed label of the icon, or %NULL.
*/
func (self *_TraitNumerableIcon) GetLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_numerable_icon_get_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the #GtkStyleContext used by the icon for theming,
or %NULL if there’s none.
*/
func (self *_TraitNumerableIcon) GetStyleContext() (return__ *StyleContext) {
	var __cgo__return__ *C.GtkStyleContext
	__cgo__return__ = C.gtk_numerable_icon_get_style_context(self.CPointer)
	return__ = NewStyleContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Updates the icon to use @icon as the base background image.
If @icon is %NULL, @self will go back using style information
or default theming for its background image.

If this method is called and an icon name was already set as
background for the icon, @icon will be used, i.e. the last method
called between gtk_numerable_icon_set_background_gicon() and
gtk_numerable_icon_set_background_icon_name() has always priority.
*/
func (self *_TraitNumerableIcon) SetBackgroundGicon(icon *C.GIcon) {
	C.gtk_numerable_icon_set_background_gicon(self.CPointer, icon)
	return
}

/*
Updates the icon to use the icon named @icon_name from the
current icon theme as the base background image. If @icon_name
is %NULL, @self will go back using style information or default
theming for its background image.

If this method is called and a #GIcon was already set as
background for the icon, @icon_name will be used, i.e. the
last method called between gtk_numerable_icon_set_background_icon_name()
and gtk_numerable_icon_set_background_gicon() has always priority.
*/
func (self *_TraitNumerableIcon) SetBackgroundIconName(icon_name string) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_numerable_icon_set_background_icon_name(self.CPointer, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

/*
Sets the currently displayed value of @self to @count.

The numeric value is always clamped to make it two digits, i.e.
between -99 and 99. Setting a count of zero removes the emblem.
If this method is called, and a label was already set on the icon,
it will automatically be reset to %NULL before rendering the number,
i.e. the last method called between gtk_numerable_icon_set_count()
and gtk_numerable_icon_set_label() has always priority.
*/
func (self *_TraitNumerableIcon) SetCount(count int) {
	C.gtk_numerable_icon_set_count(self.CPointer, C.gint(count))
	return
}

/*
Sets the currently displayed value of @self to the string
in @label. Setting an empty label removes the emblem.

Note that this is meant for displaying short labels, such as
roman numbers, or single letters. For roman numbers, consider
using the Unicode characters U+2160 - U+217F. Strings longer
than two characters will likely not be rendered very well.

If this method is called, and a number was already set on the
icon, it will automatically be reset to zero before rendering
the label, i.e. the last method called between
gtk_numerable_icon_set_label() and gtk_numerable_icon_set_count()
has always priority.
*/
func (self *_TraitNumerableIcon) SetLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_numerable_icon_set_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Updates the icon to fetch theme information from the
given #GtkStyleContext.
*/
func (self *_TraitNumerableIcon) SetStyleContext(style *StyleContext) {
	C.gtk_numerable_icon_set_style_context(self.CPointer, (*C.GtkStyleContext)(style.CPointer))
	return
}

type _TraitOffscreenWindow struct{ CPointer *C.GtkOffscreenWindow }

/*
Retrieves a snapshot of the contained widget in the form of
a #GdkPixbuf.  This is a new pixbuf with a reference count of 1,
and the application should unreference it once it is no longer
needed.
*/
func (self *_TraitOffscreenWindow) GetPixbuf() (return__ *C.GdkPixbuf) {
	return__ = C.gtk_offscreen_window_get_pixbuf(self.CPointer)
	return
}

/*
Retrieves a snapshot of the contained widget in the form of
a #cairo_surface_t.  If you need to keep this around over window
resizes then you should add a reference to it.
*/
func (self *_TraitOffscreenWindow) GetSurface() (return__ *C.cairo_surface_t) {
	return__ = C.gtk_offscreen_window_get_surface(self.CPointer)
	return
}

type _TraitOverlay struct{ CPointer *C.GtkOverlay }

/*
Adds @widget to @overlay.

The widget will be stacked on top of the main widget
added with gtk_container_add().

The position at which @widget is placed is determined
from its #GtkWidget:halign and #GtkWidget:valign properties.
*/
func (self *_TraitOverlay) AddOverlay(widget *Widget) {
	C.gtk_overlay_add_overlay(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

type _TraitPageSetup struct{ CPointer *C.GtkPageSetup }

/*
Copies a #GtkPageSetup.
*/
func (self *_TraitPageSetup) Copy() (return__ *PageSetup) {
	var __cgo__return__ *C.GtkPageSetup
	__cgo__return__ = C.gtk_page_setup_copy(self.CPointer)
	return__ = NewPageSetupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the bottom margin in units of @unit.
*/
func (self *_TraitPageSetup) GetBottomMargin(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_page_setup_get_bottom_margin(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the left margin in units of @unit.
*/
func (self *_TraitPageSetup) GetLeftMargin(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_page_setup_get_left_margin(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the page orientation of the #GtkPageSetup.
*/
func (self *_TraitPageSetup) GetOrientation() (return__ C.GtkPageOrientation) {
	return__ = C.gtk_page_setup_get_orientation(self.CPointer)
	return
}

/*
Returns the page height in units of @unit.

Note that this function takes orientation and
margins into consideration.
See gtk_page_setup_get_paper_height().
*/
func (self *_TraitPageSetup) GetPageHeight(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_page_setup_get_page_height(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Returns the page width in units of @unit.

Note that this function takes orientation and
margins into consideration.
See gtk_page_setup_get_paper_width().
*/
func (self *_TraitPageSetup) GetPageWidth(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_page_setup_get_page_width(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Returns the paper height in units of @unit.

Note that this function takes orientation, but
not margins into consideration.
See gtk_page_setup_get_page_height().
*/
func (self *_TraitPageSetup) GetPaperHeight(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_page_setup_get_paper_height(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the paper size of the #GtkPageSetup.
*/
func (self *_TraitPageSetup) GetPaperSize() (return__ *PaperSize) {
	var __cgo__return__ *C.GtkPaperSize
	__cgo__return__ = C.gtk_page_setup_get_paper_size(self.CPointer)
	return__ = (*PaperSize)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the paper width in units of @unit.

Note that this function takes orientation, but
not margins into consideration.
See gtk_page_setup_get_page_width().
*/
func (self *_TraitPageSetup) GetPaperWidth(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_page_setup_get_paper_width(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the right margin in units of @unit.
*/
func (self *_TraitPageSetup) GetRightMargin(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_page_setup_get_right_margin(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the top margin in units of @unit.
*/
func (self *_TraitPageSetup) GetTopMargin(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_page_setup_get_top_margin(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Reads the page setup from the file @file_name.
See gtk_page_setup_to_file().
*/
func (self *_TraitPageSetup) LoadFile(file_name string) (return__ bool, __err__ error) {
	__cgo__file_name := C.CString(file_name)
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_page_setup_load_file(self.CPointer, __cgo__file_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__file_name))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads the page setup from the group @group_name in the key file
@key_file.
*/
func (self *_TraitPageSetup) LoadKeyFile(key_file *C.GKeyFile, group_name string) (return__ bool, __err__ error) {
	__cgo__group_name := (*C.gchar)(unsafe.Pointer(C.CString(group_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_page_setup_load_key_file(self.CPointer, key_file, __cgo__group_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__group_name))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets the bottom margin of the #GtkPageSetup.
*/
func (self *_TraitPageSetup) SetBottomMargin(margin float64, unit C.GtkUnit) {
	C.gtk_page_setup_set_bottom_margin(self.CPointer, C.gdouble(margin), unit)
	return
}

/*
Sets the left margin of the #GtkPageSetup.
*/
func (self *_TraitPageSetup) SetLeftMargin(margin float64, unit C.GtkUnit) {
	C.gtk_page_setup_set_left_margin(self.CPointer, C.gdouble(margin), unit)
	return
}

/*
Sets the page orientation of the #GtkPageSetup.
*/
func (self *_TraitPageSetup) SetOrientation(orientation C.GtkPageOrientation) {
	C.gtk_page_setup_set_orientation(self.CPointer, orientation)
	return
}

/*
Sets the paper size of the #GtkPageSetup without
changing the margins. See
gtk_page_setup_set_paper_size_and_default_margins().
*/
func (self *_TraitPageSetup) SetPaperSize(size *PaperSize) {
	C.gtk_page_setup_set_paper_size(self.CPointer, (*C.GtkPaperSize)(unsafe.Pointer(size)))
	return
}

/*
Sets the paper size of the #GtkPageSetup and modifies
the margins according to the new paper size.
*/
func (self *_TraitPageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	C.gtk_page_setup_set_paper_size_and_default_margins(self.CPointer, (*C.GtkPaperSize)(unsafe.Pointer(size)))
	return
}

/*
Sets the right margin of the #GtkPageSetup.
*/
func (self *_TraitPageSetup) SetRightMargin(margin float64, unit C.GtkUnit) {
	C.gtk_page_setup_set_right_margin(self.CPointer, C.gdouble(margin), unit)
	return
}

/*
Sets the top margin of the #GtkPageSetup.
*/
func (self *_TraitPageSetup) SetTopMargin(margin float64, unit C.GtkUnit) {
	C.gtk_page_setup_set_top_margin(self.CPointer, C.gdouble(margin), unit)
	return
}

/*
This function saves the information from @setup to @file_name.
*/
func (self *_TraitPageSetup) ToFile(file_name string) (return__ bool, __err__ error) {
	__cgo__file_name := C.CString(file_name)
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_page_setup_to_file(self.CPointer, __cgo__file_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__file_name))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This function adds the page setup from @setup to @key_file.
*/
func (self *_TraitPageSetup) ToKeyFile(key_file *C.GKeyFile, group_name string) {
	__cgo__group_name := (*C.gchar)(unsafe.Pointer(C.CString(group_name)))
	C.gtk_page_setup_to_key_file(self.CPointer, key_file, __cgo__group_name)
	C.free(unsafe.Pointer(__cgo__group_name))
	return
}

type _TraitPaned struct{ CPointer *C.GtkPaned }

/*
Adds a child to the top or left pane with default parameters. This is
equivalent to
`gtk_paned_pack1 (paned, child, FALSE, TRUE)`.
*/
func (self *_TraitPaned) Add1(child *Widget) {
	C.gtk_paned_add1(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Adds a child to the bottom or right pane with default parameters. This
is equivalent to
`gtk_paned_pack2 (paned, child, TRUE, TRUE)`.
*/
func (self *_TraitPaned) Add2(child *Widget) {
	C.gtk_paned_add2(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Obtains the first child of the paned widget.
*/
func (self *_TraitPaned) GetChild1() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_paned_get_child1(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Obtains the second child of the paned widget.
*/
func (self *_TraitPaned) GetChild2() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_paned_get_child2(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the #GdkWindow of the handle. This function is
useful when handling button or motion events because it
enables the callback to distinguish between the window
of the paned, a child and the handle.
*/
func (self *_TraitPaned) GetHandleWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_paned_get_handle_window(self.CPointer)
	return
}

/*
Obtains the position of the divider between the two panes.
*/
func (self *_TraitPaned) GetPosition() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_paned_get_position(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Adds a child to the top or left pane.
*/
func (self *_TraitPaned) Pack1(child *Widget, resize bool, shrink bool) {
	__cgo__resize := C.gboolean(0)
	if resize {
		__cgo__resize = C.gboolean(1)
	}
	__cgo__shrink := C.gboolean(0)
	if shrink {
		__cgo__shrink = C.gboolean(1)
	}
	C.gtk_paned_pack1(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__resize, __cgo__shrink)
	return
}

/*
Adds a child to the bottom or right pane.
*/
func (self *_TraitPaned) Pack2(child *Widget, resize bool, shrink bool) {
	__cgo__resize := C.gboolean(0)
	if resize {
		__cgo__resize = C.gboolean(1)
	}
	__cgo__shrink := C.gboolean(0)
	if shrink {
		__cgo__shrink = C.gboolean(1)
	}
	C.gtk_paned_pack2(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__resize, __cgo__shrink)
	return
}

/*
Sets the position of the divider between the two panes.
*/
func (self *_TraitPaned) SetPosition(position int) {
	C.gtk_paned_set_position(self.CPointer, C.gint(position))
	return
}

type _TraitPlacesSidebar struct{ CPointer *C.GtkPlacesSidebar }

/*
Applications may want to present some folders in the places sidebar if
they could be immediately useful to users.  For example, a drawing
program could add a “/usr/share/clipart” location when the sidebar is
being used in an “Insert Clipart” dialog box.

This function adds the specified @location to a special place for immutable
shortcuts.  The shortcuts are application-specific; they are not shared
across applications, and they are not persistent.  If this function
is called multiple times with different locations, then they are added
to the sidebar’s list in the same order as the function is called.
*/
func (self *_TraitPlacesSidebar) AddShortcut(location *C.GFile) {
	C.gtk_places_sidebar_add_shortcut(self.CPointer, location)
	return
}

/*
Returns the value previously set with gtk_places_sidebar_set_local_only().
*/
func (self *_TraitPlacesSidebar) GetLocalOnly() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_places_sidebar_get_local_only(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the currently-selected location in the @sidebar.  This can be #NULL when
nothing is selected, for example, when gtk_places_sidebar_set_location() has
been called with a location that is not among the sidebar’s list of places to
show.

You can use this function to get the selection in the @sidebar.  Also, if you
connect to the #GtkPlacesSidebar::populate-popup signal, you can use this
function to get the location that is being referred to during the callbacks
for your menu items.
*/
func (self *_TraitPlacesSidebar) GetLocation() (return__ *C.GFile) {
	return__ = C.gtk_places_sidebar_get_location(self.CPointer)
	return
}

/*
This function queries the bookmarks added by the user to the places sidebar,
and returns one of them.  This function is used by #GtkFileChooser to implement
the “Alt-1”, “Alt-2”, etc. shortcuts, which activate the cooresponding bookmark.
*/
func (self *_TraitPlacesSidebar) GetNthBookmark(n int) (return__ *C.GFile) {
	return__ = C.gtk_places_sidebar_get_nth_bookmark(self.CPointer, C.gint(n))
	return
}

/*
Gets the open flags.
*/
func (self *_TraitPlacesSidebar) GetOpenFlags() (return__ C.GtkPlacesOpenFlags) {
	return__ = C.gtk_places_sidebar_get_open_flags(self.CPointer)
	return
}

/*
Returns the value previously set with gtk_places_sidebar_set_show_connect_to_server()
*/
func (self *_TraitPlacesSidebar) GetShowConnectToServer() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_places_sidebar_get_show_connect_to_server(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the value previously set with gtk_places_sidebar_set_show_desktop()
*/
func (self *_TraitPlacesSidebar) GetShowDesktop() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_places_sidebar_get_show_desktop(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the list of shortcuts.
*/
func (self *_TraitPlacesSidebar) ListShortcuts() (return__ *C.GSList) {
	return__ = C.gtk_places_sidebar_list_shortcuts(self.CPointer)
	return
}

/*
Removes an application-specific shortcut that has been previously been
inserted with gtk_places_sidebar_add_shortcut().  If the @location is not a
shortcut in the sidebar, then nothing is done.
*/
func (self *_TraitPlacesSidebar) RemoveShortcut(location *C.GFile) {
	C.gtk_places_sidebar_remove_shortcut(self.CPointer, location)
	return
}

/*
Sets whether the @sidebar should only show local files.
*/
func (self *_TraitPlacesSidebar) SetLocalOnly(local_only bool) {
	__cgo__local_only := C.gboolean(0)
	if local_only {
		__cgo__local_only = C.gboolean(1)
	}
	C.gtk_places_sidebar_set_local_only(self.CPointer, __cgo__local_only)
	return
}

/*
Sets the location that is being shown in the widgets surrounding the
@sidebar, for example, in a folder view in a file manager.  In turn, the
@sidebar will highlight that location if it is being shown in the list of
places, or it will unhighlight everything if the @location is not among the
places in the list.
*/
func (self *_TraitPlacesSidebar) SetLocation(location *C.GFile) {
	C.gtk_places_sidebar_set_location(self.CPointer, location)
	return
}

/*
Sets the way in which the calling application can open new locations from
the places sidebar.  For example, some applications only open locations
“directly” into their main view, while others may support opening locations
in a new notebook tab or a new window.

This function is used to tell the places @sidebar about the ways in which the
application can open new locations, so that the sidebar can display (or not)
the “Open in new tab” and “Open in new window” menu items as appropriate.

When the #GtkPlacesSidebar::open-location signal is emitted, its flags
argument will be set to one of the @flags that was passed in
gtk_places_sidebar_set_open_flags().

Passing 0 for @flags will cause #GTK_PLACES_OPEN_NORMAL to always be sent
to callbacks for the “open-location” signal.
*/
func (self *_TraitPlacesSidebar) SetOpenFlags(flags C.GtkPlacesOpenFlags) {
	C.gtk_places_sidebar_set_open_flags(self.CPointer, flags)
	return
}

/*
Sets whether the @sidebar should show an item for connecting to a network server; this is off by default.
An application may want to turn this on if it implements a way for the user to connect
to network servers directly.
*/
func (self *_TraitPlacesSidebar) SetShowConnectToServer(show_connect_to_server bool) {
	__cgo__show_connect_to_server := C.gboolean(0)
	if show_connect_to_server {
		__cgo__show_connect_to_server = C.gboolean(1)
	}
	C.gtk_places_sidebar_set_show_connect_to_server(self.CPointer, __cgo__show_connect_to_server)
	return
}

/*
Sets whether the @sidebar should show an item for the Desktop folder.
The default value for this option is determined by the desktop
environment and the user’s configuration, but this function can be
used to override it on a per-application basis.
*/
func (self *_TraitPlacesSidebar) SetShowDesktop(show_desktop bool) {
	__cgo__show_desktop := C.gboolean(0)
	if show_desktop {
		__cgo__show_desktop = C.gboolean(1)
	}
	C.gtk_places_sidebar_set_show_desktop(self.CPointer, __cgo__show_desktop)
	return
}

type _TraitPlug struct{ CPointer *C.GtkPlug }

/*
Finish the initialization of @plug for a given #GtkSocket identified by
@socket_id. This function will generally only be used by classes deriving from #GtkPlug.
*/
func (self *_TraitPlug) Construct(socket_id C.Window) {
	C.gtk_plug_construct(self.CPointer, socket_id)
	return
}

/*
Finish the initialization of @plug for a given #GtkSocket identified by
@socket_id which is currently displayed on @display.
This function will generally only be used by classes deriving from #GtkPlug.
*/
func (self *_TraitPlug) ConstructForDisplay(display *C.GdkDisplay, socket_id C.Window) {
	C.gtk_plug_construct_for_display(self.CPointer, display, socket_id)
	return
}

/*
Determines whether the plug is embedded in a socket.
*/
func (self *_TraitPlug) GetEmbedded() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_plug_get_embedded(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the window ID of a #GtkPlug widget, which can then
be used to embed this window inside another window, for
instance with gtk_socket_add_id().
*/
func (self *_TraitPlug) GetId() (return__ C.Window) {
	return__ = C.gtk_plug_get_id(self.CPointer)
	return
}

/*
Retrieves the socket the plug is embedded in.
*/
func (self *_TraitPlug) GetSocketWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_plug_get_socket_window(self.CPointer)
	return
}

type _TraitPopover struct{ CPointer *C.GtkPopover }

/*
Establishes a binding between a #GtkPopover and a #GMenuModel.

The contents of @popover are removed and then refilled with menu items
according to @model.  When @model changes, @popover is updated.
Calling this function twice on @popover with different @model will
cause the first binding to be replaced with a binding to the new
model. If @model is %NULL then any previous binding is undone and
all children are removed.

If @action_namespace is non-%NULL then the effect is as if all
actions mentioned in the @model have their names prefixed with the
namespace, plus a dot.  For example, if the action “quit” is
mentioned and @action_namespace is “app” then the effective action
name is “app.quit”.

This function uses #GtkActionable to define the action name and
target values on the created menu items.  If you want to use an
action group other than “app” and “win”, or if you want to use a
#GtkMenuShell outside of a #GtkApplicationWindow, then you will need
to attach your own action group to the widget hierarchy using
gtk_widget_insert_action_group().  As an example, if you created a
group with a “quit” action and inserted it with the name “mygroup”
then you would use the action name “mygroup.quit” in your
#GMenuModel.
*/
func (self *_TraitPopover) BindModel(model *C.GMenuModel, action_namespace string) {
	__cgo__action_namespace := (*C.gchar)(unsafe.Pointer(C.CString(action_namespace)))
	C.gtk_popover_bind_model(self.CPointer, model, __cgo__action_namespace)
	C.free(unsafe.Pointer(__cgo__action_namespace))
	return
}

/*
Returns whether the popover is modal, see gtk_popover_set_modal to
see the implications of this.
*/
func (self *_TraitPopover) GetModal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_popover_get_modal(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
If a rectangle to point to has been set, this function will
return %TRUE and fill in @rect with such rectangle, otherwise
it will return %FALSE and fill in @rect with the attached
widget coordinates.
*/
func (self *_TraitPopover) GetPointingTo() (rect C.GdkRectangle, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_popover_get_pointing_to(self.CPointer, &rect)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the preferred position of @popover.
*/
func (self *_TraitPopover) GetPosition() (return__ C.GtkPositionType) {
	return__ = C.gtk_popover_get_position(self.CPointer)
	return
}

/*
Returns the widget @popover is currently attached to
*/
func (self *_TraitPopover) GetRelativeTo() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_popover_get_relative_to(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Sets whether @popover is modal, a modal popover will grab all input
within the toplevel and grab the keyboard focus on it when being
displayed. Clicking outside the popover area or pressing Esc will
dismiss the popover and ungrab input.
*/
func (self *_TraitPopover) SetModal(modal bool) {
	__cgo__modal := C.gboolean(0)
	if modal {
		__cgo__modal = C.gboolean(1)
	}
	C.gtk_popover_set_modal(self.CPointer, __cgo__modal)
	return
}

/*
Sets the rectangle that @popover will point to, in the
coordinate space of the widget @popover is attached to,
see gtk_popover_set_relative_to().
*/
func (self *_TraitPopover) SetPointingTo(rect *C.GdkRectangle) {
	C.gtk_popover_set_pointing_to(self.CPointer, rect)
	return
}

/*
Sets the preferred position for @popover to appear. If the @popover
is currently visible, it will be immediately updated.

This preference will be respected where possible, although
on lack of space (eg. if close to the window edges), the
#GtkPopover may choose to appear on the opposite side
*/
func (self *_TraitPopover) SetPosition(position C.GtkPositionType) {
	C.gtk_popover_set_position(self.CPointer, position)
	return
}

/*
Sets a new widget to be attached to @popover. If @popover is
visible, the position will be updated.

Note: the ownership of popovers is always given to their @relative_to
widget, so if @relative_to is set to %NULL on an attached @popover, it
will be detached from its previous widget, and consequently destroyed
unless extra references are kept.
*/
func (self *_TraitPopover) SetRelativeTo(relative_to *Widget) {
	C.gtk_popover_set_relative_to(self.CPointer, (*C.GtkWidget)(relative_to.CPointer))
	return
}

type _TraitPrintContext struct{ CPointer *C.GtkPrintContext }

/*
Creates a new #PangoContext that can be used with the
#GtkPrintContext.
*/
func (self *_TraitPrintContext) CreatePangoContext() (return__ *C.PangoContext) {
	return__ = C.gtk_print_context_create_pango_context(self.CPointer)
	return
}

/*
Creates a new #PangoLayout that is suitable for use
with the #GtkPrintContext.
*/
func (self *_TraitPrintContext) CreatePangoLayout() (return__ *C.PangoLayout) {
	return__ = C.gtk_print_context_create_pango_layout(self.CPointer)
	return
}

/*
Obtains the cairo context that is associated with the
#GtkPrintContext.
*/
func (self *_TraitPrintContext) GetCairoContext() (return__ *C.cairo_t) {
	return__ = C.gtk_print_context_get_cairo_context(self.CPointer)
	return
}

/*
Obtains the horizontal resolution of the #GtkPrintContext,
in dots per inch.
*/
func (self *_TraitPrintContext) GetDpiX() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_context_get_dpi_x(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Obtains the vertical resolution of the #GtkPrintContext,
in dots per inch.
*/
func (self *_TraitPrintContext) GetDpiY() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_context_get_dpi_y(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Obtains the hardware printer margins of the #GtkPrintContext, in units.
*/
func (self *_TraitPrintContext) GetHardMargins() (top float64, bottom float64, left float64, right float64, return__ bool) {
	var __cgo__top C.gdouble
	var __cgo__bottom C.gdouble
	var __cgo__left C.gdouble
	var __cgo__right C.gdouble
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_context_get_hard_margins(self.CPointer, &__cgo__top, &__cgo__bottom, &__cgo__left, &__cgo__right)
	top = float64(__cgo__top)
	bottom = float64(__cgo__bottom)
	left = float64(__cgo__left)
	right = float64(__cgo__right)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Obtains the height of the #GtkPrintContext, in pixels.
*/
func (self *_TraitPrintContext) GetHeight() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_context_get_height(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Obtains the #GtkPageSetup that determines the page
dimensions of the #GtkPrintContext.
*/
func (self *_TraitPrintContext) GetPageSetup() (return__ *PageSetup) {
	var __cgo__return__ *C.GtkPageSetup
	__cgo__return__ = C.gtk_print_context_get_page_setup(self.CPointer)
	return__ = NewPageSetupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns a #PangoFontMap that is suitable for use
with the #GtkPrintContext.
*/
func (self *_TraitPrintContext) GetPangoFontmap() (return__ *C.PangoFontMap) {
	return__ = C.gtk_print_context_get_pango_fontmap(self.CPointer)
	return
}

/*
Obtains the width of the #GtkPrintContext, in pixels.
*/
func (self *_TraitPrintContext) GetWidth() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_context_get_width(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Sets a new cairo context on a print context.

This function is intended to be used when implementing
an internal print preview, it is not needed for printing,
since GTK+ itself creates a suitable cairo context in that
case.
*/
func (self *_TraitPrintContext) SetCairoContext(cr *C.cairo_t, dpi_x float64, dpi_y float64) {
	C.gtk_print_context_set_cairo_context(self.CPointer, cr, C.double(dpi_x), C.double(dpi_y))
	return
}

type _TraitPrintOperation struct{ CPointer *C.GtkPrintOperation }

/*
Cancels a running print operation. This function may
be called from a #GtkPrintOperation::begin-print,
#GtkPrintOperation::paginate or #GtkPrintOperation::draw-page
signal handler to stop the currently running print
operation.
*/
func (self *_TraitPrintOperation) Cancel() {
	C.gtk_print_operation_cancel(self.CPointer)
	return
}

/*
Signalize that drawing of particular page is complete.

It is called after completion of page drawing (e.g. drawing in another
thread).
If gtk_print_operation_set_defer_drawing() was called before, then this function
has to be called by application. In another case it is called by the library
itself.
*/
func (self *_TraitPrintOperation) DrawPageFinish() {
	C.gtk_print_operation_draw_page_finish(self.CPointer)
	return
}

/*
Returns the default page setup, see
gtk_print_operation_set_default_page_setup().
*/
func (self *_TraitPrintOperation) GetDefaultPageSetup() (return__ *PageSetup) {
	var __cgo__return__ *C.GtkPageSetup
	__cgo__return__ = C.gtk_print_operation_get_default_page_setup(self.CPointer)
	return__ = NewPageSetupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the value of #GtkPrintOperation:embed-page-setup property.
*/
func (self *_TraitPrintOperation) GetEmbedPageSetup() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_operation_get_embed_page_setup(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Call this when the result of a print operation is
%GTK_PRINT_OPERATION_RESULT_ERROR, either as returned by
gtk_print_operation_run(), or in the #GtkPrintOperation::done signal
handler. The returned #GError will contain more details on what went wrong.
*/
func (self *_TraitPrintOperation) GetError() (__err__ error) {
	var __cgo_error__ *C.GError
	C.gtk_print_operation_get_error(self.CPointer, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the value of #GtkPrintOperation:has-selection property.
*/
func (self *_TraitPrintOperation) GetHasSelection() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_operation_get_has_selection(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the number of pages that will be printed.

Note that this value is set during print preparation phase
(%GTK_PRINT_STATUS_PREPARING), so this function should never be
called before the data generation phase (%GTK_PRINT_STATUS_GENERATING_DATA).
You can connect to the #GtkPrintOperation::status-changed signal
and call gtk_print_operation_get_n_pages_to_print() when
print status is %GTK_PRINT_STATUS_GENERATING_DATA.
This is typically used to track the progress of print operation.
*/
func (self *_TraitPrintOperation) GetNPagesToPrint() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_print_operation_get_n_pages_to_print(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the current print settings.

Note that the return value is %NULL until either
gtk_print_operation_set_print_settings() or
gtk_print_operation_run() have been called.
*/
func (self *_TraitPrintOperation) GetPrintSettings() (return__ *PrintSettings) {
	var __cgo__return__ *C.GtkPrintSettings
	__cgo__return__ = C.gtk_print_operation_get_print_settings(self.CPointer)
	return__ = NewPrintSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the status of the print operation.
Also see gtk_print_operation_get_status_string().
*/
func (self *_TraitPrintOperation) GetStatus() (return__ C.GtkPrintStatus) {
	return__ = C.gtk_print_operation_get_status(self.CPointer)
	return
}

/*
Returns a string representation of the status of the
print operation. The string is translated and suitable
for displaying the print status e.g. in a #GtkStatusbar.

Use gtk_print_operation_get_status() to obtain a status
value that is suitable for programmatic use.
*/
func (self *_TraitPrintOperation) GetStatusString() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_print_operation_get_status_string(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of #GtkPrintOperation:support-selection property.
*/
func (self *_TraitPrintOperation) GetSupportSelection() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_operation_get_support_selection(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
A convenience function to find out if the print operation
is finished, either successfully (%GTK_PRINT_STATUS_FINISHED)
or unsuccessfully (%GTK_PRINT_STATUS_FINISHED_ABORTED).

Note: when you enable print status tracking the print operation
can be in a non-finished state even after done has been called, as
the operation status then tracks the print job status on the printer.
*/
func (self *_TraitPrintOperation) IsFinished() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_operation_is_finished(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Runs the print operation, by first letting the user modify
print settings in the print dialog, and then print the document.

Normally that this function does not return until the rendering of all
pages is complete. You can connect to the
#GtkPrintOperation::status-changed signal on @op to obtain some
information about the progress of the print operation.
Furthermore, it may use a recursive mainloop to show the print dialog.

If you call gtk_print_operation_set_allow_async() or set the
#GtkPrintOperation:allow-async property the operation will run
asynchronously if this is supported on the platform. The
#GtkPrintOperation::done signal will be emitted with the result of the
operation when the it is done (i.e. when the dialog is canceled, or when
the print succeeds or fails).
|[<!-- language="C" -->
if (settings != NULL)
  gtk_print_operation_set_print_settings (print, settings);

if (page_setup != NULL)
  gtk_print_operation_set_default_page_setup (print, page_setup);

g_signal_connect (print, "begin-print",
                  G_CALLBACK (begin_print), &data);
g_signal_connect (print, "draw-page",
                  G_CALLBACK (draw_page), &data);

res = gtk_print_operation_run (print,
                               GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG,
                               parent,
                               &error);

if (res == GTK_PRINT_OPERATION_RESULT_ERROR)
 {
   error_dialog = gtk_message_dialog_new (GTK_WINDOW (parent),
  			                     GTK_DIALOG_DESTROY_WITH_PARENT,
					     GTK_MESSAGE_ERROR,
					     GTK_BUTTONS_CLOSE,
					     "Error printing file:\n%s",
					     error->message);
   g_signal_connect (error_dialog, "response",
                     G_CALLBACK (gtk_widget_destroy), NULL);
   gtk_widget_show (error_dialog);
   g_error_free (error);
 }
else if (res == GTK_PRINT_OPERATION_RESULT_APPLY)
 {
   if (settings != NULL)
g_object_unref (settings);
   settings = g_object_ref (gtk_print_operation_get_print_settings (print));
 }
]|

Note that gtk_print_operation_run() can only be called once on a
given #GtkPrintOperation.
*/
func (self *_TraitPrintOperation) Run(action C.GtkPrintOperationAction, parent *Window) (return__ C.GtkPrintOperationResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.gtk_print_operation_run(self.CPointer, action, (*C.GtkWindow)(parent.CPointer), &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets whether the gtk_print_operation_run() may return
before the print operation is completed. Note that
some platforms may not allow asynchronous operation.
*/
func (self *_TraitPrintOperation) SetAllowAsync(allow_async bool) {
	__cgo__allow_async := C.gboolean(0)
	if allow_async {
		__cgo__allow_async = C.gboolean(1)
	}
	C.gtk_print_operation_set_allow_async(self.CPointer, __cgo__allow_async)
	return
}

/*
Sets the current page.

If this is called before gtk_print_operation_run(),
the user will be able to select to print only the current page.

Note that this only makes sense for pre-paginated documents.
*/
func (self *_TraitPrintOperation) SetCurrentPage(current_page int) {
	C.gtk_print_operation_set_current_page(self.CPointer, C.gint(current_page))
	return
}

/*
Sets the label for the tab holding custom widgets.
*/
func (self *_TraitPrintOperation) SetCustomTabLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_print_operation_set_custom_tab_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Makes @default_page_setup the default page setup for @op.

This page setup will be used by gtk_print_operation_run(),
but it can be overridden on a per-page basis by connecting
to the #GtkPrintOperation::request-page-setup signal.
*/
func (self *_TraitPrintOperation) SetDefaultPageSetup(default_page_setup *PageSetup) {
	C.gtk_print_operation_set_default_page_setup(self.CPointer, (*C.GtkPageSetup)(default_page_setup.CPointer))
	return
}

/*
Sets up the #GtkPrintOperation to wait for calling of
gtk_print_operation_draw_page_finish() from application. It can
be used for drawing page in another thread.

This function must be called in the callback of “draw-page” signal.
*/
func (self *_TraitPrintOperation) SetDeferDrawing() {
	C.gtk_print_operation_set_defer_drawing(self.CPointer)
	return
}

/*
Embed page size combo box and orientation combo box into page setup page.
Selected page setup is stored as default page setup in #GtkPrintOperation.
*/
func (self *_TraitPrintOperation) SetEmbedPageSetup(embed bool) {
	__cgo__embed := C.gboolean(0)
	if embed {
		__cgo__embed = C.gboolean(1)
	}
	C.gtk_print_operation_set_embed_page_setup(self.CPointer, __cgo__embed)
	return
}

/*
Sets up the #GtkPrintOperation to generate a file instead
of showing the print dialog. The indended use of this function
is for implementing “Export to PDF” actions. Currently, PDF
is the only supported format.

“Print to PDF” support is independent of this and is done
by letting the user pick the “Print to PDF” item from the list
of printers in the print dialog.
*/
func (self *_TraitPrintOperation) SetExportFilename(filename string) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	C.gtk_print_operation_set_export_filename(self.CPointer, __cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	return
}

/*
Sets whether there is a selection to print.

Application has to set number of pages to which the selection
will draw by gtk_print_operation_set_n_pages() in a callback of
#GtkPrintOperation::begin-print.
*/
func (self *_TraitPrintOperation) SetHasSelection(has_selection bool) {
	__cgo__has_selection := C.gboolean(0)
	if has_selection {
		__cgo__has_selection = C.gboolean(1)
	}
	C.gtk_print_operation_set_has_selection(self.CPointer, __cgo__has_selection)
	return
}

/*
Sets the name of the print job. The name is used to identify
the job (e.g. in monitoring applications like eggcups).

If you don’t set a job name, GTK+ picks a default one by
numbering successive print jobs.
*/
func (self *_TraitPrintOperation) SetJobName(job_name string) {
	__cgo__job_name := (*C.gchar)(unsafe.Pointer(C.CString(job_name)))
	C.gtk_print_operation_set_job_name(self.CPointer, __cgo__job_name)
	C.free(unsafe.Pointer(__cgo__job_name))
	return
}

/*
Sets the number of pages in the document.

This must be set to a positive number
before the rendering starts. It may be set in a
#GtkPrintOperation::begin-print signal hander.

Note that the page numbers passed to the
#GtkPrintOperation::request-page-setup
and #GtkPrintOperation::draw-page signals are 0-based, i.e. if
the user chooses to print all pages, the last ::draw-page signal
will be for page @n_pages - 1.
*/
func (self *_TraitPrintOperation) SetNPages(n_pages int) {
	C.gtk_print_operation_set_n_pages(self.CPointer, C.gint(n_pages))
	return
}

/*
Sets the print settings for @op. This is typically used to
re-establish print settings from a previous print operation,
see gtk_print_operation_run().
*/
func (self *_TraitPrintOperation) SetPrintSettings(print_settings *PrintSettings) {
	C.gtk_print_operation_set_print_settings(self.CPointer, (*C.GtkPrintSettings)(print_settings.CPointer))
	return
}

/*
If @show_progress is %TRUE, the print operation will show a
progress dialog during the print operation.
*/
func (self *_TraitPrintOperation) SetShowProgress(show_progress bool) {
	__cgo__show_progress := C.gboolean(0)
	if show_progress {
		__cgo__show_progress = C.gboolean(1)
	}
	C.gtk_print_operation_set_show_progress(self.CPointer, __cgo__show_progress)
	return
}

/*
Sets whether selection is supported by #GtkPrintOperation.
*/
func (self *_TraitPrintOperation) SetSupportSelection(support_selection bool) {
	__cgo__support_selection := C.gboolean(0)
	if support_selection {
		__cgo__support_selection = C.gboolean(1)
	}
	C.gtk_print_operation_set_support_selection(self.CPointer, __cgo__support_selection)
	return
}

/*
If track_status is %TRUE, the print operation will try to continue report
on the status of the print job in the printer queues and printer. This
can allow your application to show things like “out of paper” issues,
and when the print job actually reaches the printer.

This function is often implemented using some form of polling, so it should
not be enabled unless needed.
*/
func (self *_TraitPrintOperation) SetTrackPrintStatus(track_status bool) {
	__cgo__track_status := C.gboolean(0)
	if track_status {
		__cgo__track_status = C.gboolean(1)
	}
	C.gtk_print_operation_set_track_print_status(self.CPointer, __cgo__track_status)
	return
}

/*
Sets up the transformation for the cairo context obtained from
#GtkPrintContext in such a way that distances are measured in
units of @unit.
*/
func (self *_TraitPrintOperation) SetUnit(unit C.GtkUnit) {
	C.gtk_print_operation_set_unit(self.CPointer, unit)
	return
}

/*
If @full_page is %TRUE, the transformation for the cairo context
obtained from #GtkPrintContext puts the origin at the top left
corner of the page (which may not be the top left corner of the
sheet, depending on page orientation and the number of pages per
sheet). Otherwise, the origin is at the top left corner of the
imageable area (i.e. inside the margins).
*/
func (self *_TraitPrintOperation) SetUseFullPage(full_page bool) {
	__cgo__full_page := C.gboolean(0)
	if full_page {
		__cgo__full_page = C.gboolean(1)
	}
	C.gtk_print_operation_set_use_full_page(self.CPointer, __cgo__full_page)
	return
}

type _TraitPrintSettings struct{ CPointer *C.GtkPrintSettings }

/*
Copies a #GtkPrintSettings object.
*/
func (self *_TraitPrintSettings) Copy() (return__ *PrintSettings) {
	var __cgo__return__ *C.GtkPrintSettings
	__cgo__return__ = C.gtk_print_settings_copy(self.CPointer)
	return__ = NewPrintSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Calls @func for each key-value pair of @settings.
*/
func (self *_TraitPrintSettings) Foreach(func_ C.GtkPrintSettingsFunc, user_data unsafe.Pointer) {
	C.gtk_print_settings_foreach(self.CPointer, func_, (C.gpointer)(user_data))
	return
}

/*
Looks up the string value associated with @key.
*/
func (self *_TraitPrintSettings) Get(key string) (return__ string) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_print_settings_get(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the boolean represented by the value
that is associated with @key.

The string “true” represents %TRUE, any other
string %FALSE.
*/
func (self *_TraitPrintSettings) GetBool(key string) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_settings_get_bool(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_COLLATE.
*/
func (self *_TraitPrintSettings) GetCollate() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_settings_get_collate(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
*/
func (self *_TraitPrintSettings) GetDefaultSource() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_print_settings_get_default_source(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_DITHER.
*/
func (self *_TraitPrintSettings) GetDither() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_print_settings_get_dither(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the double value associated with @key, or 0.
*/
func (self *_TraitPrintSettings) GetDouble(key string) (return__ float64) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_settings_get_double(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = float64(__cgo__return__)
	return
}

/*
Returns the floating point number represented by
the value that is associated with @key, or @default_val
if the value does not represent a floating point number.

Floating point numbers are parsed with g_ascii_strtod().
*/
func (self *_TraitPrintSettings) GetDoubleWithDefault(key string, def float64) (return__ float64) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_settings_get_double_with_default(self.CPointer, __cgo__key, C.gdouble(def))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_DUPLEX.
*/
func (self *_TraitPrintSettings) GetDuplex() (return__ C.GtkPrintDuplex) {
	return__ = C.gtk_print_settings_get_duplex(self.CPointer)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_FINISHINGS.
*/
func (self *_TraitPrintSettings) GetFinishings() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_print_settings_get_finishings(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the integer value of @key, or 0.
*/
func (self *_TraitPrintSettings) GetInt(key string) (return__ int) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_print_settings_get_int(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = int(__cgo__return__)
	return
}

/*
Returns the value of @key, interpreted as
an integer, or the default value.
*/
func (self *_TraitPrintSettings) GetIntWithDefault(key string, def int) (return__ int) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_print_settings_get_int_with_default(self.CPointer, __cgo__key, C.gint(def))
	C.free(unsafe.Pointer(__cgo__key))
	return__ = int(__cgo__return__)
	return
}

/*
Returns the value associated with @key, interpreted
as a length. The returned value is converted to @units.
*/
func (self *_TraitPrintSettings) GetLength(key string, unit C.GtkUnit) (return__ float64) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_settings_get_length(self.CPointer, __cgo__key, unit)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_MEDIA_TYPE.

The set of media types is defined in PWG 5101.1-2002 PWG.
*/
func (self *_TraitPrintSettings) GetMediaType() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_print_settings_get_media_type(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_N_COPIES.
*/
func (self *_TraitPrintSettings) GetNCopies() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_print_settings_get_n_copies(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_NUMBER_UP.
*/
func (self *_TraitPrintSettings) GetNumberUp() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_print_settings_get_number_up(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
*/
func (self *_TraitPrintSettings) GetNumberUpLayout() (return__ C.GtkNumberUpLayout) {
	return__ = C.gtk_print_settings_get_number_up_layout(self.CPointer)
	return
}

/*
Get the value of %GTK_PRINT_SETTINGS_ORIENTATION,
converted to a #GtkPageOrientation.
*/
func (self *_TraitPrintSettings) GetOrientation() (return__ C.GtkPageOrientation) {
	return__ = C.gtk_print_settings_get_orientation(self.CPointer)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_OUTPUT_BIN.
*/
func (self *_TraitPrintSettings) GetOutputBin() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_print_settings_get_output_bin(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_PAGE_RANGES.
*/
func (self *_TraitPrintSettings) GetPageRanges() (num_ranges int, return__ *PageRange) {
	var __cgo__num_ranges C.gint
	var __cgo__return__ *C.GtkPageRange
	__cgo__return__ = C.gtk_print_settings_get_page_ranges(self.CPointer, &__cgo__num_ranges)
	num_ranges = int(__cgo__num_ranges)
	return__ = (*PageRange)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_PAGE_SET.
*/
func (self *_TraitPrintSettings) GetPageSet() (return__ C.GtkPageSet) {
	return__ = C.gtk_print_settings_get_page_set(self.CPointer)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_PAPER_HEIGHT,
converted to @unit.
*/
func (self *_TraitPrintSettings) GetPaperHeight(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_settings_get_paper_height(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_PAPER_FORMAT,
converted to a #GtkPaperSize.
*/
func (self *_TraitPrintSettings) GetPaperSize() (return__ *PaperSize) {
	var __cgo__return__ *C.GtkPaperSize
	__cgo__return__ = C.gtk_print_settings_get_paper_size(self.CPointer)
	return__ = (*PaperSize)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_PAPER_WIDTH,
converted to @unit.
*/
func (self *_TraitPrintSettings) GetPaperWidth(unit C.GtkUnit) (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_settings_get_paper_width(self.CPointer, unit)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_PRINT_PAGES.
*/
func (self *_TraitPrintSettings) GetPrintPages() (return__ C.GtkPrintPages) {
	return__ = C.gtk_print_settings_get_print_pages(self.CPointer)
	return
}

/*
Convenience function to obtain the value of
%GTK_PRINT_SETTINGS_PRINTER.
*/
func (self *_TraitPrintSettings) GetPrinter() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_print_settings_get_printer(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_PRINTER_LPI.
*/
func (self *_TraitPrintSettings) GetPrinterLpi() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_settings_get_printer_lpi(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_QUALITY.
*/
func (self *_TraitPrintSettings) GetQuality() (return__ C.GtkPrintQuality) {
	return__ = C.gtk_print_settings_get_quality(self.CPointer)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_RESOLUTION.
*/
func (self *_TraitPrintSettings) GetResolution() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_print_settings_get_resolution(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_RESOLUTION_X.
*/
func (self *_TraitPrintSettings) GetResolutionX() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_print_settings_get_resolution_x(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_RESOLUTION_Y.
*/
func (self *_TraitPrintSettings) GetResolutionY() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_print_settings_get_resolution_y(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_REVERSE.
*/
func (self *_TraitPrintSettings) GetReverse() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_settings_get_reverse(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_SCALE.
*/
func (self *_TraitPrintSettings) GetScale() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_print_settings_get_scale(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value of %GTK_PRINT_SETTINGS_USE_COLOR.
*/
func (self *_TraitPrintSettings) GetUseColor() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_settings_get_use_color(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE, if a value is associated with @key.
*/
func (self *_TraitPrintSettings) HasKey(key string) (return__ bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_settings_has_key(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Reads the print settings from @file_name. If the file could not be loaded
then error is set to either a #GFileError or #GKeyFileError.
See gtk_print_settings_to_file().
*/
func (self *_TraitPrintSettings) LoadFile(file_name string) (return__ bool, __err__ error) {
	__cgo__file_name := (*C.gchar)(unsafe.Pointer(C.CString(file_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_settings_load_file(self.CPointer, __cgo__file_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__file_name))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Reads the print settings from the group @group_name in @key_file. If the
file could not be loaded then error is set to either a #GFileError or
#GKeyFileError.
*/
func (self *_TraitPrintSettings) LoadKeyFile(key_file *C.GKeyFile, group_name string) (return__ bool, __err__ error) {
	__cgo__group_name := (*C.gchar)(unsafe.Pointer(C.CString(group_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_settings_load_key_file(self.CPointer, key_file, __cgo__group_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__group_name))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Associates @value with @key.
*/
func (self *_TraitPrintSettings) Set(key string, value string) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	__cgo__value := (*C.gchar)(unsafe.Pointer(C.CString(value)))
	C.gtk_print_settings_set(self.CPointer, __cgo__key, __cgo__value)
	C.free(unsafe.Pointer(__cgo__key))
	C.free(unsafe.Pointer(__cgo__value))
	return
}

/*
Sets @key to a boolean value.
*/
func (self *_TraitPrintSettings) SetBool(key string, value bool) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	__cgo__value := C.gboolean(0)
	if value {
		__cgo__value = C.gboolean(1)
	}
	C.gtk_print_settings_set_bool(self.CPointer, __cgo__key, __cgo__value)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_COLLATE.
*/
func (self *_TraitPrintSettings) SetCollate(collate bool) {
	__cgo__collate := C.gboolean(0)
	if collate {
		__cgo__collate = C.gboolean(1)
	}
	C.gtk_print_settings_set_collate(self.CPointer, __cgo__collate)
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
*/
func (self *_TraitPrintSettings) SetDefaultSource(default_source string) {
	__cgo__default_source := (*C.gchar)(unsafe.Pointer(C.CString(default_source)))
	C.gtk_print_settings_set_default_source(self.CPointer, __cgo__default_source)
	C.free(unsafe.Pointer(__cgo__default_source))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_DITHER.
*/
func (self *_TraitPrintSettings) SetDither(dither string) {
	__cgo__dither := (*C.gchar)(unsafe.Pointer(C.CString(dither)))
	C.gtk_print_settings_set_dither(self.CPointer, __cgo__dither)
	C.free(unsafe.Pointer(__cgo__dither))
	return
}

/*
Sets @key to a double value.
*/
func (self *_TraitPrintSettings) SetDouble(key string, value float64) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	C.gtk_print_settings_set_double(self.CPointer, __cgo__key, C.gdouble(value))
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_DUPLEX.
*/
func (self *_TraitPrintSettings) SetDuplex(duplex C.GtkPrintDuplex) {
	C.gtk_print_settings_set_duplex(self.CPointer, duplex)
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_FINISHINGS.
*/
func (self *_TraitPrintSettings) SetFinishings(finishings string) {
	__cgo__finishings := (*C.gchar)(unsafe.Pointer(C.CString(finishings)))
	C.gtk_print_settings_set_finishings(self.CPointer, __cgo__finishings)
	C.free(unsafe.Pointer(__cgo__finishings))
	return
}

/*
Sets @key to an integer value.
*/
func (self *_TraitPrintSettings) SetInt(key string, value int) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	C.gtk_print_settings_set_int(self.CPointer, __cgo__key, C.gint(value))
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Associates a length in units of @unit with @key.
*/
func (self *_TraitPrintSettings) SetLength(key string, value float64, unit C.GtkUnit) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	C.gtk_print_settings_set_length(self.CPointer, __cgo__key, C.gdouble(value), unit)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_MEDIA_TYPE.

The set of media types is defined in PWG 5101.1-2002 PWG.
*/
func (self *_TraitPrintSettings) SetMediaType(media_type string) {
	__cgo__media_type := (*C.gchar)(unsafe.Pointer(C.CString(media_type)))
	C.gtk_print_settings_set_media_type(self.CPointer, __cgo__media_type)
	C.free(unsafe.Pointer(__cgo__media_type))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_N_COPIES.
*/
func (self *_TraitPrintSettings) SetNCopies(num_copies int) {
	C.gtk_print_settings_set_n_copies(self.CPointer, C.gint(num_copies))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_NUMBER_UP.
*/
func (self *_TraitPrintSettings) SetNumberUp(number_up int) {
	C.gtk_print_settings_set_number_up(self.CPointer, C.gint(number_up))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
*/
func (self *_TraitPrintSettings) SetNumberUpLayout(number_up_layout C.GtkNumberUpLayout) {
	C.gtk_print_settings_set_number_up_layout(self.CPointer, number_up_layout)
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_ORIENTATION.
*/
func (self *_TraitPrintSettings) SetOrientation(orientation C.GtkPageOrientation) {
	C.gtk_print_settings_set_orientation(self.CPointer, orientation)
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_OUTPUT_BIN.
*/
func (self *_TraitPrintSettings) SetOutputBin(output_bin string) {
	__cgo__output_bin := (*C.gchar)(unsafe.Pointer(C.CString(output_bin)))
	C.gtk_print_settings_set_output_bin(self.CPointer, __cgo__output_bin)
	C.free(unsafe.Pointer(__cgo__output_bin))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_PAGE_RANGES.
*/
func (self *_TraitPrintSettings) SetPageRanges(page_ranges *PageRange, num_ranges int) {
	C.gtk_print_settings_set_page_ranges(self.CPointer, (*C.GtkPageRange)(unsafe.Pointer(page_ranges)), C.gint(num_ranges))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_PAGE_SET.
*/
func (self *_TraitPrintSettings) SetPageSet(page_set C.GtkPageSet) {
	C.gtk_print_settings_set_page_set(self.CPointer, page_set)
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_PAPER_HEIGHT.
*/
func (self *_TraitPrintSettings) SetPaperHeight(height float64, unit C.GtkUnit) {
	C.gtk_print_settings_set_paper_height(self.CPointer, C.gdouble(height), unit)
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_PAPER_FORMAT,
%GTK_PRINT_SETTINGS_PAPER_WIDTH and
%GTK_PRINT_SETTINGS_PAPER_HEIGHT.
*/
func (self *_TraitPrintSettings) SetPaperSize(paper_size *PaperSize) {
	C.gtk_print_settings_set_paper_size(self.CPointer, (*C.GtkPaperSize)(unsafe.Pointer(paper_size)))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_PAPER_WIDTH.
*/
func (self *_TraitPrintSettings) SetPaperWidth(width float64, unit C.GtkUnit) {
	C.gtk_print_settings_set_paper_width(self.CPointer, C.gdouble(width), unit)
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_PRINT_PAGES.
*/
func (self *_TraitPrintSettings) SetPrintPages(pages C.GtkPrintPages) {
	C.gtk_print_settings_set_print_pages(self.CPointer, pages)
	return
}

/*
Convenience function to set %GTK_PRINT_SETTINGS_PRINTER
to @printer.
*/
func (self *_TraitPrintSettings) SetPrinter(printer string) {
	__cgo__printer := (*C.gchar)(unsafe.Pointer(C.CString(printer)))
	C.gtk_print_settings_set_printer(self.CPointer, __cgo__printer)
	C.free(unsafe.Pointer(__cgo__printer))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_PRINTER_LPI.
*/
func (self *_TraitPrintSettings) SetPrinterLpi(lpi float64) {
	C.gtk_print_settings_set_printer_lpi(self.CPointer, C.gdouble(lpi))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_QUALITY.
*/
func (self *_TraitPrintSettings) SetQuality(quality C.GtkPrintQuality) {
	C.gtk_print_settings_set_quality(self.CPointer, quality)
	return
}

/*
Sets the values of %GTK_PRINT_SETTINGS_RESOLUTION,
%GTK_PRINT_SETTINGS_RESOLUTION_X and
%GTK_PRINT_SETTINGS_RESOLUTION_Y.
*/
func (self *_TraitPrintSettings) SetResolution(resolution int) {
	C.gtk_print_settings_set_resolution(self.CPointer, C.gint(resolution))
	return
}

/*
Sets the values of %GTK_PRINT_SETTINGS_RESOLUTION,
%GTK_PRINT_SETTINGS_RESOLUTION_X and
%GTK_PRINT_SETTINGS_RESOLUTION_Y.
*/
func (self *_TraitPrintSettings) SetResolutionXy(resolution_x int, resolution_y int) {
	C.gtk_print_settings_set_resolution_xy(self.CPointer, C.gint(resolution_x), C.gint(resolution_y))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_REVERSE.
*/
func (self *_TraitPrintSettings) SetReverse(reverse bool) {
	__cgo__reverse := C.gboolean(0)
	if reverse {
		__cgo__reverse = C.gboolean(1)
	}
	C.gtk_print_settings_set_reverse(self.CPointer, __cgo__reverse)
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_SCALE.
*/
func (self *_TraitPrintSettings) SetScale(scale float64) {
	C.gtk_print_settings_set_scale(self.CPointer, C.gdouble(scale))
	return
}

/*
Sets the value of %GTK_PRINT_SETTINGS_USE_COLOR.
*/
func (self *_TraitPrintSettings) SetUseColor(use_color bool) {
	__cgo__use_color := C.gboolean(0)
	if use_color {
		__cgo__use_color = C.gboolean(1)
	}
	C.gtk_print_settings_set_use_color(self.CPointer, __cgo__use_color)
	return
}

/*
This function saves the print settings from @settings to @file_name. If the
file could not be loaded then error is set to either a #GFileError or
#GKeyFileError.
*/
func (self *_TraitPrintSettings) ToFile(file_name string) (return__ bool, __err__ error) {
	__cgo__file_name := (*C.gchar)(unsafe.Pointer(C.CString(file_name)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_print_settings_to_file(self.CPointer, __cgo__file_name, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__file_name))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This function adds the print settings from @settings to @key_file.
*/
func (self *_TraitPrintSettings) ToKeyFile(key_file *C.GKeyFile, group_name string) {
	__cgo__group_name := (*C.gchar)(unsafe.Pointer(C.CString(group_name)))
	C.gtk_print_settings_to_key_file(self.CPointer, key_file, __cgo__group_name)
	C.free(unsafe.Pointer(__cgo__group_name))
	return
}

/*
Removes any value associated with @key.
This has the same effect as setting the value to %NULL.
*/
func (self *_TraitPrintSettings) Unset(key string) {
	__cgo__key := (*C.gchar)(unsafe.Pointer(C.CString(key)))
	C.gtk_print_settings_unset(self.CPointer, __cgo__key)
	C.free(unsafe.Pointer(__cgo__key))
	return
}

type _TraitProgressBar struct{ CPointer *C.GtkProgressBar }

/*
Returns the ellipsizing position of the progress bar.
See gtk_progress_bar_set_ellipsize().
*/
func (self *_TraitProgressBar) GetEllipsize() (return__ C.PangoEllipsizeMode) {
	return__ = C.gtk_progress_bar_get_ellipsize(self.CPointer)
	return
}

/*
Returns the current fraction of the task that’s been completed.
*/
func (self *_TraitProgressBar) GetFraction() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_progress_bar_get_fraction(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value set by gtk_progress_bar_set_inverted().
*/
func (self *_TraitProgressBar) GetInverted() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_progress_bar_get_inverted(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the pulse step set with gtk_progress_bar_set_pulse_step().
*/
func (self *_TraitProgressBar) GetPulseStep() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_progress_bar_get_pulse_step(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value of the #GtkProgressBar:show-text property.
See gtk_progress_bar_set_show_text().
*/
func (self *_TraitProgressBar) GetShowText() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_progress_bar_get_show_text(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the text displayed superimposed on the progress bar,
if any, otherwise %NULL. The return value is a reference
to the text, not a copy of it, so will become invalid
if you change the text in the progress bar.
*/
func (self *_TraitProgressBar) GetText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_progress_bar_get_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Indicates that some progress has been made, but you don’t know how much.
Causes the progress bar to enter “activity mode,” where a block
bounces back and forth. Each call to gtk_progress_bar_pulse()
causes the block to move by a little bit (the amount of movement
per pulse is determined by gtk_progress_bar_set_pulse_step()).
*/
func (self *_TraitProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(self.CPointer)
	return
}

/*
Sets the mode used to ellipsize (add an ellipsis: "...") the text
if there is not enough space to render the entire string.
*/
func (self *_TraitProgressBar) SetEllipsize(mode C.PangoEllipsizeMode) {
	C.gtk_progress_bar_set_ellipsize(self.CPointer, mode)
	return
}

/*
Causes the progress bar to “fill in” the given fraction
of the bar. The fraction should be between 0.0 and 1.0,
inclusive.
*/
func (self *_TraitProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(self.CPointer, C.gdouble(fraction))
	return
}

/*
Progress bars normally grow from top to bottom or left to right.
Inverted progress bars grow in the opposite direction.
*/
func (self *_TraitProgressBar) SetInverted(inverted bool) {
	__cgo__inverted := C.gboolean(0)
	if inverted {
		__cgo__inverted = C.gboolean(1)
	}
	C.gtk_progress_bar_set_inverted(self.CPointer, __cgo__inverted)
	return
}

/*
Sets the fraction of total progress bar length to move the
bouncing block for each call to gtk_progress_bar_pulse().
*/
func (self *_TraitProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(self.CPointer, C.gdouble(fraction))
	return
}

/*
Sets whether the progress bar will show text superimposed
over the bar. The shown text is either the value of
the #GtkProgressBar:text property or, if that is %NULL,
the #GtkProgressBar:fraction value, as a percentage.

To make a progress bar that is styled and sized suitably for containing
text (even if the actual text is blank), set #GtkProgressBar:show-text to
%TRUE and #GtkProgressBar:text to the empty string (not %NULL).
*/
func (self *_TraitProgressBar) SetShowText(show_text bool) {
	__cgo__show_text := C.gboolean(0)
	if show_text {
		__cgo__show_text = C.gboolean(1)
	}
	C.gtk_progress_bar_set_show_text(self.CPointer, __cgo__show_text)
	return
}

/*
Causes the given @text to appear superimposed on the progress bar.

If @text is %NULL and #GtkProgressBar:show-text is %TRUE, the current
value of #GtkProgressBar:fraction will be displayed as a percentage.

If @text is non-%NULL and #GtkProgressBar:show-text is %TRUE, the text will
be displayed. In this case, it will not display the progress percentage.
If @text is the empty string, the progress bar will still be styled and sized
suitably for containing text, as long as #GtkProgressBar:show-text is %TRUE.
*/
func (self *_TraitProgressBar) SetText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_progress_bar_set_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

type _TraitRadioAction struct{ CPointer *C.GtkRadioAction }

// gtk_radio_action_get_current_value is not generated due to deprecation attr

// gtk_radio_action_get_group is not generated due to deprecation attr

// gtk_radio_action_join_group is not generated due to deprecation attr

// gtk_radio_action_set_current_value is not generated due to deprecation attr

// gtk_radio_action_set_group is not generated due to deprecation attr

type _TraitRadioButton struct{ CPointer *C.GtkRadioButton }

/*
Retrieves the group assigned to a radio button.
*/
func (self *_TraitRadioButton) GetGroup() (return__ *C.GSList) {
	return__ = C.gtk_radio_button_get_group(self.CPointer)
	return
}

/*
Joins a #GtkRadioButton object to the group of another #GtkRadioButton object

Use this in language bindings instead of the gtk_radio_button_get_group()
and gtk_radio_button_set_group() methods

A common way to set up a group of radio buttons is the following:
|[<!-- language="C" -->
  GtkRadioButton *radio_button;
  GtkRadioButton *last_button;

  while ( ...more buttons to add... )
    {
       radio_button = gtk_radio_button_new (...);

       gtk_radio_button_join_group (radio_button, last_button);
       last_button = radio_button;
    }
]|
*/
func (self *_TraitRadioButton) JoinGroup(group_source *RadioButton) {
	C.gtk_radio_button_join_group(self.CPointer, (*C.GtkRadioButton)(group_source.CPointer))
	return
}

/*
Sets a #GtkRadioButton’s group. It should be noted that this does not change
the layout of your interface in any way, so if you are changing the group,
it is likely you will need to re-arrange the user interface to reflect these
changes.
*/
func (self *_TraitRadioButton) SetGroup(group *C.GSList) {
	C.gtk_radio_button_set_group(self.CPointer, group)
	return
}

type _TraitRadioMenuItem struct{ CPointer *C.GtkRadioMenuItem }

/*
Returns the group to which the radio menu item belongs, as a #GList of
#GtkRadioMenuItem. The list belongs to GTK+ and should not be freed.
*/
func (self *_TraitRadioMenuItem) GetGroup() (return__ *C.GSList) {
	return__ = C.gtk_radio_menu_item_get_group(self.CPointer)
	return
}

/*
Sets the group of a radio menu item, or changes it.
*/
func (self *_TraitRadioMenuItem) SetGroup(group *C.GSList) {
	C.gtk_radio_menu_item_set_group(self.CPointer, group)
	return
}

type _TraitRadioToolButton struct{ CPointer *C.GtkRadioToolButton }

/*
Returns the radio button group @button belongs to.
*/
func (self *_TraitRadioToolButton) GetGroup() (return__ *C.GSList) {
	return__ = C.gtk_radio_tool_button_get_group(self.CPointer)
	return
}

/*
Adds @button to @group, removing it from the group it belonged to before.
*/
func (self *_TraitRadioToolButton) SetGroup(group *C.GSList) {
	C.gtk_radio_tool_button_set_group(self.CPointer, group)
	return
}

type _TraitRange struct{ CPointer *C.GtkRange }

/*
Get the #GtkAdjustment which is the “model” object for #GtkRange.
See gtk_range_set_adjustment() for details.
The return value does not have a reference added, so should not
be unreferenced.
*/
func (self *_TraitRange) GetAdjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_range_get_adjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the current position of the fill level indicator.
*/
func (self *_TraitRange) GetFillLevel() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_range_get_fill_level(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value set by gtk_range_set_flippable().
*/
func (self *_TraitRange) GetFlippable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_range_get_flippable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_range_set_inverted().
*/
func (self *_TraitRange) GetInverted() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_range_get_inverted(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the sensitivity policy for the stepper that points to the
'lower' end of the GtkRange’s adjustment.
*/
func (self *_TraitRange) GetLowerStepperSensitivity() (return__ C.GtkSensitivityType) {
	return__ = C.gtk_range_get_lower_stepper_sensitivity(self.CPointer)
	return
}

/*
This function is useful mainly for #GtkRange subclasses.

See gtk_range_set_min_slider_size().
*/
func (self *_TraitRange) GetMinSliderSize() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_range_get_min_slider_size(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
This function returns the area that contains the range’s trough
and its steppers, in widget->window coordinates.

This function is useful mainly for #GtkRange subclasses.
*/
func (self *_TraitRange) GetRangeRect() (range_rect C.GdkRectangle) {
	C.gtk_range_get_range_rect(self.CPointer, &range_rect)
	return
}

/*
Gets whether the range is restricted to the fill level.
*/
func (self *_TraitRange) GetRestrictToFillLevel() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_range_get_restrict_to_fill_level(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the number of digits to round the value to when
it changes. See #GtkRange::change-value.
*/
func (self *_TraitRange) GetRoundDigits() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_range_get_round_digits(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets whether the range displays the fill level graphically.
*/
func (self *_TraitRange) GetShowFillLevel() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_range_get_show_fill_level(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function returns sliders range along the long dimension,
in widget->window coordinates.

This function is useful mainly for #GtkRange subclasses.
*/
func (self *_TraitRange) GetSliderRange() (slider_start int, slider_end int) {
	var __cgo__slider_start C.gint
	var __cgo__slider_end C.gint
	C.gtk_range_get_slider_range(self.CPointer, &__cgo__slider_start, &__cgo__slider_end)
	slider_start = int(__cgo__slider_start)
	slider_end = int(__cgo__slider_end)
	return
}

/*
This function is useful mainly for #GtkRange subclasses.

See gtk_range_set_slider_size_fixed().
*/
func (self *_TraitRange) GetSliderSizeFixed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_range_get_slider_size_fixed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the sensitivity policy for the stepper that points to the
'upper' end of the GtkRange’s adjustment.
*/
func (self *_TraitRange) GetUpperStepperSensitivity() (return__ C.GtkSensitivityType) {
	return__ = C.gtk_range_get_upper_stepper_sensitivity(self.CPointer)
	return
}

/*
Gets the current value of the range.
*/
func (self *_TraitRange) GetValue() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_range_get_value(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Sets the adjustment to be used as the “model” object for this range
widget. The adjustment indicates the current range value, the
minimum and maximum range values, the step/page increments used
for keybindings and scrolling, and the page size. The page size
is normally 0 for #GtkScale and nonzero for #GtkScrollbar, and
indicates the size of the visible area of the widget being scrolled.
The page size affects the size of the scrollbar slider.
*/
func (self *_TraitRange) SetAdjustment(adjustment *Adjustment) {
	C.gtk_range_set_adjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
Set the new position of the fill level indicator.

The “fill level” is probably best described by its most prominent
use case, which is an indicator for the amount of pre-buffering in
a streaming media player. In that use case, the value of the range
would indicate the current play position, and the fill level would
be the position up to which the file/stream has been downloaded.

This amount of prebuffering can be displayed on the range’s trough
and is themeable separately from the trough. To enable fill level
display, use gtk_range_set_show_fill_level(). The range defaults
to not showing the fill level.

Additionally, it’s possible to restrict the range’s slider position
to values which are smaller than the fill level. This is controller
by gtk_range_set_restrict_to_fill_level() and is by default
enabled.
*/
func (self *_TraitRange) SetFillLevel(fill_level float64) {
	C.gtk_range_set_fill_level(self.CPointer, C.gdouble(fill_level))
	return
}

/*
If a range is flippable, it will switch its direction if it is
horizontal and its direction is %GTK_TEXT_DIR_RTL.

See gtk_widget_get_direction().
*/
func (self *_TraitRange) SetFlippable(flippable bool) {
	__cgo__flippable := C.gboolean(0)
	if flippable {
		__cgo__flippable = C.gboolean(1)
	}
	C.gtk_range_set_flippable(self.CPointer, __cgo__flippable)
	return
}

/*
Sets the step and page sizes for the range.
The step size is used when the user clicks the #GtkScrollbar
arrows or moves #GtkScale via arrow keys. The page size
is used for example when moving via Page Up or Page Down keys.
*/
func (self *_TraitRange) SetIncrements(step float64, page float64) {
	C.gtk_range_set_increments(self.CPointer, C.gdouble(step), C.gdouble(page))
	return
}

/*
Ranges normally move from lower to higher values as the
slider moves from top to bottom or left to right. Inverted
ranges have higher values at the top or on the right rather than
on the bottom or left.
*/
func (self *_TraitRange) SetInverted(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_range_set_inverted(self.CPointer, __cgo__setting)
	return
}

/*
Sets the sensitivity policy for the stepper that points to the
'lower' end of the GtkRange’s adjustment.
*/
func (self *_TraitRange) SetLowerStepperSensitivity(sensitivity C.GtkSensitivityType) {
	C.gtk_range_set_lower_stepper_sensitivity(self.CPointer, sensitivity)
	return
}

/*
Sets the minimum size of the range’s slider.

This function is useful mainly for #GtkRange subclasses.
*/
func (self *_TraitRange) SetMinSliderSize(min_size int) {
	C.gtk_range_set_min_slider_size(self.CPointer, C.gint(min_size))
	return
}

/*
Sets the allowable values in the #GtkRange, and clamps the range
value to be between @min and @max. (If the range has a non-zero
page size, it is clamped between @min and @max - page-size.)
*/
func (self *_TraitRange) SetRange(min float64, max float64) {
	C.gtk_range_set_range(self.CPointer, C.gdouble(min), C.gdouble(max))
	return
}

/*
Sets whether the slider is restricted to the fill level. See
gtk_range_set_fill_level() for a general description of the fill
level concept.
*/
func (self *_TraitRange) SetRestrictToFillLevel(restrict_to_fill_level bool) {
	__cgo__restrict_to_fill_level := C.gboolean(0)
	if restrict_to_fill_level {
		__cgo__restrict_to_fill_level = C.gboolean(1)
	}
	C.gtk_range_set_restrict_to_fill_level(self.CPointer, __cgo__restrict_to_fill_level)
	return
}

/*
Sets the number of digits to round the value to when
it changes. See #GtkRange::change-value.
*/
func (self *_TraitRange) SetRoundDigits(round_digits int) {
	C.gtk_range_set_round_digits(self.CPointer, C.gint(round_digits))
	return
}

/*
Sets whether a graphical fill level is show on the trough. See
gtk_range_set_fill_level() for a general description of the fill
level concept.
*/
func (self *_TraitRange) SetShowFillLevel(show_fill_level bool) {
	__cgo__show_fill_level := C.gboolean(0)
	if show_fill_level {
		__cgo__show_fill_level = C.gboolean(1)
	}
	C.gtk_range_set_show_fill_level(self.CPointer, __cgo__show_fill_level)
	return
}

/*
Sets whether the range’s slider has a fixed size, or a size that
depends on its adjustment’s page size.

This function is useful mainly for #GtkRange subclasses.
*/
func (self *_TraitRange) SetSliderSizeFixed(size_fixed bool) {
	__cgo__size_fixed := C.gboolean(0)
	if size_fixed {
		__cgo__size_fixed = C.gboolean(1)
	}
	C.gtk_range_set_slider_size_fixed(self.CPointer, __cgo__size_fixed)
	return
}

/*
Sets the sensitivity policy for the stepper that points to the
'upper' end of the GtkRange’s adjustment.
*/
func (self *_TraitRange) SetUpperStepperSensitivity(sensitivity C.GtkSensitivityType) {
	C.gtk_range_set_upper_stepper_sensitivity(self.CPointer, sensitivity)
	return
}

/*
Sets the current value of the range; if the value is outside the
minimum or maximum range values, it will be clamped to fit inside
them. The range emits the #GtkRange::value-changed signal if the
value changes.
*/
func (self *_TraitRange) SetValue(value float64) {
	C.gtk_range_set_value(self.CPointer, C.gdouble(value))
	return
}

type _TraitRcStyle struct{ CPointer *C.GtkRcStyle }

// gtk_rc_style_copy is not generated due to deprecation attr

type _TraitRecentAction struct{ CPointer *C.GtkRecentAction }

// gtk_recent_action_get_show_numbers is not generated due to deprecation attr

// gtk_recent_action_set_show_numbers is not generated due to deprecation attr

type _TraitRecentChooserDialog struct{ CPointer *C.GtkRecentChooserDialog }

type _TraitRecentChooserMenu struct{ CPointer *C.GtkRecentChooserMenu }

/*
Returns the value set by gtk_recent_chooser_menu_set_show_numbers().
*/
func (self *_TraitRecentChooserMenu) GetShowNumbers() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_recent_chooser_menu_get_show_numbers(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets whether a number should be added to the items of @menu.  The
numbers are shown to provide a unique character for a mnemonic to
be used inside ten menu item’s label.  Only the first the items
get a number to avoid clashes.
*/
func (self *_TraitRecentChooserMenu) SetShowNumbers(show_numbers bool) {
	__cgo__show_numbers := C.gboolean(0)
	if show_numbers {
		__cgo__show_numbers = C.gboolean(1)
	}
	C.gtk_recent_chooser_menu_set_show_numbers(self.CPointer, __cgo__show_numbers)
	return
}

type _TraitRecentChooserWidget struct{ CPointer *C.GtkRecentChooserWidget }

type _TraitRecentFilter struct{ CPointer *C.GtkRecentFilter }

/*
Adds a rule that allows resources based on their age - that is, the number
of days elapsed since they were last modified.
*/
func (self *_TraitRecentFilter) AddAge(days int) {
	C.gtk_recent_filter_add_age(self.CPointer, C.gint(days))
	return
}

/*
Adds a rule that allows resources based on the name of the application
that has registered them.
*/
func (self *_TraitRecentFilter) AddApplication(application string) {
	__cgo__application := (*C.gchar)(unsafe.Pointer(C.CString(application)))
	C.gtk_recent_filter_add_application(self.CPointer, __cgo__application)
	C.free(unsafe.Pointer(__cgo__application))
	return
}

/*
Adds a rule to a filter that allows resources based on a custom callback
function. The bitfield @needed which is passed in provides information
about what sorts of information that the filter function needs;
this allows GTK+ to avoid retrieving expensive information when
it isn’t needed by the filter.
*/
func (self *_TraitRecentFilter) AddCustom(needed C.GtkRecentFilterFlags, func_ C.GtkRecentFilterFunc, data unsafe.Pointer, data_destroy C.GDestroyNotify) {
	C.gtk_recent_filter_add_custom(self.CPointer, needed, func_, (C.gpointer)(data), data_destroy)
	return
}

/*
Adds a rule that allows resources based on the name of the group
to which they belong
*/
func (self *_TraitRecentFilter) AddGroup(group string) {
	__cgo__group := (*C.gchar)(unsafe.Pointer(C.CString(group)))
	C.gtk_recent_filter_add_group(self.CPointer, __cgo__group)
	C.free(unsafe.Pointer(__cgo__group))
	return
}

/*
Adds a rule that allows resources based on their registered MIME type.
*/
func (self *_TraitRecentFilter) AddMimeType(mime_type string) {
	__cgo__mime_type := (*C.gchar)(unsafe.Pointer(C.CString(mime_type)))
	C.gtk_recent_filter_add_mime_type(self.CPointer, __cgo__mime_type)
	C.free(unsafe.Pointer(__cgo__mime_type))
	return
}

/*
Adds a rule that allows resources based on a pattern matching their
display name.
*/
func (self *_TraitRecentFilter) AddPattern(pattern string) {
	__cgo__pattern := (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	C.gtk_recent_filter_add_pattern(self.CPointer, __cgo__pattern)
	C.free(unsafe.Pointer(__cgo__pattern))
	return
}

/*
Adds a rule allowing image files in the formats supported
by GdkPixbuf.
*/
func (self *_TraitRecentFilter) AddPixbufFormats() {
	C.gtk_recent_filter_add_pixbuf_formats(self.CPointer)
	return
}

/*
Tests whether a file should be displayed according to @filter.
The #GtkRecentFilterInfo @filter_info should include
the fields returned from gtk_recent_filter_get_needed(), and
must set the #GtkRecentFilterInfo.contains field of @filter_info
to indicate which fields have been set.

This function will not typically be used by applications; it
is intended principally for use in the implementation of
#GtkRecentChooser.
*/
func (self *_TraitRecentFilter) Filter(filter_info *RecentFilterInfo) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_recent_filter_filter(self.CPointer, (*C.GtkRecentFilterInfo)(unsafe.Pointer(filter_info)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the human-readable name for the filter.
See gtk_recent_filter_set_name().
*/
func (self *_TraitRecentFilter) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_recent_filter_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the fields that need to be filled in for the #GtkRecentFilterInfo
passed to gtk_recent_filter_filter()

This function will not typically be used by applications; it
is intended principally for use in the implementation of
#GtkRecentChooser.
*/
func (self *_TraitRecentFilter) GetNeeded() (return__ C.GtkRecentFilterFlags) {
	return__ = C.gtk_recent_filter_get_needed(self.CPointer)
	return
}

/*
Sets the human-readable name of the filter; this is the string
that will be displayed in the recently used resources selector
user interface if there is a selectable list of filters.
*/
func (self *_TraitRecentFilter) SetName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_recent_filter_set_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

type _TraitRecentManager struct{ CPointer *C.GtkRecentManager }

/*
Adds a new resource, pointed by @uri, into the recently used
resources list, using the metadata specified inside the #GtkRecentData-struct
passed in @recent_data.

The passed URI will be used to identify this resource inside the
list.

In order to register the new recently used resource, metadata about
the resource must be passed as well as the URI; the metadata is
stored in a #GtkRecentData-struct, which must contain the MIME
type of the resource pointed by the URI; the name of the application
that is registering the item, and a command line to be used when
launching the item.

Optionally, a #GtkRecentData-struct might contain a UTF-8 string
to be used when viewing the item instead of the last component of the
URI; a short description of the item; whether the item should be
considered private - that is, should be displayed only by the
applications that have registered it.
*/
func (self *_TraitRecentManager) AddFull(uri string, recent_data *RecentData) (return__ bool) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_recent_manager_add_full(self.CPointer, __cgo__uri, (*C.GtkRecentData)(unsafe.Pointer(recent_data)))
	C.free(unsafe.Pointer(__cgo__uri))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Adds a new resource, pointed by @uri, into the recently used
resources list.

This function automatically retrieves some of the needed
metadata and setting other metadata to common default values; it
then feeds the data to gtk_recent_manager_add_full().

See gtk_recent_manager_add_full() if you want to explicitly
define the metadata for the resource pointed by @uri.
*/
func (self *_TraitRecentManager) AddItem(uri string) (return__ bool) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_recent_manager_add_item(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the list of recently used resources.
*/
func (self *_TraitRecentManager) GetItems() (return__ *C.GList) {
	return__ = C.gtk_recent_manager_get_items(self.CPointer)
	return
}

/*
Checks whether there is a recently used resource registered
with @uri inside the recent manager.
*/
func (self *_TraitRecentManager) HasItem(uri string) (return__ bool) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_recent_manager_has_item(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Searches for a URI inside the recently used resources list, and
returns a #GtkRecentInfo-struct containing informations about the resource
like its MIME type, or its display name.
*/
func (self *_TraitRecentManager) LookupItem(uri string) (return__ *RecentInfo, __err__ error) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.GtkRecentInfo
	__cgo__return__ = C.gtk_recent_manager_lookup_item(self.CPointer, __cgo__uri, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__uri))
	return__ = (*RecentInfo)(unsafe.Pointer(__cgo__return__))
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Changes the location of a recently used resource from @uri to @new_uri.

Please note that this function will not affect the resource pointed
by the URIs, but only the URI used in the recently used resources list.
*/
func (self *_TraitRecentManager) MoveItem(uri string, new_uri string) (return__ bool, __err__ error) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	__cgo__new_uri := (*C.gchar)(unsafe.Pointer(C.CString(new_uri)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_recent_manager_move_item(self.CPointer, __cgo__uri, __cgo__new_uri, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__uri))
	C.free(unsafe.Pointer(__cgo__new_uri))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Purges every item from the recently used resources list.
*/
func (self *_TraitRecentManager) PurgeItems() (return__ int, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_recent_manager_purge_items(self.CPointer, &__cgo_error__)
	return__ = int(__cgo__return__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Removes a resource pointed by @uri from the recently used resources
list handled by a recent manager.
*/
func (self *_TraitRecentManager) RemoveItem(uri string) (return__ bool, __err__ error) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_recent_manager_remove_item(self.CPointer, __cgo__uri, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__uri))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

type _TraitRevealer struct{ CPointer *C.GtkRevealer }

/*
Returns whether the child is fully revealed, ie wether
the transition to the revealed state is completed.
*/
func (self *_TraitRevealer) GetChildRevealed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_revealer_get_child_revealed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the child is currently
revealed. See gtk_revealer_set_reveal_child().

This function returns %TRUE as soon as the transition
is to the revealed state is started. To learn whether
the child is fully revealed (ie the transition is completed),
use gtk_revealer_get_child_revealed().
*/
func (self *_TraitRevealer) GetRevealChild() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_revealer_get_reveal_child(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the amount of time (in milliseconds) that
transitions will take.
*/
func (self *_TraitRevealer) GetTransitionDuration() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_revealer_get_transition_duration(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the type of animation that will be used
for transitions in @revealer.
*/
func (self *_TraitRevealer) GetTransitionType() (return__ C.GtkRevealerTransitionType) {
	return__ = C.gtk_revealer_get_transition_type(self.CPointer)
	return
}

/*
Tells the #GtkRevealer to reveal or conceal its child.

The transition will be animated with the current
transition type of @revealer.
*/
func (self *_TraitRevealer) SetRevealChild(reveal_child bool) {
	__cgo__reveal_child := C.gboolean(0)
	if reveal_child {
		__cgo__reveal_child = C.gboolean(1)
	}
	C.gtk_revealer_set_reveal_child(self.CPointer, __cgo__reveal_child)
	return
}

/*
Sets the duration that transitions will take.
*/
func (self *_TraitRevealer) SetTransitionDuration(duration uint) {
	C.gtk_revealer_set_transition_duration(self.CPointer, C.guint(duration))
	return
}

/*
Sets the type of animation that will be used for
transitions in @revealer. Available types include
various kinds of fades and slides.
*/
func (self *_TraitRevealer) SetTransitionType(transition C.GtkRevealerTransitionType) {
	C.gtk_revealer_set_transition_type(self.CPointer, transition)
	return
}

type _TraitScale struct{ CPointer *C.GtkScale }

/*
Adds a mark at @value.

A mark is indicated visually by drawing a tick mark next to the scale,
and GTK+ makes it easy for the user to position the scale exactly at the
marks value.

If @markup is not %NULL, text is shown next to the tick mark.

To remove marks from a scale, use gtk_scale_clear_marks().
*/
func (self *_TraitScale) AddMark(value float64, position C.GtkPositionType, markup string) {
	__cgo__markup := (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	C.gtk_scale_add_mark(self.CPointer, C.gdouble(value), position, __cgo__markup)
	C.free(unsafe.Pointer(__cgo__markup))
	return
}

/*
Removes any marks that have been added with gtk_scale_add_mark().
*/
func (self *_TraitScale) ClearMarks() {
	C.gtk_scale_clear_marks(self.CPointer)
	return
}

/*
Gets the number of decimal places that are displayed in the value.
*/
func (self *_TraitScale) GetDigits() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_scale_get_digits(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns whether the current value is displayed as a string
next to the slider.
*/
func (self *_TraitScale) GetDrawValue() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_scale_get_draw_value(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the scale has an origin.
*/
func (self *_TraitScale) GetHasOrigin() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_scale_get_has_origin(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the #PangoLayout used to display the scale. The returned
object is owned by the scale so does not need to be freed by
the caller.
*/
func (self *_TraitScale) GetLayout() (return__ *C.PangoLayout) {
	return__ = C.gtk_scale_get_layout(self.CPointer)
	return
}

/*
Obtains the coordinates where the scale will draw the
#PangoLayout representing the text in the scale. Remember
when using the #PangoLayout function you need to convert to
and from pixels using PANGO_PIXELS() or #PANGO_SCALE.

If the #GtkScale:draw-value property is %FALSE, the return
values are undefined.
*/
func (self *_TraitScale) GetLayoutOffsets() (x int, y int) {
	var __cgo__x C.gint
	var __cgo__y C.gint
	C.gtk_scale_get_layout_offsets(self.CPointer, &__cgo__x, &__cgo__y)
	x = int(__cgo__x)
	y = int(__cgo__y)
	return
}

/*
Gets the position in which the current value is displayed.
*/
func (self *_TraitScale) GetValuePos() (return__ C.GtkPositionType) {
	return__ = C.gtk_scale_get_value_pos(self.CPointer)
	return
}

/*
Sets the number of decimal places that are displayed in the value.
Also causes the value of the adjustment to be rounded off to this
number of digits, so the retrieved value matches the value the user saw.
*/
func (self *_TraitScale) SetDigits(digits int) {
	C.gtk_scale_set_digits(self.CPointer, C.gint(digits))
	return
}

/*
Specifies whether the current value is displayed as a string next
to the slider.
*/
func (self *_TraitScale) SetDrawValue(draw_value bool) {
	__cgo__draw_value := C.gboolean(0)
	if draw_value {
		__cgo__draw_value = C.gboolean(1)
	}
	C.gtk_scale_set_draw_value(self.CPointer, __cgo__draw_value)
	return
}

/*
If @has_origin is set to %TRUE (the default),
the scale will highlight the part of the scale
between the origin (bottom or left side) of the scale
and the current value.
*/
func (self *_TraitScale) SetHasOrigin(has_origin bool) {
	__cgo__has_origin := C.gboolean(0)
	if has_origin {
		__cgo__has_origin = C.gboolean(1)
	}
	C.gtk_scale_set_has_origin(self.CPointer, __cgo__has_origin)
	return
}

/*
Sets the position in which the current value is displayed.
*/
func (self *_TraitScale) SetValuePos(pos C.GtkPositionType) {
	C.gtk_scale_set_value_pos(self.CPointer, pos)
	return
}

type _TraitScaleButton struct{ CPointer *C.GtkScaleButton }

/*
Gets the #GtkAdjustment associated with the #GtkScaleButton’s scale.
See gtk_range_get_adjustment() for details.
*/
func (self *_TraitScaleButton) GetAdjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_scale_button_get_adjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the minus button of the #GtkScaleButton.
*/
func (self *_TraitScaleButton) GetMinusButton() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_scale_button_get_minus_button(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the plus button of the #GtkScaleButton.
*/
func (self *_TraitScaleButton) GetPlusButton() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_scale_button_get_plus_button(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Retrieves the popup of the #GtkScaleButton.
*/
func (self *_TraitScaleButton) GetPopup() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_scale_button_get_popup(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the current value of the scale button.
*/
func (self *_TraitScaleButton) GetValue() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_scale_button_get_value(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Sets the #GtkAdjustment to be used as a model
for the #GtkScaleButton’s scale.
See gtk_range_set_adjustment() for details.
*/
func (self *_TraitScaleButton) SetAdjustment(adjustment *Adjustment) {
	C.gtk_scale_button_set_adjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
Sets the icons to be used by the scale button.
For details, see the #GtkScaleButton:icons property.
*/
func (self *_TraitScaleButton) SetIcons(icons []string) {
	__header__icons := (*reflect.SliceHeader)(unsafe.Pointer(&icons))
	C.gtk_scale_button_set_icons(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__icons.Data)))
	return
}

/*
Sets the current value of the scale; if the value is outside
the minimum or maximum range values, it will be clamped to fit
inside them. The scale button emits the #GtkScaleButton::value-changed
signal if the value changes.
*/
func (self *_TraitScaleButton) SetValue(value float64) {
	C.gtk_scale_button_set_value(self.CPointer, C.gdouble(value))
	return
}

type _TraitScrollbar struct{ CPointer *C.GtkScrollbar }

type _TraitScrolledWindow struct{ CPointer *C.GtkScrolledWindow }

// gtk_scrolled_window_add_with_viewport is not generated due to deprecation attr

/*
Return whether button presses are captured during kinetic
scrolling. See gtk_scrolled_window_set_capture_button_press().
*/
func (self *_TraitScrolledWindow) GetCaptureButtonPress() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_scrolled_window_get_capture_button_press(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the horizontal scrollbar’s adjustment, used to connect the
horizontal scrollbar to the child widget’s horizontal scroll
functionality.
*/
func (self *_TraitScrolledWindow) GetHadjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_scrolled_window_get_hadjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the horizontal scrollbar of @scrolled_window.
*/
func (self *_TraitScrolledWindow) GetHscrollbar() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_scrolled_window_get_hscrollbar(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the specified kinetic scrolling behavior.
*/
func (self *_TraitScrolledWindow) GetKineticScrolling() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_scrolled_window_get_kinetic_scrolling(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the minimal content height of @scrolled_window, or -1 if not set.
*/
func (self *_TraitScrolledWindow) GetMinContentHeight() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_scrolled_window_get_min_content_height(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the minimum content width of @scrolled_window, or -1 if not set.
*/
func (self *_TraitScrolledWindow) GetMinContentWidth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_scrolled_window_get_min_content_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the placement of the contents with respect to the scrollbars
for the scrolled window. See gtk_scrolled_window_set_placement().
*/
func (self *_TraitScrolledWindow) GetPlacement() (return__ C.GtkCornerType) {
	return__ = C.gtk_scrolled_window_get_placement(self.CPointer)
	return
}

/*
Retrieves the current policy values for the horizontal and vertical
scrollbars. See gtk_scrolled_window_set_policy().
*/
func (self *_TraitScrolledWindow) GetPolicy() (hscrollbar_policy C.GtkPolicyType, vscrollbar_policy C.GtkPolicyType) {
	C.gtk_scrolled_window_get_policy(self.CPointer, &hscrollbar_policy, &vscrollbar_policy)
	return
}

/*
Gets the shadow type of the scrolled window. See
gtk_scrolled_window_set_shadow_type().
*/
func (self *_TraitScrolledWindow) GetShadowType() (return__ C.GtkShadowType) {
	return__ = C.gtk_scrolled_window_get_shadow_type(self.CPointer)
	return
}

/*
Returns the vertical scrollbar’s adjustment, used to connect the
vertical scrollbar to the child widget’s vertical scroll functionality.
*/
func (self *_TraitScrolledWindow) GetVadjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_scrolled_window_get_vadjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the vertical scrollbar of @scrolled_window.
*/
func (self *_TraitScrolledWindow) GetVscrollbar() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_scrolled_window_get_vscrollbar(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Changes the behaviour of @scrolled_window wrt. to the initial
event that possibly starts kinetic scrolling. When @capture_button_press
is set to %TRUE, the event is captured by the scrolled window, and
then later replayed if it is meant to go to the child widget.

This should be enabled if any child widgets perform non-reversible
actions on #GtkWidget::button-press-event. If they don't, and handle
additionally handle #GtkWidget::grab-broken-event, it might be better
to set @capture_button_press to %FALSE.

This setting only has an effect if kinetic scrolling is enabled.
*/
func (self *_TraitScrolledWindow) SetCaptureButtonPress(capture_button_press bool) {
	__cgo__capture_button_press := C.gboolean(0)
	if capture_button_press {
		__cgo__capture_button_press = C.gboolean(1)
	}
	C.gtk_scrolled_window_set_capture_button_press(self.CPointer, __cgo__capture_button_press)
	return
}

/*
Sets the #GtkAdjustment for the horizontal scrollbar.
*/
func (self *_TraitScrolledWindow) SetHadjustment(hadjustment *Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(self.CPointer, (*C.GtkAdjustment)(hadjustment.CPointer))
	return
}

/*
Turns kinetic scrolling on or off.
Kinetic scrolling only applies to devices with source
%GDK_SOURCE_TOUCHSCREEN.
*/
func (self *_TraitScrolledWindow) SetKineticScrolling(kinetic_scrolling bool) {
	__cgo__kinetic_scrolling := C.gboolean(0)
	if kinetic_scrolling {
		__cgo__kinetic_scrolling = C.gboolean(1)
	}
	C.gtk_scrolled_window_set_kinetic_scrolling(self.CPointer, __cgo__kinetic_scrolling)
	return
}

/*
Sets the minimum height that @scrolled_window should keep visible.
Note that this can and (usually will) be smaller than the minimum
size of the content.
*/
func (self *_TraitScrolledWindow) SetMinContentHeight(height int) {
	C.gtk_scrolled_window_set_min_content_height(self.CPointer, C.gint(height))
	return
}

/*
Sets the minimum width that @scrolled_window should keep visible.
Note that this can and (usually will) be smaller than the minimum
size of the content.
*/
func (self *_TraitScrolledWindow) SetMinContentWidth(width int) {
	C.gtk_scrolled_window_set_min_content_width(self.CPointer, C.gint(width))
	return
}

/*
Sets the placement of the contents with respect to the scrollbars
for the scrolled window.

The default is %GTK_CORNER_TOP_LEFT, meaning the child is
in the top left, with the scrollbars underneath and to the right.
Other values in #GtkCornerType are %GTK_CORNER_TOP_RIGHT,
%GTK_CORNER_BOTTOM_LEFT, and %GTK_CORNER_BOTTOM_RIGHT.

See also gtk_scrolled_window_get_placement() and
gtk_scrolled_window_unset_placement().
*/
func (self *_TraitScrolledWindow) SetPlacement(window_placement C.GtkCornerType) {
	C.gtk_scrolled_window_set_placement(self.CPointer, window_placement)
	return
}

/*
Sets the scrollbar policy for the horizontal and vertical scrollbars.

The policy determines when the scrollbar should appear; it is a value
from the #GtkPolicyType enumeration. If %GTK_POLICY_ALWAYS, the
scrollbar is always present; if %GTK_POLICY_NEVER, the scrollbar is
never present; if %GTK_POLICY_AUTOMATIC, the scrollbar is present only
if needed (that is, if the slider part of the bar would be smaller
than the trough - the display is larger than the page size).
*/
func (self *_TraitScrolledWindow) SetPolicy(hscrollbar_policy C.GtkPolicyType, vscrollbar_policy C.GtkPolicyType) {
	C.gtk_scrolled_window_set_policy(self.CPointer, hscrollbar_policy, vscrollbar_policy)
	return
}

/*
Changes the type of shadow drawn around the contents of
@scrolled_window.
*/
func (self *_TraitScrolledWindow) SetShadowType(type_ C.GtkShadowType) {
	C.gtk_scrolled_window_set_shadow_type(self.CPointer, type_)
	return
}

/*
Sets the #GtkAdjustment for the vertical scrollbar.
*/
func (self *_TraitScrolledWindow) SetVadjustment(vadjustment *Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(self.CPointer, (*C.GtkAdjustment)(vadjustment.CPointer))
	return
}

/*
Unsets the placement of the contents with respect to the scrollbars
for the scrolled window. If no window placement is set for a scrolled
window, it defaults to GTK_CORNER_TOP_LEFT.

See also gtk_scrolled_window_set_placement() and
gtk_scrolled_window_get_placement().
*/
func (self *_TraitScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement(self.CPointer)
	return
}

type _TraitSearchBar struct{ CPointer *C.GtkSearchBar }

/*
Connects the #GtkEntry widget passed as the one to be used in
this search bar. The entry should be a descendant of the search bar.
This is only required if the entry isn’t the direct child of the
search bar (as in our main example).
*/
func (self *_TraitSearchBar) ConnectEntry(entry *Entry) {
	C.gtk_search_bar_connect_entry(self.CPointer, (*C.GtkEntry)(entry.CPointer))
	return
}

/*
Returns whether the search mode is on or off.
*/
func (self *_TraitSearchBar) GetSearchMode() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_search_bar_get_search_mode(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the close button is shown.
*/
func (self *_TraitSearchBar) GetShowCloseButton() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_search_bar_get_show_close_button(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function should be called when the top-level
window which contains the search bar received a key event.

If the key event is handled by the search bar, the bar will
be shown, the entry populated with the entered text and %GDK_EVENT_STOP
will be returned. The caller should ensure that events are
not propagated further.

If no entry has been connected to the search bar, using
gtk_search_bar_connect_entry(), this function will return
immediately with a warning.

## Showing the search bar on key presses

|[<!-- language="C" -->
static gboolean
on_key_press_event (GtkWidget *widget,
                    GdkEvent  *event,
                    gpointer   user_data)
{
  GtkSearchBar *bar = GTK_SEARCH_BAR (user_data);
  return gtk_search_bar_handle_event (bar, event);
}

g_signal_connect (window,
                 "key-press-event",
                  G_CALLBACK (on_key_press_event),
                  search_bar);
]|
*/
func (self *_TraitSearchBar) HandleEvent(event *C.GdkEvent) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_search_bar_handle_event(self.CPointer, event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Switches the search mode on or off.
*/
func (self *_TraitSearchBar) SetSearchMode(search_mode bool) {
	__cgo__search_mode := C.gboolean(0)
	if search_mode {
		__cgo__search_mode = C.gboolean(1)
	}
	C.gtk_search_bar_set_search_mode(self.CPointer, __cgo__search_mode)
	return
}

/*
Shows or hides the close button. Applications that
already have a “search” toggle button should not show a close
button in their search bar, as it duplicates the role of the
toggle button.
*/
func (self *_TraitSearchBar) SetShowCloseButton(visible bool) {
	__cgo__visible := C.gboolean(0)
	if visible {
		__cgo__visible = C.gboolean(1)
	}
	C.gtk_search_bar_set_show_close_button(self.CPointer, __cgo__visible)
	return
}

type _TraitSearchEntry struct{ CPointer *C.GtkSearchEntry }

type _TraitSeparator struct{ CPointer *C.GtkSeparator }

type _TraitSeparatorMenuItem struct{ CPointer *C.GtkSeparatorMenuItem }

type _TraitSeparatorToolItem struct{ CPointer *C.GtkSeparatorToolItem }

/*
Returns whether @item is drawn as a line, or just blank.
See gtk_separator_tool_item_set_draw().
*/
func (self *_TraitSeparatorToolItem) GetDraw() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_separator_tool_item_get_draw(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether @item is drawn as a vertical line, or just blank.
Setting this to %FALSE along with gtk_tool_item_set_expand() is useful
to create an item that forces following items to the end of the toolbar.
*/
func (self *_TraitSeparatorToolItem) SetDraw(draw bool) {
	__cgo__draw := C.gboolean(0)
	if draw {
		__cgo__draw = C.gboolean(1)
	}
	C.gtk_separator_tool_item_set_draw(self.CPointer, __cgo__draw)
	return
}

type _TraitSettings struct{ CPointer *C.GtkSettings }

func (self *_TraitSettings) SetDoubleProperty(name string, v_double float64, origin string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__origin := (*C.gchar)(unsafe.Pointer(C.CString(origin)))
	C.gtk_settings_set_double_property(self.CPointer, __cgo__name, C.gdouble(v_double), __cgo__origin)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__origin))
	return
}

func (self *_TraitSettings) SetLongProperty(name string, v_long int64, origin string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__origin := (*C.gchar)(unsafe.Pointer(C.CString(origin)))
	C.gtk_settings_set_long_property(self.CPointer, __cgo__name, C.glong(v_long), __cgo__origin)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__origin))
	return
}

func (self *_TraitSettings) SetPropertyValue(name string, svalue *SettingsValue) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_settings_set_property_value(self.CPointer, __cgo__name, (*C.GtkSettingsValue)(unsafe.Pointer(svalue)))
	C.free(unsafe.Pointer(__cgo__name))
	return
}

func (self *_TraitSettings) SetStringProperty(name string, v_string string, origin string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__v_string := (*C.gchar)(unsafe.Pointer(C.CString(v_string)))
	__cgo__origin := (*C.gchar)(unsafe.Pointer(C.CString(origin)))
	C.gtk_settings_set_string_property(self.CPointer, __cgo__name, __cgo__v_string, __cgo__origin)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__v_string))
	C.free(unsafe.Pointer(__cgo__origin))
	return
}

type _TraitSizeGroup struct{ CPointer *C.GtkSizeGroup }

/*
Adds a widget to a #GtkSizeGroup. In the future, the requisition
of the widget will be determined as the maximum of its requisition
and the requisition of the other widgets in the size group.
Whether this applies horizontally, vertically, or in both directions
depends on the mode of the size group. See gtk_size_group_set_mode().

When the widget is destroyed or no longer referenced elsewhere, it will
be removed from the size group.
*/
func (self *_TraitSizeGroup) AddWidget(widget *Widget) {
	C.gtk_size_group_add_widget(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

/*
Returns if invisible widgets are ignored when calculating the size.
*/
func (self *_TraitSizeGroup) GetIgnoreHidden() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_size_group_get_ignore_hidden(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the current mode of the size group. See gtk_size_group_set_mode().
*/
func (self *_TraitSizeGroup) GetMode() (return__ C.GtkSizeGroupMode) {
	return__ = C.gtk_size_group_get_mode(self.CPointer)
	return
}

/*
Returns the list of widgets associated with @size_group.
*/
func (self *_TraitSizeGroup) GetWidgets() (return__ *C.GSList) {
	return__ = C.gtk_size_group_get_widgets(self.CPointer)
	return
}

/*
Removes a widget from a #GtkSizeGroup.
*/
func (self *_TraitSizeGroup) RemoveWidget(widget *Widget) {
	C.gtk_size_group_remove_widget(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

/*
Sets whether unmapped widgets should be ignored when
calculating the size.
*/
func (self *_TraitSizeGroup) SetIgnoreHidden(ignore_hidden bool) {
	__cgo__ignore_hidden := C.gboolean(0)
	if ignore_hidden {
		__cgo__ignore_hidden = C.gboolean(1)
	}
	C.gtk_size_group_set_ignore_hidden(self.CPointer, __cgo__ignore_hidden)
	return
}

/*
Sets the #GtkSizeGroupMode of the size group. The mode of the size
group determines whether the widgets in the size group should
all have the same horizontal requisition (%GTK_SIZE_GROUP_HORIZONTAL)
all have the same vertical requisition (%GTK_SIZE_GROUP_VERTICAL),
or should all have the same requisition in both directions
(%GTK_SIZE_GROUP_BOTH).
*/
func (self *_TraitSizeGroup) SetMode(mode C.GtkSizeGroupMode) {
	C.gtk_size_group_set_mode(self.CPointer, mode)
	return
}

type _TraitSocket struct{ CPointer *C.GtkSocket }

/*
Adds an XEMBED client, such as a #GtkPlug, to the #GtkSocket.  The
client may be in the same process or in a different process.

To embed a #GtkPlug in a #GtkSocket, you can either create the
#GtkPlug with `gtk_plug_new (0)`, call
gtk_plug_get_id() to get the window ID of the plug, and then pass that to the
gtk_socket_add_id(), or you can call gtk_socket_get_id() to get the
window ID for the socket, and call gtk_plug_new() passing in that
ID.

The #GtkSocket must have already be added into a toplevel window
 before you can make this call.
*/
func (self *_TraitSocket) AddId(window C.Window) {
	C.gtk_socket_add_id(self.CPointer, window)
	return
}

/*
Gets the window ID of a #GtkSocket widget, which can then
be used to create a client embedded inside the socket, for
instance with gtk_plug_new().

The #GtkSocket must have already be added into a toplevel window
before you can make this call.
*/
func (self *_TraitSocket) GetId() (return__ C.Window) {
	return__ = C.gtk_socket_get_id(self.CPointer)
	return
}

/*
Retrieves the window of the plug. Use this to check if the plug has
been created inside of the socket.
*/
func (self *_TraitSocket) GetPlugWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_socket_get_plug_window(self.CPointer)
	return
}

type _TraitSpinButton struct{ CPointer *C.GtkSpinButton }

/*
Changes the properties of an existing spin button. The adjustment,
climb rate, and number of decimal places are all changed accordingly,
after this function call.
*/
func (self *_TraitSpinButton) Configure(adjustment *Adjustment, climb_rate float64, digits uint) {
	C.gtk_spin_button_configure(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer), C.gdouble(climb_rate), C.guint(digits))
	return
}

/*
Get the adjustment associated with a #GtkSpinButton
*/
func (self *_TraitSpinButton) GetAdjustment() (return__ *Adjustment) {
	var __cgo__return__ *C.GtkAdjustment
	__cgo__return__ = C.gtk_spin_button_get_adjustment(self.CPointer)
	return__ = NewAdjustmentFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Fetches the precision of @spin_button. See gtk_spin_button_set_digits().
*/
func (self *_TraitSpinButton) GetDigits() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_spin_button_get_digits(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the current step and page the increments used by @spin_button. See
gtk_spin_button_set_increments().
*/
func (self *_TraitSpinButton) GetIncrements() (step float64, page float64) {
	var __cgo__step C.gdouble
	var __cgo__page C.gdouble
	C.gtk_spin_button_get_increments(self.CPointer, &__cgo__step, &__cgo__page)
	step = float64(__cgo__step)
	page = float64(__cgo__page)
	return
}

/*
Returns whether non-numeric text can be typed into the spin button.
See gtk_spin_button_set_numeric().
*/
func (self *_TraitSpinButton) GetNumeric() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_spin_button_get_numeric(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the range allowed for @spin_button.
See gtk_spin_button_set_range().
*/
func (self *_TraitSpinButton) GetRange() (min float64, max float64) {
	var __cgo__min C.gdouble
	var __cgo__max C.gdouble
	C.gtk_spin_button_get_range(self.CPointer, &__cgo__min, &__cgo__max)
	min = float64(__cgo__min)
	max = float64(__cgo__max)
	return
}

/*
Returns whether the values are corrected to the nearest step.
See gtk_spin_button_set_snap_to_ticks().
*/
func (self *_TraitSpinButton) GetSnapToTicks() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_spin_button_get_snap_to_ticks(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the update behavior of a spin button.
See gtk_spin_button_set_update_policy().
*/
func (self *_TraitSpinButton) GetUpdatePolicy() (return__ C.GtkSpinButtonUpdatePolicy) {
	return__ = C.gtk_spin_button_get_update_policy(self.CPointer)
	return
}

/*
Get the value in the @spin_button.
*/
func (self *_TraitSpinButton) GetValue() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.gtk_spin_button_get_value(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Get the value @spin_button represented as an integer.
*/
func (self *_TraitSpinButton) GetValueAsInt() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_spin_button_get_value_as_int(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns whether the spin button’s value wraps around to the
opposite limit when the upper or lower limit of the range is
exceeded. See gtk_spin_button_set_wrap().
*/
func (self *_TraitSpinButton) GetWrap() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_spin_button_get_wrap(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Replaces the #GtkAdjustment associated with @spin_button.
*/
func (self *_TraitSpinButton) SetAdjustment(adjustment *Adjustment) {
	C.gtk_spin_button_set_adjustment(self.CPointer, (*C.GtkAdjustment)(adjustment.CPointer))
	return
}

/*
Set the precision to be displayed by @spin_button. Up to 20 digit precision
is allowed.
*/
func (self *_TraitSpinButton) SetDigits(digits uint) {
	C.gtk_spin_button_set_digits(self.CPointer, C.guint(digits))
	return
}

/*
Sets the step and page increments for spin_button.  This affects how
quickly the value changes when the spin button’s arrows are activated.
*/
func (self *_TraitSpinButton) SetIncrements(step float64, page float64) {
	C.gtk_spin_button_set_increments(self.CPointer, C.gdouble(step), C.gdouble(page))
	return
}

/*
Sets the flag that determines if non-numeric text can be typed
into the spin button.
*/
func (self *_TraitSpinButton) SetNumeric(numeric bool) {
	__cgo__numeric := C.gboolean(0)
	if numeric {
		__cgo__numeric = C.gboolean(1)
	}
	C.gtk_spin_button_set_numeric(self.CPointer, __cgo__numeric)
	return
}

/*
Sets the minimum and maximum allowable values for @spin_button.

If the current value is outside this range, it will be adjusted
to fit within the range, otherwise it will remain unchanged.
*/
func (self *_TraitSpinButton) SetRange(min float64, max float64) {
	C.gtk_spin_button_set_range(self.CPointer, C.gdouble(min), C.gdouble(max))
	return
}

/*
Sets the policy as to whether values are corrected to the
nearest step increment when a spin button is activated after
providing an invalid value.
*/
func (self *_TraitSpinButton) SetSnapToTicks(snap_to_ticks bool) {
	__cgo__snap_to_ticks := C.gboolean(0)
	if snap_to_ticks {
		__cgo__snap_to_ticks = C.gboolean(1)
	}
	C.gtk_spin_button_set_snap_to_ticks(self.CPointer, __cgo__snap_to_ticks)
	return
}

/*
Sets the update behavior of a spin button.
This determines whether the spin button is always updated
or only when a valid value is set.
*/
func (self *_TraitSpinButton) SetUpdatePolicy(policy C.GtkSpinButtonUpdatePolicy) {
	C.gtk_spin_button_set_update_policy(self.CPointer, policy)
	return
}

/*
Sets the value of @spin_button.
*/
func (self *_TraitSpinButton) SetValue(value float64) {
	C.gtk_spin_button_set_value(self.CPointer, C.gdouble(value))
	return
}

/*
Sets the flag that determines if a spin button value wraps
around to the opposite limit when the upper or lower limit
of the range is exceeded.
*/
func (self *_TraitSpinButton) SetWrap(wrap bool) {
	__cgo__wrap := C.gboolean(0)
	if wrap {
		__cgo__wrap = C.gboolean(1)
	}
	C.gtk_spin_button_set_wrap(self.CPointer, __cgo__wrap)
	return
}

/*
Increment or decrement a spin button’s value in a specified
direction by a specified amount.
*/
func (self *_TraitSpinButton) Spin(direction C.GtkSpinType, increment float64) {
	C.gtk_spin_button_spin(self.CPointer, direction, C.gdouble(increment))
	return
}

/*
Manually force an update of the spin button.
*/
func (self *_TraitSpinButton) Update() {
	C.gtk_spin_button_update(self.CPointer)
	return
}

type _TraitSpinner struct{ CPointer *C.GtkSpinner }

/*
Starts the animation of the spinner.
*/
func (self *_TraitSpinner) Start() {
	C.gtk_spinner_start(self.CPointer)
	return
}

/*
Stops the animation of the spinner.
*/
func (self *_TraitSpinner) Stop() {
	C.gtk_spinner_stop(self.CPointer)
	return
}

type _TraitStack struct{ CPointer *C.GtkStack }

/*
Adds a child to @stack.
The child is identified by the @name.
*/
func (self *_TraitStack) AddNamed(child *Widget, name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_stack_add_named(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Adds a child to @stack.
The child is identified by the @name. The @title
will be used by #GtkStackSwitcher to represent
@child in a tab bar, so it should be short.
*/
func (self *_TraitStack) AddTitled(child *Widget, name string, title string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_stack_add_titled(self.CPointer, (*C.GtkWidget)(child.CPointer), __cgo__name, __cgo__title)
	C.free(unsafe.Pointer(__cgo__name))
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
Finds the child of the #GtkStack with the name given as
the argument. Returns %NULL if there is no child with this
name.
*/
func (self *_TraitStack) GetChildByName(name string) (return__ *Widget) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_stack_get_child_by_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets whether @stack is homogeneous.
See gtk_stack_set_homogeneous().
*/
func (self *_TraitStack) GetHomogeneous() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_stack_get_homogeneous(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the amount of time (in milliseconds) that
transitions between pages in @stack will take.
*/
func (self *_TraitStack) GetTransitionDuration() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_stack_get_transition_duration(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Returns whether the @stack is currently in a transition from one page to
another.
*/
func (self *_TraitStack) GetTransitionRunning() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_stack_get_transition_running(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the type of animation that will be used
for transitions between pages in @stack.
*/
func (self *_TraitStack) GetTransitionType() (return__ C.GtkStackTransitionType) {
	return__ = C.gtk_stack_get_transition_type(self.CPointer)
	return
}

/*
Gets the currently visible child of @stack, or %NULL if
there are no visible children.
*/
func (self *_TraitStack) GetVisibleChild() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_stack_get_visible_child(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the name of the currently visible child of @stack, or
%NULL if there is no visible child.
*/
func (self *_TraitStack) GetVisibleChildName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_stack_get_visible_child_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Sets the #GtkStack to be homogeneous or not. If it
is homogeneous, the #GtkStack will request the same
size for all its children. If it isn't, the stack
may change size when a different child becomes visible.
*/
func (self *_TraitStack) SetHomogeneous(homogeneous bool) {
	__cgo__homogeneous := C.gboolean(0)
	if homogeneous {
		__cgo__homogeneous = C.gboolean(1)
	}
	C.gtk_stack_set_homogeneous(self.CPointer, __cgo__homogeneous)
	return
}

/*
Sets the duration that transitions between pages in @stack
will take.
*/
func (self *_TraitStack) SetTransitionDuration(duration uint) {
	C.gtk_stack_set_transition_duration(self.CPointer, C.guint(duration))
	return
}

/*
Sets the type of animation that will be used for
transitions between pages in @stack. Available
types include various kinds of fades and slides.

The transition type can be changed without problems
at runtime, so it is possible to change the animation
based on the page that is about to become current.
*/
func (self *_TraitStack) SetTransitionType(transition C.GtkStackTransitionType) {
	C.gtk_stack_set_transition_type(self.CPointer, transition)
	return
}

/*
Makes @child the visible child of @stack.

If @child is different from the currently
visible child, the transition between the
two will be animated with the current
transition type of @stack.

Note that the @child widget has to be visible itself
(see gtk_widget_show()) in order to become the visible
child of @stack.
*/
func (self *_TraitStack) SetVisibleChild(child *Widget) {
	C.gtk_stack_set_visible_child(self.CPointer, (*C.GtkWidget)(child.CPointer))
	return
}

/*
Makes the child with the given name visible.

Note that the child widget has to be visible itself
(see gtk_widget_show()) in order to become the visible
child of @stack.
*/
func (self *_TraitStack) SetVisibleChildFull(name string, transition C.GtkStackTransitionType) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_stack_set_visible_child_full(self.CPointer, __cgo__name, transition)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Makes the child with the given name visible.

If @child is different from the currently
visible child, the transition between the
two will be animated with the current
transition type of @stack.

Note that the child widget has to be visible itself
(see gtk_widget_show()) in order to become the visible
child of @stack.
*/
func (self *_TraitStack) SetVisibleChildName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_stack_set_visible_child_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

type _TraitStackSwitcher struct{ CPointer *C.GtkStackSwitcher }

/*
Retrieves the stack.
See gtk_stack_switcher_set_stack().
*/
func (self *_TraitStackSwitcher) GetStack() (return__ *Stack) {
	var __cgo__return__ *C.GtkStack
	__cgo__return__ = C.gtk_stack_switcher_get_stack(self.CPointer)
	return__ = NewStackFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Sets the stack to control.
*/
func (self *_TraitStackSwitcher) SetStack(stack *Stack) {
	C.gtk_stack_switcher_set_stack(self.CPointer, (*C.GtkStack)(stack.CPointer))
	return
}

type _TraitStatusIcon struct{ CPointer *C.GtkStatusIcon }

/*
Obtains information about the location of the status icon
on screen. This information can be used to e.g. position
popups like notification bubbles.

See gtk_status_icon_position_menu() for a more convenient
alternative for positioning menus.

Note that some platforms do not allow GTK+ to provide
this information, and even on platforms that do allow it,
the information is not reliable unless the status icon
is embedded in a notification area, see
gtk_status_icon_is_embedded().
*/
func (self *_TraitStatusIcon) GetGeometry() (screen *C.GdkScreen, area C.GdkRectangle, orientation C.GtkOrientation, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_status_icon_get_geometry(self.CPointer, &screen, &area, &orientation)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the #GIcon being displayed by the #GtkStatusIcon.
The storage type of the status icon must be %GTK_IMAGE_EMPTY or
%GTK_IMAGE_GICON (see gtk_status_icon_get_storage_type()).
The caller of this function does not own a reference to the
returned #GIcon.

If this function fails, @icon is left unchanged;
*/
func (self *_TraitStatusIcon) GetGicon() (return__ *C.GIcon) {
	return__ = C.gtk_status_icon_get_gicon(self.CPointer)
	return
}

/*
Returns the current value of the has-tooltip property.
See #GtkStatusIcon:has-tooltip for more information.
*/
func (self *_TraitStatusIcon) GetHasTooltip() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_status_icon_get_has_tooltip(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the name of the icon being displayed by the #GtkStatusIcon.
The storage type of the status icon must be %GTK_IMAGE_EMPTY or
%GTK_IMAGE_ICON_NAME (see gtk_status_icon_get_storage_type()).
The returned string is owned by the #GtkStatusIcon and should not
be freed or modified.
*/
func (self *_TraitStatusIcon) GetIconName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_status_icon_get_icon_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #GdkPixbuf being displayed by the #GtkStatusIcon.
The storage type of the status icon must be %GTK_IMAGE_EMPTY or
%GTK_IMAGE_PIXBUF (see gtk_status_icon_get_storage_type()).
The caller of this function does not own a reference to the
returned pixbuf.
*/
func (self *_TraitStatusIcon) GetPixbuf() (return__ *C.GdkPixbuf) {
	return__ = C.gtk_status_icon_get_pixbuf(self.CPointer)
	return
}

/*
Returns the #GdkScreen associated with @status_icon.
*/
func (self *_TraitStatusIcon) GetScreen() (return__ *C.GdkScreen) {
	return__ = C.gtk_status_icon_get_screen(self.CPointer)
	return
}

/*
Gets the size in pixels that is available for the image.
Stock icons and named icons adapt their size automatically
if the size of the notification area changes. For other
storage types, the size-changed signal can be used to
react to size changes.

Note that the returned size is only meaningful while the
status icon is embedded (see gtk_status_icon_is_embedded()).
*/
func (self *_TraitStatusIcon) GetSize() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_status_icon_get_size(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

// gtk_status_icon_get_stock is not generated due to deprecation attr

/*
Gets the type of representation being used by the #GtkStatusIcon
to store image data. If the #GtkStatusIcon has no image data,
the return value will be %GTK_IMAGE_EMPTY.
*/
func (self *_TraitStatusIcon) GetStorageType() (return__ C.GtkImageType) {
	return__ = C.gtk_status_icon_get_storage_type(self.CPointer)
	return
}

/*
Gets the title of this tray icon. See gtk_status_icon_set_title().
*/
func (self *_TraitStatusIcon) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_status_icon_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the contents of the tooltip for @status_icon.
*/
func (self *_TraitStatusIcon) GetTooltipMarkup() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_status_icon_get_tooltip_markup(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the contents of the tooltip for @status_icon.
*/
func (self *_TraitStatusIcon) GetTooltipText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_status_icon_get_tooltip_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns whether the status icon is visible or not.
Note that being visible does not guarantee that
the user can actually see the icon, see also
gtk_status_icon_is_embedded().
*/
func (self *_TraitStatusIcon) GetVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_status_icon_get_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function is only useful on the X11/freedesktop.org platform.
It returns a window ID for the widget in the underlying
status icon implementation.  This is useful for the Galago
notification service, which can send a window ID in the protocol
in order for the server to position notification windows
pointing to a status icon reliably.

This function is not intended for other use cases which are
more likely to be met by one of the non-X11 specific methods, such
as gtk_status_icon_position_menu().
*/
func (self *_TraitStatusIcon) GetX11WindowId() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.gtk_status_icon_get_x11_window_id(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Returns whether the status icon is embedded in a notification
area.
*/
func (self *_TraitStatusIcon) IsEmbedded() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_status_icon_is_embedded(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Makes @status_icon display the file @filename.
See gtk_status_icon_new_from_file() for details.
*/
func (self *_TraitStatusIcon) SetFromFile(filename string) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	C.gtk_status_icon_set_from_file(self.CPointer, __cgo__filename)
	C.free(unsafe.Pointer(__cgo__filename))
	return
}

/*
Makes @status_icon display the #GIcon.
See gtk_status_icon_new_from_gicon() for details.
*/
func (self *_TraitStatusIcon) SetFromGicon(icon *C.GIcon) {
	C.gtk_status_icon_set_from_gicon(self.CPointer, icon)
	return
}

/*
Makes @status_icon display the icon named @icon_name from the
current icon theme.
See gtk_status_icon_new_from_icon_name() for details.
*/
func (self *_TraitStatusIcon) SetFromIconName(icon_name string) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_status_icon_set_from_icon_name(self.CPointer, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

/*
Makes @status_icon display @pixbuf.
See gtk_status_icon_new_from_pixbuf() for details.
*/
func (self *_TraitStatusIcon) SetFromPixbuf(pixbuf *C.GdkPixbuf) {
	C.gtk_status_icon_set_from_pixbuf(self.CPointer, pixbuf)
	return
}

// gtk_status_icon_set_from_stock is not generated due to deprecation attr

/*
Sets the has-tooltip property on @status_icon to @has_tooltip.
See #GtkStatusIcon:has-tooltip for more information.
*/
func (self *_TraitStatusIcon) SetHasTooltip(has_tooltip bool) {
	__cgo__has_tooltip := C.gboolean(0)
	if has_tooltip {
		__cgo__has_tooltip = C.gboolean(1)
	}
	C.gtk_status_icon_set_has_tooltip(self.CPointer, __cgo__has_tooltip)
	return
}

/*
Sets the name of this tray icon.
This should be a string identifying this icon. It is may be
used for sorting the icons in the tray and will not be shown to
the user.
*/
func (self *_TraitStatusIcon) SetName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_status_icon_set_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Sets the #GdkScreen where @status_icon is displayed; if
the icon is already mapped, it will be unmapped, and
then remapped on the new screen.
*/
func (self *_TraitStatusIcon) SetScreen(screen *C.GdkScreen) {
	C.gtk_status_icon_set_screen(self.CPointer, screen)
	return
}

/*
Sets the title of this tray icon.
This should be a short, human-readable, localized string
describing the tray icon. It may be used by tools like screen
readers to render the tray icon.
*/
func (self *_TraitStatusIcon) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_status_icon_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
Sets @markup as the contents of the tooltip, which is marked up with
 the [Pango text markup language][PangoMarkupFormat].

This function will take care of setting #GtkStatusIcon:has-tooltip to %TRUE
and of the default handler for the #GtkStatusIcon::query-tooltip signal.

See also the #GtkStatusIcon:tooltip-markup property and
gtk_tooltip_set_markup().
*/
func (self *_TraitStatusIcon) SetTooltipMarkup(markup string) {
	__cgo__markup := (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	C.gtk_status_icon_set_tooltip_markup(self.CPointer, __cgo__markup)
	C.free(unsafe.Pointer(__cgo__markup))
	return
}

/*
Sets @text as the contents of the tooltip.

This function will take care of setting #GtkStatusIcon:has-tooltip to
%TRUE and of the default handler for the #GtkStatusIcon::query-tooltip
signal.

See also the #GtkStatusIcon:tooltip-text property and
gtk_tooltip_set_text().
*/
func (self *_TraitStatusIcon) SetTooltipText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_status_icon_set_tooltip_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Shows or hides a status icon.
*/
func (self *_TraitStatusIcon) SetVisible(visible bool) {
	__cgo__visible := C.gboolean(0)
	if visible {
		__cgo__visible = C.gboolean(1)
	}
	C.gtk_status_icon_set_visible(self.CPointer, __cgo__visible)
	return
}

type _TraitStatusbar struct{ CPointer *C.GtkStatusbar }

/*
Returns a new context identifier, given a description
of the actual context. Note that the description is
not shown in the UI.
*/
func (self *_TraitStatusbar) GetContextId(context_description string) (return__ uint) {
	__cgo__context_description := (*C.gchar)(unsafe.Pointer(C.CString(context_description)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_statusbar_get_context_id(self.CPointer, __cgo__context_description)
	C.free(unsafe.Pointer(__cgo__context_description))
	return__ = uint(__cgo__return__)
	return
}

/*
Retrieves the box containing the label widget.
*/
func (self *_TraitStatusbar) GetMessageArea() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_statusbar_get_message_area(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Removes the first message in the #GtkStatusbar’s stack
with the given context id.

Note that this may not change the displayed message, if
the message at the top of the stack has a different
context id.
*/
func (self *_TraitStatusbar) Pop(context_id uint) {
	C.gtk_statusbar_pop(self.CPointer, C.guint(context_id))
	return
}

/*
Pushes a new message onto a statusbar’s stack.
*/
func (self *_TraitStatusbar) Push(context_id uint, text string) (return__ uint) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_statusbar_push(self.CPointer, C.guint(context_id), __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return__ = uint(__cgo__return__)
	return
}

/*
Forces the removal of a message from a statusbar’s stack.
The exact @context_id and @message_id must be specified.
*/
func (self *_TraitStatusbar) Remove(context_id uint, message_id uint) {
	C.gtk_statusbar_remove(self.CPointer, C.guint(context_id), C.guint(message_id))
	return
}

/*
Forces the removal of all messages from a statusbar's
stack with the exact @context_id.
*/
func (self *_TraitStatusbar) RemoveAll(context_id uint) {
	C.gtk_statusbar_remove_all(self.CPointer, C.guint(context_id))
	return
}

type _TraitStyle struct{ CPointer *C.GtkStyle }

// gtk_style_apply_default_background is not generated due to deprecation attr

// gtk_style_attach is not generated due to deprecation attr

// gtk_style_copy is not generated due to deprecation attr

// gtk_style_detach is not generated due to deprecation attr

// gtk_style_get is not generated due to varargs

// gtk_style_get_style_property is not generated due to explicit ignore

// gtk_style_get_valist is not generated due to varargs

// gtk_style_has_context is not generated due to explicit ignore

// gtk_style_lookup_color is not generated due to deprecation attr

// gtk_style_lookup_icon_set is not generated due to deprecation attr

// gtk_style_render_icon is not generated due to deprecation attr

// gtk_style_set_background is not generated due to deprecation attr

type _TraitStyleContext struct{ CPointer *C.GtkStyleContext }

/*
Adds a style class to @context, so posterior calls to
gtk_style_context_get() or any of the gtk_render_*()
functions will make use of this new class for styling.

In the CSS file format, a #GtkEntry defining an “entry”
class, would be matched by:

|[
GtkEntry.entry { ... }
]|

While any widget defining an “entry” class would be
matched by:
|[
.entry { ... }
]|
*/
func (self *_TraitStyleContext) AddClass(class_name string) {
	__cgo__class_name := (*C.gchar)(unsafe.Pointer(C.CString(class_name)))
	C.gtk_style_context_add_class(self.CPointer, __cgo__class_name)
	C.free(unsafe.Pointer(__cgo__class_name))
	return
}

/*
Adds a style provider to @context, to be used in style construction.
Note that a style provider added by this function only affects
the style of the widget to which @context belongs. If you want
to affect the style of all widgets, use
gtk_style_context_add_provider_for_screen().

Note: If both priorities are the same, a #GtkStyleProvider
added through this function takes precedence over another added
through gtk_style_context_add_provider_for_screen().
*/
func (self *_TraitStyleContext) AddProvider(provider *C.GtkStyleProvider, priority uint) {
	C.gtk_style_context_add_provider(self.CPointer, provider, C.guint(priority))
	return
}

/*
Adds a region to @context, so posterior calls to
gtk_style_context_get() or any of the gtk_render_*()
functions will make use of this new region for styling.

In the CSS file format, a #GtkTreeView defining a “row”
region, would be matched by:

|[
GtkTreeView row { ... }
]|

Pseudo-classes are used for matching @flags, so the two
following rules:
|[
GtkTreeView row:nth-child(even) { ... }
GtkTreeView row:nth-child(odd) { ... }
]|

would apply to even and odd rows, respectively.

Region names must only contain lowercase letters
and “-”, starting always with a lowercase letter.
*/
func (self *_TraitStyleContext) AddRegion(region_name string, flags C.GtkRegionFlags) {
	__cgo__region_name := (*C.gchar)(unsafe.Pointer(C.CString(region_name)))
	C.gtk_style_context_add_region(self.CPointer, __cgo__region_name, flags)
	C.free(unsafe.Pointer(__cgo__region_name))
	return
}

// gtk_style_context_cancel_animations is not generated due to deprecation attr

// gtk_style_context_get is not generated due to varargs

/*
Gets the background color for a given state.
*/
func (self *_TraitStyleContext) GetBackgroundColor(state C.GtkStateFlags) (color C.GdkRGBA) {
	C.gtk_style_context_get_background_color(self.CPointer, state, &color)
	return
}

/*
Gets the border for a given state as a #GtkBorder.
See %GTK_STYLE_PROPERTY_BORDER_WIDTH.
*/
func (self *_TraitStyleContext) GetBorder(state C.GtkStateFlags) (border C.GtkBorder) {
	C.gtk_style_context_get_border(self.CPointer, state, &border)
	return
}

/*
Gets the border color for a given state.
*/
func (self *_TraitStyleContext) GetBorderColor(state C.GtkStateFlags) (color C.GdkRGBA) {
	C.gtk_style_context_get_border_color(self.CPointer, state, &color)
	return
}

/*
Gets the foreground color for a given state.
*/
func (self *_TraitStyleContext) GetColor(state C.GtkStateFlags) (color C.GdkRGBA) {
	C.gtk_style_context_get_color(self.CPointer, state, &color)
	return
}

// gtk_style_context_get_direction is not generated due to deprecation attr

// gtk_style_context_get_font is not generated due to deprecation attr

/*
Returns the #GdkFrameClock to which @context is attached.
*/
func (self *_TraitStyleContext) GetFrameClock() (return__ *C.GdkFrameClock) {
	return__ = C.gtk_style_context_get_frame_clock(self.CPointer)
	return
}

/*
Returns the sides where rendered elements connect visually with others.
*/
func (self *_TraitStyleContext) GetJunctionSides() (return__ C.GtkJunctionSides) {
	return__ = C.gtk_style_context_get_junction_sides(self.CPointer)
	return
}

/*
Gets the margin for a given state as a #GtkBorder.
See %GTK_STYLE_PROPERTY_MARGIN.
*/
func (self *_TraitStyleContext) GetMargin(state C.GtkStateFlags) (margin C.GtkBorder) {
	C.gtk_style_context_get_margin(self.CPointer, state, &margin)
	return
}

/*
Gets the padding for a given state as a #GtkBorder.
See %GTK_STYLE_PROPERTY_PADDING.
*/
func (self *_TraitStyleContext) GetPadding(state C.GtkStateFlags) (padding C.GtkBorder) {
	C.gtk_style_context_get_padding(self.CPointer, state, &padding)
	return
}

/*
Gets the parent context set via gtk_style_context_set_parent().
See that function for details.
*/
func (self *_TraitStyleContext) GetParent() (return__ *StyleContext) {
	var __cgo__return__ *C.GtkStyleContext
	__cgo__return__ = C.gtk_style_context_get_parent(self.CPointer)
	return__ = NewStyleContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the widget path used for style matching.
*/
func (self *_TraitStyleContext) GetPath() (return__ *WidgetPath) {
	var __cgo__return__ *C.GtkWidgetPath
	__cgo__return__ = C.gtk_style_context_get_path(self.CPointer)
	return__ = (*WidgetPath)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Gets a style property from @context for the given state.

When @value is no longer needed, g_value_unset() must be called
to free any allocated memory.
*/
func (self *_TraitStyleContext) GetProperty(property string, state C.GtkStateFlags) (value C.GValue) {
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	C.gtk_style_context_get_property(self.CPointer, __cgo__property, state, &value)
	C.free(unsafe.Pointer(__cgo__property))
	return
}

/*
Returns the scale used for assets.
*/
func (self *_TraitStyleContext) GetScale() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_style_context_get_scale(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the #GdkScreen to which @context is attached.
*/
func (self *_TraitStyleContext) GetScreen() (return__ *C.GdkScreen) {
	return__ = C.gtk_style_context_get_screen(self.CPointer)
	return
}

/*
Queries the location in the CSS where @property was defined for the
current @context. Note that the state to be queried is taken from
gtk_style_context_get_state().

If the location is not available, %NULL will be returned. The
location might not be available for various reasons, such as the
property being overridden, @property not naming a supported CSS
property or tracking of definitions being disabled for performance
reasons.

Shorthand CSS properties cannot be queried for a location and will
always return %NULL.
*/
func (self *_TraitStyleContext) GetSection(property string) (return__ *CssSection) {
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	var __cgo__return__ *C.GtkCssSection
	__cgo__return__ = C.gtk_style_context_get_section(self.CPointer, __cgo__property)
	C.free(unsafe.Pointer(__cgo__property))
	return__ = (*CssSection)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the state used when rendering.
*/
func (self *_TraitStyleContext) GetState() (return__ C.GtkStateFlags) {
	return__ = C.gtk_style_context_get_state(self.CPointer)
	return
}

// gtk_style_context_get_style is not generated due to varargs

/*
Gets the value for a widget style property.

When @value is no longer needed, g_value_unset() must be called
to free any allocated memory.
*/
func (self *_TraitStyleContext) GetStyleProperty(property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.gtk_style_context_get_style_property(self.CPointer, __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

// gtk_style_context_get_style_valist is not generated due to varargs

// gtk_style_context_get_valist is not generated due to varargs

/*
Returns %TRUE if @context currently has defined the
given class name
*/
func (self *_TraitStyleContext) HasClass(class_name string) (return__ bool) {
	__cgo__class_name := (*C.gchar)(unsafe.Pointer(C.CString(class_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_style_context_has_class(self.CPointer, __cgo__class_name)
	C.free(unsafe.Pointer(__cgo__class_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if @context has the region defined.
If @flags_return is not %NULL, it is set to the flags
affecting the region.
*/
func (self *_TraitStyleContext) HasRegion(region_name string) (flags_return C.GtkRegionFlags, return__ bool) {
	__cgo__region_name := (*C.gchar)(unsafe.Pointer(C.CString(region_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_style_context_has_region(self.CPointer, __cgo__region_name, &flags_return)
	C.free(unsafe.Pointer(__cgo__region_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_style_context_invalidate is not generated due to deprecation attr

/*
Returns the list of classes currently defined in @context.
*/
func (self *_TraitStyleContext) ListClasses() (return__ *C.GList) {
	return__ = C.gtk_style_context_list_classes(self.CPointer)
	return
}

/*
Returns the list of regions currently defined in @context.
*/
func (self *_TraitStyleContext) ListRegions() (return__ *C.GList) {
	return__ = C.gtk_style_context_list_regions(self.CPointer)
	return
}

/*
Looks up and resolves a color name in the @context color map.
*/
func (self *_TraitStyleContext) LookupColor(color_name string) (color C.GdkRGBA, return__ bool) {
	__cgo__color_name := (*C.gchar)(unsafe.Pointer(C.CString(color_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_style_context_lookup_color(self.CPointer, __cgo__color_name, &color)
	C.free(unsafe.Pointer(__cgo__color_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_style_context_lookup_icon_set is not generated due to deprecation attr

// gtk_style_context_notify_state_change is not generated due to deprecation attr

// gtk_style_context_pop_animatable_region is not generated due to deprecation attr

// gtk_style_context_push_animatable_region is not generated due to deprecation attr

/*
Removes @class_name from @context.
*/
func (self *_TraitStyleContext) RemoveClass(class_name string) {
	__cgo__class_name := (*C.gchar)(unsafe.Pointer(C.CString(class_name)))
	C.gtk_style_context_remove_class(self.CPointer, __cgo__class_name)
	C.free(unsafe.Pointer(__cgo__class_name))
	return
}

/*
Removes @provider from the style providers list in @context.
*/
func (self *_TraitStyleContext) RemoveProvider(provider *C.GtkStyleProvider) {
	C.gtk_style_context_remove_provider(self.CPointer, provider)
	return
}

/*
Removes a region from @context.
*/
func (self *_TraitStyleContext) RemoveRegion(region_name string) {
	__cgo__region_name := (*C.gchar)(unsafe.Pointer(C.CString(region_name)))
	C.gtk_style_context_remove_region(self.CPointer, __cgo__region_name)
	C.free(unsafe.Pointer(__cgo__region_name))
	return
}

/*
Restores @context state to a previous stage.
See gtk_style_context_save().
*/
func (self *_TraitStyleContext) Restore() {
	C.gtk_style_context_restore(self.CPointer)
	return
}

/*
Saves the @context state, so all modifications done through
gtk_style_context_add_class(), gtk_style_context_remove_class(),
gtk_style_context_add_region(), gtk_style_context_remove_region()
or gtk_style_context_set_junction_sides() can be reverted in one
go through gtk_style_context_restore().
*/
func (self *_TraitStyleContext) Save() {
	C.gtk_style_context_save(self.CPointer)
	return
}

// gtk_style_context_scroll_animations is not generated due to deprecation attr

/*
Sets the background of @window to the background pattern or
color specified in @context for its current state.
*/
func (self *_TraitStyleContext) SetBackground(window *C.GdkWindow) {
	C.gtk_style_context_set_background(self.CPointer, window)
	return
}

// gtk_style_context_set_direction is not generated due to deprecation attr

/*
Attaches @context to the given frame clock.

The frame clock is used for the timing of animations.

If you are using a #GtkStyleContext returned from
gtk_widget_get_style_context(), you do not need to
call this yourself.
*/
func (self *_TraitStyleContext) SetFrameClock(frame_clock *C.GdkFrameClock) {
	C.gtk_style_context_set_frame_clock(self.CPointer, frame_clock)
	return
}

/*
Sets the sides where rendered elements (mostly through
gtk_render_frame()) will visually connect with other visual elements.

This is merely a hint that may or may not be honored
by theming engines.

Container widgets are expected to set junction hints as appropriate
for their children, so it should not normally be necessary to call
this function manually.
*/
func (self *_TraitStyleContext) SetJunctionSides(sides C.GtkJunctionSides) {
	C.gtk_style_context_set_junction_sides(self.CPointer, sides)
	return
}

/*
Sets the parent style context for @context. The parent style
context is used to implement
[inheritance](http://www.w3.org/TR/css3-cascade/#inheritance)
of properties.

If you are using a #GtkStyleContext returned from
gtk_widget_get_style_context(), the parent will be set for you.
*/
func (self *_TraitStyleContext) SetParent(parent *StyleContext) {
	C.gtk_style_context_set_parent(self.CPointer, (*C.GtkStyleContext)(parent.CPointer))
	return
}

/*
Sets the #GtkWidgetPath used for style matching. As a
consequence, the style will be regenerated to match
the new given path.

If you are using a #GtkStyleContext returned from
gtk_widget_get_style_context(), you do not need to call
this yourself.
*/
func (self *_TraitStyleContext) SetPath(path *WidgetPath) {
	C.gtk_style_context_set_path(self.CPointer, (*C.GtkWidgetPath)(unsafe.Pointer(path)))
	return
}

/*
Sets the scale to use when getting image assets for the style .
*/
func (self *_TraitStyleContext) SetScale(scale int) {
	C.gtk_style_context_set_scale(self.CPointer, C.gint(scale))
	return
}

/*
Attaches @context to the given screen.

The screen is used to add style information from “global” style
providers, such as the screens #GtkSettings instance.

If you are using a #GtkStyleContext returned from
gtk_widget_get_style_context(), you do not need to
call this yourself.
*/
func (self *_TraitStyleContext) SetScreen(screen *C.GdkScreen) {
	C.gtk_style_context_set_screen(self.CPointer, screen)
	return
}

/*
Sets the state to be used when rendering with any
of the gtk_render_*() functions.
*/
func (self *_TraitStyleContext) SetState(flags C.GtkStateFlags) {
	C.gtk_style_context_set_state(self.CPointer, flags)
	return
}

// gtk_style_context_state_is_running is not generated due to deprecation attr

type _TraitStyleProperties struct{ CPointer *C.GtkStyleProperties }

/*
Clears all style information from @props.
*/
func (self *_TraitStyleProperties) Clear() {
	C.gtk_style_properties_clear(self.CPointer)
	return
}

// gtk_style_properties_get is not generated due to varargs

/*
Gets a style property from @props for the given state. When done with @value,
g_value_unset() needs to be called to free any allocated memory.
*/
func (self *_TraitStyleProperties) GetProperty(property string, state C.GtkStateFlags) (value C.GValue, return__ bool) {
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_style_properties_get_property(self.CPointer, __cgo__property, state, &value)
	C.free(unsafe.Pointer(__cgo__property))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_style_properties_get_valist is not generated due to varargs

// gtk_style_properties_lookup_color is not generated due to deprecation attr

// gtk_style_properties_map_color is not generated due to deprecation attr

/*
Merges into @props all the style information contained
in @props_to_merge. If @replace is %TRUE, the values
will be overwritten, if it is %FALSE, the older values
will prevail.
*/
func (self *_TraitStyleProperties) Merge(props_to_merge *StyleProperties, replace bool) {
	__cgo__replace := C.gboolean(0)
	if replace {
		__cgo__replace = C.gboolean(1)
	}
	C.gtk_style_properties_merge(self.CPointer, (*C.GtkStyleProperties)(props_to_merge.CPointer), __cgo__replace)
	return
}

// gtk_style_properties_set is not generated due to varargs

/*
Sets a styling property in @props.
*/
func (self *_TraitStyleProperties) SetProperty(property string, state C.GtkStateFlags, value *C.GValue) {
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	C.gtk_style_properties_set_property(self.CPointer, __cgo__property, state, value)
	C.free(unsafe.Pointer(__cgo__property))
	return
}

// gtk_style_properties_set_valist is not generated due to varargs

/*
Unsets a style property in @props.
*/
func (self *_TraitStyleProperties) UnsetProperty(property string, state C.GtkStateFlags) {
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	C.gtk_style_properties_unset_property(self.CPointer, __cgo__property, state)
	C.free(unsafe.Pointer(__cgo__property))
	return
}

type _TraitSwitch struct{ CPointer *C.GtkSwitch }

/*
Gets whether the #GtkSwitch is in its “on” or “off” state.
*/
func (self *_TraitSwitch) GetActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_switch_get_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Changes the state of @sw to the desired one.
*/
func (self *_TraitSwitch) SetActive(is_active bool) {
	__cgo__is_active := C.gboolean(0)
	if is_active {
		__cgo__is_active = C.gboolean(1)
	}
	C.gtk_switch_set_active(self.CPointer, __cgo__is_active)
	return
}

type _TraitTable struct{ CPointer *C.GtkTable }

// gtk_table_attach is not generated due to deprecation attr

// gtk_table_attach_defaults is not generated due to deprecation attr

// gtk_table_get_col_spacing is not generated due to deprecation attr

// gtk_table_get_default_col_spacing is not generated due to deprecation attr

// gtk_table_get_default_row_spacing is not generated due to deprecation attr

// gtk_table_get_homogeneous is not generated due to deprecation attr

// gtk_table_get_row_spacing is not generated due to deprecation attr

// gtk_table_get_size is not generated due to deprecation attr

// gtk_table_resize is not generated due to deprecation attr

// gtk_table_set_col_spacing is not generated due to deprecation attr

// gtk_table_set_col_spacings is not generated due to deprecation attr

// gtk_table_set_homogeneous is not generated due to deprecation attr

// gtk_table_set_row_spacing is not generated due to deprecation attr

// gtk_table_set_row_spacings is not generated due to deprecation attr

type _TraitTearoffMenuItem struct{ CPointer *C.GtkTearoffMenuItem }

type _TraitTextBuffer struct{ CPointer *C.GtkTextBuffer }

/*
Adds the mark at position @where. The mark must not be added to
another buffer, and if its name is not %NULL then there must not
be another mark in the buffer with the same name.

Emits the #GtkTextBuffer::mark-set signal as notification of the mark's
initial placement.
*/
func (self *_TraitTextBuffer) AddMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_add_mark(self.CPointer, (*C.GtkTextMark)(mark.CPointer), (*C.GtkTextIter)(unsafe.Pointer(where)))
	return
}

/*
Adds @clipboard to the list of clipboards in which the selection
contents of @buffer are available. In most cases, @clipboard will be
the #GtkClipboard of type %GDK_SELECTION_PRIMARY for a view of @buffer.
*/
func (self *_TraitTextBuffer) AddSelectionClipboard(clipboard *Clipboard) {
	C.gtk_text_buffer_add_selection_clipboard(self.CPointer, (*C.GtkClipboard)(clipboard.CPointer))
	return
}

/*
Emits the “apply-tag” signal on @buffer. The default
handler for the signal applies @tag to the given range.
@start and @end do not have to be in order.
*/
func (self *_TraitTextBuffer) ApplyTag(tag *TextTag, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_apply_tag(self.CPointer, (*C.GtkTextTag)(tag.CPointer), (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)))
	return
}

/*
Calls gtk_text_tag_table_lookup() on the buffer’s tag table to
get a #GtkTextTag, then calls gtk_text_buffer_apply_tag().
*/
func (self *_TraitTextBuffer) ApplyTagByName(name string, start *TextIter, end *TextIter) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_text_buffer_apply_tag_by_name(self.CPointer, __cgo__name, (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)))
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Performs the appropriate action as if the user hit the delete
key with the cursor at the position specified by @iter. In the
normal case a single character will be deleted, but when
combining accents are involved, more than one character can
be deleted, and when precomposed character and accent combinations
are involved, less than one character will be deleted.

Because the buffer is modified, all outstanding iterators become
invalid after calling this function; however, the @iter will be
re-initialized to point to the location where text was deleted.
*/
func (self *_TraitTextBuffer) Backspace(iter *TextIter, interactive bool, default_editable bool) (return__ bool) {
	__cgo__interactive := C.gboolean(0)
	if interactive {
		__cgo__interactive = C.gboolean(1)
	}
	__cgo__default_editable := C.gboolean(0)
	if default_editable {
		__cgo__default_editable = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_backspace(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), __cgo__interactive, __cgo__default_editable)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Called to indicate that the buffer operations between here and a
call to gtk_text_buffer_end_user_action() are part of a single
user-visible operation. The operations between
gtk_text_buffer_begin_user_action() and
gtk_text_buffer_end_user_action() can then be grouped when creating
an undo stack. #GtkTextBuffer maintains a count of calls to
gtk_text_buffer_begin_user_action() that have not been closed with
a call to gtk_text_buffer_end_user_action(), and emits the
“begin-user-action” and “end-user-action” signals only for the
outermost pair of calls. This allows you to build user actions
from other user actions.

The “interactive” buffer mutation functions, such as
gtk_text_buffer_insert_interactive(), automatically call begin/end
user action around the buffer operations they perform, so there's
no need to add extra calls if you user action consists solely of a
single call to one of those functions.
*/
func (self *_TraitTextBuffer) BeginUserAction() {
	C.gtk_text_buffer_begin_user_action(self.CPointer)
	return
}

/*
Copies the currently-selected text to a clipboard.
*/
func (self *_TraitTextBuffer) CopyClipboard(clipboard *Clipboard) {
	C.gtk_text_buffer_copy_clipboard(self.CPointer, (*C.GtkClipboard)(clipboard.CPointer))
	return
}

/*
This is a convenience function which simply creates a child anchor
with gtk_text_child_anchor_new() and inserts it into the buffer
with gtk_text_buffer_insert_child_anchor(). The new anchor is
owned by the buffer; no reference count is returned to
the caller of gtk_text_buffer_create_child_anchor().
*/
func (self *_TraitTextBuffer) CreateChildAnchor(iter *TextIter) (return__ *TextChildAnchor) {
	var __cgo__return__ *C.GtkTextChildAnchor
	__cgo__return__ = C.gtk_text_buffer_create_child_anchor(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)))
	return__ = NewTextChildAnchorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Creates a mark at position @where. If @mark_name is %NULL, the mark
is anonymous; otherwise, the mark can be retrieved by name using
gtk_text_buffer_get_mark(). If a mark has left gravity, and text is
inserted at the mark’s current location, the mark will be moved to
the left of the newly-inserted text. If the mark has right gravity
(@left_gravity = %FALSE), the mark will end up on the right of
newly-inserted text. The standard left-to-right cursor is a mark
with right gravity (when you type, the cursor stays on the right
side of the text you’re typing).

The caller of this function does not own a
reference to the returned #GtkTextMark, so you can ignore the
return value if you like. Marks are owned by the buffer and go
away when the buffer does.

Emits the #GtkTextBuffer::mark-set signal as notification of the mark's
initial placement.
*/
func (self *_TraitTextBuffer) CreateMark(mark_name string, where *TextIter, left_gravity bool) (return__ *TextMark) {
	__cgo__mark_name := (*C.gchar)(unsafe.Pointer(C.CString(mark_name)))
	__cgo__left_gravity := C.gboolean(0)
	if left_gravity {
		__cgo__left_gravity = C.gboolean(1)
	}
	var __cgo__return__ *C.GtkTextMark
	__cgo__return__ = C.gtk_text_buffer_create_mark(self.CPointer, __cgo__mark_name, (*C.GtkTextIter)(unsafe.Pointer(where)), __cgo__left_gravity)
	C.free(unsafe.Pointer(__cgo__mark_name))
	return__ = NewTextMarkFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

// gtk_text_buffer_create_tag is not generated due to varargs

/*
Copies the currently-selected text to a clipboard, then deletes
said text if it’s editable.
*/
func (self *_TraitTextBuffer) CutClipboard(clipboard *Clipboard, default_editable bool) {
	__cgo__default_editable := C.gboolean(0)
	if default_editable {
		__cgo__default_editable = C.gboolean(1)
	}
	C.gtk_text_buffer_cut_clipboard(self.CPointer, (*C.GtkClipboard)(clipboard.CPointer), __cgo__default_editable)
	return
}

/*
Deletes text between @start and @end. The order of @start and @end
is not actually relevant; gtk_text_buffer_delete() will reorder
them. This function actually emits the “delete-range” signal, and
the default handler of that signal deletes the text. Because the
buffer is modified, all outstanding iterators become invalid after
calling this function; however, the @start and @end will be
re-initialized to point to the location where text was deleted.
*/
func (self *_TraitTextBuffer) Delete(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_delete(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)))
	return
}

/*
Deletes all editable text in the given range.
Calls gtk_text_buffer_delete() for each editable sub-range of
[@start,@end). @start and @end are revalidated to point to
the location of the last deleted range, or left untouched if
no text was deleted.
*/
func (self *_TraitTextBuffer) DeleteInteractive(start_iter *TextIter, end_iter *TextIter, default_editable bool) (return__ bool) {
	__cgo__default_editable := C.gboolean(0)
	if default_editable {
		__cgo__default_editable = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_delete_interactive(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(start_iter)), (*C.GtkTextIter)(unsafe.Pointer(end_iter)), __cgo__default_editable)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Deletes @mark, so that it’s no longer located anywhere in the
buffer. Removes the reference the buffer holds to the mark, so if
you haven’t called g_object_ref() on the mark, it will be freed. Even
if the mark isn’t freed, most operations on @mark become
invalid, until it gets added to a buffer again with
gtk_text_buffer_add_mark(). Use gtk_text_mark_get_deleted() to
find out if a mark has been removed from its buffer.
The #GtkTextBuffer::mark-deleted signal will be emitted as notification after
the mark is deleted.
*/
func (self *_TraitTextBuffer) DeleteMark(mark *TextMark) {
	C.gtk_text_buffer_delete_mark(self.CPointer, (*C.GtkTextMark)(mark.CPointer))
	return
}

/*
Deletes the mark named @name; the mark must exist. See
gtk_text_buffer_delete_mark() for details.
*/
func (self *_TraitTextBuffer) DeleteMarkByName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_text_buffer_delete_mark_by_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Deletes the range between the “insert” and “selection_bound” marks,
that is, the currently-selected text. If @interactive is %TRUE,
the editability of the selection will be considered (users can’t delete
uneditable text).
*/
func (self *_TraitTextBuffer) DeleteSelection(interactive bool, default_editable bool) (return__ bool) {
	__cgo__interactive := C.gboolean(0)
	if interactive {
		__cgo__interactive = C.gboolean(1)
	}
	__cgo__default_editable := C.gboolean(0)
	if default_editable {
		__cgo__default_editable = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_delete_selection(self.CPointer, __cgo__interactive, __cgo__default_editable)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function deserializes rich text in format @format and inserts
it at @iter.

@formats to be used must be registered using
gtk_text_buffer_register_deserialize_format() or
gtk_text_buffer_register_deserialize_tagset() beforehand.
*/
func (self *_TraitTextBuffer) Deserialize(content_buffer *TextBuffer, format C.GdkAtom, iter *TextIter, data []byte, length int64) (return__ bool, __err__ error) {
	__header__data := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_deserialize(self.CPointer, (*C.GtkTextBuffer)(content_buffer.CPointer), format, (*C.GtkTextIter)(unsafe.Pointer(iter)), (*C.guint8)(unsafe.Pointer(__header__data.Data)), C.gsize(length), &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
This functions returns the value set with
gtk_text_buffer_deserialize_set_can_create_tags()
*/
func (self *_TraitTextBuffer) DeserializeGetCanCreateTags(format C.GdkAtom) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_deserialize_get_can_create_tags(self.CPointer, format)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Use this function to allow a rich text deserialization function to
create new tags in the receiving buffer. Note that using this
function is almost always a bad idea, because the rich text
functions you register should know how to map the rich text format
they handler to your text buffers set of tags.

The ability of creating new (arbitrary!) tags in the receiving buffer
is meant for special rich text formats like the internal one that
is registered using gtk_text_buffer_register_deserialize_tagset(),
because that format is essentially a dump of the internal structure
of the source buffer, including its tag names.

You should allow creation of tags only if you know what you are
doing, e.g. if you defined a tagset name for your application
suite’s text buffers and you know that it’s fine to receive new
tags from these buffers, because you know that your application can
handle the newly created tags.
*/
func (self *_TraitTextBuffer) DeserializeSetCanCreateTags(format C.GdkAtom, can_create_tags bool) {
	__cgo__can_create_tags := C.gboolean(0)
	if can_create_tags {
		__cgo__can_create_tags = C.gboolean(1)
	}
	C.gtk_text_buffer_deserialize_set_can_create_tags(self.CPointer, format, __cgo__can_create_tags)
	return
}

/*
Should be paired with a call to gtk_text_buffer_begin_user_action().
See that function for a full explanation.
*/
func (self *_TraitTextBuffer) EndUserAction() {
	C.gtk_text_buffer_end_user_action(self.CPointer)
	return
}

/*
Retrieves the first and last iterators in the buffer, i.e. the
entire buffer lies within the range [@start,@end).
*/
func (self *_TraitTextBuffer) GetBounds() (start C.GtkTextIter, end C.GtkTextIter) {
	C.gtk_text_buffer_get_bounds(self.CPointer, &start, &end)
	return
}

/*
Gets the number of characters in the buffer; note that characters
and bytes are not the same, you can’t e.g. expect the contents of
the buffer in string form to be this many bytes long. The character
count is cached, so this function is very fast.
*/
func (self *_TraitTextBuffer) GetCharCount() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_buffer_get_char_count(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
This function returns the list of targets this text buffer can
provide for copying and as DND source. The targets in the list are
added with @info values from the #GtkTextBufferTargetInfo enum,
using gtk_target_list_add_rich_text_targets() and
gtk_target_list_add_text_targets().
*/
func (self *_TraitTextBuffer) GetCopyTargetList() (return__ *TargetList) {
	var __cgo__return__ *C.GtkTargetList
	__cgo__return__ = C.gtk_text_buffer_get_copy_target_list(self.CPointer)
	return__ = (*TargetList)(unsafe.Pointer(__cgo__return__))
	return
}

/*
This function returns the rich text deserialize formats registered
with @buffer using gtk_text_buffer_register_deserialize_format() or
gtk_text_buffer_register_deserialize_tagset()
*/
func (self *_TraitTextBuffer) GetDeserializeFormats() (n_formats int, return__ *C.GdkAtom) {
	var __cgo__n_formats C.gint
	return__ = C.gtk_text_buffer_get_deserialize_formats(self.CPointer, &__cgo__n_formats)
	n_formats = int(__cgo__n_formats)
	return
}

/*
Initializes @iter with the “end iterator,” one past the last valid
character in the text buffer. If dereferenced with
gtk_text_iter_get_char(), the end iterator has a character value of 0.
The entire buffer lies in the range from the first position in
the buffer (call gtk_text_buffer_get_start_iter() to get
character position 0) to the end iterator.
*/
func (self *_TraitTextBuffer) GetEndIter() (iter C.GtkTextIter) {
	C.gtk_text_buffer_get_end_iter(self.CPointer, &iter)
	return
}

/*
Indicates whether the buffer has some text currently selected.
*/
func (self *_TraitTextBuffer) GetHasSelection() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_get_has_selection(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the mark that represents the cursor (insertion point).
Equivalent to calling gtk_text_buffer_get_mark() to get the mark
named “insert”, but very slightly more efficient, and involves less
typing.
*/
func (self *_TraitTextBuffer) GetInsert() (return__ *TextMark) {
	var __cgo__return__ *C.GtkTextMark
	__cgo__return__ = C.gtk_text_buffer_get_insert(self.CPointer)
	return__ = NewTextMarkFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Obtains the location of @anchor within @buffer.
*/
func (self *_TraitTextBuffer) GetIterAtChildAnchor(anchor *TextChildAnchor) (iter C.GtkTextIter) {
	C.gtk_text_buffer_get_iter_at_child_anchor(self.CPointer, &iter, (*C.GtkTextChildAnchor)(anchor.CPointer))
	return
}

/*
Initializes @iter to the start of the given line. If @line_number is greater
than the number of lines in the @buffer, the end iterator is returned.
*/
func (self *_TraitTextBuffer) GetIterAtLine(line_number int) (iter C.GtkTextIter) {
	C.gtk_text_buffer_get_iter_at_line(self.CPointer, &iter, C.gint(line_number))
	return
}

/*
Obtains an iterator pointing to @byte_index within the given line.
@byte_index must be the start of a UTF-8 character, and must not be
beyond the end of the line.  Note bytes, not
characters; UTF-8 may encode one character as multiple bytes.
*/
func (self *_TraitTextBuffer) GetIterAtLineIndex(line_number int, byte_index int) (iter C.GtkTextIter) {
	C.gtk_text_buffer_get_iter_at_line_index(self.CPointer, &iter, C.gint(line_number), C.gint(byte_index))
	return
}

/*
Obtains an iterator pointing to @char_offset within the given
line. The @char_offset must exist, offsets off the end of the line
are not allowed. Note characters, not bytes;
UTF-8 may encode one character as multiple bytes.
*/
func (self *_TraitTextBuffer) GetIterAtLineOffset(line_number int, char_offset int) (iter C.GtkTextIter) {
	C.gtk_text_buffer_get_iter_at_line_offset(self.CPointer, &iter, C.gint(line_number), C.gint(char_offset))
	return
}

/*
Initializes @iter with the current position of @mark.
*/
func (self *_TraitTextBuffer) GetIterAtMark(mark *TextMark) (iter C.GtkTextIter) {
	C.gtk_text_buffer_get_iter_at_mark(self.CPointer, &iter, (*C.GtkTextMark)(mark.CPointer))
	return
}

/*
Initializes @iter to a position @char_offset chars from the start
of the entire buffer. If @char_offset is -1 or greater than the number
of characters in the buffer, @iter is initialized to the end iterator,
the iterator one past the last valid character in the buffer.
*/
func (self *_TraitTextBuffer) GetIterAtOffset(char_offset int) (iter C.GtkTextIter) {
	C.gtk_text_buffer_get_iter_at_offset(self.CPointer, &iter, C.gint(char_offset))
	return
}

/*
Obtains the number of lines in the buffer. This value is cached, so
the function is very fast.
*/
func (self *_TraitTextBuffer) GetLineCount() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_buffer_get_line_count(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the mark named @name in buffer @buffer, or %NULL if no such
mark exists in the buffer.
*/
func (self *_TraitTextBuffer) GetMark(name string) (return__ *TextMark) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ *C.GtkTextMark
	__cgo__return__ = C.gtk_text_buffer_get_mark(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = NewTextMarkFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Indicates whether the buffer has been modified since the last call
to gtk_text_buffer_set_modified() set the modification flag to
%FALSE. Used for example to enable a “save” function in a text
editor.
*/
func (self *_TraitTextBuffer) GetModified() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_get_modified(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function returns the list of targets this text buffer supports
for pasting and as DND destination. The targets in the list are
added with @info values from the #GtkTextBufferTargetInfo enum,
using gtk_target_list_add_rich_text_targets() and
gtk_target_list_add_text_targets().
*/
func (self *_TraitTextBuffer) GetPasteTargetList() (return__ *TargetList) {
	var __cgo__return__ *C.GtkTargetList
	__cgo__return__ = C.gtk_text_buffer_get_paste_target_list(self.CPointer)
	return__ = (*TargetList)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the mark that represents the selection bound.  Equivalent
to calling gtk_text_buffer_get_mark() to get the mark named
“selection_bound”, but very slightly more efficient, and involves
less typing.

The currently-selected text in @buffer is the region between the
“selection_bound” and “insert” marks. If “selection_bound” and
“insert” are in the same place, then there is no current selection.
gtk_text_buffer_get_selection_bounds() is another convenient function
for handling the selection, if you just want to know whether there’s a
selection and what its bounds are.
*/
func (self *_TraitTextBuffer) GetSelectionBound() (return__ *TextMark) {
	var __cgo__return__ *C.GtkTextMark
	__cgo__return__ = C.gtk_text_buffer_get_selection_bound(self.CPointer)
	return__ = NewTextMarkFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns %TRUE if some text is selected; places the bounds
of the selection in @start and @end (if the selection has length 0,
then @start and @end are filled in with the same value).
@start and @end will be in ascending order. If @start and @end are
NULL, then they are not filled in, but the return value still indicates
whether text is selected.
*/
func (self *_TraitTextBuffer) GetSelectionBounds() (start C.GtkTextIter, end C.GtkTextIter, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_get_selection_bounds(self.CPointer, &start, &end)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function returns the rich text serialize formats registered
with @buffer using gtk_text_buffer_register_serialize_format() or
gtk_text_buffer_register_serialize_tagset()
*/
func (self *_TraitTextBuffer) GetSerializeFormats() (n_formats int, return__ *C.GdkAtom) {
	var __cgo__n_formats C.gint
	return__ = C.gtk_text_buffer_get_serialize_formats(self.CPointer, &__cgo__n_formats)
	n_formats = int(__cgo__n_formats)
	return
}

/*
Returns the text in the range [@start,@end). Excludes undisplayed
text (text marked with tags that set the invisibility attribute) if
@include_hidden_chars is %FALSE. The returned string includes a
0xFFFC character whenever the buffer contains
embedded images, so byte and character indexes into
the returned string do correspond to byte
and character indexes into the buffer. Contrast with
gtk_text_buffer_get_text(). Note that 0xFFFC can occur in normal
text as well, so it is not a reliable indicator that a pixbuf or
widget is in the buffer.
*/
func (self *_TraitTextBuffer) GetSlice(start *TextIter, end *TextIter, include_hidden_chars bool) (return__ string) {
	__cgo__include_hidden_chars := C.gboolean(0)
	if include_hidden_chars {
		__cgo__include_hidden_chars = C.gboolean(1)
	}
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_text_buffer_get_slice(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)), __cgo__include_hidden_chars)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Initialized @iter with the first position in the text buffer. This
is the same as using gtk_text_buffer_get_iter_at_offset() to get
the iter at character offset 0.
*/
func (self *_TraitTextBuffer) GetStartIter() (iter C.GtkTextIter) {
	C.gtk_text_buffer_get_start_iter(self.CPointer, &iter)
	return
}

/*
Get the #GtkTextTagTable associated with this buffer.
*/
func (self *_TraitTextBuffer) GetTagTable() (return__ *TextTagTable) {
	var __cgo__return__ *C.GtkTextTagTable
	__cgo__return__ = C.gtk_text_buffer_get_tag_table(self.CPointer)
	return__ = NewTextTagTableFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the text in the range [@start,@end). Excludes undisplayed
text (text marked with tags that set the invisibility attribute) if
@include_hidden_chars is %FALSE. Does not include characters
representing embedded images, so byte and character indexes into
the returned string do not correspond to byte
and character indexes into the buffer. Contrast with
gtk_text_buffer_get_slice().
*/
func (self *_TraitTextBuffer) GetText(start *TextIter, end *TextIter, include_hidden_chars bool) (return__ string) {
	__cgo__include_hidden_chars := C.gboolean(0)
	if include_hidden_chars {
		__cgo__include_hidden_chars = C.gboolean(1)
	}
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_text_buffer_get_text(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)), __cgo__include_hidden_chars)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Inserts @len bytes of @text at position @iter.  If @len is -1,
@text must be nul-terminated and will be inserted in its
entirety. Emits the “insert-text” signal; insertion actually occurs
in the default handler for the signal. @iter is invalidated when
insertion occurs (because the buffer contents change), but the
default signal handler revalidates it to point to the end of the
inserted text.
*/
func (self *_TraitTextBuffer) Insert(iter *TextIter, text string, len_ int) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_text_buffer_insert(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), __cgo__text, C.gint(len_))
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Simply calls gtk_text_buffer_insert(), using the current
cursor position as the insertion point.
*/
func (self *_TraitTextBuffer) InsertAtCursor(text string, len_ int) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_text_buffer_insert_at_cursor(self.CPointer, __cgo__text, C.gint(len_))
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Inserts a child widget anchor into the text buffer at @iter. The
anchor will be counted as one character in character counts, and
when obtaining the buffer contents as a string, will be represented
by the Unicode “object replacement character” 0xFFFC. Note that the
“slice” variants for obtaining portions of the buffer as a string
include this character for child anchors, but the “text” variants do
not. E.g. see gtk_text_buffer_get_slice() and
gtk_text_buffer_get_text(). Consider
gtk_text_buffer_create_child_anchor() as a more convenient
alternative to this function. The buffer will add a reference to
the anchor, so you can unref it after insertion.
*/
func (self *_TraitTextBuffer) InsertChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	C.gtk_text_buffer_insert_child_anchor(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), (*C.GtkTextChildAnchor)(anchor.CPointer))
	return
}

/*
Like gtk_text_buffer_insert(), but the insertion will not occur if
@iter is at a non-editable location in the buffer. Usually you
want to prevent insertions at ineditable locations if the insertion
results from a user action (is interactive).

@default_editable indicates the editability of text that doesn't
have a tag affecting editability applied to it. Typically the
result of gtk_text_view_get_editable() is appropriate here.
*/
func (self *_TraitTextBuffer) InsertInteractive(iter *TextIter, text string, len_ int, default_editable bool) (return__ bool) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	__cgo__default_editable := C.gboolean(0)
	if default_editable {
		__cgo__default_editable = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_insert_interactive(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), __cgo__text, C.gint(len_), __cgo__default_editable)
	C.free(unsafe.Pointer(__cgo__text))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Calls gtk_text_buffer_insert_interactive() at the cursor
position.

@default_editable indicates the editability of text that doesn't
have a tag affecting editability applied to it. Typically the
result of gtk_text_view_get_editable() is appropriate here.
*/
func (self *_TraitTextBuffer) InsertInteractiveAtCursor(text string, len_ int, default_editable bool) (return__ bool) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	__cgo__default_editable := C.gboolean(0)
	if default_editable {
		__cgo__default_editable = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_insert_interactive_at_cursor(self.CPointer, __cgo__text, C.gint(len_), __cgo__default_editable)
	C.free(unsafe.Pointer(__cgo__text))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Inserts an image into the text buffer at @iter. The image will be
counted as one character in character counts, and when obtaining
the buffer contents as a string, will be represented by the Unicode
“object replacement character” 0xFFFC. Note that the “slice”
variants for obtaining portions of the buffer as a string include
this character for pixbufs, but the “text” variants do
not. e.g. see gtk_text_buffer_get_slice() and
gtk_text_buffer_get_text().
*/
func (self *_TraitTextBuffer) InsertPixbuf(iter *TextIter, pixbuf *C.GdkPixbuf) {
	C.gtk_text_buffer_insert_pixbuf(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), pixbuf)
	return
}

/*
Copies text, tags, and pixbufs between @start and @end (the order
of @start and @end doesn’t matter) and inserts the copy at @iter.
Used instead of simply getting/inserting text because it preserves
images and tags. If @start and @end are in a different buffer from
@buffer, the two buffers must share the same tag table.

Implemented via emissions of the insert_text and apply_tag signals,
so expect those.
*/
func (self *_TraitTextBuffer) InsertRange(iter *TextIter, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_insert_range(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)))
	return
}

/*
Same as gtk_text_buffer_insert_range(), but does nothing if the
insertion point isn’t editable. The @default_editable parameter
indicates whether the text is editable at @iter if no tags
enclosing @iter affect editability. Typically the result of
gtk_text_view_get_editable() is appropriate here.
*/
func (self *_TraitTextBuffer) InsertRangeInteractive(iter *TextIter, start *TextIter, end *TextIter, default_editable bool) (return__ bool) {
	__cgo__default_editable := C.gboolean(0)
	if default_editable {
		__cgo__default_editable = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_buffer_insert_range_interactive(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)), __cgo__default_editable)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_text_buffer_insert_with_tags is not generated due to varargs

// gtk_text_buffer_insert_with_tags_by_name is not generated due to varargs

/*
Moves @mark to the new location @where. Emits the #GtkTextBuffer::mark-set
signal as notification of the move.
*/
func (self *_TraitTextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_move_mark(self.CPointer, (*C.GtkTextMark)(mark.CPointer), (*C.GtkTextIter)(unsafe.Pointer(where)))
	return
}

/*
Moves the mark named @name (which must exist) to location @where.
See gtk_text_buffer_move_mark() for details.
*/
func (self *_TraitTextBuffer) MoveMarkByName(name string, where *TextIter) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_text_buffer_move_mark_by_name(self.CPointer, __cgo__name, (*C.GtkTextIter)(unsafe.Pointer(where)))
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Pastes the contents of a clipboard. If @override_location is %NULL, the
pasted text will be inserted at the cursor position, or the buffer selection
will be replaced if the selection is non-empty.

Note: pasting is asynchronous, that is, we’ll ask for the paste data and
return, and at some point later after the main loop runs, the paste data will
be inserted.
*/
func (self *_TraitTextBuffer) PasteClipboard(clipboard *Clipboard, override_location *TextIter, default_editable bool) {
	__cgo__default_editable := C.gboolean(0)
	if default_editable {
		__cgo__default_editable = C.gboolean(1)
	}
	C.gtk_text_buffer_paste_clipboard(self.CPointer, (*C.GtkClipboard)(clipboard.CPointer), (*C.GtkTextIter)(unsafe.Pointer(override_location)), __cgo__default_editable)
	return
}

/*
This function moves the “insert” and “selection_bound” marks
simultaneously.  If you move them to the same place in two steps
with gtk_text_buffer_move_mark(), you will temporarily select a
region in between their old and new locations, which can be pretty
inefficient since the temporarily-selected region will force stuff
to be recalculated. This function moves them as a unit, which can
be optimized.
*/
func (self *_TraitTextBuffer) PlaceCursor(where *TextIter) {
	C.gtk_text_buffer_place_cursor(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(where)))
	return
}

/*
This function registers a rich text deserialization @function along with
its @mime_type with the passed @buffer.
*/
func (self *_TraitTextBuffer) RegisterDeserializeFormat(mime_type string, function C.GtkTextBufferDeserializeFunc, user_data unsafe.Pointer, user_data_destroy C.GDestroyNotify) (return__ C.GdkAtom) {
	__cgo__mime_type := (*C.gchar)(unsafe.Pointer(C.CString(mime_type)))
	return__ = C.gtk_text_buffer_register_deserialize_format(self.CPointer, __cgo__mime_type, function, (C.gpointer)(user_data), user_data_destroy)
	C.free(unsafe.Pointer(__cgo__mime_type))
	return
}

/*
This function registers GTK+’s internal rich text serialization
format with the passed @buffer. See
gtk_text_buffer_register_serialize_tagset() for details.
*/
func (self *_TraitTextBuffer) RegisterDeserializeTagset(tagset_name string) (return__ C.GdkAtom) {
	__cgo__tagset_name := (*C.gchar)(unsafe.Pointer(C.CString(tagset_name)))
	return__ = C.gtk_text_buffer_register_deserialize_tagset(self.CPointer, __cgo__tagset_name)
	C.free(unsafe.Pointer(__cgo__tagset_name))
	return
}

/*
This function registers a rich text serialization @function along with
its @mime_type with the passed @buffer.
*/
func (self *_TraitTextBuffer) RegisterSerializeFormat(mime_type string, function C.GtkTextBufferSerializeFunc, user_data unsafe.Pointer, user_data_destroy C.GDestroyNotify) (return__ C.GdkAtom) {
	__cgo__mime_type := (*C.gchar)(unsafe.Pointer(C.CString(mime_type)))
	return__ = C.gtk_text_buffer_register_serialize_format(self.CPointer, __cgo__mime_type, function, (C.gpointer)(user_data), user_data_destroy)
	C.free(unsafe.Pointer(__cgo__mime_type))
	return
}

/*
This function registers GTK+’s internal rich text serialization
format with the passed @buffer. The internal format does not comply
to any standard rich text format and only works between #GtkTextBuffer
instances. It is capable of serializing all of a text buffer’s tags
and embedded pixbufs.

This function is just a wrapper around
gtk_text_buffer_register_serialize_format(). The mime type used
for registering is “application/x-gtk-text-buffer-rich-text”, or
“application/x-gtk-text-buffer-rich-text;format=@tagset_name” if a
@tagset_name was passed.

The @tagset_name can be used to restrict the transfer of rich text
to buffers with compatible sets of tags, in order to avoid unknown
tags from being pasted. It is probably the common case to pass an
identifier != %NULL here, since the %NULL tagset requires the
receiving buffer to deal with with pasting of arbitrary tags.
*/
func (self *_TraitTextBuffer) RegisterSerializeTagset(tagset_name string) (return__ C.GdkAtom) {
	__cgo__tagset_name := (*C.gchar)(unsafe.Pointer(C.CString(tagset_name)))
	return__ = C.gtk_text_buffer_register_serialize_tagset(self.CPointer, __cgo__tagset_name)
	C.free(unsafe.Pointer(__cgo__tagset_name))
	return
}

/*
Removes all tags in the range between @start and @end.  Be careful
with this function; it could remove tags added in code unrelated to
the code you’re currently writing. That is, using this function is
probably a bad idea if you have two or more unrelated code sections
that add tags.
*/
func (self *_TraitTextBuffer) RemoveAllTags(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_remove_all_tags(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)))
	return
}

/*
Removes a #GtkClipboard added with
gtk_text_buffer_add_selection_clipboard().
*/
func (self *_TraitTextBuffer) RemoveSelectionClipboard(clipboard *Clipboard) {
	C.gtk_text_buffer_remove_selection_clipboard(self.CPointer, (*C.GtkClipboard)(clipboard.CPointer))
	return
}

/*
Emits the “remove-tag” signal. The default handler for the signal
removes all occurrences of @tag from the given range. @start and
@end don’t have to be in order.
*/
func (self *_TraitTextBuffer) RemoveTag(tag *TextTag, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_remove_tag(self.CPointer, (*C.GtkTextTag)(tag.CPointer), (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)))
	return
}

/*
Calls gtk_text_tag_table_lookup() on the buffer’s tag table to
get a #GtkTextTag, then calls gtk_text_buffer_remove_tag().
*/
func (self *_TraitTextBuffer) RemoveTagByName(name string, start *TextIter, end *TextIter) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_text_buffer_remove_tag_by_name(self.CPointer, __cgo__name, (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)))
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
This function moves the “insert” and “selection_bound” marks
simultaneously.  If you move them in two steps
with gtk_text_buffer_move_mark(), you will temporarily select a
region in between their old and new locations, which can be pretty
inefficient since the temporarily-selected region will force stuff
to be recalculated. This function moves them as a unit, which can
be optimized.
*/
func (self *_TraitTextBuffer) SelectRange(ins *TextIter, bound *TextIter) {
	C.gtk_text_buffer_select_range(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(ins)), (*C.GtkTextIter)(unsafe.Pointer(bound)))
	return
}

/*
This function serializes the portion of text between @start
and @end in the rich text format represented by @format.

@formats to be used must be registered using
gtk_text_buffer_register_serialize_format() or
gtk_text_buffer_register_serialize_tagset() beforehand.
*/
func (self *_TraitTextBuffer) Serialize(content_buffer *TextBuffer, format C.GdkAtom, start *TextIter, end *TextIter) (length int64, return__ []byte) {
	var __cgo__length C.gsize
	var __cgo__return__ *C.guint8
	__cgo__return__ = C.gtk_text_buffer_serialize(self.CPointer, (*C.GtkTextBuffer)(content_buffer.CPointer), format, (*C.GtkTextIter)(unsafe.Pointer(start)), (*C.GtkTextIter)(unsafe.Pointer(end)), &__cgo__length)
	length = int64(__cgo__length)
	defer func() { return__ = C.GoBytes(unsafe.Pointer(__cgo__return__), C.int(length)) }()
	return
}

/*
Used to keep track of whether the buffer has been modified since the
last time it was saved. Whenever the buffer is saved to disk, call
gtk_text_buffer_set_modified (@buffer, FALSE). When the buffer is modified,
it will automatically toggled on the modified bit again. When the modified
bit flips, the buffer emits the #GtkTextBuffer::modified-changed signal.
*/
func (self *_TraitTextBuffer) SetModified(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_text_buffer_set_modified(self.CPointer, __cgo__setting)
	return
}

/*
Deletes current contents of @buffer, and inserts @text instead. If
@len is -1, @text must be nul-terminated. @text must be valid UTF-8.
*/
func (self *_TraitTextBuffer) SetText(text string, len_ int) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_text_buffer_set_text(self.CPointer, __cgo__text, C.gint(len_))
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
This function unregisters a rich text format that was previously
registered using gtk_text_buffer_register_deserialize_format() or
gtk_text_buffer_register_deserialize_tagset().
*/
func (self *_TraitTextBuffer) UnregisterDeserializeFormat(format C.GdkAtom) {
	C.gtk_text_buffer_unregister_deserialize_format(self.CPointer, format)
	return
}

/*
This function unregisters a rich text format that was previously
registered using gtk_text_buffer_register_serialize_format() or
gtk_text_buffer_register_serialize_tagset()
*/
func (self *_TraitTextBuffer) UnregisterSerializeFormat(format C.GdkAtom) {
	C.gtk_text_buffer_unregister_serialize_format(self.CPointer, format)
	return
}

type _TraitTextChildAnchor struct{ CPointer *C.GtkTextChildAnchor }

/*
Determines whether a child anchor has been deleted from
the buffer. Keep in mind that the child anchor will be
unreferenced when removed from the buffer, so you need to
hold your own reference (with g_object_ref()) if you plan
to use this function &mdash; otherwise all deleted child anchors
will also be finalized.
*/
func (self *_TraitTextChildAnchor) GetDeleted() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_child_anchor_get_deleted(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets a list of all widgets anchored at this child anchor.
The returned list should be freed with g_list_free().
*/
func (self *_TraitTextChildAnchor) GetWidgets() (return__ *C.GList) {
	return__ = C.gtk_text_child_anchor_get_widgets(self.CPointer)
	return
}

type _TraitTextMark struct{ CPointer *C.GtkTextMark }

/*
Gets the buffer this mark is located inside,
or %NULL if the mark is deleted.
*/
func (self *_TraitTextMark) GetBuffer() (return__ *TextBuffer) {
	var __cgo__return__ *C.GtkTextBuffer
	__cgo__return__ = C.gtk_text_mark_get_buffer(self.CPointer)
	return__ = NewTextBufferFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns %TRUE if the mark has been removed from its buffer
with gtk_text_buffer_delete_mark(). See gtk_text_buffer_add_mark()
for a way to add it to a buffer again.
*/
func (self *_TraitTextMark) GetDeleted() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_mark_get_deleted(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether the mark has left gravity.
*/
func (self *_TraitTextMark) GetLeftGravity() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_mark_get_left_gravity(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the mark name; returns NULL for anonymous marks.
*/
func (self *_TraitTextMark) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_text_mark_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns %TRUE if the mark is visible (i.e. a cursor is displayed
for it).
*/
func (self *_TraitTextMark) GetVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_mark_get_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the visibility of @mark; the insertion point is normally
visible, i.e. you can see it as a vertical bar. Also, the text
widget uses a visible mark to indicate where a drop will occur when
dragging-and-dropping text. Most other marks are not visible.
Marks are not visible by default.
*/
func (self *_TraitTextMark) SetVisible(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_text_mark_set_visible(self.CPointer, __cgo__setting)
	return
}

type _TraitTextTag struct{ CPointer *C.GtkTextTag }

/*
Emits the “event” signal on the #GtkTextTag.
*/
func (self *_TraitTextTag) Event(event_object *C.GObject, event *C.GdkEvent, iter *TextIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_tag_event(self.CPointer, event_object, event, (*C.GtkTextIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the tag priority.
*/
func (self *_TraitTextTag) GetPriority() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_tag_get_priority(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Sets the priority of a #GtkTextTag. Valid priorities
start at 0 and go to one less than gtk_text_tag_table_get_size().
Each tag in a table has a unique priority; setting the priority
of one tag shifts the priorities of all the other tags in the
table to maintain a unique priority for each tag. Higher priority
tags “win” if two tags both set the same text attribute. When adding
a tag to a tag table, it will be assigned the highest priority in
the table by default; so normally the precedence of a set of tags
is the order in which they were added to the table, or created with
gtk_text_buffer_create_tag(), which adds the tag to the buffer’s table
automatically.
*/
func (self *_TraitTextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(self.CPointer, C.gint(priority))
	return
}

type _TraitTextTagTable struct{ CPointer *C.GtkTextTagTable }

/*
Add a tag to the table. The tag is assigned the highest priority
in the table.

@tag must not be in a tag table already, and may not have
the same name as an already-added tag.
*/
func (self *_TraitTextTagTable) Add(tag *TextTag) {
	C.gtk_text_tag_table_add(self.CPointer, (*C.GtkTextTag)(tag.CPointer))
	return
}

/*
Calls @func on each tag in @table, with user data @data.
Note that the table may not be modified while iterating
over it (you can’t add/remove tags).
*/
func (self *_TraitTextTagTable) Foreach(func_ C.GtkTextTagTableForeach, data unsafe.Pointer) {
	C.gtk_text_tag_table_foreach(self.CPointer, func_, (C.gpointer)(data))
	return
}

/*
Returns the size of the table (number of tags)
*/
func (self *_TraitTextTagTable) GetSize() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_tag_table_get_size(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Look up a named tag.
*/
func (self *_TraitTextTagTable) Lookup(name string) (return__ *TextTag) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ *C.GtkTextTag
	__cgo__return__ = C.gtk_text_tag_table_lookup(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return__ = NewTextTagFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Remove a tag from the table. If a #GtkTextBuffer has @table as its tag table,
the tag is removed from the buffer. The table’s reference to the tag is
removed, so the tag will end up destroyed if you don’t have a reference to
it.
*/
func (self *_TraitTextTagTable) Remove(tag *TextTag) {
	C.gtk_text_tag_table_remove(self.CPointer, (*C.GtkTextTag)(tag.CPointer))
	return
}

type _TraitTextView struct{ CPointer *C.GtkTextView }

/*
Adds a child widget in the text buffer, at the given @anchor.
*/
func (self *_TraitTextView) AddChildAtAnchor(child *Widget, anchor *TextChildAnchor) {
	C.gtk_text_view_add_child_at_anchor(self.CPointer, (*C.GtkWidget)(child.CPointer), (*C.GtkTextChildAnchor)(anchor.CPointer))
	return
}

/*
Adds a child at fixed coordinates in one of the text widget's
windows.

The window must have nonzero size (see
gtk_text_view_set_border_window_size()). Note that the child
coordinates are given relative to scrolling. When
placing a child in #GTK_TEXT_WINDOW_WIDGET, scrolling is
irrelevant, the child floats above all scrollable areas. But when
placing a child in one of the scrollable windows (border windows or
text window) it will move with the scrolling as needed.
*/
func (self *_TraitTextView) AddChildInWindow(child *Widget, which_window C.GtkTextWindowType, xpos int, ypos int) {
	C.gtk_text_view_add_child_in_window(self.CPointer, (*C.GtkWidget)(child.CPointer), which_window, C.gint(xpos), C.gint(ypos))
	return
}

/*
Moves the given @iter backward by one display (wrapped) line.
A display line is different from a paragraph. Paragraphs are
separated by newlines or other paragraph separator characters.
Display lines are created by line-wrapping a paragraph. If
wrapping is turned off, display lines and paragraphs will be the
same. Display lines are divided differently for each view, since
they depend on the view’s width; paragraphs are the same in all
views, since they depend on the contents of the #GtkTextBuffer.
*/
func (self *_TraitTextView) BackwardDisplayLine(iter *TextIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_backward_display_line(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Moves the given @iter backward to the next display line start.
A display line is different from a paragraph. Paragraphs are
separated by newlines or other paragraph separator characters.
Display lines are created by line-wrapping a paragraph. If
wrapping is turned off, display lines and paragraphs will be the
same. Display lines are divided differently for each view, since
they depend on the view’s width; paragraphs are the same in all
views, since they depend on the contents of the #GtkTextBuffer.
*/
func (self *_TraitTextView) BackwardDisplayLineStart(iter *TextIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_backward_display_line_start(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts coordinate (@buffer_x, @buffer_y) to coordinates for the window
@win, and stores the result in (@window_x, @window_y).

Note that you can’t convert coordinates for a nonexisting window (see
gtk_text_view_set_border_window_size()).
*/
func (self *_TraitTextView) BufferToWindowCoords(win C.GtkTextWindowType, buffer_x int, buffer_y int) (window_x int, window_y int) {
	var __cgo__window_x C.gint
	var __cgo__window_y C.gint
	C.gtk_text_view_buffer_to_window_coords(self.CPointer, win, C.gint(buffer_x), C.gint(buffer_y), &__cgo__window_x, &__cgo__window_y)
	window_x = int(__cgo__window_x)
	window_y = int(__cgo__window_y)
	return
}

/*
Moves the given @iter forward by one display (wrapped) line.
A display line is different from a paragraph. Paragraphs are
separated by newlines or other paragraph separator characters.
Display lines are created by line-wrapping a paragraph. If
wrapping is turned off, display lines and paragraphs will be the
same. Display lines are divided differently for each view, since
they depend on the view’s width; paragraphs are the same in all
views, since they depend on the contents of the #GtkTextBuffer.
*/
func (self *_TraitTextView) ForwardDisplayLine(iter *TextIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_forward_display_line(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Moves the given @iter forward to the next display line end.
A display line is different from a paragraph. Paragraphs are
separated by newlines or other paragraph separator characters.
Display lines are created by line-wrapping a paragraph. If
wrapping is turned off, display lines and paragraphs will be the
same. Display lines are divided differently for each view, since
they depend on the view’s width; paragraphs are the same in all
views, since they depend on the contents of the #GtkTextBuffer.
*/
func (self *_TraitTextView) ForwardDisplayLineEnd(iter *TextIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_forward_display_line_end(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether pressing the Tab key inserts a tab characters.
gtk_text_view_set_accepts_tab().
*/
func (self *_TraitTextView) GetAcceptsTab() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_get_accepts_tab(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the width of the specified border window. See
gtk_text_view_set_border_window_size().
*/
func (self *_TraitTextView) GetBorderWindowSize(type_ C.GtkTextWindowType) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_view_get_border_window_size(self.CPointer, type_)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the #GtkTextBuffer being displayed by this text view.
The reference count on the buffer is not incremented; the caller
of this function won’t own a new reference.
*/
func (self *_TraitTextView) GetBuffer() (return__ *TextBuffer) {
	var __cgo__return__ *C.GtkTextBuffer
	__cgo__return__ = C.gtk_text_view_get_buffer(self.CPointer)
	return__ = NewTextBufferFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Given an @iter within a text layout, determine the positions of the
strong and weak cursors if the insertion point is at that
iterator. The position of each cursor is stored as a zero-width
rectangle. The strong cursor location is the location where
characters of the directionality equal to the base direction of the
paragraph are inserted.  The weak cursor location is the location
where characters of the directionality opposite to the base
direction of the paragraph are inserted.

If @iter is %NULL, the actual cursor position is used.

Note that if @iter happens to be the actual cursor position, and
there is currently an IM preedit sequence being entered, the
returned locations will be adjusted to account for the preedit
cursor’s offset within the preedit sequence.

The rectangle position is in buffer coordinates; use
gtk_text_view_buffer_to_window_coords() to convert these
coordinates to coordinates for one of the windows in the text view.
*/
func (self *_TraitTextView) GetCursorLocations(iter *TextIter) (strong C.GdkRectangle, weak C.GdkRectangle) {
	C.gtk_text_view_get_cursor_locations(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), &strong, &weak)
	return
}

/*
Find out whether the cursor is being displayed.
*/
func (self *_TraitTextView) GetCursorVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_get_cursor_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Obtains a copy of the default text attributes. These are the
attributes used for text unless a tag overrides them.
You’d typically pass the default attributes in to
gtk_text_iter_get_attributes() in order to get the
attributes in effect at a given text position.

The return value is a copy owned by the caller of this function,
and should be freed.
*/
func (self *_TraitTextView) GetDefaultAttributes() (return__ *TextAttributes) {
	var __cgo__return__ *C.GtkTextAttributes
	__cgo__return__ = C.gtk_text_view_get_default_attributes(self.CPointer)
	return__ = (*TextAttributes)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the default editability of the #GtkTextView. Tags in the
buffer may override this setting for some ranges of text.
*/
func (self *_TraitTextView) GetEditable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_get_editable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_text_view_get_hadjustment is not generated due to deprecation attr

/*
Gets the default indentation of paragraphs in @text_view.
Tags in the view’s buffer may override the default.
The indentation may be negative.
*/
func (self *_TraitTextView) GetIndent() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_view_get_indent(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value of the #GtkTextView:input-hints property.
*/
func (self *_TraitTextView) GetInputHints() (return__ C.GtkInputHints) {
	return__ = C.gtk_text_view_get_input_hints(self.CPointer)
	return
}

/*
Gets the value of the #GtkTextView:input-purpose property.
*/
func (self *_TraitTextView) GetInputPurpose() (return__ C.GtkInputPurpose) {
	return__ = C.gtk_text_view_get_input_purpose(self.CPointer)
	return
}

/*
Retrieves the iterator at buffer coordinates @x and @y. Buffer
coordinates are coordinates for the entire buffer, not just the
currently-displayed portion.  If you have coordinates from an
event, you have to convert those to buffer coordinates with
gtk_text_view_window_to_buffer_coords().
*/
func (self *_TraitTextView) GetIterAtLocation(x int, y int) (iter C.GtkTextIter) {
	C.gtk_text_view_get_iter_at_location(self.CPointer, &iter, C.gint(x), C.gint(y))
	return
}

/*
Retrieves the iterator pointing to the character at buffer
coordinates @x and @y. Buffer coordinates are coordinates for
the entire buffer, not just the currently-displayed portion.
If you have coordinates from an event, you have to convert
those to buffer coordinates with
gtk_text_view_window_to_buffer_coords().

Note that this is different from gtk_text_view_get_iter_at_location(),
which returns cursor locations, i.e. positions between
characters.
*/
func (self *_TraitTextView) GetIterAtPosition(x int, y int) (iter C.GtkTextIter, trailing int) {
	var __cgo__trailing C.gint
	C.gtk_text_view_get_iter_at_position(self.CPointer, &iter, &__cgo__trailing, C.gint(x), C.gint(y))
	trailing = int(__cgo__trailing)
	return
}

/*
Gets a rectangle which roughly contains the character at @iter.
The rectangle position is in buffer coordinates; use
gtk_text_view_buffer_to_window_coords() to convert these
coordinates to coordinates for one of the windows in the text view.
*/
func (self *_TraitTextView) GetIterLocation(iter *TextIter) (location C.GdkRectangle) {
	C.gtk_text_view_get_iter_location(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), &location)
	return
}

/*
Gets the default justification of paragraphs in @text_view.
Tags in the buffer may override the default.
*/
func (self *_TraitTextView) GetJustification() (return__ C.GtkJustification) {
	return__ = C.gtk_text_view_get_justification(self.CPointer)
	return
}

/*
Gets the default left margin size of paragraphs in the @text_view.
Tags in the buffer may override the default.
*/
func (self *_TraitTextView) GetLeftMargin() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_view_get_left_margin(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the #GtkTextIter at the start of the line containing
the coordinate @y. @y is in buffer coordinates, convert from
window coordinates with gtk_text_view_window_to_buffer_coords().
If non-%NULL, @line_top will be filled with the coordinate of the top
edge of the line.
*/
func (self *_TraitTextView) GetLineAtY(y int) (target_iter C.GtkTextIter, line_top int) {
	var __cgo__line_top C.gint
	C.gtk_text_view_get_line_at_y(self.CPointer, &target_iter, C.gint(y), &__cgo__line_top)
	line_top = int(__cgo__line_top)
	return
}

/*
Gets the y coordinate of the top of the line containing @iter,
and the height of the line. The coordinate is a buffer coordinate;
convert to window coordinates with gtk_text_view_buffer_to_window_coords().
*/
func (self *_TraitTextView) GetLineYrange(iter *TextIter) (y int, height int) {
	var __cgo__y C.gint
	var __cgo__height C.gint
	C.gtk_text_view_get_line_yrange(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), &__cgo__y, &__cgo__height)
	y = int(__cgo__y)
	height = int(__cgo__height)
	return
}

/*
Returns whether the #GtkTextView is in overwrite mode or not.
*/
func (self *_TraitTextView) GetOverwrite() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_get_overwrite(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the default number of pixels to put above paragraphs.
*/
func (self *_TraitTextView) GetPixelsAboveLines() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_view_get_pixels_above_lines(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value set by gtk_text_view_set_pixels_below_lines().
*/
func (self *_TraitTextView) GetPixelsBelowLines() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_view_get_pixels_below_lines(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value set by gtk_text_view_set_pixels_inside_wrap().
*/
func (self *_TraitTextView) GetPixelsInsideWrap() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_view_get_pixels_inside_wrap(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the default right margin for text in @text_view. Tags
in the buffer may override the default.
*/
func (self *_TraitTextView) GetRightMargin() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_text_view_get_right_margin(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the default tabs for @text_view. Tags in the buffer may
override the defaults. The returned array will be %NULL if
“standard” (8-space) tabs are used. Free the return value
with pango_tab_array_free().
*/
func (self *_TraitTextView) GetTabs() (return__ *C.PangoTabArray) {
	return__ = C.gtk_text_view_get_tabs(self.CPointer)
	return
}

// gtk_text_view_get_vadjustment is not generated due to deprecation attr

/*
Fills @visible_rect with the currently-visible
region of the buffer, in buffer coordinates. Convert to window coordinates
with gtk_text_view_buffer_to_window_coords().
*/
func (self *_TraitTextView) GetVisibleRect() (visible_rect C.GdkRectangle) {
	C.gtk_text_view_get_visible_rect(self.CPointer, &visible_rect)
	return
}

/*
Retrieves the #GdkWindow corresponding to an area of the text view;
possible windows include the overall widget window, child windows
on the left, right, top, bottom, and the window that displays the
text buffer. Windows are %NULL and nonexistent if their width or
height is 0, and are nonexistent before the widget has been
realized.
*/
func (self *_TraitTextView) GetWindow(win C.GtkTextWindowType) (return__ *C.GdkWindow) {
	return__ = C.gtk_text_view_get_window(self.CPointer, win)
	return
}

/*
Usually used to find out which window an event corresponds to.
If you connect to an event signal on @text_view, this function
should be called on `event-&gt;window` to
see which window it was.
*/
func (self *_TraitTextView) GetWindowType(window *C.GdkWindow) (return__ C.GtkTextWindowType) {
	return__ = C.gtk_text_view_get_window_type(self.CPointer, window)
	return
}

/*
Gets the line wrapping for the view.
*/
func (self *_TraitTextView) GetWrapMode() (return__ C.GtkWrapMode) {
	return__ = C.gtk_text_view_get_wrap_mode(self.CPointer)
	return
}

/*
Allow the #GtkTextView input method to internally handle key press
and release events. If this function returns %TRUE, then no further
processing should be done for this key event. See
gtk_im_context_filter_keypress().

Note that you are expected to call this function from your handler
when overriding key event handling. This is needed in the case when
you need to insert your own key handling between the input method
and the default key event handling of the #GtkTextView.

|[<!-- language="C" -->
static gboolean
gtk_foo_bar_key_press_event (GtkWidget   *widget,
                             GdkEventKey *event)
{
  if ((key->keyval == GDK_KEY_Return || key->keyval == GDK_KEY_KP_Enter))
    {
      if (gtk_text_view_im_context_filter_keypress (GTK_TEXT_VIEW (view), event))
        return TRUE;
    }

    // Do some stuff

  return GTK_WIDGET_CLASS (gtk_foo_bar_parent_class)->key_press_event (widget, event);
}
]|
*/
func (self *_TraitTextView) ImContextFilterKeypress(event *C.GdkEventKey) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_im_context_filter_keypress(self.CPointer, event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Updates the position of a child, as for gtk_text_view_add_child_in_window().
*/
func (self *_TraitTextView) MoveChild(child *Widget, xpos int, ypos int) {
	C.gtk_text_view_move_child(self.CPointer, (*C.GtkWidget)(child.CPointer), C.gint(xpos), C.gint(ypos))
	return
}

/*
Moves a mark within the buffer so that it's
located within the currently-visible text area.
*/
func (self *_TraitTextView) MoveMarkOnscreen(mark *TextMark) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_move_mark_onscreen(self.CPointer, (*C.GtkTextMark)(mark.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Move the iterator a given number of characters visually, treating
it as the strong cursor position. If @count is positive, then the
new strong cursor position will be @count positions to the right of
the old cursor position. If @count is negative then the new strong
cursor position will be @count positions to the left of the old
cursor position.

In the presence of bi-directional text, the correspondence
between logical and visual order will depend on the direction
of the current run, and there may be jumps when the cursor
is moved off of the end of a run.
*/
func (self *_TraitTextView) MoveVisually(iter *TextIter, count int) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_move_visually(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), C.gint(count))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Moves the cursor to the currently visible region of the
buffer, it it isn’t there already.
*/
func (self *_TraitTextView) PlaceCursorOnscreen() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_place_cursor_onscreen(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Reset the input method context of the text view if needed.

This can be necessary in the case where modifying the buffer
would confuse on-going input method behavior.
*/
func (self *_TraitTextView) ResetImContext() {
	C.gtk_text_view_reset_im_context(self.CPointer)
	return
}

/*
Scrolls @text_view the minimum distance such that @mark is contained
within the visible area of the widget.
*/
func (self *_TraitTextView) ScrollMarkOnscreen(mark *TextMark) {
	C.gtk_text_view_scroll_mark_onscreen(self.CPointer, (*C.GtkTextMark)(mark.CPointer))
	return
}

/*
Scrolls @text_view so that @iter is on the screen in the position
indicated by @xalign and @yalign. An alignment of 0.0 indicates
left or top, 1.0 indicates right or bottom, 0.5 means center.
If @use_align is %FALSE, the text scrolls the minimal distance to
get the mark onscreen, possibly not scrolling at all. The effective
screen for purposes of this function is reduced by a margin of size
@within_margin.

Note that this function uses the currently-computed height of the
lines in the text buffer. Line heights are computed in an idle
handler; so this function may not have the desired effect if it’s
called before the height computations. To avoid oddness, consider
using gtk_text_view_scroll_to_mark() which saves a point to be
scrolled to after line validation.
*/
func (self *_TraitTextView) ScrollToIter(iter *TextIter, within_margin float64, use_align bool, xalign float64, yalign float64) (return__ bool) {
	__cgo__use_align := C.gboolean(0)
	if use_align {
		__cgo__use_align = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_scroll_to_iter(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)), C.gdouble(within_margin), __cgo__use_align, C.gdouble(xalign), C.gdouble(yalign))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Scrolls @text_view so that @mark is on the screen in the position
indicated by @xalign and @yalign. An alignment of 0.0 indicates
left or top, 1.0 indicates right or bottom, 0.5 means center.
If @use_align is %FALSE, the text scrolls the minimal distance to
get the mark onscreen, possibly not scrolling at all. The effective
screen for purposes of this function is reduced by a margin of size
@within_margin.
*/
func (self *_TraitTextView) ScrollToMark(mark *TextMark, within_margin float64, use_align bool, xalign float64, yalign float64) {
	__cgo__use_align := C.gboolean(0)
	if use_align {
		__cgo__use_align = C.gboolean(1)
	}
	C.gtk_text_view_scroll_to_mark(self.CPointer, (*C.GtkTextMark)(mark.CPointer), C.gdouble(within_margin), __cgo__use_align, C.gdouble(xalign), C.gdouble(yalign))
	return
}

/*
Sets the behavior of the text widget when the Tab key is pressed.
If @accepts_tab is %TRUE, a tab character is inserted. If @accepts_tab
is %FALSE the keyboard focus is moved to the next widget in the focus
chain.
*/
func (self *_TraitTextView) SetAcceptsTab(accepts_tab bool) {
	__cgo__accepts_tab := C.gboolean(0)
	if accepts_tab {
		__cgo__accepts_tab = C.gboolean(1)
	}
	C.gtk_text_view_set_accepts_tab(self.CPointer, __cgo__accepts_tab)
	return
}

/*
Sets the width of %GTK_TEXT_WINDOW_LEFT or %GTK_TEXT_WINDOW_RIGHT,
or the height of %GTK_TEXT_WINDOW_TOP or %GTK_TEXT_WINDOW_BOTTOM.
Automatically destroys the corresponding window if the size is set
to 0, and creates the window if the size is set to non-zero.  This
function can only be used for the “border windows,” it doesn’t work
with #GTK_TEXT_WINDOW_WIDGET, #GTK_TEXT_WINDOW_TEXT, or
#GTK_TEXT_WINDOW_PRIVATE.
*/
func (self *_TraitTextView) SetBorderWindowSize(type_ C.GtkTextWindowType, size int) {
	C.gtk_text_view_set_border_window_size(self.CPointer, type_, C.gint(size))
	return
}

/*
Sets @buffer as the buffer being displayed by @text_view. The previous
buffer displayed by the text view is unreferenced, and a reference is
added to @buffer. If you owned a reference to @buffer before passing it
to this function, you must remove that reference yourself; #GtkTextView
will not “adopt” it.
*/
func (self *_TraitTextView) SetBuffer(buffer *TextBuffer) {
	C.gtk_text_view_set_buffer(self.CPointer, (*C.GtkTextBuffer)(buffer.CPointer))
	return
}

/*
Toggles whether the insertion point is displayed. A buffer with no editable
text probably shouldn’t have a visible cursor, so you may want to turn
the cursor off.
*/
func (self *_TraitTextView) SetCursorVisible(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_text_view_set_cursor_visible(self.CPointer, __cgo__setting)
	return
}

/*
Sets the default editability of the #GtkTextView. You can override
this default setting with tags in the buffer, using the “editable”
attribute of tags.
*/
func (self *_TraitTextView) SetEditable(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_text_view_set_editable(self.CPointer, __cgo__setting)
	return
}

/*
Sets the default indentation for paragraphs in @text_view.
Tags in the buffer may override the default.
*/
func (self *_TraitTextView) SetIndent(indent int) {
	C.gtk_text_view_set_indent(self.CPointer, C.gint(indent))
	return
}

/*
Sets the #GtkTextView:input-hints property, which
allows input methods to fine-tune their behaviour.
*/
func (self *_TraitTextView) SetInputHints(hints C.GtkInputHints) {
	C.gtk_text_view_set_input_hints(self.CPointer, hints)
	return
}

/*
Sets the #GtkTextView:input-purpose property which
can be used by on-screen keyboards and other input
methods to adjust their behaviour.
*/
func (self *_TraitTextView) SetInputPurpose(purpose C.GtkInputPurpose) {
	C.gtk_text_view_set_input_purpose(self.CPointer, purpose)
	return
}

/*
Sets the default justification of text in @text_view.
Tags in the view’s buffer may override the default.
*/
func (self *_TraitTextView) SetJustification(justification C.GtkJustification) {
	C.gtk_text_view_set_justification(self.CPointer, justification)
	return
}

/*
Sets the default left margin for text in @text_view.
Tags in the buffer may override the default.
*/
func (self *_TraitTextView) SetLeftMargin(left_margin int) {
	C.gtk_text_view_set_left_margin(self.CPointer, C.gint(left_margin))
	return
}

/*
Changes the #GtkTextView overwrite mode.
*/
func (self *_TraitTextView) SetOverwrite(overwrite bool) {
	__cgo__overwrite := C.gboolean(0)
	if overwrite {
		__cgo__overwrite = C.gboolean(1)
	}
	C.gtk_text_view_set_overwrite(self.CPointer, __cgo__overwrite)
	return
}

/*
Sets the default number of blank pixels above paragraphs in @text_view.
Tags in the buffer for @text_view may override the defaults.
*/
func (self *_TraitTextView) SetPixelsAboveLines(pixels_above_lines int) {
	C.gtk_text_view_set_pixels_above_lines(self.CPointer, C.gint(pixels_above_lines))
	return
}

/*
Sets the default number of pixels of blank space
to put below paragraphs in @text_view. May be overridden
by tags applied to @text_view’s buffer.
*/
func (self *_TraitTextView) SetPixelsBelowLines(pixels_below_lines int) {
	C.gtk_text_view_set_pixels_below_lines(self.CPointer, C.gint(pixels_below_lines))
	return
}

/*
Sets the default number of pixels of blank space to leave between
display/wrapped lines within a paragraph. May be overridden by
tags in @text_view’s buffer.
*/
func (self *_TraitTextView) SetPixelsInsideWrap(pixels_inside_wrap int) {
	C.gtk_text_view_set_pixels_inside_wrap(self.CPointer, C.gint(pixels_inside_wrap))
	return
}

/*
Sets the default right margin for text in the text view.
Tags in the buffer may override the default.
*/
func (self *_TraitTextView) SetRightMargin(right_margin int) {
	C.gtk_text_view_set_right_margin(self.CPointer, C.gint(right_margin))
	return
}

/*
Sets the default tab stops for paragraphs in @text_view.
Tags in the buffer may override the default.
*/
func (self *_TraitTextView) SetTabs(tabs *C.PangoTabArray) {
	C.gtk_text_view_set_tabs(self.CPointer, tabs)
	return
}

/*
Sets the line wrapping for the view.
*/
func (self *_TraitTextView) SetWrapMode(wrap_mode C.GtkWrapMode) {
	C.gtk_text_view_set_wrap_mode(self.CPointer, wrap_mode)
	return
}

/*
Determines whether @iter is at the start of a display line.
See gtk_text_view_forward_display_line() for an explanation of
display lines vs. paragraphs.
*/
func (self *_TraitTextView) StartsDisplayLine(iter *TextIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_text_view_starts_display_line(self.CPointer, (*C.GtkTextIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts coordinates on the window identified by @win to buffer
coordinates, storing the result in (@buffer_x,@buffer_y).

Note that you can’t convert coordinates for a nonexisting window (see
gtk_text_view_set_border_window_size()).
*/
func (self *_TraitTextView) WindowToBufferCoords(win C.GtkTextWindowType, window_x int, window_y int) (buffer_x int, buffer_y int) {
	var __cgo__buffer_x C.gint
	var __cgo__buffer_y C.gint
	C.gtk_text_view_window_to_buffer_coords(self.CPointer, win, C.gint(window_x), C.gint(window_y), &__cgo__buffer_x, &__cgo__buffer_y)
	buffer_x = int(__cgo__buffer_x)
	buffer_y = int(__cgo__buffer_y)
	return
}

type _TraitThemingEngine struct{ CPointer *C.GtkThemingEngine }

// gtk_theming_engine_get is not generated due to varargs

/*
Gets the background color for a given state.
*/
func (self *_TraitThemingEngine) GetBackgroundColor(state C.GtkStateFlags) (color C.GdkRGBA) {
	C.gtk_theming_engine_get_background_color(self.CPointer, state, &color)
	return
}

/*
Gets the border for a given state as a #GtkBorder.
*/
func (self *_TraitThemingEngine) GetBorder(state C.GtkStateFlags) (border C.GtkBorder) {
	C.gtk_theming_engine_get_border(self.CPointer, state, &border)
	return
}

/*
Gets the border color for a given state.
*/
func (self *_TraitThemingEngine) GetBorderColor(state C.GtkStateFlags) (color C.GdkRGBA) {
	C.gtk_theming_engine_get_border_color(self.CPointer, state, &color)
	return
}

/*
Gets the foreground color for a given state.
*/
func (self *_TraitThemingEngine) GetColor(state C.GtkStateFlags) (color C.GdkRGBA) {
	C.gtk_theming_engine_get_color(self.CPointer, state, &color)
	return
}

// gtk_theming_engine_get_direction is not generated due to deprecation attr

// gtk_theming_engine_get_font is not generated due to deprecation attr

/*
Returns the widget direction used for rendering.
*/
func (self *_TraitThemingEngine) GetJunctionSides() (return__ C.GtkJunctionSides) {
	return__ = C.gtk_theming_engine_get_junction_sides(self.CPointer)
	return
}

/*
Gets the margin for a given state as a #GtkBorder.
*/
func (self *_TraitThemingEngine) GetMargin(state C.GtkStateFlags) (margin C.GtkBorder) {
	C.gtk_theming_engine_get_margin(self.CPointer, state, &margin)
	return
}

/*
Gets the padding for a given state as a #GtkBorder.
*/
func (self *_TraitThemingEngine) GetPadding(state C.GtkStateFlags) (padding C.GtkBorder) {
	C.gtk_theming_engine_get_padding(self.CPointer, state, &padding)
	return
}

/*
Returns the widget path used for style matching.
*/
func (self *_TraitThemingEngine) GetPath() (return__ *WidgetPath) {
	var __cgo__return__ *C.GtkWidgetPath
	__cgo__return__ = C.gtk_theming_engine_get_path(self.CPointer)
	return__ = (*WidgetPath)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Gets a property value as retrieved from the style settings that apply
to the currently rendered element.
*/
func (self *_TraitThemingEngine) GetProperty(property string, state C.GtkStateFlags) (value C.GValue) {
	__cgo__property := (*C.gchar)(unsafe.Pointer(C.CString(property)))
	C.gtk_theming_engine_get_property(self.CPointer, __cgo__property, state, &value)
	C.free(unsafe.Pointer(__cgo__property))
	return
}

/*
Returns the #GdkScreen to which @engine currently rendering to.
*/
func (self *_TraitThemingEngine) GetScreen() (return__ *C.GdkScreen) {
	return__ = C.gtk_theming_engine_get_screen(self.CPointer)
	return
}

/*
returns the state used when rendering.
*/
func (self *_TraitThemingEngine) GetState() (return__ C.GtkStateFlags) {
	return__ = C.gtk_theming_engine_get_state(self.CPointer)
	return
}

// gtk_theming_engine_get_style is not generated due to varargs

/*
Gets the value for a widget style property.
*/
func (self *_TraitThemingEngine) GetStyleProperty(property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.gtk_theming_engine_get_style_property(self.CPointer, __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

// gtk_theming_engine_get_style_valist is not generated due to varargs

// gtk_theming_engine_get_valist is not generated due to varargs

/*
Returns %TRUE if the currently rendered contents have
defined the given class name.
*/
func (self *_TraitThemingEngine) HasClass(style_class string) (return__ bool) {
	__cgo__style_class := (*C.gchar)(unsafe.Pointer(C.CString(style_class)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_theming_engine_has_class(self.CPointer, __cgo__style_class)
	C.free(unsafe.Pointer(__cgo__style_class))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the currently rendered contents have the
region defined. If @flags_return is not %NULL, it is set
to the flags affecting the region.
*/
func (self *_TraitThemingEngine) HasRegion(style_region string) (flags C.GtkRegionFlags, return__ bool) {
	__cgo__style_region := (*C.gchar)(unsafe.Pointer(C.CString(style_region)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_theming_engine_has_region(self.CPointer, __cgo__style_region, &flags)
	C.free(unsafe.Pointer(__cgo__style_region))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Looks up and resolves a color name in the current style’s color map.
*/
func (self *_TraitThemingEngine) LookupColor(color_name string) (color C.GdkRGBA, return__ bool) {
	__cgo__color_name := (*C.gchar)(unsafe.Pointer(C.CString(color_name)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_theming_engine_lookup_color(self.CPointer, __cgo__color_name, &color)
	C.free(unsafe.Pointer(__cgo__color_name))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_theming_engine_state_is_running is not generated due to deprecation attr

type _TraitToggleAction struct{ CPointer *C.GtkToggleAction }

// gtk_toggle_action_get_active is not generated due to deprecation attr

// gtk_toggle_action_get_draw_as_radio is not generated due to deprecation attr

// gtk_toggle_action_set_active is not generated due to deprecation attr

// gtk_toggle_action_set_draw_as_radio is not generated due to deprecation attr

// gtk_toggle_action_toggled is not generated due to deprecation attr

type _TraitToggleButton struct{ CPointer *C.GtkToggleButton }

/*
Queries a #GtkToggleButton and returns its current state. Returns %TRUE if
the toggle button is pressed in and %FALSE if it is raised.
*/
func (self *_TraitToggleButton) GetActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_toggle_button_get_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_toggle_button_set_inconsistent().
*/
func (self *_TraitToggleButton) GetInconsistent() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_toggle_button_get_inconsistent(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves whether the button is displayed as a separate indicator
and label. See gtk_toggle_button_set_mode().
*/
func (self *_TraitToggleButton) GetMode() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_toggle_button_get_mode(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the status of the toggle button. Set to %TRUE if you want the
GtkToggleButton to be “pressed in”, and %FALSE to raise it.
This action causes the #GtkToggleButton::toggled signal and the
#GtkButton::clicked signal to be emitted.
*/
func (self *_TraitToggleButton) SetActive(is_active bool) {
	__cgo__is_active := C.gboolean(0)
	if is_active {
		__cgo__is_active = C.gboolean(1)
	}
	C.gtk_toggle_button_set_active(self.CPointer, __cgo__is_active)
	return
}

/*
If the user has selected a range of elements (such as some text or
spreadsheet cells) that are affected by a toggle button, and the
current values in that range are inconsistent, you may want to
display the toggle in an “in between” state. This function turns on
“in between” display.  Normally you would turn off the inconsistent
state again if the user toggles the toggle button. This has to be
done manually, gtk_toggle_button_set_inconsistent() only affects
visual appearance, it doesn’t affect the semantics of the button.
*/
func (self *_TraitToggleButton) SetInconsistent(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_toggle_button_set_inconsistent(self.CPointer, __cgo__setting)
	return
}

/*
Sets whether the button is displayed as a separate indicator and label.
You can call this function on a checkbutton or a radiobutton with
@draw_indicator = %FALSE to make the button look like a normal button

This function only affects instances of classes like #GtkCheckButton
and #GtkRadioButton that derive from #GtkToggleButton,
not instances of #GtkToggleButton itself.
*/
func (self *_TraitToggleButton) SetMode(draw_indicator bool) {
	__cgo__draw_indicator := C.gboolean(0)
	if draw_indicator {
		__cgo__draw_indicator = C.gboolean(1)
	}
	C.gtk_toggle_button_set_mode(self.CPointer, __cgo__draw_indicator)
	return
}

/*
Emits the #GtkToggleButton::toggled signal on the
#GtkToggleButton. There is no good reason for an
application ever to call this function.
*/
func (self *_TraitToggleButton) Toggled() {
	C.gtk_toggle_button_toggled(self.CPointer)
	return
}

type _TraitToggleToolButton struct{ CPointer *C.GtkToggleToolButton }

/*
Queries a #GtkToggleToolButton and returns its current state.
Returns %TRUE if the toggle button is pressed in and %FALSE if it is raised.
*/
func (self *_TraitToggleToolButton) GetActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_toggle_tool_button_get_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the status of the toggle tool button. Set to %TRUE if you
want the GtkToggleButton to be “pressed in”, and %FALSE to raise it.
This action causes the toggled signal to be emitted.
*/
func (self *_TraitToggleToolButton) SetActive(is_active bool) {
	__cgo__is_active := C.gboolean(0)
	if is_active {
		__cgo__is_active = C.gboolean(1)
	}
	C.gtk_toggle_tool_button_set_active(self.CPointer, __cgo__is_active)
	return
}

type _TraitToolButton struct{ CPointer *C.GtkToolButton }

/*
Returns the name of the themed icon for the tool button,
see gtk_tool_button_set_icon_name().
*/
func (self *_TraitToolButton) GetIconName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_tool_button_get_icon_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Return the widget used as icon widget on @button.
See gtk_tool_button_set_icon_widget().
*/
func (self *_TraitToolButton) GetIconWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tool_button_get_icon_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the label used by the tool button, or %NULL if the tool button
doesn’t have a label. or uses a the label from a stock item. The returned
string is owned by GTK+, and must not be modified or freed.
*/
func (self *_TraitToolButton) GetLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_tool_button_get_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the widget used as label on @button.
See gtk_tool_button_set_label_widget().
*/
func (self *_TraitToolButton) GetLabelWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tool_button_get_label_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

// gtk_tool_button_get_stock_id is not generated due to deprecation attr

/*
Returns whether underscores in the label property are used as mnemonics
on menu items on the overflow menu. See gtk_tool_button_set_use_underline().
*/
func (self *_TraitToolButton) GetUseUnderline() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_button_get_use_underline(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the icon for the tool button from a named themed icon.
See the docs for #GtkIconTheme for more details.
The “icon_name” property only has an effect if not
overridden by non-%NULL “label”, “icon_widget” and “stock_id”
properties.
*/
func (self *_TraitToolButton) SetIconName(icon_name string) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_tool_button_set_icon_name(self.CPointer, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

/*
Sets @icon as the widget used as icon on @button. If @icon_widget is
%NULL the icon is determined by the “stock_id” property. If the
“stock_id” property is also %NULL, @button will not have an icon.
*/
func (self *_TraitToolButton) SetIconWidget(icon_widget *Widget) {
	C.gtk_tool_button_set_icon_widget(self.CPointer, (*C.GtkWidget)(icon_widget.CPointer))
	return
}

/*
Sets @label as the label used for the tool button. The “label” property
only has an effect if not overridden by a non-%NULL “label_widget” property.
If both the “label_widget” and “label” properties are %NULL, the label
is determined by the “stock_id” property. If the “stock_id” property is also
%NULL, @button will not have a label.
*/
func (self *_TraitToolButton) SetLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_tool_button_set_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Sets @label_widget as the widget that will be used as the label
for @button. If @label_widget is %NULL the “label” property is used
as label. If “label” is also %NULL, the label in the stock item
determined by the “stock_id” property is used as label. If
“stock_id” is also %NULL, @button does not have a label.
*/
func (self *_TraitToolButton) SetLabelWidget(label_widget *Widget) {
	C.gtk_tool_button_set_label_widget(self.CPointer, (*C.GtkWidget)(label_widget.CPointer))
	return
}

// gtk_tool_button_set_stock_id is not generated due to deprecation attr

/*
If set, an underline in the label property indicates that the next character
should be used for the mnemonic accelerator key in the overflow menu. For
example, if the label property is “_Open” and @use_underline is %TRUE,
the label on the tool button will be “Open” and the item on the overflow
menu will have an underlined “O”.

Labels shown on tool buttons never have mnemonics on them; this property
only affects the menu item on the overflow menu.
*/
func (self *_TraitToolButton) SetUseUnderline(use_underline bool) {
	__cgo__use_underline := C.gboolean(0)
	if use_underline {
		__cgo__use_underline = C.gboolean(1)
	}
	C.gtk_tool_button_set_use_underline(self.CPointer, __cgo__use_underline)
	return
}

type _TraitToolItem struct{ CPointer *C.GtkToolItem }

/*
Returns the ellipsize mode used for @tool_item. Custom subclasses of
#GtkToolItem should call this function to find out how text should
be ellipsized.
*/
func (self *_TraitToolItem) GetEllipsizeMode() (return__ C.PangoEllipsizeMode) {
	return__ = C.gtk_tool_item_get_ellipsize_mode(self.CPointer)
	return
}

/*
Returns whether @tool_item is allocated extra space.
See gtk_tool_item_set_expand().
*/
func (self *_TraitToolItem) GetExpand() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_item_get_expand(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether @tool_item is the same size as other homogeneous
items. See gtk_tool_item_set_homogeneous().
*/
func (self *_TraitToolItem) GetHomogeneous() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_item_get_homogeneous(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the icon size used for @tool_item. Custom subclasses of
#GtkToolItem should call this function to find out what size icons
they should use.
*/
func (self *_TraitToolItem) GetIconSize() (return__ C.GtkIconSize) {
	return__ = C.gtk_tool_item_get_icon_size(self.CPointer)
	return
}

/*
Returns whether @tool_item is considered important. See
gtk_tool_item_set_is_important()
*/
func (self *_TraitToolItem) GetIsImportant() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_item_get_is_important(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the orientation used for @tool_item. Custom subclasses of
#GtkToolItem should call this function to find out what size icons
they should use.
*/
func (self *_TraitToolItem) GetOrientation() (return__ C.GtkOrientation) {
	return__ = C.gtk_tool_item_get_orientation(self.CPointer)
	return
}

/*
If @menu_item_id matches the string passed to
gtk_tool_item_set_proxy_menu_item() return the corresponding #GtkMenuItem.

Custom subclasses of #GtkToolItem should use this function to
update their menu item when the #GtkToolItem changes. That the
@menu_item_ids must match ensures that a #GtkToolItem
will not inadvertently change a menu item that they did not create.
*/
func (self *_TraitToolItem) GetProxyMenuItem(menu_item_id string) (return__ *Widget) {
	__cgo__menu_item_id := (*C.gchar)(unsafe.Pointer(C.CString(menu_item_id)))
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tool_item_get_proxy_menu_item(self.CPointer, __cgo__menu_item_id)
	C.free(unsafe.Pointer(__cgo__menu_item_id))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the relief style of @tool_item. See gtk_button_set_relief().
Custom subclasses of #GtkToolItem should call this function in the handler
of the #GtkToolItem::toolbar_reconfigured signal to find out the
relief style of buttons.
*/
func (self *_TraitToolItem) GetReliefStyle() (return__ C.GtkReliefStyle) {
	return__ = C.gtk_tool_item_get_relief_style(self.CPointer)
	return
}

/*
Returns the text alignment used for @tool_item. Custom subclasses of
#GtkToolItem should call this function to find out how text should
be aligned.
*/
func (self *_TraitToolItem) GetTextAlignment() (return__ float32) {
	var __cgo__return__ C.gfloat
	__cgo__return__ = C.gtk_tool_item_get_text_alignment(self.CPointer)
	return__ = float32(__cgo__return__)
	return
}

/*
Returns the text orientation used for @tool_item. Custom subclasses of
#GtkToolItem should call this function to find out how text should
be orientated.
*/
func (self *_TraitToolItem) GetTextOrientation() (return__ C.GtkOrientation) {
	return__ = C.gtk_tool_item_get_text_orientation(self.CPointer)
	return
}

/*
Returns the size group used for labels in @tool_item.
Custom subclasses of #GtkToolItem should call this function
and use the size group for labels.
*/
func (self *_TraitToolItem) GetTextSizeGroup() (return__ *SizeGroup) {
	var __cgo__return__ *C.GtkSizeGroup
	__cgo__return__ = C.gtk_tool_item_get_text_size_group(self.CPointer)
	return__ = NewSizeGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the toolbar style used for @tool_item. Custom subclasses of
#GtkToolItem should call this function in the handler of the
GtkToolItem::toolbar_reconfigured signal to find out in what style
the toolbar is displayed and change themselves accordingly

Possibilities are:
- %GTK_TOOLBAR_BOTH, meaning the tool item should show
  both an icon and a label, stacked vertically
- %GTK_TOOLBAR_ICONS, meaning the toolbar shows only icons
- %GTK_TOOLBAR_TEXT, meaning the tool item should only show text
- %GTK_TOOLBAR_BOTH_HORIZ, meaning the tool item should show
  both an icon and a label, arranged horizontally
*/
func (self *_TraitToolItem) GetToolbarStyle() (return__ C.GtkToolbarStyle) {
	return__ = C.gtk_tool_item_get_toolbar_style(self.CPointer)
	return
}

/*
Returns whether @tool_item has a drag window. See
gtk_tool_item_set_use_drag_window().
*/
func (self *_TraitToolItem) GetUseDragWindow() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_item_get_use_drag_window(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the @tool_item is visible on toolbars that are
docked horizontally.
*/
func (self *_TraitToolItem) GetVisibleHorizontal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_item_get_visible_horizontal(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether @tool_item is visible when the toolbar is docked vertically.
See gtk_tool_item_set_visible_vertical().
*/
func (self *_TraitToolItem) GetVisibleVertical() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_item_get_visible_vertical(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Calling this function signals to the toolbar that the
overflow menu item for @tool_item has changed. If the
overflow menu is visible when this function it called,
the menu will be rebuilt.

The function must be called when the tool item changes what it
will do in response to the #GtkToolItem::create-menu-proxy signal.
*/
func (self *_TraitToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu(self.CPointer)
	return
}

/*
Returns the #GtkMenuItem that was last set by
gtk_tool_item_set_proxy_menu_item(), ie. the #GtkMenuItem
that is going to appear in the overflow menu.
*/
func (self *_TraitToolItem) RetrieveProxyMenuItem() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tool_item_retrieve_proxy_menu_item(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Sets whether @tool_item is allocated extra space when there
is more room on the toolbar then needed for the items. The
effect is that the item gets bigger when the toolbar gets bigger
and smaller when the toolbar gets smaller.
*/
func (self *_TraitToolItem) SetExpand(expand bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	C.gtk_tool_item_set_expand(self.CPointer, __cgo__expand)
	return
}

/*
Sets whether @tool_item is to be allocated the same size as other
homogeneous items. The effect is that all homogeneous items will have
the same width as the widest of the items.
*/
func (self *_TraitToolItem) SetHomogeneous(homogeneous bool) {
	__cgo__homogeneous := C.gboolean(0)
	if homogeneous {
		__cgo__homogeneous = C.gboolean(1)
	}
	C.gtk_tool_item_set_homogeneous(self.CPointer, __cgo__homogeneous)
	return
}

/*
Sets whether @tool_item should be considered important. The #GtkToolButton
class uses this property to determine whether to show or hide its label
when the toolbar style is %GTK_TOOLBAR_BOTH_HORIZ. The result is that
only tool buttons with the “is_important” property set have labels, an
effect known as “priority text”
*/
func (self *_TraitToolItem) SetIsImportant(is_important bool) {
	__cgo__is_important := C.gboolean(0)
	if is_important {
		__cgo__is_important = C.gboolean(1)
	}
	C.gtk_tool_item_set_is_important(self.CPointer, __cgo__is_important)
	return
}

/*
Sets the #GtkMenuItem used in the toolbar overflow menu. The
@menu_item_id is used to identify the caller of this function and
should also be used with gtk_tool_item_get_proxy_menu_item().
*/
func (self *_TraitToolItem) SetProxyMenuItem(menu_item_id string, menu_item *Widget) {
	__cgo__menu_item_id := (*C.gchar)(unsafe.Pointer(C.CString(menu_item_id)))
	C.gtk_tool_item_set_proxy_menu_item(self.CPointer, __cgo__menu_item_id, (*C.GtkWidget)(menu_item.CPointer))
	C.free(unsafe.Pointer(__cgo__menu_item_id))
	return
}

/*
Sets the markup text to be displayed as tooltip on the item.
See gtk_widget_set_tooltip_markup().
*/
func (self *_TraitToolItem) SetTooltipMarkup(markup string) {
	__cgo__markup := (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	C.gtk_tool_item_set_tooltip_markup(self.CPointer, __cgo__markup)
	C.free(unsafe.Pointer(__cgo__markup))
	return
}

/*
Sets the text to be displayed as tooltip on the item.
See gtk_widget_set_tooltip_text().
*/
func (self *_TraitToolItem) SetTooltipText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_tool_item_set_tooltip_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Sets whether @tool_item has a drag window. When %TRUE the
toolitem can be used as a drag source through gtk_drag_source_set().
When @tool_item has a drag window it will intercept all events,
even those that would otherwise be sent to a child of @tool_item.
*/
func (self *_TraitToolItem) SetUseDragWindow(use_drag_window bool) {
	__cgo__use_drag_window := C.gboolean(0)
	if use_drag_window {
		__cgo__use_drag_window = C.gboolean(1)
	}
	C.gtk_tool_item_set_use_drag_window(self.CPointer, __cgo__use_drag_window)
	return
}

/*
Sets whether @tool_item is visible when the toolbar is docked horizontally.
*/
func (self *_TraitToolItem) SetVisibleHorizontal(visible_horizontal bool) {
	__cgo__visible_horizontal := C.gboolean(0)
	if visible_horizontal {
		__cgo__visible_horizontal = C.gboolean(1)
	}
	C.gtk_tool_item_set_visible_horizontal(self.CPointer, __cgo__visible_horizontal)
	return
}

/*
Sets whether @tool_item is visible when the toolbar is docked
vertically. Some tool items, such as text entries, are too wide to be
useful on a vertically docked toolbar. If @visible_vertical is %FALSE
@tool_item will not appear on toolbars that are docked vertically.
*/
func (self *_TraitToolItem) SetVisibleVertical(visible_vertical bool) {
	__cgo__visible_vertical := C.gboolean(0)
	if visible_vertical {
		__cgo__visible_vertical = C.gboolean(1)
	}
	C.gtk_tool_item_set_visible_vertical(self.CPointer, __cgo__visible_vertical)
	return
}

/*
Emits the signal #GtkToolItem::toolbar_reconfigured on @tool_item.
#GtkToolbar and other #GtkToolShell implementations use this function
to notify children, when some aspect of their configuration changes.
*/
func (self *_TraitToolItem) ToolbarReconfigured() {
	C.gtk_tool_item_toolbar_reconfigured(self.CPointer)
	return
}

type _TraitToolItemGroup struct{ CPointer *C.GtkToolItemGroup }

/*
Gets whether @group is collapsed or expanded.
*/
func (self *_TraitToolItemGroup) GetCollapsed() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_item_group_get_collapsed(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the tool item at position (x, y).
*/
func (self *_TraitToolItemGroup) GetDropItem(x int, y int) (return__ *ToolItem) {
	var __cgo__return__ *C.GtkToolItem
	__cgo__return__ = C.gtk_tool_item_group_get_drop_item(self.CPointer, C.gint(x), C.gint(y))
	return__ = NewToolItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the ellipsization mode of @group.
*/
func (self *_TraitToolItemGroup) GetEllipsize() (return__ C.PangoEllipsizeMode) {
	return__ = C.gtk_tool_item_group_get_ellipsize(self.CPointer)
	return
}

/*
Gets the relief mode of the header button of @group.
*/
func (self *_TraitToolItemGroup) GetHeaderRelief() (return__ C.GtkReliefStyle) {
	return__ = C.gtk_tool_item_group_get_header_relief(self.CPointer)
	return
}

/*
Gets the position of @item in @group as index.
*/
func (self *_TraitToolItemGroup) GetItemPosition(item *ToolItem) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tool_item_group_get_item_position(self.CPointer, (*C.GtkToolItem)(item.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Gets the label of @group.
*/
func (self *_TraitToolItemGroup) GetLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_tool_item_group_get_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the label widget of @group.
See gtk_tool_item_group_set_label_widget().
*/
func (self *_TraitToolItemGroup) GetLabelWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tool_item_group_get_label_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the number of tool items in @group.
*/
func (self *_TraitToolItemGroup) GetNItems() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_tool_item_group_get_n_items(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the tool item at @index in group.
*/
func (self *_TraitToolItemGroup) GetNthItem(index uint) (return__ *ToolItem) {
	var __cgo__return__ *C.GtkToolItem
	__cgo__return__ = C.gtk_tool_item_group_get_nth_item(self.CPointer, C.guint(index))
	return__ = NewToolItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Inserts @item at @position in the list of children of @group.
*/
func (self *_TraitToolItemGroup) Insert(item *ToolItem, position int) {
	C.gtk_tool_item_group_insert(self.CPointer, (*C.GtkToolItem)(item.CPointer), C.gint(position))
	return
}

/*
Sets whether the @group should be collapsed or expanded.
*/
func (self *_TraitToolItemGroup) SetCollapsed(collapsed bool) {
	__cgo__collapsed := C.gboolean(0)
	if collapsed {
		__cgo__collapsed = C.gboolean(1)
	}
	C.gtk_tool_item_group_set_collapsed(self.CPointer, __cgo__collapsed)
	return
}

/*
Sets the ellipsization mode which should be used by labels in @group.
*/
func (self *_TraitToolItemGroup) SetEllipsize(ellipsize C.PangoEllipsizeMode) {
	C.gtk_tool_item_group_set_ellipsize(self.CPointer, ellipsize)
	return
}

/*
Set the button relief of the group header.
See gtk_button_set_relief() for details.
*/
func (self *_TraitToolItemGroup) SetHeaderRelief(style C.GtkReliefStyle) {
	C.gtk_tool_item_group_set_header_relief(self.CPointer, style)
	return
}

/*
Sets the position of @item in the list of children of @group.
*/
func (self *_TraitToolItemGroup) SetItemPosition(item *ToolItem, position int) {
	C.gtk_tool_item_group_set_item_position(self.CPointer, (*C.GtkToolItem)(item.CPointer), C.gint(position))
	return
}

/*
Sets the label of the tool item group. The label is displayed in the header
of the group.
*/
func (self *_TraitToolItemGroup) SetLabel(label string) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	C.gtk_tool_item_group_set_label(self.CPointer, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	return
}

/*
Sets the label of the tool item group.
The label widget is displayed in the header of the group, in place
of the usual label.
*/
func (self *_TraitToolItemGroup) SetLabelWidget(label_widget *Widget) {
	C.gtk_tool_item_group_set_label_widget(self.CPointer, (*C.GtkWidget)(label_widget.CPointer))
	return
}

type _TraitToolPalette struct{ CPointer *C.GtkToolPalette }

/*
Sets @palette as drag source (see gtk_tool_palette_set_drag_source())
and sets @widget as a drag destination for drags from @palette.
See gtk_drag_dest_set().
*/
func (self *_TraitToolPalette) AddDragDest(widget *Widget, flags C.GtkDestDefaults, targets C.GtkToolPaletteDragTargets, actions C.GdkDragAction) {
	C.gtk_tool_palette_add_drag_dest(self.CPointer, (*C.GtkWidget)(widget.CPointer), flags, targets, actions)
	return
}

/*
Get the dragged item from the selection.
This could be a #GtkToolItem or a #GtkToolItemGroup.
*/
func (self *_TraitToolPalette) GetDragItem(selection *SelectionData) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tool_palette_get_drag_item(self.CPointer, (*C.GtkSelectionData)(unsafe.Pointer(selection)))
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the group at position (x, y).
*/
func (self *_TraitToolPalette) GetDropGroup(x int, y int) (return__ *ToolItemGroup) {
	var __cgo__return__ *C.GtkToolItemGroup
	__cgo__return__ = C.gtk_tool_palette_get_drop_group(self.CPointer, C.gint(x), C.gint(y))
	return__ = NewToolItemGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the item at position (x, y).
See gtk_tool_palette_get_drop_group().
*/
func (self *_TraitToolPalette) GetDropItem(x int, y int) (return__ *ToolItem) {
	var __cgo__return__ *C.GtkToolItem
	__cgo__return__ = C.gtk_tool_palette_get_drop_item(self.CPointer, C.gint(x), C.gint(y))
	return__ = NewToolItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets whether @group is exclusive or not.
See gtk_tool_palette_set_exclusive().
*/
func (self *_TraitToolPalette) GetExclusive(group *ToolItemGroup) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_palette_get_exclusive(self.CPointer, (*C.GtkToolItemGroup)(group.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether group should be given extra space.
See gtk_tool_palette_set_expand().
*/
func (self *_TraitToolPalette) GetExpand(group *ToolItemGroup) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tool_palette_get_expand(self.CPointer, (*C.GtkToolItemGroup)(group.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the position of @group in @palette as index.
See gtk_tool_palette_set_group_position().
*/
func (self *_TraitToolPalette) GetGroupPosition(group *ToolItemGroup) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tool_palette_get_group_position(self.CPointer, (*C.GtkToolItemGroup)(group.CPointer))
	return__ = int(__cgo__return__)
	return
}

// gtk_tool_palette_get_hadjustment is not generated due to deprecation attr

/*
Gets the size of icons in the tool palette.
See gtk_tool_palette_set_icon_size().
*/
func (self *_TraitToolPalette) GetIconSize() (return__ C.GtkIconSize) {
	return__ = C.gtk_tool_palette_get_icon_size(self.CPointer)
	return
}

/*
Gets the style (icons, text or both) of items in the tool palette.
*/
func (self *_TraitToolPalette) GetStyle() (return__ C.GtkToolbarStyle) {
	return__ = C.gtk_tool_palette_get_style(self.CPointer)
	return
}

// gtk_tool_palette_get_vadjustment is not generated due to deprecation attr

/*
Sets the tool palette as a drag source.
Enables all groups and items in the tool palette as drag sources
on button 1 and button 3 press with copy and move actions.
See gtk_drag_source_set().
*/
func (self *_TraitToolPalette) SetDragSource(targets C.GtkToolPaletteDragTargets) {
	C.gtk_tool_palette_set_drag_source(self.CPointer, targets)
	return
}

/*
Sets whether the group should be exclusive or not.
If an exclusive group is expanded all other groups are collapsed.
*/
func (self *_TraitToolPalette) SetExclusive(group *ToolItemGroup, exclusive bool) {
	__cgo__exclusive := C.gboolean(0)
	if exclusive {
		__cgo__exclusive = C.gboolean(1)
	}
	C.gtk_tool_palette_set_exclusive(self.CPointer, (*C.GtkToolItemGroup)(group.CPointer), __cgo__exclusive)
	return
}

/*
Sets whether the group should be given extra space.
*/
func (self *_TraitToolPalette) SetExpand(group *ToolItemGroup, expand bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	C.gtk_tool_palette_set_expand(self.CPointer, (*C.GtkToolItemGroup)(group.CPointer), __cgo__expand)
	return
}

/*
Sets the position of the group as an index of the tool palette.
If position is 0 the group will become the first child, if position is
-1 it will become the last child.
*/
func (self *_TraitToolPalette) SetGroupPosition(group *ToolItemGroup, position int) {
	C.gtk_tool_palette_set_group_position(self.CPointer, (*C.GtkToolItemGroup)(group.CPointer), C.gint(position))
	return
}

/*
Sets the size of icons in the tool palette.
*/
func (self *_TraitToolPalette) SetIconSize(icon_size C.GtkIconSize) {
	C.gtk_tool_palette_set_icon_size(self.CPointer, icon_size)
	return
}

/*
Sets the style (text, icons or both) of items in the tool palette.
*/
func (self *_TraitToolPalette) SetStyle(style C.GtkToolbarStyle) {
	C.gtk_tool_palette_set_style(self.CPointer, style)
	return
}

/*
Unsets the tool palette icon size set with gtk_tool_palette_set_icon_size(),
so that user preferences will be used to determine the icon size.
*/
func (self *_TraitToolPalette) UnsetIconSize() {
	C.gtk_tool_palette_unset_icon_size(self.CPointer)
	return
}

/*
Unsets a toolbar style set with gtk_tool_palette_set_style(),
so that user preferences will be used to determine the toolbar style.
*/
func (self *_TraitToolPalette) UnsetStyle() {
	C.gtk_tool_palette_unset_style(self.CPointer)
	return
}

type _TraitToolbar struct{ CPointer *C.GtkToolbar }

/*
Returns the position corresponding to the indicated point on
@toolbar. This is useful when dragging items to the toolbar:
this function returns the position a new item should be
inserted.

@x and @y are in @toolbar coordinates.
*/
func (self *_TraitToolbar) GetDropIndex(x int, y int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_toolbar_get_drop_index(self.CPointer, C.gint(x), C.gint(y))
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the icon size for the toolbar. See gtk_toolbar_set_icon_size().
*/
func (self *_TraitToolbar) GetIconSize() (return__ C.GtkIconSize) {
	return__ = C.gtk_toolbar_get_icon_size(self.CPointer)
	return
}

/*
Returns the position of @item on the toolbar, starting from 0.
It is an error if @item is not a child of the toolbar.
*/
func (self *_TraitToolbar) GetItemIndex(item *ToolItem) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_toolbar_get_item_index(self.CPointer, (*C.GtkToolItem)(item.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Returns the number of items on the toolbar.
*/
func (self *_TraitToolbar) GetNItems() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_toolbar_get_n_items(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the @n'th item on @toolbar, or %NULL if the
toolbar does not contain an @n'th item.
*/
func (self *_TraitToolbar) GetNthItem(n int) (return__ *ToolItem) {
	var __cgo__return__ *C.GtkToolItem
	__cgo__return__ = C.gtk_toolbar_get_nth_item(self.CPointer, C.gint(n))
	return__ = NewToolItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the relief style of buttons on @toolbar. See
gtk_button_set_relief().
*/
func (self *_TraitToolbar) GetReliefStyle() (return__ C.GtkReliefStyle) {
	return__ = C.gtk_toolbar_get_relief_style(self.CPointer)
	return
}

/*
Returns whether the toolbar has an overflow menu.
See gtk_toolbar_set_show_arrow().
*/
func (self *_TraitToolbar) GetShowArrow() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_toolbar_get_show_arrow(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves whether the toolbar has text, icons, or both . See
gtk_toolbar_set_style().
*/
func (self *_TraitToolbar) GetStyle() (return__ C.GtkToolbarStyle) {
	return__ = C.gtk_toolbar_get_style(self.CPointer)
	return
}

/*
Insert a #GtkToolItem into the toolbar at position @pos. If @pos is
0 the item is prepended to the start of the toolbar. If @pos is
negative, the item is appended to the end of the toolbar.
*/
func (self *_TraitToolbar) Insert(item *ToolItem, pos int) {
	C.gtk_toolbar_insert(self.CPointer, (*C.GtkToolItem)(item.CPointer), C.gint(pos))
	return
}

/*
Highlights @toolbar to give an idea of what it would look like
if @item was added to @toolbar at the position indicated by @index_.
If @item is %NULL, highlighting is turned off. In that case @index_
is ignored.

The @tool_item passed to this function must not be part of any widget
hierarchy. When an item is set as drop highlight item it can not
added to any widget hierarchy or used as highlight item for another
toolbar.
*/
func (self *_TraitToolbar) SetDropHighlightItem(tool_item *ToolItem, index_ int) {
	C.gtk_toolbar_set_drop_highlight_item(self.CPointer, (*C.GtkToolItem)(tool_item.CPointer), C.gint(index_))
	return
}

/*
This function sets the size of stock icons in the toolbar. You
can call it both before you add the icons and after they’ve been
added. The size you set will override user preferences for the default
icon size.

This should only be used for special-purpose toolbars, normal
application toolbars should respect the user preferences for the
size of icons.
*/
func (self *_TraitToolbar) SetIconSize(icon_size C.GtkIconSize) {
	C.gtk_toolbar_set_icon_size(self.CPointer, icon_size)
	return
}

/*
Sets whether to show an overflow menu when
@toolbar doesn’t have room for all items on it. If %TRUE,
items that there are not room are available through an
overflow menu.
*/
func (self *_TraitToolbar) SetShowArrow(show_arrow bool) {
	__cgo__show_arrow := C.gboolean(0)
	if show_arrow {
		__cgo__show_arrow = C.gboolean(1)
	}
	C.gtk_toolbar_set_show_arrow(self.CPointer, __cgo__show_arrow)
	return
}

/*
Alters the view of @toolbar to display either icons only, text only, or both.
*/
func (self *_TraitToolbar) SetStyle(style C.GtkToolbarStyle) {
	C.gtk_toolbar_set_style(self.CPointer, style)
	return
}

/*
Unsets toolbar icon size set with gtk_toolbar_set_icon_size(), so that
user preferences will be used to determine the icon size.
*/
func (self *_TraitToolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size(self.CPointer)
	return
}

/*
Unsets a toolbar style set with gtk_toolbar_set_style(), so that
user preferences will be used to determine the toolbar style.
*/
func (self *_TraitToolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style(self.CPointer)
	return
}

type _TraitTooltip struct{ CPointer *C.GtkTooltip }

/*
Replaces the widget packed into the tooltip with
@custom_widget. @custom_widget does not get destroyed when the tooltip goes
away.
By default a box with a #GtkImage and #GtkLabel is embedded in
the tooltip, which can be configured using gtk_tooltip_set_markup()
and gtk_tooltip_set_icon().
*/
func (self *_TraitTooltip) SetCustom(custom_widget *Widget) {
	C.gtk_tooltip_set_custom(self.CPointer, (*C.GtkWidget)(custom_widget.CPointer))
	return
}

/*
Sets the icon of the tooltip (which is in front of the text) to be
@pixbuf.  If @pixbuf is %NULL, the image will be hidden.
*/
func (self *_TraitTooltip) SetIcon(pixbuf *C.GdkPixbuf) {
	C.gtk_tooltip_set_icon(self.CPointer, pixbuf)
	return
}

/*
Sets the icon of the tooltip (which is in front of the text)
to be the icon indicated by @gicon with the size indicated
by @size. If @gicon is %NULL, the image will be hidden.
*/
func (self *_TraitTooltip) SetIconFromGicon(gicon *C.GIcon, size C.GtkIconSize) {
	C.gtk_tooltip_set_icon_from_gicon(self.CPointer, gicon, size)
	return
}

/*
Sets the icon of the tooltip (which is in front of the text) to be
the icon indicated by @icon_name with the size indicated
by @size.  If @icon_name is %NULL, the image will be hidden.
*/
func (self *_TraitTooltip) SetIconFromIconName(icon_name string, size C.GtkIconSize) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_tooltip_set_icon_from_icon_name(self.CPointer, __cgo__icon_name, size)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

// gtk_tooltip_set_icon_from_stock is not generated due to deprecation attr

/*
Sets the text of the tooltip to be @markup, which is marked up
with the [Pango text markup language][PangoMarkupFormat].
If @markup is %NULL, the label will be hidden.
*/
func (self *_TraitTooltip) SetMarkup(markup string) {
	__cgo__markup := (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	C.gtk_tooltip_set_markup(self.CPointer, __cgo__markup)
	C.free(unsafe.Pointer(__cgo__markup))
	return
}

/*
Sets the text of the tooltip to be @text. If @text is %NULL, the label
will be hidden. See also gtk_tooltip_set_markup().
*/
func (self *_TraitTooltip) SetText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_tooltip_set_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Sets the area of the widget, where the contents of this tooltip apply,
to be @rect (in widget coordinates).  This is especially useful for
properly setting tooltips on #GtkTreeView rows and cells, #GtkIconViews,
etc.

For setting tooltips on #GtkTreeView, please refer to the convenience
functions for this: gtk_tree_view_set_tooltip_row() and
gtk_tree_view_set_tooltip_cell().
*/
func (self *_TraitTooltip) SetTipArea(rect *C.GdkRectangle) {
	C.gtk_tooltip_set_tip_area(self.CPointer, rect)
	return
}

type _TraitTreeModelFilter struct{ CPointer *C.GtkTreeModelFilter }

/*
This function should almost never be called. It clears the @filter
of any cached iterators that haven’t been reffed with
gtk_tree_model_ref_node(). This might be useful if the child model
being filtered is static (and doesn’t change often) and there has been
a lot of unreffed access to nodes. As a side effect of this function,
all unreffed iters will be invalid.
*/
func (self *_TraitTreeModelFilter) ClearCache() {
	C.gtk_tree_model_filter_clear_cache(self.CPointer)
	return
}

/*
Sets @filter_iter to point to the row in @filter that corresponds to the
row pointed at by @child_iter.  If @filter_iter was not set, %FALSE is
returned.
*/
func (self *_TraitTreeModelFilter) ConvertChildIterToIter(child_iter *TreeIter) (filter_iter C.GtkTreeIter, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_model_filter_convert_child_iter_to_iter(self.CPointer, &filter_iter, (*C.GtkTreeIter)(unsafe.Pointer(child_iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts @child_path to a path relative to @filter. That is, @child_path
points to a path in the child model. The rerturned path will point to the
same row in the filtered model. If @child_path isn’t a valid path on the
child model or points to a row which is not visible in @filter, then %NULL
is returned.
*/
func (self *_TraitTreeModelFilter) ConvertChildPathToPath(child_path *TreePath) (return__ *TreePath) {
	var __cgo__return__ *C.GtkTreePath
	__cgo__return__ = C.gtk_tree_model_filter_convert_child_path_to_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(child_path)))
	return__ = (*TreePath)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Sets @child_iter to point to the row pointed to by @filter_iter.
*/
func (self *_TraitTreeModelFilter) ConvertIterToChildIter(filter_iter *TreeIter) (child_iter C.GtkTreeIter) {
	C.gtk_tree_model_filter_convert_iter_to_child_iter(self.CPointer, &child_iter, (*C.GtkTreeIter)(unsafe.Pointer(filter_iter)))
	return
}

/*
Converts @filter_path to a path on the child model of @filter. That is,
@filter_path points to a location in @filter. The returned path will
point to the same location in the model not being filtered. If @filter_path
does not point to a location in the child model, %NULL is returned.
*/
func (self *_TraitTreeModelFilter) ConvertPathToChildPath(filter_path *TreePath) (return__ *TreePath) {
	var __cgo__return__ *C.GtkTreePath
	__cgo__return__ = C.gtk_tree_model_filter_convert_path_to_child_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(filter_path)))
	return__ = (*TreePath)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns a pointer to the child model of @filter.
*/
func (self *_TraitTreeModelFilter) GetModel() (return__ *C.GtkTreeModel) {
	return__ = C.gtk_tree_model_filter_get_model(self.CPointer)
	return
}

/*
Emits ::row_changed for each row in the child model, which causes
the filter to re-evaluate whether a row is visible or not.
*/
func (self *_TraitTreeModelFilter) Refilter() {
	C.gtk_tree_model_filter_refilter(self.CPointer)
	return
}

/*
With the @n_columns and @types parameters, you give an array of column
types for this model (which will be exposed to the parent model/view).
The @func, @data and @destroy parameters are for specifying the modify
function. The modify function will get called for each
data access, the goal of the modify function is to return the data which
should be displayed at the location specified using the parameters of the
modify function.
*/
func (self *_TraitTreeModelFilter) SetModifyFunc(n_columns int, types *C.GType, func_ C.GtkTreeModelFilterModifyFunc, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_tree_model_filter_set_modify_func(self.CPointer, C.gint(n_columns), types, func_, (C.gpointer)(data), destroy)
	return
}

/*
Sets @column of the child_model to be the column where @filter should
look for visibility information. @columns should be a column of type
%G_TYPE_BOOLEAN, where %TRUE means that a row is visible, and %FALSE
if not.
*/
func (self *_TraitTreeModelFilter) SetVisibleColumn(column int) {
	C.gtk_tree_model_filter_set_visible_column(self.CPointer, C.gint(column))
	return
}

/*
Sets the visible function used when filtering the @filter to be @func. The
function should return %TRUE if the given row should be visible and
%FALSE otherwise.

If the condition calculated by the function changes over time (e.g. because
it depends on some global parameters), you must call
gtk_tree_model_filter_refilter() to keep the visibility information of
the model uptodate.

Note that @func is called whenever a row is inserted, when it may still be
empty. The visible function should therefore take special care of empty
rows, like in the example below.

|[<!-- language="C" -->
static gboolean
visible_func (GtkTreeModel *model,
              GtkTreeIter  *iter,
              gpointer      data)
{
  // Visible if row is non-empty and first column is “HI”
  gchar *str;
  gboolean visible = FALSE;

  gtk_tree_model_get (model, iter, 0, &str, -1);
  if (str && strcmp (str, "HI") == 0)
    visible = TRUE;
  g_free (str);

  return visible;
}
]|
*/
func (self *_TraitTreeModelFilter) SetVisibleFunc(func_ C.GtkTreeModelFilterVisibleFunc, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_tree_model_filter_set_visible_func(self.CPointer, func_, (C.gpointer)(data), destroy)
	return
}

type _TraitTreeModelSort struct{ CPointer *C.GtkTreeModelSort }

/*
This function should almost never be called.  It clears the @tree_model_sort
of any cached iterators that haven’t been reffed with
gtk_tree_model_ref_node().  This might be useful if the child model being
sorted is static (and doesn’t change often) and there has been a lot of
unreffed access to nodes.  As a side effect of this function, all unreffed
iters will be invalid.
*/
func (self *_TraitTreeModelSort) ClearCache() {
	C.gtk_tree_model_sort_clear_cache(self.CPointer)
	return
}

/*
Sets @sort_iter to point to the row in @tree_model_sort that corresponds to
the row pointed at by @child_iter.  If @sort_iter was not set, %FALSE
is returned.  Note: a boolean is only returned since 2.14.
*/
func (self *_TraitTreeModelSort) ConvertChildIterToIter(child_iter *TreeIter) (sort_iter C.GtkTreeIter, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_model_sort_convert_child_iter_to_iter(self.CPointer, &sort_iter, (*C.GtkTreeIter)(unsafe.Pointer(child_iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Converts @child_path to a path relative to @tree_model_sort.  That is,
@child_path points to a path in the child model.  The returned path will
point to the same row in the sorted model.  If @child_path isn’t a valid
path on the child model, then %NULL is returned.
*/
func (self *_TraitTreeModelSort) ConvertChildPathToPath(child_path *TreePath) (return__ *TreePath) {
	var __cgo__return__ *C.GtkTreePath
	__cgo__return__ = C.gtk_tree_model_sort_convert_child_path_to_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(child_path)))
	return__ = (*TreePath)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Sets @child_iter to point to the row pointed to by @sorted_iter.
*/
func (self *_TraitTreeModelSort) ConvertIterToChildIter(sorted_iter *TreeIter) (child_iter C.GtkTreeIter) {
	C.gtk_tree_model_sort_convert_iter_to_child_iter(self.CPointer, &child_iter, (*C.GtkTreeIter)(unsafe.Pointer(sorted_iter)))
	return
}

/*
Converts @sorted_path to a path on the child model of @tree_model_sort.
That is, @sorted_path points to a location in @tree_model_sort.  The
returned path will point to the same location in the model not being
sorted.  If @sorted_path does not point to a location in the child model,
%NULL is returned.
*/
func (self *_TraitTreeModelSort) ConvertPathToChildPath(sorted_path *TreePath) (return__ *TreePath) {
	var __cgo__return__ *C.GtkTreePath
	__cgo__return__ = C.gtk_tree_model_sort_convert_path_to_child_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(sorted_path)))
	return__ = (*TreePath)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns the model the #GtkTreeModelSort is sorting.
*/
func (self *_TraitTreeModelSort) GetModel() (return__ *C.GtkTreeModel) {
	return__ = C.gtk_tree_model_sort_get_model(self.CPointer)
	return
}

/*
> This function is slow. Only use it for debugging and/or testing
> purposes.

Checks if the given iter is a valid iter for this #GtkTreeModelSort.
*/
func (self *_TraitTreeModelSort) IterIsValid(iter *TreeIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_model_sort_iter_is_valid(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This resets the default sort function to be in the “unsorted” state.  That
is, it is in the same order as the child model. It will re-sort the model
to be in the same order as the child model only if the #GtkTreeModelSort
is in “unsorted” state.
*/
func (self *_TraitTreeModelSort) ResetDefaultSortFunc() {
	C.gtk_tree_model_sort_reset_default_sort_func(self.CPointer)
	return
}

type _TraitTreeSelection struct{ CPointer *C.GtkTreeSelection }

/*
Returns the number of rows that have been selected in @tree.
*/
func (self *_TraitTreeSelection) CountSelectedRows() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_selection_count_selected_rows(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the selection mode for @selection. See
gtk_tree_selection_set_mode().
*/
func (self *_TraitTreeSelection) GetMode() (return__ C.GtkSelectionMode) {
	return__ = C.gtk_tree_selection_get_mode(self.CPointer)
	return
}

/*
Returns the current selection function.
*/
func (self *_TraitTreeSelection) GetSelectFunction() (return__ C.GtkTreeSelectionFunc) {
	return__ = C.gtk_tree_selection_get_select_function(self.CPointer)
	return
}

/*
Sets @iter to the currently selected node if @selection is set to
#GTK_SELECTION_SINGLE or #GTK_SELECTION_BROWSE.  @iter may be NULL if you
just want to test if @selection has any selected nodes.  @model is filled
with the current model as a convenience.  This function will not work if you
use @selection is #GTK_SELECTION_MULTIPLE.
*/
func (self *_TraitTreeSelection) GetSelected() (model *C.GtkTreeModel, iter C.GtkTreeIter, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_selection_get_selected(self.CPointer, &model, &iter)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Creates a list of path of all selected rows. Additionally, if you are
planning on modifying the model after calling this function, you may
want to convert the returned list into a list of #GtkTreeRowReferences.
To do this, you can use gtk_tree_row_reference_new().

To free the return value, use:
|[<!-- language="C" -->
g_list_free_full (list, (GDestroyNotify) gtk_tree_path_free);
]|
*/
func (self *_TraitTreeSelection) GetSelectedRows() (model *C.GtkTreeModel, return__ *C.GList) {
	return__ = C.gtk_tree_selection_get_selected_rows(self.CPointer, &model)
	return
}

/*
Returns the tree view associated with @selection.
*/
func (self *_TraitTreeSelection) GetTreeView() (return__ *TreeView) {
	var __cgo__return__ *C.GtkTreeView
	__cgo__return__ = C.gtk_tree_selection_get_tree_view(self.CPointer)
	return__ = NewTreeViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the user data for the selection function.
*/
func (self *_TraitTreeSelection) GetUserData() (return__ unsafe.Pointer) {
	var __cgo__return__ C.gpointer
	__cgo__return__ = C.gtk_tree_selection_get_user_data(self.CPointer)
	return__ = unsafe.Pointer(__cgo__return__)
	return
}

/*
Returns %TRUE if the row at @iter is currently selected.
*/
func (self *_TraitTreeSelection) IterIsSelected(iter *TreeIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_selection_iter_is_selected(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the row pointed to by @path is currently selected.  If @path
does not point to a valid location, %FALSE is returned
*/
func (self *_TraitTreeSelection) PathIsSelected(path *TreePath) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_selection_path_is_selected(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Selects all the nodes. @selection must be set to #GTK_SELECTION_MULTIPLE
mode.
*/
func (self *_TraitTreeSelection) SelectAll() {
	C.gtk_tree_selection_select_all(self.CPointer)
	return
}

/*
Selects the specified iterator.
*/
func (self *_TraitTreeSelection) SelectIter(iter *TreeIter) {
	C.gtk_tree_selection_select_iter(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return
}

/*
Select the row at @path.
*/
func (self *_TraitTreeSelection) SelectPath(path *TreePath) {
	C.gtk_tree_selection_select_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Selects a range of nodes, determined by @start_path and @end_path inclusive.
@selection must be set to #GTK_SELECTION_MULTIPLE mode.
*/
func (self *_TraitTreeSelection) SelectRange(start_path *TreePath, end_path *TreePath) {
	C.gtk_tree_selection_select_range(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(start_path)), (*C.GtkTreePath)(unsafe.Pointer(end_path)))
	return
}

/*
Calls a function for each selected node. Note that you cannot modify
the tree or selection from within this function. As a result,
gtk_tree_selection_get_selected_rows() might be more useful.
*/
func (self *_TraitTreeSelection) SelectedForeach(func_ C.GtkTreeSelectionForeachFunc, data unsafe.Pointer) {
	C.gtk_tree_selection_selected_foreach(self.CPointer, func_, (C.gpointer)(data))
	return
}

/*
Sets the selection mode of the @selection.  If the previous type was
#GTK_SELECTION_MULTIPLE, then the anchor is kept selected, if it was
previously selected.
*/
func (self *_TraitTreeSelection) SetMode(type_ C.GtkSelectionMode) {
	C.gtk_tree_selection_set_mode(self.CPointer, type_)
	return
}

/*
Sets the selection function.

If set, this function is called before any node is selected or unselected,
giving some control over which nodes are selected. The select function
should return %TRUE if the state of the node may be toggled, and %FALSE
if the state of the node should be left unchanged.
*/
func (self *_TraitTreeSelection) SetSelectFunction(func_ C.GtkTreeSelectionFunc, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_tree_selection_set_select_function(self.CPointer, func_, (C.gpointer)(data), destroy)
	return
}

/*
Unselects all the nodes.
*/
func (self *_TraitTreeSelection) UnselectAll() {
	C.gtk_tree_selection_unselect_all(self.CPointer)
	return
}

/*
Unselects the specified iterator.
*/
func (self *_TraitTreeSelection) UnselectIter(iter *TreeIter) {
	C.gtk_tree_selection_unselect_iter(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return
}

/*
Unselects the row at @path.
*/
func (self *_TraitTreeSelection) UnselectPath(path *TreePath) {
	C.gtk_tree_selection_unselect_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Unselects a range of nodes, determined by @start_path and @end_path
inclusive.
*/
func (self *_TraitTreeSelection) UnselectRange(start_path *TreePath, end_path *TreePath) {
	C.gtk_tree_selection_unselect_range(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(start_path)), (*C.GtkTreePath)(unsafe.Pointer(end_path)))
	return
}

type _TraitTreeStore struct{ CPointer *C.GtkTreeStore }

/*
Appends a new row to @tree_store.  If @parent is non-%NULL, then it will append the
new row after the last child of @parent, otherwise it will append a row to
the top level.  @iter will be changed to point to this new row.  The row will
be empty after this function is called.  To fill in values, you need to call
gtk_tree_store_set() or gtk_tree_store_set_value().
*/
func (self *_TraitTreeStore) Append(parent *TreeIter) (iter C.GtkTreeIter) {
	C.gtk_tree_store_append(self.CPointer, &iter, (*C.GtkTreeIter)(unsafe.Pointer(parent)))
	return
}

/*
Removes all rows from @tree_store
*/
func (self *_TraitTreeStore) Clear() {
	C.gtk_tree_store_clear(self.CPointer)
	return
}

/*
Creates a new row at @position.  If parent is non-%NULL, then the row will be
made a child of @parent.  Otherwise, the row will be created at the toplevel.
If @position is -1 or is larger than the number of rows at that level, then
the new row will be inserted to the end of the list.  @iter will be changed
to point to this new row.  The row will be empty after this function is
called.  To fill in values, you need to call gtk_tree_store_set() or
gtk_tree_store_set_value().
*/
func (self *_TraitTreeStore) Insert(parent *TreeIter, position int) (iter C.GtkTreeIter) {
	C.gtk_tree_store_insert(self.CPointer, &iter, (*C.GtkTreeIter)(unsafe.Pointer(parent)), C.gint(position))
	return
}

/*
Inserts a new row after @sibling.  If @sibling is %NULL, then the row will be
prepended to @parent ’s children.  If @parent and @sibling are %NULL, then
the row will be prepended to the toplevel.  If both @sibling and @parent are
set, then @parent must be the parent of @sibling.  When @sibling is set,
@parent is optional.

@iter will be changed to point to this new row.  The row will be empty after
this function is called.  To fill in values, you need to call
gtk_tree_store_set() or gtk_tree_store_set_value().
*/
func (self *_TraitTreeStore) InsertAfter(parent *TreeIter, sibling *TreeIter) (iter C.GtkTreeIter) {
	C.gtk_tree_store_insert_after(self.CPointer, &iter, (*C.GtkTreeIter)(unsafe.Pointer(parent)), (*C.GtkTreeIter)(unsafe.Pointer(sibling)))
	return
}

/*
Inserts a new row before @sibling.  If @sibling is %NULL, then the row will
be appended to @parent ’s children.  If @parent and @sibling are %NULL, then
the row will be appended to the toplevel.  If both @sibling and @parent are
set, then @parent must be the parent of @sibling.  When @sibling is set,
@parent is optional.

@iter will be changed to point to this new row.  The row will be empty after
this function is called.  To fill in values, you need to call
gtk_tree_store_set() or gtk_tree_store_set_value().
*/
func (self *_TraitTreeStore) InsertBefore(parent *TreeIter, sibling *TreeIter) (iter C.GtkTreeIter) {
	C.gtk_tree_store_insert_before(self.CPointer, &iter, (*C.GtkTreeIter)(unsafe.Pointer(parent)), (*C.GtkTreeIter)(unsafe.Pointer(sibling)))
	return
}

// gtk_tree_store_insert_with_values is not generated due to varargs

/*
A variant of gtk_tree_store_insert_with_values() which takes
the columns and values as two arrays, instead of varargs.  This
function is mainly intended for language bindings.
*/
func (self *_TraitTreeStore) InsertWithValuesv(parent *TreeIter, position int, columns []int, values *C.GValue, n_values int) (iter C.GtkTreeIter) {
	__header__columns := (*reflect.SliceHeader)(unsafe.Pointer(&columns))
	C.gtk_tree_store_insert_with_valuesv(self.CPointer, &iter, (*C.GtkTreeIter)(unsafe.Pointer(parent)), C.gint(position), (*C.gint)(unsafe.Pointer(__header__columns.Data)), values, C.gint(n_values))
	return
}

/*
Returns %TRUE if @iter is an ancestor of @descendant.  That is, @iter is the
parent (or grandparent or great-grandparent) of @descendant.
*/
func (self *_TraitTreeStore) IsAncestor(iter *TreeIter, descendant *TreeIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_store_is_ancestor(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), (*C.GtkTreeIter)(unsafe.Pointer(descendant)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the depth of @iter.  This will be 0 for anything on the root level, 1
for anything down a level, etc.
*/
func (self *_TraitTreeStore) IterDepth(iter *TreeIter) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_store_iter_depth(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return__ = int(__cgo__return__)
	return
}

/*
WARNING: This function is slow. Only use it for debugging and/or testing
purposes.

Checks if the given iter is a valid iter for this #GtkTreeStore.
*/
func (self *_TraitTreeStore) IterIsValid(iter *TreeIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_store_iter_is_valid(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Moves @iter in @tree_store to the position after @position. @iter and
@position should be in the same level. Note that this function only
works with unsorted stores. If @position is %NULL, @iter will be moved
to the start of the level.
*/
func (self *_TraitTreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	C.gtk_tree_store_move_after(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), (*C.GtkTreeIter)(unsafe.Pointer(position)))
	return
}

/*
Moves @iter in @tree_store to the position before @position. @iter and
@position should be in the same level. Note that this function only
works with unsorted stores. If @position is %NULL, @iter will be
moved to the end of the level.
*/
func (self *_TraitTreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	C.gtk_tree_store_move_before(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), (*C.GtkTreeIter)(unsafe.Pointer(position)))
	return
}

/*
Prepends a new row to @tree_store.  If @parent is non-%NULL, then it will prepend
the new row before the first child of @parent, otherwise it will prepend a row
to the top level.  @iter will be changed to point to this new row.  The row
will be empty after this function is called.  To fill in values, you need to
call gtk_tree_store_set() or gtk_tree_store_set_value().
*/
func (self *_TraitTreeStore) Prepend(parent *TreeIter) (iter C.GtkTreeIter) {
	C.gtk_tree_store_prepend(self.CPointer, &iter, (*C.GtkTreeIter)(unsafe.Pointer(parent)))
	return
}

/*
Removes @iter from @tree_store.  After being removed, @iter is set to the
next valid row at that level, or invalidated if it previously pointed to the
last one.
*/
func (self *_TraitTreeStore) Remove(iter *TreeIter) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_store_remove(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Reorders the children of @parent in @tree_store to follow the order
indicated by @new_order. Note that this function only works with
unsorted stores.
*/
func (self *_TraitTreeStore) Reorder(parent *TreeIter, new_order *C.gint) {
	C.gtk_tree_store_reorder(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(parent)), new_order)
	return
}

// gtk_tree_store_set is not generated due to varargs

/*
This function is meant primarily for #GObjects that inherit from
#GtkTreeStore, and should only be used when constructing a new
#GtkTreeStore.  It will not function after a row has been added,
or a method on the #GtkTreeModel interface is called.
*/
func (self *_TraitTreeStore) SetColumnTypes(n_columns int, types *C.GType) {
	C.gtk_tree_store_set_column_types(self.CPointer, C.gint(n_columns), types)
	return
}

// gtk_tree_store_set_valist is not generated due to varargs

/*
Sets the data in the cell specified by @iter and @column.
The type of @value must be convertible to the type of the
column.
*/
func (self *_TraitTreeStore) SetValue(iter *TreeIter, column int, value *C.GValue) {
	C.gtk_tree_store_set_value(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), C.gint(column), value)
	return
}

/*
A variant of gtk_tree_store_set_valist() which takes
the columns and values as two arrays, instead of varargs.  This
function is mainly intended for language bindings or in case
the number of columns to change is not known until run-time.
*/
func (self *_TraitTreeStore) SetValuesv(iter *TreeIter, columns []int, values *C.GValue, n_values int) {
	__header__columns := (*reflect.SliceHeader)(unsafe.Pointer(&columns))
	C.gtk_tree_store_set_valuesv(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(iter)), (*C.gint)(unsafe.Pointer(__header__columns.Data)), values, C.gint(n_values))
	return
}

/*
Swaps @a and @b in the same level of @tree_store. Note that this function
only works with unsorted stores.
*/
func (self *_TraitTreeStore) Swap(a *TreeIter, b *TreeIter) {
	C.gtk_tree_store_swap(self.CPointer, (*C.GtkTreeIter)(unsafe.Pointer(a)), (*C.GtkTreeIter)(unsafe.Pointer(b)))
	return
}

type _TraitTreeView struct{ CPointer *C.GtkTreeView }

/*
Appends @column to the list of columns. If @tree_view has “fixed_height”
mode enabled, then @column must have its “sizing” property set to be
GTK_TREE_VIEW_COLUMN_FIXED.
*/
func (self *_TraitTreeView) AppendColumn(column *TreeViewColumn) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_append_column(self.CPointer, (*C.GtkTreeViewColumn)(column.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Recursively collapses all visible, expanded nodes in @tree_view.
*/
func (self *_TraitTreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(self.CPointer)
	return
}

/*
Collapses a row (hides its child rows, if they exist).
*/
func (self *_TraitTreeView) CollapseRow(path *TreePath) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_collapse_row(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Resizes all columns to their optimal width. Only works after the
treeview has been realized.
*/
func (self *_TraitTreeView) ColumnsAutosize() {
	C.gtk_tree_view_columns_autosize(self.CPointer)
	return
}

/*
Converts bin_window coordinates to coordinates for the
tree (the full scrollable area of the tree).
*/
func (self *_TraitTreeView) ConvertBinWindowToTreeCoords(bx int, by int) (tx int, ty int) {
	var __cgo__tx C.gint
	var __cgo__ty C.gint
	C.gtk_tree_view_convert_bin_window_to_tree_coords(self.CPointer, C.gint(bx), C.gint(by), &__cgo__tx, &__cgo__ty)
	tx = int(__cgo__tx)
	ty = int(__cgo__ty)
	return
}

/*
Converts bin_window coordinates (see gtk_tree_view_get_bin_window())
to widget relative coordinates.
*/
func (self *_TraitTreeView) ConvertBinWindowToWidgetCoords(bx int, by int) (wx int, wy int) {
	var __cgo__wx C.gint
	var __cgo__wy C.gint
	C.gtk_tree_view_convert_bin_window_to_widget_coords(self.CPointer, C.gint(bx), C.gint(by), &__cgo__wx, &__cgo__wy)
	wx = int(__cgo__wx)
	wy = int(__cgo__wy)
	return
}

/*
Converts tree coordinates (coordinates in full scrollable area of the tree)
to bin_window coordinates.
*/
func (self *_TraitTreeView) ConvertTreeToBinWindowCoords(tx int, ty int) (bx int, by int) {
	var __cgo__bx C.gint
	var __cgo__by C.gint
	C.gtk_tree_view_convert_tree_to_bin_window_coords(self.CPointer, C.gint(tx), C.gint(ty), &__cgo__bx, &__cgo__by)
	bx = int(__cgo__bx)
	by = int(__cgo__by)
	return
}

/*
Converts tree coordinates (coordinates in full scrollable area of the tree)
to widget coordinates.
*/
func (self *_TraitTreeView) ConvertTreeToWidgetCoords(tx int, ty int) (wx int, wy int) {
	var __cgo__wx C.gint
	var __cgo__wy C.gint
	C.gtk_tree_view_convert_tree_to_widget_coords(self.CPointer, C.gint(tx), C.gint(ty), &__cgo__wx, &__cgo__wy)
	wx = int(__cgo__wx)
	wy = int(__cgo__wy)
	return
}

/*
Converts widget coordinates to coordinates for the bin_window
(see gtk_tree_view_get_bin_window()).
*/
func (self *_TraitTreeView) ConvertWidgetToBinWindowCoords(wx int, wy int) (bx int, by int) {
	var __cgo__bx C.gint
	var __cgo__by C.gint
	C.gtk_tree_view_convert_widget_to_bin_window_coords(self.CPointer, C.gint(wx), C.gint(wy), &__cgo__bx, &__cgo__by)
	bx = int(__cgo__bx)
	by = int(__cgo__by)
	return
}

/*
Converts widget coordinates to coordinates for the
tree (the full scrollable area of the tree).
*/
func (self *_TraitTreeView) ConvertWidgetToTreeCoords(wx int, wy int) (tx int, ty int) {
	var __cgo__tx C.gint
	var __cgo__ty C.gint
	C.gtk_tree_view_convert_widget_to_tree_coords(self.CPointer, C.gint(wx), C.gint(wy), &__cgo__tx, &__cgo__ty)
	tx = int(__cgo__tx)
	ty = int(__cgo__ty)
	return
}

/*
Creates a #cairo_surface_t representation of the row at @path.
This image is used for a drag icon.
*/
func (self *_TraitTreeView) CreateRowDragIcon(path *TreePath) (return__ *C.cairo_surface_t) {
	return__ = C.gtk_tree_view_create_row_drag_icon(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Turns @tree_view into a drop destination for automatic DND. Calling
this method sets #GtkTreeView:reorderable to %FALSE.
*/
func (self *_TraitTreeView) EnableModelDragDest(targets *TargetEntry, n_targets int, actions C.GdkDragAction) {
	C.gtk_tree_view_enable_model_drag_dest(self.CPointer, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.gint(n_targets), actions)
	return
}

/*
Turns @tree_view into a drag source for automatic DND. Calling this
method sets #GtkTreeView:reorderable to %FALSE.
*/
func (self *_TraitTreeView) EnableModelDragSource(start_button_mask C.GdkModifierType, targets *TargetEntry, n_targets int, actions C.GdkDragAction) {
	C.gtk_tree_view_enable_model_drag_source(self.CPointer, start_button_mask, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.gint(n_targets), actions)
	return
}

/*
Recursively expands all nodes in the @tree_view.
*/
func (self *_TraitTreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(self.CPointer)
	return
}

/*
Opens the row so its children are visible.
*/
func (self *_TraitTreeView) ExpandRow(path *TreePath, open_all bool) (return__ bool) {
	__cgo__open_all := C.gboolean(0)
	if open_all {
		__cgo__open_all = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_expand_row(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), __cgo__open_all)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Expands the row at @path. This will also expand all parent rows of
@path as necessary.
*/
func (self *_TraitTreeView) ExpandToPath(path *TreePath) {
	C.gtk_tree_view_expand_to_path(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

/*
Gets the setting set by gtk_tree_view_set_activate_on_single_click().
*/
func (self *_TraitTreeView) GetActivateOnSingleClick() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_activate_on_single_click(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Fills the bounding rectangle in bin_window coordinates for the cell at the
row specified by @path and the column specified by @column.  If @path is
%NULL, or points to a node not found in the tree, the @y and @height fields of
the rectangle will be filled with 0. If @column is %NULL, the @x and @width
fields will be filled with 0.  The returned rectangle is equivalent to the
@background_area passed to gtk_cell_renderer_render().  These background
areas tile to cover the entire bin window.  Contrast with the @cell_area,
returned by gtk_tree_view_get_cell_area(), which returns only the cell
itself, excluding surrounding borders and the tree expander area.
*/
func (self *_TraitTreeView) GetBackgroundArea(path *TreePath, column *TreeViewColumn) (rect C.GdkRectangle) {
	C.gtk_tree_view_get_background_area(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkTreeViewColumn)(column.CPointer), &rect)
	return
}

/*
Returns the window that @tree_view renders to.
This is used primarily to compare to `event->window`
to confirm that the event on @tree_view is on the right window.
*/
func (self *_TraitTreeView) GetBinWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_tree_view_get_bin_window(self.CPointer)
	return
}

/*
Fills the bounding rectangle in bin_window coordinates for the cell at the
row specified by @path and the column specified by @column.  If @path is
%NULL, or points to a path not currently displayed, the @y and @height fields
of the rectangle will be filled with 0. If @column is %NULL, the @x and @width
fields will be filled with 0.  The sum of all cell rects does not cover the
entire tree; there are extra pixels in between rows, for example. The
returned rectangle is equivalent to the @cell_area passed to
gtk_cell_renderer_render().  This function is only valid if @tree_view is
realized.
*/
func (self *_TraitTreeView) GetCellArea(path *TreePath, column *TreeViewColumn) (rect C.GdkRectangle) {
	C.gtk_tree_view_get_cell_area(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkTreeViewColumn)(column.CPointer), &rect)
	return
}

/*
Gets the #GtkTreeViewColumn at the given position in the #tree_view.
*/
func (self *_TraitTreeView) GetColumn(n int) (return__ *TreeViewColumn) {
	var __cgo__return__ *C.GtkTreeViewColumn
	__cgo__return__ = C.gtk_tree_view_get_column(self.CPointer, C.gint(n))
	return__ = NewTreeViewColumnFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns a #GList of all the #GtkTreeViewColumn s currently in @tree_view.
The returned list must be freed with g_list_free ().
*/
func (self *_TraitTreeView) GetColumns() (return__ *C.GList) {
	return__ = C.gtk_tree_view_get_columns(self.CPointer)
	return
}

/*
Fills in @path and @focus_column with the current path and focus column.  If
the cursor isn’t currently set, then *@path will be %NULL.  If no column
currently has focus, then *@focus_column will be %NULL.

The returned #GtkTreePath must be freed with gtk_tree_path_free() when
you are done with it.
*/
func (self *_TraitTreeView) GetCursor() (path *TreePath, focus_column *TreeViewColumn) {
	var __cgo__path *C.GtkTreePath
	var __cgo__focus_column *C.GtkTreeViewColumn
	C.gtk_tree_view_get_cursor(self.CPointer, &__cgo__path, &__cgo__focus_column)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	focus_column = NewTreeViewColumnFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__focus_column).Pointer()))
	return
}

/*
Determines the destination row for a given position.  @drag_x and
@drag_y are expected to be in widget coordinates.  This function is only
meaningful if @tree_view is realized.  Therefore this function will always
return %FALSE if @tree_view is not realized or does not have a model.
*/
func (self *_TraitTreeView) GetDestRowAtPos(drag_x int, drag_y int) (path *TreePath, pos C.GtkTreeViewDropPosition, return__ bool) {
	var __cgo__path *C.GtkTreePath
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_dest_row_at_pos(self.CPointer, C.gint(drag_x), C.gint(drag_y), &__cgo__path, &pos)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets information about the row that is highlighted for feedback.
*/
func (self *_TraitTreeView) GetDragDestRow() (path *TreePath, pos C.GtkTreeViewDropPosition) {
	var __cgo__path *C.GtkTreePath
	C.gtk_tree_view_get_drag_dest_row(self.CPointer, &__cgo__path, &pos)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	return
}

/*
Returns whether or not the tree allows to start interactive searching
by typing in text.
*/
func (self *_TraitTreeView) GetEnableSearch() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_enable_search(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether or not tree lines are drawn in @tree_view.
*/
func (self *_TraitTreeView) GetEnableTreeLines() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_enable_tree_lines(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the column that is the current expander column.
This column has the expander arrow drawn next to it.
*/
func (self *_TraitTreeView) GetExpanderColumn() (return__ *TreeViewColumn) {
	var __cgo__return__ *C.GtkTreeViewColumn
	__cgo__return__ = C.gtk_tree_view_get_expander_column(self.CPointer)
	return__ = NewTreeViewColumnFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether fixed height mode is turned on for @tree_view.
*/
func (self *_TraitTreeView) GetFixedHeightMode() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_fixed_height_mode(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns which grid lines are enabled in @tree_view.
*/
func (self *_TraitTreeView) GetGridLines() (return__ C.GtkTreeViewGridLines) {
	return__ = C.gtk_tree_view_get_grid_lines(self.CPointer)
	return
}

// gtk_tree_view_get_hadjustment is not generated due to deprecation attr

/*
Returns whether all header columns are clickable.
*/
func (self *_TraitTreeView) GetHeadersClickable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_headers_clickable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the headers on the @tree_view are visible.
*/
func (self *_TraitTreeView) GetHeadersVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_headers_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether hover expansion mode is turned on for @tree_view.
*/
func (self *_TraitTreeView) GetHoverExpand() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_hover_expand(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether hover selection mode is turned on for @tree_view.
*/
func (self *_TraitTreeView) GetHoverSelection() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_hover_selection(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the amount, in pixels, of extra indentation for child levels
in @tree_view.
*/
func (self *_TraitTreeView) GetLevelIndentation() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_get_level_indentation(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the model the #GtkTreeView is based on.  Returns %NULL if the
model is unset.
*/
func (self *_TraitTreeView) GetModel() (return__ *C.GtkTreeModel) {
	return__ = C.gtk_tree_view_get_model(self.CPointer)
	return
}

/*
Queries the number of columns in the given @tree_view.
*/
func (self *_TraitTreeView) GetNColumns() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_tree_view_get_n_columns(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Finds the path at the point (@x, @y), relative to bin_window coordinates
(please see gtk_tree_view_get_bin_window()).
That is, @x and @y are relative to an events coordinates. @x and @y must
come from an event on the @tree_view only where `event->window ==
gtk_tree_view_get_bin_window ()`. It is primarily for
things like popup menus. If @path is non-%NULL, then it will be filled
with the #GtkTreePath at that point.  This path should be freed with
gtk_tree_path_free().  If @column is non-%NULL, then it will be filled
with the column at that point.  @cell_x and @cell_y return the coordinates
relative to the cell background (i.e. the @background_area passed to
gtk_cell_renderer_render()).  This function is only meaningful if
@tree_view is realized.  Therefore this function will always return %FALSE
if @tree_view is not realized or does not have a model.

For converting widget coordinates (eg. the ones you get from
GtkWidget::query-tooltip), please see
gtk_tree_view_convert_widget_to_bin_window_coords().
*/
func (self *_TraitTreeView) GetPathAtPos(x int, y int) (path *TreePath, column *TreeViewColumn, cell_x int, cell_y int, return__ bool) {
	var __cgo__path *C.GtkTreePath
	var __cgo__column *C.GtkTreeViewColumn
	var __cgo__cell_x C.gint
	var __cgo__cell_y C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_path_at_pos(self.CPointer, C.gint(x), C.gint(y), &__cgo__path, &__cgo__column, &__cgo__cell_x, &__cgo__cell_y)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	column = NewTreeViewColumnFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__column).Pointer()))
	cell_x = int(__cgo__cell_x)
	cell_y = int(__cgo__cell_y)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves whether the user can reorder the tree via drag-and-drop. See
gtk_tree_view_set_reorderable().
*/
func (self *_TraitTreeView) GetReorderable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_reorderable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current row separator function.
*/
func (self *_TraitTreeView) GetRowSeparatorFunc() (return__ C.GtkTreeViewRowSeparatorFunc) {
	return__ = C.gtk_tree_view_get_row_separator_func(self.CPointer)
	return
}

/*
Returns whether rubber banding is turned on for @tree_view.  If the
selection mode is #GTK_SELECTION_MULTIPLE, rubber banding will allow the
user to select multiple rows by dragging the mouse.
*/
func (self *_TraitTreeView) GetRubberBanding() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_rubber_banding(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the setting set by gtk_tree_view_set_rules_hint().
*/
func (self *_TraitTreeView) GetRulesHint() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_rules_hint(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the column searched on by the interactive search code.
*/
func (self *_TraitTreeView) GetSearchColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_get_search_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the #GtkEntry which is currently in use as interactive search
entry for @tree_view.  In case the built-in entry is being used, %NULL
will be returned.
*/
func (self *_TraitTreeView) GetSearchEntry() (return__ *Entry) {
	var __cgo__return__ *C.GtkEntry
	__cgo__return__ = C.gtk_tree_view_get_search_entry(self.CPointer)
	return__ = NewEntryFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the compare function currently in use.
*/
func (self *_TraitTreeView) GetSearchEqualFunc() (return__ C.GtkTreeViewSearchEqualFunc) {
	return__ = C.gtk_tree_view_get_search_equal_func(self.CPointer)
	return
}

/*
Returns the positioning function currently in use.
*/
func (self *_TraitTreeView) GetSearchPositionFunc() (return__ C.GtkTreeViewSearchPositionFunc) {
	return__ = C.gtk_tree_view_get_search_position_func(self.CPointer)
	return
}

/*
Gets the #GtkTreeSelection associated with @tree_view.
*/
func (self *_TraitTreeView) GetSelection() (return__ *TreeSelection) {
	var __cgo__return__ *C.GtkTreeSelection
	__cgo__return__ = C.gtk_tree_view_get_selection(self.CPointer)
	return__ = NewTreeSelectionFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether or not expanders are drawn in @tree_view.
*/
func (self *_TraitTreeView) GetShowExpanders() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_show_expanders(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the column of @tree_view’s model which is being used for
displaying tooltips on @tree_view’s rows.
*/
func (self *_TraitTreeView) GetTooltipColumn() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_get_tooltip_column(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

// gtk_tree_view_get_tooltip_context is not generated due to inout param

// gtk_tree_view_get_vadjustment is not generated due to deprecation attr

/*
Sets @start_path and @end_path to be the first and last visible path.
Note that there may be invisible paths in between.

The paths should be freed with gtk_tree_path_free() after use.
*/
func (self *_TraitTreeView) GetVisibleRange() (start_path *TreePath, end_path *TreePath, return__ bool) {
	var __cgo__start_path *C.GtkTreePath
	var __cgo__end_path *C.GtkTreePath
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_get_visible_range(self.CPointer, &__cgo__start_path, &__cgo__end_path)
	start_path = (*TreePath)(unsafe.Pointer(__cgo__start_path))
	end_path = (*TreePath)(unsafe.Pointer(__cgo__end_path))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Fills @visible_rect with the currently-visible region of the
buffer, in tree coordinates. Convert to bin_window coordinates with
gtk_tree_view_convert_tree_to_bin_window_coords().
Tree coordinates start at 0,0 for row 0 of the tree, and cover the entire
scrollable area of the tree.
*/
func (self *_TraitTreeView) GetVisibleRect() (visible_rect C.GdkRectangle) {
	C.gtk_tree_view_get_visible_rect(self.CPointer, &visible_rect)
	return
}

/*
This inserts the @column into the @tree_view at @position.  If @position is
-1, then the column is inserted at the end. If @tree_view has
“fixed_height” mode enabled, then @column must have its “sizing” property
set to be GTK_TREE_VIEW_COLUMN_FIXED.
*/
func (self *_TraitTreeView) InsertColumn(column *TreeViewColumn, position int) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_insert_column(self.CPointer, (*C.GtkTreeViewColumn)(column.CPointer), C.gint(position))
	return__ = int(__cgo__return__)
	return
}

// gtk_tree_view_insert_column_with_attributes is not generated due to varargs

/*
Convenience function that inserts a new column into the #GtkTreeView
with the given cell renderer and a #GtkTreeCellDataFunc to set cell renderer
attributes (normally using data from the model). See also
gtk_tree_view_column_set_cell_data_func(), gtk_tree_view_column_pack_start().
If @tree_view has “fixed_height” mode enabled, then the new column will have its
“sizing” property set to be GTK_TREE_VIEW_COLUMN_FIXED.
*/
func (self *_TraitTreeView) InsertColumnWithDataFunc(position int, title string, cell *CellRenderer, func_ C.GtkTreeCellDataFunc, data unsafe.Pointer, dnotify C.GDestroyNotify) (return__ int) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_insert_column_with_data_func(self.CPointer, C.gint(position), __cgo__title, (*C.GtkCellRenderer)(cell.CPointer), func_, (C.gpointer)(data), dnotify)
	C.free(unsafe.Pointer(__cgo__title))
	return__ = int(__cgo__return__)
	return
}

/*
Determine whether the point (@x, @y) in @tree_view is blank, that is no
cell content nor an expander arrow is drawn at the location. If so, the
location can be considered as the background. You might wish to take
special action on clicks on the background, such as clearing a current
selection, having a custom context menu or starting rubber banding.

The @x and @y coordinate that are provided must be relative to bin_window
coordinates.  That is, @x and @y must come from an event on @tree_view
where `event->window == gtk_tree_view_get_bin_window ()`.

For converting widget coordinates (eg. the ones you get from
GtkWidget::query-tooltip), please see
gtk_tree_view_convert_widget_to_bin_window_coords().

The @path, @column, @cell_x and @cell_y arguments will be filled in
likewise as for gtk_tree_view_get_path_at_pos().  Please see
gtk_tree_view_get_path_at_pos() for more information.
*/
func (self *_TraitTreeView) IsBlankAtPos(x int, y int) (path *TreePath, column *TreeViewColumn, cell_x int, cell_y int, return__ bool) {
	var __cgo__path *C.GtkTreePath
	var __cgo__column *C.GtkTreeViewColumn
	var __cgo__cell_x C.gint
	var __cgo__cell_y C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_is_blank_at_pos(self.CPointer, C.gint(x), C.gint(y), &__cgo__path, &__cgo__column, &__cgo__cell_x, &__cgo__cell_y)
	path = (*TreePath)(unsafe.Pointer(__cgo__path))
	column = NewTreeViewColumnFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__column).Pointer()))
	cell_x = int(__cgo__cell_x)
	cell_y = int(__cgo__cell_y)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether a rubber banding operation is currently being done
in @tree_view.
*/
func (self *_TraitTreeView) IsRubberBandingActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_is_rubber_banding_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Calls @func on all expanded rows.
*/
func (self *_TraitTreeView) MapExpandedRows(func_ C.GtkTreeViewMappingFunc, data unsafe.Pointer) {
	C.gtk_tree_view_map_expanded_rows(self.CPointer, func_, (C.gpointer)(data))
	return
}

/*
Moves @column to be after to @base_column.  If @base_column is %NULL, then
@column is placed in the first position.
*/
func (self *_TraitTreeView) MoveColumnAfter(column *TreeViewColumn, base_column *TreeViewColumn) {
	C.gtk_tree_view_move_column_after(self.CPointer, (*C.GtkTreeViewColumn)(column.CPointer), (*C.GtkTreeViewColumn)(base_column.CPointer))
	return
}

/*
Removes @column from @tree_view.
*/
func (self *_TraitTreeView) RemoveColumn(column *TreeViewColumn) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_remove_column(self.CPointer, (*C.GtkTreeViewColumn)(column.CPointer))
	return__ = int(__cgo__return__)
	return
}

/*
Activates the cell determined by @path and @column.
*/
func (self *_TraitTreeView) RowActivated(path *TreePath, column *TreeViewColumn) {
	C.gtk_tree_view_row_activated(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkTreeViewColumn)(column.CPointer))
	return
}

/*
Returns %TRUE if the node pointed to by @path is expanded in @tree_view.
*/
func (self *_TraitTreeView) RowExpanded(path *TreePath) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_row_expanded(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Moves the alignments of @tree_view to the position specified by @column and
@path.  If @column is %NULL, then no horizontal scrolling occurs.  Likewise,
if @path is %NULL no vertical scrolling occurs.  At a minimum, one of @column
or @path need to be non-%NULL.  @row_align determines where the row is
placed, and @col_align determines where @column is placed.  Both are expected
to be between 0.0 and 1.0. 0.0 means left/top alignment, 1.0 means
right/bottom alignment, 0.5 means center.

If @use_align is %FALSE, then the alignment arguments are ignored, and the
tree does the minimum amount of work to scroll the cell onto the screen.
This means that the cell will be scrolled to the edge closest to its current
position.  If the cell is currently visible on the screen, nothing is done.

This function only works if the model is set, and @path is a valid row on the
model.  If the model changes before the @tree_view is realized, the centered
path will be modified to reflect this change.
*/
func (self *_TraitTreeView) ScrollToCell(path *TreePath, column *TreeViewColumn, use_align bool, row_align float32, col_align float32) {
	__cgo__use_align := C.gboolean(0)
	if use_align {
		__cgo__use_align = C.gboolean(1)
	}
	C.gtk_tree_view_scroll_to_cell(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkTreeViewColumn)(column.CPointer), __cgo__use_align, C.gfloat(row_align), C.gfloat(col_align))
	return
}

/*
Scrolls the tree view such that the top-left corner of the visible
area is @tree_x, @tree_y, where @tree_x and @tree_y are specified
in tree coordinates.  The @tree_view must be realized before
this function is called.  If it isn't, you probably want to be
using gtk_tree_view_scroll_to_cell().

If either @tree_x or @tree_y are -1, then that direction isn’t scrolled.
*/
func (self *_TraitTreeView) ScrollToPoint(tree_x int, tree_y int) {
	C.gtk_tree_view_scroll_to_point(self.CPointer, C.gint(tree_x), C.gint(tree_y))
	return
}

/*
Cause the #GtkTreeView::row-activated signal to be emitted
on a single click instead of a double click.
*/
func (self *_TraitTreeView) SetActivateOnSingleClick(single bool) {
	__cgo__single := C.gboolean(0)
	if single {
		__cgo__single = C.gboolean(1)
	}
	C.gtk_tree_view_set_activate_on_single_click(self.CPointer, __cgo__single)
	return
}

/*
Sets a user function for determining where a column may be dropped when
dragged.  This function is called on every column pair in turn at the
beginning of a column drag to determine where a drop can take place.  The
arguments passed to @func are: the @tree_view, the #GtkTreeViewColumn being
dragged, the two #GtkTreeViewColumn s determining the drop spot, and
@user_data.  If either of the #GtkTreeViewColumn arguments for the drop spot
are %NULL, then they indicate an edge.  If @func is set to be %NULL, then
@tree_view reverts to the default behavior of allowing all columns to be
dropped everywhere.
*/
func (self *_TraitTreeView) SetColumnDragFunction(func_ C.GtkTreeViewColumnDropFunc, user_data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_tree_view_set_column_drag_function(self.CPointer, func_, (C.gpointer)(user_data), destroy)
	return
}

/*
Sets the current keyboard focus to be at @path, and selects it.  This is
useful when you want to focus the user’s attention on a particular row.  If
@focus_column is not %NULL, then focus is given to the column specified by
it. Additionally, if @focus_column is specified, and @start_editing is
%TRUE, then editing should be started in the specified cell.
This function is often followed by @gtk_widget_grab_focus (@tree_view)
in order to give keyboard focus to the widget.  Please note that editing
can only happen when the widget is realized.

If @path is invalid for @model, the current cursor (if any) will be unset
and the function will return without failing.
*/
func (self *_TraitTreeView) SetCursor(path *TreePath, focus_column *TreeViewColumn, start_editing bool) {
	__cgo__start_editing := C.gboolean(0)
	if start_editing {
		__cgo__start_editing = C.gboolean(1)
	}
	C.gtk_tree_view_set_cursor(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkTreeViewColumn)(focus_column.CPointer), __cgo__start_editing)
	return
}

/*
Sets the current keyboard focus to be at @path, and selects it.  This is
useful when you want to focus the user’s attention on a particular row.  If
@focus_column is not %NULL, then focus is given to the column specified by
it. If @focus_column and @focus_cell are not %NULL, and @focus_column
contains 2 or more editable or activatable cells, then focus is given to
the cell specified by @focus_cell. Additionally, if @focus_column is
specified, and @start_editing is %TRUE, then editing should be started in
the specified cell.  This function is often followed by
@gtk_widget_grab_focus (@tree_view) in order to give keyboard focus to the
widget.  Please note that editing can only happen when the widget is
realized.

If @path is invalid for @model, the current cursor (if any) will be unset
and the function will return without failing.
*/
func (self *_TraitTreeView) SetCursorOnCell(path *TreePath, focus_column *TreeViewColumn, focus_cell *CellRenderer, start_editing bool) {
	__cgo__start_editing := C.gboolean(0)
	if start_editing {
		__cgo__start_editing = C.gboolean(1)
	}
	C.gtk_tree_view_set_cursor_on_cell(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkTreeViewColumn)(focus_column.CPointer), (*C.GtkCellRenderer)(focus_cell.CPointer), __cgo__start_editing)
	return
}

// gtk_tree_view_set_destroy_count_func is not generated due to deprecation attr

/*
Sets the row that is highlighted for feedback.
If @path is %NULL, an existing highlight is removed.
*/
func (self *_TraitTreeView) SetDragDestRow(path *TreePath, pos C.GtkTreeViewDropPosition) {
	C.gtk_tree_view_set_drag_dest_row(self.CPointer, (*C.GtkTreePath)(unsafe.Pointer(path)), pos)
	return
}

/*
If @enable_search is set, then the user can type in text to search through
the tree interactively (this is sometimes called "typeahead find").

Note that even if this is %FALSE, the user can still initiate a search
using the “start-interactive-search” key binding.
*/
func (self *_TraitTreeView) SetEnableSearch(enable_search bool) {
	__cgo__enable_search := C.gboolean(0)
	if enable_search {
		__cgo__enable_search = C.gboolean(1)
	}
	C.gtk_tree_view_set_enable_search(self.CPointer, __cgo__enable_search)
	return
}

/*
Sets whether to draw lines interconnecting the expanders in @tree_view.
This does not have any visible effects for lists.
*/
func (self *_TraitTreeView) SetEnableTreeLines(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.gtk_tree_view_set_enable_tree_lines(self.CPointer, __cgo__enabled)
	return
}

/*
Sets the column to draw the expander arrow at. It must be in @tree_view.
If @column is %NULL, then the expander arrow is always at the first
visible column.

If you do not want expander arrow to appear in your tree, set the
expander column to a hidden column.
*/
func (self *_TraitTreeView) SetExpanderColumn(column *TreeViewColumn) {
	C.gtk_tree_view_set_expander_column(self.CPointer, (*C.GtkTreeViewColumn)(column.CPointer))
	return
}

/*
Enables or disables the fixed height mode of @tree_view.
Fixed height mode speeds up #GtkTreeView by assuming that all
rows have the same height.
Only enable this option if all rows are the same height and all
columns are of type %GTK_TREE_VIEW_COLUMN_FIXED.
*/
func (self *_TraitTreeView) SetFixedHeightMode(enable bool) {
	__cgo__enable := C.gboolean(0)
	if enable {
		__cgo__enable = C.gboolean(1)
	}
	C.gtk_tree_view_set_fixed_height_mode(self.CPointer, __cgo__enable)
	return
}

/*
Sets which grid lines to draw in @tree_view.
*/
func (self *_TraitTreeView) SetGridLines(grid_lines C.GtkTreeViewGridLines) {
	C.gtk_tree_view_set_grid_lines(self.CPointer, grid_lines)
	return
}

// gtk_tree_view_set_hadjustment is not generated due to deprecation attr

/*
Allow the column title buttons to be clicked.
*/
func (self *_TraitTreeView) SetHeadersClickable(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_tree_view_set_headers_clickable(self.CPointer, __cgo__setting)
	return
}

/*
Sets the visibility state of the headers.
*/
func (self *_TraitTreeView) SetHeadersVisible(headers_visible bool) {
	__cgo__headers_visible := C.gboolean(0)
	if headers_visible {
		__cgo__headers_visible = C.gboolean(1)
	}
	C.gtk_tree_view_set_headers_visible(self.CPointer, __cgo__headers_visible)
	return
}

/*
Enables or disables the hover expansion mode of @tree_view.
Hover expansion makes rows expand or collapse if the pointer
moves over them.
*/
func (self *_TraitTreeView) SetHoverExpand(expand bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	C.gtk_tree_view_set_hover_expand(self.CPointer, __cgo__expand)
	return
}

/*
Enables or disables the hover selection mode of @tree_view.
Hover selection makes the selected row follow the pointer.
Currently, this works only for the selection modes
%GTK_SELECTION_SINGLE and %GTK_SELECTION_BROWSE.
*/
func (self *_TraitTreeView) SetHoverSelection(hover bool) {
	__cgo__hover := C.gboolean(0)
	if hover {
		__cgo__hover = C.gboolean(1)
	}
	C.gtk_tree_view_set_hover_selection(self.CPointer, __cgo__hover)
	return
}

/*
Sets the amount of extra indentation for child levels to use in @tree_view
in addition to the default indentation.  The value should be specified in
pixels, a value of 0 disables this feature and in this case only the default
indentation will be used.
This does not have any visible effects for lists.
*/
func (self *_TraitTreeView) SetLevelIndentation(indentation int) {
	C.gtk_tree_view_set_level_indentation(self.CPointer, C.gint(indentation))
	return
}

/*
Sets the model for a #GtkTreeView.  If the @tree_view already has a model
set, it will remove it before setting the new model.  If @model is %NULL,
then it will unset the old model.
*/
func (self *_TraitTreeView) SetModel(model *C.GtkTreeModel) {
	C.gtk_tree_view_set_model(self.CPointer, model)
	return
}

/*
This function is a convenience function to allow you to reorder
models that support the #GtkTreeDragSourceIface and the
#GtkTreeDragDestIface.  Both #GtkTreeStore and #GtkListStore support
these.  If @reorderable is %TRUE, then the user can reorder the
model by dragging and dropping rows. The developer can listen to
these changes by connecting to the model’s #GtkTreeModel::row-inserted
and #GtkTreeModel::row-deleted signals. The reordering is implemented
by setting up the tree view as a drag source and destination.
Therefore, drag and drop can not be used in a reorderable view for any
other purpose.

This function does not give you any degree of control over the order -- any
reordering is allowed.  If more control is needed, you should probably
handle drag and drop manually.
*/
func (self *_TraitTreeView) SetReorderable(reorderable bool) {
	__cgo__reorderable := C.gboolean(0)
	if reorderable {
		__cgo__reorderable = C.gboolean(1)
	}
	C.gtk_tree_view_set_reorderable(self.CPointer, __cgo__reorderable)
	return
}

/*
Sets the row separator function, which is used to determine
whether a row should be drawn as a separator. If the row separator
function is %NULL, no separators are drawn. This is the default value.
*/
func (self *_TraitTreeView) SetRowSeparatorFunc(func_ C.GtkTreeViewRowSeparatorFunc, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_tree_view_set_row_separator_func(self.CPointer, func_, (C.gpointer)(data), destroy)
	return
}

/*
Enables or disables rubber banding in @tree_view.  If the selection mode
is #GTK_SELECTION_MULTIPLE, rubber banding will allow the user to select
multiple rows by dragging the mouse.
*/
func (self *_TraitTreeView) SetRubberBanding(enable bool) {
	__cgo__enable := C.gboolean(0)
	if enable {
		__cgo__enable = C.gboolean(1)
	}
	C.gtk_tree_view_set_rubber_banding(self.CPointer, __cgo__enable)
	return
}

/*
This function tells GTK+ that the user interface for your
application requires users to read across tree rows and associate
cells with one another. By default, GTK+ will then render the tree
with alternating row colors. Do not use it
just because you prefer the appearance of the ruled tree; that’s a
question for the theme. Some themes will draw tree rows in
alternating colors even when rules are turned off, and users who
prefer that appearance all the time can choose those themes. You
should call this function only as a semantic
hint to the theme engine that your tree makes alternating colors
useful from a functional standpoint (since it has lots of columns,
generally).
*/
func (self *_TraitTreeView) SetRulesHint(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_tree_view_set_rules_hint(self.CPointer, __cgo__setting)
	return
}

/*
Sets @column as the column where the interactive search code should
search in for the current model.

If the search column is set, users can use the “start-interactive-search”
key binding to bring up search popup. The enable-search property controls
whether simply typing text will also start an interactive search.

Note that @column refers to a column of the current model. The search
column is reset to -1 when the model is changed.
*/
func (self *_TraitTreeView) SetSearchColumn(column int) {
	C.gtk_tree_view_set_search_column(self.CPointer, C.gint(column))
	return
}

/*
Sets the entry which the interactive search code will use for this
@tree_view.  This is useful when you want to provide a search entry
in our interface at all time at a fixed position.  Passing %NULL for
@entry will make the interactive search code use the built-in popup
entry again.
*/
func (self *_TraitTreeView) SetSearchEntry(entry *Entry) {
	C.gtk_tree_view_set_search_entry(self.CPointer, (*C.GtkEntry)(entry.CPointer))
	return
}

/*
Sets the compare function for the interactive search capabilities; note
that somewhat like strcmp() returning 0 for equality
#GtkTreeViewSearchEqualFunc returns %FALSE on matches.
*/
func (self *_TraitTreeView) SetSearchEqualFunc(search_equal_func C.GtkTreeViewSearchEqualFunc, search_user_data unsafe.Pointer, search_destroy C.GDestroyNotify) {
	C.gtk_tree_view_set_search_equal_func(self.CPointer, search_equal_func, (C.gpointer)(search_user_data), search_destroy)
	return
}

/*
Sets the function to use when positioning the search dialog.
*/
func (self *_TraitTreeView) SetSearchPositionFunc(func_ C.GtkTreeViewSearchPositionFunc, data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_tree_view_set_search_position_func(self.CPointer, func_, (C.gpointer)(data), destroy)
	return
}

/*
Sets whether to draw and enable expanders and indent child rows in
@tree_view.  When disabled there will be no expanders visible in trees
and there will be no way to expand and collapse rows by default.  Also
note that hiding the expanders will disable the default indentation.  You
can set a custom indentation in this case using
gtk_tree_view_set_level_indentation().
This does not have any visible effects for lists.
*/
func (self *_TraitTreeView) SetShowExpanders(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.gtk_tree_view_set_show_expanders(self.CPointer, __cgo__enabled)
	return
}

/*
Sets the tip area of @tooltip to the area @path, @column and @cell have
in common.  For example if @path is %NULL and @column is set, the tip
area will be set to the full area covered by @column.  See also
gtk_tooltip_set_tip_area().

Note that if @path is not specified and @cell is set and part of a column
containing the expander, the tooltip might not show and hide at the correct
position.  In such cases @path must be set to the current node under the
mouse cursor for this function to operate correctly.

See also gtk_tree_view_set_tooltip_column() for a simpler alternative.
*/
func (self *_TraitTreeView) SetTooltipCell(tooltip *Tooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer) {
	C.gtk_tree_view_set_tooltip_cell(self.CPointer, (*C.GtkTooltip)(tooltip.CPointer), (*C.GtkTreePath)(unsafe.Pointer(path)), (*C.GtkTreeViewColumn)(column.CPointer), (*C.GtkCellRenderer)(cell.CPointer))
	return
}

/*
If you only plan to have simple (text-only) tooltips on full rows, you
can use this function to have #GtkTreeView handle these automatically
for you. @column should be set to the column in @tree_view’s model
containing the tooltip texts, or -1 to disable this feature.

When enabled, #GtkWidget:has-tooltip will be set to %TRUE and
@tree_view will connect a #GtkWidget::query-tooltip signal handler.

Note that the signal handler sets the text with gtk_tooltip_set_markup(),
so &, <, etc have to be escaped in the text.
*/
func (self *_TraitTreeView) SetTooltipColumn(column int) {
	C.gtk_tree_view_set_tooltip_column(self.CPointer, C.gint(column))
	return
}

/*
Sets the tip area of @tooltip to be the area covered by the row at @path.
See also gtk_tree_view_set_tooltip_column() for a simpler alternative.
See also gtk_tooltip_set_tip_area().
*/
func (self *_TraitTreeView) SetTooltipRow(tooltip *Tooltip, path *TreePath) {
	C.gtk_tree_view_set_tooltip_row(self.CPointer, (*C.GtkTooltip)(tooltip.CPointer), (*C.GtkTreePath)(unsafe.Pointer(path)))
	return
}

// gtk_tree_view_set_vadjustment is not generated due to deprecation attr

/*
Undoes the effect of
gtk_tree_view_enable_model_drag_dest(). Calling this method sets
#GtkTreeView:reorderable to %FALSE.
*/
func (self *_TraitTreeView) UnsetRowsDragDest() {
	C.gtk_tree_view_unset_rows_drag_dest(self.CPointer)
	return
}

/*
Undoes the effect of
gtk_tree_view_enable_model_drag_source(). Calling this method sets
#GtkTreeView:reorderable to %FALSE.
*/
func (self *_TraitTreeView) UnsetRowsDragSource() {
	C.gtk_tree_view_unset_rows_drag_source(self.CPointer)
	return
}

type _TraitTreeViewColumn struct{ CPointer *C.GtkTreeViewColumn }

/*
Adds an attribute mapping to the list in @tree_column.  The @column is the
column of the model to get a value from, and the @attribute is the
parameter on @cell_renderer to be set from the value. So for example
if column 2 of the model contains strings, you could have the
“text” attribute of a #GtkCellRendererText get its values from
column 2.
*/
func (self *_TraitTreeViewColumn) AddAttribute(cell_renderer *CellRenderer, attribute string, column int) {
	__cgo__attribute := (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	C.gtk_tree_view_column_add_attribute(self.CPointer, (*C.GtkCellRenderer)(cell_renderer.CPointer), __cgo__attribute, C.gint(column))
	C.free(unsafe.Pointer(__cgo__attribute))
	return
}

/*
Obtains the horizontal position and size of a cell in a column. If the
cell is not found in the column, @start_pos and @width are not changed and
%FALSE is returned.
*/
func (self *_TraitTreeViewColumn) CellGetPosition(cell_renderer *CellRenderer) (x_offset int, width int, return__ bool) {
	var __cgo__x_offset C.gint
	var __cgo__width C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_column_cell_get_position(self.CPointer, (*C.GtkCellRenderer)(cell_renderer.CPointer), &__cgo__x_offset, &__cgo__width)
	x_offset = int(__cgo__x_offset)
	width = int(__cgo__width)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Obtains the width and height needed to render the column.  This is used
primarily by the #GtkTreeView.
*/
func (self *_TraitTreeViewColumn) CellGetSize(cell_area *C.GdkRectangle) (x_offset int, y_offset int, width int, height int) {
	var __cgo__x_offset C.gint
	var __cgo__y_offset C.gint
	var __cgo__width C.gint
	var __cgo__height C.gint
	C.gtk_tree_view_column_cell_get_size(self.CPointer, cell_area, &__cgo__x_offset, &__cgo__y_offset, &__cgo__width, &__cgo__height)
	x_offset = int(__cgo__x_offset)
	y_offset = int(__cgo__y_offset)
	width = int(__cgo__width)
	height = int(__cgo__height)
	return
}

/*
Returns %TRUE if any of the cells packed into the @tree_column are visible.
For this to be meaningful, you must first initialize the cells with
gtk_tree_view_column_cell_set_cell_data()
*/
func (self *_TraitTreeViewColumn) CellIsVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_column_cell_is_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets the cell renderer based on the @tree_model and @iter.  That is, for
every attribute mapping in @tree_column, it will get a value from the set
column on the @iter, and use that value to set the attribute on the cell
renderer.  This is used primarily by the #GtkTreeView.
*/
func (self *_TraitTreeViewColumn) CellSetCellData(tree_model *C.GtkTreeModel, iter *TreeIter, is_expander bool, is_expanded bool) {
	__cgo__is_expander := C.gboolean(0)
	if is_expander {
		__cgo__is_expander = C.gboolean(1)
	}
	__cgo__is_expanded := C.gboolean(0)
	if is_expanded {
		__cgo__is_expanded = C.gboolean(1)
	}
	C.gtk_tree_view_column_cell_set_cell_data(self.CPointer, tree_model, (*C.GtkTreeIter)(unsafe.Pointer(iter)), __cgo__is_expander, __cgo__is_expanded)
	return
}

/*
Unsets all the mappings on all renderers on the @tree_column.
*/
func (self *_TraitTreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(self.CPointer)
	return
}

/*
Clears all existing attributes previously set with
gtk_tree_view_column_set_attributes().
*/
func (self *_TraitTreeViewColumn) ClearAttributes(cell_renderer *CellRenderer) {
	C.gtk_tree_view_column_clear_attributes(self.CPointer, (*C.GtkCellRenderer)(cell_renderer.CPointer))
	return
}

/*
Emits the “clicked” signal on the column.  This function will only work if
@tree_column is clickable.
*/
func (self *_TraitTreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(self.CPointer)
	return
}

/*
Sets the current keyboard focus to be at @cell, if the column contains
2 or more editable and activatable cells.
*/
func (self *_TraitTreeViewColumn) FocusCell(cell *CellRenderer) {
	C.gtk_tree_view_column_focus_cell(self.CPointer, (*C.GtkCellRenderer)(cell.CPointer))
	return
}

/*
Returns the current x alignment of @tree_column.  This value can range
between 0.0 and 1.0.
*/
func (self *_TraitTreeViewColumn) GetAlignment() (return__ float32) {
	var __cgo__return__ C.gfloat
	__cgo__return__ = C.gtk_tree_view_column_get_alignment(self.CPointer)
	return__ = float32(__cgo__return__)
	return
}

/*
Returns the button used in the treeview column header
*/
func (self *_TraitTreeViewColumn) GetButton() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tree_view_column_get_button(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns %TRUE if the user can click on the header for the column.
*/
func (self *_TraitTreeViewColumn) GetClickable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_column_get_clickable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the column expands to fill available space.
*/
func (self *_TraitTreeViewColumn) GetExpand() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_column_get_expand(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the fixed width of the column.  This may not be the actual displayed
width of the column; for that, use gtk_tree_view_column_get_width().
*/
func (self *_TraitTreeViewColumn) GetFixedWidth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_column_get_fixed_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the maximum width in pixels of the @tree_column, or -1 if no maximum
width is set.
*/
func (self *_TraitTreeViewColumn) GetMaxWidth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_column_get_max_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the minimum width in pixels of the @tree_column, or -1 if no minimum
width is set.
*/
func (self *_TraitTreeViewColumn) GetMinWidth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_column_get_min_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns %TRUE if the @tree_column can be reordered by the user.
*/
func (self *_TraitTreeViewColumn) GetReorderable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_column_get_reorderable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns %TRUE if the @tree_column can be resized by the end user.
*/
func (self *_TraitTreeViewColumn) GetResizable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_column_get_resizable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current type of @tree_column.
*/
func (self *_TraitTreeViewColumn) GetSizing() (return__ C.GtkTreeViewColumnSizing) {
	return__ = C.gtk_tree_view_column_get_sizing(self.CPointer)
	return
}

/*
Gets the logical @sort_column_id that the model sorts on when this
column is selected for sorting.
See gtk_tree_view_column_set_sort_column_id().
*/
func (self *_TraitTreeViewColumn) GetSortColumnId() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_column_get_sort_column_id(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value set by gtk_tree_view_column_set_sort_indicator().
*/
func (self *_TraitTreeViewColumn) GetSortIndicator() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_column_get_sort_indicator(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_tree_view_column_set_sort_order().
*/
func (self *_TraitTreeViewColumn) GetSortOrder() (return__ C.GtkSortType) {
	return__ = C.gtk_tree_view_column_get_sort_order(self.CPointer)
	return
}

/*
Returns the spacing of @tree_column.
*/
func (self *_TraitTreeViewColumn) GetSpacing() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_column_get_spacing(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the title of the widget.
*/
func (self *_TraitTreeViewColumn) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_tree_view_column_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the #GtkTreeView wherein @tree_column has been inserted.
If @column is currently not inserted in any tree view, %NULL is
returned.
*/
func (self *_TraitTreeViewColumn) GetTreeView() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tree_view_column_get_tree_view(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns %TRUE if @tree_column is visible.
*/
func (self *_TraitTreeViewColumn) GetVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_tree_view_column_get_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the #GtkWidget in the button on the column header.
If a custom widget has not been set then %NULL is returned.
*/
func (self *_TraitTreeViewColumn) GetWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_tree_view_column_get_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns the current size of @tree_column in pixels.
*/
func (self *_TraitTreeViewColumn) GetWidth() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_column_get_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the current X offset of @tree_column in pixels.
*/
func (self *_TraitTreeViewColumn) GetXOffset() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_tree_view_column_get_x_offset(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Adds the @cell to end of the column. If @expand is %FALSE, then the @cell
is allocated no more space than it needs. Any unused space is divided
evenly between cells for which @expand is %TRUE.
*/
func (self *_TraitTreeViewColumn) PackEnd(cell *CellRenderer, expand bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	C.gtk_tree_view_column_pack_end(self.CPointer, (*C.GtkCellRenderer)(cell.CPointer), __cgo__expand)
	return
}

/*
Packs the @cell into the beginning of the column. If @expand is %FALSE, then
the @cell is allocated no more space than it needs. Any unused space is divided
evenly between cells for which @expand is %TRUE.
*/
func (self *_TraitTreeViewColumn) PackStart(cell *CellRenderer, expand bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	C.gtk_tree_view_column_pack_start(self.CPointer, (*C.GtkCellRenderer)(cell.CPointer), __cgo__expand)
	return
}

/*
Flags the column, and the cell renderers added to this column, to have
their sizes renegotiated.
*/
func (self *_TraitTreeViewColumn) QueueResize() {
	C.gtk_tree_view_column_queue_resize(self.CPointer)
	return
}

/*
Sets the alignment of the title or custom widget inside the column header.
The alignment determines its location inside the button -- 0.0 for left, 0.5
for center, 1.0 for right.
*/
func (self *_TraitTreeViewColumn) SetAlignment(xalign float32) {
	C.gtk_tree_view_column_set_alignment(self.CPointer, C.gfloat(xalign))
	return
}

// gtk_tree_view_column_set_attributes is not generated due to varargs

/*
Sets the #GtkTreeCellDataFunc to use for the column.  This
function is used instead of the standard attributes mapping for
setting the column value, and should set the value of @tree_column's
cell renderer as appropriate.  @func may be %NULL to remove an
older one.
*/
func (self *_TraitTreeViewColumn) SetCellDataFunc(cell_renderer *CellRenderer, func_ C.GtkTreeCellDataFunc, func_data unsafe.Pointer, destroy C.GDestroyNotify) {
	C.gtk_tree_view_column_set_cell_data_func(self.CPointer, (*C.GtkCellRenderer)(cell_renderer.CPointer), func_, (C.gpointer)(func_data), destroy)
	return
}

/*
Sets the header to be active if @clickable is %TRUE.  When the header is
active, then it can take keyboard focus, and can be clicked.
*/
func (self *_TraitTreeViewColumn) SetClickable(clickable bool) {
	__cgo__clickable := C.gboolean(0)
	if clickable {
		__cgo__clickable = C.gboolean(1)
	}
	C.gtk_tree_view_column_set_clickable(self.CPointer, __cgo__clickable)
	return
}

/*
Sets the column to take available extra space.  This space is shared equally
amongst all columns that have the expand set to %TRUE.  If no column has this
option set, then the last column gets all extra space.  By default, every
column is created with this %FALSE.

Along with “fixed-width”, the “expand” property changes when the column is
resized by the user.
*/
func (self *_TraitTreeViewColumn) SetExpand(expand bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	C.gtk_tree_view_column_set_expand(self.CPointer, __cgo__expand)
	return
}

/*
If @fixed_width is not -1, sets the fixed width of @tree_column; otherwise
unsets it.  The effective value of @fixed_width is clamped between the
minumum and maximum width of the column; however, the value stored in the
“fixed-width” property is not clamped.  If the column sizing is
#GTK_TREE_VIEW_COLUMN_GROW_ONLY or #GTK_TREE_VIEW_COLUMN_AUTOSIZE, setting a
fixed width overrides the automatically calculated width.  Note that
@fixed_width is only a hint to GTK+; the width actually allocated to the
column may be greater or less than requested.

Along with “expand”, the “fixed-width” property changes when the column is
resized by the user.
*/
func (self *_TraitTreeViewColumn) SetFixedWidth(fixed_width int) {
	C.gtk_tree_view_column_set_fixed_width(self.CPointer, C.gint(fixed_width))
	return
}

/*
Sets the maximum width of the @tree_column.  If @max_width is -1, then the
maximum width is unset.  Note, the column can actually be wider than max
width if it’s the last column in a view.  In this case, the column expands to
fill any extra space.
*/
func (self *_TraitTreeViewColumn) SetMaxWidth(max_width int) {
	C.gtk_tree_view_column_set_max_width(self.CPointer, C.gint(max_width))
	return
}

/*
Sets the minimum width of the @tree_column.  If @min_width is -1, then the
minimum width is unset.
*/
func (self *_TraitTreeViewColumn) SetMinWidth(min_width int) {
	C.gtk_tree_view_column_set_min_width(self.CPointer, C.gint(min_width))
	return
}

/*
If @reorderable is %TRUE, then the column can be reordered by the end user
dragging the header.
*/
func (self *_TraitTreeViewColumn) SetReorderable(reorderable bool) {
	__cgo__reorderable := C.gboolean(0)
	if reorderable {
		__cgo__reorderable = C.gboolean(1)
	}
	C.gtk_tree_view_column_set_reorderable(self.CPointer, __cgo__reorderable)
	return
}

/*
If @resizable is %TRUE, then the user can explicitly resize the column by
grabbing the outer edge of the column button.  If resizable is %TRUE and
sizing mode of the column is #GTK_TREE_VIEW_COLUMN_AUTOSIZE, then the sizing
mode is changed to #GTK_TREE_VIEW_COLUMN_GROW_ONLY.
*/
func (self *_TraitTreeViewColumn) SetResizable(resizable bool) {
	__cgo__resizable := C.gboolean(0)
	if resizable {
		__cgo__resizable = C.gboolean(1)
	}
	C.gtk_tree_view_column_set_resizable(self.CPointer, __cgo__resizable)
	return
}

/*
Sets the growth behavior of @tree_column to @type.
*/
func (self *_TraitTreeViewColumn) SetSizing(type_ C.GtkTreeViewColumnSizing) {
	C.gtk_tree_view_column_set_sizing(self.CPointer, type_)
	return
}

/*
Sets the logical @sort_column_id that this column sorts on when this column
is selected for sorting.  Doing so makes the column header clickable.
*/
func (self *_TraitTreeViewColumn) SetSortColumnId(sort_column_id int) {
	C.gtk_tree_view_column_set_sort_column_id(self.CPointer, C.gint(sort_column_id))
	return
}

/*
Call this function with a @setting of %TRUE to display an arrow in
the header button indicating the column is sorted. Call
gtk_tree_view_column_set_sort_order() to change the direction of
the arrow.
*/
func (self *_TraitTreeViewColumn) SetSortIndicator(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_tree_view_column_set_sort_indicator(self.CPointer, __cgo__setting)
	return
}

/*
Changes the appearance of the sort indicator.

This does not actually sort the model.  Use
gtk_tree_view_column_set_sort_column_id() if you want automatic sorting
support.  This function is primarily for custom sorting behavior, and should
be used in conjunction with gtk_tree_sortable_set_sort_column_id() to do
that. For custom models, the mechanism will vary.

The sort indicator changes direction to indicate normal sort or reverse sort.
Note that you must have the sort indicator enabled to see anything when
calling this function; see gtk_tree_view_column_set_sort_indicator().
*/
func (self *_TraitTreeViewColumn) SetSortOrder(order C.GtkSortType) {
	C.gtk_tree_view_column_set_sort_order(self.CPointer, order)
	return
}

/*
Sets the spacing field of @tree_column, which is the number of pixels to
place between cell renderers packed into it.
*/
func (self *_TraitTreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(self.CPointer, C.gint(spacing))
	return
}

/*
Sets the title of the @tree_column.  If a custom widget has been set, then
this value is ignored.
*/
func (self *_TraitTreeViewColumn) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_tree_view_column_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
Sets the visibility of @tree_column.
*/
func (self *_TraitTreeViewColumn) SetVisible(visible bool) {
	__cgo__visible := C.gboolean(0)
	if visible {
		__cgo__visible = C.gboolean(1)
	}
	C.gtk_tree_view_column_set_visible(self.CPointer, __cgo__visible)
	return
}

/*
Sets the widget in the header to be @widget.  If widget is %NULL, then the
header button is set with a #GtkLabel set to the title of @tree_column.
*/
func (self *_TraitTreeViewColumn) SetWidget(widget *Widget) {
	C.gtk_tree_view_column_set_widget(self.CPointer, (*C.GtkWidget)(widget.CPointer))
	return
}

type _TraitUIManager struct{ CPointer *C.GtkUIManager }

// gtk_ui_manager_add_ui is not generated due to deprecation attr

// gtk_ui_manager_add_ui_from_file is not generated due to deprecation attr

// gtk_ui_manager_add_ui_from_resource is not generated due to deprecation attr

// gtk_ui_manager_add_ui_from_string is not generated due to deprecation attr

// gtk_ui_manager_ensure_update is not generated due to deprecation attr

// gtk_ui_manager_get_accel_group is not generated due to deprecation attr

// gtk_ui_manager_get_action is not generated due to deprecation attr

// gtk_ui_manager_get_action_groups is not generated due to deprecation attr

// gtk_ui_manager_get_add_tearoffs is not generated due to deprecation attr

// gtk_ui_manager_get_toplevels is not generated due to deprecation attr

// gtk_ui_manager_get_ui is not generated due to deprecation attr

// gtk_ui_manager_get_widget is not generated due to deprecation attr

// gtk_ui_manager_insert_action_group is not generated due to deprecation attr

// gtk_ui_manager_new_merge_id is not generated due to deprecation attr

// gtk_ui_manager_remove_action_group is not generated due to deprecation attr

// gtk_ui_manager_remove_ui is not generated due to deprecation attr

// gtk_ui_manager_set_add_tearoffs is not generated due to deprecation attr

type _TraitVBox struct{ CPointer *C.GtkVBox }

type _TraitVButtonBox struct{ CPointer *C.GtkVButtonBox }

type _TraitVPaned struct{ CPointer *C.GtkVPaned }

type _TraitVScale struct{ CPointer *C.GtkVScale }

type _TraitVScrollbar struct{ CPointer *C.GtkVScrollbar }

type _TraitVSeparator struct{ CPointer *C.GtkVSeparator }

type _TraitViewport struct{ CPointer *C.GtkViewport }

/*
Gets the bin window of the #GtkViewport.
*/
func (self *_TraitViewport) GetBinWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_viewport_get_bin_window(self.CPointer)
	return
}

// gtk_viewport_get_hadjustment is not generated due to deprecation attr

/*
Gets the shadow type of the #GtkViewport. See
gtk_viewport_set_shadow_type().
*/
func (self *_TraitViewport) GetShadowType() (return__ C.GtkShadowType) {
	return__ = C.gtk_viewport_get_shadow_type(self.CPointer)
	return
}

// gtk_viewport_get_vadjustment is not generated due to deprecation attr

/*
Gets the view window of the #GtkViewport.
*/
func (self *_TraitViewport) GetViewWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_viewport_get_view_window(self.CPointer)
	return
}

// gtk_viewport_set_hadjustment is not generated due to deprecation attr

/*
Sets the shadow type of the viewport.
*/
func (self *_TraitViewport) SetShadowType(type_ C.GtkShadowType) {
	C.gtk_viewport_set_shadow_type(self.CPointer, type_)
	return
}

// gtk_viewport_set_vadjustment is not generated due to deprecation attr

type _TraitVolumeButton struct{ CPointer *C.GtkVolumeButton }

type _TraitWidget struct{ CPointer *C.GtkWidget }

/*
For widgets that can be “activated” (buttons, menu items, etc.)
this function activates them. Activation is what happens when you
press Enter on a widget during key navigation. If @widget isn't
activatable, the function returns %FALSE.
*/
func (self *_TraitWidget) Activate() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_activate(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Installs an accelerator for this @widget in @accel_group that causes
@accel_signal to be emitted if the accelerator is activated.
The @accel_group needs to be added to the widget’s toplevel via
gtk_window_add_accel_group(), and the signal must be of type %G_SIGNAL_ACTION.
Accelerators added through this function are not user changeable during
runtime. If you want to support accelerators that can be changed by the
user, use gtk_accel_map_add_entry() and gtk_widget_set_accel_path() or
gtk_menu_item_set_accel_path() instead.
*/
func (self *_TraitWidget) AddAccelerator(accel_signal string, accel_group *AccelGroup, accel_key uint, accel_mods C.GdkModifierType, accel_flags C.GtkAccelFlags) {
	__cgo__accel_signal := (*C.gchar)(unsafe.Pointer(C.CString(accel_signal)))
	C.gtk_widget_add_accelerator(self.CPointer, __cgo__accel_signal, (*C.GtkAccelGroup)(accel_group.CPointer), C.guint(accel_key), accel_mods, accel_flags)
	C.free(unsafe.Pointer(__cgo__accel_signal))
	return
}

/*
Adds the device events in the bitfield @events to the event mask for
@widget. See gtk_widget_set_device_events() for details.
*/
func (self *_TraitWidget) AddDeviceEvents(device *C.GdkDevice, events C.GdkEventMask) {
	C.gtk_widget_add_device_events(self.CPointer, device, events)
	return
}

/*
Adds the events in the bitfield @events to the event mask for
@widget. See gtk_widget_set_events() for details.
*/
func (self *_TraitWidget) AddEvents(events int) {
	C.gtk_widget_add_events(self.CPointer, C.gint(events))
	return
}

/*
Adds a widget to the list of mnemonic labels for
this widget. (See gtk_widget_list_mnemonic_labels()). Note the
list of mnemonic labels for the widget is cleared when the
widget is destroyed, so the caller must make sure to update
its internal state at this point as well, by using a connection
to the #GtkWidget::destroy signal or a weak notifier.
*/
func (self *_TraitWidget) AddMnemonicLabel(label *Widget) {
	C.gtk_widget_add_mnemonic_label(self.CPointer, (*C.GtkWidget)(label.CPointer))
	return
}

/*
Queues an animation frame update and adds a callback to be called
before each frame. Until the tick callback is removed, it will be
called frequently (usually at the frame rate of the output device
or as quickly as the application can be repainted, whichever is
slower). For this reason, is most suitable for handling graphics
that change every frame or every few frames. The tick callback does
not automatically imply a relayout or repaint. If you want a
repaint or relayout, and aren’t changing widget properties that
would trigger that (for example, changing the text of a #GtkLabel),
then you will have to call gtk_widget_queue_resize() or
gtk_widget_queue_draw_area() yourself.

gdk_frame_clock_get_frame_time() should generally be used for timing
continuous animations and
gdk_frame_timings_get_predicted_presentation_time() if you are
trying to display isolated frames at particular times.

This is a more convenient alternative to connecting directly to the
#GdkFrameClock::update signal of #GdkFrameClock, since you don't
have to worry about when a #GdkFrameClock is assigned to a widget.
*/
func (self *_TraitWidget) AddTickCallback(callback C.GtkTickCallback, user_data unsafe.Pointer, notify C.GDestroyNotify) (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.gtk_widget_add_tick_callback(self.CPointer, callback, (C.gpointer)(user_data), notify)
	return__ = uint(__cgo__return__)
	return
}

/*
Determines whether an accelerator that activates the signal
identified by @signal_id can currently be activated.
This is done by emitting the #GtkWidget::can-activate-accel
signal on @widget; if the signal isn’t overridden by a
handler or in a derived widget, then the default check is
that the widget must be sensitive, and the widget and all
its ancestors mapped.
*/
func (self *_TraitWidget) CanActivateAccel(signal_id uint) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_can_activate_accel(self.CPointer, C.guint(signal_id))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function is used by custom widget implementations; if you're
writing an app, you’d use gtk_widget_grab_focus() to move the focus
to a particular widget, and gtk_container_set_focus_chain() to
change the focus tab order. So you may want to investigate those
functions instead.

gtk_widget_child_focus() is called by containers as the user moves
around the window using keyboard shortcuts. @direction indicates
what kind of motion is taking place (up, down, left, right, tab
forward, tab backward). gtk_widget_child_focus() emits the
#GtkWidget::focus signal; widgets override the default handler
for this signal in order to implement appropriate focus behavior.

The default ::focus handler for a widget should return %TRUE if
moving in @direction left the focus on a focusable location inside
that widget, and %FALSE if moving in @direction moved the focus
outside the widget. If returning %TRUE, widgets normally
call gtk_widget_grab_focus() to place the focus accordingly;
if returning %FALSE, they don’t modify the current focus location.
*/
func (self *_TraitWidget) ChildFocus(direction C.GtkDirectionType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_child_focus(self.CPointer, direction)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Emits a #GtkWidget::child-notify signal for the
[child property][child-properties] @child_property
on @widget.

This is the analogue of g_object_notify() for child properties.

Also see gtk_container_child_notify().
*/
func (self *_TraitWidget) ChildNotify(child_property string) {
	__cgo__child_property := (*C.gchar)(unsafe.Pointer(C.CString(child_property)))
	C.gtk_widget_child_notify(self.CPointer, __cgo__child_property)
	C.free(unsafe.Pointer(__cgo__child_property))
	return
}

// gtk_widget_class_path is not generated due to deprecation attr

/*
Computes whether a container should give this widget extra space
when possible. Containers should check this, rather than
looking at gtk_widget_get_hexpand() or gtk_widget_get_vexpand().

This function already checks whether the widget is visible, so
visibility does not need to be checked separately. Non-visible
widgets are not expanded.

The computed expand value uses either the expand setting explicitly
set on the widget itself, or, if none has been explicitly set,
the widget may expand if some of its children do.
*/
func (self *_TraitWidget) ComputeExpand(orientation C.GtkOrientation) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_compute_expand(self.CPointer, orientation)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Creates a new #PangoContext with the appropriate font map,
font description, and base direction for drawing text for
this widget. See also gtk_widget_get_pango_context().
*/
func (self *_TraitWidget) CreatePangoContext() (return__ *C.PangoContext) {
	return__ = C.gtk_widget_create_pango_context(self.CPointer)
	return
}

/*
Creates a new #PangoLayout with the appropriate font map,
font description, and base direction for drawing text for
this widget.

If you keep a #PangoLayout created in this way around, you need
to re-create it when the widget #PangoContext is replaced.
This can be tracked by using the #GtkWidget::screen-changed signal
on the widget.
*/
func (self *_TraitWidget) CreatePangoLayout(text string) (return__ *C.PangoLayout) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	return__ = C.gtk_widget_create_pango_layout(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Destroys a widget.

When a widget is
destroyed, it will break any references it holds to other objects.
If the widget is inside a container, the widget will be removed
from the container. If the widget is a toplevel (derived from
#GtkWindow), it will be removed from the list of toplevels, and the
reference GTK+ holds to it will be removed. Removing a
widget from its container or the list of toplevels results in the
widget being finalized, unless you’ve added additional references
to the widget with g_object_ref().

In most cases, only toplevel widgets (windows) require explicit
destruction, because when you destroy a toplevel its children will
be destroyed as well.
*/
func (self *_TraitWidget) Destroy() {
	C.gtk_widget_destroy(self.CPointer)
	return
}

// gtk_widget_destroyed is not generated due to inout param

/*
Returns %TRUE if @device has been shadowed by a GTK+
device grab on another widget, so it would stop sending
events to @widget. This may be used in the
#GtkWidget::grab-notify signal to check for specific
devices. See gtk_device_grab_add().
*/
func (self *_TraitWidget) DeviceIsShadowed(device *C.GdkDevice) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_device_is_shadowed(self.CPointer, device)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_drag_begin is not generated due to deprecation attr

/*
Initiates a drag on the source side. The function only needs to be used
when the application is starting drags itself, and is not needed when
gtk_drag_source_set() is used.

The @event is used to retrieve the timestamp that will be used internally to
grab the pointer.  If @event is %NULL, then %GDK_CURRENT_TIME will be used.
However, you should try to pass a real event in all cases, since that can be
used to get information about the drag.

Generally there are three cases when you want to start a drag by hand by
calling this function:

1. During a #GtkWidget::button-press-event handler, if you want to start a drag
immediately when the user presses the mouse button.  Pass the @event
that you have in your #GtkWidget::button-press-event handler.

2. During a #GtkWidget::motion-notify-event handler, if you want to start a drag
when the mouse moves past a certain threshold distance after a button-press.
Pass the @event that you have in your #GtkWidget::motion-notify-event handler.

3. During a timeout handler, if you want to start a drag after the mouse
button is held down for some time.  Try to save the last event that you got
from the mouse, using gdk_event_copy(), and pass it to this function
(remember to free the event with gdk_event_free() when you are done).
If you can really not pass a real event, pass #NULL instead.
*/
func (self *_TraitWidget) DragBeginWithCoordinates(targets *TargetList, actions C.GdkDragAction, button int, event *C.GdkEvent, x int, y int) (return__ *C.GdkDragContext) {
	return__ = C.gtk_drag_begin_with_coordinates(self.CPointer, (*C.GtkTargetList)(unsafe.Pointer(targets)), actions, C.gint(button), event, C.gint(x), C.gint(y))
	return
}

/*
Checks to see if a mouse drag starting at (@start_x, @start_y) and ending
at (@current_x, @current_y) has passed the GTK+ drag threshold, and thus
should trigger the beginning of a drag-and-drop operation.
*/
func (self *_TraitWidget) DragCheckThreshold(start_x int, start_y int, current_x int, current_y int) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_drag_check_threshold(self.CPointer, C.gint(start_x), C.gint(start_y), C.gint(current_x), C.gint(current_y))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Add the image targets supported by #GtkSelectionData to
the target list of the drag destination. The targets
are added with @info = 0. If you need another value,
use gtk_target_list_add_image_targets() and
gtk_drag_dest_set_target_list().
*/
func (self *_TraitWidget) DragDestAddImageTargets() {
	C.gtk_drag_dest_add_image_targets(self.CPointer)
	return
}

/*
Add the text targets supported by #GtkSelectionData to
the target list of the drag destination. The targets
are added with @info = 0. If you need another value,
use gtk_target_list_add_text_targets() and
gtk_drag_dest_set_target_list().
*/
func (self *_TraitWidget) DragDestAddTextTargets() {
	C.gtk_drag_dest_add_text_targets(self.CPointer)
	return
}

/*
Add the URI targets supported by #GtkSelectionData to
the target list of the drag destination. The targets
are added with @info = 0. If you need another value,
use gtk_target_list_add_uri_targets() and
gtk_drag_dest_set_target_list().
*/
func (self *_TraitWidget) DragDestAddUriTargets() {
	C.gtk_drag_dest_add_uri_targets(self.CPointer)
	return
}

/*
Looks for a match between the supported targets of @context and the
@dest_target_list, returning the first matching target, otherwise
returning %GDK_NONE. @dest_target_list should usually be the return
value from gtk_drag_dest_get_target_list(), but some widgets may
have different valid targets for different parts of the widget; in
that case, they will have to implement a drag_motion handler that
passes the correct target list to this function.
*/
func (self *_TraitWidget) DragDestFindTarget(context *C.GdkDragContext, target_list *TargetList) (return__ C.GdkAtom) {
	return__ = C.gtk_drag_dest_find_target(self.CPointer, context, (*C.GtkTargetList)(unsafe.Pointer(target_list)))
	return
}

/*
Returns the list of targets this widget can accept from
drag-and-drop.
*/
func (self *_TraitWidget) DragDestGetTargetList() (return__ *TargetList) {
	var __cgo__return__ *C.GtkTargetList
	__cgo__return__ = C.gtk_drag_dest_get_target_list(self.CPointer)
	return__ = (*TargetList)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Returns whether the widget has been configured to always
emit #GtkWidget::drag-motion signals.
*/
func (self *_TraitWidget) DragDestGetTrackMotion() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_drag_dest_get_track_motion(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets a widget as a potential drop destination, and adds default behaviors.

The default behaviors listed in @flags have an effect similar
to installing default handlers for the widget’s drag-and-drop signals
(#GtkWidget::drag-motion, #GtkWidget::drag-drop, ...). They all exist
for convenience. When passing #GTK_DEST_DEFAULT_ALL for instance it is
sufficient to connect to the widget’s #GtkWidget::drag-data-received
signal to get primitive, but consistent drag-and-drop support.

Things become more complicated when you try to preview the dragged data,
as described in the documentation for #GtkWidget::drag-motion. The default
behaviors described by @flags make some assumptions, that can conflict
with your own signal handlers. For instance #GTK_DEST_DEFAULT_DROP causes
invokations of gdk_drag_status() in the context of #GtkWidget::drag-motion,
and invokations of gtk_drag_finish() in #GtkWidget::drag-data-received.
Especially the later is dramatic, when your own #GtkWidget::drag-motion
handler calls gtk_drag_get_data() to inspect the dragged data.

There’s no way to set a default action here, you can use the
#GtkWidget::drag-motion callback for that. Here’s an example which selects
the action to use depending on whether the control key is pressed or not:
|[<!-- language="C" -->
static void
drag_motion (GtkWidget *widget,
             GdkDragContext *context,
             gint x,
             gint y,
             guint time)
{
  GdkModifierType mask;

  gdk_window_get_pointer (gtk_widget_get_window (widget),
                          NULL, NULL, &mask);
  if (mask & GDK_CONTROL_MASK)
    gdk_drag_status (context, GDK_ACTION_COPY, time);
  else
    gdk_drag_status (context, GDK_ACTION_MOVE, time);
}
]|
*/
func (self *_TraitWidget) DragDestSet(flags C.GtkDestDefaults, targets *TargetEntry, n_targets int, actions C.GdkDragAction) {
	C.gtk_drag_dest_set(self.CPointer, flags, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.gint(n_targets), actions)
	return
}

/*
Sets this widget as a proxy for drops to another window.
*/
func (self *_TraitWidget) DragDestSetProxy(proxy_window *C.GdkWindow, protocol C.GdkDragProtocol, use_coordinates bool) {
	__cgo__use_coordinates := C.gboolean(0)
	if use_coordinates {
		__cgo__use_coordinates = C.gboolean(1)
	}
	C.gtk_drag_dest_set_proxy(self.CPointer, proxy_window, protocol, __cgo__use_coordinates)
	return
}

/*
Sets the target types that this widget can accept from drag-and-drop.
The widget must first be made into a drag destination with
gtk_drag_dest_set().
*/
func (self *_TraitWidget) DragDestSetTargetList(target_list *TargetList) {
	C.gtk_drag_dest_set_target_list(self.CPointer, (*C.GtkTargetList)(unsafe.Pointer(target_list)))
	return
}

/*
Tells the widget to emit #GtkWidget::drag-motion and
#GtkWidget::drag-leave events regardless of the targets and the
%GTK_DEST_DEFAULT_MOTION flag.

This may be used when a widget wants to do generic
actions regardless of the targets that the source offers.
*/
func (self *_TraitWidget) DragDestSetTrackMotion(track_motion bool) {
	__cgo__track_motion := C.gboolean(0)
	if track_motion {
		__cgo__track_motion = C.gboolean(1)
	}
	C.gtk_drag_dest_set_track_motion(self.CPointer, __cgo__track_motion)
	return
}

/*
Clears information about a drop destination set with
gtk_drag_dest_set(). The widget will no longer receive
notification of drags.
*/
func (self *_TraitWidget) DragDestUnset() {
	C.gtk_drag_dest_unset(self.CPointer)
	return
}

/*
Gets the data associated with a drag. When the data
is received or the retrieval fails, GTK+ will emit a
#GtkWidget::drag-data-received signal. Failure of the retrieval
is indicated by the length field of the @selection_data
signal parameter being negative. However, when gtk_drag_get_data()
is called implicitely because the %GTK_DEST_DEFAULT_DROP was set,
then the widget will not receive notification of failed
drops.
*/
func (self *_TraitWidget) DragGetData(context *C.GdkDragContext, target C.GdkAtom, time_ uint32) {
	C.gtk_drag_get_data(self.CPointer, context, target, C.guint32(time_))
	return
}

/*
Draws a highlight around a widget. This will attach
handlers to #GtkWidget::draw, so the highlight
will continue to be displayed until gtk_drag_unhighlight()
is called.
*/
func (self *_TraitWidget) DragHighlight() {
	C.gtk_drag_highlight(self.CPointer)
	return
}

/*
Add the writable image targets supported by #GtkSelectionData to
the target list of the drag source. The targets
are added with @info = 0. If you need another value,
use gtk_target_list_add_image_targets() and
gtk_drag_source_set_target_list().
*/
func (self *_TraitWidget) DragSourceAddImageTargets() {
	C.gtk_drag_source_add_image_targets(self.CPointer)
	return
}

/*
Add the text targets supported by #GtkSelectionData to
the target list of the drag source.  The targets
are added with @info = 0. If you need another value,
use gtk_target_list_add_text_targets() and
gtk_drag_source_set_target_list().
*/
func (self *_TraitWidget) DragSourceAddTextTargets() {
	C.gtk_drag_source_add_text_targets(self.CPointer)
	return
}

/*
Add the URI targets supported by #GtkSelectionData to
the target list of the drag source.  The targets
are added with @info = 0. If you need another value,
use gtk_target_list_add_uri_targets() and
gtk_drag_source_set_target_list().
*/
func (self *_TraitWidget) DragSourceAddUriTargets() {
	C.gtk_drag_source_add_uri_targets(self.CPointer)
	return
}

/*
Gets the list of targets this widget can provide for
drag-and-drop.
*/
func (self *_TraitWidget) DragSourceGetTargetList() (return__ *TargetList) {
	var __cgo__return__ *C.GtkTargetList
	__cgo__return__ = C.gtk_drag_source_get_target_list(self.CPointer)
	return__ = (*TargetList)(unsafe.Pointer(__cgo__return__))
	return
}

/*
Sets up a widget so that GTK+ will start a drag operation when the user
clicks and drags on the widget. The widget must have a window.
*/
func (self *_TraitWidget) DragSourceSet(start_button_mask C.GdkModifierType, targets *TargetEntry, n_targets int, actions C.GdkDragAction) {
	C.gtk_drag_source_set(self.CPointer, start_button_mask, (*C.GtkTargetEntry)(unsafe.Pointer(targets)), C.gint(n_targets), actions)
	return
}

/*
Sets the icon that will be used for drags from a particular source
to @icon. See the docs for #GtkIconTheme for more details.
*/
func (self *_TraitWidget) DragSourceSetIconGicon(icon *C.GIcon) {
	C.gtk_drag_source_set_icon_gicon(self.CPointer, icon)
	return
}

/*
Sets the icon that will be used for drags from a particular source
to a themed icon. See the docs for #GtkIconTheme for more details.
*/
func (self *_TraitWidget) DragSourceSetIconName(icon_name string) {
	__cgo__icon_name := (*C.gchar)(unsafe.Pointer(C.CString(icon_name)))
	C.gtk_drag_source_set_icon_name(self.CPointer, __cgo__icon_name)
	C.free(unsafe.Pointer(__cgo__icon_name))
	return
}

/*
Sets the icon that will be used for drags from a particular widget
from a #GdkPixbuf. GTK+ retains a reference for @pixbuf and will
release it when it is no longer needed.
*/
func (self *_TraitWidget) DragSourceSetIconPixbuf(pixbuf *C.GdkPixbuf) {
	C.gtk_drag_source_set_icon_pixbuf(self.CPointer, pixbuf)
	return
}

// gtk_drag_source_set_icon_stock is not generated due to deprecation attr

/*
Changes the target types that this widget offers for drag-and-drop.
The widget must first be made into a drag source with
gtk_drag_source_set().
*/
func (self *_TraitWidget) DragSourceSetTargetList(target_list *TargetList) {
	C.gtk_drag_source_set_target_list(self.CPointer, (*C.GtkTargetList)(unsafe.Pointer(target_list)))
	return
}

/*
Undoes the effects of gtk_drag_source_set().
*/
func (self *_TraitWidget) DragSourceUnset() {
	C.gtk_drag_source_unset(self.CPointer)
	return
}

/*
Removes a highlight set by gtk_drag_highlight() from
a widget.
*/
func (self *_TraitWidget) DragUnhighlight() {
	C.gtk_drag_unhighlight(self.CPointer)
	return
}

/*
Draws @widget to @cr. The top left corner of the widget will be
drawn to the currently set origin point of @cr.

You should pass a cairo context as @cr argument that is in an
original state. Otherwise the resulting drawing is undefined. For
example changing the operator using cairo_set_operator() or the
line width using cairo_set_line_width() might have unwanted side
effects.
You may however change the context’s transform matrix - like with
cairo_scale(), cairo_translate() or cairo_set_matrix() and clip
region with cairo_clip() prior to calling this function. Also, it
is fine to modify the context with cairo_save() and
cairo_push_group() prior to calling this function.

Note that special-purpose widgets may contain special code for
rendering to the screen and might appear differently on screen
and when rendered using gtk_widget_draw().
*/
func (self *_TraitWidget) Draw(cr *C.cairo_t) {
	C.gtk_widget_draw(self.CPointer, cr)
	return
}

// gtk_widget_ensure_style is not generated due to deprecation attr

/*
Notifies the user about an input-related error on this widget.
If the #GtkSettings:gtk-error-bell setting is %TRUE, it calls
gdk_window_beep(), otherwise it does nothing.

Note that the effect of gdk_window_beep() can be configured in many
ways, depending on the windowing backend and the desktop environment
or window manager that is used.
*/
func (self *_TraitWidget) ErrorBell() {
	C.gtk_widget_error_bell(self.CPointer)
	return
}

/*
Rarely-used function. This function is used to emit
the event signals on a widget (those signals should never
be emitted without using this function to do so).
If you want to synthesize an event though, don’t use this function;
instead, use gtk_main_do_event() so the event will behave as if
it were in the event queue. Don’t synthesize expose events; instead,
use gdk_window_invalidate_rect() to invalidate a region of the
window.
*/
func (self *_TraitWidget) Event(event *C.GdkEvent) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_event(self.CPointer, event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Stops emission of #GtkWidget::child-notify signals on @widget. The
signals are queued until gtk_widget_thaw_child_notify() is called
on @widget.

This is the analogue of g_object_freeze_notify() for child properties.
*/
func (self *_TraitWidget) FreezeChildNotify() {
	C.gtk_widget_freeze_child_notify(self.CPointer)
	return
}

/*
Returns the accessible object that describes the widget to an
assistive technology.

If accessibility support is not available, this #AtkObject
instance may be a no-op. Likewise, if no class-specific #AtkObject
implementation is available for the widget instance in question,
it will inherit an #AtkObject implementation from the first ancestor
class for which such an implementation is defined.

The documentation of the
[ATK](http://developer.gnome.org/atk/stable/)
library contains more information about accessible objects and their uses.
*/
func (self *_TraitWidget) GetAccessible() (return__ *C.AtkObject) {
	return__ = C.gtk_widget_get_accessible(self.CPointer)
	return
}

/*
Returns the baseline that has currently been allocated to @widget.
This function is intended to be used when implementing handlers
for the #GtkWidget::draw function, and when allocating child
widgets in #GtkWidget::size_allocate.
*/
func (self *_TraitWidget) GetAllocatedBaseline() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gtk_widget_get_allocated_baseline(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the height that has currently been allocated to @widget.
This function is intended to be used when implementing handlers
for the #GtkWidget::draw function.
*/
func (self *_TraitWidget) GetAllocatedHeight() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gtk_widget_get_allocated_height(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the width that has currently been allocated to @widget.
This function is intended to be used when implementing handlers
for the #GtkWidget::draw function.
*/
func (self *_TraitWidget) GetAllocatedWidth() (return__ int) {
	var __cgo__return__ C.int
	__cgo__return__ = C.gtk_widget_get_allocated_width(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Retrieves the widget’s allocation.

Note, when implementing a #GtkContainer: a widget’s allocation will
be its “adjusted” allocation, that is, the widget’s parent
container typically calls gtk_widget_size_allocate() with an
allocation, and that allocation is then adjusted (to handle margin
and alignment for example) before assignment to the widget.
gtk_widget_get_allocation() returns the adjusted allocation that
was actually assigned to the widget. The adjusted allocation is
guaranteed to be completely contained within the
gtk_widget_size_allocate() allocation, however. So a #GtkContainer
is guaranteed that its children stay inside the assigned bounds,
but not that they have exactly the bounds the container assigned.
There is no way to get the original allocation assigned by
gtk_widget_size_allocate(), since it isn’t stored; if a container
implementation needs that information it will have to track it itself.
*/
func (self *_TraitWidget) GetAllocation() (allocation C.GtkAllocation) {
	C.gtk_widget_get_allocation(self.CPointer, &allocation)
	return
}

/*
Gets the first ancestor of @widget with type @widget_type. For example,
`gtk_widget_get_ancestor (widget, GTK_TYPE_BOX)` gets
the first #GtkBox that’s an ancestor of @widget. No reference will be
added to the returned widget; it should not be unreferenced. See note
about checking for a toplevel #GtkWindow in the docs for
gtk_widget_get_toplevel().

Note that unlike gtk_widget_is_ancestor(), gtk_widget_get_ancestor()
considers @widget to be an ancestor of itself.
*/
func (self *_TraitWidget) GetAncestor(widget_type C.GType) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_widget_get_ancestor(self.CPointer, widget_type)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Determines whether the application intends to draw on the widget in
an #GtkWidget::draw handler.

See gtk_widget_set_app_paintable()
*/
func (self *_TraitWidget) GetAppPaintable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_app_paintable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether @widget can be a default widget. See
gtk_widget_set_can_default().
*/
func (self *_TraitWidget) GetCanDefault() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_can_default(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether @widget can own the input focus. See
gtk_widget_set_can_focus().
*/
func (self *_TraitWidget) GetCanFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_can_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_widget_get_child_requisition is not generated due to deprecation attr

/*
Gets the value set with gtk_widget_set_child_visible().
If you feel a need to use this function, your code probably
needs reorganization.

This function is only useful for container implementations and
never should be called by an application.
*/
func (self *_TraitWidget) GetChildVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_child_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the clipboard object for the given selection to
be used with @widget. @widget must have a #GdkDisplay
associated with it, so must be attached to a toplevel
window.
*/
func (self *_TraitWidget) GetClipboard(selection C.GdkAtom) (return__ *Clipboard) {
	var __cgo__return__ *C.GtkClipboard
	__cgo__return__ = C.gtk_widget_get_clipboard(self.CPointer, selection)
	return__ = NewClipboardFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

// gtk_widget_get_composite_name is not generated due to deprecation attr

/*
Returns whether @device can interact with @widget and its
children. See gtk_widget_set_device_enabled().
*/
func (self *_TraitWidget) GetDeviceEnabled(device *C.GdkDevice) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_device_enabled(self.CPointer, device)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the events mask for the widget corresponding to an specific device. These
are the events that the widget will receive when @device operates on it.
*/
func (self *_TraitWidget) GetDeviceEvents(device *C.GdkDevice) (return__ C.GdkEventMask) {
	return__ = C.gtk_widget_get_device_events(self.CPointer, device)
	return
}

/*
Gets the reading direction for a particular widget. See
gtk_widget_set_direction().
*/
func (self *_TraitWidget) GetDirection() (return__ C.GtkTextDirection) {
	return__ = C.gtk_widget_get_direction(self.CPointer)
	return
}

/*
Get the #GdkDisplay for the toplevel window associated with
this widget. This function can only be called after the widget
has been added to a widget hierarchy with a #GtkWindow at the top.

In general, you should only create display specific
resources when a widget has been realized, and you should
free those resources when the widget is unrealized.
*/
func (self *_TraitWidget) GetDisplay() (return__ *C.GdkDisplay) {
	return__ = C.gtk_widget_get_display(self.CPointer)
	return
}

/*
Determines whether the widget is double buffered.

See gtk_widget_set_double_buffered()
*/
func (self *_TraitWidget) GetDoubleBuffered() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_double_buffered(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the event mask for the widget (a bitfield containing flags
from the #GdkEventMask enumeration). These are the events that the widget
will receive.
*/
func (self *_TraitWidget) GetEvents() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_widget_get_events(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Obtains the frame clock for a widget. The frame clock is a global
“ticker” that can be used to drive animations and repaints.  The
most common reason to get the frame clock is to call
gdk_frame_clock_get_frame_time(), in order to get a time to use for
animating. For example you might record the start of the animation
with an initial value from gdk_frame_clock_get_frame_time(), and
then update the animation by calling
gdk_frame_clock_get_frame_time() again during each repaint.

gdk_frame_clock_request_phase() will result in a new frame on the
clock, but won’t necessarily repaint any widgets. To repaint a
widget, you have to use gtk_widget_queue_draw() which invalidates
the widget (thus scheduling it to receive a draw on the next
frame). gtk_widget_queue_draw() will also end up requesting a frame
on the appropriate frame clock.

A widget’s frame clock will not change while the widget is
mapped. Reparenting a widget (which implies a temporary unmap) can
change the widget’s frame clock.

Unrealized widgets do not have a frame clock.
*/
func (self *_TraitWidget) GetFrameClock() (return__ *C.GdkFrameClock) {
	return__ = C.gtk_widget_get_frame_clock(self.CPointer)
	return
}

/*
Gets the value of the #GtkWidget:halign property.

For backwards compatibility reasons this method will never return
%GTK_ALIGN_BASELINE, but instead it will convert it to
%GTK_ALIGN_FILL. Baselines are not supported for horizontal
alignment.
*/
func (self *_TraitWidget) GetHalign() (return__ C.GtkAlign) {
	return__ = C.gtk_widget_get_halign(self.CPointer)
	return
}

/*
Returns the current value of the has-tooltip property.  See
#GtkWidget:has-tooltip for more information.
*/
func (self *_TraitWidget) GetHasTooltip() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_has_tooltip(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether @widget has a #GdkWindow of its own. See
gtk_widget_set_has_window().
*/
func (self *_TraitWidget) GetHasWindow() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_has_window(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether the widget would like any available extra horizontal
space. When a user resizes a #GtkWindow, widgets with expand=TRUE
generally receive the extra space. For example, a list or
scrollable area or document in your window would often be set to
expand.

Containers should use gtk_widget_compute_expand() rather than
this function, to see whether a widget, or any of its children,
has the expand flag set. If any child of a widget wants to
expand, the parent may ask to expand also.

This function only looks at the widget’s own hexpand flag, rather
than computing whether the entire widget tree rooted at this widget
wants to expand.
*/
func (self *_TraitWidget) GetHexpand() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_hexpand(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether gtk_widget_set_hexpand() has been used to
explicitly set the expand flag on this widget.

If hexpand is set, then it overrides any computed
expand value based on child widgets. If hexpand is not
set, then the expand value depends on whether any
children of the widget would like to expand.

There are few reasons to use this function, but it’s here
for completeness and consistency.
*/
func (self *_TraitWidget) GetHexpandSet() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_hexpand_set(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether the widget is mapped.
*/
func (self *_TraitWidget) GetMapped() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_mapped(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of the #GtkWidget:margin-bottom property.
*/
func (self *_TraitWidget) GetMarginBottom() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_widget_get_margin_bottom(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value of the #GtkWidget:margin-end property.
*/
func (self *_TraitWidget) GetMarginEnd() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_widget_get_margin_end(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

// gtk_widget_get_margin_left is not generated due to deprecation attr

// gtk_widget_get_margin_right is not generated due to deprecation attr

/*
Gets the value of the #GtkWidget:margin-start property.
*/
func (self *_TraitWidget) GetMarginStart() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_widget_get_margin_start(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Gets the value of the #GtkWidget:margin-top property.
*/
func (self *_TraitWidget) GetMarginTop() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_widget_get_margin_top(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Returns the modifier mask the @widget’s windowing system backend
uses for a particular purpose.

See gdk_keymap_get_modifier_mask().
*/
func (self *_TraitWidget) GetModifierMask(intent C.GdkModifierIntent) (return__ C.GdkModifierType) {
	return__ = C.gtk_widget_get_modifier_mask(self.CPointer, intent)
	return
}

// gtk_widget_get_modifier_style is not generated due to deprecation attr

/*
Retrieves the name of a widget. See gtk_widget_set_name() for the
significance of widget names.
*/
func (self *_TraitWidget) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_widget_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the current value of the #GtkWidget:no-show-all property,
which determines whether calls to gtk_widget_show_all()
will affect this widget.
*/
func (self *_TraitWidget) GetNoShowAll() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_no_show_all(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Fetches the requested opacity for this widget. See
gtk_widget_set_opacity().
*/
func (self *_TraitWidget) GetOpacity() (return__ float64) {
	var __cgo__return__ C.double
	__cgo__return__ = C.gtk_widget_get_opacity(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets a #PangoContext with the appropriate font map, font description,
and base direction for this widget. Unlike the context returned
by gtk_widget_create_pango_context(), this context is owned by
the widget (it can be used until the screen for the widget changes
or the widget is removed from its toplevel), and will be updated to
match any changes to the widget’s attributes. This can be tracked
by using the #GtkWidget::screen-changed signal on the widget.
*/
func (self *_TraitWidget) GetPangoContext() (return__ *C.PangoContext) {
	return__ = C.gtk_widget_get_pango_context(self.CPointer)
	return
}

/*
Returns the parent container of @widget.
*/
func (self *_TraitWidget) GetParent() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_widget_get_parent(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets @widget’s parent window.
*/
func (self *_TraitWidget) GetParentWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_widget_get_parent_window(self.CPointer)
	return
}

/*
Returns the #GtkWidgetPath representing @widget, if the widget
is not connected to a toplevel widget, a partial path will be
created.
*/
func (self *_TraitWidget) GetPath() (return__ *WidgetPath) {
	var __cgo__return__ *C.GtkWidgetPath
	__cgo__return__ = C.gtk_widget_get_path(self.CPointer)
	return__ = (*WidgetPath)(unsafe.Pointer(__cgo__return__))
	return
}

// gtk_widget_get_pointer is not generated due to deprecation attr

/*
Retrieves a widget’s initial minimum and natural height.

This call is specific to width-for-height requests.

The returned request will be modified by the
GtkWidgetClass::adjust_size_request virtual method and by any
#GtkSizeGroups that have been applied. That is, the returned request
is the one that should be used for layout, not necessarily the one
returned by the widget itself.
*/
func (self *_TraitWidget) GetPreferredHeight() (minimum_height int, natural_height int) {
	var __cgo__minimum_height C.gint
	var __cgo__natural_height C.gint
	C.gtk_widget_get_preferred_height(self.CPointer, &__cgo__minimum_height, &__cgo__natural_height)
	minimum_height = int(__cgo__minimum_height)
	natural_height = int(__cgo__natural_height)
	return
}

/*
Retrieves a widget’s minimum and natural height and the corresponding baselines if it would be given
the specified @width, or the default height if @width is -1. The baselines may be -1 which means
that no baseline is requested for this widget.

The returned request will be modified by the
GtkWidgetClass::adjust_size_request and GtkWidgetClass::adjust_baseline_request virtual methods
and by any #GtkSizeGroups that have been applied. That is, the returned request
is the one that should be used for layout, not necessarily the one
returned by the widget itself.
*/
func (self *_TraitWidget) GetPreferredHeightAndBaselineForWidth(width int) (minimum_height int, natural_height int, minimum_baseline int, natural_baseline int) {
	var __cgo__minimum_height C.gint
	var __cgo__natural_height C.gint
	var __cgo__minimum_baseline C.gint
	var __cgo__natural_baseline C.gint
	C.gtk_widget_get_preferred_height_and_baseline_for_width(self.CPointer, C.gint(width), &__cgo__minimum_height, &__cgo__natural_height, &__cgo__minimum_baseline, &__cgo__natural_baseline)
	minimum_height = int(__cgo__minimum_height)
	natural_height = int(__cgo__natural_height)
	minimum_baseline = int(__cgo__minimum_baseline)
	natural_baseline = int(__cgo__natural_baseline)
	return
}

/*
Retrieves a widget’s minimum and natural height if it would be given
the specified @width.

The returned request will be modified by the
GtkWidgetClass::adjust_size_request virtual method and by any
#GtkSizeGroups that have been applied. That is, the returned request
is the one that should be used for layout, not necessarily the one
returned by the widget itself.
*/
func (self *_TraitWidget) GetPreferredHeightForWidth(width int) (minimum_height int, natural_height int) {
	var __cgo__minimum_height C.gint
	var __cgo__natural_height C.gint
	C.gtk_widget_get_preferred_height_for_width(self.CPointer, C.gint(width), &__cgo__minimum_height, &__cgo__natural_height)
	minimum_height = int(__cgo__minimum_height)
	natural_height = int(__cgo__natural_height)
	return
}

/*
Retrieves the minimum and natural size of a widget, taking
into account the widget’s preference for height-for-width management.

This is used to retrieve a suitable size by container widgets which do
not impose any restrictions on the child placement. It can be used
to deduce toplevel window and menu sizes as well as child widgets in
free-form containers such as GtkLayout.

Handle with care. Note that the natural height of a height-for-width
widget will generally be a smaller size than the minimum height, since the required
height for the natural width is generally smaller than the required height for
the minimum width.

Use gtk_widget_get_preferred_height_and_baseline_for_width() if you want to support
baseline alignment.
*/
func (self *_TraitWidget) GetPreferredSize() (minimum_size C.GtkRequisition, natural_size C.GtkRequisition) {
	C.gtk_widget_get_preferred_size(self.CPointer, &minimum_size, &natural_size)
	return
}

/*
Retrieves a widget’s initial minimum and natural width.

This call is specific to height-for-width requests.

The returned request will be modified by the
GtkWidgetClass::adjust_size_request virtual method and by any
#GtkSizeGroups that have been applied. That is, the returned request
is the one that should be used for layout, not necessarily the one
returned by the widget itself.
*/
func (self *_TraitWidget) GetPreferredWidth() (minimum_width int, natural_width int) {
	var __cgo__minimum_width C.gint
	var __cgo__natural_width C.gint
	C.gtk_widget_get_preferred_width(self.CPointer, &__cgo__minimum_width, &__cgo__natural_width)
	minimum_width = int(__cgo__minimum_width)
	natural_width = int(__cgo__natural_width)
	return
}

/*
Retrieves a widget’s minimum and natural width if it would be given
the specified @height.

The returned request will be modified by the
GtkWidgetClass::adjust_size_request virtual method and by any
#GtkSizeGroups that have been applied. That is, the returned request
is the one that should be used for layout, not necessarily the one
returned by the widget itself.
*/
func (self *_TraitWidget) GetPreferredWidthForHeight(height int) (minimum_width int, natural_width int) {
	var __cgo__minimum_width C.gint
	var __cgo__natural_width C.gint
	C.gtk_widget_get_preferred_width_for_height(self.CPointer, C.gint(height), &__cgo__minimum_width, &__cgo__natural_width)
	minimum_width = int(__cgo__minimum_width)
	natural_width = int(__cgo__natural_width)
	return
}

/*
Determines whether @widget is realized.
*/
func (self *_TraitWidget) GetRealized() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_realized(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether @widget is always treated as the default widget
within its toplevel when it has the focus, even if another widget
is the default.

See gtk_widget_set_receives_default().
*/
func (self *_TraitWidget) GetReceivesDefault() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_receives_default(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether the widget prefers a height-for-width layout
or a width-for-height layout.

#GtkBin widgets generally propagate the preference of
their child, container widgets need to request something either in
context of their children or in context of their allocation
capabilities.
*/
func (self *_TraitWidget) GetRequestMode() (return__ C.GtkSizeRequestMode) {
	return__ = C.gtk_widget_get_request_mode(self.CPointer)
	return
}

// gtk_widget_get_requisition is not generated due to deprecation attr

// gtk_widget_get_root_window is not generated due to deprecation attr

/*
Retrieves the internal scale factor that maps from window coordinates
to the actual device pixels. On traditional systems this is 1, on
high density outputs, it can be a higher value (typically 2).

See gdk_window_get_scale_factor().
*/
func (self *_TraitWidget) GetScaleFactor() (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_widget_get_scale_factor(self.CPointer)
	return__ = int(__cgo__return__)
	return
}

/*
Get the #GdkScreen from the toplevel window associated with
this widget. This function can only be called after the widget
has been added to a widget hierarchy with a #GtkWindow
at the top.

In general, you should only create screen specific
resources when a widget has been realized, and you should
free those resources when the widget is unrealized.
*/
func (self *_TraitWidget) GetScreen() (return__ *C.GdkScreen) {
	return__ = C.gtk_widget_get_screen(self.CPointer)
	return
}

/*
Returns the widget’s sensitivity (in the sense of returning
the value that has been set using gtk_widget_set_sensitive()).

The effective sensitivity of a widget is however determined by both its
own and its parent widget’s sensitivity. See gtk_widget_is_sensitive().
*/
func (self *_TraitWidget) GetSensitive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_sensitive(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the settings object holding the settings used for this widget.

Note that this function can only be called when the #GtkWidget
is attached to a toplevel, since the settings object is specific
to a particular #GdkScreen.
*/
func (self *_TraitWidget) GetSettings() (return__ *Settings) {
	var __cgo__return__ *C.GtkSettings
	__cgo__return__ = C.gtk_widget_get_settings(self.CPointer)
	return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the size request that was explicitly set for the widget using
gtk_widget_set_size_request(). A value of -1 stored in @width or
@height indicates that that dimension has not been set explicitly
and the natural requisition of the widget will be used intead. See
gtk_widget_set_size_request(). To get the size a widget will
actually request, call gtk_widget_get_preferred_size() instead of
this function.
*/
func (self *_TraitWidget) GetSizeRequest() (width int, height int) {
	var __cgo__width C.gint
	var __cgo__height C.gint
	C.gtk_widget_get_size_request(self.CPointer, &__cgo__width, &__cgo__height)
	width = int(__cgo__width)
	height = int(__cgo__height)
	return
}

// gtk_widget_get_state is not generated due to deprecation attr

/*
Returns the widget state as a flag set. It is worth mentioning
that the effective %GTK_STATE_FLAG_INSENSITIVE state will be
returned, that is, also based on parent insensitivity, even if
@widget itself is sensitive.
*/
func (self *_TraitWidget) GetStateFlags() (return__ C.GtkStateFlags) {
	return__ = C.gtk_widget_get_state_flags(self.CPointer)
	return
}

// gtk_widget_get_style is not generated due to deprecation attr

/*
Returns the style context associated to @widget.
*/
func (self *_TraitWidget) GetStyleContext() (return__ *StyleContext) {
	var __cgo__return__ *C.GtkStyleContext
	__cgo__return__ = C.gtk_widget_get_style_context(self.CPointer)
	return__ = NewStyleContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns %TRUE if @widget is multiple pointer aware. See
gtk_widget_set_support_multidevice() for more information.
*/
func (self *_TraitWidget) GetSupportMultidevice() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_support_multidevice(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Fetch an object build from the template XML for @widget_type in this @widget instance.

This will only report children which were previously declared with
gtk_widget_class_bind_template_child_full() or one of its
variants.

This function is only meant to be called for code which is private to the @widget_type which
declared the child and is meant for language bindings which cannot easily make use
of the GObject structure offsets.
*/
func (self *_TraitWidget) GetTemplateChild(widget_type C.GType, name string) (return__ *C.GObject) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	return__ = C.gtk_widget_get_template_child(self.CPointer, widget_type, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Gets the contents of the tooltip for @widget.
*/
func (self *_TraitWidget) GetTooltipMarkup() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_widget_get_tooltip_markup(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the contents of the tooltip for @widget.
*/
func (self *_TraitWidget) GetTooltipText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_widget_get_tooltip_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the #GtkWindow of the current tooltip. This can be the
GtkWindow created by default, or the custom tooltip window set
using gtk_widget_set_tooltip_window().
*/
func (self *_TraitWidget) GetTooltipWindow() (return__ *Window) {
	var __cgo__return__ *C.GtkWindow
	__cgo__return__ = C.gtk_widget_get_tooltip_window(self.CPointer)
	return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
This function returns the topmost widget in the container hierarchy
@widget is a part of. If @widget has no parent widgets, it will be
returned as the topmost widget. No reference will be added to the
returned widget; it should not be unreferenced.

Note the difference in behavior vs. gtk_widget_get_ancestor();
`gtk_widget_get_ancestor (widget, GTK_TYPE_WINDOW)`
would return
%NULL if @widget wasn’t inside a toplevel window, and if the
window was inside a #GtkWindow-derived widget which was in turn
inside the toplevel #GtkWindow. While the second case may
seem unlikely, it actually happens when a #GtkPlug is embedded
inside a #GtkSocket within the same application.

To reliably find the toplevel #GtkWindow, use
gtk_widget_get_toplevel() and call gtk_widget_is_toplevel()
on the result.
|[<!-- language="C" -->
 GtkWidget *toplevel = gtk_widget_get_toplevel (widget);
 if (gtk_widget_is_toplevel (toplevel))
   {
     // Perform action on toplevel.
   }
]|
*/
func (self *_TraitWidget) GetToplevel() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_widget_get_toplevel(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the value of the #GtkWidget:valign property.

For backwards compatibility reasons this method will never return
%GTK_ALIGN_BASELINE, but instead it will convert it to
%GTK_ALIGN_FILL. If your widget want to support baseline aligned
children it must use gtk_widget_get_valign_with_baseline().
*/
func (self *_TraitWidget) GetValign() (return__ C.GtkAlign) {
	return__ = C.gtk_widget_get_valign(self.CPointer)
	return
}

/*
Gets the value of the #GtkWidget:valign property, including
%GTK_ALIGN_BASELINE.
*/
func (self *_TraitWidget) GetValignWithBaseline() (return__ C.GtkAlign) {
	return__ = C.gtk_widget_get_valign_with_baseline(self.CPointer)
	return
}

/*
Gets whether the widget would like any available extra vertical
space.

See gtk_widget_get_hexpand() for more detail.
*/
func (self *_TraitWidget) GetVexpand() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_vexpand(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether gtk_widget_set_vexpand() has been used to
explicitly set the expand flag on this widget.

See gtk_widget_get_hexpand_set() for more detail.
*/
func (self *_TraitWidget) GetVexpandSet() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_vexpand_set(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether the widget is visible. If you want to
take into account whether the widget’s parent is also marked as
visible, use gtk_widget_is_visible() instead.

This function does not check if the widget is obscured in any way.

See gtk_widget_set_visible().
*/
func (self *_TraitWidget) GetVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_get_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the visual that will be used to render @widget.
*/
func (self *_TraitWidget) GetVisual() (return__ *C.GdkVisual) {
	return__ = C.gtk_widget_get_visual(self.CPointer)
	return
}

/*
Returns the widget’s window if it is realized, %NULL otherwise
*/
func (self *_TraitWidget) GetWindow() (return__ *C.GdkWindow) {
	return__ = C.gtk_widget_get_window(self.CPointer)
	return
}

/*
Makes @widget the current grabbed widget.

This means that interaction with other widgets in the same
application is blocked and mouse as well as keyboard events
are delivered to this widget.

If @widget is not sensitive, it is not set as the current
grabbed widget and this function does nothing.
*/
func (self *_TraitWidget) GrabAdd() {
	C.gtk_grab_add(self.CPointer)
	return
}

/*
Causes @widget to become the default widget. @widget must be able to be
a default widget; typically you would ensure this yourself
by calling gtk_widget_set_can_default() with a %TRUE value.
The default widget is activated when
the user presses Enter in a window. Default widgets must be
activatable, that is, gtk_widget_activate() should affect them. Note
that #GtkEntry widgets require the “activates-default” property
set to %TRUE before they activate the default widget when Enter
is pressed and the #GtkEntry is focused.
*/
func (self *_TraitWidget) GrabDefault() {
	C.gtk_widget_grab_default(self.CPointer)
	return
}

/*
Causes @widget to have the keyboard focus for the #GtkWindow it's
inside. @widget must be a focusable widget, such as a #GtkEntry;
something like #GtkFrame won’t work.

More precisely, it must have the %GTK_CAN_FOCUS flag set. Use
gtk_widget_set_can_focus() to modify that flag.

The widget also needs to be realized and mapped. This is indicated by the
related signals. Grabbing the focus immediately after creating the widget
will likely fail and cause critical warnings.
*/
func (self *_TraitWidget) GrabFocus() {
	C.gtk_widget_grab_focus(self.CPointer)
	return
}

/*
Removes the grab from the given widget.

You have to pair calls to gtk_grab_add() and gtk_grab_remove().

If @widget does not have the grab, this function does nothing.
*/
func (self *_TraitWidget) GrabRemove() {
	C.gtk_grab_remove(self.CPointer)
	return
}

/*
Determines whether @widget is the current default widget within its
toplevel. See gtk_widget_set_can_default().
*/
func (self *_TraitWidget) HasDefault() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_has_default(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if the widget has the global input focus. See
gtk_widget_is_focus() for the difference between having the global
input focus, and only having the focus within a toplevel.
*/
func (self *_TraitWidget) HasFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_has_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether the widget is currently grabbing events, so it
is the only widget receiving input events (keyboard and mouse).

See also gtk_grab_add().
*/
func (self *_TraitWidget) HasGrab() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_has_grab(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_widget_has_rc_style is not generated due to deprecation attr

/*
Checks whether there is a #GdkScreen is associated with
this widget. All toplevel widgets have an associated
screen, and all widgets added into a hierarchy with a toplevel
window at the top.
*/
func (self *_TraitWidget) HasScreen() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_has_screen(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if the widget should show a visible indication that
it has the global input focus. This is a convenience function for
use in ::draw handlers that takes into account whether focus
indication should currently be shown in the toplevel window of
@widget. See gtk_window_get_focus_visible() for more information
about focus indication.

To find out if the widget has the global input focus, use
gtk_widget_has_focus().
*/
func (self *_TraitWidget) HasVisibleFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_has_visible_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Reverses the effects of gtk_widget_show(), causing the widget to be
hidden (invisible to the user).
*/
func (self *_TraitWidget) Hide() {
	C.gtk_widget_hide(self.CPointer)
	return
}

/*
Utility function; intended to be connected to the #GtkWidget::delete-event
signal on a #GtkWindow. The function calls gtk_widget_hide() on its
argument, then returns %TRUE. If connected to ::delete-event, the
result is that clicking the close button for a window (on the
window frame, top right corner usually) will hide but not destroy
the window. By default, GTK+ destroys windows when ::delete-event
is received.
*/
func (self *_TraitWidget) HideOnDelete() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_hide_on_delete(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the widget is currently being destroyed.
This information can sometimes be used to avoid doing
unnecessary work.
*/
func (self *_TraitWidget) InDestruction() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_in_destruction(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Creates and initializes child widgets defined in templates. This
function must be called in the instance initializer for any
class which assigned itself a template using gtk_widget_class_set_template()

It is important to call this function in the instance initializer
of a #GtkWidget subclass and not in #GObject.constructed() or
#GObject.constructor() for two reasons.

One reason is that generally derived widgets will assume that parent
class composite widgets have been created in their instance
initializers.

Another reason is that when calling g_object_new() on a widget with
composite templates, it’s important to build the composite widgets
before the construct properties are set. Properties passed to g_object_new()
should take precedence over properties set in the private template XML.
*/
func (self *_TraitWidget) InitTemplate() {
	C.gtk_widget_init_template(self.CPointer)
	return
}

/*
Sets an input shape for this widget’s GDK window. This allows for
windows which react to mouse click in a nonrectangular region, see
gdk_window_input_shape_combine_region() for more information.
*/
func (self *_TraitWidget) InputShapeCombineRegion(region *C.cairo_region_t) {
	C.gtk_widget_input_shape_combine_region(self.CPointer, region)
	return
}

/*
Inserts @group into @widget. Children of @widget that implement
#GtkActionable can then be associated with actions in @group by
setting their “action-name” to
@prefix.`action-name`.

If @group is %NULL, a previously inserted group for @name is removed
from @widget.
*/
func (self *_TraitWidget) InsertActionGroup(name string, group *C.GActionGroup) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_widget_insert_action_group(self.CPointer, __cgo__name, group)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Computes the intersection of a @widget’s area and @area, storing
the intersection in @intersection, and returns %TRUE if there was
an intersection.  @intersection may be %NULL if you’re only
interested in whether there was an intersection.
*/
func (self *_TraitWidget) Intersect(area *C.GdkRectangle, intersection *C.GdkRectangle) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_intersect(self.CPointer, area, intersection)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether @widget is somewhere inside @ancestor, possibly with
intermediate containers.
*/
func (self *_TraitWidget) IsAncestor(ancestor *Widget) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_is_ancestor(self.CPointer, (*C.GtkWidget)(ancestor.CPointer))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether @widget can rely on having its alpha channel
drawn correctly. On X11 this function returns whether a
compositing manager is running for @widget’s screen.

Please note that the semantics of this call will change
in the future if used on a widget that has a composited
window in its hierarchy (as set by gdk_window_set_composited()).
*/
func (self *_TraitWidget) IsComposited() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_is_composited(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether @widget can be drawn to. A widget can be drawn
to if it is mapped and visible.
*/
func (self *_TraitWidget) IsDrawable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_is_drawable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines if the widget is the focus widget within its
toplevel. (This does not mean that the #GtkWidget:has-focus property is
necessarily set; #GtkWidget:has-focus will only be set if the
toplevel widget additionally has the global input focus.)
*/
func (self *_TraitWidget) IsFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_is_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the widget’s effective sensitivity, which means
it is sensitive itself and also its parent widget is sensitive
*/
func (self *_TraitWidget) IsSensitive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_is_sensitive(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether @widget is a toplevel widget.

Currently only #GtkWindow and #GtkInvisible (and out-of-process
#GtkPlugs) are toplevel widgets. Toplevel widgets have no parent
widget.
*/
func (self *_TraitWidget) IsToplevel() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_is_toplevel(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether the widget and all its parents are marked as
visible.

This function does not check if the widget is obscured in any way.

See also gtk_widget_get_visible() and gtk_widget_set_visible()
*/
func (self *_TraitWidget) IsVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_is_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
This function should be called whenever keyboard navigation within
a single widget hits a boundary. The function emits the
#GtkWidget::keynav-failed signal on the widget and its return
value should be interpreted in a way similar to the return value of
gtk_widget_child_focus():

When %TRUE is returned, stay in the widget, the failed keyboard
navigation is Ok and/or there is nowhere we can/should move the
focus to.

When %FALSE is returned, the caller should continue with keyboard
navigation outside the widget, e.g. by calling
gtk_widget_child_focus() on the widget’s toplevel.

The default ::keynav-failed handler returns %TRUE for
%GTK_DIR_TAB_FORWARD and %GTK_DIR_TAB_BACKWARD. For the other
values of #GtkDirectionType it returns %FALSE.

Whenever the default handler returns %TRUE, it also calls
gtk_widget_error_bell() to notify the user of the failed keyboard
navigation.

A use case for providing an own implementation of ::keynav-failed
(either by connecting to it or by overriding it) would be a row of
#GtkEntry widgets where the user should be able to navigate the
entire row with the cursor keys, as e.g. known from user interfaces
that require entering license keys.
*/
func (self *_TraitWidget) KeynavFailed(direction C.GtkDirectionType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_keynav_failed(self.CPointer, direction)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Lists the closures used by @widget for accelerator group connections
with gtk_accel_group_connect_by_path() or gtk_accel_group_connect().
The closures can be used to monitor accelerator changes on @widget,
by connecting to the @GtkAccelGroup::accel-changed signal of the
#GtkAccelGroup of a closure which can be found out with
gtk_accel_group_from_accel_closure().
*/
func (self *_TraitWidget) ListAccelClosures() (return__ *C.GList) {
	return__ = C.gtk_widget_list_accel_closures(self.CPointer)
	return
}

/*
Returns a newly allocated list of the widgets, normally labels, for
which this widget is the target of a mnemonic (see for example,
gtk_label_set_mnemonic_widget()).

The widgets in the list are not individually referenced. If you
want to iterate through the list and perform actions involving
callbacks that might destroy the widgets, you
must call `g_list_foreach (result,
(GFunc)g_object_ref, NULL)` first, and then unref all the
widgets afterwards.
*/
func (self *_TraitWidget) ListMnemonicLabels() (return__ *C.GList) {
	return__ = C.gtk_widget_list_mnemonic_labels(self.CPointer)
	return
}

/*
This function is only for use in widget implementations. Causes
a widget to be mapped if it isn’t already.
*/
func (self *_TraitWidget) Map() {
	C.gtk_widget_map(self.CPointer)
	return
}

/*
Emits the #GtkWidget::mnemonic-activate signal.

The default handler for this signal activates the @widget if
@group_cycling is %FALSE, and just grabs the focus if @group_cycling
is %TRUE.
*/
func (self *_TraitWidget) MnemonicActivate(group_cycling bool) (return__ bool) {
	__cgo__group_cycling := C.gboolean(0)
	if group_cycling {
		__cgo__group_cycling = C.gboolean(1)
	}
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_mnemonic_activate(self.CPointer, __cgo__group_cycling)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_widget_modify_base is not generated due to deprecation attr

// gtk_widget_modify_bg is not generated due to deprecation attr

// gtk_widget_modify_cursor is not generated due to deprecation attr

// gtk_widget_modify_fg is not generated due to deprecation attr

// gtk_widget_modify_font is not generated due to deprecation attr

// gtk_widget_modify_style is not generated due to deprecation attr

// gtk_widget_modify_text is not generated due to deprecation attr

/*
Sets the background color to use for a widget.

All other style values are left untouched.
See gtk_widget_override_color().
*/
func (self *_TraitWidget) OverrideBackgroundColor(state C.GtkStateFlags, color *C.GdkRGBA) {
	C.gtk_widget_override_background_color(self.CPointer, state, color)
	return
}

/*
Sets the color to use for a widget.

All other style values are left untouched.

This function does not act recursively. Setting the color of a
container does not affect its children. Note that some widgets that
you may not think of as containers, for instance #GtkButtons,
are actually containers.

This API is mostly meant as a quick way for applications to
change a widget appearance. If you are developing a widgets
library and intend this change to be themeable, it is better
done by setting meaningful CSS classes and regions in your
widget/container implementation through gtk_style_context_add_class()
and gtk_style_context_add_region().

This way, your widget library can install a #GtkCssProvider
with the %GTK_STYLE_PROVIDER_PRIORITY_FALLBACK priority in order
to provide a default styling for those widgets that need so, and
this theming may fully overridden by the user’s theme.

Note that for complex widgets this may bring in undesired
results (such as uniform background color everywhere), in
these cases it is better to fully style such widgets through a
#GtkCssProvider with the %GTK_STYLE_PROVIDER_PRIORITY_APPLICATION
priority.
*/
func (self *_TraitWidget) OverrideColor(state C.GtkStateFlags, color *C.GdkRGBA) {
	C.gtk_widget_override_color(self.CPointer, state, color)
	return
}

/*
Sets the cursor color to use in a widget, overriding the
cursor-color and secondary-cursor-color
style properties. All other style values are left untouched.
See also gtk_widget_modify_style().

Note that the underlying properties have the #GdkColor type,
so the alpha value in @primary and @secondary will be ignored.
*/
func (self *_TraitWidget) OverrideCursor(cursor *C.GdkRGBA, secondary_cursor *C.GdkRGBA) {
	C.gtk_widget_override_cursor(self.CPointer, cursor, secondary_cursor)
	return
}

/*
Sets the font to use for a widget. All other style values are
left untouched. See gtk_widget_override_color().
*/
func (self *_TraitWidget) OverrideFont(font_desc *C.PangoFontDescription) {
	C.gtk_widget_override_font(self.CPointer, font_desc)
	return
}

/*
Sets a symbolic color for a widget.

All other style values are left untouched.
See gtk_widget_override_color() for overriding the foreground
or background color.
*/
func (self *_TraitWidget) OverrideSymbolicColor(name string, color *C.GdkRGBA) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_widget_override_symbolic_color(self.CPointer, __cgo__name, color)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

// gtk_widget_path is not generated due to deprecation attr

/*
Mark @widget as needing to recompute its expand flags. Call
this function when setting legacy expand child properties
on the child of a container.

See gtk_widget_compute_expand().
*/
func (self *_TraitWidget) QueueComputeExpand() {
	C.gtk_widget_queue_compute_expand(self.CPointer)
	return
}

/*
Equivalent to calling gtk_widget_queue_draw_area() for the
entire area of a widget.
*/
func (self *_TraitWidget) QueueDraw() {
	C.gtk_widget_queue_draw(self.CPointer)
	return
}

/*
Convenience function that calls gtk_widget_queue_draw_region() on
the region created from the given coordinates.

The region here is specified in widget coordinates.
Widget coordinates are a bit odd; for historical reasons, they are
defined as @widget->window coordinates for widgets that are not
#GTK_NO_WINDOW widgets, and are relative to @widget->allocation.x,
@widget->allocation.y for widgets that are #GTK_NO_WINDOW widgets.
*/
func (self *_TraitWidget) QueueDrawArea(x int, y int, width int, height int) {
	C.gtk_widget_queue_draw_area(self.CPointer, C.gint(x), C.gint(y), C.gint(width), C.gint(height))
	return
}

/*
Invalidates the area of @widget defined by @region by calling
gdk_window_invalidate_region() on the widget’s window and all its
child windows. Once the main loop becomes idle (after the current
batch of events has been processed, roughly), the window will
receive expose events for the union of all regions that have been
invalidated.

Normally you would only use this function in widget
implementations. You might also use it to schedule a redraw of a
#GtkDrawingArea or some portion thereof.
*/
func (self *_TraitWidget) QueueDrawRegion(region *C.cairo_region_t) {
	C.gtk_widget_queue_draw_region(self.CPointer, region)
	return
}

/*
This function is only for use in widget implementations.
Flags a widget to have its size renegotiated; should
be called when a widget for some reason has a new size request.
For example, when you change the text in a #GtkLabel, #GtkLabel
queues a resize to ensure there’s enough space for the new text.

Note that you cannot call gtk_widget_queue_resize() on a widget
from inside its implementation of the GtkWidgetClass::size_allocate
virtual method. Calls to gtk_widget_queue_resize() from inside
GtkWidgetClass::size_allocate will be silently ignored.
*/
func (self *_TraitWidget) QueueResize() {
	C.gtk_widget_queue_resize(self.CPointer)
	return
}

/*
This function works like gtk_widget_queue_resize(),
except that the widget is not invalidated.
*/
func (self *_TraitWidget) QueueResizeNoRedraw() {
	C.gtk_widget_queue_resize_no_redraw(self.CPointer)
	return
}

/*
Creates the GDK (windowing system) resources associated with a
widget.  For example, @widget->window will be created when a widget
is realized.  Normally realization happens implicitly; if you show
a widget and all its parent containers, then the widget will be
realized and mapped automatically.

Realizing a widget requires all
the widget’s parent widgets to be realized; calling
gtk_widget_realize() realizes the widget’s parents in addition to
@widget itself. If a widget is not yet inside a toplevel window
when you realize it, bad things will happen.

This function is primarily used in widget implementations, and
isn’t very useful otherwise. Many times when you think you might
need it, a better approach is to connect to a signal that will be
called after the widget is realized automatically, such as
#GtkWidget::draw. Or simply g_signal_connect () to the
#GtkWidget::realize signal.
*/
func (self *_TraitWidget) Realize() {
	C.gtk_widget_realize(self.CPointer)
	return
}

/*
Computes the intersection of a @widget’s area and @region, returning
the intersection. The result may be empty, use cairo_region_is_empty() to
check.
*/
func (self *_TraitWidget) RegionIntersect(region *C.cairo_region_t) (return__ *C.cairo_region_t) {
	return__ = C.gtk_widget_region_intersect(self.CPointer, region)
	return
}

/*
Registers a #GdkWindow with the widget and sets it up so that
the widget receives events for it. Call gtk_widget_unregister_window()
when destroying the window.

Before 3.8 you needed to call gdk_window_set_user_data() directly to set
this up. This is now deprecated and you should use gtk_widget_register_window()
instead. Old code will keep working as is, although some new features like
transparency might not work perfectly.
*/
func (self *_TraitWidget) RegisterWindow(window *C.GdkWindow) {
	C.gtk_widget_register_window(self.CPointer, window)
	return
}

/*
Removes an accelerator from @widget, previously installed with
gtk_widget_add_accelerator().
*/
func (self *_TraitWidget) RemoveAccelerator(accel_group *AccelGroup, accel_key uint, accel_mods C.GdkModifierType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_remove_accelerator(self.CPointer, (*C.GtkAccelGroup)(accel_group.CPointer), C.guint(accel_key), accel_mods)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Removes a widget from the list of mnemonic labels for
this widget. (See gtk_widget_list_mnemonic_labels()). The widget
must have previously been added to the list with
gtk_widget_add_mnemonic_label().
*/
func (self *_TraitWidget) RemoveMnemonicLabel(label *Widget) {
	C.gtk_widget_remove_mnemonic_label(self.CPointer, (*C.GtkWidget)(label.CPointer))
	return
}

/*
Removes a tick callback previously registered with
gtk_widget_add_tick_callback().
*/
func (self *_TraitWidget) RemoveTickCallback(id uint) {
	C.gtk_widget_remove_tick_callback(self.CPointer, C.guint(id))
	return
}

// gtk_widget_render_icon is not generated due to deprecation attr

// gtk_widget_render_icon_pixbuf is not generated due to deprecation attr

/*
Moves a widget from one #GtkContainer to another, handling reference
count issues to avoid destroying the widget.
*/
func (self *_TraitWidget) Reparent(new_parent *Widget) {
	C.gtk_widget_reparent(self.CPointer, (*C.GtkWidget)(new_parent.CPointer))
	return
}

// gtk_widget_reset_rc_styles is not generated due to deprecation attr

/*
Updates the style context of @widget and all descendents
by updating its widget path. #GtkContainers may want
to use this on a child when reordering it in a way that a different
style might apply to it. See also gtk_container_get_path_for_child().
*/
func (self *_TraitWidget) ResetStyle() {
	C.gtk_widget_reset_style(self.CPointer)
	return
}

/*
Very rarely-used function. This function is used to emit
an expose event on a widget. This function is not normally used
directly. The only time it is used is when propagating an expose
event to a child %NO_WINDOW widget, and that is normally done
using gtk_container_propagate_draw().

If you want to force an area of a window to be redrawn,
use gdk_window_invalidate_rect() or gdk_window_invalidate_region().
To cause the redraw to be done immediately, follow that call
with a call to gdk_window_process_updates().
*/
func (self *_TraitWidget) SendExpose(event *C.GdkEvent) (return__ int) {
	var __cgo__return__ C.gint
	__cgo__return__ = C.gtk_widget_send_expose(self.CPointer, event)
	return__ = int(__cgo__return__)
	return
}

/*
Sends the focus change @event to @widget

This function is not meant to be used by applications. The only time it
should be used is when it is necessary for a #GtkWidget to assign focus
to a widget that is semantically owned by the first widget even though
it’s not a direct child - for instance, a search entry in a floating
window similar to the quick search in #GtkTreeView.

An example of its usage is:

|[<!-- language="C" -->
  GdkEvent *fevent = gdk_event_new (GDK_FOCUS_CHANGE);

  fevent->focus_change.type = GDK_FOCUS_CHANGE;
  fevent->focus_change.in = TRUE;
  fevent->focus_change.window = gtk_widget_get_window (widget);
  if (fevent->focus_change.window != NULL)
    g_object_ref (fevent->focus_change.window);

  gtk_widget_send_focus_change (widget, fevent);

  gdk_event_free (event);
]|
*/
func (self *_TraitWidget) SendFocusChange(event *C.GdkEvent) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_send_focus_change(self.CPointer, event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Given an accelerator group, @accel_group, and an accelerator path,
@accel_path, sets up an accelerator in @accel_group so whenever the
key binding that is defined for @accel_path is pressed, @widget
will be activated.  This removes any accelerators (for any
accelerator group) installed by previous calls to
gtk_widget_set_accel_path(). Associating accelerators with
paths allows them to be modified by the user and the modifications
to be saved for future use. (See gtk_accel_map_save().)

This function is a low level function that would most likely
be used by a menu creation system like #GtkUIManager. If you
use #GtkUIManager, setting up accelerator paths will be done
automatically.

Even when you you aren’t using #GtkUIManager, if you only want to
set up accelerators on menu items gtk_menu_item_set_accel_path()
provides a somewhat more convenient interface.

Note that @accel_path string will be stored in a #GQuark. Therefore, if you
pass a static string, you can save some memory by interning it first with
g_intern_static_string().
*/
func (self *_TraitWidget) SetAccelPath(accel_path string, accel_group *AccelGroup) {
	__cgo__accel_path := (*C.gchar)(unsafe.Pointer(C.CString(accel_path)))
	C.gtk_widget_set_accel_path(self.CPointer, __cgo__accel_path, (*C.GtkAccelGroup)(accel_group.CPointer))
	C.free(unsafe.Pointer(__cgo__accel_path))
	return
}

/*
Sets the widget’s allocation.  This should not be used
directly, but from within a widget’s size_allocate method.

The allocation set should be the “adjusted” or actual
allocation. If you’re implementing a #GtkContainer, you want to use
gtk_widget_size_allocate() instead of gtk_widget_set_allocation().
The GtkWidgetClass::adjust_size_allocation virtual method adjusts the
allocation inside gtk_widget_size_allocate() to create an adjusted
allocation.
*/
func (self *_TraitWidget) SetAllocation(allocation *C.GtkAllocation) {
	C.gtk_widget_set_allocation(self.CPointer, allocation)
	return
}

/*
Sets whether the application intends to draw on the widget in
an #GtkWidget::draw handler.

This is a hint to the widget and does not affect the behavior of
the GTK+ core; many widgets ignore this flag entirely. For widgets
that do pay attention to the flag, such as #GtkEventBox and #GtkWindow,
the effect is to suppress default themed drawing of the widget's
background. (Children of the widget will still be drawn.) The application
is then entirely responsible for drawing the widget background.

Note that the background is still drawn when the widget is mapped.
*/
func (self *_TraitWidget) SetAppPaintable(app_paintable bool) {
	__cgo__app_paintable := C.gboolean(0)
	if app_paintable {
		__cgo__app_paintable = C.gboolean(1)
	}
	C.gtk_widget_set_app_paintable(self.CPointer, __cgo__app_paintable)
	return
}

/*
Specifies whether @widget can be a default widget. See
gtk_widget_grab_default() for details about the meaning of
“default”.
*/
func (self *_TraitWidget) SetCanDefault(can_default bool) {
	__cgo__can_default := C.gboolean(0)
	if can_default {
		__cgo__can_default = C.gboolean(1)
	}
	C.gtk_widget_set_can_default(self.CPointer, __cgo__can_default)
	return
}

/*
Specifies whether @widget can own the input focus. See
gtk_widget_grab_focus() for actually setting the input focus on a
widget.
*/
func (self *_TraitWidget) SetCanFocus(can_focus bool) {
	__cgo__can_focus := C.gboolean(0)
	if can_focus {
		__cgo__can_focus = C.gboolean(1)
	}
	C.gtk_widget_set_can_focus(self.CPointer, __cgo__can_focus)
	return
}

/*
Sets whether @widget should be mapped along with its when its parent
is mapped and @widget has been shown with gtk_widget_show().

The child visibility can be set for widget before it is added to
a container with gtk_widget_set_parent(), to avoid mapping
children unnecessary before immediately unmapping them. However
it will be reset to its default state of %TRUE when the widget
is removed from a container.

Note that changing the child visibility of a widget does not
queue a resize on the widget. Most of the time, the size of
a widget is computed from all visible children, whether or
not they are mapped. If this is not the case, the container
can queue a resize itself.

This function is only useful for container implementations and
never should be called by an application.
*/
func (self *_TraitWidget) SetChildVisible(is_visible bool) {
	__cgo__is_visible := C.gboolean(0)
	if is_visible {
		__cgo__is_visible = C.gboolean(1)
	}
	C.gtk_widget_set_child_visible(self.CPointer, __cgo__is_visible)
	return
}

// gtk_widget_set_composite_name is not generated due to deprecation attr

/*
Enables or disables a #GdkDevice to interact with @widget
and all its children.

It does so by descending through the #GdkWindow hierarchy
and enabling the same mask that is has for core events
(i.e. the one that gdk_window_get_events() returns).
*/
func (self *_TraitWidget) SetDeviceEnabled(device *C.GdkDevice, enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.gtk_widget_set_device_enabled(self.CPointer, device, __cgo__enabled)
	return
}

/*
Sets the device event mask (see #GdkEventMask) for a widget. The event
mask determines which events a widget will receive from @device. Keep
in mind that different widgets have different default event masks, and by
changing the event mask you may disrupt a widget’s functionality,
so be careful. This function must be called while a widget is
unrealized. Consider gtk_widget_add_device_events() for widgets that are
already realized, or if you want to preserve the existing event
mask. This function can’t be used with #GTK_NO_WINDOW widgets;
to get events on those widgets, place them inside a #GtkEventBox
and receive events on the event box.
*/
func (self *_TraitWidget) SetDeviceEvents(device *C.GdkDevice, events C.GdkEventMask) {
	C.gtk_widget_set_device_events(self.CPointer, device, events)
	return
}

/*
Sets the reading direction on a particular widget. This direction
controls the primary direction for widgets containing text,
and also the direction in which the children of a container are
packed. The ability to set the direction is present in order
so that correct localization into languages with right-to-left
reading directions can be done. Generally, applications will
let the default reading direction present, except for containers
where the containers are arranged in an order that is explicitly
visual rather than logical (such as buttons for text justification).

If the direction is set to %GTK_TEXT_DIR_NONE, then the value
set by gtk_widget_set_default_direction() will be used.
*/
func (self *_TraitWidget) SetDirection(dir C.GtkTextDirection) {
	C.gtk_widget_set_direction(self.CPointer, dir)
	return
}

/*
Widgets are double buffered by default; you can use this function
to turn off the buffering. “Double buffered” simply means that
gdk_window_begin_paint_region() and gdk_window_end_paint() are called
automatically around expose events sent to the
widget. gdk_window_begin_paint_region() diverts all drawing to a widget's
window to an offscreen buffer, and gdk_window_end_paint() draws the
buffer to the screen. The result is that users see the window
update in one smooth step, and don’t see individual graphics
primitives being rendered.

In very simple terms, double buffered widgets don’t flicker,
so you would only use this function to turn off double buffering
if you had special needs and really knew what you were doing.

Note: if you turn off double-buffering, you have to handle
expose events, since even the clearing to the background color or
pixmap will not happen automatically (as it is done in
gdk_window_begin_paint_region()).

Since 3.10 this function only works for widgets with native
windows.
*/
func (self *_TraitWidget) SetDoubleBuffered(double_buffered bool) {
	__cgo__double_buffered := C.gboolean(0)
	if double_buffered {
		__cgo__double_buffered = C.gboolean(1)
	}
	C.gtk_widget_set_double_buffered(self.CPointer, __cgo__double_buffered)
	return
}

/*
Sets the event mask (see #GdkEventMask) for a widget. The event
mask determines which events a widget will receive. Keep in mind
that different widgets have different default event masks, and by
changing the event mask you may disrupt a widget’s functionality,
so be careful. This function must be called while a widget is
unrealized. Consider gtk_widget_add_events() for widgets that are
already realized, or if you want to preserve the existing event
mask. This function can’t be used with widgets that have no window.
(See gtk_widget_get_has_window()).  To get events on those widgets,
place them inside a #GtkEventBox and receive events on the event
box.
*/
func (self *_TraitWidget) SetEvents(events int) {
	C.gtk_widget_set_events(self.CPointer, C.gint(events))
	return
}

/*
Sets the horizontal alignment of @widget.
See the #GtkWidget:halign property.
*/
func (self *_TraitWidget) SetHalign(align C.GtkAlign) {
	C.gtk_widget_set_halign(self.CPointer, align)
	return
}

/*
Sets the has-tooltip property on @widget to @has_tooltip.  See
#GtkWidget:has-tooltip for more information.
*/
func (self *_TraitWidget) SetHasTooltip(has_tooltip bool) {
	__cgo__has_tooltip := C.gboolean(0)
	if has_tooltip {
		__cgo__has_tooltip = C.gboolean(1)
	}
	C.gtk_widget_set_has_tooltip(self.CPointer, __cgo__has_tooltip)
	return
}

/*
Specifies whether @widget has a #GdkWindow of its own. Note that
all realized widgets have a non-%NULL “window” pointer
(gtk_widget_get_window() never returns a %NULL window when a widget
is realized), but for many of them it’s actually the #GdkWindow of
one of its parent widgets. Widgets that do not create a %window for
themselves in #GtkWidget::realize must announce this by
calling this function with @has_window = %FALSE.

This function should only be called by widget implementations,
and they should call it in their init() function.
*/
func (self *_TraitWidget) SetHasWindow(has_window bool) {
	__cgo__has_window := C.gboolean(0)
	if has_window {
		__cgo__has_window = C.gboolean(1)
	}
	C.gtk_widget_set_has_window(self.CPointer, __cgo__has_window)
	return
}

/*
Sets whether the widget would like any available extra horizontal
space. When a user resizes a #GtkWindow, widgets with expand=TRUE
generally receive the extra space. For example, a list or
scrollable area or document in your window would often be set to
expand.

Call this function to set the expand flag if you would like your
widget to become larger horizontally when the window has extra
room.

By default, widgets automatically expand if any of their children
want to expand. (To see if a widget will automatically expand given
its current children and state, call gtk_widget_compute_expand(). A
container can decide how the expandability of children affects the
expansion of the container by overriding the compute_expand virtual
method on #GtkWidget.).

Setting hexpand explicitly with this function will override the
automatic expand behavior.

This function forces the widget to expand or not to expand,
regardless of children.  The override occurs because
gtk_widget_set_hexpand() sets the hexpand-set property (see
gtk_widget_set_hexpand_set()) which causes the widget’s hexpand
value to be used, rather than looking at children and widget state.
*/
func (self *_TraitWidget) SetHexpand(expand bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	C.gtk_widget_set_hexpand(self.CPointer, __cgo__expand)
	return
}

/*
Sets whether the hexpand flag (see gtk_widget_get_hexpand()) will
be used.

The hexpand-set property will be set automatically when you call
gtk_widget_set_hexpand() to set hexpand, so the most likely
reason to use this function would be to unset an explicit expand
flag.

If hexpand is set, then it overrides any computed
expand value based on child widgets. If hexpand is not
set, then the expand value depends on whether any
children of the widget would like to expand.

There are few reasons to use this function, but it’s here
for completeness and consistency.
*/
func (self *_TraitWidget) SetHexpandSet(set bool) {
	__cgo__set := C.gboolean(0)
	if set {
		__cgo__set = C.gboolean(1)
	}
	C.gtk_widget_set_hexpand_set(self.CPointer, __cgo__set)
	return
}

/*
Marks the widget as being realized.

This function should only ever be called in a derived widget's
“map” or “unmap” implementation.
*/
func (self *_TraitWidget) SetMapped(mapped bool) {
	__cgo__mapped := C.gboolean(0)
	if mapped {
		__cgo__mapped = C.gboolean(1)
	}
	C.gtk_widget_set_mapped(self.CPointer, __cgo__mapped)
	return
}

/*
Sets the bottom margin of @widget.
See the #GtkWidget:margin-bottom property.
*/
func (self *_TraitWidget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(self.CPointer, C.gint(margin))
	return
}

/*
Sets the end margin of @widget.
See the #GtkWidget:margin-end property.
*/
func (self *_TraitWidget) SetMarginEnd(margin int) {
	C.gtk_widget_set_margin_end(self.CPointer, C.gint(margin))
	return
}

// gtk_widget_set_margin_left is not generated due to deprecation attr

// gtk_widget_set_margin_right is not generated due to deprecation attr

/*
Sets the start margin of @widget.
See the #GtkWidget:margin-start property.
*/
func (self *_TraitWidget) SetMarginStart(margin int) {
	C.gtk_widget_set_margin_start(self.CPointer, C.gint(margin))
	return
}

/*
Sets the top margin of @widget.
See the #GtkWidget:margin-top property.
*/
func (self *_TraitWidget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(self.CPointer, C.gint(margin))
	return
}

/*
Widgets can be named, which allows you to refer to them from a
CSS file. You can apply a style to widgets with a particular name
in the CSS file. See the documentation for the CSS syntax (on the
same page as the docs for #GtkStyleContext).

Note that the CSS syntax has certain special characters to delimit
and represent elements in a selector (period, #, >, *...), so using
these will make your widget impossible to match by name. Any combination
of alphanumeric symbols, dashes and underscores will suffice.
*/
func (self *_TraitWidget) SetName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_widget_set_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Sets the #GtkWidget:no-show-all property, which determines whether
calls to gtk_widget_show_all() will affect this widget.

This is mostly for use in constructing widget hierarchies with externally
controlled visibility, see #GtkUIManager.
*/
func (self *_TraitWidget) SetNoShowAll(no_show_all bool) {
	__cgo__no_show_all := C.gboolean(0)
	if no_show_all {
		__cgo__no_show_all = C.gboolean(1)
	}
	C.gtk_widget_set_no_show_all(self.CPointer, __cgo__no_show_all)
	return
}

/*
Request the @widget to be rendered partially transparent,
with opacity 0 being fully transparent and 1 fully opaque. (Opacity values
are clamped to the [0,1] range.).
This works on both toplevel widget, and child widgets, although there
are some limitations:

For toplevel widgets this depends on the capabilities of the windowing
system. On X11 this has any effect only on X screens with a compositing manager
running. See gtk_widget_is_composited(). On Windows it should work
always, although setting a window’s opacity after the window has been
shown causes it to flicker once on Windows.

For child widgets it doesn’t work if any affected widget has a native window, or
disables double buffering.
*/
func (self *_TraitWidget) SetOpacity(opacity float64) {
	C.gtk_widget_set_opacity(self.CPointer, C.double(opacity))
	return
}

/*
This function is useful only when implementing subclasses of
#GtkContainer.
Sets the container as the parent of @widget, and takes care of
some details such as updating the state and style of the child
to reflect its new location. The opposite function is
gtk_widget_unparent().
*/
func (self *_TraitWidget) SetParent(parent *Widget) {
	C.gtk_widget_set_parent(self.CPointer, (*C.GtkWidget)(parent.CPointer))
	return
}

/*
Sets a non default parent window for @widget.

For #GtkWindow classes, setting a @parent_window effects whether
the window is a toplevel window or can be embedded into other
widgets.

For #GtkWindow classes, this needs to be called before the
window is realized.
*/
func (self *_TraitWidget) SetParentWindow(parent_window *C.GdkWindow) {
	C.gtk_widget_set_parent_window(self.CPointer, parent_window)
	return
}

/*
Marks the widget as being realized. This function must only be
called after all #GdkWindows for the @widget have been created
and registered.

This function should only ever be called in a derived widget's
“realize” or “unrealize” implementation.
*/
func (self *_TraitWidget) SetRealized(realized bool) {
	__cgo__realized := C.gboolean(0)
	if realized {
		__cgo__realized = C.gboolean(1)
	}
	C.gtk_widget_set_realized(self.CPointer, __cgo__realized)
	return
}

/*
Specifies whether @widget will be treated as the default widget
within its toplevel when it has the focus, even if another widget
is the default.

See gtk_widget_grab_default() for details about the meaning of
“default”.
*/
func (self *_TraitWidget) SetReceivesDefault(receives_default bool) {
	__cgo__receives_default := C.gboolean(0)
	if receives_default {
		__cgo__receives_default = C.gboolean(1)
	}
	C.gtk_widget_set_receives_default(self.CPointer, __cgo__receives_default)
	return
}

/*
Sets whether the entire widget is queued for drawing when its size
allocation changes. By default, this setting is %TRUE and
the entire widget is redrawn on every size change. If your widget
leaves the upper left unchanged when made bigger, turning this
setting off will improve performance.

Note that for widgets where gtk_widget_get_has_window() is %FALSE
setting this flag to %FALSE turns off all allocation on resizing:
the widget will not even redraw if its position changes; this is to
allow containers that don’t draw anything to avoid excess
invalidations. If you set this flag on a widget with no window that
does draw on @widget->window, you are
responsible for invalidating both the old and new allocation of the
widget when the widget is moved and responsible for invalidating
regions newly when the widget increases size.
*/
func (self *_TraitWidget) SetRedrawOnAllocate(redraw_on_allocate bool) {
	__cgo__redraw_on_allocate := C.gboolean(0)
	if redraw_on_allocate {
		__cgo__redraw_on_allocate = C.gboolean(1)
	}
	C.gtk_widget_set_redraw_on_allocate(self.CPointer, __cgo__redraw_on_allocate)
	return
}

/*
Sets the sensitivity of a widget. A widget is sensitive if the user
can interact with it. Insensitive widgets are “grayed out” and the
user can’t interact with them. Insensitive widgets are known as
“inactive”, “disabled”, or “ghosted” in some other toolkits.
*/
func (self *_TraitWidget) SetSensitive(sensitive bool) {
	__cgo__sensitive := C.gboolean(0)
	if sensitive {
		__cgo__sensitive = C.gboolean(1)
	}
	C.gtk_widget_set_sensitive(self.CPointer, __cgo__sensitive)
	return
}

/*
Sets the minimum size of a widget; that is, the widget’s size
request will be at least @width by @height. You can use this
function to force a widget to be larger than it normally would be.

In most cases, gtk_window_set_default_size() is a better choice for
toplevel windows than this function; setting the default size will
still allow users to shrink the window. Setting the size request
will force them to leave the window at least as large as the size
request. When dealing with window sizes,
gtk_window_set_geometry_hints() can be a useful function as well.

Note the inherent danger of setting any fixed size - themes,
translations into other languages, different fonts, and user action
can all change the appropriate size for a given widget. So, it's
basically impossible to hardcode a size that will always be
correct.

The size request of a widget is the smallest size a widget can
accept while still functioning well and drawing itself correctly.
However in some strange cases a widget may be allocated less than
its requested size, and in many cases a widget may be allocated more
space than it requested.

If the size request in a given direction is -1 (unset), then
the “natural” size request of the widget will be used instead.

The size request set here does not include any margin from the
#GtkWidget properties margin-left, margin-right, margin-top, and
margin-bottom, but it does include pretty much all other padding
or border properties set by any subclass of #GtkWidget.
*/
func (self *_TraitWidget) SetSizeRequest(width int, height int) {
	C.gtk_widget_set_size_request(self.CPointer, C.gint(width), C.gint(height))
	return
}

// gtk_widget_set_state is not generated due to deprecation attr

/*
This function is for use in widget implementations. Turns on flag
values in the current widget state (insensitive, prelighted, etc.).

This function accepts the values %GTK_STATE_FLAG_DIR_LTR and
%GTK_STATE_FLAG_DIR_RTL but ignores them. If you want to set the widget's
direction, use gtk_widget_set_direction().

It is worth mentioning that any other state than %GTK_STATE_FLAG_INSENSITIVE,
will be propagated down to all non-internal children if @widget is a
#GtkContainer, while %GTK_STATE_FLAG_INSENSITIVE itself will be propagated
down to all #GtkContainer children by different means than turning on the
state flag down the hierarchy, both gtk_widget_get_state_flags() and
gtk_widget_is_sensitive() will make use of these.
*/
func (self *_TraitWidget) SetStateFlags(flags C.GtkStateFlags, clear bool) {
	__cgo__clear := C.gboolean(0)
	if clear {
		__cgo__clear = C.gboolean(1)
	}
	C.gtk_widget_set_state_flags(self.CPointer, flags, __cgo__clear)
	return
}

// gtk_widget_set_style is not generated due to deprecation attr

/*
Enables or disables multiple pointer awareness. If this setting is %TRUE,
@widget will start receiving multiple, per device enter/leave events. Note
that if custom #GdkWindows are created in #GtkWidget::realize,
gdk_window_set_support_multidevice() will have to be called manually on them.
*/
func (self *_TraitWidget) SetSupportMultidevice(support_multidevice bool) {
	__cgo__support_multidevice := C.gboolean(0)
	if support_multidevice {
		__cgo__support_multidevice = C.gboolean(1)
	}
	C.gtk_widget_set_support_multidevice(self.CPointer, __cgo__support_multidevice)
	return
}

/*
Sets @markup as the contents of the tooltip, which is marked up with
 the [Pango text markup language][PangoMarkupFormat].

This function will take care of setting #GtkWidget:has-tooltip to %TRUE
and of the default handler for the #GtkWidget::query-tooltip signal.

See also the #GtkWidget:tooltip-markup property and
gtk_tooltip_set_markup().
*/
func (self *_TraitWidget) SetTooltipMarkup(markup string) {
	__cgo__markup := (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	C.gtk_widget_set_tooltip_markup(self.CPointer, __cgo__markup)
	C.free(unsafe.Pointer(__cgo__markup))
	return
}

/*
Sets @text as the contents of the tooltip. This function will take
care of setting #GtkWidget:has-tooltip to %TRUE and of the default
handler for the #GtkWidget::query-tooltip signal.

See also the #GtkWidget:tooltip-text property and gtk_tooltip_set_text().
*/
func (self *_TraitWidget) SetTooltipText(text string) {
	__cgo__text := (*C.gchar)(unsafe.Pointer(C.CString(text)))
	C.gtk_widget_set_tooltip_text(self.CPointer, __cgo__text)
	C.free(unsafe.Pointer(__cgo__text))
	return
}

/*
Replaces the default, usually yellow, window used for displaying
tooltips with @custom_window. GTK+ will take care of showing and
hiding @custom_window at the right moment, to behave likewise as
the default tooltip window. If @custom_window is %NULL, the default
tooltip window will be used.

If the custom window should have the default theming it needs to
have the name “gtk-tooltip”, see gtk_widget_set_name().
*/
func (self *_TraitWidget) SetTooltipWindow(custom_window *Window) {
	C.gtk_widget_set_tooltip_window(self.CPointer, (*C.GtkWindow)(custom_window.CPointer))
	return
}

/*
Sets the vertical alignment of @widget.
See the #GtkWidget:valign property.
*/
func (self *_TraitWidget) SetValign(align C.GtkAlign) {
	C.gtk_widget_set_valign(self.CPointer, align)
	return
}

/*
Sets whether the widget would like any available extra vertical
space.

See gtk_widget_set_hexpand() for more detail.
*/
func (self *_TraitWidget) SetVexpand(expand bool) {
	__cgo__expand := C.gboolean(0)
	if expand {
		__cgo__expand = C.gboolean(1)
	}
	C.gtk_widget_set_vexpand(self.CPointer, __cgo__expand)
	return
}

/*
Sets whether the vexpand flag (see gtk_widget_get_vexpand()) will
be used.

See gtk_widget_set_hexpand_set() for more detail.
*/
func (self *_TraitWidget) SetVexpandSet(set bool) {
	__cgo__set := C.gboolean(0)
	if set {
		__cgo__set = C.gboolean(1)
	}
	C.gtk_widget_set_vexpand_set(self.CPointer, __cgo__set)
	return
}

/*
Sets the visibility state of @widget. Note that setting this to
%TRUE doesn’t mean the widget is actually viewable, see
gtk_widget_get_visible().

This function simply calls gtk_widget_show() or gtk_widget_hide()
but is nicer to use when the visibility of the widget depends on
some condition.
*/
func (self *_TraitWidget) SetVisible(visible bool) {
	__cgo__visible := C.gboolean(0)
	if visible {
		__cgo__visible = C.gboolean(1)
	}
	C.gtk_widget_set_visible(self.CPointer, __cgo__visible)
	return
}

/*
Sets the visual that should be used for by widget and its children for
creating #GdkWindows. The visual must be on the same #GdkScreen as
returned by gtk_widget_get_screen(), so handling the
#GtkWidget::screen-changed signal is necessary.

Setting a new @visual will not cause @widget to recreate its windows,
so you should call this function before @widget is realized.
*/
func (self *_TraitWidget) SetVisual(visual *C.GdkVisual) {
	C.gtk_widget_set_visual(self.CPointer, visual)
	return
}

/*
Sets a widget’s window. This function should only be used in a
widget’s #GtkWidget::realize implementation. The %window passed is
usually either new window created with gdk_window_new(), or the
window of its parent widget as returned by
gtk_widget_get_parent_window().

Widgets must indicate whether they will create their own #GdkWindow
by calling gtk_widget_set_has_window(). This is usually done in the
widget’s init() function.

Note that this function does not add any reference to @window.
*/
func (self *_TraitWidget) SetWindow(window *C.GdkWindow) {
	C.gtk_widget_set_window(self.CPointer, window)
	return
}

/*
Sets a shape for this widget’s GDK window. This allows for
transparent windows etc., see gdk_window_shape_combine_region()
for more information.
*/
func (self *_TraitWidget) ShapeCombineRegion(region *C.cairo_region_t) {
	C.gtk_widget_shape_combine_region(self.CPointer, region)
	return
}

/*
Flags a widget to be displayed. Any widget that isn’t shown will
not appear on the screen. If you want to show all the widgets in a
container, it’s easier to call gtk_widget_show_all() on the
container, instead of individually showing the widgets.

Remember that you have to show the containers containing a widget,
in addition to the widget itself, before it will appear onscreen.

When a toplevel container is shown, it is immediately realized and
mapped; other shown widgets are realized and mapped when their
toplevel container is realized and mapped.
*/
func (self *_TraitWidget) Show() {
	C.gtk_widget_show(self.CPointer)
	return
}

/*
Recursively shows a widget, and any child widgets (if the widget is
a container).
*/
func (self *_TraitWidget) ShowAll() {
	C.gtk_widget_show_all(self.CPointer)
	return
}

/*
Shows a widget. If the widget is an unmapped toplevel widget
(i.e. a #GtkWindow that has not yet been shown), enter the main
loop and wait for the window to actually be mapped. Be careful;
because the main loop is running, anything can happen during
this function.
*/
func (self *_TraitWidget) ShowNow() {
	C.gtk_widget_show_now(self.CPointer)
	return
}

/*
This function is only used by #GtkContainer subclasses, to assign a size
and position to their child widgets.

In this function, the allocation may be adjusted. It will be forced
to a 1x1 minimum size, and the adjust_size_allocation virtual
method on the child will be used to adjust the allocation. Standard
adjustments include removing the widget’s margins, and applying the
widget’s #GtkWidget:halign and #GtkWidget:valign properties.

For baseline support in containers you need to use gtk_widget_size_allocate_with_baseline()
instead.
*/
func (self *_TraitWidget) SizeAllocate(allocation *C.GtkAllocation) {
	C.gtk_widget_size_allocate(self.CPointer, allocation)
	return
}

/*
This function is only used by #GtkContainer subclasses, to assign a size,
position and (optionally) baseline to their child widgets.

In this function, the allocation and baseline may be adjusted. It
will be forced to a 1x1 minimum size, and the
adjust_size_allocation virtual and adjust_baseline_allocation
methods on the child will be used to adjust the allocation and
baseline. Standard adjustments include removing the widget's
margins, and applying the widget’s #GtkWidget:halign and
#GtkWidget:valign properties.

If the child widget does not have a valign of %GTK_ALIGN_BASELINE the
baseline argument is ignored and -1 is used instead.
*/
func (self *_TraitWidget) SizeAllocateWithBaseline(allocation *C.GtkAllocation, baseline int) {
	C.gtk_widget_size_allocate_with_baseline(self.CPointer, allocation, C.gint(baseline))
	return
}

// gtk_widget_size_request is not generated due to deprecation attr

// gtk_widget_style_attach is not generated due to deprecation attr

// gtk_widget_style_get is not generated due to varargs

/*
Gets the value of a style property of @widget.
*/
func (self *_TraitWidget) StyleGetProperty(property_name string, value *C.GValue) {
	__cgo__property_name := (*C.gchar)(unsafe.Pointer(C.CString(property_name)))
	C.gtk_widget_style_get_property(self.CPointer, __cgo__property_name, value)
	C.free(unsafe.Pointer(__cgo__property_name))
	return
}

// gtk_widget_style_get_valist is not generated due to varargs

/*
Reverts the effect of a previous call to gtk_widget_freeze_child_notify().
This causes all queued #GtkWidget::child-notify signals on @widget to be
emitted.
*/
func (self *_TraitWidget) ThawChildNotify() {
	C.gtk_widget_thaw_child_notify(self.CPointer)
	return
}

/*
Translate coordinates relative to @src_widget’s allocation to coordinates
relative to @dest_widget’s allocations. In order to perform this
operation, both widgets must be realized, and must share a common
toplevel.
*/
func (self *_TraitWidget) TranslateCoordinates(dest_widget *Widget, src_x int, src_y int) (dest_x int, dest_y int, return__ bool) {
	var __cgo__dest_x C.gint
	var __cgo__dest_y C.gint
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_widget_translate_coordinates(self.CPointer, (*C.GtkWidget)(dest_widget.CPointer), C.gint(src_x), C.gint(src_y), &__cgo__dest_x, &__cgo__dest_y)
	dest_x = int(__cgo__dest_x)
	dest_y = int(__cgo__dest_y)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Triggers a tooltip query on the display where the toplevel of @widget
is located. See gtk_tooltip_trigger_tooltip_query() for more
information.
*/
func (self *_TraitWidget) TriggerTooltipQuery() {
	C.gtk_widget_trigger_tooltip_query(self.CPointer)
	return
}

/*
This function is only for use in widget implementations. Causes
a widget to be unmapped if it’s currently mapped.
*/
func (self *_TraitWidget) Unmap() {
	C.gtk_widget_unmap(self.CPointer)
	return
}

/*
This function is only for use in widget implementations.
Should be called by implementations of the remove method
on #GtkContainer, to dissociate a child from the container.
*/
func (self *_TraitWidget) Unparent() {
	C.gtk_widget_unparent(self.CPointer)
	return
}

/*
This function is only useful in widget implementations.
Causes a widget to be unrealized (frees all GDK resources
associated with the widget, such as @widget->window).
*/
func (self *_TraitWidget) Unrealize() {
	C.gtk_widget_unrealize(self.CPointer)
	return
}

/*
Unregisters a #GdkWindow from the widget that was previously set up with
gtk_widget_register_window(). You need to call this when the window is
no longer used by the widget, such as when you destroy it.
*/
func (self *_TraitWidget) UnregisterWindow(window *C.GdkWindow) {
	C.gtk_widget_unregister_window(self.CPointer, window)
	return
}

/*
This function is for use in widget implementations. Turns off flag
values for the current widget state (insensitive, prelighted, etc.).
See gtk_widget_set_state_flags().
*/
func (self *_TraitWidget) UnsetStateFlags(flags C.GtkStateFlags) {
	C.gtk_widget_unset_state_flags(self.CPointer, flags)
	return
}

type _TraitWindow struct{ CPointer *C.GtkWindow }

/*
Activates the default widget for the window, unless the current
focused widget has been configured to receive the default action
(see gtk_widget_set_receives_default()), in which case the
focused widget is activated.
*/
func (self *_TraitWindow) ActivateDefault() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_activate_default(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Activates the current focused widget within the window.
*/
func (self *_TraitWindow) ActivateFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_activate_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Activates mnemonics and accelerators for this #GtkWindow. This is normally
called by the default ::key_press_event handler for toplevel windows,
however in some cases it may be useful to call this directly when
overriding the standard key handling for a toplevel window.
*/
func (self *_TraitWindow) ActivateKey(event *C.GdkEventKey) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_activate_key(self.CPointer, event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Associate @accel_group with @window, such that calling
gtk_accel_groups_activate() on @window will activate accelerators
in @accel_group.
*/
func (self *_TraitWindow) AddAccelGroup(accel_group *AccelGroup) {
	C.gtk_window_add_accel_group(self.CPointer, (*C.GtkAccelGroup)(accel_group.CPointer))
	return
}

/*
Adds a mnemonic to this window.
*/
func (self *_TraitWindow) AddMnemonic(keyval uint, target *Widget) {
	C.gtk_window_add_mnemonic(self.CPointer, C.guint(keyval), (*C.GtkWidget)(target.CPointer))
	return
}

/*
Starts moving a window. This function is used if an application has
window movement grips. When GDK can support it, the window movement
will be done using the standard mechanism for the
[window manager][gtk-X11-arch] or windowing
system. Otherwise, GDK will try to emulate window movement,
potentially not all that well, depending on the windowing system.
*/
func (self *_TraitWindow) BeginMoveDrag(button int, root_x int, root_y int, timestamp uint32) {
	C.gtk_window_begin_move_drag(self.CPointer, C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
	return
}

/*
Starts resizing a window. This function is used if an application
has window resizing controls. When GDK can support it, the resize
will be done using the standard mechanism for the
[window manager][gtk-X11-arch] or windowing
system. Otherwise, GDK will try to emulate window resizing,
potentially not all that well, depending on the windowing system.
*/
func (self *_TraitWindow) BeginResizeDrag(edge C.GdkWindowEdge, button int, root_x int, root_y int, timestamp uint32) {
	C.gtk_window_begin_resize_drag(self.CPointer, edge, C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
	return
}

/*
Requests that the window is closed, similar to what happens
when a window manager close button is clicked.

This function can be used with close buttons in custom
titlebars.
*/
func (self *_TraitWindow) Close() {
	C.gtk_window_close(self.CPointer)
	return
}

/*
Asks to deiconify (i.e. unminimize) the specified @window. Note
that you shouldn’t assume the window is definitely deiconified
afterward, because other entities (e.g. the user or
[window manager][gtk-X11-arch])) could iconify it
again before your code which assumes deiconification gets to run.

You can track iconification via the “window-state-event” signal
on #GtkWidget.
*/
func (self *_TraitWindow) Deiconify() {
	C.gtk_window_deiconify(self.CPointer)
	return
}

/*
Asks to place @window in the fullscreen state. Note that you
shouldn’t assume the window is definitely full screen afterward,
because other entities (e.g. the user or
[window manager][gtk-X11-arch]) could unfullscreen it
again, and not all window managers honor requests to fullscreen
windows. But normally the window will end up fullscreen. Just
don’t write code that crashes if not.

You can track the fullscreen state via the “window-state-event” signal
on #GtkWidget.
*/
func (self *_TraitWindow) Fullscreen() {
	C.gtk_window_fullscreen(self.CPointer)
	return
}

/*
Gets the value set by gtk_window_set_accept_focus().
*/
func (self *_TraitWindow) GetAcceptFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_accept_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the #GtkApplication associated with the window (if any).
*/
func (self *_TraitWindow) GetApplication() (return__ *Application) {
	var __cgo__return__ *C.GtkApplication
	__cgo__return__ = C.gtk_window_get_application(self.CPointer)
	return__ = NewApplicationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Fetches the attach widget for this window. See
gtk_window_set_attached_to().
*/
func (self *_TraitWindow) GetAttachedTo() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_window_get_attached_to(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether the window has been set to have decorations
such as a title bar via gtk_window_set_decorated().
*/
func (self *_TraitWindow) GetDecorated() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_decorated(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the default size of the window. A value of -1 for the width or
height indicates that a default size has not been explicitly set
for that dimension, so the “natural” size of the window will be
used.
*/
func (self *_TraitWindow) GetDefaultSize() (width int, height int) {
	var __cgo__width C.gint
	var __cgo__height C.gint
	C.gtk_window_get_default_size(self.CPointer, &__cgo__width, &__cgo__height)
	width = int(__cgo__width)
	height = int(__cgo__height)
	return
}

/*
Returns the default widget for @window. See gtk_window_set_default()
for more details.
*/
func (self *_TraitWindow) GetDefaultWidget() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_window_get_default_widget(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns whether the window has been set to have a close button
via gtk_window_set_deletable().
*/
func (self *_TraitWindow) GetDeletable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_deletable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the window will be destroyed with its transient parent. See
gtk_window_set_destroy_with_parent ().
*/
func (self *_TraitWindow) GetDestroyWithParent() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_destroy_with_parent(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the current focused widget within the window.
Note that this is the widget that would have the focus
if the toplevel window focused; if the toplevel window
is not focused then  `gtk_widget_has_focus (widget)` will
not be %TRUE for the widget.
*/
func (self *_TraitWindow) GetFocus() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_window_get_focus(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the value set by gtk_window_set_focus_on_map().
*/
func (self *_TraitWindow) GetFocusOnMap() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_focus_on_map(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of the #GtkWindow:focus-visible property.
*/
func (self *_TraitWindow) GetFocusVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_focus_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_window_set_gravity().
*/
func (self *_TraitWindow) GetGravity() (return__ C.GdkGravity) {
	return__ = C.gtk_window_get_gravity(self.CPointer)
	return
}

/*
Returns the group for @window or the default group, if
@window is %NULL or if @window does not have an explicit
window group.
*/
func (self *_TraitWindow) GetGroup() (return__ *WindowGroup) {
	var __cgo__return__ *C.GtkWindowGroup
	__cgo__return__ = C.gtk_window_get_group(self.CPointer)
	return__ = NewWindowGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Determines whether the window may have a resize grip.
*/
func (self *_TraitWindow) GetHasResizeGrip() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_has_resize_grip(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the window has requested to have its titlebar hidden
when maximized. See gtk_window_set_hide_titlebar_when_maximized ().
*/
func (self *_TraitWindow) GetHideTitlebarWhenMaximized() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_hide_titlebar_when_maximized(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_window_set_icon() (or if you've
called gtk_window_set_icon_list(), gets the first icon in
the icon list).
*/
func (self *_TraitWindow) GetIcon() (return__ *C.GdkPixbuf) {
	return__ = C.gtk_window_get_icon(self.CPointer)
	return
}

/*
Retrieves the list of icons set by gtk_window_set_icon_list().
The list is copied, but the reference count on each
member won’t be incremented.
*/
func (self *_TraitWindow) GetIconList() (return__ *C.GList) {
	return__ = C.gtk_window_get_icon_list(self.CPointer)
	return
}

/*
Returns the name of the themed icon for the window,
see gtk_window_set_icon_name().
*/
func (self *_TraitWindow) GetIconName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_window_get_icon_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the mnemonic modifier for this window. See
gtk_window_set_mnemonic_modifier().
*/
func (self *_TraitWindow) GetMnemonicModifier() (return__ C.GdkModifierType) {
	return__ = C.gtk_window_get_mnemonic_modifier(self.CPointer)
	return
}

/*
Gets the value of the #GtkWindow:mnemonics-visible property.
*/
func (self *_TraitWindow) GetMnemonicsVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_mnemonics_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the window is modal. See gtk_window_set_modal().
*/
func (self *_TraitWindow) GetModal() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_modal(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

// gtk_window_get_opacity is not generated due to deprecation attr

/*
This function returns the position you need to pass to
gtk_window_move() to keep @window in its current position.
This means that the meaning of the returned value varies with
window gravity. See gtk_window_move() for more details.

If you haven’t changed the window gravity, its gravity will be
#GDK_GRAVITY_NORTH_WEST. This means that gtk_window_get_position()
gets the position of the top-left corner of the window manager
frame for the window. gtk_window_move() sets the position of this
same top-left corner.

gtk_window_get_position() is not 100% reliable because the X Window System
does not specify a way to obtain the geometry of the
decorations placed on a window by the window manager.
Thus GTK+ is using a “best guess” that works with most
window managers.

Moreover, nearly all window managers are historically broken with
respect to their handling of window gravity. So moving a window to
its current position as returned by gtk_window_get_position() tends
to result in moving the window slightly. Window managers are
slowly getting better over time.

If a window has gravity #GDK_GRAVITY_STATIC the window manager
frame is not relevant, and thus gtk_window_get_position() will
always produce accurate results. However you can’t use static
gravity to do things like place a window in a corner of the screen,
because static gravity ignores the window manager decorations.

If you are saving and restoring your application’s window
positions, you should know that it’s impossible for applications to
do this without getting it somewhat wrong because applications do
not have sufficient knowledge of window manager state. The Correct
Mechanism is to support the session management protocol (see the
“GnomeClient” object in the GNOME libraries for example) and allow
the window manager to save your window sizes and positions.
*/
func (self *_TraitWindow) GetPosition() (root_x int, root_y int) {
	var __cgo__root_x C.gint
	var __cgo__root_y C.gint
	C.gtk_window_get_position(self.CPointer, &__cgo__root_x, &__cgo__root_y)
	root_x = int(__cgo__root_x)
	root_y = int(__cgo__root_y)
	return
}

/*
Gets the value set by gtk_window_set_resizable().
*/
func (self *_TraitWindow) GetResizable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_resizable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
If a window has a resize grip, this will retrieve the grip
position, width and height into the specified #GdkRectangle.
*/
func (self *_TraitWindow) GetResizeGripArea() (rect C.GdkRectangle, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_resize_grip_area(self.CPointer, &rect)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the role of the window. See gtk_window_set_role() for
further explanation.
*/
func (self *_TraitWindow) GetRole() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_window_get_role(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Returns the #GdkScreen associated with @window.
*/
func (self *_TraitWindow) GetScreen() (return__ *C.GdkScreen) {
	return__ = C.gtk_window_get_screen(self.CPointer)
	return
}

/*
Obtains the current size of @window. If @window is not onscreen,
it returns the size GTK+ will suggest to the
[window manager][gtk-X11-arch]
for the initial window
size (but this is not reliably the same as the size the window
manager will actually select). The size obtained by
gtk_window_get_size() is the last size received in a
#GdkEventConfigure, that is, GTK+ uses its locally-stored size,
rather than querying the X server for the size. As a result, if you
call gtk_window_resize() then immediately call
gtk_window_get_size(), the size won’t have taken effect yet. After
the window manager processes the resize request, GTK+ receives
notification that the size has changed via a configure event, and
the size of the window gets updated.

Note 1: Nearly any use of this function creates a race condition,
because the size of the window may change between the time that you
get the size and the time that you perform some action assuming
that size is the current size. To avoid race conditions, connect to
“configure-event” on the window and adjust your size-dependent
state to match the size delivered in the #GdkEventConfigure.

Note 2: The returned size does not include the
size of the window manager decorations (aka the window frame or
border). Those are not drawn by GTK+ and GTK+ has no reliable
method of determining their size.

Note 3: If you are getting a window size in order to position
the window onscreen, there may be a better way. The preferred
way is to simply set the window’s semantic type with
gtk_window_set_type_hint(), which allows the window manager to
e.g. center dialogs. Also, if you set the transient parent of
dialogs with gtk_window_set_transient_for() window managers
will often center the dialog over its parent window. It's
much preferred to let the window manager handle these
things rather than doing it yourself, because all apps will
behave consistently and according to user prefs if the window
manager handles it. Also, the window manager can take the size
of the window decorations/border into account, while your
application cannot.

In any case, if you insist on application-specified window
positioning, there’s still a better way than
doing it yourself - gtk_window_set_position() will frequently
handle the details for you.
*/
func (self *_TraitWindow) GetSize() (width int, height int) {
	var __cgo__width C.gint
	var __cgo__height C.gint
	C.gtk_window_get_size(self.CPointer, &__cgo__width, &__cgo__height)
	width = int(__cgo__width)
	height = int(__cgo__height)
	return
}

/*
Gets the value set by gtk_window_set_skip_pager_hint().
*/
func (self *_TraitWindow) GetSkipPagerHint() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_skip_pager_hint(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value set by gtk_window_set_skip_taskbar_hint()
*/
func (self *_TraitWindow) GetSkipTaskbarHint() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_skip_taskbar_hint(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the title of the window. See gtk_window_set_title().
*/
func (self *_TraitWindow) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.gtk_window_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Fetches the transient parent for this window. See
gtk_window_set_transient_for().
*/
func (self *_TraitWindow) GetTransientFor() (return__ *Window) {
	var __cgo__return__ *C.GtkWindow
	__cgo__return__ = C.gtk_window_get_transient_for(self.CPointer)
	return__ = NewWindowFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the type hint for this window. See gtk_window_set_type_hint().
*/
func (self *_TraitWindow) GetTypeHint() (return__ C.GdkWindowTypeHint) {
	return__ = C.gtk_window_get_type_hint(self.CPointer)
	return
}

/*
Gets the value set by gtk_window_set_urgency_hint()
*/
func (self *_TraitWindow) GetUrgencyHint() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_get_urgency_hint(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the type of the window. See #GtkWindowType.
*/
func (self *_TraitWindow) GetWindowType() (return__ C.GtkWindowType) {
	return__ = C.gtk_window_get_window_type(self.CPointer)
	return
}

/*
Returns whether @window has an explicit window group.
*/
func (self *_TraitWindow) HasGroup() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_has_group(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns whether the input focus is within this GtkWindow.
For real toplevel windows, this is identical to gtk_window_is_active(),
but for embedded windows, like #GtkPlug, the results will differ.
*/
func (self *_TraitWindow) HasToplevelFocus() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_has_toplevel_focus(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Asks to iconify (i.e. minimize) the specified @window. Note that
you shouldn’t assume the window is definitely iconified afterward,
because other entities (e.g. the user or
[window manager][gtk-X11-arch]) could deiconify it
again, or there may not be a window manager in which case
iconification isn’t possible, etc. But normally the window will end
up iconified. Just don’t write code that crashes if not.

It’s permitted to call this function before showing a window,
in which case the window will be iconified before it ever appears
onscreen.

You can track iconification via the “window-state-event” signal
on #GtkWidget.
*/
func (self *_TraitWindow) Iconify() {
	C.gtk_window_iconify(self.CPointer)
	return
}

/*
Returns whether the window is part of the current active toplevel.
(That is, the toplevel window receiving keystrokes.)
The return value is %TRUE if the window is active toplevel
itself, but also if it is, say, a #GtkPlug embedded in the active toplevel.
You might use this function if you wanted to draw a widget
differently in an active window from a widget in an inactive window.
See gtk_window_has_toplevel_focus()
*/
func (self *_TraitWindow) IsActive() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_is_active(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Retrieves the current maximized state of @window.

Note that since maximization is ultimately handled by the window
manager and happens asynchronously to an application request, you
shouldn’t assume the return value of this function changing
immediately (or at all), as an effect of calling
gtk_window_maximize() or gtk_window_unmaximize().
*/
func (self *_TraitWindow) IsMaximized() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_is_maximized(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Asks to maximize @window, so that it becomes full-screen. Note that
you shouldn’t assume the window is definitely maximized afterward,
because other entities (e.g. the user or
[window manager][gtk-X11-arch]) could unmaximize it
again, and not all window managers support maximization. But
normally the window will end up maximized. Just don’t write code
that crashes if not.

It’s permitted to call this function before showing a window,
in which case the window will be maximized when it appears onscreen
initially.

You can track maximization via the “window-state-event” signal
on #GtkWidget, or by listening to notifications on the
#GtkWindow:is-maximized property.
*/
func (self *_TraitWindow) Maximize() {
	C.gtk_window_maximize(self.CPointer)
	return
}

/*
Activates the targets associated with the mnemonic.
*/
func (self *_TraitWindow) MnemonicActivate(keyval uint, modifier C.GdkModifierType) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_mnemonic_activate(self.CPointer, C.guint(keyval), modifier)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Asks the [window manager][gtk-X11-arch] to move
@window to the given position.  Window managers are free to ignore
this; most window managers ignore requests for initial window
positions (instead using a user-defined placement algorithm) and
honor requests after the window has already been shown.

Note: the position is the position of the gravity-determined
reference point for the window. The gravity determines two things:
first, the location of the reference point in root window
coordinates; and second, which point on the window is positioned at
the reference point.

By default the gravity is #GDK_GRAVITY_NORTH_WEST, so the reference
point is simply the @x, @y supplied to gtk_window_move(). The
top-left corner of the window decorations (aka window frame or
border) will be placed at @x, @y.  Therefore, to position a window
at the top left of the screen, you want to use the default gravity
(which is #GDK_GRAVITY_NORTH_WEST) and move the window to 0,0.

To position a window at the bottom right corner of the screen, you
would set #GDK_GRAVITY_SOUTH_EAST, which means that the reference
point is at @x + the window width and @y + the window height, and
the bottom-right corner of the window border will be placed at that
reference point. So, to place a window in the bottom right corner
you would first set gravity to south east, then write:
`gtk_window_move (window, gdk_screen_width () - window_width,
gdk_screen_height () - window_height)` (note that this
example does not take multi-head scenarios into account).

The [Extended Window Manager Hints Specification](http://www.freedesktop.org/Standards/wm-spec)
has a nice table of gravities in the “implementation notes” section.

The gtk_window_get_position() documentation may also be relevant.
*/
func (self *_TraitWindow) Move(x int, y int) {
	C.gtk_window_move(self.CPointer, C.gint(x), C.gint(y))
	return
}

/*
Parses a standard X Window System geometry string - see the
manual page for X (type “man X”) for details on this.
gtk_window_parse_geometry() does work on all GTK+ ports
including Win32 but is primarily intended for an X environment.

If either a size or a position can be extracted from the
geometry string, gtk_window_parse_geometry() returns %TRUE
and calls gtk_window_set_default_size() and/or gtk_window_move()
to resize/move the window.

If gtk_window_parse_geometry() returns %TRUE, it will also
set the #GDK_HINT_USER_POS and/or #GDK_HINT_USER_SIZE hints
indicating to the window manager that the size/position of
the window was user-specified. This causes most window
managers to honor the geometry.

Note that for gtk_window_parse_geometry() to work as expected, it has
to be called when the window has its “final” size, i.e. after calling
gtk_widget_show_all() on the contents and gtk_window_set_geometry_hints()
on the window.
|[<!-- language="C" -->
#include <gtk/gtk.h>

static void
fill_with_content (GtkWidget *vbox)
{
  // fill with content...
}

int
main (int argc, char *argv[])
{
  GtkWidget *window, *vbox;
  GdkGeometry size_hints = {
    100, 50, 0, 0, 100, 50, 10,
    10, 0.0, 0.0, GDK_GRAVITY_NORTH_WEST
  };

  gtk_init (&argc, &argv);

  window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
  vbox = gtk_box_new (GTK_ORIENTATION_VERTICAL,
                      FALSE, 0);

  gtk_container_add (GTK_CONTAINER (window), vbox);
  fill_with_content (vbox);
  gtk_widget_show_all (vbox);

  gtk_window_set_geometry_hints (GTK_WINDOW (window),
	  			    window,
				    &size_hints,
				    GDK_HINT_MIN_SIZE |
				    GDK_HINT_BASE_SIZE |
				    GDK_HINT_RESIZE_INC);

  if (argc > 1)
    {
      gboolean res;
      res = gtk_window_parse_geometry (GTK_WINDOW (window),
                                       argv[1]);
      if (! res)
        fprintf (stderr,
                 "Failed to parse “%s”\n",
                 argv[1]);
    }

  gtk_widget_show_all (window);
  gtk_main ();

  return 0;
}
]|
*/
func (self *_TraitWindow) ParseGeometry(geometry string) (return__ bool) {
	__cgo__geometry := (*C.gchar)(unsafe.Pointer(C.CString(geometry)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_parse_geometry(self.CPointer, __cgo__geometry)
	C.free(unsafe.Pointer(__cgo__geometry))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Presents a window to the user. This may mean raising the window
in the stacking order, deiconifying it, moving it to the current
desktop, and/or giving it the keyboard focus, possibly dependent
on the user’s platform, window manager, and preferences.

If @window is hidden, this function calls gtk_widget_show()
as well.

This function should be used when the user tries to open a window
that’s already open. Say for example the preferences dialog is
currently open, and the user chooses Preferences from the menu
a second time; use gtk_window_present() to move the already-open dialog
where the user can see it.

If you are calling this function in response to a user interaction,
it is preferable to use gtk_window_present_with_time().
*/
func (self *_TraitWindow) Present() {
	C.gtk_window_present(self.CPointer)
	return
}

/*
Presents a window to the user in response to a user interaction.
If you need to present a window without a timestamp, use
gtk_window_present(). See gtk_window_present() for details.
*/
func (self *_TraitWindow) PresentWithTime(timestamp uint32) {
	C.gtk_window_present_with_time(self.CPointer, C.guint32(timestamp))
	return
}

/*
Propagate a key press or release event to the focus widget and
up the focus container chain until a widget handles @event.
This is normally called by the default ::key_press_event and
::key_release_event handlers for toplevel windows,
however in some cases it may be useful to call this directly when
overriding the standard key handling for a toplevel window.
*/
func (self *_TraitWindow) PropagateKeyEvent(event *C.GdkEventKey) (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_propagate_key_event(self.CPointer, event)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Reverses the effects of gtk_window_add_accel_group().
*/
func (self *_TraitWindow) RemoveAccelGroup(accel_group *AccelGroup) {
	C.gtk_window_remove_accel_group(self.CPointer, (*C.GtkAccelGroup)(accel_group.CPointer))
	return
}

/*
Removes a mnemonic from this window.
*/
func (self *_TraitWindow) RemoveMnemonic(keyval uint, target *Widget) {
	C.gtk_window_remove_mnemonic(self.CPointer, C.guint(keyval), (*C.GtkWidget)(target.CPointer))
	return
}

// gtk_window_reshow_with_initial_size is not generated due to deprecation attr

/*
Resizes the window as if the user had done so, obeying geometry
constraints. The default geometry constraint is that windows may
not be smaller than their size request; to override this
constraint, call gtk_widget_set_size_request() to set the window's
request to a smaller value.

If gtk_window_resize() is called before showing a window for the
first time, it overrides any default size set with
gtk_window_set_default_size().

Windows may not be resized smaller than 1 by 1 pixels.
*/
func (self *_TraitWindow) Resize(width int, height int) {
	C.gtk_window_resize(self.CPointer, C.gint(width), C.gint(height))
	return
}

/*
Determines whether a resize grip is visible for the specified window.
*/
func (self *_TraitWindow) ResizeGripIsVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_resize_grip_is_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Like gtk_window_resize(), but @width and @height are interpreted
in terms of the base size and increment set with
gtk_window_set_geometry_hints.
*/
func (self *_TraitWindow) ResizeToGeometry(width int, height int) {
	C.gtk_window_resize_to_geometry(self.CPointer, C.gint(width), C.gint(height))
	return
}

/*
Windows may set a hint asking the desktop environment not to receive
the input focus. This function sets this hint.
*/
func (self *_TraitWindow) SetAcceptFocus(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_accept_focus(self.CPointer, __cgo__setting)
	return
}

/*
Sets or unsets the #GtkApplication associated with the window.

The application will be kept alive for at least as long as the window
is open.
*/
func (self *_TraitWindow) SetApplication(application *Application) {
	C.gtk_window_set_application(self.CPointer, (*C.GtkApplication)(application.CPointer))
	return
}

/*
Marks @window as attached to @attach_widget. This creates a logical binding
between the window and the widget it belongs to, which is used by GTK+ to
propagate information such as styling or accessibility to @window as if it
was a children of @attach_widget.

Examples of places where specifying this relation is useful are for instance
a #GtkMenu created by a #GtkComboBox, a completion popup window
created by #GtkEntry or a typeahead search entry created by #GtkTreeView.

Note that this function should not be confused with
gtk_window_set_transient_for(), which specifies a window manager relation
between two toplevels instead.

Passing %NULL for @attach_widget detaches the window.
*/
func (self *_TraitWindow) SetAttachedTo(attach_widget *Widget) {
	C.gtk_window_set_attached_to(self.CPointer, (*C.GtkWidget)(attach_widget.CPointer))
	return
}

/*
By default, windows are decorated with a title bar, resize
controls, etc.  Some [window managers][gtk-X11-arch]
allow GTK+ to disable these decorations, creating a
borderless window. If you set the decorated property to %FALSE
using this function, GTK+ will do its best to convince the window
manager not to decorate the window. Depending on the system, this
function may not have any effect when called on a window that is
already visible, so you should call it before calling gtk_widget_show().

On Windows, this function always works, since there’s no window manager
policy involved.
*/
func (self *_TraitWindow) SetDecorated(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_decorated(self.CPointer, __cgo__setting)
	return
}

/*
The default widget is the widget that’s activated when the user
presses Enter in a dialog (for example). This function sets or
unsets the default widget for a #GtkWindow. When setting (rather
than unsetting) the default widget it’s generally easier to call
gtk_widget_grab_default() on the widget. Before making a widget
the default widget, you must call gtk_widget_set_can_default() on
the widget you’d like to make the default.
*/
func (self *_TraitWindow) SetDefault(default_widget *Widget) {
	C.gtk_window_set_default(self.CPointer, (*C.GtkWidget)(default_widget.CPointer))
	return
}

/*
Like gtk_window_set_default_size(), but @width and @height are interpreted
in terms of the base size and increment set with
gtk_window_set_geometry_hints.
*/
func (self *_TraitWindow) SetDefaultGeometry(width int, height int) {
	C.gtk_window_set_default_geometry(self.CPointer, C.gint(width), C.gint(height))
	return
}

/*
Sets the default size of a window. If the window’s “natural” size
(its size request) is larger than the default, the default will be
ignored. More generally, if the default size does not obey the
geometry hints for the window (gtk_window_set_geometry_hints() can
be used to set these explicitly), the default size will be clamped
to the nearest permitted size.

Unlike gtk_widget_set_size_request(), which sets a size request for
a widget and thus would keep users from shrinking the window, this
function only sets the initial size, just as if the user had
resized the window themselves. Users can still shrink the window
again as they normally would. Setting a default size of -1 means to
use the “natural” default size (the size request of the window).

For more control over a window’s initial size and how resizing works,
investigate gtk_window_set_geometry_hints().

For some uses, gtk_window_resize() is a more appropriate function.
gtk_window_resize() changes the current size of the window, rather
than the size to be used on initial display. gtk_window_resize() always
affects the window itself, not the geometry widget.

The default size of a window only affects the first time a window is
shown; if a window is hidden and re-shown, it will remember the size
it had prior to hiding, rather than using the default size.

Windows can’t actually be 0x0 in size, they must be at least 1x1, but
passing 0 for @width and @height is OK, resulting in a 1x1 default size.
*/
func (self *_TraitWindow) SetDefaultSize(width int, height int) {
	C.gtk_window_set_default_size(self.CPointer, C.gint(width), C.gint(height))
	return
}

/*
By default, windows have a close button in the window frame. Some
[window managers][gtk-X11-arch] allow GTK+ to
disable this button. If you set the deletable property to %FALSE
using this function, GTK+ will do its best to convince the window
manager not to show a close button. Depending on the system, this
function may not have any effect when called on a window that is
already visible, so you should call it before calling gtk_widget_show().

On Windows, this function always works, since there’s no window manager
policy involved.
*/
func (self *_TraitWindow) SetDeletable(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_deletable(self.CPointer, __cgo__setting)
	return
}

/*
If @setting is %TRUE, then destroying the transient parent of @window
will also destroy @window itself. This is useful for dialogs that
shouldn’t persist beyond the lifetime of the main window they're
associated with, for example.
*/
func (self *_TraitWindow) SetDestroyWithParent(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_destroy_with_parent(self.CPointer, __cgo__setting)
	return
}

/*
If @focus is not the current focus widget, and is focusable, sets
it as the focus widget for the window. If @focus is %NULL, unsets
the focus widget for this window. To set the focus to a particular
widget in the toplevel, it is usually more convenient to use
gtk_widget_grab_focus() instead of this function.
*/
func (self *_TraitWindow) SetFocus(focus *Widget) {
	C.gtk_window_set_focus(self.CPointer, (*C.GtkWidget)(focus.CPointer))
	return
}

/*
Windows may set a hint asking the desktop environment not to receive
the input focus when the window is mapped.  This function sets this
hint.
*/
func (self *_TraitWindow) SetFocusOnMap(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_focus_on_map(self.CPointer, __cgo__setting)
	return
}

/*
Sets the #GtkWindow:focus-visible property.
*/
func (self *_TraitWindow) SetFocusVisible(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_focus_visible(self.CPointer, __cgo__setting)
	return
}

/*
This function sets up hints about how a window can be resized by
the user.  You can set a minimum and maximum size; allowed resize
increments (e.g. for xterm, you can only resize by the size of a
character); aspect ratios; and more. See the #GdkGeometry struct.
*/
func (self *_TraitWindow) SetGeometryHints(geometry_widget *Widget, geometry *C.GdkGeometry, geom_mask C.GdkWindowHints) {
	C.gtk_window_set_geometry_hints(self.CPointer, (*C.GtkWidget)(geometry_widget.CPointer), geometry, geom_mask)
	return
}

/*
Window gravity defines the meaning of coordinates passed to
gtk_window_move(). See gtk_window_move() and #GdkGravity for
more details.

The default window gravity is #GDK_GRAVITY_NORTH_WEST which will
typically “do what you mean.”
*/
func (self *_TraitWindow) SetGravity(gravity C.GdkGravity) {
	C.gtk_window_set_gravity(self.CPointer, gravity)
	return
}

/*
Sets whether @window has a corner resize grip.

Note that the resize grip is only shown if the window
is actually resizable and not maximized. Use
gtk_window_resize_grip_is_visible() to find out if the
resize grip is currently shown.
*/
func (self *_TraitWindow) SetHasResizeGrip(value bool) {
	__cgo__value := C.gboolean(0)
	if value {
		__cgo__value = C.gboolean(1)
	}
	C.gtk_window_set_has_resize_grip(self.CPointer, __cgo__value)
	return
}

/*
Tells GTK+ whether to drop its extra reference to the window
when gtk_widget_destroy() is called.

This function is only exported for the benefit of language
bindings which may need to keep the window alive until their
wrapper object is garbage collected. There is no justification
for ever calling this function in an application.
*/
func (self *_TraitWindow) SetHasUserRefCount(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_has_user_ref_count(self.CPointer, __cgo__setting)
	return
}

/*
If @setting is %TRUE, then @window will request that it’s titlebar
should be hidden when maximized.
This is useful for windows that don’t convey any information other
than the application name in the titlebar, to put the available
screen space to better use. If the underlying window system does not
support the request, the setting will not have any effect.

Note that custom titlebars set with gtk_window_set_titlebar() are
not affected by this. The application is in full control of their
content and visibility anyway.
*/
func (self *_TraitWindow) SetHideTitlebarWhenMaximized(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_hide_titlebar_when_maximized(self.CPointer, __cgo__setting)
	return
}

/*
Sets up the icon representing a #GtkWindow. This icon is used when
the window is minimized (also known as iconified).  Some window
managers or desktop environments may also place it in the window
frame, or display it in other contexts.

The icon should be provided in whatever size it was naturally
drawn; that is, don’t scale the image before passing it to
GTK+. Scaling is postponed until the last minute, when the desired
final size is known, to allow best quality.

If you have your icon hand-drawn in multiple sizes, use
gtk_window_set_icon_list(). Then the best size will be used.

This function is equivalent to calling gtk_window_set_icon_list()
with a 1-element list.

See also gtk_window_set_default_icon_list() to set the icon
for all windows in your application in one go.
*/
func (self *_TraitWindow) SetIcon(icon *C.GdkPixbuf) {
	C.gtk_window_set_icon(self.CPointer, icon)
	return
}

/*
Sets the icon for @window.
Warns on failure if @err is %NULL.

This function is equivalent to calling gtk_window_set_icon()
with a pixbuf created by loading the image from @filename.
*/
func (self *_TraitWindow) SetIconFromFile(filename string) (return__ bool, __err__ error) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.gtk_window_set_icon_from_file(self.CPointer, __cgo__filename, &__cgo_error__)
	C.free(unsafe.Pointer(__cgo__filename))
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets up the icon representing a #GtkWindow. The icon is used when
the window is minimized (also known as iconified).  Some window
managers or desktop environments may also place it in the window
frame, or display it in other contexts.

gtk_window_set_icon_list() allows you to pass in the same icon in
several hand-drawn sizes. The list should contain the natural sizes
your icon is available in; that is, don’t scale the image before
passing it to GTK+. Scaling is postponed until the last minute,
when the desired final size is known, to allow best quality.

By passing several sizes, you may improve the final image quality
of the icon, by reducing or eliminating automatic image scaling.

Recommended sizes to provide: 16x16, 32x32, 48x48 at minimum, and
larger images (64x64, 128x128) if you have them.

See also gtk_window_set_default_icon_list() to set the icon
for all windows in your application in one go.

Note that transient windows (those who have been set transient for another
window using gtk_window_set_transient_for()) will inherit their
icon from their transient parent. So there’s no need to explicitly
set the icon on transient windows.
*/
func (self *_TraitWindow) SetIconList(list *C.GList) {
	C.gtk_window_set_icon_list(self.CPointer, list)
	return
}

/*
Sets the icon for the window from a named themed icon. See
the docs for #GtkIconTheme for more details.

Note that this has nothing to do with the WM_ICON_NAME
property which is mentioned in the ICCCM.
*/
func (self *_TraitWindow) SetIconName(name string) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	C.gtk_window_set_icon_name(self.CPointer, __cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	return
}

/*
Asks to keep @window above, so that it stays on top. Note that
you shouldn’t assume the window is definitely above afterward,
because other entities (e.g. the user or
[window manager][gtk-X11-arch]) could not keep it above,
and not all window managers support keeping windows above. But
normally the window will end kept above. Just don’t write code
that crashes if not.

It’s permitted to call this function before showing a window,
in which case the window will be kept above when it appears onscreen
initially.

You can track the above state via the “window-state-event” signal
on #GtkWidget.

Note that, according to the
[Extended Window Manager Hints Specification](http://www.freedesktop.org/Standards/wm-spec),
the above state is mainly meant for user preferences and should not
be used by applications e.g. for drawing attention to their
dialogs.
*/
func (self *_TraitWindow) SetKeepAbove(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_keep_above(self.CPointer, __cgo__setting)
	return
}

/*
Asks to keep @window below, so that it stays in bottom. Note that
you shouldn’t assume the window is definitely below afterward,
because other entities (e.g. the user or
[window manager][gtk-X11-arch]) could not keep it below,
and not all window managers support putting windows below. But
normally the window will be kept below. Just don’t write code
that crashes if not.

It’s permitted to call this function before showing a window,
in which case the window will be kept below when it appears onscreen
initially.

You can track the below state via the “window-state-event” signal
on #GtkWidget.

Note that, according to the
[Extended Window Manager Hints Specification](http://www.freedesktop.org/Standards/wm-spec),
the above state is mainly meant for user preferences and should not
be used by applications e.g. for drawing attention to their
dialogs.
*/
func (self *_TraitWindow) SetKeepBelow(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_keep_below(self.CPointer, __cgo__setting)
	return
}

/*
Sets the mnemonic modifier for this window.
*/
func (self *_TraitWindow) SetMnemonicModifier(modifier C.GdkModifierType) {
	C.gtk_window_set_mnemonic_modifier(self.CPointer, modifier)
	return
}

/*
Sets the #GtkWindow:mnemonics-visible property.
*/
func (self *_TraitWindow) SetMnemonicsVisible(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_mnemonics_visible(self.CPointer, __cgo__setting)
	return
}

/*
Sets a window modal or non-modal. Modal windows prevent interaction
with other windows in the same application. To keep modal dialogs
on top of main application windows, use
gtk_window_set_transient_for() to make the dialog transient for the
parent; most [window managers][gtk-X11-arch]
will then disallow lowering the dialog below the parent.
*/
func (self *_TraitWindow) SetModal(modal bool) {
	__cgo__modal := C.gboolean(0)
	if modal {
		__cgo__modal = C.gboolean(1)
	}
	C.gtk_window_set_modal(self.CPointer, __cgo__modal)
	return
}

// gtk_window_set_opacity is not generated due to deprecation attr

/*
Sets a position constraint for this window. If the old or new
constraint is %GTK_WIN_POS_CENTER_ALWAYS, this will also cause
the window to be repositioned to satisfy the new constraint.
*/
func (self *_TraitWindow) SetPosition(position C.GtkWindowPosition) {
	C.gtk_window_set_position(self.CPointer, position)
	return
}

/*
Sets whether the user can resize a window. Windows are user resizable
by default.
*/
func (self *_TraitWindow) SetResizable(resizable bool) {
	__cgo__resizable := C.gboolean(0)
	if resizable {
		__cgo__resizable = C.gboolean(1)
	}
	C.gtk_window_set_resizable(self.CPointer, __cgo__resizable)
	return
}

/*
This function is only useful on X11, not with other GTK+ targets.

In combination with the window title, the window role allows a
[window manager][gtk-X11-arch] to identify "the
same" window when an application is restarted. So for example you
might set the “toolbox” role on your app’s toolbox window, so that
when the user restarts their session, the window manager can put
the toolbox back in the same place.

If a window already has a unique title, you don’t need to set the
role, since the WM can use the title to identify the window when
restoring the session.
*/
func (self *_TraitWindow) SetRole(role string) {
	__cgo__role := (*C.gchar)(unsafe.Pointer(C.CString(role)))
	C.gtk_window_set_role(self.CPointer, __cgo__role)
	C.free(unsafe.Pointer(__cgo__role))
	return
}

/*
Sets the #GdkScreen where the @window is displayed; if
the window is already mapped, it will be unmapped, and
then remapped on the new screen.
*/
func (self *_TraitWindow) SetScreen(screen *C.GdkScreen) {
	C.gtk_window_set_screen(self.CPointer, screen)
	return
}

/*
Windows may set a hint asking the desktop environment not to display
the window in the pager. This function sets this hint.
(A "pager" is any desktop navigation tool such as a workspace
switcher that displays a thumbnail representation of the windows
on the screen.)
*/
func (self *_TraitWindow) SetSkipPagerHint(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_skip_pager_hint(self.CPointer, __cgo__setting)
	return
}

/*
Windows may set a hint asking the desktop environment not to display
the window in the task bar. This function sets this hint.
*/
func (self *_TraitWindow) SetSkipTaskbarHint(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_skip_taskbar_hint(self.CPointer, __cgo__setting)
	return
}

/*
Startup notification identifiers are used by desktop environment to
track application startup, to provide user feedback and other
features. This function changes the corresponding property on the
underlying GdkWindow. Normally, startup identifier is managed
automatically and you should only use this function in special cases
like transferring focus from other processes. You should use this
function before calling gtk_window_present() or any equivalent
function generating a window map event.

This function is only useful on X11, not with other GTK+ targets.
*/
func (self *_TraitWindow) SetStartupId(startup_id string) {
	__cgo__startup_id := (*C.gchar)(unsafe.Pointer(C.CString(startup_id)))
	C.gtk_window_set_startup_id(self.CPointer, __cgo__startup_id)
	C.free(unsafe.Pointer(__cgo__startup_id))
	return
}

/*
Sets the title of the #GtkWindow. The title of a window will be
displayed in its title bar; on the X Window System, the title bar
is rendered by the [window manager][gtk-X11-arch],
so exactly how the title appears to users may vary
according to a user’s exact configuration. The title should help a
user distinguish this window from other windows they may have
open. A good title might include the application name and current
document filename, for example.
*/
func (self *_TraitWindow) SetTitle(title string) {
	__cgo__title := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	C.gtk_window_set_title(self.CPointer, __cgo__title)
	C.free(unsafe.Pointer(__cgo__title))
	return
}

/*
Sets a custom titlebar for @window.

If you set a custom titlebar, GTK+ will do its best to convince
the window manager not to put its own titlebar on the window.
Depending on the system, this function may not work for a window
that is already visible, so you set the titlebar before calling
gtk_widget_show().
*/
func (self *_TraitWindow) SetTitlebar(titlebar *Widget) {
	C.gtk_window_set_titlebar(self.CPointer, (*C.GtkWidget)(titlebar.CPointer))
	return
}

/*
Dialog windows should be set transient for the main application
window they were spawned from. This allows
[window managers][gtk-X11-arch] to e.g. keep the
dialog on top of the main window, or center the dialog over the
main window. gtk_dialog_new_with_buttons() and other convenience
functions in GTK+ will sometimes call
gtk_window_set_transient_for() on your behalf.

Passing %NULL for @parent unsets the current transient window.

On Windows, this function puts the child window on top of the parent,
much as the window manager would have done on X.
*/
func (self *_TraitWindow) SetTransientFor(parent *Window) {
	C.gtk_window_set_transient_for(self.CPointer, (*C.GtkWindow)(parent.CPointer))
	return
}

/*
By setting the type hint for the window, you allow the window
manager to decorate and handle the window in a way which is
suitable to the function of the window in your application.

This function should be called before the window becomes visible.

gtk_dialog_new_with_buttons() and other convenience functions in GTK+
will sometimes call gtk_window_set_type_hint() on your behalf.
*/
func (self *_TraitWindow) SetTypeHint(hint C.GdkWindowTypeHint) {
	C.gtk_window_set_type_hint(self.CPointer, hint)
	return
}

/*
Windows may set a hint asking the desktop environment to draw
the users attention to the window. This function sets this hint.
*/
func (self *_TraitWindow) SetUrgencyHint(setting bool) {
	__cgo__setting := C.gboolean(0)
	if setting {
		__cgo__setting = C.gboolean(1)
	}
	C.gtk_window_set_urgency_hint(self.CPointer, __cgo__setting)
	return
}

/*
Don’t use this function. It sets the X Window System “class” and
“name” hints for a window.  According to the ICCCM, you should
always set these to the same value for all windows in an
application, and GTK+ sets them to that value by default, so calling
this function is sort of pointless. However, you may want to call
gtk_window_set_role() on each window in your application, for the
benefit of the session manager. Setting the role allows the window
manager to restore window positions when loading a saved session.
*/
func (self *_TraitWindow) SetWmclass(wmclass_name string, wmclass_class string) {
	__cgo__wmclass_name := (*C.gchar)(unsafe.Pointer(C.CString(wmclass_name)))
	__cgo__wmclass_class := (*C.gchar)(unsafe.Pointer(C.CString(wmclass_class)))
	C.gtk_window_set_wmclass(self.CPointer, __cgo__wmclass_name, __cgo__wmclass_class)
	C.free(unsafe.Pointer(__cgo__wmclass_name))
	C.free(unsafe.Pointer(__cgo__wmclass_class))
	return
}

/*
Asks to stick @window, which means that it will appear on all user
desktops. Note that you shouldn’t assume the window is definitely
stuck afterward, because other entities (e.g. the user or
[window manager][gtk-X11-arch] could unstick it
again, and some window managers do not support sticking
windows. But normally the window will end up stuck. Just don't
write code that crashes if not.

It’s permitted to call this function before showing a window.

You can track stickiness via the “window-state-event” signal
on #GtkWidget.
*/
func (self *_TraitWindow) Stick() {
	C.gtk_window_stick(self.CPointer)
	return
}

/*
Asks to toggle off the fullscreen state for @window. Note that you
shouldn’t assume the window is definitely not full screen
afterward, because other entities (e.g. the user or
[window manager][gtk-X11-arch]) could fullscreen it
again, and not all window managers honor requests to unfullscreen
windows. But normally the window will end up restored to its normal
state. Just don’t write code that crashes if not.

You can track the fullscreen state via the “window-state-event” signal
on #GtkWidget.
*/
func (self *_TraitWindow) Unfullscreen() {
	C.gtk_window_unfullscreen(self.CPointer)
	return
}

/*
Asks to unmaximize @window. Note that you shouldn’t assume the
window is definitely unmaximized afterward, because other entities
(e.g. the user or [window manager][gtk-X11-arch])
could maximize it again, and not all window
managers honor requests to unmaximize. But normally the window will
end up unmaximized. Just don’t write code that crashes if not.

You can track maximization via the “window-state-event” signal
on #GtkWidget.
*/
func (self *_TraitWindow) Unmaximize() {
	C.gtk_window_unmaximize(self.CPointer)
	return
}

/*
Asks to unstick @window, which means that it will appear on only
one of the user’s desktops. Note that you shouldn’t assume the
window is definitely unstuck afterward, because other entities
(e.g. the user or [window manager][gtk-X11-arch]) could
stick it again. But normally the window will
end up stuck. Just don’t write code that crashes if not.

You can track stickiness via the “window-state-event” signal
on #GtkWidget.
*/
func (self *_TraitWindow) Unstick() {
	C.gtk_window_unstick(self.CPointer)
	return
}

type _TraitWindowGroup struct{ CPointer *C.GtkWindowGroup }

/*
Adds a window to a #GtkWindowGroup.
*/
func (self *_TraitWindowGroup) AddWindow(window *Window) {
	C.gtk_window_group_add_window(self.CPointer, (*C.GtkWindow)(window.CPointer))
	return
}

/*
Returns the current grab widget for @device, or %NULL if none.
*/
func (self *_TraitWindowGroup) GetCurrentDeviceGrab(device *C.GdkDevice) (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_window_group_get_current_device_grab(self.CPointer, device)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Gets the current grab widget of the given group,
see gtk_grab_add().
*/
func (self *_TraitWindowGroup) GetCurrentGrab() (return__ *Widget) {
	var __cgo__return__ *C.GtkWidget
	__cgo__return__ = C.gtk_window_group_get_current_grab(self.CPointer)
	return__ = NewWidgetFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	return
}

/*
Returns a list of the #GtkWindows that belong to @window_group.
*/
func (self *_TraitWindowGroup) ListWindows() (return__ *C.GList) {
	return__ = C.gtk_window_group_list_windows(self.CPointer)
	return
}

/*
Removes a window from a #GtkWindowGroup.
*/
func (self *_TraitWindowGroup) RemoveWindow(window *Window) {
	C.gtk_window_group_remove_window(self.CPointer, (*C.GtkWindow)(window.CPointer))
	return
}
