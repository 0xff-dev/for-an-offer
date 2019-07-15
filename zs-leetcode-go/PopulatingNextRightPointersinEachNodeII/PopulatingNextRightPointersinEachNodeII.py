class Solution:
    def connect(self, root: 'Node') -> 'Node':
        if root is None:
            return None
        queue = [root]
        while queue:
            length = len(queue)
            for i in range(length):
                obj = queue.pop(0)
                if i != length-1:
                    obj.next = queue[0]
                else:
                    obj.next = None
                if obj.left:
                    queue.append(obj.left)
                if obj.right:
                    queue.append(obj.right)
        return root