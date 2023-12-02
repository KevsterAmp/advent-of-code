with open("input.txt", "r") as file:
    sum = 0
    for line in file:
        possible = True
        game_num, game = line.split(":")
        game_num = int(game_num.split(" ")[1])
        sets = game.split(";")
        cubes = [j.strip() for x in sets for j in x.split(",")]
        green_list, red_list, blue_list = [], [], []
        for cube in cubes:
            num, color = cube.split(" ")
            num = int(num)

            if color == "green": green_list.append(num)
            elif color == "red": red_list.append(num)
            elif color == "blue": blue_list.append(num)

        min_green_list = max(green_list) if len(green_list) > 0 else 0
        min_red_list = max(red_list) if len(red_list) > 0 else 0
        min_blue_list = max(blue_list) if len(blue_list) > 0 else 0
        power = min_green_list * min_red_list * min_blue_list
        print(line)
        print(power)
        print(f"min green: {min_green_list}")
        print(f"min blue: {min_blue_list}")
        print(f"min red: {min_red_list}")
        sum += power


    print(sum)
