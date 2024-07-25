class Person:
  def __init__(self, name, age):
    self.name = name  # public attribute
    self._age = age  # protected attribute
    self.__bank_amount = 10000 # private attribute

  def display_info(self):
    print(f"{self.name} {self._age}")

  def get_bank_amount(self):
    return self.__bank_amount
  
  def set_bank_amount(self, amount):
    if amount >= 0:
      self.__bank_amount = amount
    else:
      print("Invalid amount")

person = Person("Ashish", 3)
print(person.name)
print(person._age)

# Accessing private attribute (not directly possible)
# print(person.__bank_account)  # This will raise an AttributeError

# Using getter and setter methods to access private attribute
print(person.get_bank_amount())
person.set_bank_amount(20000)
print(person.get_bank_amount())