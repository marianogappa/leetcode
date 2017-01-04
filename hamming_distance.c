#include <stdio.h>

int hammingDistance(int x, int y) {
  int i, c = 0, xor = x ^ y;
  for (i = (sizeof(int) * 8) - 1; i >= 0; i--)
    if (xor & (1u << i))
      c++;

  return c;
}

int main()
{
  int a = 7, b = 0;
  printf("distance(%d, %d) = %d\n", a, b, hammingDistance(a, b));
}
