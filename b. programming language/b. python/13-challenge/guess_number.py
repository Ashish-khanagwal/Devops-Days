import sys
import random

def guess_number(name='PlayerOne'):
  game_count = 0
  you_wins = 0
  python_wins = 0

  def play_guess_the_number():
    nonlocal name
    nonlocal you_wins
    nonlocal python_wins

    playerchoose = input(f"\n{name} guess which number i'm thinking of... 1, 2, or 3. \n")

    if playerchoose not in ["1", "2", "3"]:
      print(f"{name}, please choose: 1, 2, or 3")
      return play_guess_the_number()
    
    player = int(playerchoose)
    pythonchoose = random.choice("123")
    python = int(pythonchoose)

    def decide_winner(player, python):
      nonlocal name
      nonlocal you_wins 
      nonlocal python_wins
      if player == python:
        you_wins += 1
        print(f"{name}, you choses {player} 🙋‍♂️🙋‍♂️")
        print(f"I was thinking of {python} 😎😎")
        return f"{name}, you win! 🎉🎉"
      else:
        python_wins += 1
        print(f"{name}, you choses {player}🙋‍♂️🙋‍♂️")
        print(f"I was thinking of {python} 😎😎")
        return f"{name}, you lose! 😉😉"

    game_result = decide_winner(player, python)
    print(game_result)

    nonlocal game_count
    game_count += 1

    def game_data():
      nonlocal game_count
      nonlocal you_wins
      nonlocal name

      score = (you_wins / game_count) * 100
      print(f"\n{name}, your winning percentage is: {score}%")

    print(f"\nGame count: {game_count}")
    print(f"\n{name} wins: {you_wins}")
    print(f"\nPython wins: {python_wins}")
    game_data()

    print(f"\nplay again, {name}?")

    while True:
      playagain = input("\nY for Yes\nQ for Quit\n")
      if playagain.lower() not in ["y", "q"]:
        print("Invalid input. Please enter 'y' or 'q'")
        continue
      else:
        break

    if playagain.lower() == 'y':
      return play_guess_the_number()
    else:
      print("\nThank you for playing\n")
      if __name__ == "__main__":
        sys.exit(f"Bye {name}! 👋")
      else:
        return

  return play_guess_the_number

if __name__ == "__main__":
  import argparse

  parser = argparse.ArgumentParser(
    description="Provides a personalized gaming experience!"
  )

  parser.add_argument(
      "-n", "--name", type=str, required=True, help="The name of the person playing the game"
    )

  args = parser.parse_args()
  guess_my_number = guess_number(args.name)
  guess_my_number()