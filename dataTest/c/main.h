#ifndef MAIN_H
	#define MAIN_H

	#include <stdio.h>

	// The hello message
	#define HELLO "Hello World\n"
	#define YOLO "Yolo!!!!!!!!!!!!!!!!!!!!!!\n"

	// A printer of error
	#define ERR(xxx ...)	fprintf(stderr, xxx)

	// A cystom type for testing
	typedef struct {
		int yolo;
		char swag;
	} customType ;

	int hello();

#endif
