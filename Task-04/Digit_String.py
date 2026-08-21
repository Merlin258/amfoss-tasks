t=int(input())
for _ in range(t):
    s=input()
    pre=0
    suf=0
    for i in s:
        if i in "13":
            suf+=1
    x=pre+suf
    for i in s:
        if i=="2":
            pre+=1
        if i in "13":
            suf-=1
        x=max(x,pre+suf)
    print(len(s)-x)