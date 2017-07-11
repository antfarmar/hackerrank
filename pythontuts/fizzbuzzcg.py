#score = 0.58
#i=0
#while i<100:i+=1;print("".join("BzuzzizF"[::2*j]for j in(-1,1)if 1>i%(4+j))or i)

#score = 0.68
#for x in range(1,101):print('Fizz'*(x%3==0)+'Buzz'*(x%5==0)or x)

#score = 0.685
i=0
while i<100:i+=1;print('Fizz'*(i%3==0)+'Buzz'*(i%5==0)or i)
