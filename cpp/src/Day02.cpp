#include <algorithm>
#include <cstdio>
#include <cstring>
#include <utility>
#include <vector>
using namespace std;

int main() {
  freopen("Day02.txt", "r", stdin);
  vector<pair<int, int>> v;
  int acc = 0;
  char a, b;
  while (scanf("%c %c\n", &a, &b) != EOF) {
    v.push_back({a - 'A', b - 'X'});
  }

  int score1 = 0;
  for (auto [x, y] : v) {
    score1 += y + 1;
    int mod = (y - x + 3) % 3;
    if (mod == 1)
      score1 += 6;
    else if (mod == 0)
      score1 += 3;
    else
      score1 += 0;
  }
  printf("Part 1: %d\n", score1);

  int score2 = 0;
  for (auto [x, y] : v) {
    int playScore = (x + y + 2) % 3 + 1;
    int outcomeScore = y * 3;
    score2 += playScore + outcomeScore;
  }
  printf("Part 2: %d\n", score2);
}
