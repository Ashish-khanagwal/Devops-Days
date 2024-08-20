print("Lambda Function")

# lambda arguments: expression
# Lambda functions can have any number of arguments but only one expression.

def squared(num):
  return num **2
print(squared(4))

# Using lambda function
# A lambda function in Python is a small, anonymous function defined using the lambda keyword

squared = lambda num : num * num
print(squared(4))

sum = lambda x, y: x + y

print(sum(45,54))

################################

def functionBuilder(x):
  return lambda num : num + x

addTen = functionBuilder(10)
addTwenty = functionBuilder(20)

print(addTen(5))
print(addTwenty(5))

################################