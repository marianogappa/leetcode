#include <stdio.h>
#include <stdint.h>

int hammingWeight(uint32_t n) {
  int i, acc = 0;
  for (i = (sizeof(int) * 8) - 1; i >= 0; i--)
    if (n & (1u << i))
      acc++;
  return acc;
}

int main()
{
  printf("hammingWeight = %d\n", hammingWeight(11));
}
