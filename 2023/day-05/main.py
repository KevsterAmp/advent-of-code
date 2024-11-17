def main():
    with open("input.in", "r") as file:
        lines = file.read().split("\n")

    seeds, maps = group_maps(lines)

    # turn seed to soil
    for sub_maps in maps:
        seeds = transform(seeds, sub_maps)   
    print(seeds) 
    print(min(seeds)) # type: ignore 

def group_maps(lines):
    seeds = list(map(lambda x: int(x), lines[0].replace("seeds:", "").strip().split()))
    flag = False
    maps = []
    sub_maps = []
    for line in lines[2:]:
        line = line.strip()
        if line != "" and line[0].isalpha():
            flag = True
        
        if line == "" or line == lines[-1]: 
            flag = False
            maps.append(sub_maps)
            sub_maps= []

        if flag and line[0].isdigit():
            sub_maps.append(line.split())    

    return seeds, maps


def transform(seeds, sub_maps):
    source_list = {}
    output = [None] * len(seeds)
    # get all maps and calculate
    for line in sub_maps:
        # output += [seed + diff for seed in seeds if source + range_n >= seed]
    #     source_list.update({x:y for x, y in zip(range(source, source + range_n), range(desti, desti + range_n))})
    # output = [source_list[seed] if seed in source_list else seed for seed in seeds]

        desti, source, range_n = list(map(lambda x: int(x), line))
        range_n -= 1
        diff = desti - source 
        for i, seed in enumerate(seeds):
            if seed and source + range_n >=seed and seed >= source:
                output[i] = seed+diff

    for i in range(len(seeds)): 
        if output[i] is None: 
            output[i] = seeds[i]

    return output

main()
