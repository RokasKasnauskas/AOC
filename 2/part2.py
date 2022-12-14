f = open("input.txt", "r")
inputs = f.readlines()
point_system = {
    "Rock": 1,
    "Paper": 2,
    "Scissors": 3,
    "Lose": 0,
    "Draw": 3,
    "Win": 6,
}
letter_meaning = {
    "A": "Rock",
    "B": "Paper",
    "C": "Scissors",
    "X": "Lose",
    "Y": "Draw",
    "Z": "Win",
}
beating = {
    "Rock": "Paper",
    "Paper": "Scissors",
    "Scissors": "Rock"
}
all_posibilites = ["Rock", "Paper", "Scissors"]
def choose_shape(opponent, match_ends_with):
    if match_ends_with == "Draw":
        return opponent
    if match_ends_with == "Win":
        return beating[opponent]
    return [a for a in all_posibilites if a not in [opponent, beating[opponent]]][0]

total_points = 0

for input in inputs:
    splited_inputs = input.strip().split(' ')
    opponent = letter_meaning[splited_inputs[0]]
    match_end_with =  letter_meaning[splited_inputs[1]]
    player_shape = choose_shape(opponent, match_end_with)
    points_for_match_result = point_system[match_end_with]
    points_for_chosen_hand = point_system[player_shape]
    total_match_point = points_for_chosen_hand + points_for_match_result
    total_points += total_match_point
print(total_points)