#include <stdio.h>
#include <stdlib.h>

int factorial(int n) {
  int acc = 1;
  for (int i = n; i > 0; i--)
    acc *= i;
  return acc;
}

int* removeIth(int* ns, int l, int ith) {
  int* nns = malloc((l - 1) * sizeof(int*));
  for (int i = 0, j = 0; i < l; i++) {
    if (i == ith)
      continue;
    nns[j] = *(ns + i);
    j++;
  }
  return nns;
}

void fillPs(int* nums, int numsSize, int** ps, int offset) {
  if (numsSize == 0)
    return;

  for (int i = 0; i < numsSize; i++) {
    int f = factorial(numsSize - 1);

    for (int j = 0; j < f; j++) {
      *(*(ps + i * f + j) + offset) = *(nums + i);
    }

    fillPs(removeIth(nums, numsSize, i), numsSize - 1, ps + i * f, offset + 1);
  }
}

int** permute(int* nums, int numsSize, int* returnSize) {
  if (numsSize == 0) {
    *returnSize = 0;
    return NULL;
  }

  *returnSize = factorial(numsSize);
  int** ps = malloc(*returnSize * sizeof(int**));
  for (int i = 0; i < *returnSize; i++)
    *(ps + i) = malloc(numsSize * sizeof(int));

  fillPs(nums, numsSize, ps, 0);

  return ps;
}

int main()
{
  int* rs = malloc(sizeof(int));
  int l = 2;
  int ns[2] = {1, 2};
  int** ps = permute(ns, l, rs);

  printf("return size = %d\n", *rs);
  for (int i = 0; i < *rs; i++) {
    for (int j = 0; j < l; j++)
      printf("%d", *(*(ps + i) + j));
    printf("\n");
  }
}

