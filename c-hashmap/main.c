#include <stdio.h>
#include <memory.h>
#include "eprintf.h"

typedef struct Nameval Nameval;
struct Nameval {
    char *name;
    int value;
    Nameval *next; /* следующий в цепочке */
};

const int NHASH = 11;

Nameval *symtab[NHASH]; /* таблица символов */

enum {
    MULTIPLIER = 31
};

unsigned int hash(char *str) {
    unsigned int h;
    unsigned char *p;

    h = 0;
    for (p = (unsigned char *) str; *p != '\0'; p++) {
        h = MULTIPLIER * h + *p;
    }

    return h % NHASH;
}

Nameval *lookup(char *name, int create, int value) {
    int h;
    Nameval *sym;

    h = hash(name);
    for (sym = symtab[h]; sym != NULL; sym = sym->next) {
        if (strcmp(name, sym->name) == 0) {
            return sym;
        }
    }
    if (create) {
        sym = (Nameval *) emalloc(sizeof(Nameval));
        sym->name = name; /* размещается в другом месте */
        sym->value = value;
        sym->next = symtab[h];
        symtab[h] = sym;
    }

    return sym;
}

int main() {
    lookup("AElig", 1, 0x00c);
    lookup("zeta", 1, 0x03b6);
    lookup("Acicrc", 1, 0x00c2);
    lookup("AAcute", 1, 0x00c1);

    Nameval *zeta = lookup("zeta", 0, 0);
    printf("\nlookup: %s %x", zeta->name, zeta->value);

    return 0;
}

