// Code generated by 'ccgo errno/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o errno/errno_netbsd_amd64.go -pkgname errno', DO NOT EDIT.

package errno

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	E2BIG              = 7  // errno.h:48:1:
	EACCES             = 13 // errno.h:55:1:
	EADDRINUSE         = 48 // errno.h:97:1:
	EADDRNOTAVAIL      = 49 // errno.h:98:1:
	EAFNOSUPPORT       = 47 // errno.h:96:1:
	EAGAIN             = 35 // errno.h:81:1:
	EALREADY           = 37 // errno.h:84:1:
	EAUTH              = 80 // errno.h:140:1:
	EBADF              = 9  // errno.h:50:1:
	EBADMSG            = 88 // errno.h:159:1:
	EBADRPC            = 72 // errno.h:130:1:
	EBUSY              = 16 // errno.h:58:1:
	ECANCELED          = 87 // errno.h:156:1:
	ECHILD             = 10 // errno.h:51:1:
	ECONNABORTED       = 53 // errno.h:104:1:
	ECONNREFUSED       = 61 // errno.h:112:1:
	ECONNRESET         = 54 // errno.h:105:1:
	EDEADLK            = 11 // errno.h:52:1:
	EDESTADDRREQ       = 39 // errno.h:88:1:
	EDOM               = 33 // errno.h:77:1:
	EDQUOT             = 69 // errno.h:125:1:
	EEXIST             = 17 // errno.h:59:1:
	EFAULT             = 14 // errno.h:56:1:
	EFBIG              = 27 // errno.h:69:1:
	EFTYPE             = 79 // errno.h:139:1:
	EHOSTDOWN          = 64 // errno.h:118:1:
	EHOSTUNREACH       = 65 // errno.h:119:1:
	EIDRM              = 82 // errno.h:144:1:
	EILSEQ             = 85 // errno.h:149:1:
	EINPROGRESS        = 36 // errno.h:83:1:
	EINTR              = 4  // errno.h:45:1:
	EINVAL             = 22 // errno.h:64:1:
	EIO                = 5  // errno.h:46:1:
	EISCONN            = 56 // errno.h:107:1:
	EISDIR             = 21 // errno.h:63:1:
	ELAST              = 96 // errno.h:175:1:
	ELOOP              = 62 // errno.h:114:1:
	EMFILE             = 24 // errno.h:66:1:
	EMLINK             = 31 // errno.h:73:1:
	EMSGSIZE           = 40 // errno.h:89:1:
	EMULTIHOP          = 94 // errno.h:171:1:
	ENAMETOOLONG       = 63 // errno.h:115:1:
	ENEEDAUTH          = 81 // errno.h:141:1:
	ENETDOWN           = 50 // errno.h:101:1:
	ENETRESET          = 52 // errno.h:103:1:
	ENETUNREACH        = 51 // errno.h:102:1:
	ENFILE             = 23 // errno.h:65:1:
	ENOATTR            = 93 // errno.h:168:1:
	ENOBUFS            = 55 // errno.h:106:1:
	ENODATA            = 89 // errno.h:162:1:
	ENODEV             = 19 // errno.h:61:1:
	ENOENT             = 2  // errno.h:43:1:
	ENOEXEC            = 8  // errno.h:49:1:
	ENOLCK             = 77 // errno.h:136:1:
	ENOLINK            = 95 // errno.h:172:1:
	ENOMEM             = 12 // errno.h:54:1:
	ENOMSG             = 83 // errno.h:145:1:
	ENOPROTOOPT        = 42 // errno.h:91:1:
	ENOSPC             = 28 // errno.h:70:1:
	ENOSR              = 90 // errno.h:163:1:
	ENOSTR             = 91 // errno.h:164:1:
	ENOSYS             = 78 // errno.h:137:1:
	ENOTBLK            = 15 // errno.h:57:1:
	ENOTCONN           = 57 // errno.h:108:1:
	ENOTDIR            = 20 // errno.h:62:1:
	ENOTEMPTY          = 66 // errno.h:120:1:
	ENOTSOCK           = 38 // errno.h:87:1:
	ENOTSUP            = 86 // errno.h:153:1:
	ENOTTY             = 25 // errno.h:67:1:
	ENXIO              = 6  // errno.h:47:1:
	EOPNOTSUPP         = 45 // errno.h:94:1:
	EOVERFLOW          = 84 // errno.h:146:1:
	EPERM              = 1  // errno.h:42:1:
	EPFNOSUPPORT       = 46 // errno.h:95:1:
	EPIPE              = 32 // errno.h:74:1:
	EPROCLIM           = 67 // errno.h:123:1:
	EPROCUNAVAIL       = 76 // errno.h:134:1:
	EPROGMISMATCH      = 75 // errno.h:133:1:
	EPROGUNAVAIL       = 74 // errno.h:132:1:
	EPROTO             = 96 // errno.h:173:1:
	EPROTONOSUPPORT    = 43 // errno.h:92:1:
	EPROTOTYPE         = 41 // errno.h:90:1:
	ERANGE             = 34 // errno.h:78:1:
	EREMOTE            = 71 // errno.h:129:1:
	EROFS              = 30 // errno.h:72:1:
	ERPCMISMATCH       = 73 // errno.h:131:1:
	ESHUTDOWN          = 58 // errno.h:109:1:
	ESOCKTNOSUPPORT    = 44 // errno.h:93:1:
	ESPIPE             = 29 // errno.h:71:1:
	ESRCH              = 3  // errno.h:44:1:
	ESTALE             = 70 // errno.h:128:1:
	ETIME              = 92 // errno.h:165:1:
	ETIMEDOUT          = 60 // errno.h:111:1:
	ETOOMANYREFS       = 59 // errno.h:110:1:
	ETXTBSY            = 26 // errno.h:68:1:
	EUSERS             = 68 // errno.h:124:1:
	EWOULDBLOCK        = 35 // errno.h:82:1:
	EXDEV              = 18 // errno.h:60:1:
	X_ERRNO_H_         = 0  // errno.h:40:1:
	X_FILE_OFFSET_BITS = 64 // <builtin>:25:1:
	X_LP64             = 1  // <predefined>:268:1:
	X_NETBSD_SOURCE    = 1  // featuretest.h:70:1:
	X_SYS_CDEFS_ELF_H_ = 0  // cdefs_elf.h:31:1:
	X_SYS_CDEFS_H_     = 0  // cdefs.h:37:1:
	X_SYS_ERRNO_H_     = 0  // errno.h:40:1:
	X_X86_64_CDEFS_H_  = 0  // cdefs.h:4:1:
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

var _ int8 /* gen.c:2:13: */
