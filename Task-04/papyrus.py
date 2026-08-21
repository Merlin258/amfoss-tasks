for _ in range(int(input())):
    n,c=map(int,input().split())
    l1=list(map(int,input().split()))
    l2=list(map(int,input().split()))
    ans=0
    flag1=True
    for i in range(n):
        if l2[i]>l1[i]:
            flag1=False
        ans+=l1[i]-l2[i]
    if flag1:
        print(ans)
    else:
        l1.sort()
        l2.sort()
        flag2=True
        for i in range(n):
            if l2[i]>l1[i]:
                flag2=False
        if flag2:
            print(c+ans)
        else:
            print(-1)