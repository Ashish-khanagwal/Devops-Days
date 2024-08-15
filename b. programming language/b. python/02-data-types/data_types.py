print("Data Types")

# Literal assignment
# String
first = 'Ashish'
last = 'Khanagwal'

print(type(first))
print(type(first) == str)
print(isinstance(first, str))

# Constructor function
pizza = str("Pepperoni")
print(type(pizza))
print(type(pizza) == str)
print(isinstance(pizza, str))

# Concatenation
fullname = first + " " + last
print(fullname)

# casting a number to a string
decade = str(2000)
print(type(decade))
print(decade)

statement = "I like the music from " + decade + "s."
print(statement)

# Multiline string
multiline = """
This is a multiline 
                  string"""
print(multiline)
print('\n\n')

# Escaping speical characters
sentence = 'I\'m back at work!\t Hey!\n\nWhere\'s this at\\located'
print(sentence)