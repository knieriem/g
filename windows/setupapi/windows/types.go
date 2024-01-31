//go:build ignore

package setupapi

/*
#include "windows/c.h"
#define SP_DEVINFO_DATA_SZ sizeof(SP_DEVINFO_DATA)
*/
import "C"

type SpDevinfoData C.SP_DEVINFO_DATA
type Guid C.GUID

const (
	SpDevinfoDataSz = C.SP_DEVINFO_DATA_SZ
)

type DevPropKey C.DEVPROPKEY
type DevPropType C.DEVPROPTYPE
