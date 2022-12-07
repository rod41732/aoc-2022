#include <algorithm>
#include <cstdio>
#include <cstring>
#include <vector>
using namespace std;

int main() {
  freopen("Day01.txt", "r", stdin);
  vector<int> v;
  int acc = 0;
  int num;
  char line[100];
  while (fgets(line, 100, stdin) != NULL) {
    int r = sscanf(line, "%d", &num);
    if (r == EOF) {
      v.push_back(acc);
      acc = 0;
    } else {
      acc += num;
    }
  }

  sort(v.begin(), v.end());
  printf("Part 1: %d\n", v[v.size() - 1]);
  printf("Part 2: %d\n", v[v.size() - 3] + v[v.size() - 2] + v[v.size() - 1]);
}
