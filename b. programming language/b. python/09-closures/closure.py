print("Closures")

def parent_function(person, coins):

  def play_game():
    nonlocal coins
    coins -= 1

    if coins > 1:
      print(f"{person} has {coins} coins left")
    elif coins == 1:
      print(f"{person} has {coins} coin left")
    else:
      print(f"{person} has no more coins left")

  return play_game

Ashish = parent_function("Ashish", 5)
Jimmy = parent_function("Jimmy", 3)

Ashish()
Ashish()
Ashish()
Jimmy()
Jimmy()
Jimmy()
Jimmy()
