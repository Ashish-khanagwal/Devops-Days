import sys
import random
from enum import Enum


def game():
  game_count = 0
  you_wins = 0
  python_wins = 0

  def play_rps():

    nonlocal you_wins
    nonlocal python_wins

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

    print(f"\nYou chose {str(RPS(player)).replace('RPS.', "")}.")
    print(f"\nPython chose {str(RPS(computer)).replace('RPS.', "")}.")
    
    # Game logic
    def decide_winner(player, computer):
      nonlocal you_wins 
      nonlocal python_wins
      if player == computer:
        return "Draw!! 🤓🤓"
      elif player == 1 and computer == 3:
        you_wins += 1
        print()
        return "You win! 🎉🎉"
      elif player == 2 and computer == 1:
        you_wins += 1
        return "You win! 🎉🎉"
      elif player == 3 and computer == 2:
        you_wins += 1
        return "You win! 🎉🎉"
      else:
        python_wins += 1
        return "You lose! 😢😢"
    
    game_result = decide_winner(player, computer)
    print(game_result)

    nonlocal game_count
    game_count += 1

    print(f"\nGame count: {str(game_count)}")
    print(f"\nYou wins: {str(you_wins)}")
    print(f"\nPython wins: {str(python_wins)}")

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

  return play_rps

play = game()
play()