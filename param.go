package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func (self *Param) GetGoName() string {
	goName := self.Name
	if isGoReservedWord(goName) {
		goName += "_"
	}
	return goName
}

func (self *Param) CollectInfo(isReturn bool, fn *Function) {
	// name
	self.GoName = self.GetGoName()
	if isReturn {
		self.GoName = "return__"
	}

	// type info
	if self.Array != nil { // array type
		self.IsArray = true
		self.CType = self.Array.CType
		self.CTypeName = self.Array.Name // for spec only
		self.ElementCType = self.Array.Type.CType
		self.ElementCTypeName = self.Array.Type.Name // for spec only
		self.ElementGoType = cTypeToGoType(self.ElementCType)
		if self.Array.Length != "" { // length param
			n, err := strconv.Atoi(self.Array.Length)
			checkError(err)
			self.LenParamName = fn.Params[n].GetGoName()
		}
		if self.Array.ZeroTerminated == "1" {
			self.IsZeroTerminated = true
		}
	} else if self.Type != nil { // non-array type
		self.CType = self.Type.CType
		self.CTypeName = self.Type.Name // for spec only
	}
	self.GoType = cTypeToGoType(self.CType) // go type in function signature
	if self.GoType == "" {
		log.Fatal("ctype is not converted to gotype " + self.CType)
	}

	// skip void type
	if self.CTypeName == "none" || self.CType == "void" || self.GoType == "C.void" {
		self.IsVoid = true
		return
	}

	// check direction
	if self.Direction == "out" { // out param
		if !strings.HasSuffix(self.CType, "*") || !strings.HasPrefix(self.GoType, "*") {
			log.Fatal("non-pointer out param " + self.CType)
		}
		// convert to non-pointer type
		self.CType = self.CType[:len(self.CType)-1]
		self.GoType = self.GoType[1:]
	} else if self.Direction == "" { // in param
		self.Direction = "in"
	}
	if isReturn {
		self.Direction = "out"
	}

	// type spec, for type mapping
	spec := fs("%s", self.Direction)
	if self.GoType != "" {
		spec += " GoType " + self.GoType
	}
	if self.CTypeName != "" {
		spec += " TypeName " + self.CTypeName
	}
	if self.IsArray {
		spec += " IsArray"
	}
	if self.ElementGoType != "" {
		spec += " ElemType " + self.ElementGoType
	}
	if self.ElementCTypeName != "" {
		spec += " ElemName " + self.ElementCTypeName
	}
	if self.LenParamName != "" {
		spec += " HasLenParam"
	}
	if self.IsZeroTerminated {
		spec += " ZeroTerminated"
	}
	self.TypeSpec = spec

	// type mapping
	self.MappedType = self.MapType()

	// cgo
	if self.Direction == "in" {
		self.PrepareInParam() // convert go value to c value
	} else if self.Direction == "out" {
		self.PrepareOutParam() // convert c value to go value
	}
}

var TypeMapping = make(map[string]string)

func (self *Param) MapType() (ret string) {
	// query type mapping dict
	var ok bool
	if ret, ok = TypeMapping[self.TypeSpec]; ok {
		return
	}

	// skip complex types
	if strings.HasPrefix(self.GoType, "C.") && unicode.IsUpper(rune(self.GoType[2])) {
		ret = self.GoType
		return
	}
	if strings.HasPrefix(self.GoType, "*C.") && unicode.IsUpper(rune(self.GoType[3])) {
		ret = self.GoType
		return
	}
	if strings.HasPrefix(self.GoType, "**C.") && unicode.IsUpper(rune(self.GoType[4])) {
		ret = self.GoType
		return
	}

	switch self.TypeSpec {

	// bool
	case "in GoType C.gboolean TypeName gboolean",
		"out GoType C.gboolean TypeName gboolean":
		ret = "bool"

	// numeric
	case "in GoType C.gint TypeName gint",
		"in GoType C.int TypeName gint",
		"out GoType C.int TypeName gint",
		"out GoType C.gint TypeName gint":
		ret = "int"
	case "in GoType C.guint TypeName guint",
		"out GoType C.guint TypeName guint":
		ret = "uint"
	case "out GoType C.gunichar TypeName gunichar",
		"in GoType C.gunichar TypeName gunichar":
		ret = "rune"
	case "out GoType C.guint16 TypeName guint16":
		ret = "uint16"
	case "in GoType C.guint32 TypeName guint32",
		"out GoType C.guint32 TypeName guint32":
		ret = "uint32"
	case "in GoType C.glong TypeName glong",
		"in GoType C.gsize TypeName gsize",
		"in GoType C.gssize TypeName gssize",
		"out GoType C.gsize TypeName gsize":
		ret = "int64"
	case "in GoType C.gfloat TypeName gfloat",
		"out GoType C.gfloat TypeName gfloat":
		ret = "float32"
	case "in GoType C.gdouble TypeName gdouble",
		"in GoType C.double TypeName gdouble",
		"out GoType C.double TypeName gdouble",
		"out GoType C.gdouble TypeName gdouble":
		ret = "float64"

	// string
	case "in GoType *C.gchar TypeName utf8",
		"in GoType *C.char TypeName utf8",
		"in GoType *C.gchar TypeName filename",
		"out GoType *C.char TypeName utf8",
		"out GoType *C.gchar TypeName filename",
		"in GoType *C.char TypeName filename",
		"out GoType *C.gchar TypeName utf8":
		ret = "string"

	// untyped pointer
	case "in GoType C.gpointer TypeName gpointer",
		"out GoType C.gpointer TypeName gpointer":
		ret = "unsafe.Pointer"

	// string slice
	case "in GoType **C.gchar IsArray ElemType *C.gchar ElemName utf8",
		"in GoType **C.gchar IsArray ElemName utf8":
		ret = "[]string"

	// numeric slice
	case "in GoType *C.gint IsArray ElemType C.gint ElemName gint",
		"in GoType *C.gint IsArray ElemType C.gint ElemName gint HasLenParam":
		ret = "[]int"

	// bytes
	case "out GoType *C.guint8 IsArray ElemType C.guint8 ElemName guint8 HasLenParam",
		"in GoType *C.guint8 IsArray ElemType C.guint8 ElemName guint8 HasLenParam",
		"in GoType *C.gchar IsArray ElemName guint8 HasLenParam":
		ret = "[]byte"

	// pointer to fundamental types
	case "in GoType *C.gint TypeName gint",
		"in GoType *C.guint16 TypeName guint16",
		"in GoType *C.int TypeName gint":
		ret = self.GoType

	// not enough info FIXME patch it
	case "out GoType *C.gint IsArray ElemType C.gint ElemName gint",
		"out GoType **C.gchar IsArray ElemName utf8",
		"out GoType *C.guint IsArray ElemType *C.guint ElemName guint":
		ret = self.GoType

	// inout to in param
	case "in GoType ***C.char IsArray ElemType **C.char ElemName utf8 HasLenParam",
		"in GoType ***C.gchar IsArray ElemType **C.gchar ElemName utf8 HasLenParam":
		ret = self.GoType

	// cairo TODO
	case "in GoType *C.cairo_surface_t TypeName cairo.Surface",
		"out GoType *C.cairo_region_t TypeName cairo.Region",
		"in GoType *C.cairo_t TypeName cairo.Context",
		"out GoType *C.cairo_surface_t TypeName cairo.Surface",
		"out GoType *C.cairo_t TypeName cairo.Context",
		"in GoType *C.cairo_region_t TypeName cairo.Region":
		ret = self.GoType

	default:
		p("FIXME type mapping %s @ %s -> %s\n", self.Name, self.Function.CIdentifier, self.TypeSpec)
		ret = self.GoType

	}
	return
}

var InParamMapping = make(map[string]func(*Param))

func (self *Param) PrepareInParam() {
	if self.GoType == self.MappedType { // not mapped
		self.CgoParam = self.GoName
	} else { // do type cast
		spec := fmt.Sprintf("%s -> %s", self.MappedType, self.GoType)

		// query mapping
		if f, ok := InParamMapping[spec]; ok {
			f(self)
			return
		}

		// helpers
		sliceToPointer := func(p string) {
			self.CgoBeforeStmt = fs("__header__%s := (*reflect.SliceHeader)(unsafe.Pointer(&%s))",
				self.GoName, self.GoName)
			self.CgoParam = fs("(%s)(unsafe.Pointer(__header__%s.Data))", p, self.GoName)
		}

		switch spec {
		// bool
		case "bool -> C.gboolean":
			self.CgoBeforeStmt = fs("__cgo__%s := C.gboolean(0); if %s { __cgo__%s = C.gboolean(1)};",
				self.GoName, self.GoName, self.GoName)
			self.CgoParam = fs("__cgo__%s", self.GoName)

		// char
		case "byte -> C.char":
			self.CgoParam = fs("C.char(%s)", self.GoName)
		case "byte -> C.gchar":
			self.CgoParam = fs("C.gchar(%s)", self.GoName)
		case "rune -> C.gunichar":
			self.CgoParam = fs("C.gunichar(%s)", self.GoName)

		// string
		case "string -> *C.char":
			self.CgoBeforeStmt = fs("__cgo__%s := C.CString(%s);", self.GoName, self.GoName)
			self.CgoParam = fs("__cgo__%s", self.GoName)
			self.CgoAfterStmt += fs("C.free(unsafe.Pointer(__cgo__%s));", self.GoName)
		case "string -> *C.gchar":
			self.CgoBeforeStmt = fs("__cgo__%s := (*C.gchar)(unsafe.Pointer(C.CString(%s)));", self.GoName, self.GoName)
			self.CgoParam = fs("__cgo__%s", self.GoName)
			self.CgoAfterStmt += fs("C.free(unsafe.Pointer(__cgo__%s));", self.GoName)

		// int
		case "int -> C.gint":
			self.CgoParam = fs("C.gint(%s)", self.GoName)
		case "int -> C.int":
			self.CgoParam = fs("C.int(%s)", self.GoName)
		case "uint -> C.guint":
			self.CgoParam = fs("C.guint(%s)", self.GoName)
		case "int8 -> C.gint8":
			self.CgoParam = fs("C.gint8(%s)", self.GoName)
		case "uint8 -> C.guint8":
			self.CgoParam = fs("C.guint8(%s)", self.GoName)
		case "uint16 -> C.guint16":
			self.CgoParam = fs("C.guint16(%s)", self.GoName)
		case "int32 -> C.gint32":
			self.CgoParam = fs("C.gint32(%s)", self.GoName)
		case "uint32 -> C.guint32":
			self.CgoParam = fs("C.guint32(%s)", self.GoName)
		case "int64 -> C.gint64":
			self.CgoParam = fs("C.gint64(%s)", self.GoName)
		case "uint64 -> C.guint64":
			self.CgoParam = fs("C.guint64(%s)", self.GoName)
		case "uint64 -> C.gulong":
			self.CgoParam = fs("C.gulong(%s)", self.GoName)
		case "int64 -> C.glong":
			self.CgoParam = fs("C.glong(%s)", self.GoName)
		case "int64 -> C.gsize":
			self.CgoParam = fs("C.gsize(%s)", self.GoName)
		case "int64 -> C.gssize":
			self.CgoParam = fs("C.gssize(%s)", self.GoName)
		case "int64 -> C.gulong":
			self.CgoParam = fs("C.gulong(%s)", self.GoName)

		// float
		case "float32 -> C.gfloat":
			self.CgoParam = fs("C.gfloat(%s)", self.GoName)
		case "float64 -> C.double":
			self.CgoParam = fs("C.double(%s)", self.GoName)
		case "float64 -> C.gdouble":
			self.CgoParam = fs("C.gdouble(%s)", self.GoName)

		// slice
		case "[]byte -> *C.gchar":
			sliceToPointer("*C.gchar")
		case "[]byte -> *C.guchar":
			sliceToPointer("*C.guchar")
		case "[]byte -> *C.guint8":
			sliceToPointer("*C.guint8")
		case "[]string -> **C.gchar":
			sliceToPointer("**C.gchar")
		case "[]int -> *C.gint":
			sliceToPointer("*C.gint")

		// untyped pointer
		case "unsafe.Pointer -> C.gpointer":
			self.CgoParam = fs("(C.gpointer)(%s)", self.GoName)
		case "unsafe.Pointer -> *C.void":
			self.CgoParam = fs("(*C.void)(%s)", self.GoName)
		case "unsafe.Pointer -> C.gconstpointer":
			self.CgoParam = fs("(C.gconstpointer)(%s)", self.GoName)

		default:
			p("FIXME in param -> %s\n", spec)
		}
	}
}

var OutParamMapping = make(map[string]func(*Param))

func (self *Param) PrepareOutParam() {
	if self.GoType == self.MappedType { // not mapped
		self.CgoParam = "&" + self.GoName
	} else { // mapped
		self.CgoBeforeStmt = fs("var __cgo__%s %s", self.GoName, self.GoType)
		self.CgoParam = "&__cgo__" + self.GoName

		spec := fs("%s -> %s", self.GoType, self.MappedType)

		if f, ok := OutParamMapping[spec]; ok {
			f(self)
			return
		}

		switch spec {

		// bool
		case "C.gboolean -> bool":
			self.CgoAfterStmt += fs("%s = __cgo__%s == C.gboolean(1);", self.GoName, self.GoName)

		// byte
		case "C.gchar -> byte":
			self.CgoAfterStmt += fs("%s = byte(__cgo__%s);", self.GoName, self.GoName)
		case "C.gunichar -> rune":
			self.CgoAfterStmt += fs("%s = rune(__cgo__%s);", self.GoName, self.GoName)

		// string
		case "*C.gchar -> string":
			self.CgoAfterStmt += fs("%s = C.GoString((*C.char)(unsafe.Pointer(__cgo__%s)));",
				self.GoName, self.GoName)
		case "*C.char -> string":
			self.CgoAfterStmt += fs("%s = C.GoString(__cgo__%s);", self.GoName, self.GoName)

		// numeric
		case "C.gint -> int",
			"C.int -> int":
			self.CgoAfterStmt += fs("%s = int(__cgo__%s);", self.GoName, self.GoName)
		case "C.guint -> uint":
			self.CgoAfterStmt += fs("%s = uint(__cgo__%s);", self.GoName, self.GoName)
		case "C.guint8 -> uint8":
			self.CgoAfterStmt += fs("%s = uint8(__cgo__%s);", self.GoName, self.GoName)
		case "C.guint16 -> uint16":
			self.CgoAfterStmt += fs("%s = uint16(__cgo__%s);", self.GoName, self.GoName)
		case "C.gint32 -> int32":
			self.CgoAfterStmt += fs("%s = int32(__cgo__%s);", self.GoName, self.GoName)
		case "C.guint32 -> uint32":
			self.CgoAfterStmt += fs("%s = uint32(__cgo__%s);", self.GoName, self.GoName)
		case "C.gsize -> int64",
			"C.gint64 -> int64",
			"C.gssize -> int64",
			"C.glong -> int64":
			self.CgoAfterStmt += fs("%s = int64(__cgo__%s);", self.GoName, self.GoName)
		case "C.guint64 -> uint64",
			"C.gulong -> uint64":
			self.CgoAfterStmt += fs("%s = uint64(__cgo__%s);", self.GoName, self.GoName)
		case "C.gfloat -> float32":
			self.CgoAfterStmt += fs("%s = float32(__cgo__%s);", self.GoName, self.GoName)
		case "C.double -> float64",
			"C.gdouble -> float64":
			self.CgoAfterStmt += fs("%s = float64(__cgo__%s);", self.GoName, self.GoName)

		// bytes
		case "*C.gchar -> []byte",
			"*C.guchar -> []byte",
			"*C.guint8 -> []byte":
			if self.LenParamName != "" { // len param
				// defer is needed because length param may be later.
				self.CgoAfterStmt += fs(`defer func() { %s = C.GoBytes(unsafe.Pointer(__cgo__%s), C.int(%s)) }();`,
					self.GoName, self.GoName, self.LenParamName)
			} else {
				p("no len param\n")
			}

		// string slice
		case "**C.gchar -> []string":
			self.CgoAfterStmt += fs(`var __slice__%s []*C.gchar
				__header__%s := (*reflect.SliceHeader)(unsafe.Pointer(&__slice__%s))
				__header__%s.Len = 4294967296
				__header__%s.Cap = 4294967296
				__header__%s.Data = uintptr(unsafe.Pointer(__cgo__%s))
				for _, p := range __slice__%s {
					if p == nil { break }
					%s = append(%s, C.GoString((*C.char)(unsafe.Pointer(p))))
				};`, self.GoName,
				self.GoName, self.GoName,
				self.GoName,
				self.GoName,
				self.GoName, self.GoName,
				self.GoName,
				self.GoName, self.GoName)

		// uint slice
		case "*C.guint -> []uint":
			self.CgoAfterStmt += fs(`defer func() {
				__header__%s := (*reflect.SliceHeader)(unsafe.Pointer(&%s))
				__header__%s.Len = int(%s)
				__header__%s.Cap = int(%s)
				__header__%s.Data = uintptr(unsafe.Pointer(__cgo__%s))
			}();`, self.GoName, self.GoName,
				self.GoName, self.LenParamName,
				self.GoName, self.LenParamName,
				self.GoName, self.GoName)

		// untyped pointer
		case "C.gpointer -> unsafe.Pointer":
			self.CgoAfterStmt += fs("%s = unsafe.Pointer(__cgo__%s);", self.GoName, self.GoName)

		default:
			p("FIXME out param -> %s\n", spec)
		}
	}
}
