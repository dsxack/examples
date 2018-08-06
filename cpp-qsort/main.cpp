#include <cstdlib>
#include <cstdio>

#define NELEMS(array) (sizeof(array) / sizeof(array[0]))

void swap(int v[], int i, int j) {
    int temp;

    temp = v[i];
    v[i] = v[j];
    v[j] = temp;
}

void quicksort(int v[], int n) {
    int i, last;

    if (n <= 1) {   /* ничего не нужно делать */
        return;
    }

    swap(v, 0, rand() % n);     /* переместить опору в v[0] */
    last = 0;
    for (i = 0; i < n; ++i) {   /* разбиение */
        if (v[i] < v[0]) {
            swap(v, ++last, i);
        }
    }
    swap(v, 0, last);       /* восстановить опору */
    quicksort(v, last);     /* рекурсивная сортировка */
    quicksort(v+last, n-last-1);    /* каждой из частей */
}

int main() {
    int values[] = {5, 3, 4, 1, 2, 9, 8, 6, 7};

    quicksort(values, NELEMS(values));

    for (int value : values) {
        printf("%d ", value);
    }

    return 0;
}
