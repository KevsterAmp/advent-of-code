with open("input.txt", "r") as file:
    sum = 0
    for line in file:
        possible = True
        game_num, game = line.split(":")
        game_num = int(game_num.split(" ")[1])
        sets = game.split(";")
        cubes = [j.strip() for x in sets for j in x.split(",")]
        for cube in cubes:
            num, color = cube.split(" ")
            num = int(num)
            if (color == "green" and num > 13) or (color == "red" and num > 12) or (color == "blue" and num > 14):
                possible = False
                break
            
        if possible == True:
            # add sum to game_id
            sum += game_num

    print(sum)
