#include <stdio.h>
#include <stdlib.h>
#include <memory.h>

typedef struct Nameval Nameval;
struct Nameval {
    char *name;
    int value;
};

struct Nvtab {
    int nval;
    int max;
    Nameval *nameval;
} nvtab;

enum {
    NVINIT = 1, NVGROW = 2
};

int addname(Nameval newname) {
    Nameval *nvp;

    if (nvtab.nameval == NULL) {
        nvtab.nameval = (Nameval *) malloc(NVINIT * sizeof(Nameval));
        if (nvtab.nameval == NULL) {
            return -1;
        }
        nvtab.max = NVINIT;
        nvtab.nval = 0;
    } else if (nvtab.nval >= nvtab.max) {
        nvp = (Nameval *) realloc(nvtab.nameval, (NVGROW * nvtab.max) * sizeof(Nameval));
        if (nvp == NULL) {
            return -1;
        }
        nvtab.max *= NVGROW;
        nvtab.nameval = nvp;
    }
    nvtab.nameval[nvtab.nval] = newname;
    return nvtab.nval++;
}

int delname(char *name) {
    int i;
    for (i = 0; i < nvtab.nval; ++i) {
        if (strcmp(nvtab.nameval[i].name, name) == 0) {
            memmove(nvtab.nameval + i, nvtab.nameval + i + 1,
                    (nvtab.nval - (i + 1)) * sizeof(Nameval));
            nvtab.nval--;
            return 1;
        }
    }

    return 0;
}

int main() {
    addname((Nameval) {.name = "AElig", .value = 0x00c6});
    addname((Nameval) {.name = "AAcute", .value = 0x00c1});
    addname((Nameval) {.name = "Acicrc", .value = 0x00c2});
    addname((Nameval) {.name = "zeta", .value = 0x03b6});

    delname("AAcute");

    for (int i = 0; i < nvtab.nval; ++i) {
        Nameval nameval = nvtab.nameval[i];
        printf("%s - %d\n", nameval.name, nameval.value);
    }

    return 0;
}