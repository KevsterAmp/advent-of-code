from math import sqrt

    
def  main():
    with open("test.in", "r") as file:
        lines = file.read().replace("Time:", "").replace("Distance:", "").strip().split("\n")
    
    part2_timex = int("".join(lines[0].split()))
    part2_dist = int("".join(lines[1].split()))

    print(lines)
    timex = list(map(lambda x: int(x), lines[0].split()))  
    dist = list(map(lambda x: int(x), lines[1].split()))  

    # part 2
    x = sing_record(part2_timex,part2_dist)
    print(x)

    # part 1
    out = get_record(timex, dist)
    total = 1
    for x in out:
        total *= x

    print(total)
    

def sing_record(timex, dist):
    pb = dist
    count = 1
    prev = 0
    flag = False
    for s in range(1, timex, round(sqrt(timex))):
        curr_dist = (timex - s) * s

        if curr_dist > pb:
            gotcha = s - (round(sqrt(timex)) + 1)
            print(f"gotcha: {gotcha}")
            break
            
    for j in range(gotcha, timex):
        curr_dist = (timex - j) * j
        if curr_dist > pb and flag == False:
            flag = True    
        
        elif prev > curr_dist:
            break

        if flag:
            count += 1
        
        prev = curr_dist
                
    return count

    
def get_record(timex, dist):
    out = []
    count = 0
    for i, sec in enumerate(timex):
        pb = dist[i]

        # calculate time
        for s in range(1, sec):
            # dist = sec - range
            curr_dist = (sec - s) * s

            if curr_dist > pb:
                count += 1
        out.append(count)
        count = 0
    
    return out
        

main()
