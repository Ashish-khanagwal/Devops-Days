# **init** Method is known as Constructor

## Explanation

Special Method:

The **init** method is part of Python's special methods, also known as "magic methods" or "dunder methods" (short for "double underscore"). These methods have special significance and are typically invoked in specific situations.
Constructor:

As a constructor, **init** sets up the initial state of the object by initializing its attributes. This is crucial because it ensures that each object starts with a valid state and that its attributes are properly configured.
Syntax:

The **init** method is defined using the following syntax:
python
Copy code

```
class ClassName:
    def __init__(self, parameters):
        # Initialization code
```

It must be defined with the first parameter named self, which refers to the instance being created. Following self, additional parameters can be specified to accept data that will be used to initialize the object.
Initialization of Attributes:

Inside the **init** method, instance attributes are typically assigned using the self keyword. These attributes are unique to each instance of the class.
No Return Value:

The **init** method does not return a value (it implicitly returns None). It is solely used for initialization and setup.

# self

## Explanation

In the **init** method, self refers to the newly created object. It allows you to set the attributes (brand, model) for that specific instance. For example, when you do self.brand = brand, you are setting the brand attribute of the specific instance to the value passed in the constructor.

In the display_info method, self refers to the instance of the Car class from which the method is called. It allows the method to access the instance's attributes (self.year, self.brand, self.model) and print them.

# Inheritance

## Explanation

Inheritance is a fundamental concept in object-oriented programming that allows a new class to inherit the properties and behaviors (attributes and methods) of an existing class. The new class, called the derived or child class, extends the functionality of the existing class, known as the base or parent class.

```
# Parent class
class Vehicle:
    def __init__(self, brand, model):
        self.brand = brand
        self.model = model

    def start(self):
        print(f"{self.brand} {self.model} is starting.")

# Child class
class Car(Vehicle):
    def __init__(self, brand, model, num_doors):
        super().__init__(brand, model)  # Call the constructor of the parent class
        self.num_doors = num_doors

    def display_info(self):
        print(f"Car: {self.brand} {self.model}, Doors: {self.num_doors}")

# Creating an object of the Child class
my_car = Car("Toyota", "Camry", 4)

# Accessing methods from both Parent and Child classes
my_car.start()          # Output: Toyota Camry is starting.
my_car.display_info()   # Output: Car: Toyota Camry, Doors: 4
```

- Parent Class (Vehicle): This class has an **init** method to initialize brand and model attributes, and a start method.
- Child Class (Car): This class inherits from Vehicle and adds an additional attribute num_doors. It also defines a method display_info to display details about the car.
- The super().**init**(brand, model) call in the Car class constructor invokes the parent class's constructor, initializing the brand and model attributes using the parent class's logic.
