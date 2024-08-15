import math
print("String Methods")

first = "Ashish"
last = "Khanagwal"

multiline = """ Hey, how are you?
                                all good?"""

print(first)
print(first.lower())
print(first.upper())

print("\n\n")
print(multiline)
print(multiline.title())
print(multiline.replace("good", "ok"))

print("\n\n")

title = "menu".upper()
print(title.center(20, "-"))
print("Coffee".ljust(16, ".") + "$5".rjust(4))
print("Tea".ljust(16, ".") + "$3".rjust(4))
print("Muffin".ljust(16, ".") + "$7".rjust(4))

print("\n\n")

# String index values
print(first[1])
print(first[-1])
print(first[1:-1])
print(first[1:])

# Some methods return boolean data
print(first.lower().startswith('a'))
print(first.endswith('a'))

print('\n')

# Boolean data type
print("Boolean Data Type")
myvalue = False
print(type(myvalue))
x = bool(False)
print(x)
print(type(x))
print(isinstance(x, bool))

print("\n")

# Numeric data types
print("Numeric Data Types")

# Integer Data types
print("Integer Data types")
num1 = 80
num2 = int(100)
print(type(num2))
print(num2)
print(isinstance(num2, int))

print("\n")

# Float Data Types
print("Float Data Types")
gpa = 8.14
print(type(gpa))

print("\n")

# Complex Data Types
print("Complex Data Types")
c = 3 + 5j
print(type(c))
print(c.real)
print(c.imag)

print("\n")

# built in functions for numbers
print("built in functions")
print(abs(gpa))
print(abs(gpa * -1))
print(round(gpa))
print(round(gpa, 1))

print("\n")

# Math module
print("Math module")
a = 19
b = 25

print(max(a, b))
print(math.pi)
print(math.ceil(gpa))
print(math.floor(gpa))

print("\n")

# Casting a String to a number
print("Casting a String to a number")
zipcode = "124001"
print(type(zipcode))
zip_code = int(zipcode)
print(type(zip_code))

# Error if you attempt to cast incorrect data
# zip_value = int("India")