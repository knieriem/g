## Utility packages for Go.

*	__go9p__

	Utility functions for [go9p][]


*	__text__

	Plain text processing utilities.

	Contains `Tokenize`, an implementation of an
	equally named [function of Plan 9's libc][tokenize]
	(similar to `string.Fields`, but with interpretation of
	single quotes).

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

	System functions for Linux and Windows that were
	needed to implement the above packages.

	The make use of the `mksyscall*.sh` scripts from `$GOROOT/src/pkg/syscall`.

[9P]: http://plan9.bell-labs.com/sys/man/5/INDEX.html
[go9p]: http://code.google.com/p/go9p/
[hg-git]: http://hg-git.github.com/
[tokenize]: http://plan9.bell-labs.com/magic/man2html/2/getfields
