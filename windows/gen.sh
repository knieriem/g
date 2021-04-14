# depends on environment variables:
#	PKG, ZDIR, GOARCH

set -e

pkg=$PKG
OS=$GOOS
GOROOT=`go env GOROOT`

case $GOARCH in
386)
	gccarch=i686-w64-mingw32
	;;
amd64)
	gccarch=x86_64-w64-mingw32
	;;
*)
	echo GOARCH $GOARCH not supported
	exit 1
	;;
esac

GCC=/usr/bin/$gccarch-gcc

SFX=_${OS}_$GOARCH.go

src=${pkg}_$OS.go
mv $src _$src
sed '/^package/s,syscall,none,' <_$src >$src
GOARCH= GOOS= go run golang.org/x/sys/windows/mkwinsyscall $src |
	sed 's/^package.*none/package '$pkg'/' |
	gofmt > z$pkg$SFX
rm -f $src
mv _$src $src

if test -f $OS/types.go; then
	# note: cgo execution depends on $GOARCH value
	CC=$GCC go tool cgo -godefs $OS/types.go  |
		sed '/^.. cgo -godefs/s,[^ ]\+types.go,windows/types.go,' |
		sed '/Pad_cgo_0/c\
		Flags	uint32' |
		awk -f $ZDIR/fixtype.awk |
		gofmt >ztypes$SFX
fi

if test -f $OS/const; then :
else
	exit 0
fi

(
	cat <<EOF
package $pkg

/*
#include "$OS/c.h"
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

CC=$GCC go tool cgo -godefs ,,const.go |
	sed '/^.. cgo -godefs/s/[^ ]\+const.go/,,const.go/' |
	awk -f $ZDIR/fixtype.awk |
	gofmt > zconst$SFX
rm -f ,,const.go
