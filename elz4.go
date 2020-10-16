package elz4

/*
#include "./C/include/elz4.h"
#include "./C/include/etype.h"

#cgo LDFLAGS: -L${SRCDIR}/C/lib -letk
*/
import "C"
import "unsafe"
import "fmt"

func EncB2B(in []byte, out* []byte) (error) {

	var ret C.int;
	
	ret = C.elz4_encb2b((C.constr)(unsafe.Pointer(&in    [0])), C.uint(len(in)), 
	                    (C.cstr  )(unsafe.Pointer(&(*out)[0])), C.uint(cap(*out)));

	if ret >= 0 {

		*out = (*out)[0:ret]
		return nil 
	}

	*out = (*out)[0:0]
	return fmt.Errorf("%s", C.GoString(C.elz4_err(ret)));
} 

func DecB2B(in []byte, out* []byte) (error) {
	var ret C.int;

	ret = C.elz4_decb2b((C.constr)(unsafe.Pointer(&in    [0])), C.uint(len(in)), 
	                    (C.cstr  )(unsafe.Pointer(&(*out)[0])), C.uint(cap(*out)));

	if ret >= 0 {
		*out = (*out)[0:ret]
		return nil
	}

	*out = (*out)[0:0]
	return fmt.Errorf("%s", C.GoString(C.elz4_err(ret)));
}

func FEncB2B(in []byte, out* []byte) (error) {

	var ret C.int;

	ret = C.elz4f_encb2b((C.constr)(unsafe.Pointer(&in    [0])), C.uint(len(in)), 
	                     (C.cstr  )(unsafe.Pointer(&(*out)[0])), C.uint(cap(*out)));

	if ret >= 0 {
		*out = (*out)[0:ret]
		return nil
	}

	*out = (*out)[0:0]
	return fmt.Errorf("%s", C.GoString(C.elz4_err(ret)));
} 

func FDecB2B(in []byte, out* []byte) (error) {
	var ret C.int;

	ret = C.elz4f_decb2b((C.constr)(unsafe.Pointer(&in    [0])), C.uint(len(in)), 
	                     (C.cstr  )(unsafe.Pointer(&(*out)[0])), C.uint(cap(*out)));

	if ret >= 0 {
		*out = (*out)[0:ret]
		return nil
	}

	*out = (*out)[0:0]
	return fmt.Errorf("%s", C.GoString(C.elz4_err(ret)));
} 