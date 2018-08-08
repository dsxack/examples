//
// Created by Dmitriy Smotrov on 09/08/2018.
//

#ifndef C_BTREE_EPRINTF_H
#define C_BTREE_EPRINTF_H

#include <zconf.h>
#include <stdlib.h>
#include <stdio.h>
#include <memory.h>
#include <errno.h>

void eprintf(char *fmt, ...);
void *emalloc(size_t n);
void weprintf(char *fmt, ...);

#endif //C_BTREE_EPRINTF_H
