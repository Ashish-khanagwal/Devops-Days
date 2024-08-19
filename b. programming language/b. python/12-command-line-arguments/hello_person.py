def hello(name, lang):
  greetings = {
    "English": "Hello!",
    "Hindi": "नमस्ते!",
    "Spanish": "Hola!"
  }
  msg = f"{greetings[lang]} {name}!"
  print(msg)

if __name__ == "__main__":
  import argparse

  parser = argparse.ArgumentParser(
    description="Provides a personal greeting!"
  )

  parser.add_argument(
    "-n", "--name", type=str, required=True, help="The name of the person to greet."
  )

  parser.add_argument(
    "-l", "--lang", type=str, choices=["English", "Hindi", "Spanish"], 
    required=True, help="The language to greet the person."
  )

  args = parser.parse_args()

  hello(args.name, args.lang)