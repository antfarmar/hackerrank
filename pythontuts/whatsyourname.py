#import sys
#file = sys.stdin
file = open( r".\data\whatsyourname.txt" )
data = file.read().strip().split()

msg = "Hello {0} {1}! You just delved into python."

print(data)
print(msg.format(*data))
