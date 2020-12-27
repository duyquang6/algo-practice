class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        up = 0
        
        while students and up < len(students):
            s = students.pop(0)
            if s == sandwiches[0]:
                sandwiches.pop(0)
                up = 0
                continue
            up += 1
            students.append(s)
            
        return len(students)