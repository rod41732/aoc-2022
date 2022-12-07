#include <algorithm>
#include <cstdio>
#include <functional>
#include <set>
#include <string>
#include <vector>

int main() {
  std::vector<std::string> rucksacks;
  freopen("Day03.txt", "r", stdin);
  // freopen("Day03_test.txt", "r", stdin);

  char buf[100];
  while (scanf("%s\n", buf) != EOF) {
    rucksacks.push_back(buf);
  }

  int ans1 = 0;
  for (auto &ruck : rucksacks) {
    char common;
    std::set<char> acc;
    acc.insert(ruck.begin(), ruck.begin() + ruck.size() / 2);

    for (auto c : ruck.substr(ruck.size() / 2)) {
      if (acc.count(c)) {
        common = c;
        break;
      }
    }
    ans1 += ((common | 32) - 'a' + 1) + (common & 32 ? 0 : 26);
  }
  printf("Part1: %d\n", ans1);

  int ans2 = 0;
  for (int i = 0; i < rucksacks.size(); i += 3) {
    std::set<char> s1;
    s1.insert(rucksacks[i].begin(), rucksacks[i].end());
    std::set<char> s2;
    s2.insert(rucksacks[i + 1].begin(), rucksacks[i + 1].end());
    std::set<char> s3;
    s3.insert(rucksacks[i + 2].begin(), rucksacks[i + 2].end());

    char common = 0;
    for (auto c : s1) {
      if (s2.count(c) && s3.count(c)) {
        common = c;
        break;
      }
    }
    if (!common) {
      for (auto c : s2) {
        if (s1.count(c) && s3.count(c)) {
          common = c;
          break;
        }
      }
    }
    if (!common) {
      for (auto c : s3) {
        if (s1.count(c) && s2.count(c)) {
          common = c;
          break;
        }
      }
    }

    ans2 += ((common | 32) - 'a' + 1) + (common & 32 ? 0 : 26);
  }
  printf("Part2: %d\n", ans2);
}
