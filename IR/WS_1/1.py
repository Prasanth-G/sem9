import re, os
from nltk.stem import PorterStemmer
from os import listdir
from os.path import isfile, join
from collections import defaultdict
from pprint import pprint

home_dir = "Z:\\data"
onlyfiles = [join(home_dir, f) for f in listdir(home_dir) if isfile(join(home_dir, f)) and f.endswith(".txt")]

tokenizer = re.compile('\W')
def getTokensStream(stream):
    for line in stream:
        yield [each for each in tokenizer.split(line) if each]

ps = PorterStemmer()
def stemmer(stream):
    for each in stream:
        yield ps.stem(each)

if __name__ == '__main__':
    index = 0
    file_index = []
    file_terms = defaultdict(lambda : set())
    term_index = defaultdict(lambda : set())
    for filepath in onlyfiles:
        #filepath = os.path.join(os.getcwd(), file_)
        print(filepath)
        with open(filepath) as file:
            if filepath not in file_index:
                file_index.append(filepath)
            tokenStream = getTokensStream(file)
            for line in tokenStream:
                for word in stemmer(line):
                    term_index[word].add(index)
                    file_terms[index].add(word)
            index = index + 1
    pprint(file_index)
    pprint(file_terms)
    pprint(term_index)