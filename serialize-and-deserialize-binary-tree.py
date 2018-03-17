# This solution works but is not performant for some Python-bs reason I can't understand. Please look at my go solution which is O(n)

class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

def readNum(s):
	if len(s) == 0:
		return None, s
	sign = 1
	bs = []
	ok = False
	i = 0
	if s[i] == '-':
		i = 1
		sign =-1
	for i in range(i, len(s)):
		if s[i] == ',':
			i+=1
			break
		ok = True
		bs.append(s[i])
	d = 1
	r = 0
	for j in range(len(bs)-1, -1, -1):
		r+= int(bs[j])*d
		d*=10
	if not ok:
		return None, s[i:]
	return TreeNode(r*sign), s[i:]


def bfs(ts, res):
	nts = []
	for t in ts:
		if t != None:
			nts.append(t.left)
			nts.append(t.right)
			res.append('{}'.format(t.val))
		res.append(',')
	if len(nts) > 0:
		bfs(nts, res)
	i = 0
	while len(res)>0 and res[len(res)-1] == ',':
		res.pop()
	res.append(',')

class Codec:

	def serialize(self, root):
		"""Encodes a tree to a single string.
		:type root: TreeNode
		:rtype: str
		"""
		if root == None:
			return ""
		bs = []
		bfs([root], bs)
		return ''.join(bs)

	def deserialize(self, data):
		"""Decodes your encoded data to tree.
		:type data: str
		:rtype: TreeNode
		"""
		root, data = readNum(data)
		if root == None:
			return None
		nodes = [root]
		while len(nodes) > 0 and len(data) > 0:
			nns = []
			for node in nodes:
				node.left, data = readNum(data)
				node.right, data = readNum(data)
				if node.left != None:
					nns.append(node.left)
				if node.right != None:
					nns.append(node.right)
			nodes = nns
		return root

print serialize(deserialize("-1,2,3,,4,5,,6,"))
print serialize(deserialize("1,"))
print serialize(deserialize(""))
