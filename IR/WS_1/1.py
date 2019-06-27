import re, os
from nltk.stem import PorterStemmer

tokenizer = re.compile('\W')
def getTokensStream(stream):
    for line in stream:
        yield [each for each in tokenizer.split(line) if each]

ps = PorterStemmer()
def stemmer(stream):
    for each in stream:
        yield ps.stem(each)

if __name__ == '__main__':
    filepath = os.path.join(os.getcwd(), r'sample.txt')
    with open(filepath) as file:
        tokenStream = getTokensStream(file)
        for line in tokenStream:
            for word in stemmer(line):
                print(word)

