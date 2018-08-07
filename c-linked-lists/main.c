#include <stdio.h>
#include <stdlib.h>
#include <zconf.h>
#include <memory.h>
#include <errno.h>

typedef struct Nameval Nameval;
struct Nameval {
    char *name;
    int value;
    Nameval *next; /* следующий в списке */
};

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

/* newitem: создает новый элемент по имени и значению */
Nameval *newitem(char *name, int value) {
    Nameval *newp;

    newp = (Nameval *) emalloc(sizeof(Nameval));
    newp->name = name;
    newp->value = value;
    newp->next = NULL;
    return newp;
}

/* addfront: добавляет newp в голову списка listp */
Nameval *addfront(Nameval *listp, Nameval *newp) {
    newp->next = listp;
    return newp;
}

/* addend: добавляет newp в конец списка listp */
Nameval *addend(Nameval *listp, Nameval *newp) {
    Nameval *p;

    if (listp == NULL) {
        return newp;
    }
    for (p = listp; p->next != NULL; p = p->next)
        ;
    p->next = newp;
    return listp;
}

/* lookup: последовательный поиск имени name в списке listp */
Nameval *lookup(Nameval *listp, char *name) {
    for (; listp != NULL; listp = listp->next) {
        if (strcmp(name, listp->name) == 0) {
            return listp;
        }
    }

    return NULL;
}

/* apply: выполняет fn для каждого элемента списка listp */
void apply(Nameval *listp,
           void (*fn)(Nameval *, void *), void *arg) {
    for (; listp != NULL; listp = listp->next) {
        (*fn)(listp, arg); /* вызов функции */
    }
}

/* printv: вывести имя и значения по строке формата arg */
void printv(Nameval *p, void *arg) {
    char *fmt;

    fmt = (char *) arg;
    printf(fmt, p->name, p->value);
}

/* inccounter: инкрементирует счетчик *arg */
void inccounter(Nameval *p, void *arg) {
    int *ip;

    /* p не используется */
    ip = (int *) arg;
    (*ip)++;
}

/* freaall: освобождение всех элементов списка listp */
void freeall(Nameval *listp) {
    Nameval *next;

    for (; listp != NULL; listp = next) {
        next = listp->next;
        /* name освобождается в другом месте */
        free(listp);
    }
}

/* delitem: удаляет первое имя name из списка listp */
Nameval *delitem(Nameval *listp, char *name) {
    Nameval *p, *prev;

    prev = NULL;
    for (p = listp; p != NULL; p = p->next) {
        if (strcmp(name, p->name) == 0) {
            if (prev == NULL) {
                listp = p->next;
            } else {
                prev->next = p->next;
            }
            free(p);
            return listp;
        }
        prev = p;
    }
    eprintf("delitem: %s not in list", name);
    return NULL; /* сюда управление не доходит */
}

int main() {
    Nameval *list = newitem("AElig", 0x00c6);

    list = addfront(list, newitem("AAcute", 0x00c1));
    list = addfront(list, newitem("Acicrc", 0x00c2));
    list = addend(list, newitem("zeta", 0x03b6));

    printf("\nApply printv: \n");
    apply(list, printv, "%s: %x\n");

    Nameval *zeta = lookup(list, "zeta");
    printf("\nLookup: %s: %x\n", zeta->name, zeta->value);

    int n;
    n = 0;
    apply(list, inccounter, &n);
    printf("\nApply inccounter: %d \n", n);

    list = delitem(list, "Acicrc");
    printf("\nDelitem: \n");
    apply(list, printv, "%s: %x\n");

    freeall(list);
    n = 0;
    apply(list, inccounter, &n);
    printf("\nFreeall:\n");
    apply(list, printv, "%s: %x\n");

    return 0;
}