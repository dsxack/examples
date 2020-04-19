#include <stdlib.h>
#include "rustscryptbindings.h"

int main() {
    unsigned char* result = malloc(sizeof(unsigned char) * 32);
    char* password = "test1";
    char* salt = "salt1";

    scrypt_key(password, salt, 8, 1, 16, result, 32);
}
