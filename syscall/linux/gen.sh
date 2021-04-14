pkg=syscall
OS=$GOOS
ARCH=$GOARCH

if test $GOARCH = arm; then
	CC=arm-linux-gnueabi-gcc
	export CC
fi

mksyscall=`go env GOROOT`/src/syscall/mksyscall.pl

perl $mksyscall -tags ${OS},$ARCH ${pkg}_$OS.go |
	sed 's/^package.*syscall$$/package $*/' |
	sed '/^import/a \
		import "syscall"' |
	sed 's/Syscall/syscall.Syscall/' |
	sed 's/SYS_/syscall.SYS_/' |
	gofmt > z${pkg}_${OS}_$ARCH.go

# note: cgo execution depends on $GOARCH value
go tool cgo -godefs $OS/types.go  |
	sed '/^.. cgo -godefs/s,[^ ]\+/types.go,linux/types.go,' |
	gofmt >ztypes_${OS}_$ARCH.go


(
	cat <<EOF
package $pkg
/*
#include <unistd.h>
#include <termios.h>
#include <linux/tty_flags.h>
#include <sys/ioctl.h>
*/
import "C"

const (
EOF
	<$OS/const awk '
		/^[^\/]/ { print "\t" $1 "= C." $1 ; next}
		{ print }
	'
	echo ')'
) > ,,const.go

go tool cgo -godefs ,,const.go |
	sed '/^.. cgo -godefs/s/[^ ]\+const.go/,,const.go/' |
	gofmt > zconst_${OS}_$ARCH.go
rm -f ,,const.go
