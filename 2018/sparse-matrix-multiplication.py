
def multiply(A, B):
	if len(A) == 0 or len(B) == 0 or len(A[0]) == 0 or len(B[0]) == 0:
		return [[]]

	# build sparse versions of a & b: []int{x, y, val}
	# key: a is sorted by y <- not really; doesn't matter
	sa = []
	sb = []
	for y in range(0, len(A)):
		for x in range(0, len(A[y])):
			if A[y][x] != 0:
				sa.append([x, y, A[y][x]])

	for y in range(0, len(B)):
		for x in range(0, len(B[y])):
			if B[y][x] != 0:
				sb.append([x, y, B[y][x]])

	# initialize result matrix
	res = [[0] * len(B[0]) for i in range(len(A))]

	# fill result matrix of cardinality y(a)*x(b)
	for A in sa:
		for B in sb:
			if A[0] == B[1]:
				res[A[1]][B[0]] += A[2] * B[2]

	return res

if __name__ == "__main__":
	res = multiply(
		[
			[1, 0, 0],
			[-1, 0, 3],
		],
		[
			[7, 0, 0],
			[0, 0, 0],
			[0, 0, 1],
		],
	)
	# res = multiply(
	# 	[
	# 		[1, -5],
	# 	],
	# 	[
	# 		[12],
	# 		[-1],
	# 	],
	# )
	for r in res:
		print r
