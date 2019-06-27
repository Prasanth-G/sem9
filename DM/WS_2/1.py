
attributes_type = [(0, 'animal name', str), (1, 'hair', bool), (2, 'feathers', bool), (3, 'eggs', bool), (4, 'milk', bool), 
                (5, 'airborne', bool), (6, 'aquatic', bool), (7, 'predator', bool), (8, 'toothed', bool), (9, 'backbone', bool), 
                (10, 'breathes', bool), (11, 'venomous', bool), (12, 'fins', bool), (13, 'legs', int), (14, 'tail', bool), 
                (15, 'domestic', bool), (16, 'catsize', bool), (17, 'type',int)]

mapping = {
    bool : lambda data: 'True' if data == '1' else 'False',
}

def human_readable_representation(stream, sep=',', def_col_name="COLUMN_"):
    first_line = stream.readline().split(sep)
    header = []
    for i in range(len(first_line)):
        header.append(def_col_name + str(i))
    for attr in attributes_type:
        if len(header) > attr[0]:
            header[attr[0]] = attr[1]
    yield header

    for attr in attributes_type:
        if attr[2] in mapping:
            first_line[attr[0]] = mapping[attr[2]](first_line[attr[0]])
    yield first_line

    for line_str in stream:
        line = line_str.split(sep)
        for attr in attributes_type:
            if attr[2] in mapping:
                line[attr[0]] = mapping[attr[2]](line[attr[0]])
        yield line


if __name__ == '__main__':
    input_file = '''Z:\9 Sem\DM\WS_2\zoo.data'''
    output_file = '''Z:\9 Sem\DM\WS_2\zoo_refined.csv'''
    with open(input_file, 'rt') as ifile:
        with open(output_file, 'wt') as ofile:
            for line in human_readable_representation(ifile):
                ofile.write(",".join(line))
            

