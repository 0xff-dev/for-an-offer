class Solution:
    def frequencySort(self, s: str) -> str:
        return ''.join(sorted(set(map(lambda x:f'{x}'*s.count(x), list(s))),key=lambda x: len(x), reverse=True))


solution = Solution()
print(solution.frequencySort("tree"))
print(solution.frequencySort("Aabb"))
print(solution.frequencySort("cccaaa"))
