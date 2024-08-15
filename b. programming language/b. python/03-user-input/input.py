name = input("Enter your name: ")
print(f"hello! {name}")

age = int(input("Enter your age: "))
print(f"your age is {age}")

# Taking multiple inputs separated by spaces
x, y = input("Enter two numbers separated by a space: ").split()
print(f"First number: {x}, Second number: {y}")