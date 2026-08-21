from math import sqrt
t=int(input())
def next_prime(n):
    n+=1
    while True:
        flag=True
        if n%2==0 or n%3==0:
            flag=False
        for i in range(6,int(sqrt(n))+1,6):
            if n%(i+1)==0 or n%(i-1)==0:
                flag=False
        if flag:
            return n
        n+=1
for _ in range(t):
    prime_window=[3,5]
    n=int(input())
    print("6",end="")
    for x in range(n-1):
        print(" ",end="")
        print(prime_window[0]*prime_window[1],end="")
        prime_window[0],prime_window[1]=prime_window[1],next_prime(prime_window[1])
    print()