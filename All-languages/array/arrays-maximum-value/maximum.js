var numbers = [1,2,34,65,3223,3434]

var maximum = numbers[0]

for (var i=0;i<=numbers.length;i++) {
	var num = numbers[i]
	if (num > maximum) {
		maximum = num
	}
}
console.log(maximum)
