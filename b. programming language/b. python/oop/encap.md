# Encapsulation

Encapsulation is one of the core principles of object-oriented programming (OOP). It refers to the bundling of data (attributes) and methods (functions) that operate on the data into a single unit, called a class. Encapsulation also involves restricting direct access to some of an object's components, which is a way of preventing accidental or unauthorized interference and misuse of the data.

Why Use Encapsulation?

- Data Protection: Encapsulation helps in protecting the internal state of an object from unintended or harmful changes
- Control: It allows the control of how the data is accessed and modified, often through getter and setter methods
- Flexibility: The internal implementation of a class can be changed without affecting the code that uses the class
- Modularity: By keeping the data and methods within a class, encapsulation promotes a modular design, making the code easier to understand and maintain

- self.name is a public attribute and can be accessed from outside the class.
- self.\_age is a protected attribute, indicating that it should not be accessed directly, though it's not enforced by Python.
- self.**bank_account is a private attribute, and Python uses name mangling to change its name to \_Person**bank_account to prevent direct access from outside the class.

The methods get_bank_account() and set_bank_account() are used to access and modify the private attribute \_\_bank_account. This approach ensures that the bank account value can only be changed through the set_bank_account method, which includes a check to prevent negative amounts
