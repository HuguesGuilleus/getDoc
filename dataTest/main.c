#include <stdio.h>

// A cystom type for testing
typedef struct {
	int yolo;
	char swag;
} customType ;

int hello();

// C'est la fête!
//
// The famous main function
int main(int argc, char const *argv[]) {
	hello();
	return 0;
}

// My function por Say Hello Wordl.
int hello() {
	printf("Hello World\n");
	return 0;
}
