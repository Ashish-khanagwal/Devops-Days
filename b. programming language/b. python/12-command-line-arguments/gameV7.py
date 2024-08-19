import sys
import random
from enum import Enum


def game(name='PlayerOne'):
  game_count = 0
  you_wins = 0
  python_wins = 0

  def play_rps():
    nonlocal name
    nonlocal you_wins
    nonlocal python_wins

    class RPS(Enum):
      ROCK = 1
      PAPER = 2
      SCISSORS = 3

    playerchoice = input(f"\n{name}, please enter....\n1 for Rock, \n2 for Paper, \n3 for Scissors:\n\n")

    if playerchoice not in ['1', '2', '3']:
      print(f"{name}, please enter choose: 1, 2, or 3")
      return play_rps()

    player = int(playerchoice)

    computerchoice = random.choice("123")
    computer = int(computerchoice)

    print(f"\n{name}, you chose {str(RPS(player)).replace('RPS.', "")}.")
    print(f"\nPython chose {str(RPS(computer)).replace('RPS.', "")}.")
    
    # Game logic
    def decide_winner(player, computer):
      nonlocal name
      nonlocal you_wins 
      nonlocal python_wins
      if player == computer:
        return "Draw!! 🤓🤓"
      elif player == 1 and computer == 3:
        you_wins += 1
        print()
        return f"{name}, you win! 🎉🎉"
      elif player == 2 and computer == 1:
        you_wins += 1
        return f"{name}, you win! 🎉🎉"
      elif player == 3 and computer == 2:
        you_wins += 1
        return f"{name}, you win! 🎉🎉"
      else:
        python_wins += 1
        return f"{name}, you lose! 😢😢"
    
    game_result = decide_winner(player, computer)
    print(game_result)

    nonlocal game_count
    game_count += 1

    print(f"\nGame count: {game_count}")
    print(f"\n{name} wins: {you_wins}")
    print(f"\nPython wins: {python_wins}")

    print(f"\nplay again, {name}?")

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
      sys.exit(f"Bye! {name}👋")

  return play_rps

if __name__ == "__main__":
  import argparse

  parser = argparse.ArgumentParser(
    description="Provides a personalized gaming experience!"
  )

  parser.add_argument(
    "-n", "--name", type=str, required=True, help="The name of the person playing the game"
  )

  args = parser.parse_args()
  rock_paper_scissiors = game(args.name)
  rock_paper_scissiors()