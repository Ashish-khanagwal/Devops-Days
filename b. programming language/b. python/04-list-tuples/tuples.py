mytuple = tuple((32, "Ashish", True))

anothertuple = (1, 4, 62, 87, 4, 4)

print(mytuple)
print(type(mytuple))
print(type(anothertuple))

newlist = list(mytuple)
newlist.append("Khanagwal")
newtuple = tuple(newlist)
print(newtuple)

(one, *two, hey) = anothertuple
print(one)
print(two)
print(hey)

print(anothertuple.count(4))