def replaceWords(dictionary: list[str], sentence: str):
    sentence = sentence.split()
    length = len(sentence)

    for i in range(length):
        rootAppended = False
        for key in dictionary:
            if sentence[i].startswith(key):
                if not rootAppended:
                    sentence[i] = key
                    rootAppended = True
                elif rootAppended and len(sentence[i]) > len(key):
                    sentence[i] = key

    return ' '.join(sentence)


print(replaceWords(['cat', 'bat', 'rat'],
      "the cattle was rattled by the battery"))
