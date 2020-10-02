from typing import List
import functools
import copy
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        @functools.lru_cache(None)
        def dfs(target):
            nonlocal candidates
            dup = set()
            for c in candidates:
                if c < target:
                    temp = dfs(target - c)
                    if not temp:
                        continue
                    for t in copy.deepcopy(temp):
                        t.append(c)
                        t.sort()
                        dup.add(tuple(t))
                if target == c:
                    dup.add((c,))
                    break
                if target < c:
                    break
            return [list(v) for v in dup]
        return dfs(target)