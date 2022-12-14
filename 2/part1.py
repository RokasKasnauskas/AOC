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
    "X": "Rock",
    "Y": "Paper",
    "Z": "Scissors",
}
def define_winner(opponent, player):
    if opponent == player:
        return "Draw"
    if player == "Rock" and opponent == "Scissors":
        return "Win"
    if player == "Scissors" and opponent == "Paper":
        return "Win"
    if player == "Paper" and opponent == "Rock":
        return "Win"
    return "Lose"

total_points = 0

for input in inputs:
    splited_inputs = input.strip().split(' ')
    opponent = letter_meaning[splited_inputs[0]]
    player =  letter_meaning[splited_inputs[1]]
    match_result = define_winner(opponent, player)
    points_for_match_result = point_system[match_result]
    points_for_chosen_hand = point_system[player]
    total_match_point = points_for_chosen_hand + points_for_match_result
    total_points += total_match_point
print(total_points)