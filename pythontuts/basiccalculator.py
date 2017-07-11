#import sys
#file = sys.stdin
file = open( r".\data\basiccalculator.txt" )
data = file.read().strip().split()

ops = ['+', '-', '*', '/', '//']
exp = [data[0] + op + data[1]    for op in ops]
ans = ['{:.2f}'.format(eval(ex)) for ex in exp]

print('\n'.join(ans))


# Others

#print('%.2f'%(a+b))

#a, b = float(input()), float(input())
#print(*('{:.2f}'.format(x) for x in (a + b, a - b, a * b, a / b, a // b)), sep='\n')

#sys.stdout.write('{0:.2f}\n{1:.2f}\n{2:.2f}\n{3:.2f}\n{4:.2f}\n'.format(*output))
