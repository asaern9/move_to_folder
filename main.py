import json
import sys

from guessit import guessit


def main(filename):
    filename_comp = guessit(filename)
    try:
        filename_dict = {
            "title": filename_comp["title"],
            "season": filename_comp["season"],
            "episode": filename_comp["episode"],
        }
    except:
        return print("invalid file format")
    print(json.dumps(filename_dict))
    return True


if __name__ == "__main__":
    main(sys.argv[1])
