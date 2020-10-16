
# https://leetcode.com/problems/thousand-separator/
class Solution:
    def thousandSeparator(self, n: int) -> str:
        if n == 0:
            return '0'
        ans = ''
        i = 0
        while n > 0:
            num = n % 10
            n //= 10
            if i % 3 == 0 and i != 0:
                ans+='.'
            ans += str(num)
            i+=1
        return ans[::-1]