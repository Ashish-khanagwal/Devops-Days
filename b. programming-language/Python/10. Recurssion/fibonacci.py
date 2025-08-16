line = "********************"
print(line)
print("Fibonacci".center(20))
print(line)

# fibonacci series - 0, 1, 1, 2, 3, 5, 8, 13, 21


def fibonacci(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fibonacci(n - 1) + fibonacci(n - 2)


n = int(input("Enter the number you want fibonacci number for: "))
print(fibonacci(n))
