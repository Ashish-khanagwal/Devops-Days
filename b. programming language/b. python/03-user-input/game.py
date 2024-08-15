import sys
import random
from enum import Enum

class RPS(Enum):
  ROCK = 1
  PAPER = 2
  SCISSORS = 3

playerchoice = input("Enter....\n1 for Rock, \n2 for Paper, \n3 for Scissors:\n\n")
player = int(playerchoice)

if player < 1 or player > 3:
  sys.exit("You must choose: 1, 2, or 3")

computerchoice = random.choice("123")
computer = int(computerchoice)

print("")
print("You chose " + str(RPS(player)).replace('RPS.', ""))
print("Python chose " + str(RPS(computer)).replace('RPS.', ""))
print("")

# Game logic

if player == computer:
  sys.exit("Draw!! 🤓🤓")
elif player == 1 and computer == 3:
  sys.exit("You win! 🎉🎉")
elif player == 2 and computer == 1:
  sys.exit("You win! 🎉🎉")
elif player == 3 and computer == 2:
  sys.exit("You win! 🎉🎉")
else:
  sys.exit("You lose! 😢😢")
