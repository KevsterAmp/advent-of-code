from itertools import product


with open("input.in", "r") as file:
    lines = file.read().split("\n")


def part1():
    def calibrate(lines):
        # calibrate lines and split them to records and guide
        output = []
        for line in lines:
            records, guide = line.split()
            clean_guide = list(map(lambda x: int(x), guide.split(",")))
            output.append([records, clean_guide])

        return output

    def find_arrangements(line):
        arr_num = 0
        record = line[0]
        guide = line[1]
        wildcard_num = 0
        for x in record:
            if x == "?":
                wildcard_num += 1

        # get the algo
        algos = create_algo(wildcard_num)

        for algo in algos:
            temp = record
            # substitute the ? on the record on the algo
            for element in algo:
                temp = temp.replace("?", element, 1)

            # split the record by .
            temp = temp.split(".")

            # turn the combinations into number
            algo_num = hash2num(temp)
            if algo_num == guide:
                arr_num += 1
            
        return arr_num
                

    # create a list that changes every time that is the length of wild_cardnum
    def create_algo(wildcard_num):
        symbols = ["#", "."]
        if wildcard_num < 1:
            return []

        # create an algorithm that changes the question marks as . or # 
        combinations = list(product(symbols, repeat=wildcard_num))
        return ["".join(combination) for combination in combinations]


    def hash2num(hashes):
        # turn the hashes [##, ###, #] to [2, 3, 1]
        output = []
        for x in hashes:
            if x == "":
                continue
            elif len(x) == 1:
                output.append(1)
            else:
                output.append(len(x))
        
        return output

    # calibrate
    calibrated = calibrate(lines)
    total = 0
    for i, x in enumerate(calibrated):
        total +=  find_arrangements(x)
       
    return total


def part2():
    pass


if __name__ in "__main__":
    print(f"{part1() = }")
    print(f"{part2() = }")
