#include "main.h"

int yolo6Var = 42;

/**
 * C'est la fête!
 *
 * The famous main function
 */
int main(int argc, char const *argv[]) {
	hello();
	ERR("C'est la merde %d!!!\n", 156);
	return 0;
}

// My function por Say Hello World.
int hello() {
	printf(HELLO);
	printf(YOLO);
	return 0;
}

int *yolo() {
	printf(YOLO);
	printf(YOLO);
}
