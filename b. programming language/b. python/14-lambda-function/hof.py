print("Higher Order Function")

# A higher-order function is a function that does one or both of the following:

# Takes one or more functions as arguments.
# Returns a function as its result.

################################

# map()
# Purpose: Applies a function to every item in an iterable (like a list) and returns an iterator.

numbers = [2, 4, 5, 6, 21, 34, 57, 91, 879]

squared = list(map(lambda num: num * num, numbers))
print(squared)

################################

# filter()
# Purpose: Filters elements from an iterable based on a function that returns True or False.

odd_nums = list(filter(lambda num: num % 2 != 0, numbers))
print(odd_nums)

################################

# reduce() (from functools module)
# Purpose: Applies a function cumulatively to the items of an iterable, reducing it to a single value.

from functools import reduce

product = reduce(lambda x, y: x + y, numbers)
# It basicly sum all numbers in the list.
print(product)

names = ["Ashish Khanagwal", "Dave Gray", "Sara Ito"]
char_count = reduce(lambda x, y: x + len(y), names, 0)
print(char_count)