import sys
import random
from enum import Enum


def play_rps():
  class RPS(Enum):
    ROCK = 1
    PAPER = 2
    SCISSORS = 3

  playerchoice = input("\nEnter....\n1 for Rock, \n2 for Paper, \n3 for Scissors:\n\n")

  if playerchoice not in ['1', '2', '3']:
    print("You must choose: 1, 2, or 3")
    return play_rps()

  player = int(playerchoice)

  computerchoice = random.choice("123")
  computer = int(computerchoice)

  print("You chose " + str(RPS(player)).replace('RPS.', ""))
  print("Python chose " + str(RPS(computer)).replace('RPS.', ""))
  
  # Game logic
  def decide_winner(player, computer):
    if player == computer:
      return "Draw!! 🤓🤓"
    elif player == 1 and computer == 3:
      return "You win! 🎉🎉"
    elif player == 2 and computer == 1:
      return "You win! 🎉🎉"
    elif player == 3 and computer == 2:
      return "You win! 🎉🎉"
    else:
      return "You lose! 😢😢"
  
  game_result = decide_winner(player, computer)
  print(game_result)

  print("\nplay again?")

  while True:
    playagain = input("\nY for Yes\nQ for Quit\n")
    if playagain.lower() not in ["y", "q"]:
      print("Invalid input. Please enter 'y' or 'q'")
      continue
    else:
      break

  if playagain.lower() == 'y':
    return play_rps()
  else:
    print("\nThank you for playing\n")
    sys.exit("Bye! 👋")

play_rps()