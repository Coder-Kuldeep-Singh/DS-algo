numbers = [10,20,30,110,102,202]

#lets assume the index 0 is high
maximum = numbers[0]

#linear search
for num in numbers:
    if num > maximum:
        maximum = num

print(maximum)

