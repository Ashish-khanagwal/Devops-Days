# Polymorphism

Polymorphism is another fundamental concept in object-oriented programming (OOP) that allows objects of different classes to be treated as objects of a common superclass. It enables a single interface to represent different underlying data types or classes. Polymorphism is mainly used to allow flexibility and the ability to define generic interfaces to different types of objects

Types of Polymorphism

- Compile-time Polymorphism (Static Polymorphism): This occurs when the method to be invoked is determined at compile time. This is typically achieved through method overloading, where multiple methods have the same name but differ in the type or number of parameters

- Runtime Polymorphism (Dynamic Polymorphism): This occurs when the method to be invoked is determined at runtime. It is typically achieved through method overriding, where a subclass provides a specific implementation of a method that is already defined in its superclass

1. Method Overloading in Python:
   Python does not support method overloading in the traditional sense because you cannot have multiple methods with the same name differing only in the number or type of parameters. However, you can achieve a similar effect using default parameters or variable-length arguments

2. Method Overriding in Python:
   Method overriding allows a subclass to provide a specific implementation of a method that is already defined in its superclass. The subclass's version of the method overrides the superclass's version

## Polymorphism through Duck Typing in Python

Duck Typing in Python is a concept related to polymorphism that focuses on what an object can do rather than what type it is. The term comes from the saying, "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck." In other words, if an object behaves like a certain type (e.g., has certain methods), it can be treated as that type, regardless of what it actually is
