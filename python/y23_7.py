# Cards
# A K Q J T 9 8 7 6 5 4 3 2
#
# Hand types (ordered by strength)
#
# 5OAK: AAAAA (7) [1 type]
# 4OAK: AAAA8 (6) [2 types (4 & 1)]
# FH  : 33322 (5) [2 types (3 & 2)]
# 3OAK: AAA23 (4) [3 types (3, 1, 1)]
# 2P  : AA223 (3) [3 types (2, 2, 1)]
# 1P  : AA234 (2) [4 types]
# HC  : A2345 (1) [5 types]
#
# Ex. hands
#
# Hand  Bid
# 32T3K 765
# T55J5 684
#
# Must put the hands in order of strength to determine their rank
#
# Winnings = bid * rank
#
# For answer: add up all the winnings from each hand together
#

SCORE_5OAK = 7
SCORE_4OAK = 6
SCORE_FH = 5
SCORE_3OAK = 4
SCORE_2P = 3
SCORE_1P = 2
SCORE_HC = 1

score_string_map = {
    SCORE_5OAK: "five-of-a-kind (7)",
    SCORE_4OAK: "four-of-a-kind (6)",
    SCORE_FH: "full-house (5)",
    SCORE_3OAK: "three-of-a-kind (4)",
    SCORE_2P: "two-pair (3)",
    SCORE_1P: "one-pair (2)",
    SCORE_HC: "high-card (1)"
}

card_value_map = {
    "A": 14,
    "K": 13,
    "Q": 12,
    "J": 11,
    "T": 10,
    "9": 9,
    "8": 8,
    "7": 7,
    "6": 6,
    "5": 5,
    "4": 4,
    "3": 3,
    "2": 2
}

#  [10, 5, 5, 11, 5]
# Should become:
# '1005051105'
card_static_score_digits_map = {
    "A": "14",
    "K": "13",
    "Q": "12",
    "J": "11",
    "T": "10",
    "9": "09",
    "8": "08",
    "7": "07",
    "6": "06",
    "5": "05",
    "4": "04",
    "3": "03",
    "2": "02"
}


def parse_line(line: str):
    hand_str, bid_str = line.split()

    hand_list = list(map(lambda card: card_value_map[card], hand_str))
    hand_list_strs = list(map(lambda card: card_static_score_digits_map[card], hand_str))
    static_score = int("".join(hand_list_strs))

    return {
        "hand": hand_list,
        "score": get_hand_score(hand_list),
        "static_score": static_score,
        "bid": int(bid_str)
    }


def get_hand_score(hand: list):

    # We know the first one won't be encountered yet, no need to loop
    # "counter" keeps track of how many of each card we encounter.
    # Example: [ [10, 3], [4, 2] ] means "encountered a ten 3 times, and a four twice"
    # [ [card_val, count] ]
    counter = [[hand[0], 1]]

    for card_value in hand[1:]:
        got_one = False

        # Check if this card value was encountered already
        # If so, add to its current encounter value
        for i, encountered in enumerate(counter):
            if card_value == encountered[0]:
                counter[i][1] += 1
                got_one = True
                break

        # If not, add to the counters
        if not got_one:
            counter.append([card_value, 1])

    # Analyze the counter to figure out the hand score (type)
    # 5OAK: AAAAA (7) [1 type]
    # 1P  : AA234 (2) [4 types]
    # HC  : A2345 (1) [5 types]
    if len(counter) == 1:
        return SCORE_5OAK
    if len(counter) == 4:
        return SCORE_1P
    if len(counter) == 5:
        return SCORE_HC

    # If 2 different types - 4OAK or FH
    # 4OAK: AAAA8 (6) [2 types (4 & 1)]
    # FH  : 33322 (5) [2 types (3 & 2)]
    if len(counter) == 2:
        if counter[0][1] == 4 or counter[0][1] == 1:
            return SCORE_4OAK
        else:
            return SCORE_FH

    # If 3 different types - 3OAK or 2P
    # 2P  : AA223 (3) [3 types (2, 2, 1)]
    if counter[0][1] == 2 or counter[1][1] == 2:
        return SCORE_2P

    # Last possibility: 3OAK
    # 3OAK: AAA23 (4) [3 types (3, 1, 1)]
    return SCORE_3OAK


def part1(test_input, real_input):
    use_real_input = True
    input = real_input if use_real_input else test_input
    # print(f"Test input: {test_input}")

    # [{hand, score, bid}]
    hands = []

    for line in input.split("\n"):
        hands.append(
            parse_line(line)
        )

    # Sort by scores
    hands.sort(key=lambda hand_obj: hand_obj["score"])

    # List of lists of hands
    # Each list contains all hands tied with the same score
    tied_hands_lists = {}

    # Isolate into different lists
    for hand in hands:
        hand_score = hand["score"]

        # Set initial list for this score
        if hand_score not in tied_hands_lists:
            tied_hands_lists[hand_score] = []

        # Add to tied_hands_lists
        tied_hands_lists[hand_score].append(hand)

    # Loop through each list
    for score, tied_hands in sorted(tied_hands_lists.items()):
        # Sort all elements in each list
        tied_hands_lists[score].sort(key=lambda hand_obj: hand_obj["static_score"])

        # tied_hands.sort(key=lambda hand_tuple: hand_tuple["score"])

        # print(f"Score: {score}")
        # for tied_hand in tied_hands:
        #     print(tied_hand["hand"])
        #     print(tied_hand["static_score"])
        # print(f"===========================")

    total_score = 0
    rank = 1

    # Loop through everything again, now gather score
    for score, hands_list in sorted(tied_hands_lists.items()):
        for hand_obj in hands_list:
            hand = hand_obj["hand"]
            hand_score = hand_obj["score"]
            hand_bid = hand_obj["bid"]
            hand_score_str = score_string_map[hand_score]

            print(f"Rank {rank} - {hand_score_str} - {hand}")

            # Winnings for hand = bid * rank
            total_score += (hand_bid * rank)
            rank += 1

    return total_score


def part2(test_input, real_input):
    use_real_input = False
    input = real_input if use_real_input else test_input
