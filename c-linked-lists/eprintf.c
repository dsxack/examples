#include "eprintf.h"

/* eprintf: выводит сообщение об ошибке и выходит */
void eprintf(char *fmt, ...) {
    va_list args;

    fflush(stdout);
    if (getprogname() != NULL) {
        fprintf(stderr, "%s: ", getprogname());
    }

    va_start(args, fmt);
    vfprintf(stderr, fmt, args);
    va_end(args);

    if (fmt[0] != '\0' && fmt[strlen(fmt) - 1] == ':') {
        fprintf(stderr, " %s", strerror(errno));
    }

    fprintf(stderr, "\n");
    exit(2); /* условный код аварийного выхода */
}

/* emalloc: вызывает malloc, сообщает об ошибке */
void *emalloc(size_t n) {
    void *p;

    p = malloc(n);
    if (p == NULL) {
        eprintf("malloc of %u bytes failed:", n);
    }

    return p;
}

/* weprintf: выводит сообщение об ошибке */
void weprintf(char *fmt, ...) {
    va_list args;

    fflush(stdout);
    if (getprogname() != NULL) {
        fprintf(stderr, "%s: ", getprogname());
    }

    va_start(args, fmt);
    vfprintf(stderr, fmt, args);
    va_end(args);

    if (fmt[0] != '\0' && fmt[strlen(fmt) - 1] == ':') {
        fprintf(stderr, " %s", strerror(errno));
    }

    fprintf(stderr, "\n");
}
