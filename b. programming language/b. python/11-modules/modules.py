import math
import sys
import random # --> also, we can do "import random as <name> e.g -> rdm" and then use this rdm in code as "rdm.choice("123")"
from math import pi # --> we can import pi from math module
import kansas
from gamev6 import rock_paper_scissiors

print(round(math.pi))

result = math.sqrt(math.pi)
formatted_result = f"{result:.2f}"

print(formatted_result)

print(random.choice("123"))

print(pi)

# num1 = input("Enter first number: ")
# num2 = input("Enter second number: ")
# print(type(num1))
# maximum = max(num1, num2) --> String type
# print(maximum)

# int_num1 = int(num1)
# int_num2 = int(num2)
# print(type(int_num1))
# maximum = max(int_num1, int_num2) --> Int type
# print(maximum)

# print(dir(random)) # --> to know the functions of specific module in dictionary format

# for item in dir(random):
#   print(item)

print(kansas.capital)
kansas.randomfunfact()

print(kansas.__name__) # --> Output will be "kansas"
print(__name__) # --> Output will be "__main__", this is the name of the modules.py file

rock_paper_scissiors() # --> Calling the game function from gamev6.py file

sys.exit()