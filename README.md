## Utility packages for Go.

*	__go9p__

	Utility functions for [go9p]

*	__image__

	A `Bitmap` image of `BinaryColor` values implementing the
	`image.Image` interface. It is used by packages *pnm* and *xbm*,
	which contain decoders for raw PBM and XBM images.

*	__ioutil__

	`ChannelizeReader` and `IsTerminal`.

*	__windows/registry__

	Access to Windows' registry database (still read-only). 


*	__windows/setupapi__

	Access to some SetupDi functions. 

*	__syscall__

	System functions for Linux and Windows that are
	not defined by golang.org/x/sys; some of them are
	needed to implement the above packages.

	They make use of `mksyscall.pl` resp. `mkwinsyscall` from `$GOROOT/src/pkg/syscall`.


[go9p]: http://code.google.com/p/go9p/
