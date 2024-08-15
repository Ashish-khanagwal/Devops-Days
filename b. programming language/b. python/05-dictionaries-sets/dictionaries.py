print("Dictionaries")

band = {
  "vocals": "Plant",
  "guitar": "Page"
}

band2 = dict(vocals = "Plant", guitar = "Page")

print(band)
print(band2)

# Access items
print(band["vocals"])
print(band.get("guitar"))

# list all keys
print(band.keys())

# list all values
print(band.values())

# list of key/value pairs as tuples
print(band.items())

# verify a key exits
print("guitar" in band)
print("tree" in band)

# Change values
band["guitar"] = "Pagey"
band.update({"bass": "JPJ"})
print(band)

# remove items
print(band.pop("bass"))
print(band)

band["drums"] = "Boham"
print(band)

print(band.popitem()) # returns a tuple
print(band)

# Delete and Clear

band["drums"] = "Bonham"
del band["drums"]
print(band)

band2.clear()
print(band2)

del band2

# Copy

# band2 = band # create a reference
# print("Bad copy mechanism!!")

# band2["bass"] = "JPJ"
# print(band)

band2 = band.copy()
print("Good copy mechanism!!")

band2["bass"] = "JPJ"
print(band)
print(band2)

# Or use the dict() constructor function
band3 = dict(band)
print("Good copy mechanism")
print(band3)

# Nested Dictionary

member1 = {
  "firstname" : "Ashish",
  "lastname" : "khanagwal",
  "age": 23
}
member2 = {
  "firstname" : "John",
  "lastname" : "Doe",
  "age": 30
}

data = {
  "member1" : member1,
  "member2" : member2
}

print(data)
print(data["member1"]["firstname"])