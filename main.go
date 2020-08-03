package main

// #include <string.h>
// #include <stdbool.h>
// #include <mysql.h>
// #cgo CFLAGS: -O3 -I/usr/include/mysql -fno-omit-frame-pointer
import "C"
import (
	"net/url"
	"unsafe"
)

func msg(message *C.char, s string) {
	m := C.CString(s)
	defer C.free(unsafe.Pointer(m))

	C.strcpy(message, m)
}

//export query_escape_init
func query_escape_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 1 {
		msg(message, "`query_escape` requires 1 parameter: the string to be escaped")
		return C.bool(true)
	}

	argsTypes := (*[2]uint32)(unsafe.Pointer(args.arg_type))

	argsTypes[0] = C.STRING_RESULT
	initid.maybe_null = 1

	return C.bool(false)
}

//export query_escape
func query_escape(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, message *C.char) *C.char {
	c := 1
	argsArgs := (*[1 << 30]*C.char)(unsafe.Pointer(args.args))[:c:c]

	if argsArgs[0] == nil {
		*length = 0
		*isNull = 1
		return nil
	}

	a := make([]string, c, c)
	for i, argsArg := range argsArgs {
		// This should be the correct way, but lengths come through as "0"
		// for everything after the first argument, so hopefully we don't
		// encounter any URLs or param names with null bytes in them (not really that worried)
		// a[i] = C.GoStringN(argsArg, C.int(argsLengths[i]))

		a[i] = C.GoString(argsArg)
	}

	s := url.QueryEscape(a[0])

	*length = uint64(len(s))
	return C.CString(s)
}

//export query_unescape_init
func query_unescape_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 1 {
		msg(message, "`query_unescape` requires 1 parameter: the string to be unescaped")
		return C.bool(true)
	}

	argsTypes := (*[2]uint32)(unsafe.Pointer(args.arg_type))

	argsTypes[0] = C.STRING_RESULT
	initid.maybe_null = 1

	return C.bool(false)
}

//export query_unescape
func query_unescape(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, message *C.char) *C.char {
	c := 1
	argsArgs := (*[1 << 30]*C.char)(unsafe.Pointer(args.args))[:c:c]

	if argsArgs[0] == nil {
		*length = 0
		*isNull = 1
		return nil
	}

	a := make([]string, c, c)
	for i, argsArg := range argsArgs {
		// This should be the correct way, but lengths come through as "0"
		// for everything after the first argument, so hopefully we don't
		// encounter any URLs or param names with null bytes in them (not really that worried)
		// a[i] = C.GoStringN(argsArg, C.int(argsLengths[i]))

		a[i] = C.GoString(argsArg)
	}

	s, err := url.QueryUnescape(a[0])
	if err != nil {
		*length = 0
		*isNull = 1
		return nil
	}

	*length = uint64(len(s))
	return C.CString(s)
}

func main() {}
