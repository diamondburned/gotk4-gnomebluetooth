// Code generated by girgen. DO NOT EDIT.

package gnomebluetooth

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// extern void _gotk4_gnomebluetooth1_ChooserButtonClass_chooser_created(BluetoothChooserButton*, GtkWidget*);
// extern void _gotk4_gnomebluetooth1_ChooserButton_ConnectChooserCreated(gpointer, GObject, guintptr);
import "C"

// glib.Type values for bluetooth-chooser-button.go.
var GTypeChooserButton = externglib.Type(C.bluetooth_chooser_button_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeChooserButton, F: marshalChooserButton},
	})
}

// ChooserButtonOverrider contains methods that are overridable.
type ChooserButtonOverrider interface {
	// The function takes the following parameters:
	//
	ChooserCreated(chooser gtk.Widgetter)
}

// ChooserButton: <structname>BluetoothChooserButton</structname> struct
// contains only private fields and should not be directly accessed.
type ChooserButton struct {
	_ [0]func() // equal guard
	gtk.Button
}

var (
	_ gtk.Binner          = (*ChooserButton)(nil)
	_ externglib.Objector = (*ChooserButton)(nil)
)

func classInitChooserButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.BluetoothChooserButtonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.BluetoothChooserButtonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ChooserCreated(chooser gtk.Widgetter) }); ok {
		pclass.chooser_created = (*[0]byte)(C._gotk4_gnomebluetooth1_ChooserButtonClass_chooser_created)
	}
}

//export _gotk4_gnomebluetooth1_ChooserButtonClass_chooser_created
func _gotk4_gnomebluetooth1_ChooserButtonClass_chooser_created(arg0 *C.BluetoothChooserButton, arg1 *C.GtkWidget) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ChooserCreated(chooser gtk.Widgetter) })

	var _chooser gtk.Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gtk.Widgetter)
			return ok
		})
		rv, ok := casted.(gtk.Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_chooser = rv
	}

	iface.ChooserCreated(_chooser)
}

func wrapChooserButton(obj *externglib.Object) *ChooserButton {
	return &ChooserButton{
		Button: gtk.Button{
			Bin: gtk.Bin{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: gtk.Buildable{
							Object: obj,
						},
					},
				},
			},
			Object: obj,
			Actionable: gtk.Actionable{
				Widget: gtk.Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: gtk.Buildable{
						Object: obj,
					},
				},
			},
			Activatable: gtk.Activatable{
				Object: obj,
			},
		},
	}
}

func marshalChooserButton(p uintptr) (interface{}, error) {
	return wrapChooserButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gnomebluetooth1_ChooserButton_ConnectChooserCreated
func _gotk4_gnomebluetooth1_ChooserButton_ConnectChooserCreated(arg0 C.gpointer, arg1 C.GObject, arg2 C.guintptr) {
	var f func(chooser *externglib.Object)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(chooser *externglib.Object))
	}

	var _chooser *externglib.Object // out

	_chooser = externglib.Take(unsafe.Pointer(&arg1))

	f(_chooser)
}

// ConnectChooserCreated: signal is sent when a popup dialogue is created for
// the user to select a device. This signal allows you to change the
// configuration and filtering of the tree from its defaults.
func (button *ChooserButton) ConnectChooserCreated(f func(chooser *externglib.Object)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(button, "chooser-created", false, unsafe.Pointer(C._gotk4_gnomebluetooth1_ChooserButton_ConnectChooserCreated), f)
}

// NewChooserButton returns a new ChooserButton widget.
//
// The function returns the following values:
//
//    - chooserButton: ChooserButton widget.
//
func NewChooserButton() *ChooserButton {
	var _cret *C.GtkWidget // in

	_cret = C.bluetooth_chooser_button_new()

	var _chooserButton *ChooserButton // out

	_chooserButton = wrapChooserButton(externglib.Take(unsafe.Pointer(_cret)))

	return _chooserButton
}

// Available returns whether there is a powered Bluetooth adapter.
//
// The function returns the following values:
//
//    - ok: TRUE if there is a powered Bluetooth adapter available, and the
//      button should be sensitive.
//
func (button *ChooserButton) Available() bool {
	var _arg0 *C.BluetoothChooserButton // out
	var _cret C.gboolean                // in

	_arg0 = (*C.BluetoothChooserButton)(unsafe.Pointer(externglib.InternObject(button).Native()))

	_cret = C.bluetooth_chooser_button_available(_arg0)
	runtime.KeepAlive(button)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
