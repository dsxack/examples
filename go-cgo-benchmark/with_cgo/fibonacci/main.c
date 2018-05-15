#include "./fibonacci.h"
#include "./fibonacci.c"

int main() {
    int i = 0;

    for (i = 1; i <= 10000; i++) {
        fibonacci(i);
    }

    return 0;
}
