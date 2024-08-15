users = ["Ashish", "John", "Homelendar"]

name = ["Sara", "Ashish"]

data = []

print("Ashish" in users)
print("Ashish" in name)
print("Ashish" in data)

print(users[0])
print(users[2])

print(users[-1])
print(users[0:2])
print(users[-3:-1])
print(users[:-1])

print(len(users))
print(users.index("John"))

users.append("Butcher")
print(users)

users += ["Hughie"]
print(users)

users.extend(["starlight", "MJ"])
print(users)

print(users.index("starlight"))
users.insert(5, "A-train")
print(users)

users[2:2] = ["Eddie", "Alex"]
print(users)

users[1:3] = ["Webvewer", "Webdeveloper"]
print(users)

users.remove('Webdeveloper')
print(users)

print(users.pop())
print(users)

del users[0]
print(users)

users[1:2] = ["ashish"] # lowercase comes after the uppercase
users.sort()
print(users)

users.sort(key=str.lower) # this won't work if different data types are present in list
print(users)

nums = [4, 23, 2, 65, 1, 76, 5]
nums.reverse()
print(nums)

nums.sort()
print(nums)

# nums.sort(reverse=True)
# print(nums)
# alternative to above method, without modifying the original list

print(sorted(nums, reverse=True))
print(nums)

# Copying the list

numscopy = nums.copy()
mycopy = list(nums)
mynums = nums[:]

print(numscopy)
print(sorted(mycopy, reverse=True))
print(mynums)

mylist = list([1, "Ashish", True])
print(mylist)