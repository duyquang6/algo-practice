class Solution:
    def reformatNumber(self, number: str) -> str:
        filter_input = [c for c in number if c not in [' ', '-']]        
        rem = len(filter_input)
        res = []
        while rem > 4:
            index = len(filter_input) - rem
            res.extend(filter_input[index: index + 3])
            rem -= 3
            res.append('-')        
        if rem == 3:        
            res.extend(filter_input[-rem:])
        else:
            while rem > 0:                
                print(rem)
                if -rem + 2 == 0:
                    res.extend(filter_input[-rem:])
                else:
                    res.extend(filter_input[-rem:-rem + 2])
                res.append('-')
                rem -= 2
            res = res[:-1]
        
        return ''.join(res)

print(Solution().reformatNumber("1-23-45 6"))
print(Solution().reformatNumber("123 4-567"))
print(Solution().reformatNumber("123 4-5678"))
print(Solution().reformatNumber("--17-5 229 35-39475 "))