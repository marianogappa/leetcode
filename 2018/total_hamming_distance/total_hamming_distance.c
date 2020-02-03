#include <stdio.h>
#include <stdlib.h>

int totalHammingDistance(int* nums, int numsSize) {
  int i, o, acc = 0;
  for (i = 1; i; i <<= 1) {
    int os = 0;
    for (o = 0; o < numsSize; o++)
      if (*(nums + o)&i)
        os++;

    acc += os * (numsSize - os);
  }

  return acc;
}

int main()
{
  int ns[3] = {4, 14, 2};
  printf("totalHammingDistance = %d\n", totalHammingDistance(&ns[0], 3));
}
