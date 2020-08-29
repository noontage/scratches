def solution(n)
    ary = []
    t = 0
    s = n.to_s(2)

    s.each_char do |c|
        if c == '1'
            unless t.zero?
                ary.push(t)
                t = 0
            end
        else
            t+=1
        end
    end

    return 0 if ary.empty?
    ary.max
end

p solution(0b11111111111111111)