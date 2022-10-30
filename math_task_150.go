package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func One() { //Вывести на экран текст "Silence is golden".
	fmt.Println("Silence is golden")
}

func Two() { //Вывести на экран текущую дату.
	today := time.Now()
	//fmt.Println(today)
	fmt.Printf("The time is: %s\n", today.Format(time.RFC1123))
}

func Tree() { //Вывести на экран пять строк из нулей, причем количество нулей в каждой строке равно номеру строки.
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("0")
		}
		fmt.Print("\n")
	}
}

func Four() { //Вывести на экран прямоугольник, заполненный буквами А. Количество строк в прямоугольнике равно 5, количество столбцов равно 8.
	fmt.Print("\n")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("A")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func Five() { //Вывести на экран букву "W" из символов "*" (звездочка).
	fmt.Println("*     *     *")
	fmt.Println(" *   * *   * ")
	fmt.Println("  * *   * *  ")
	fmt.Println("   *     *   ")
}

func Six() { //Вывести на экран результат вычисления 1+2−4
	fmt.Println(1 + 2 - 4)
}

func Seven() { // Вычислите 1/2+1/4. Ответ: 0.75
	fmt.Println(1./2. + 1./4.)
}

func Eight(a float64, b float64) { //Вычислите значение выражения (a+4b)(a−3b)+a**2 при a=2 и b=3. Ответ:  -94
	c := (a+4*b)*(a-3*b) + math.Pow(a, 2)
	fmt.Println(c)
}

func Nine(x float64) { //Вычислите |x|+x**5, если x=−2. Ответ: -30
	ans := math.Abs(x) + math.Pow(x, 5)
	fmt.Println(ans)
}

func Ten(x float64) { // Вычислите значение выражения (x+1)**2+3*(x+1) при а) x=1.7; б) x=3. Ответ: а) 15.39 б) 28
	ans := x + 1
	ans = math.Pow(ans, 2) + 3*ans
	fmt.Printf("%.2f", ans)
	fmt.Print("\n")
}

func Eleven(x float64) { //Вычислите значение выражения ---------------- при x=−2.34. Ответ: -1.76911 (проверено!)
	ans := (math.Abs(x-5)-math.Sin(x))/3 + math.Pow(math.Pow(x, 2)+2014, 0.5)*math.Cos(2*x) - 3
	fmt.Printf("%.5f", ans)
	fmt.Print("\n")
}

func Twelve(x float64) { //Вычислите значение выражения e**(x−2)+|sin(x)|−x**4⋅cos(1/x) при x=3.6 Ответ: -156.1276
	ans := math.Exp(x-2) + math.Abs(math.Sin(x)) - math.Pow(x, 4)*math.Cos(1./x)
	fmt.Printf("%.4f", ans)
	fmt.Print("\n")
}

func Thirteen(a, b, x float64) { //Вычислите значение выражения ------------ при a=0.1, b=0.2 и x=1 Ответ: 1.0088
	ans := math.Pow(math.Pow(x, 2)+b, 1./5.) - (math.Pow(b, 2)*math.Pow(math.Sin(x+a), 3))/x
	fmt.Printf("%.4f", ans)
	fmt.Print("\n")
}

func Fourteen() { //Пользователь вводит два числа. Найдите сумму и произведение данных чисел.
	var a, b float64
	fmt.Print("Введите Первую переменную: ")
	fmt.Scan(&a)
	fmt.Print("Введите Вторую переменную: ")
	fmt.Scan(&b)
	fmt.Println("Сумма чисел: ", a+b)
	fmt.Println("Произведение чисел: ", a*b)
}

func Fifteen() { //Пользователь вводит число. Выведите на экран квадрат этого числа, куб этого числа.
	var a float64
	fmt.Print("Введите число: ")
	fmt.Scan(&a)
	fmt.Println("Квадрат числа равен", math.Pow(a, 2))
	fmt.Println("Куб числа равен", math.Pow(a, 3))
}

func Sixteen(a, b, c float64) { //Пользователь вводит три числа. Увеличьте первое число в два раза, второе числа уменьшите на 3, третье число возведите в квадрат и затем найдите сумму новых трех чисел.
	a *= 2
	b -= 3
	c *= c
	fmt.Println(a + b + c)
}

func Seventeen(a, b, c float64) { //Пользователь вводит три числа. Найдите среднее арифметическое этих чисел, а также разность удвоенной суммы первого и третьего чисел и утроенного второго числа.
	fmt.Println((a + b + c) / 3.)
	fmt.Println(2*(a+c) - 3*b)
}

func Eighteen() { //Пользователь вводит сторону квадрата. Найдите периметр и площадь квадрата.
	var a float64
	fmt.Print("Введите сторону квадрата: ")
	fmt.Scan(&a)
	fmt.Println("Периметр квадрата: ", 4*a)
	fmt.Println("Площадь квадрата: ", a*a)
}

func Nineteen() { //Пользователь вводит цены 1 кг конфет и 1 кг печенья. Найдите стоимость: а) одной покупки из 300 г конфет и 400 г печенья; б) трех покупок, каждая из 2 кг печенья и 1 кг 800 г конфет.
	var coo_p, can_p float64
	fmt.Print("Введите цену 1 килограмма печенья: ")
	fmt.Scan(&coo_p)
	fmt.Print("Введите цену 1 килограмма конфет: ")
	fmt.Scan(&can_p)
	fmt.Println("Цена первой покупки: ", can_p*0.3+coo_p*0.4)
	fmt.Println("Цена второй покупки: ", 3.*(can_p*1.8+coo_p*2.))
}

func Twenty() { //Пользователь вводит время в минутах и расстояние в километрах. Найдите скорость в м/c.) { //Пользователь вводит время в минутах и расстояние в километрах. Найдите скорость в м/c.
	var min, kilo float64
	fmt.Print("Введите время в минутах: ")
	fmt.Scan(&min)
	fmt.Print("Введите расстояние в километрах: ")
	fmt.Scan(&kilo)
	fmt.Println(1000*kilo/(60*min), " м/с")
}

func Twenty_one(a, b float64) { //Даны катеты прямоугольного треугольника. Найдите площадь, периметр и гипотенузу треугольника.
	fmt.Println("Площадь: ", a*b*0.5)
	c := math.Pow((math.Pow(a, 2) + math.Pow(b, 2)), 0.5)
	fmt.Println("Гипотенуза: ", c)
	fmt.Println("Периметр: ", a+b+c)
}

func Twenty_two(tc float64) { //Дано значение температуры в градусах Цельсия. Вывести температуру  в градусах Фаренгейта.
	fmt.Println("Температура в Фаренгейтах: ", (9./5.)*tc+32)
}

func Twenty_tree(x, a, y, k float64) { // Известно, что x кг конфет стоит a рублей. Определите, сколько стоит y кг этих конфет, а также сколько кг конфет можно купить на k рублей. Все значения вводит пользователь.
	fmt.Println(y, " кг конфет стоит ", y*a/x, " рублей.")
	fmt.Println("На ", k, " рублей можно купить ", k/(a/x), " кг конфет.")
}

func Twenty_four(d, p, s float64) { // Пользователь вводит количество дней, указывает процент скидки и вводит сумму. Рассчитать прибыль, если за каждый день сумма увеличивается на 3 $  и затем применяется скидка, то есть итоговая сумма еще увеличивается на данное число процентов.
	fmt.Println("Итоговая сумма: ", (s+3*d)+(s+3*d)*p/100)
}

func Twenty_five(d, m, y int64) { // Пользователь вводит количество недель, месяцев, лет и получает количество дней за это время. Считать, что в месяце 30 дней
	fmt.Println("Дней за это время : ", d+m*30+y*12*30)
}

func Twenty_six(a, b *int64) { //Даны две переменных с некоторыми значениями. Поменять местами значения этих переменных
	fmt.Println(*a, "  ", *b)
	*a, *b = *b, *a
	fmt.Println(*a, "  ", *b)
}

func Twenty_seven(a, b, c *float64) { //Даны три переменные a, b и c. Изменить значения этих переменных так, чтобы в a хранилось значение a+b, в b хранилась разность старых значений c−a, а в c хранилось сумма старых значений a+b+c. Например, a=0, b=2, c=5, тогда новые значения a=2, b=5 и c=7.
	*a, *b, *c = *a+*b, *c-*a, *a+*b+*c
}

func Thirty(a float64) { //Дано число a. Не пользуясь никакими арифметическими операциями кроме умножения, получите а)a*4 за две операции; б) a**6 за три операции; в) a**15 за пять операций.
	var na, b float64
	na = a
	b = 0
	fmt.Println("Начальное число: ", a)
	a *= a
	a *= a
	fmt.Println("a**4 = ", a)
	a = na
	a *= a
	b = a
	a *= a
	a *= b
	fmt.Println("a**6 = ", a)
	a, b = na, na
	a *= a
	a *= a
	a *= b
	b = a
	a *= a
	a *= b
	fmt.Println("a**15 = ", a)
}

func Thirty_two(a int64) { //Из трехзначного числа x вычли его последнюю цифру. Когда результат разделили на 10, а к частному слева приписали последнюю цифру числа x, то получилось число 237. Найти число x.
	var d int64 = 10
	for ; a/d > 0; d *= 10 {
	}
	d /= 10
	fmt.Println("Изначальное число ", (a/d)+(a%d)*10)
}

func Thirty_tree(x, y float64) { //Вычислите , если x и y вводит пользователь. Перед вычислением выполнить проверку на существование квадратных корней.
	if (y < 0) || (x-math.Pow(y, 0.5) < 0) {
		fmt.Println("Wrong operands")
	} else {
		fmt.Println("Answer: ", math.Pow(x-math.Pow(y, 0.5), 0.5))
	}
}

func Thirty_four(x *float64) { //Дано число. Если оно больше 3, то увеличить число на 10, иначе уменьшить на 10.
	if *x > 3 {
		*x += 10
	} else {
		*x -= 10
	}
}

func Thirty_five(x float64) { //Дано число. Если оно меньше 7, то вывести Yes, если больше 10, то вывести No, если равно 9, то вывести Error.
	if x < 7 {
		fmt.Println("Yes")
	} else if x == 9 {
		fmt.Println("Error")
	} else if x > 10 {
		fmt.Println("No")
	}
}

func Trirty_six(x float64) { //Пользователь вводит номер месяца, вывести название месяца.
	switch x {
	case 1:
		fmt.Println("Январь")
	case 2:
		fmt.Println("Февраль")
	case 3:
		fmt.Println("Март")
	case 4:
		fmt.Println("Апрель")
	case 5:
		fmt.Println("Май")
	case 6:
		fmt.Println("Июнь")
	case 7:
		fmt.Println("Июль")
	case 8:
		fmt.Println("Август")
	case 9:
		fmt.Println("Сентябрь")
	case 10:
		fmt.Println("Октябрь")
	case 11:
		fmt.Println("Ноябрь")
	case 12:
		fmt.Println("Декабрь")
	default:
		fmt.Println("Такого месяца не существует")
	}
}

func Thirty_seven(a, b float64) { //Дано два числа. Вывести наибольшее из них.
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}

func Thirty_eight(a, b float64) { //Дано два числа. Вывести yes, если они отличаются на 100, иначе вывести No.
	if math.Abs(a-b) == 100 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Thrity_nine(a, b int64) { //Даны два числа. Если первое число больше второго, то вывести yes, иначе поменять значения этих переменных и вывести их на экран.
	if a > b {
		fmt.Println("Yes")
	} else {
		Twenty_six(&a, &b)
	}
}

func Forty(b *float64) { //Дано число. Если оно от -10 до 10 не включительно, то увеличить его на 5, иначе уменьшить на 10
	if (-10 < *b) && (*b < 10) {
		*b += 5
	} else {
		*b -= 10
	}
}

func Forty_one(b *float64) { //Дано число. Если оно от -10 до 10 не включительно, то увеличить его на 5, иначе уменьшить на 10
	if (-100 <= *b) && (*b <= 100) {
		*b += 1
	} else {
		*b = 0
	}
}

func Forty_two(b *float64) { //Дано число. Если оно от 2 до 5 включительно, то увеличить его на 10. Если оно от 7 до 40, то уменьшить на 100. Если оно не более 0 или более 3000, то увеличить в 3 раза (то есть умножить на 3). Иначе занулить это число.
	if (2 < *b) && (*b <= 5) {
		*b += 10
	} else if (7 < *b) && (*b < 40) {
		*b -= 100
	} else if (*b < 0) || (*b > 3000) {
		*b *= 3
	} else {
		*b = 0
	}
}

func Forty_tree() { //Пользователь вводит номер месяца. Вывести название поры года (весна, лето и т.д.)
	var x float64
	fmt.Println("Введите номер месяца: ")
	fmt.Scan(&x)
	switch x {
	case 1, 2, 12:
		fmt.Println("Зима")
	case 3, 4, 5:
		fmt.Println("Весна")
	case 6, 7, 8:
		fmt.Println("Лето")
	case 9, 10, 11:
		fmt.Println("Осень")
	default:
		fmt.Println("Такого месяца не существует")
	}
}

func Forty_four() { //Пользователь вводит два числа. Если они не равны 10 и первое число четное, то вывести их сумму, иначе вывести их произведение.
	var x, y int64
	fmt.Println("Введите два числа ")
	fmt.Scan(&x, &y)
	if (x != 10 && y != 10) && x%2 != 1 {
		fmt.Println(x + y)
	} else {
		fmt.Println(x * y)
	}
}

func Forty_five() { //Пользователь вводит три числа. Если все числа больше 10 и первые два числа делятся на 3, то вывести yes, иначе no
	var x, y, z int64
	fmt.Println("Введите три числа ")
	fmt.Scan(&x, &y, &z)
	if x > 10 && y > 10 && z > 10 && x%3 == 0 && y%3 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Forty_six() { //Пользователь вводит три числа. Найти сумму тех чисел, которые делятся на 5. Если таких чисел нет, то вывести error
	var x, y, z, ans int64
	fmt.Println("Введите три числа ")
	fmt.Scan(&x, &y, &z)
	ans = 0
	if x%5 == 0 {
		ans += x
	}
	if y%5 == 0 {
		ans += y
	}
	if z%5 == 0 {
		ans += z
	}
	fmt.Println(ans)
}

func Forty_seven() { //Даны три числа. Найдите наибольшее число из них.
	var x, y, z, ans int64
	fmt.Println("Введите три числа ")
	fmt.Scan(&x, &y, &z)
	ans = x
	if y > ans {
		ans = y
	}
	if z > ans {
		ans = z
	}
	fmt.Println(ans)
}

func Forty_eight() { //Даны три числа. Найдите те два из них, сумма которых наибольшая.
	var x, y, z, ans int64
	fmt.Println("Введите три числа ")
	fmt.Scan(&x, &y, &z)
	ans = x
	if y < ans {
		ans = y
	}
	if z < ans {
		ans = z
	}
	ans = x + y + z - ans
	fmt.Println(ans)
}

func Forty_nine() { // Пользователь вводит четыре числа. Найдите наибольшее четное число среди них. Если оно не существует, выведите фразу "not found"
	var x, y, z, f, ans int64
	var b bool = false
	fmt.Println("Введите четыре числа ")
	fmt.Scan(&x, &y, &z, &f)
	if x > ans && x%2 == 0 {
		ans = x
		b = true
	}
	if y > ans && y%2 == 0 {
		ans = y
		b = true
	}
	if z > ans && z%2 == 0 {
		ans = z
		b = true
	}
	if f > ans && f%2 == 0 {
		ans = f
		b = true
	}
	if b {
		fmt.Println(ans)
	} else {
		fmt.Println("not found")
	}
}

func Fifty() { //Даны три числа. Написать "yes", если среди них есть одинаковые.
	var x, y, z int64
	fmt.Println("Введите три числа ")
	fmt.Scan(&x, &y, &z)
	if x == y {
		fmt.Println("Yes")
	} else if y == z {
		fmt.Println("Yes")
	} else if x == z {
		fmt.Println("Yes")
	}
}

func Fifty_one() { //Даны три числа. Написать "yes", если можно взять какие-то два из них и в сумме получить третье
	var x, y, z int64
	fmt.Println("Введите три числа ")
	fmt.Scan(&x, &y, &z)
	if x+z == y {
		fmt.Println("Yes")
	} else if y+x == z {
		fmt.Println("Yes")
	} else if y+z == x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Fifty_two() { //Дано четыре числа, если первые два числа больше 5, третье число делится на 6, четвертое число не делится на 3, то вывести yes, иначе no.
	var x, y, z, c int64
	fmt.Println("Введите четыре числа ")
	fmt.Scan(&x, &y, &z, &c)
	if x > 5 && y > 5 && z%6 == 0 && c%3 != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Fifty_tree() { //Дано два числа. Если хотя бы одно из них больше 30, то вывести yes, иначе no.
	var x, y int64
	fmt.Println("Введите два числа ")
	fmt.Scan(&x, &y)
	if x > 30 || y > 30 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Fifty_four() { //Дано три числа. Если ровно два из них  меньше 5, то вывести yes, иначе вывести no
	var x, y, z, f int64
	fmt.Println("Введите три числа ")
	fmt.Scan(&x, &y, &z)
	f = 0
	if x < 5 {
		f += 1
	}
	if y < 5 {
		f += 1
	}
	if z < 5 {
		f += 1
	}
	if f == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Fifty_five() { //Дано три числа. Найти количество положительных чисел среди них.
	var x, y, z, ans int64
	fmt.Println("Введите три числа ")
	fmt.Scan(&x, &y, &z)
	ans = 0
	if x > 0 {
		ans += 1
	}
	if y > 0 {
		ans += 1
	}
	if z > 0 {
		ans += 1
	}
	fmt.Println(ans)
}

func Fifty_six() { //Робот может перемещаться в четырех направлениях («11» — север, «12» — запад, «13» — юг, «14» — восток) и принимать три цифровые команды: 0 — продолжать движение, 1 — поворот налево, –1 — поворот направо. Дан число (11, 12, 13 или 14) — исходное направление робота и целое число N (0, 1 или -1) — посланная ему команда. Вывести направление робота после выполнения полученной команды (то есть север, запад, юг или восток).
	var x, y int64
	fmt.Println("Введите два числа ")
	fmt.Scan(&x, &y)
	if !(x == 11 || x == 12 || x == 13 || x == 14) && (y == 0 || y == 1 || y == -1) {
		fmt.Println("Error input")
		return
	} else if (x == 11) && (y == -1) {
		x = 14
	} else if (x == 14) && (y == 1) {
		x = 11
	} else {
		x = y + x
	}
	switch x {
	case 11:
		fmt.Println("Север")
	case 12:
		fmt.Println("запад")
	case 13:
		fmt.Println("юг")
	case 14:
		fmt.Println("восток")
	}
}

func Fifty_seven() { // Дана дата из трех чисел (день, месяц и год). Вывести yes, если такая дата существует (например, 12 02 1999 - yes, 22 13 2001 - no). Считать, что в феврале всегда 28 дней.
	Error := "32/12/2017"
	layout := "02/01/2006"
	var input string
	fmt.Println("Введите дату в формате dd/mm/yy :")
	fmt.Scan(&input)

	e, _ := time.Parse(layout, Error)
	t, _ := time.Parse(layout, input)

	if t == e {
		fmt.Println("Такой даты не существует no")
	} else {
		fmt.Println("Yes")
	}
}

func Fifty_eight() { // Дано две даты, каждая из которых состоит из трех чисел (день, месяц и год). Вывести yes, если первая дата раньше второй, иначе вывести no.
	Error := "32/12/2017"
	layout := "02/01/2006"
	var input1, input2 string
	fmt.Println("Введите две даты в формате dd/mm/yy :")
	fmt.Scan(&input1, &input2)

	e, _ := time.Parse(layout, Error)
	t1, _ := time.Parse(layout, input1)
	t2, _ := time.Parse(layout, input2)

	if t1 == e || t2 == e {
		fmt.Println("Такой даты не существует no")
	} else if t1.Sub(t2) < 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Fifty_nine() { // Дано четырехзначное число. Верно ли, что цифр в нем расположены по убыванию?Например, 4311 - нет, 4321 - да, 5542 - нет, 5631 - нет, 9871 - да.
	var input, x2 int64
	fmt.Println("Введите четырехзначное число :")
	fmt.Scan(&input)
	x2 = 100
	for x2 != 0 {
		if (input/(10*x2))%10 <= (input/x2)%10 {
			fmt.Println("No")
			return
		} else {
			x2 /= 10
		}
	}
	fmt.Println("Yes")
}

func Sixty() { //Дано трехзначное число. Переставьте первую и последнюю цифры.
	var input int64
	fmt.Println("Введите трехзначное число :")
	fmt.Scan(&input)
	fmt.Println(input/100 + ((input%100)/10)*10 + (input%10)*100)
}

func Sixty_one() { //Дано четырехзначное число. Определите, есть ли одинаковые цифры в нем
	var input int64
	fmt.Println("Введите четырехзначное число :")
	fmt.Scan(&input)
	i1 := input / 1000
	i2 := (input / 100) % 10
	i3 := (input / 10) % 10
	i4 := input % 10
	switch i1 {
	case i2, i3, i4:
		fmt.Println("Есть одинаковые цифры")
	default:
		switch i2 {
		case i3, i4:
			fmt.Println("Есть одинаковые цифры")
		default:
			if i3 == i4 {
				fmt.Println("Есть одинаковые цифры")
			} else {
				fmt.Println("Одинаковых цифр нет")
			}
		}
	}
}

func Sixty_two() { //Дано пятизначное число. Цифры на четных позициях занулить. Например, из 12345 получается число 10305.
	var input int64
	fmt.Println("Введите пятизначное число :")
	fmt.Scan(&input)
	fmt.Println(input - (input/1000)%10*1000 - (input/10)%10*10)
}

func Sixty_tree() { //Даны два трехзначных числа. Найдите шестизначное число, образованное из двух данных чисел путем дописывания второго числа к первому справа.
	var input1, input2 int64
	fmt.Println("Введите два трехзначных числа: ")
	fmt.Scan(&input1, &input2)
	fmt.Println(input1*1000 + input2)
}

func Sixty_four() { //Дано четырехзначное число. Если оно читается слева направо и справа налево одинаково, то вывести yes, иначе no.
	var input, reverse int64
	fmt.Println("Введите четырехзначное число :")
	fmt.Scan(&input)
	i1 := input / 1000
	i2 := (input / 100) % 10
	i3 := (input / 10) % 10
	i4 := input % 10
	reverse = i4*1000 + i3*100 + i2*10 + i1
	if input == reverse {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Sixty_five() { //Дано четырехзначное число. Переставьте местами цифры так, чтобы сначала оказались цифры, меньшие пяти.
	var input, x1, ans1, x2, ans2 int64
	fmt.Println("Введите четырехзначное число :")
	fmt.Scan(&input)
	x1 = 1000
	x2 = 1
	i1 := input / 1000
	i2 := (input / 100) % 10
	i3 := (input / 10) % 10
	i4 := input % 10
	ans1 = 0
	switch i1 {
	case 0, 1, 2, 3, 4:
		ans1 += i1 * x1
		x1 /= 10
	case 5, 6, 7, 8, 9:
		ans2 += i1 * x2
		x2 *= 10
	}
	switch i2 {
	case 0, 1, 2, 3, 4:
		ans1 += i2 * x1
		x1 /= 10
	case 5, 6, 7, 8, 9:
		ans2 += i2 * x2
		x2 *= 10
	}
	switch i3 {
	case 0, 1, 2, 3, 4:
		ans1 += i3 * x1
		x1 /= 10
	case 5, 6, 7, 8, 9:
		ans2 += i3 * x2
		x2 *= 10
	}
	switch i4 {
	case 0, 1, 2, 3, 4:
		ans1 += i4 * x1
		x1 /= 10
	case 5, 6, 7, 8, 9:
		ans2 += i4 * x2
		x2 *= 10
	}
	fmt.Println(ans1 + ans2)
}

func Sixty_six() { //Даны два трехзначных числа. Получите новое число присоединением второго числа справа к первому без последних цифр у каждого. Например, 123 и 456 Ответ: 1245
	var input1, input2 int64
	fmt.Println("Введите два трехзначных числа: ")
	fmt.Scan(&input1, &input2)
	fmt.Println((input1/10)*100 + (input2 / 10))
}

func Sixty_seven() { //Дано четырехзначное число. Поменяйте местами наименьшую и наибольшую цифры.
	var input1, min, max, x, ans int64
	var max_f, min_f bool = true, true
	fmt.Println("Введите четырехзначное число: ")
	fmt.Scan(&input1)
	max = 0
	min = 9
	x = 1000
	ans = 0
	for ; x != 0; x /= 10 {
		if (input1/x)%10 > max {
			max = (input1 / x) % 10
		}
		if (input1/x)%10 < min {
			min = (input1 / x) % 10
		}
	}
	x = 1000
	for ; x != 0; x /= 10 {
		if ((input1/x)%10 == max) && max_f {
			max_f = false
			ans += min * x
		} else if ((input1/x)%10 == min) && min_f {
			min_f = false
			ans += max * x
		} else {
			ans += ((input1 / x) % 10) * x
		}
	}
	fmt.Println(ans)
}

func Sixty_eight() { //Даны коэффициенты a,b,c уравнения ax2+bx+c=0. Найти решение. Проверить ответы можно здесь. Как решать квадратные уравнения можно прочитать здесь.
	var a, b, c float64
	fmt.Println("Введите коэффициенты a,b,c  ")
	fmt.Scan(&a, &b, &c)
	D := math.Pow(b, 2) - 4*a*c
	if a == 0 && b == 0 {
		fmt.Println("Выражение бессмысленно")
		return
	}
	if D > 0 && a != 0 {
		fmt.Println((-b+math.Pow(D, 0.5))/(2*a), "  ", (-b-math.Pow(D, 0.5))/(2*a))
	} else if D == 0 {
		fmt.Println((-b + math.Pow(D, 0.5)) / (2 * a))
	} else if D < 0 {
		fmt.Println("Корней нет")
		return
	} else if a == 0 {
		fmt.Println(-c / b)
	}
}

func Sixty_nine() { //Пользователь вводит три числа - длины сторон треугольника. Найти площадь треугольника. Сделать проверку на существование треугольника (например, 1, 2, 3 - такого треугольника не существует). Проверить ответы можно здесь
	var a, b, c float64
	fmt.Println("Введите стороны треугольника ")
	fmt.Scan(&a, &b, &c)
	if (a >= b+c) || (b >= a+c) || (c >= a+b) {
		fmt.Println("Такого треугольника не существует")
		return
	}
	p := (a + b + c) / 2
	s := math.Pow(p*(p-a)*(p-b)*(p-c), 0.5)
	fmt.Println("Площадь треугольника: ", s)
}

func Seventy() { //Даны целочисленные координаты трех вершин прямоугольника, стороны которого параллельны координатным осям. Найдите координаты его четвертой вершины (после проверки введенных данных на правильность).
	var ax, bx, cx, ay, by, cy, x, y float64
	fmt.Println("Введите Координаты вершины a: ")
	fmt.Scan(&ax, &ay)
	fmt.Println("Введите Координаты вершины b: ")
	fmt.Scan(&bx, &by)
	fmt.Println("Введите Координаты вершины c: ")
	fmt.Scan(&cx, &cy)
	if !((ax == bx && cx != ax) || (bx == cx && ax != cx) || (ax == cx && bx != ax)) {
		fmt.Println("Wrong input")
		return
	}
	if !((ay == by && cy != ay) || (by == cy && ay != cy) || (ay == cy && by != ay)) {
		fmt.Println("Wrong input")
		return
	}
	x = ax
	y = ay
	if x == bx {
		x = cx
	} else if x == cx {
		x = bx
	}
	if y == by {
		y = cy
	} else if y == cy {
		y = by
	}
	fmt.Println(x, y)
}

func Seventy_one() { //Даны числа h и m, где h - количество часов, m - количество минут некоторого момента времени. Найдите угол между часовой и минутной стрелками в этот момент времени.
	var h, m float64
	fmt.Println("Введите часы и минуты: ")
	fmt.Scan(&h, &m)
	if h < 0 || h > 12 || m < 0 || m > 60 {
		fmt.Println("Ошибка ввода")
		return
	} // Каждую минуту часовая стрелка смещается на 30°:60 = 0.5°
	a := math.Abs(h*30 - m*5.5)
	if a > 180 {
		a = 360 - a
	}
	fmt.Println("Между часовой и минутной стрелкой ", a, " градусов")
}

func Seventy_five() { // Выведите на экран 10 раз фразу "You are welcome!"
	for i := 0; i < 10; i++ {
		fmt.Println(i+1, " You are welcome!")
	}
}

func Seventy_six() { //Выведите на экран n раз фразу "Silence is golden". Число n вводит пользователь.
	var n int
	fmt.Println("Введите колличество строк n: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println("Silence is golden")
	}
}

func Seventy_seven() { //Выведите на экран прямоугольник из нулей. Количество строк вводит пользователь, количество столбцов равно 5.
	var n int
	fmt.Println("Введите колличество строк n: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println("00000")
	}
}

func Seventy_eight() { //Вывести на экран фигуру из звездочек:(квадрат из n строк, в каждой строке n звездочек)
	var n int
	fmt.Println("Введит сторону квадрата n: ")
	fmt.Scan(&n)
	fmt.Println()
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func Seventy_nine() { // Выведите на экран числа 1, 2, 3, 4, ..., 20.
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}
}

func Eighty_eighty_two() { //Вывести на экран ряд чисел 1001,  1004,  1007,  ... 1025.
	//Вывести на экран числа 100, 96, 92, ... до последнего положительного включительно.
	//Выведите на экран числа 1.2, 1.4, 1.6, ..., 2.8.
	var str, stp, fin float64
	fmt.Println("Введите начальную точку, шаг и конечную : ")
	fmt.Scan(&str, &stp, &fin)
	fmt.Println()
	if str > fin {
		for i := str; i >= fin; i += stp {
			fmt.Print(i, " ")
		}
	} else {
		for i := str; i <= fin; i += stp {
			fmt.Print(i, " ")
		}
	}
}

func Eighty_tree() { //Выведите следующие строки. Первая: 25  25.5  24.8. Вторая: 26 26.5 25.8.  И так далее. Последняя строка: 35 35.5 34.8.
	var i float64
	for i = 25; i <= 35; i++ {
		fmt.Println(i, i+0.5, i-0.2)
	}
}

func Eighty_four() { //Пользователь вводит курс доллара в рублях. Показать таблицу цен 1$, 2$, ..., 100$ в рублях, третьим столбцом добавить количество кг конфет, которые можно купить на данные суммы, если цена 1 кг конфет равна 20 руб. Пример: 1$ - 70 р - 3.5 кг и так далее (всего 100 строк).
	var der, d float64
	fmt.Println("Введите курс доллара: ")
	fmt.Scan(&der)
	fmt.Println("$     руб    кг Конфет")
	for d = 1; d <= 100; d++ {
		fmt.Println(d, "   ", d*der, "   ", d*der/20)
	}
}

func Eighty_six() { //Для данного n найти сумму 1+2+3+...+n. Например, для n=10 ответ равен 55.
	var d, i, ans int64
	fmt.Println("Введите число: ")
	fmt.Scan(&d)
	ans = 0
	for i = 1; i <= d; i++ {
		ans += i
	}
	fmt.Println("Ответ: ", ans)
}

func Eighty_seven() { //Найти сумму 10+11+12+13+...+88. Ответ: 3871
	var i, ans int64
	ans = 0
	for i = 10; i <= 88; i++ {
		ans += i
	}
	fmt.Println("Ответ: ", ans)
}

func Eighty_eight() { //Найти произведение 5⋅6⋅7⋅...⋅13.
	var i, ans int64
	ans = 1
	for i = 5; i <= 13; i++ {
		ans *= i
	}
	fmt.Println("Ответ: ", ans)
}

func Eighty_nine() { // 1+4+7+11+...+112.
	var ans, i int64
	ans = 0
	for i = 1; i <= 112; i += 3 {
		ans += i
	}
	fmt.Println("Ответ: ", ans)
}

func Ninety() { // Найти сумму  cos3/5+cos5/7+cos7/9+...+cos111/113
	var n, ans float64 = 0, 0
	for n = 5; n <= 113; n += 2 {
		ans += math.Cos((n - 2) / n)
	}
	fmt.Println(ans)
}

func Ninety_one() { // Найти сумму 23+34+45+...+910
	var n, ans float64 = 0, 0
	for n = 3; n <= 10; n += 1 {
		ans += (n - 1) / n
	}
	fmt.Println(ans)
}

func Ninety_two() { // Вывести на экран сто первых сумм вида 1+2+3+...+n.
	var n, ans, i float64 = 0, 0, 1
	for n = 1; n <= 100; n++ {
		ans = 0
		for i = 1; i <= n; i += 1 {
			ans += i
		}
		fmt.Println(ans)
	}
}

func Ninety_tree() { //Найдите сумму квадратов первых n натуральных чисел
	var n, ans, i float64 = 0, 0, 1
	fmt.Println("Введите натуральное число: ")
	fmt.Scan(&n)
	for i = 1; i <= n; i += 1 {
		ans += math.Pow(i, 2)
	}
	fmt.Println(ans)

}

func Ninety_four() { //Найдите сумму 1+1/2+1/3+…+1/n
	var n, ans, i float64 = 0, 0, 1
	fmt.Println("Введите натуральное число: ")
	fmt.Scan(&n)
	for i = 1; i <= n; i += 1 {
		ans += 1 / i
	}
	fmt.Println(ans)

}

func Ninety_five() { // Даны a и n. Вычислите p=(a+1)**2(a+2)**2⋅…⋅(a+n)**2
	var n, ans, i, a float64 = 0, 1, 1, 0
	fmt.Println("Введите два числа: ")
	fmt.Scan(&n, &a)
	for i = 1; i <= n; i += 1 {
		ans *= math.Pow(a+i, 2)
	}
	fmt.Println(ans)
}

func Ninety_six() { //ано натуральное число n. Вычислите 1/cosx+1/cosx**2+…+1/cosx**n
	var n, ans, i, x float64 = 0, 0, 1, 1
	fmt.Println("Введите натуральное число и вещественное: ")
	fmt.Scan(&n, &x)
	for i = 1; i <= n; i += 1 {
		ans += 1 / math.Cos(math.Pow(x, i))
	}
	fmt.Println(ans)
}

func Ninety_seven() { //Вычислите 1⋅2+2⋅3⋅4+...+n⋅(n+1)⋅…⋅2n.
	var n, ans1, ans2, i, j float64 = 0, 0, 1, 1, 1
	fmt.Println("Введите натуральное число: ")
	fmt.Scan(&n)
	for i = 1; i <= n; i += 1 {
		ans2 = 1
		for j = i; j <= 2*i; j++ {
			ans2 *= j
		}
		ans1 += ans2
	}
	fmt.Println(ans1)
}

func Ninety_eight() { // Начав тренировки, лыжник в первый день пробежал 10 км. Каждый следующий день он увеличивал пробег на 10% от пробега предыдущего дня. Определите: а) пробег лыжника за второй, третий, ..., десятый день тренировок; б) какой суммарный путь он пробежал за первые 7 дней тренировок. в) суммарный путь за n дней тренировок; г) в какой день ему следует прекратить увеличивать пробег, если он не должен превышать 80 км?
	var n, ans1, ans2, i float64 = 0, 10, 0, 1
	var j bool = true
	fmt.Println("Введите натуральное число: ")
	fmt.Scan(&n)
	for i = 1; i <= 10; i++ {
		if i <= 7 {
			ans2 += ans1
		}
		fmt.Println(i, " пробежал ", ans1)
		ans1 = ans1 * 1.1
	}
	fmt.Println("За первые семь дней он пробежал ", ans2)
	ans1 = 10
	ans2 = 0
	for i = 1; i <= n; i++ {
		if ans1 >= 80 && j {
			fmt.Println("На ", i, "тий день можно прекратить")
			j = false
		}
		ans2 += ans1
		ans1 = ans1 * 1.1
	}
	fmt.Println(ans1, " пробежал за ", n, "дней")
}

func Ninety_nine() { // Вывести на экран числа от 1000 до 9999 такие, что все цифры различны.
	var numb, n1, n2, n3, n4 int64
	for numb = 1000; numb <= 9999; numb++ {
		n1 = numb % 10
		n2 = (numb / 10) % 10
		n3 = (numb / 100) % 10
		n4 = numb / 1000
		if n1 != n2 && n1 != n3 && n1 != n4 && n2 != n3 && n2 != n4 && n3 != n4 {
			fmt.Println(numb)
		}
	}
}

func Hundred() { //Вывести на экран числа от 1000 до 9999 такие, что среди цифр нет цифр 5 и цифры 6.
	var numb, n1, n2, n3, n4 int64
	for numb = 1000; numb <= 9999; numb++ {
		n1 = numb % 10
		n2 = (numb / 10) % 10
		n3 = (numb / 100) % 10
		n4 = numb / 1000
		if n1 != 5 && n1 != 6 && n2 != 5 && n2 != 6 && n3 != 5 && n3 != 6 && n4 != 5 && n4 != 6 {
			fmt.Println(numb)
		}
	}
}

func One_hundred_one() { //Вывести все пятизначные числа, которые делятся на 2, у которых средняя цифра нечетная, и сумма всех цифр делится на 4. Материал сайта www.itmathrepetitor.ru
	var numb, n1, n2, n3, n4, n5 int64
	for numb = 10000; numb <= 99999; numb++ {
		n1 = numb % 10
		n2 = (numb / 10) % 10
		n3 = (numb / 100) % 10
		n4 = (numb / 1000) % 10
		n5 = numb / 10000
		if numb%2 == 0 && n3%2 == 1 && (n1+n2+n3+n4+n5)%4 == 0 {
			fmt.Println(numb)
		}
	}
}

func One_hundred_two() { // Вывести на экран числа от 1000 до 9999 такие, что среди цифр есть цифра 3.
	var numb, n1, n2, n3, n4 int64
	for numb = 1000; numb <= 9999; numb++ {
		n1 = numb % 10
		n2 = (numb / 10) % 10
		n3 = (numb / 100) % 10
		n4 = numb / 1000
		if n1 == 3 || n2 == 3 || n3 == 3 || n4 == 3 {
			fmt.Println(numb)
		}
	}
}

func One_hundred_tree() { //Найдите трехзначные числа, равные сумме кубов своих цифр.
	var numb, n1, n2, n3 int64
	for numb = 100; numb <= 999; numb++ {
		n1 = numb % 10
		n2 = (numb / 10) % 10
		n3 = numb / 100
		if float64(numb) == math.Pow(float64(n1), 3)+math.Pow(float64(n2), 3)+math.Pow(float64(n3), 3) {
			fmt.Println(numb)
		}
	}
}

func One_hundred_four() { //Сколько существует четырехзначных чисел, которые в 600 раз больше суммы своих цифр?
	var numb, n1, n2, n3, n4, i int64
	i = 0
	for numb = 1000; numb <= 9999; numb++ {
		n1 = numb % 10
		n2 = (numb / 10) % 10
		n3 = (numb / 100) % 10
		n4 = numb / 1000
		if numb/(n1+n2+n3+n4) >= 600 {
			fmt.Println(numb)
			i++
		}
	}
	fmt.Println(i)
}

func One_hundred_five() { //Найдите хотя одно натуральное число, которое делится на 11, а при делении на 2, 3, 4, ..., 10 дает в остатке 1
	var numb, i int64
	var f bool = true
	for numb = 1; numb < math.MaxInt64; numb++ {
		f = true
		if numb%11 == 0 {
			for i = 2; i <= 10; i++ {
				if numb%i != 1 {
					f = false
					break
				}
			}
			if f {
				fmt.Println(numb)
				return
			}
		}
	}
}

func One_hundred_six() { //Вывести на экран n единиц, затем 2n двоек, затем 3n троек. Число n вводит пользователь
	var n, i int64
	fmt.Println("Введите число: ")
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Print("1")
	}
	fmt.Println("")
	for i = 1; i <= 2*n; i++ {
		fmt.Print("2")
	}
	fmt.Println("")
	for i = 1; i <= 3*n; i++ {
		fmt.Print("3")
	}
	fmt.Println("")
}

func One_hundred_seven() { //Вывести ряд чисел: десять десяток, девять девяток, восемь восьмерок, ... , одну единицу. Найти сумму всех этих чисел.
	var i, j, ans int64 = 0, 0, 0
	for i = 10; i >= 1; i-- {
		for j = 1; j <= i; j++ {
			ans += i
			fmt.Print(i, " ")
		}
		fmt.Print("\n")
	}
	fmt.Println(ans)
}

func One_hundred_eight() { // Выведите  на экран строки (в последней строке n звездочек)
	var i, j, n int64 = 0, 0, 0
	fmt.Println("Введите число: ")
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func One_hundred_nine() { // (всего n строк, звездочек или 7, или 4 по очереди)
	var i, n int64 = 0, 0
	fmt.Println("Введите число: ")
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Println("*******")
		} else {
			fmt.Println("****")
		}
	}
}

func One_hundred_ten() { // Вывести на экран:
	//AAABBBAAABBBAAABBB
	//BBBAAABBBAAABBBAAA
	//AAABBBAAABBBAAABBB
	//(таких строк n, в каждой строке m троек AAA)
	var i, j, n, m int64 = 0, 0, 0, 0
	fmt.Println("Введите два числа: ")
	fmt.Scan(&n, &m)
	for i = 1; i <= n; i++ {
		if i%2 == 1 {
			for j = 1; j <= 3*m; j++ {
				fmt.Print("AAABBB")
			}
		} else {
			for j = 1; j <= 3*m; j++ {
				fmt.Print("BBBAAA")
			}
		}
		fmt.Print("\n")
	}
}

func One_hundred_eleven() { //Вывести на экран:
	//AAAAAAAAAAAAAAAA
	//ABBBBBBBBBBBBBBA
	//ABBBBBBBBBBBBBBA
	//ABBBBBBBBBBBBBBA
	//AAAAAAAAAAAAAAAA
	//(количество строк вводит пользователь, ширина прямоугольника в два раза больше высоты)
	var h, i, j int64 = 0, 0, 0
	fmt.Println("Введите количество строк: ")
	fmt.Scan(&h)
	for i = 1; i <= 2*h; i++ {
		fmt.Print("A")
	}
	for i = 2; i <= h-1; i++ {
		fmt.Print("\n")
		fmt.Print("A")
		for j = 2; j <= 2*h-1; j++ {
			fmt.Print("B")
		}
		fmt.Print("A")
	}
	fmt.Print("\n")
	for i = 1; i <= 2*h; i++ {
		fmt.Print("A")
	}
	fmt.Print("\n")
}

func One_hundred_twelve() { //Выведите на экран квадрат из нулей и единиц, причем нули находятся только на диагонали квадрата. Всего в квадрате сто цифр.
	var i, j int64
	for i = 1; i <= 10; i++ {
		fmt.Print("\n")
		for j = 1; j <= 10; j++ {
			if i == j || i == 10-j+1 {
				fmt.Print("0")
			} else {
				fmt.Print("1")
			}
		}
	}
	fmt.Print("\n")
}

func One_hundred_thirteen() { //Вывести на экран 20 строк. В строках с четными номерами вывести по 10 чисел, равных номеру строки. В строках с нечетными номерами вывести десять единиц.
	var i, j int64
	for i = 1; i <= 20; i++ {
		fmt.Print("\n")
		if i%2 == 0 {
			for j = 1; j <= 10; j++ {
				fmt.Print(i, " ")
			}
		} else {
			for j = 1; j <= 10; j++ {
				fmt.Print("1 ")
			}
		}
	}
	fmt.Print("\n")
}

func One_hundred_fourteen() { //Вывести 30 строк. Нечетные строки содержат натуральные числа от 1 до номера текущей строки включительно с шагом 1, четные строки состоят  из пяти единиц.
	var i, j int64
	for i = 1; i <= 30; i++ {
		fmt.Print("\n")
		if i%2 == 0 {
			for j = 1; j <= 5; j++ {
				fmt.Print("1 ")
			}
		} else {
			for j = 1; j <= i; j++ {
				fmt.Print(j, " ")
			}
		}
	}
	fmt.Print("\n")
}

func One_hundred_fifteen() { //Выведите на экран таблицу умножения для чисел от 1 до 10.
	var i, j int64
	for i = 1; i <= 10; i++ {
		for j = 1; j <= 10; j++ {
			fmt.Println(i, "x", j, "=", i*j)
		}
		fmt.Println("")
	}

}

func One_hundred_sixteen() { // Найдите количество целых чисел от a до b включительно, которые делятся на 12
	var a, b, i, ans int64
	ans = 0
	fmt.Println("введите промежуток")
	fmt.Scan(&a, &b)
	for i = a; i <= b; i++ {
		if i%12 == 0 {
			ans += 1
		}
	}
	fmt.Println(ans)
}

func One_hundred_seventeen() { //Пользователь вводит ненулевые числа до тех пор пока не введет ноль. Найдите сумму этих чисе
	var i, ans int64 = 1, 0
	for i != 0 {
		fmt.Scan(&i)
		ans += i
	}
	fmt.Println(ans)
}

func One_hundred_eighteen() { //Пользователь вводит ненулевые целые числа до тех пор, пока не введет ноль. Найдите количество четных чисел, которые он ввел.
	var i, ans int64 = 1, 0
	for i != 0 {
		fmt.Scan(&i)
		if i%2 == 0 {
			ans += 1
		}
	}
	fmt.Println(ans - 1)
}

func One_hundred_nineteen() { //Найдите четырехзначные числа, сумма цифр которых равна 15.
	var numb, n1, n2, n3, n4 int64
	for numb = 1000; numb <= 9999; numb++ {
		n1 = numb % 10
		n2 = (numb / 10) % 10
		n3 = (numb / 100) % 10
		n4 = numb / 1000
		if n1+n2+n3+n4 == 15 {
			fmt.Println(numb)
		}
	}
}

func One_hundred_twenty() { // Найдите наибольшую цифру в данном натуральном числе.
	var numb, max, dec int64
	fmt.Println("Введите натуральное число ")
	fmt.Scan(&numb)
	dec, max = 1, 0
	for dec = 1; dec <= numb; dec *= 10 {
		if (numb/dec)%10 > max {
			max = (numb / dec) % 10
		}
	}
	fmt.Println(max)
}

func One_hundred_twenty_one() { //Дано натуральное число. Найдите количество четных цифр
	var numb, ans, dec int64
	fmt.Println("Введите натуральное число ")
	fmt.Scan(&numb)
	dec, ans = 1, 0
	for dec = 1; dec <= numb; dec *= 10 {
		if ((numb/dec)%10)%2 == 0 {
			ans += 1
		}
	}
	fmt.Println(ans)
}

func One_hundred_twenty_two() { // В данном натуральном числе найдите количество цифр, которые больше 3, но меньше 8.
	var numb, ans, dec int64
	fmt.Println("Введите натуральное число ")
	fmt.Scan(&numb)
	dec, ans = 1, 0
	for dec = 1; dec <= numb; dec *= 10 {
		if ((numb/dec)%10) > 3 && ((numb/dec)%10) < 8 {
			ans += 1
		}
	}
	fmt.Println(ans)
}

func One_hundred_twenty_tree() { // Для данного натурального числа найдите число, цифры которого записаны в обратном порядке.
	var numb, ans, dec, decmax int64
	fmt.Println("Введите натуральное число ")
	fmt.Scan(&numb)
	dec, ans = 1, 0
	for decmax = 1; decmax <= numb; decmax *= 10 {
	}
	for dec = 1; dec <= numb; dec *= 10 {
		decmax /= 10
		ans += decmax * ((numb / dec) % 10)
	}
	fmt.Println(ans)
}

func One_hundred_twenty_four() { // Найдите n-ое число Фибоначчи
	var fn, fn_1, i, n int64 = 1, 0, 0, 0
	fmt.Println("Введите номер числа")
	fmt.Scan(&n)
	for i = 2; i <= n; i++ {
		fn, fn_1 = fn+fn_1, fn
	}
	fmt.Println(fn)
}

func fx1(x float64) float64 { // функция для One_hundred_twenty_five()
	return math.Pow(x, 2) - math.Sin(x)
}

func One_hundred_twenty_five() { //вычислите значения функции f(x)=x**2−sinx на отрезке [a;b] с шагом h. Результат представить в виде таблицы.
	var a, b, h, i float64
	fmt.Println("Введите отрезок и шаг ")
	fmt.Scan(&a, &b, &h)
	fmt.Println("x  ", "  f(x)")
	for i = 0; i <= (b-a)/h; i++ {
		fmt.Println(a+i*h, "    ", fx1(a+i*h))
	}
}

func One_hundred_twenty_six() { //Найдите все делители данного натурального числа
	var numb, i int64
	fmt.Println("Введите натуральное число ")
	fmt.Scan(&numb)
	for i = 1; i <= numb/2; i++ {
		if numb%i == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(numb)
}

func One_hundred_twenty_seven() { // Определите, является ли данное число простым
	var numb, i int64
	var flag bool = false
	fmt.Println("Введите натуральное число ")
	fmt.Scan(&numb)
	for i = 2; i <= numb/2; i++ {
		if numb%i == 0 {
			flag = true
		}
	}
	if flag {
		fmt.Println("Данное число не простое")
	} else {
		fmt.Println("Данное число простое")
	}
}

func One_hundred_twenty_eight() { // Два числа называются дружественными, если каждое из них равно сумме всех делителей второго не считая самого этого числа. Найдите все пары дружественных чисел на отрезке [a;b].
	var a, b, i, j, del, delrev int64
	fmt.Println("Введите отрезок")
	fmt.Scan(&a, &b)
	for j = a; j <= b; j++ {
		del = 0
		for i = 1; i <= j/2; i++ {
			if j%i == 0 {
				del += i
			}
		}
		if del >= a && del <= b && del > j {
			delrev = 0
			for i = 1; i <= del/2; i++ {
				if del%i == 0 {
					delrev += i
				}
			}
			if delrev == j {
				fmt.Println(j, "   ", del)
			}
		}
	}
}

func One_hundred_twenty_nine() { // Натуральное число называется совершенным, если оно равно сумме всех своих делителей, не равных самому числу. Найдите все совершенные числа, меньшие данного натурального числа n.
	var n, i, j, del int64
	fmt.Println("Введите отрезок")
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		del = 0
		for j = 1; j <= i/2; j++ {
			if i%j == 0 {
				del += j
			}
		}
		if i == del {
			fmt.Println(i)
		}
	}
}

func One_hundred_thirty() { //назовем автобусный билет несчастливым, если сумма цифр его шестизначного номера делится на 13. Могут ли два идущих подряд билета оказаться несчастливыми?

}

func One_hundred_thirty_one() { //Найдите n пар простых чисел, которые отличаются друг от друга на 2.
	var numb, i, j, n int64
	var flag1, flag2 bool = true, true
	fmt.Println("Введите натуральное число ")
	fmt.Scan(&n)
	j = 0
	for numb = 1; numb < math.MaxInt64-3 && j < n; numb++ {
		flag1, flag2 = true, true
		for i = 2; i <= numb/2; i++ {
			if numb%i == 0 {
				flag1 = false
			}
		}
		if flag1 {
			for i = 2; i <= (numb+2)/2; i++ {
				if (numb+2)%i == 0 {
					flag2 = false
				}
			}
		}
		if flag1 && flag2 {
			fmt.Println(numb, "  ", numb+2)
			j += 1
		}
	}
}

func One_hundred_thirty_two() { // Найдите все натуральные числа, не превосходящие 10000, сумма цифр каждого из которых в некоторой степени равна самому числу.
	var i, j, sum int64
	for i = 1; i <= 10000; i++ {
		sum = 0
		for j = 1; j <= i; j *= 10 {
			sum += (i / j) % 10
		}
		for j = 1; math.Pow(float64(sum), float64(j)) < float64(i) && sum != 1; j++ {
		}
		if math.Pow(float64(sum), float64(j)) == float64(i) {
			fmt.Println(i)
		}
	}
}

func One_hundred_thirty_tree() { // Дано число k. Определите, существует ли такое число n, что 1+2+3+...+n=k
	var sum, n, k int64 = 0, 0, 1
	fmt.Println("Введите проверяемое число")
	fmt.Scan(&n)
	for k = 1; sum < n; k++ {
		sum += k
	}
	if sum == n {
		fmt.Println("существует ", k-1)
	} else {
		fmt.Println("не существует ")
	}
}

func f2(xc int64, yc int64, xn int64, yn int64) float64 {
	return math.Pow(math.Pow(float64(xc-xn), 2)+math.Pow(float64(yc-yn), 2), 0.5)
}

func One_hundred_thirty_four() { // Найдите, сколько точек с целочисленными координатами попадает в круг радиуса r  с центром в точке (x,y).
	var xc, yc, i, j, ans, z int64
	var r float64
	fmt.Println("Введите параметры круга: ") // вообщем выделяем весь квадрат в котором находится окружность и вычетаем угловые точки которые не лежат внутри
	fmt.Scan(&xc, &yc, &r)
	i = int64(r) + xc
	j = int64(r) + yc
	ans = 0
	for f2(xc, yc, i, j) > r {
		for z = 0; f2(xc, yc, i, j-z-1) > r; z++ {
		}
		ans += z*2 + 1
		i -= 1
		j -= 1
		fmt.Println(z)
	}
	fmt.Println("Колличество точек: ", (int64(r)*2+1)*(int64(r)*2+1)-ans*4)
}

func One_hundred_thirty_five() { //Выведите случайную серию чисел из 0 и 1 такую, что сумма чисел в ней больше 10.
	var i int64 = 0
	rand.Seed(time.Now().UnixNano())
	n := rand.Int63()
	for i <= (n%89 + 11) {
		n = rand.Int63()
		if n%2 == 1 {
			i++
		}
		fmt.Print(n % 2)
	}
	fmt.Print("\n")
}

func One_hundred_thirty_six() { // Дано n кирпичей. Вы и компьютер ходите поочередно. за ход можно взять 1, 2 или 3 кирпича. Проиграл тот, кому нечего брать. Реализуйте игру с компьютером. Компьютер ходит случайно (без анализа выигрышной стратегии), однако если у него есть ход, который является последним для его выигрыша, то он его совершает.
	var n, coin, r int64
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Введите колличество кирпечей (в случае если введен 0 или отрицательное число колличество кирпечей случайное)")
	fmt.Scan(&n)
	if n <= 0 {
		n = rand.Int63() % 31
	}
	fmt.Println("Введите очередность хода 1 или 2 (в случае если введен 0 или отрицательное число или превышающее число сторона рандомится)")
	fmt.Scan(&coin)
	if 0 >= coin && coin >= 2 {
		coin = rand.Int63()%2 + 1
	}
	if coin == 1 {
		for n > 0 {
			fmt.Println("Осталось ", n, " камней")
			fmt.Println("Введите колличество камней 1, 2 или 3")
			fmt.Scan(&r)
			for r != 1 && r != 2 && r != 3 && r > n {
				fmt.Println("Нет введите 1, 2 или 3 или число не превышающее оставшиеся камни.")
				fmt.Scan(&r)
			}
			n -= r
			if n == 0 {
				fmt.Println("Вы победили!")
				return
			}
			if n != 1 && n != 2 && n != 3 {
				r = rand.Int63()%3 + 1
				fmt.Println("Компьютер берет ", r, " камней.")
				n -= r
			} else {
				fmt.Println("Компьютер берет ", n, " камней и побеждает!.")
				return
			}
		}
	} else {
		for n > 0 {
			if n != 1 && n != 2 && n != 3 {
				r = rand.Int63()%3 + 1
				fmt.Println("Компьютер берет ", r, " камней.")
				n -= r
			} else {
				fmt.Println("Компьютер берет ", n, " камней и побеждает!.")
				return
			}
			fmt.Println("Осталось ", n, " камней")
			fmt.Println("Введите колличество камней 1, 2 или 3")
			fmt.Scan(&r)
			for r != 1 && r != 2 && r != 3 && r > n {
				fmt.Println("Нет введите 1, 2 или 3 или число не превышающее оставшиеся камни.")
				fmt.Scan(&r)
			}
			n -= r
			if n == 0 {
				fmt.Println("Вы победили!")
				return
			}
		}
	}
}

func One_hundred_thirty_seven() {

}

func One_hundred_thirty_eight() { //Сгенерировать случайную серию из 15 чисел, в которой ровно 3 единицы, остальные нули.
	var n1, n2, n3, i int64
	rand.Seed(time.Now().UnixNano())
	n1 = rand.Int63()%15 + 1
	n2 = rand.Int63()%15 + 1
	n3 = rand.Int63()%15 + 1
	for n2 == n1 || n3 == n2 || n3 == n1 {
		n1 = rand.Int63()%15 + 1
		n2 = rand.Int63()%15 + 1
		n3 = rand.Int63()%15 + 1
	}
	for i = 1; i <= 15; i++ {
		if i == n1 || i == n2 || i == n3 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Print("\n")
}

func One_hundred_thirty_nine() { // Сгенерируйте серию из 10 случайных чисел от 1 до 3 и найдите: а) на сколько количество двоек больше/меньше количества троек, б) количество троек, стоящих на четных местах, в)количество двоек среди первых пяти чисел серии.
	var n, numb_a, numb_b, numb_c, i int64 = 0, 0, 0, 0, 0
	rand.Seed(time.Now().UnixNano())
	for i = 1; i <= 10; i++ {
		n = rand.Int63()%3 + 1
		if n == 2 {
			numb_a += 1
		} else if n == 3 {
			numb_a -= 1
		}
		if i%2 == 0 && n == 3 {
			numb_b += 1
		}
		if i < 6 && n == 2 {
			numb_c += 1
		}
		fmt.Print(n)
	}
	fmt.Print("\n")
	if numb_a > 0 {
		fmt.Println("Колличество двоек больше колличества троек на ", numb_a)
	} else if numb_a == 0 {
		fmt.Println("Колличество двоек равно колличеству троек ")
	} else {
		fmt.Println("Колличество двоек меньше колличества троек на ", -numb_a)
	}
	fmt.Println("Колличество троек на четных местах = ", numb_b)
	fmt.Println("Колличество двоек среди первых пяти чисел серии = ", numb_c)
}

func One_hundred_forty() { // Сгенерировать 20 серий из 0, 1, 2 таких, что сумма чисел в каждой серии равна 12. Найти количество единиц в каждой серии, количество двоек в каждой серии, длину каждой серии, среднее количество двоек в сериях, среднюю длину серий и наибольшее количество ненулевых чисел в сериях.
	numbers_one := [20]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	numbers_two := [20]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	long_serias := [20]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	rand.Seed(time.Now().UnixNano())
	var sum, num, i, max_nonz int64
	var ave_long, ave_one, ave_two float64 = 0, 0, 0
	for i = 1; i <= 20; i++ {
		sum = 0
		for sum < 12 {
			if 12-sum == 1 {
				num = rand.Int63() % 2
			} else {
				num = rand.Int63() % 3
			}
			sum += num
			long_serias[i-1] += 1
			if num == 1 {
				numbers_one[i-1] += 1
			} else if num == 2 {
				numbers_two[i-1] += 1
			}
		}
	}
	max_nonz = 0
	fmt.Println("№ серии     единицы    двойки    длинна")
	for i = 0; i <= 19; i++ {
		fmt.Println("  ", i+1, "         ", numbers_one[i], "        ", numbers_two[i], "        ", long_serias[i])
		ave_long += float64(long_serias[i])
		ave_one += float64(numbers_one[i])
		ave_two += float64(numbers_two[i])
		if max_nonz < numbers_one[i]+numbers_two[i] {
			max_nonz = numbers_one[i] + numbers_two[i]
		}
	}
	ave_long /= 20
	ave_one /= 20
	ave_two /= 20
	fmt.Println("Средняя длинна ", ave_long)
	fmt.Println("Среднее колличество единиц", ave_one)
	fmt.Println("Среднее колличество двоек", ave_two)
}

func One_hundred_forty_one() { //Сгенерируйте серию случайных чисел из 0, 1, 2 так, чтобы количество двоек было равно количеству нулей.
	var n_two, n_zero, n_numb, i, numb int64
	rand.Seed(time.Now().UnixNano())
	n_zero = 0
	n_two = 0
	n_numb = rand.Int63()%10 + 10
	fmt.Print("\n")
	for i = 1; n_two != n_zero || i <= n_numb; i++ {
		numb = rand.Int63() % 3
		if numb == 0 {
			n_zero += 1
		} else if numb == 2 {
			n_two += 1
		}
		fmt.Print(numb, " ")
	}
	fmt.Print("\n")
}

func One_hundred_forty_two() { //Компьютер загадывает число от 1 до 100. У пользователя три попытки отгадать. После каждой неудачной попытки компьютер сообщает меньше или больше загаданное число.
	var numb, i, gad int64
	rand.Seed(time.Now().UnixNano())
	numb = rand.Int63()%100 + 1
	fmt.Println("Угадайте число от одного до ста")
	fmt.Scan(&gad)
	for i = 1; i <= 2; i++ {
		if gad < i {
			fmt.Println("Указанное число меньше")
		} else if gad > i {
			fmt.Println("Указанное число больше")
		} else {
			fmt.Println("Вы угадали")
			return
		}
		fmt.Println("Введите новое")
		fmt.Scan(&gad)
	}
	fmt.Println("Вы проиграли указанное число: ", numb)
}

func One_hundred_forty_tree() { //Вывести 3 случайных числа от 0 до 100 без повторений.
	var n1, n2, n3 int64
	n1, n2, n3 = 1, 1, 1
	rand.Seed(time.Now().UnixNano())
	for n1 == n2 || n2 == n3 || n3 == n1 {
		n1 = rand.Int63() % 101
		n2 = rand.Int63() % 101
		n3 = rand.Int63() % 101
	}
	fmt.Println("Числа  ", n1, "  ", n2, "  ", n3)
}

func One_hundred_forty_five() { //Пользователь вводит англ. букву, вывести следующие три по алфавиту. Если алфавит закончился, то вывести циклично с начала алфавита, то есть если z, то a b c.  Вывод только маленьких букв. Учесть, что пользователь может ввести заглавную
	var c string
	fmt.Println("Введите маленькую букву")
	fmt.Scan(&c)
	s := []rune(c)
	if 'A' <= s[0] && s[0] <= 'Z' {
		fmt.Println("Вы ввели заглавную букву")
		return
	}
	if s[0] == 'x' {
		fmt.Println(string(s[0] + 1))
		fmt.Println(string(s[0] + 2))
		fmt.Println("a")
	} else if s[0] == 'y' {
		fmt.Println(string(s[0] + 1))
		fmt.Println("a")
		fmt.Println("b")
	} else if s[0] == 'z' {
		fmt.Println("a")
		fmt.Println("b")
		fmt.Println("c")
	} else {
		fmt.Println(string(s[0] + 1))
		fmt.Println(string(s[0] + 2))
		fmt.Println(string(s[0] + 3))
	}
}

func One_hundred_forty_six() { //Вывести англ. алфавит по 5 букв в строке.
	for i := 0; 'a'+i <= 'z'; i++ {
		if i%5 == 0 {
			fmt.Print("\n")
		}
		fmt.Print(string('a'+i), " ")
	}
	fmt.Print("\n")
}

func main() {
	//One()
	//Two()
	//Tree()
	//Four()
	//Five()
	//Six()
	//Seven()
	//Eight(2,3)
	//Nine(-2)
	//Ten(1.7)
	//Ten(3)
	//Eleven(-2.34)
	//Twelve(3.6)
	//Thirteen(0.1, 0.2, 1)
	//Fourteen()
	//Fifteen()
	//Sixteen(3,3,3)
	//Seventeen(2,2,1)
	//Eighteen()
	//Nineteen()
	//Twenty()
	//Twenty_one(1,1)
	//Twenty_two(14)
	//Twenty_tree(3, 300, 2, 250)
	//Twenty_four(10, 10, 70)
	//Twenty_five(0,6,1)

	/*var a, b int64
	a, b = 2 , 3
	Twenty_six(&a, &b)
	fmt.Println(a,"  ",b)*/

	/*var a, b, c float64
	a, b, c = 0, 2, 5
	Twenty_seven(&a, &b, &c)
	fmt.Println(a,"  ",b,"  ",c)*/

	//Thirty(2)
	//Thirty_two(15312)
	//Thirty_tree(5, 4)

	/*var a float64 = 2
	Thirty_four(&a)
	fmt.Println(a)*/

	//Thirty_five(8)
	//Thirty_six(1)
	//Thirty_seven(3,3)
	//Thrity_nine(3,2)

	/*var a float64 = 11
	Forty(&a)
	fmt.Println(a)*/

	/*var a float64 = -99
	Forty_one(&a)
	fmt.Println(a)*/

	/*var a float64 = 30
	Forty_two(&a)
	fmt.Println(a)*/

	//Forty_tree()
	//Forty_seven()
	//Forty_eight()
	//Forty_nine()
	//Fifty()
	//Fifty_two()
	//Fifty_tree()
	//Fifty_five()
	//Fifty_six()
	//Fifty_seven()
	//Fifty_eight()
	//Fifty_nine()
	//Sixty()
	//Sixty_one()
	//Sixty_two()
	//Sixty_tree()
	//Sixty_four()
	//Sixty_five()
	//Sixty_six()
	//Sixty_seven()
	//Sixty_eight()
	//Sixty_nine()
	//Seventy()
	//Seventy_one()

	//Seventy_five()
	//Seventy_six()
	//Seventy_seven()
	//Seventy_eight()
	//Seventy_nine()
	//Eighty_eighty_two()
	//Eighty_tree()
	//Eighty_four()
	//Eighty_seven()
	//Eighty_eight()
	//Eighty_nine()
	//Ninety
	//Ninety_one()
	//Ninety_two()
	//Ninety_tree()
	//Ninety_four()
	//Ninety_seven()
	//Ninety_eight()
	//Ninety_nine()
	//Hundred()
	//One_hundred_one()
	//One_hundred_two()
	//One_hundred_tree()
	//One_hundred_four()
	//One_hundred_five()
	//One_hundred_six()
	//One_hundred_seven()
	//One_hundred_eight()
	//One_hundred_nine()
	//One_hundred_ten()
	//One_hundred_eleven()
	//One_hundred_twelve()
	//One_hundred_thirteen()
	//One_hundred_fourteen()
	//One_hundred_fifteen()
	//One_hundred_sixteen()
	//One_hundred_eighteen()
	//One_hundred_nineteen()
	//One_hundred_twenty()
	//One_hundred_twenty_one()
	//One_hundred_twenty_two()
	//One_hundred_twenty_tree()
	//One_hundred_twenty_four()
	//One_hundred_twenty_five()
	//One_hundred_twenty_six()
	//One_hundred_twenty_seven()
	//One_hundred_twenty_eight()
	//One_hundred_twenty_nine()
	//One_hundred_thirty_one()
	//One_hundred_thirty_two()
	//One_hundred_thirty_tree()
	//One_hundred_thirty_four()
	//One_hundred_thirty_five()
	//One_hundred_thirty_six()
	//One_hundred_thirty_eight()
	//One_hundred_thirty_nine()
	//One_hundred_forty()
	//One_hundred_forty_one()
	//One_hundred_forty_two()
	//One_hundred_forty_tree()
	//One_hundred_forty_five()
	One_hundred_forty_six()
}
