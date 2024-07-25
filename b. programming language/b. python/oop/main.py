class Car:
  def __init__(self, brand, model):
    self.brand = brand
    self.model = model
  
  def full_name(self):
    return f"{self.brand} {self.model}"
  
# Inheritance

class ElectricCar(Car):
  def __init__(self, brand, model, battery_size):
    super().__init__(brand, model)
    self.battery_size = battery_size

my_electric_car = ElectricCar("Tesla", "Modle S", "80KWH")
print(my_electric_car.model)
print(my_electric_car.full_name())

my_car = Car("Toyota", "Corolla")
print(my_car.brand)
print(my_car.model)
print(my_car.full_name())