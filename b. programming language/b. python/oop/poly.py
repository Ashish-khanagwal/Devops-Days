# Method Overloading

class MathOperations:
  def add(self, a, b, c=0):
    return a + b + c
  
math = MathOperations()
print(math.add(1, 2))
print(math.add(1, 4, 5))

# Method Overriding

class Animal:
  def sound(self):
    print("Animal makes a sound")

class Dog(Animal):
  def sound(self):
    print("Dog barks")
  
class Cat(Animal):
  def sound(self):
    print("Cat meows")

dog = Dog()
cat = Cat()

dog.sound()  # Output: Dog barks
cat.sound()  # Output: Cat meows

# Duck Typing

class Bird:
  def fly(self):
    print("This bird can fly")

class Airplane:
  def fly(self):
    print("This airplane can fly")

class Kite:
  def fly(self):
    print("This kite can fly")

bird = Bird()
airplane = Airplane()
kite = Kite()

bird.fly()
airplane.fly()
kite.fly()  