def soltion(a)
    k = nil
    cnt = 0
    a.sort.each do |e|
        if k != e
            k = e
            cnt+=1
        end
    end
    cnt
end

p soltion([1,2,3,1,2,1,2,3,5,1,2,6,3,5])