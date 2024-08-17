print("String Formatting")

person = "Ashish"
coins = 4

print("\n" + person + " has " + str(coins) + " coins left")

print("\n%s has %s coins left" % (person, coins))

print(f"\n{person} has {coins} coins left")

print("\n{} has {} coins left".format(person, coins))

print("\n{1} has {0} coins left".format(coins, person))

print("\n{person} has {coins} coins left".format(person = person, coins = coins))

player = {"person": "Ashish", "coins": 5}
message = "\n{person} has {coins} coins left".format(**player)
print(message)

##############
# f-strings is tha way

message = f"\n{person} has {coins} coins left"
print(message)

message = f"\n{person} has {2*5} coins left"
print(message)

message = f"\n{person.lower()} has {2*5} coins left"
print(message)

message = f"\n{player['person']} has {2*5} coins left"
print(message)

##############
# we can pass formatting options too

num = 123456789
print(f"\n2.25 times of {num} is {2.25*num:,.2f}\n")

num_list = []
for i in range(1,11):
  message = 2.25*i
  num_list.append(message)
print(f"\n{num_list}")

for num in range(1, 11):
  print(f"\n{num} divided by 4.52 is {num / 4.52:.2%}\n")