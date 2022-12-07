#include <cstdio>
#include <utility>
#include <vector>

bool contains(int a, int b, int c, int d) { return a <= c && b >= d; }
bool containsPoint(int a, int point, int b) { return a <= point && point <= b; }

bool overlaps(int a, int b, int c, int d) {
  return containsPoint(a, c, b) || containsPoint(a, d, b) ||
         contains(c, d, b, a);
}

int main() {
  freopen("Day04.txt", "r", stdin);
  int ans1 = 0;
  int ans2 = 0;
  int a, b, c, d;
  while (scanf("%d-%d,%d-%d", &a, &b, &c, &d) != EOF) {
    ans1 += contains(a, b, c, d) || contains(c, d, a, b);
    ans2 += overlaps(a, b, c, d);
  }
  printf("Part 1: %d\n", ans1);
  printf("Part 2: %d\n", ans2);
}
