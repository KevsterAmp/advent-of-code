card_str = {
    "A": 13,
    "K": 12,
    "Q": 11,
    "J": 1,
    "T": 10,
}


def main():
    with open("input.in", "r") as file:
        lines = file.read().split("\n")
    card2bid = [line.split() for line in lines]
    print(len(lines))
    x = calibrate(card2bid)
    print(len(x))
    y = get_dupes(x)
    print(len(y))
    c = get_rankings(y)
    print(c)
    print(compute_bids(c))


def calibrate(lines):
    # for each line, turn letters to int
    output = []
    for cards, bid in lines:
        num_cards = []
        for card in cards:
            if card in card_str:
                num_cards.append(card_str[card])
            else:
                num_cards.append(int(card))

    # output should be a list containing bid and another list of numbers
        output.append((num_cards, bid))
    return output


def get_dupes(lines):
    # determine number of pairs
    output = []
    for line in lines:
        count_dict = {}
        strength = 0
        cards = line[0]
        bid = line[1]
        for card in cards:
            if card == 1:
                card = max(cards)

            if card not in count_dict:
                count_dict[card] = 1
            else:
                count_dict[card] += 1


        if len(count_dict) == 1:
            strength = 7
        elif len(count_dict) == 2 and 1 in count_dict.values():
            strength =6
        elif len(count_dict) == 2:
            strength =5
        elif len(count_dict) == 3 and 3 in count_dict.values():
            strength = 4
        # two pair
        elif len(count_dict) == 3:
            strength = 3
        # one pair
        elif len(count_dict) == 4:
            strength = 2
        # high pair
        else:
            strength = 1

        output.append((strength, cards, bid))
    return output



def get_rankings(lines):
    ranking = []
    # find equal strength and sort em
    # find all 5
    all_ranks = []
    for i in range(1, 8):
        rank = [x for x in lines if x[0] == i]
        all_ranks.append(rank)

    for rank in all_ranks:
        rank.sort(reverse=True, key=get_elem)
        ranking += rank

    return ranking

    # rank_7 = [x for x in lines if x[0] == 7]
    # rank_6 = [x for x in lines if x[0] == 6]
    # rank_5 = [x for x in lines if x[0] == 5]
    # rank_4 = [x for x in lines if x[0] == 4]
    # rank_3 = [x for x in lines if x[0] == 3]
    # rank_2 = [x for x in lines if x[0] == 2]
    # rank_1 = [x for x in lines if x[0] == 1]

    # all_ranks = [rank_7, rank_6, rank_5, rank_4, rank_3, rank_2, rank_1]


def get_elem(line):
    return line[1]


def compute_bids(lines):
    i = len(lines)
    output = 0
    for i, line in enumerate(lines):
        win = int(line[2]) * i+1
        output += win

    return output

main()
