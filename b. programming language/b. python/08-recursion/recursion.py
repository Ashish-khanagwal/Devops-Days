print("Recursion")

def add_one(num):
  if num >= 9:
    return num + 1
  
  total = num + 1
  print(total)

  return add_one(total)

mytotal = add_one(0)
print(mytotal)


print("\n")

# Same behaiour using loops
print("Loop")

def add_one_loop(num):
    while num < 9:
        num += 1
        print(num)
    
    return num + 1

mytotal = add_one_loop(0)
print(mytotal)