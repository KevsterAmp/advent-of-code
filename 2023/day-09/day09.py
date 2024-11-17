with open("input.in", "r") as file:
    lines = file.read().split("\n")

def calculate_diff(numbers):
    max_i = len(numbers)
    output = []
    for i, n in enumerate(numbers):
        if i == max_i-1: break
        diff = numbers[i+1] - numbers[i]
        output.append(diff)
    
    return output

# split numbers on lines 
def calibrate(lines):
    return [list(map(lambda x: int(x),line.split())) for line in lines]

def get_histories(line):   
# while list is not equal to each other
    curr_list = line
    all_histories = [curr_list]
    while len(set(curr_list)) != 1:
        # get the differences of each of the consecutive lists
        curr_list = calculate_diff(curr_list)
        all_histories.append(curr_list)

    # return as list of lists
    return all_histories


def part1():
    # function to calculate the total score of the top list
    def get_top_score(histories):
        # calculate top score starting from -1 of list
        histories.reverse()
        for i, numbers in enumerate(histories):
            # if all elements in numbers is equal, the top score is the same
            if i == 0: histories[i].append(numbers[0])
    
            # else, the topscore is by numbers[-1] + histories[i+1][-1] 
            else:
                end = numbers[-1] + histories[i-1][-1]
                histories[i].append(end)

        return histories[-1][-1]

    # first, calibrate lines
    calibrated = calibrate(lines)
    output_sum = 0
    # per line 
    for line in calibrated:
        histories = get_histories(line)
        test = get_top_score(histories)
        output_sum += test

    return output_sum 


def part2():    
    # function to get_bottom_score
    def get_bottom_score(histories):
        # calculate top score starting from -1 of list
        histories.reverse()
        for i, numbers in enumerate(histories):
            # if all elements in numbers is equal, the bottom score is the same
            if i == 0: histories[i].insert(0, numbers[0])
    
            # else, the bottom score is by numbers[-1] + histories[i+1][-1] 
            else:
                end = numbers[0] - histories[i-1][0]
                histories[i].insert(0, end)

        return histories[-1][0]

    # first, calibrate lines
    calibrated = calibrate(lines)
    output_sum = 0
    # per line 
    for line in calibrated:
        # print(f"line: {line}")
        histories = get_histories(line)
        test = get_bottom_score(histories)
        # print(f"histories: {histories}")
        # print(f"bottom score: {test}")
        # print()
        output_sum += test

    return output_sum 


if __name__ in "__main__":
    print(part1())
    print(part2())