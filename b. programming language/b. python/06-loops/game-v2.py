import sys
import random
from enum import Enum

class RPS(Enum):
  ROCK = 1
  PAPER = 2
  SCISSORS = 3

playagain = True

while playagain:

  playerchoice = input("\nEnter....\n1 for Rock, \n2 for Paper, \n3 for Scissors:\n\n")
  player = int(playerchoice)

  if player < 1 or player > 3:
    sys.exit("You must choose: 1, 2, or 3")

  computerchoice = random.choice("123")
  computer = int(computerchoice)

  # print("You chose " + str(RPS(player)).replace('RPS.', ""))
  # print("Python chose " + str(RPS(computer)).replace('RPS.', ""))

  print("\nYou chose " + playerchoice + ".")
  print("\nPython chose " + computerchoice + ".\n")
  

  # Game logic

  if player == computer:
    print("Draw!! 🤓🤓")
  elif player == 1 and computer == 3:
    print("You win! 🎉🎉")
  elif player == 2 and computer == 1:
    print("You win! 🎉🎉")
  elif player == 3 and computer == 2:
    print("You win! 🎉🎉")
  else:
    print("You lose! 😢😢")

  playagain = input("\nDo you want to play again? (y/n): ")
  if playagain.lower() == 'y':
    continue
  else:
    print("\nThank you for playing\n")
    playagain = False

sys.exit("Bye! 👋")