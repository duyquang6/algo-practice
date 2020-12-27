class Solution:
    def maskPII(self, s: str) -> str:
        res = []
        a_id = s.find('@')
        if a_id != -1:
            # email case
            res.append(s[0].lower() + 5*'*' + s[a_id - 1].lower())
            res.append('@')
            res.append(s[a_id + 1:].lower())
            return ''.join(res)

        # phone case
        s = [c for c in s if '0' <= c <= '9']
        has_country_code = len(s) > 10
        if has_country_code:
            res.append('+')
            cc_len = len(s) - 10
            res.extend(['*'*cc_len, '-'])
        res.extend([3*'*', '-', 3*'*', '-'])
        res.extend(s[-4:])
        return ''.join(res)
