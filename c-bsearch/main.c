#include "string.h"
#include "stdio.h"

#define NELEMS(array) (sizeof(array) / sizeof(array[0]))

typedef struct Nameval Nameval;
struct Nameval {
    const char *name;
    int     value;
};

/* lookup: двоичный поиск строки name в массиве tab; возвращает индекс */
int lookup(const char *name, Nameval tab[], int ntab) {
    int low, high, mid, cmp;

    low = 0;
    high = ntab - 1;
    while (low <= high) {
        mid = (low + high) / 2;
        cmp = strcmp(name, tab[mid].name);
        if (cmp < 0) {
            high = mid - 1;
        } else if (cmp > 0) {
            low = mid + 1;
        } else { /* элемент найден */
            return mid;
        }
    }

    return -1; /* элемент не найден */
}

int main() {
    Nameval htmlchars[] = {
        "AElig",    0x00c6,
        "AAcute",   0x00c1,
        "Acicrc",   0x00c2,
        /* ... */
        "zeta",     0x03b6,
    };

    int index = lookup("zeta", htmlchars, NELEMS(htmlchars));

    printf("%x", htmlchars[index].value);

    return 0;
}
