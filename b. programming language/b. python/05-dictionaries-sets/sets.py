print("Sets")

nums = {1, 2, 3, 4}
nums2 = set((1, 2, 3, 4))

print(nums)
print(nums2)
print(type(nums))
print(len(nums))

# No duplicates are allowed
num3 = {1, 2, 3, 4, 5, 2}
print(num3)

# True is a dupe of 1 and Fasle is a dupe of 0
num4 = {1, 2, 3, 4, True, False}
print(num4)

# Check if value is in set
print(1 in num4)
print(0 in num4)

# we can not refert to any element in the set with an index position or key.

# Add a new element to a set
num4.add(5)
print(num4)

# Add elements from one set to another
morenums = {6, 7, 8}
num4.update(morenums)
print(num4)

# we can use update with dictionaries, sets, tuples, and list too.

# merge two sets to create a new set
one = {1, 2, 3}
two = {4, 5, 6}
mynewset = one.union(two)
print(mynewset)

# Keep only the duplicates
one = {1, 2, 3}
two = {2, 3, 4}
one.intersection_update(two)
print(one)

# Keep everything except the duplicates
one = {1, 2, 3}
two = {2, 3, 4}
one.symmetric_difference_update(two)
print(one)