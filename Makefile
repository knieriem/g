include $(GOROOT)/src/Make.inc 

ifeq ($(GOOS),windows)
REGISTRY=registry
endif

DIRS=\
	syscall\
	go9p\
	ioutil\
	$(REGISTRY)\
	text\
	sercom\
	\
	examples\

include Make.dirs
