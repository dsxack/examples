#include <stdio.h>
#include <memory.h>
#include <stdlib.h>
#include <zconf.h>
#include <errno.h>

typedef struct Nameval Nameval;
struct Nameval {
    char *name;
    int value;
    Nameval *left;  /* меньшее значение */
    Nameval *right; /* большее значение */
};

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

/* insert: вставляет newp в дерево treep, возвращает treep */
Nameval *insert(Nameval *treep, Nameval *newp) {
    int cmp;

    if (treep == NULL) {
        return newp;
    }

    cmp = strcmp(newp->name, treep->name);
    if (cmp == 0) {
        weprintf("insert: duplicate entry %s ignored",
                newp->name);
    } else if (cmp < 0) {
        treep->left = insert(treep->left, newp);
    } else {
        treep->right = insert(treep->right, newp);
    }

    return treep;
}

/* lookup: ищет имя name в дереве treep */
Nameval *lookup(Nameval *treep, char *name) {
    int cmp;

    if (treep == NULL) {
        return NULL;
    }

    cmp = strcmp(name, treep->name);
    if (cmp == 0) {
        return treep;
    } else if (cmp < 0) {
        return lookup(treep->left, name);
    } else {
        return lookup(treep->right, name);
    }
}

/* nrlookup: нерекурсивный поиск имени name в дереве treep */
Nameval *nrlookup(Nameval *treep, char *name) {
    int cmp;

    while (treep != NULL) {
        cmp = strcmp(name, treep->name);

        if (cmp == 0) {
            return treep;
        } else if (cmp < 0) {
            treep = treep->left;
        } else {
            treep = treep->right;
        }
    }

    return NULL;
}

/* applyinorder: симметричное применение функции fn к treep */
void applyinorder(Nameval *treep,
        void (*fn)(Nameval*, void*), void *arg) {
    if (treep == NULL) {
        return;
    }

    applyinorder(treep->left, fn, arg);
    (*fn)(treep, arg);
    applyinorder(treep->right, fn, arg);
}

/* applypostorder: концевой обход с вызовом fn */
void applypostorder(Nameval *treep,
                  void (*fn)(Nameval*, void*), void *arg) {
    if (treep == NULL) {
        return;
    }

    applypostorder(treep->left, fn, arg);
    applypostorder(treep->right, fn, arg);
    (*fn)(treep, arg);
}

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

/* printv: вывести имя и значения по строке формата arg */
void printv(Nameval *p, void *arg) {
    char *fmt;

    fmt = (char *) arg;
    printf(fmt, p->name, p->value);
}

/* newitem: создает новый элемент по имени и значению */
Nameval *newitem(char *name, int value) {
    Nameval *newp;

    newp = (Nameval *) emalloc(sizeof(Nameval));
    newp->name = name;
    newp->value = value;
    newp->right = NULL;
    newp->left = NULL;
    return newp;
}

int main() {
    Nameval *treep = newitem("AElig", 0x00c6);

    treep = insert(treep, newitem("zeta", 0x03b6));
    treep = insert(treep, newitem("Acicrc", 0x00c2));
    treep = insert(treep, newitem("AAcute", 0x00c1));

    printf("\nApply in order printv: \n");
    applyinorder(treep, printv, "%s: %x\n");

    printf("\nApply post order printv: \n");
    applypostorder(treep, printv, "%s: %x\n");

    Nameval *zeta = lookup(treep, "zeta");
    printf("\nlookup: %s: %x\n", zeta->name, zeta->value);

    Nameval *aacute = nrlookup(treep, "AAcute");
    printf("\nnrlookup: %s: %x\n", aacute->name, aacute->value);

    return 0;
}
