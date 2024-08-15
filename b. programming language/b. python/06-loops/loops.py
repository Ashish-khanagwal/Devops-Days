print("Loops")

# While Loop
print("While Loop")

value = 1
while value <= 10:
  print(value)
  if value == 5:
    break
  value += 1

print("...............")

while value <= 10:
  value += 1
  if value == 5:
    continue
  print(value)
else:
  print("value is now equal to " + str(value))

print("\n")

# For Loop
print("For Loop")
names = ["Ashish", "Dave", "John"]
for name in names:
  print(name)

x = []
for char in "Mississippi":
  x.append(char)
print(x)

for name in names:
  if name == "Dave":
    break
  print(name)

for name in names:
  if name == "Dave":
    continue
  print(name)

for i in range(5):
  print(i)

for i in range(1, 5):
  print(i)

for i in range(5, 101, 5):
  print(i)
else:
  print("Glad that\'s over!😤")

i = []
for num in range(5, 101, 5):
  i.append(num)
print(i)

names = ["Ashish", "Dave", "John"]
actions = ["Codes", "Eats", "Sleeps"]

info = []
for name in names:
  for action in actions:
    info.append(name + " " + action)
print(info)

info = []
for action in actions:
  for name in names:
    info.append(name + " " + action)
print(info)