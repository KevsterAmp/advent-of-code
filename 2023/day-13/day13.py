with open("input.in", "r") as file:
    line = file.read()


def part1():
    def calibrate(line):
        # split the lines per pattern
        lines = line.split("\n\n")
        return [line.split("\n") for line in lines]

    # get the horizontal or vertical match
    def horizontal_check(horizons):
        for i, line in enumerate(horizons):
            # print(f"{line = }")
            if i == 0:
                continue
            
            else:
                past = horizons[i-1]  
            
                if line == past:
                    # check if perfect reflection
                    match = True
                    bottom = 1
                    while True:
                        if i+bottom >= len(horizons) or i-bottom+1 < 0:
                            break
                        elif horizons[i+bottom] == horizons[i-bottom-1]:
                            bottom+=1
                            continue
                        else: 
                            match = False
                            break
                        
                    
                    if match == True:
                        return i
                    else: 
                        continue
        
        return None

    def vertical2horizon(box):
        # calibrate to vertical list
        verticals = [ [] for _ in range(len(box[0]))]
        for line in box:
            for i in range(len(box[0])):
                verticals[i].append(line[i])
        return ["".join(x) for x in verticals]
            
    output = 0
    boxes = calibrate(line)
    for box in boxes:
        num = horizontal_check(box)
        #print(f"{box = }")
        # print(f"{num = }")
        if num is None:
            calibrated = vertical2horizon(box) 
            num = horizontal_check(calibrated)
            if num == None:
                print(f"{boxes.index(box)= }")
                print("ERROR")
            else:
               output += num 
        else:
            output += 100*num

    
    return output


def part2():
    pass


if __name__ in "__main__":
    print(f"{part1() = }")
    print(f"{part2() = }")
    