const input = ""
const elfsItemsCalories =  input.split('\n\n').map((elf) => elf.split('\n'))
const elfsCalories = elfsItemsCalories.map((elf) => elf.reduce((a, b) => Number(a) +Number(b), 0))
const sortedElfsCaloreis = elfsCalories.sort((a,b) => a - b)
const elfsWithMostCalories = sortedElfsCaloreis.slice(-3)
const sumOfMostCalories = elfsWithMostCalories.reduce((a,b) => a + b, 0)
console.log(sumOfMostCalories)
