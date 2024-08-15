print("Functions")

def hello():
  print("hello world!")

hello()

def sum(num1=0, num2=0):
  if(type(num1) is not int and type(num2) is not int):
    return
  return num1 + num2

total = sum(5, 3)
print(total)

# Variable-Length Arguments

def name_fruits(*args):
  for fruit in args:
    print(fruit)
  print(type(args))

name_fruits("apple", "banana", "cherry")

def multi_named_items(**kwargs):
  for key, value in kwargs.items():
    print(f"{key}: {value}")
  print(type(kwargs))

multi_named_items(fruit="apple", color="red")
  