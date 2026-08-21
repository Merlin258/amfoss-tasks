for _ in range(int(input())):
    n=int(input())
    s=list(map(int,input().split()))
    for i in range(n-1):
        if s[i]>s[i+1]:
            s[i],s[i+1]=s[i+1],s[i+1]+s[i]
    print(s[-1])