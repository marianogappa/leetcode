#include <stdio.h>

int isBadVersion(int version) {
  return version >=666;
}

int doFirstBadVersion(int floor, int ceil) {
  if (ceil <= floor) {
    return floor;
  }
  return isBadVersion((ceil-floor)/2+floor) ? doFirstBadVersion(floor, (ceil-floor)/2+floor) : doFirstBadVersion((ceil-floor)/2+floor+1, ceil);
}

int firstBadVersion(int n) {
  return doFirstBadVersion(0, n);
}

int main() {
  printf("%d\n", firstBadVersion(1000));
}
